package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// config captures all runtime knobs so the script stays easily testable and DRY.
type runMode string

const (
	modeSingle runMode = "single"
	modeBulk   runMode = "bulk"
	modeBoth   runMode = "both"
)

type config struct {
	baseURL         string
	email           string
	secret          string
	timeout         time.Duration
	pollDelay       time.Duration
	mode            runMode
	bulkEmails      []string
	bulkPollDelay   time.Duration
	bulkTimeout     time.Duration
	bulkResultLimit int
}

func main() {
	cfg := loadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	if err := waitForService(ctx, cfg); err != nil {
		fail("service did not become ready", err)
	}

	if cfg.mode == modeSingle || cfg.mode == modeBoth {
		if err := runSingleShot(ctx, cfg); err != nil {
			fail("single-shot verification failed", err)
		}
	}
	if cfg.mode == modeBulk || cfg.mode == modeBoth {
		if err := runBulkWorkflow(ctx, cfg); err != nil {
			fail("bulk verification failed", err)
		}
	}
}

// loadConfig builds the runtime configuration from environment variables with sensible defaults.
func loadConfig() config {
	mode := runMode(strings.ToLower(envOrDefault("CHECK_EMAIL_MODE", string(modeBoth))))
	switch mode {
	case modeSingle, modeBulk, modeBoth:
	default:
		mode = modeBoth
	}

	return config{
		baseURL:         envOrDefault("CHECK_EMAIL_BASE_URL", "http://localhost:8080"),
		email:           envOrDefault("CHECK_EMAIL_TEST_EMAIL", "demo@example.com"),
		secret:          os.Getenv("CHECK_EMAIL_SECRET"),
		timeout:         parseDurationOrDefault("CHECK_EMAIL_TIMEOUT", 2*time.Minute),
		pollDelay:       parseDurationOrDefault("CHECK_EMAIL_POLL_DELAY", 3*time.Second),
		mode:            mode,
		bulkEmails:      parseEmailList(envOrDefault("CHECK_EMAIL_BULK_EMAILS", "demo@example.com,yoanyombapro@gmail.com")),
		bulkPollDelay:   parseDurationOrDefault("CHECK_EMAIL_BULK_POLL_DELAY", 5*time.Second),
		bulkTimeout:     parseDurationOrDefault("CHECK_EMAIL_BULK_TIMEOUT", 5*time.Minute),
		bulkResultLimit: parseIntOrDefault("CHECK_EMAIL_BULK_RESULT_LIMIT", 50),
	}
}

// waitForService polls the /version endpoint until the backend reports ready or the context expires.
func waitForService(ctx context.Context, cfg config) error {
	versionURL := fmt.Sprintf("%s/version", cfg.baseURL)
	client := http.Client{Timeout: 5 * time.Second}

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, versionURL, nil)
		if err != nil {
			return err
		}

		res, err := client.Do(req)
		if err == nil && res.StatusCode == http.StatusOK {
			_ = res.Body.Close()
			return nil
		}
		if res != nil {
			_ = res.Body.Close()
		}

		select {
		case <-ctx.Done():
			return errors.New("timeout waiting for backend readiness")
		case <-time.After(cfg.pollDelay):
		}
	}
}

// runSingleShot hits /v1/check_email and prints the response.
func runSingleShot(ctx context.Context, cfg config) error {
	payload, err := json.Marshal(map[string]string{
		"to_email": cfg.email,
	})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/check_email", cfg.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if cfg.secret != "" {
		req.Header.Set("x-reacher-secret", cfg.secret)
	}

	client := http.Client{Timeout: 30 * time.Second}
	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	printed, err := prettyJSON(body)
	if err != nil {
		return err
	}
	fmt.Println("Single-shot response:")
	fmt.Println(printed)
	fmt.Printf("Single-shot duration: %s\n\n", time.Since(start).Truncate(time.Millisecond))
	return nil
}

