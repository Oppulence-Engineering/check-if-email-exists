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
	"time"
)

// config captures all runtime knobs so the script stays easily testable and DRY.
type config struct {
	baseURL   string
	email     string
	secret    string
	timeout   time.Duration
	pollDelay time.Duration
}

func main() {
	cfg := loadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	if err := waitForService(ctx, cfg); err != nil {
		fail("service did not become ready", err)
	}

	respBody, err := verifyEmail(ctx, cfg)
	if err != nil {
		fail("email verification failed", err)
	}

	fmt.Println("Backend response:")
	fmt.Println(respBody)
}

// loadConfig builds the runtime configuration from environment variables with sensible defaults.
func loadConfig() config {
	return config{
		baseURL:   envOrDefault("CHECK_EMAIL_BASE_URL", "http://localhost:8080"),
		email:     envOrDefault("CHECK_EMAIL_TEST_EMAIL", "demo@example.com"),
		secret:    os.Getenv("CHECK_EMAIL_SECRET"),
		timeout:   parseDurationOrDefault("CHECK_EMAIL_TIMEOUT", 2*time.Minute),
		pollDelay: parseDurationOrDefault("CHECK_EMAIL_POLL_DELAY", 3*time.Second),
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

// verifyEmail sends the sample verification request and returns the pretty-printed JSON body.
func verifyEmail(ctx context.Context, cfg config) (string, error) {
	payload, err := json.Marshal(map[string]string{
		"to_email": cfg.email,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/v1/check_email", cfg.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if cfg.secret != "" {
		req.Header.Set("x-reacher-secret", cfg.secret)
	}

	client := http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("unexpected status %d: %s", res.StatusCode, string(body))
	}

	return prettyJSON(body)
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