func runBulkWorkflow(parent context.Context, cfg config) error {
	if len(cfg.bulkEmails) == 0 {
		return errors.New("CHECK_EMAIL_BULK_EMAILS is empty")
	}

	fmt.Printf("Starting bulk workflow for %d addresses: %s\n", len(cfg.bulkEmails), strings.Join(cfg.bulkEmails, ", "))

	bulkCtx, cancel := context.WithTimeout(parent, cfg.bulkTimeout)
	defer cancel()

	client := &http.Client{Timeout: 30 * time.Second}
	payload, err := json.Marshal(map[string]any{"input": cfg.bulkEmails})
	if err != nil {
		return err
	}

	createReq, err := http.NewRequestWithContext(bulkCtx, http.MethodPost, fmt.Sprintf("%s/v1/bulk", cfg.baseURL), bytes.NewReader(payload))
	if err != nil {
		return err
	}
	createReq.Header.Set("Content-Type", "application/json")
	applySecret(createReq, cfg)

	createRes, err := client.Do(createReq)
	if err != nil {
		return err
	}
	defer createRes.Body.Close()
	if createRes.StatusCode >= 400 {
		body, _ := io.ReadAll(createRes.Body)
		return fmt.Errorf("bulk create failed (%d): %s", createRes.StatusCode, string(body))
	}

	var created struct {
		JobID int `json:"job_id"`
	}
	if err := json.NewDecoder(createRes.Body).Decode(&created); err != nil {
		return err
	}
	fmt.Printf("Created bulk job %d\n", created.JobID)

	progress, err := pollBulkJob(bulkCtx, client, cfg, created.JobID)
	if err != nil {
		return err
	}

	if err := fetchBulkResults(bulkCtx, client, cfg, created.JobID); err != nil {
		return err
	}

	fmt.Printf("Bulk job %d summary: safe=%d risky=%d invalid=%d unknown=%d processed=%d/%d\n\n",
		progress.JobID,
		progress.Summary.TotalSafe,
		progress.Summary.TotalRisky,
		progress.Summary.TotalInvalid,
		progress.Summary.TotalUnknown,
		progress.TotalProcessed,
		progress.TotalRecords,
	)
	return nil
}

type bulkProgressResponse struct {
	JobID          int    `json:"job_id"`
	TotalRecords   int    `json:"total_records"`
	TotalProcessed int    `json:"total_processed"`
	JobStatus      string `json:"job_status"`
	Summary        struct {
		TotalSafe    int `json:"total_safe"`
		TotalRisky   int `json:"total_risky"`
		TotalInvalid int `json:"total_invalid"`
		TotalUnknown int `json:"total_unknown"`
	} `json:"summary"`
}

func pollBulkJob(ctx context.Context, client *http.Client, cfg config, jobID int) (bulkProgressResponse, error) {
	url := fmt.Sprintf("%s/v1/bulk/%d", cfg.baseURL, jobID)
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return bulkProgressResponse{}, err
		}
		applySecret(req, cfg)

		res, err := client.Do(req)
		if err != nil {
			return bulkProgressResponse{}, err
		}

		var payload bulkProgressResponse
		decodeErr := json.NewDecoder(res.Body).Decode(&payload)
		res.Body.Close()
		if decodeErr != nil {
			return bulkProgressResponse{}, decodeErr
		}
		if res.StatusCode >= 400 {
			return bulkProgressResponse{}, fmt.Errorf("bulk progress error (%d): job %d", res.StatusCode, jobID)
		}

		fmt.Printf("Job %d status=%s processed=%d/%d\r", jobID, payload.JobStatus, payload.TotalProcessed, payload.TotalRecords)
		if strings.EqualFold(payload.JobStatus, "completed") {
			fmt.Println()
			return payload, nil
		}

		select {
		case <-ctx.Done():
			return bulkProgressResponse{}, ctx.Err()
		case <-time.After(cfg.bulkPollDelay):
		}
	}
}

func fetchBulkResults(ctx context.Context, client *http.Client, cfg config, jobID int) error {
	url := fmt.Sprintf("%s/v1/bulk/%d/results?format=json", cfg.baseURL, jobID)
	if cfg.bulkResultLimit > 0 {
		url = fmt.Sprintf("%s&limit=%d", url, cfg.bulkResultLimit)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	applySecret(req, cfg)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("bulk results error (%d): %s", res.StatusCode, string(body))
	}

	var parsed struct {
		Results []json.RawMessage `json:"results"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return err
	}

	fmt.Printf("Fetched %d bulk results (limit=%d)\n", len(parsed.Results), cfg.bulkResultLimit)
	maxDisplay := parsed.Results
	if cfg.bulkResultLimit > 0 && len(maxDisplay) > cfg.bulkResultLimit {
		maxDisplay = parsed.Results[:cfg.bulkResultLimit]
	}
	display := len(maxDisplay)
	if display > 3 {
		display = 3
	}
	for i := 0; i < display; i++ {
		rendered, err := prettyJSON(parsed.Results[i])
		if err != nil {
			return err
		}
		fmt.Printf("Result #%d:\n%s\n", i+1, rendered)
	}
	return nil
}

// prettyJSON indents JSON responses for readability while keeping raw output if unmarshalling fails.
func prettyJSON(raw []byte) (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		// Return the raw payload if we cannot pretty-print (still useful for debugging).
		return string(raw), nil
	}
	return buf.String(), nil
}

func envOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func parseDurationOrDefault(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return d
}

func fail(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}

func parseEmailList(raw string) []string {
	parts := strings.Split(raw, ",")
	emails := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			emails = append(emails, trimmed)
		}
	}
	return emails
}

func parseIntOrDefault(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return fallback
}

func applySecret(req *http.Request, cfg config) {
	if cfg.secret != "" {
		req.Header.Set("x-reacher-secret", cfg.secret)
	}
}
