# RFC-001: Reacher Platform Evolution

| Field | Value |
|-------|-------|
| **RFC Number** | 001 |
| **Title** | Reacher Platform Evolution: From Email Verification Engine to Full-Stack Prospecting Platform |
| **Author** | Oppulence Engineering |
| **Status** | Draft |
| **Created** | 2024-12-03 |
| **Updated** | 2024-12-03 |

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Background and Motivation](#2-background-and-motivation)
3. [Goals and Non-Goals](#3-goals-and-non-goals)
4. [Feature Specifications](#4-feature-specifications)
   - [4.1 Email List Management Dashboard](#41-email-list-management-dashboard)
   - [4.2 Contact Enrichment Module](#42-contact-enrichment-module)
   - [4.3 Usage Analytics & Billing Dashboard](#43-usage-analytics--billing-dashboard)
   - [4.4 Native CRM Integrations](#44-native-crm-integrations)
   - [4.5 Compliance & Audit Logging](#45-compliance--audit-logging)
   - [4.6 Spam Trap & Honeypot Detection](#46-spam-trap--honeypot-detection)
   - [4.7 Scheduled Re-verification](#47-scheduled-re-verification)
   - [4.8 Bulk Domain Verification](#48-bulk-domain-verification)
   - [4.9 Sender Reputation Scoring](#49-sender-reputation-scoring)
   - [4.10 Advanced Segmentation & Scoring Engine](#410-advanced-segmentation--scoring-engine)
5. [System Architecture](#5-system-architecture)
6. [Database Schema Changes](#6-database-schema-changes)
7. [API Design](#7-api-design)
8. [Security Considerations](#8-security-considerations)
9. [Implementation Roadmap](#9-implementation-roadmap)
10. [Success Metrics](#10-success-metrics)
11. [Open Questions](#11-open-questions)
12. [Appendices](#12-appendices)

---

## 1. Executive Summary

This RFC proposes a comprehensive evolution of the Reacher email verification platform from a developer-focused API service into a full-stack B2B prospecting and email deliverability platform. The proposal encompasses 10 major features that address critical gaps in the current offering and position Reacher to compete effectively with established players like Hunter.io, ZeroBounce, and Apollo.io.

### Current State

Reacher is a high-accuracy, open-source email verification engine with:
- Robust SMTP-based verification with provider-specific strategies
- Async worker queue system (RabbitMQ) for scalable bulk processing
- REST API with webhook support
- Self-hosted deployment option

### Proposed End State

A comprehensive email intelligence platform offering:
- Web-based dashboard for non-technical users
- Contact enrichment (company, title, social profiles)
- CRM-native integrations (Salesforce, HubSpot, Pipedrive)
- Compliance and audit logging for enterprise customers
- Sender reputation and deliverability scoring
- Scheduled monitoring and re-verification

### Strategic Impact

| Metric | Current | Projected (12 months) |
|--------|---------|----------------------|
| Target Market | Developers only | SMB + Mid-Market + Enterprise |
| Use Case | One-time list cleaning | Ongoing prospecting platform |
| Revenue Model | Open source / Self-hosted | Freemium SaaS + Enterprise |
| Stickiness | Low (single-use) | High (daily workflow tool) |
| Competitive Position | Verification-only alternative | Full-stack competitor to Hunter/Apollo |

---

## 2. Background and Motivation

### 2.1 Market Context

The email verification and B2B data market is valued at $3.2B (2024) with 12% CAGR. Key players include:

| Competitor | Verification | Enrichment | Dashboard | CRM Integration | Pricing |
|------------|--------------|------------|-----------|-----------------|---------|
| Hunter.io | ✅ | ✅ | ✅ | ✅ | $49-399/mo |
| ZeroBounce | ✅ | ✅ | ✅ | ✅ | Pay-per-use |
| Apollo.io | ✅ | ✅ | ✅ | ✅ | $49-99/mo |
| NeverBounce | ✅ | ❌ | ✅ | ✅ | Pay-per-use |
| **Reacher** | ✅ | ❌ | ❌ | ❌ | Free/Self-hosted |

### 2.2 Current Limitations

1. **No visual interface**: Requires API integration, limiting adoption to technical users
2. **Verification-only**: Customers need separate tools for enrichment
3. **Single-use pattern**: No recurring engagement after initial list cleaning
4. **No enterprise features**: Missing compliance, audit logs, SSO
5. **No usage visibility**: Customers can't track consumption or costs
6. **Manual integrations**: Requires custom webhook implementations

### 2.3 Customer Feedback Themes

Based on support inquiries and community discussions:

1. "Can I upload a CSV without writing code?" (45% of inquiries)
2. "Do you have company/title data?" (30% of inquiries)
3. "How do I connect to Salesforce/HubSpot?" (15% of inquiries)
4. "Is this GDPR compliant?" (10% of inquiries)

### 2.4 Business Case

| Investment | Expected Return |
|------------|-----------------|
| Dashboard + Analytics | 5-10x user adoption (non-technical segment) |
| Enrichment Module | 3x revenue per customer (higher ARPU) |
| CRM Integrations | 2x retention (workflow stickiness) |
| Compliance Features | 2-5x pricing power (enterprise segment) |

---

## 3. Goals and Non-Goals

### 3.1 Goals

1. **Expand addressable market** from developers to marketing/sales teams
2. **Increase customer lifetime value** through recurring use cases
3. **Enable SaaS revenue model** with usage-based pricing
4. **Achieve feature parity** with Hunter.io and ZeroBounce
5. **Maintain open-source core** while offering commercial features
6. **Support enterprise requirements** (compliance, SSO, audit logs)

### 3.2 Non-Goals

1. **Replace CRMs**: We integrate with CRMs, not compete with them
2. **Build email sending**: We verify deliverability, not send emails
3. **Provide intent data**: Focus on contact data, not buying signals
4. **Build marketing automation**: Partner with, not compete against, ESPs
5. **Abandon self-hosted option**: Commercial features layer on top

### 3.3 Success Criteria

| Metric | Target (6 months) | Target (12 months) |
|--------|-------------------|-------------------|
| Monthly Active Users | 1,000 | 10,000 |
| Paid Customers | 100 | 1,000 |
| Monthly Recurring Revenue | $10,000 | $100,000 |
| Enterprise Customers (>$1k/mo) | 5 | 25 |
| Average Revenue Per User | $50 | $100 |
| Net Promoter Score | 40 | 60 |

---

## 4. Feature Specifications

### 4.1 Email List Management Dashboard

#### 4.1.1 Problem Statement

Currently, users must interact with Reacher via API, requiring:
- Technical knowledge (HTTP clients, JSON parsing)
- Custom code for CSV import/export
- Manual result aggregation and filtering

This excludes 80%+ of potential users (marketing teams, sales ops, founders).

#### 4.1.2 Proposed Solution

A web-based dashboard providing:

1. **List Upload & Management**
   - Drag-and-drop CSV/XLSX upload
   - Column mapping interface
   - Duplicate detection and removal
   - List naming, tagging, and organization

2. **Verification Progress**
   - Real-time progress indicators
   - Estimated completion time
   - Pause/resume/cancel controls
   - Error summaries and retry options

3. **Results Visualization**
   - Summary statistics (safe/risky/invalid/unknown)
   - Filterable results table
   - Detailed per-email breakdown
   - Risk score distribution charts

4. **Export & Actions**
   - Download cleaned list (CSV/XLSX)
   - Export by status (safe only, exclude invalid, etc.)
   - Direct CRM push (Phase 2)
   - Webhook configuration

#### 4.1.3 Technical Design

```
┌─────────────────────────────────────────────────────────────────┐
│                        Frontend (React)                         │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌───────────┐ │
│  │ List Upload │ │  Progress   │ │   Results   │ │  Export   │ │
│  │  Component  │ │  Tracker    │ │   Table     │ │  Actions  │ │
│  └──────┬──────┘ └──────┬──────┘ └──────┬──────┘ └─────┬─────┘ │
└─────────┼───────────────┼───────────────┼──────────────┼───────┘
          │               │               │              │
          ▼               ▼               ▼              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API Gateway (New)                          │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌───────────┐ │
│  │ POST /lists │ │ GET /lists  │ │ GET /lists/ │ │ GET /lists│ │
│  │   /upload   │ │  /{id}/     │ │  /{id}/     │ │  /{id}/   │ │
│  │             │ │  progress   │ │  results    │ │  export   │ │
│  └──────┬──────┘ └──────┬──────┘ └──────┬──────┘ └─────┬─────┘ │
└─────────┼───────────────┼───────────────┼──────────────┼───────┘
          │               │               │              │
          ▼               ▼               ▼              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Existing Backend                             │
│  ┌─────────────────────────────────────────────────────────────┐│
│  │              Bulk Verification Engine (v1)                  ││
│  │  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────────────┐││
│  │  │RabbitMQ │──│ Workers │──│PostgreSQL│──│ Result Storage ││ │
│  │  └─────────┘  └─────────┘  └─────────┘  └─────────────────┘││
│  └─────────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

#### 4.1.4 Frontend Technology Stack

| Component | Technology | Rationale |
|-----------|------------|-----------|
| Framework | React 18 + TypeScript | Industry standard, large ecosystem |
| State Management | TanStack Query | Server state, caching, optimistic updates |
| UI Components | shadcn/ui + Tailwind | Accessible, customizable, fast development |
| Charts | Recharts | Lightweight, React-native charting |
| File Upload | react-dropzone | Mature, accessible file handling |
| Tables | TanStack Table | Virtualization, sorting, filtering |
| Build Tool | Vite | Fast HMR, optimized builds |

#### 4.1.5 New API Endpoints

```yaml
# List Management
POST   /v2/lists                    # Create new list
GET    /v2/lists                    # List all lists (paginated)
GET    /v2/lists/{id}               # Get list details
DELETE /v2/lists/{id}               # Delete list
PATCH  /v2/lists/{id}               # Update list metadata

# File Upload
POST   /v2/lists/{id}/upload        # Upload CSV/XLSX file
GET    /v2/lists/{id}/upload/status # Upload processing status

# Verification
POST   /v2/lists/{id}/verify        # Start verification job
GET    /v2/lists/{id}/verify/status # Verification progress
POST   /v2/lists/{id}/verify/pause  # Pause verification
POST   /v2/lists/{id}/verify/resume # Resume verification
POST   /v2/lists/{id}/verify/cancel # Cancel verification

# Results
GET    /v2/lists/{id}/results       # Get results (paginated)
GET    /v2/lists/{id}/results/stats # Get summary statistics
GET    /v2/lists/{id}/export        # Export results (CSV/XLSX)
```

#### 4.1.6 Database Schema

```sql
-- New table: email_lists
CREATE TABLE email_lists (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    tags TEXT[] DEFAULT '{}',

    -- Upload metadata
    original_filename VARCHAR(255),
    file_size_bytes BIGINT,
    total_rows INTEGER,
    column_mapping JSONB, -- {"email": "Email Address", "name": "Full Name", ...}

    -- Verification status
    status VARCHAR(50) DEFAULT 'pending', -- pending, uploading, processing, verifying, completed, failed, cancelled
    verified_count INTEGER DEFAULT 0,
    safe_count INTEGER DEFAULT 0,
    risky_count INTEGER DEFAULT 0,
    invalid_count INTEGER DEFAULT 0,
    unknown_count INTEGER DEFAULT 0,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    verified_at TIMESTAMPTZ,

    -- Soft delete
    deleted_at TIMESTAMPTZ
);

-- New table: email_list_items
CREATE TABLE email_list_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    list_id UUID NOT NULL REFERENCES email_lists(id) ON DELETE CASCADE,

    -- Original data
    email VARCHAR(255) NOT NULL,
    original_data JSONB, -- Other columns from CSV
    row_number INTEGER,

    -- Verification result
    verification_id UUID REFERENCES v1_task_result(id),
    reachability VARCHAR(50), -- safe, risky, invalid, unknown
    is_disposable BOOLEAN,
    is_role_account BOOLEAN,
    is_free_provider BOOLEAN,
    mx_records TEXT[],
    smtp_response TEXT,

    -- Enrichment data (Phase 2)
    enrichment_data JSONB,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    verified_at TIMESTAMPTZ,

    -- Indexes
    CONSTRAINT unique_email_per_list UNIQUE (list_id, email)
);

CREATE INDEX idx_list_items_list_id ON email_list_items(list_id);
CREATE INDEX idx_list_items_reachability ON email_list_items(reachability);
CREATE INDEX idx_list_items_email ON email_list_items(email);
```

#### 4.1.7 User Interface Mockups

```
┌─────────────────────────────────────────────────────────────────────────┐
│  Reacher                                    [+ New List]  [Account ▼]  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │                                                                  │   │
│  │     ┌──────────────────────────────────────────────────────┐    │   │
│  │     │                                                       │    │   │
│  │     │        📁 Drag & drop your CSV file here             │    │   │
│  │     │              or click to browse                       │    │   │
│  │     │                                                       │    │   │
│  │     │         Supports: CSV, XLSX (max 100,000 rows)       │    │   │
│  │     │                                                       │    │   │
│  │     └──────────────────────────────────────────────────────┘    │   │
│  │                                                                  │   │
│  └─────────────────────────────────────────────────────────────────┘   │
│                                                                         │
│  Recent Lists                                                           │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │ Name              │ Emails  │ Status     │ Safe │ Created      │   │
│  ├───────────────────┼─────────┼────────────┼──────┼──────────────┤   │
│  │ Q4 Prospects      │ 5,432   │ ✅ Complete │ 78%  │ 2 hours ago  │   │
│  │ Conference Leads  │ 1,205   │ 🔄 45%      │ --   │ 10 mins ago  │   │
│  │ Newsletter Import │ 12,891  │ ✅ Complete │ 65%  │ Yesterday    │   │
│  └─────────────────────────────────────────────────────────────────┘   │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

```
┌─────────────────────────────────────────────────────────────────────────┐
│  Reacher  ›  Q4 Prospects                              [Export ▼]      │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ┌────────────┐ ┌────────────┐ ┌────────────┐ ┌────────────┐           │
│  │    5,432   │ │    4,237   │ │      892   │ │      303   │           │
│  │   Total    │ │    Safe    │ │    Risky   │ │   Invalid  │           │
│  │            │ │    78%     │ │    16%     │ │     6%     │           │
│  └────────────┘ └────────────┘ └────────────┘ └────────────┘           │
│                                                                         │
│  Filter: [All ▼]  [Search emails...]                                   │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │ ☐ │ Email                  │ Status  │ Provider │ Details      │   │
│  ├───┼────────────────────────┼─────────┼──────────┼──────────────┤   │
│  │ ☐ │ john@acme.com          │ ✅ Safe  │ Google   │ Deliverable  │   │
│  │ ☐ │ info@example.com       │ ⚠️ Risky │ Custom   │ Role account │   │
│  │ ☐ │ test@tempmail.com      │ ❌ Invalid│ Disposable│ Disposable  │   │
│  │ ☐ │ jane@corp.io           │ ✅ Safe  │ Microsoft│ Deliverable  │   │
│  └─────────────────────────────────────────────────────────────────┘   │
│                                                                         │
│  Showing 1-50 of 5,432  [◀ Prev] [1] [2] [3] ... [109] [Next ▶]       │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### 4.1.8 Implementation Phases

| Phase | Scope | Duration |
|-------|-------|----------|
| 1.1 | List CRUD, file upload, column mapping | 2 weeks |
| 1.2 | Verification integration, progress tracking | 2 weeks |
| 1.3 | Results table, filtering, pagination | 2 weeks |
| 1.4 | Export functionality, charts | 1 week |
| 1.5 | Polish, testing, documentation | 1 week |

**Total: 8 weeks**

---

### 4.2 Contact Enrichment Module

#### 4.2.1 Problem Statement

After verifying an email is deliverable, customers need additional context:
- Who is this person? (name, title, seniority)
- What company do they work for? (name, size, industry)
- How can I reach them? (LinkedIn, Twitter, phone)

Currently, customers must use separate services (Hunter, Apollo, Clearbit) for this data, creating friction and additional costs.

#### 4.2.2 Proposed Solution

An enrichment engine that augments verified emails with:

1. **Person Data**
   - Full name (first, last)
   - Job title
   - Seniority level (C-level, VP, Director, Manager, Individual Contributor)
   - Department (Sales, Marketing, Engineering, etc.)
   - LinkedIn profile URL
   - Twitter handle
   - Phone number (when available)
   - Profile photo URL

2. **Company Data**
   - Company name
   - Domain
   - Industry
   - Employee count range
   - Revenue range
   - Headquarters location
   - LinkedIn company page
   - Technologies used (optional)

3. **Enrichment Sources**
   - Primary: Licensed data partnerships (Apollo, RocketReach, or similar)
   - Secondary: Public web scraping (LinkedIn public profiles, company websites)
   - Tertiary: User-contributed data (with consent)

#### 4.2.3 Technical Design

```
┌─────────────────────────────────────────────────────────────────┐
│                    Enrichment Request Flow                       │
└─────────────────────────────────────────────────────────────────┘

  ┌──────────┐         ┌──────────────┐         ┌─────────────────┐
  │  Client  │────────▶│  Enrichment  │────────▶│   Data Source   │
  │          │         │   Gateway    │         │    Router       │
  └──────────┘         └──────────────┘         └────────┬────────┘
                                                         │
                       ┌─────────────────────────────────┼─────────────────────────────────┐
                       │                                 │                                 │
                       ▼                                 ▼                                 ▼
              ┌─────────────────┐             ┌─────────────────┐             ┌─────────────────┐
              │  Apollo API     │             │  RocketReach    │             │  Web Scraper    │
              │  (Primary)      │             │  (Secondary)    │             │  (Tertiary)     │
              └────────┬────────┘             └────────┬────────┘             └────────┬────────┘
                       │                               │                               │
                       └───────────────────────────────┼───────────────────────────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │  Data Merger &  │
                                              │  Normalizer     │
                                              └────────┬────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │  Enrichment     │
                                              │  Cache (Redis)  │
                                              └────────┬────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │  PostgreSQL     │
                                              │  (Persistent)   │
                                              └─────────────────┘
```

#### 4.2.4 Data Model

```rust
/// Person enrichment data
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PersonEnrichment {
    /// Full name
    pub full_name: Option<String>,
    /// First name
    pub first_name: Option<String>,
    /// Last name
    pub last_name: Option<String>,
    /// Job title
    pub title: Option<String>,
    /// Seniority level
    pub seniority: Option<Seniority>,
    /// Department
    pub department: Option<Department>,
    /// LinkedIn profile URL
    pub linkedin_url: Option<String>,
    /// Twitter handle (without @)
    pub twitter_handle: Option<String>,
    /// Phone number (E.164 format)
    pub phone: Option<String>,
    /// Profile photo URL
    pub photo_url: Option<String>,
    /// Location
    pub location: Option<Location>,
    /// Data freshness
    pub last_updated: DateTime<Utc>,
    /// Confidence score (0.0 - 1.0)
    pub confidence: f32,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum Seniority {
    CLevel,
    VicePresident,
    Director,
    Manager,
    Senior,
    Individual,
    Intern,
    Unknown,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum Department {
    Sales,
    Marketing,
    Engineering,
    Product,
    Design,
    Finance,
    HR,
    Legal,
    Operations,
    CustomerSuccess,
    Other(String),
}

/// Company enrichment data
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CompanyEnrichment {
    /// Company name
    pub name: String,
    /// Primary domain
    pub domain: String,
    /// Industry
    pub industry: Option<String>,
    /// Sub-industry
    pub sub_industry: Option<String>,
    /// Employee count range
    pub employee_count: Option<EmployeeRange>,
    /// Annual revenue range
    pub revenue: Option<RevenueRange>,
    /// Year founded
    pub founded_year: Option<i32>,
    /// Headquarters location
    pub headquarters: Option<Location>,
    /// Company LinkedIn URL
    pub linkedin_url: Option<String>,
    /// Company logo URL
    pub logo_url: Option<String>,
    /// Technologies used
    pub technologies: Vec<String>,
    /// Data freshness
    pub last_updated: DateTime<Utc>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct EmployeeRange {
    pub min: i32,
    pub max: Option<i32>,
    pub label: String, // "1-10", "11-50", "51-200", etc.
}
```

#### 4.2.5 Database Schema

```sql
-- Person enrichment cache
CREATE TABLE person_enrichment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    email_domain VARCHAR(255) NOT NULL,

    -- Person data
    full_name VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    title VARCHAR(255),
    seniority VARCHAR(50),
    department VARCHAR(100),
    linkedin_url TEXT,
    twitter_handle VARCHAR(100),
    phone VARCHAR(50),
    photo_url TEXT,
    location_city VARCHAR(100),
    location_state VARCHAR(100),
    location_country VARCHAR(100),

    -- Metadata
    source VARCHAR(50), -- 'apollo', 'rocketreach', 'scraper', 'user'
    confidence DECIMAL(3,2),
    raw_data JSONB,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ DEFAULT NOW() + INTERVAL '90 days'
);

-- Company enrichment cache
CREATE TABLE company_enrichment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    domain VARCHAR(255) NOT NULL UNIQUE,

    -- Company data
    name VARCHAR(255) NOT NULL,
    industry VARCHAR(100),
    sub_industry VARCHAR(100),
    employee_count_min INTEGER,
    employee_count_max INTEGER,
    employee_count_label VARCHAR(50),
    revenue_min BIGINT,
    revenue_max BIGINT,
    revenue_label VARCHAR(50),
    founded_year INTEGER,
    headquarters_city VARCHAR(100),
    headquarters_state VARCHAR(100),
    headquarters_country VARCHAR(100),
    linkedin_url TEXT,
    logo_url TEXT,
    technologies TEXT[],

    -- Metadata
    source VARCHAR(50),
    raw_data JSONB,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ DEFAULT NOW() + INTERVAL '90 days'
);

CREATE INDEX idx_person_email_domain ON person_enrichment(email_domain);
CREATE INDEX idx_person_expires ON person_enrichment(expires_at);
CREATE INDEX idx_company_expires ON company_enrichment(expires_at);
```

#### 4.2.6 API Endpoints

```yaml
# Single email enrichment
POST /v2/enrich
Content-Type: application/json
{
  "email": "john@acme.com",
  "include": ["person", "company"],  # Optional, defaults to both
  "force_refresh": false              # Optional, bypass cache
}

Response:
{
  "email": "john@acme.com",
  "person": {
    "full_name": "John Smith",
    "first_name": "John",
    "last_name": "Smith",
    "title": "VP of Sales",
    "seniority": "vice_president",
    "department": "sales",
    "linkedin_url": "https://linkedin.com/in/johnsmith",
    "phone": "+1-555-123-4567",
    "confidence": 0.92
  },
  "company": {
    "name": "Acme Corporation",
    "domain": "acme.com",
    "industry": "Software",
    "employee_count": {"min": 51, "max": 200, "label": "51-200"},
    "headquarters": {"city": "San Francisco", "state": "CA", "country": "US"},
    "linkedin_url": "https://linkedin.com/company/acme",
    "logo_url": "https://logo.clearbit.com/acme.com"
  },
  "credits_used": 1
}

# Bulk enrichment
POST /v2/enrich/bulk
Content-Type: application/json
{
  "emails": ["john@acme.com", "jane@corp.io", ...],
  "include": ["person", "company"],
  "webhook_url": "https://your-app.com/webhook"  # Optional
}

Response:
{
  "job_id": "enrich_abc123",
  "total": 100,
  "estimated_seconds": 30,
  "status_url": "/v2/enrich/bulk/enrich_abc123"
}

# Combined verify + enrich
POST /v2/check_email
Content-Type: application/json
{
  "to_email": "john@acme.com",
  "enrich": true  # New optional parameter
}
```

#### 4.2.7 Data Source Integration

**Option A: Apollo.io API**
- Coverage: 250M+ contacts
- Pricing: ~$0.01-0.03 per enrichment
- Quality: High accuracy, frequently updated
- Integration: REST API with bulk support

**Option B: RocketReach API**
- Coverage: 700M+ profiles
- Pricing: ~$0.02-0.05 per enrichment
- Quality: Good for tech industry
- Integration: REST API

**Option C: Build Own Scraper**
- Coverage: Limited to public data
- Pricing: Infrastructure costs only
- Quality: Variable, requires maintenance
- Integration: Custom development

**Recommended Approach**: Start with Apollo.io API for primary data, with RocketReach as fallback. Build scraper for supplemental data and cost optimization at scale.

#### 4.2.8 Pricing Model

| Tier | Enrichments/Month | Price | Per-Enrichment |
|------|-------------------|-------|----------------|
| Free | 100 | $0 | $0.00 |
| Starter | 1,000 | $29 | $0.029 |
| Growth | 10,000 | $149 | $0.015 |
| Pro | 50,000 | $499 | $0.010 |
| Enterprise | Unlimited | Custom | Negotiated |

---

### 4.3 Usage Analytics & Billing Dashboard

#### 4.3.1 Problem Statement

Current limitations:
- Single API key with no usage visibility
- No way to track consumption by endpoint, time period, or user
- No billing integration for metered usage
- No alerts for quota approaching/exceeded

#### 4.3.2 Proposed Solution

1. **Usage Tracking**
   - Per-API-key usage metrics
   - Breakdown by endpoint, status, response time
   - Historical charts (hourly, daily, monthly)
   - Real-time usage counters

2. **Billing Integration**
   - Stripe-based subscription management
   - Usage-based billing with metered products
   - Invoice generation and history
   - Payment method management

3. **Alerting**
   - Quota threshold alerts (50%, 80%, 100%)
   - Error rate alerts
   - Email and webhook notifications

#### 4.3.3 Database Schema

```sql
-- API keys with organization association
CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    key_hash VARCHAR(64) NOT NULL UNIQUE, -- SHA-256 of actual key
    key_prefix VARCHAR(8) NOT NULL, -- First 8 chars for identification
    name VARCHAR(255),

    -- Permissions
    scopes TEXT[] DEFAULT '{verify,enrich,bulk}',
    rate_limit_per_second INTEGER DEFAULT 10,
    rate_limit_per_month INTEGER,

    -- Status
    is_active BOOLEAN DEFAULT true,
    last_used_at TIMESTAMPTZ,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ,
    revoked_at TIMESTAMPTZ
);

-- Usage events (high-volume, consider TimescaleDB)
CREATE TABLE usage_events (
    id BIGSERIAL PRIMARY KEY,
    api_key_id UUID NOT NULL REFERENCES api_keys(id),
    organization_id UUID NOT NULL,

    -- Request details
    endpoint VARCHAR(100) NOT NULL,
    method VARCHAR(10) NOT NULL,
    status_code INTEGER NOT NULL,
    response_time_ms INTEGER,

    -- Metering
    emails_verified INTEGER DEFAULT 0,
    emails_enriched INTEGER DEFAULT 0,
    credits_used INTEGER DEFAULT 0,

    -- Context
    ip_address INET,
    user_agent TEXT,

    -- Timestamp
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Partition by month for performance
CREATE INDEX idx_usage_events_org_time ON usage_events(organization_id, created_at);
CREATE INDEX idx_usage_events_key_time ON usage_events(api_key_id, created_at);

-- Aggregated usage (materialized for performance)
CREATE TABLE usage_daily (
    id BIGSERIAL PRIMARY KEY,
    organization_id UUID NOT NULL,
    api_key_id UUID,
    date DATE NOT NULL,

    -- Aggregates
    total_requests INTEGER DEFAULT 0,
    successful_requests INTEGER DEFAULT 0,
    failed_requests INTEGER DEFAULT 0,
    emails_verified INTEGER DEFAULT 0,
    emails_enriched INTEGER DEFAULT 0,
    credits_used INTEGER DEFAULT 0,
    avg_response_time_ms INTEGER,

    -- Unique constraint
    UNIQUE(organization_id, api_key_id, date)
);

-- Billing
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    stripe_subscription_id VARCHAR(255) UNIQUE,
    stripe_customer_id VARCHAR(255),

    -- Plan details
    plan_id VARCHAR(50) NOT NULL,
    plan_name VARCHAR(100),
    monthly_credits INTEGER,
    credits_remaining INTEGER,
    credits_reset_at TIMESTAMPTZ,

    -- Status
    status VARCHAR(50) DEFAULT 'active', -- active, past_due, canceled, trialing
    trial_ends_at TIMESTAMPTZ,
    current_period_start TIMESTAMPTZ,
    current_period_end TIMESTAMPTZ,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    canceled_at TIMESTAMPTZ
);
```

#### 4.3.4 API Endpoints

```yaml
# Usage endpoints
GET /v2/usage/summary
  Query: period=day|week|month|year, start_date, end_date

GET /v2/usage/breakdown
  Query: group_by=endpoint|status|api_key, period, start_date, end_date

GET /v2/usage/credits
  Response: { used: 5000, remaining: 15000, resets_at: "2024-02-01" }

# API key management
GET    /v2/api-keys
POST   /v2/api-keys
GET    /v2/api-keys/{id}
PATCH  /v2/api-keys/{id}
DELETE /v2/api-keys/{id}
POST   /v2/api-keys/{id}/rotate

# Billing (Stripe integration)
GET  /v2/billing/subscription
POST /v2/billing/subscription          # Create/update subscription
GET  /v2/billing/invoices
GET  /v2/billing/payment-methods
POST /v2/billing/payment-methods
POST /v2/billing/portal-session        # Stripe customer portal
```

#### 4.3.5 Dashboard UI

```
┌─────────────────────────────────────────────────────────────────────────┐
│  Reacher  ›  Usage & Billing                                           │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  Current Period: Dec 1 - Dec 31, 2024                                  │
│                                                                         │
│  ┌────────────────────────────────────────────────────────────────┐    │
│  │  Credits Used                                    Plan: Growth   │    │
│  │  ████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░  5,432 / 10,000 │    │
│  │                                                     54% used    │    │
│  └────────────────────────────────────────────────────────────────┘    │
│                                                                         │
│  ┌──────────────────────────────┐  ┌──────────────────────────────┐   │
│  │  Verifications This Month    │  │  Enrichments This Month      │   │
│  │         4,892                │  │         540                   │   │
│  │  ▲ 23% vs last month         │  │  ▲ 45% vs last month         │   │
│  └──────────────────────────────┘  └──────────────────────────────┘   │
│                                                                         │
│  Usage Over Time                                        [Day ▼]        │
│  ┌────────────────────────────────────────────────────────────────┐    │
│  │     ╭─╮                                                         │    │
│  │   ╭─╯ ╰─╮     ╭──╮                         ╭─╮                 │    │
│  │ ╭─╯     ╰─────╯  ╰─────────────────────────╯ ╰─────            │    │
│  │─────────────────────────────────────────────────────────────── │    │
│  │ Dec 1    Dec 5    Dec 10    Dec 15    Dec 20    Dec 25        │    │
│  └────────────────────────────────────────────────────────────────┘    │
│                                                                         │
│  API Keys                                              [+ Create Key]  │
│  ┌────────────────────────────────────────────────────────────────┐    │
│  │ Name            │ Key          │ Requests │ Last Used          │    │
│  ├─────────────────┼──────────────┼──────────┼────────────────────┤    │
│  │ Production      │ rch_abc1...  │ 4,521    │ 2 minutes ago      │    │
│  │ Development     │ rch_xyz9...  │ 371      │ 3 days ago         │    │
│  └────────────────────────────────────────────────────────────────┘    │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

### 4.4 Native CRM Integrations

#### 4.4.1 Problem Statement

Currently, users must:
1. Export contacts from CRM
2. Upload to Reacher
3. Download verified results
4. Re-import to CRM
5. Manually update records

This creates friction and limits adoption in sales/marketing workflows.

#### 4.4.2 Proposed Solution

Native integrations with major CRMs:

1. **Salesforce**
   - AppExchange listing
   - Real-time verification on Lead/Contact creation
   - Bulk verification of existing records
   - Custom field mapping
   - Verification status synced to CRM

2. **HubSpot**
   - App Marketplace listing
   - Workflow action for verification
   - Property syncing
   - Contact enrichment integration

3. **Pipedrive**
   - Marketplace listing
   - Person field verification
   - Deal enrichment

4. **Outreach/Salesloft** (Future)
   - Prospect verification before sequences
   - Bounce prevention

#### 4.4.3 Technical Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    CRM Integration Layer                         │
└─────────────────────────────────────────────────────────────────┘

┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────────┐
│  Salesforce │ │   HubSpot   │ │  Pipedrive  │ │ Outreach/Loft   │
│   Adapter   │ │   Adapter   │ │   Adapter   │ │    Adapter      │
└──────┬──────┘ └──────┬──────┘ └──────┬──────┘ └────────┬────────┘
       │               │               │                  │
       └───────────────┴───────┬───────┴──────────────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │  Integration Core   │
                    │  ┌───────────────┐  │
                    │  │ OAuth Manager │  │
                    │  ├───────────────┤  │
                    │  │ Field Mapper  │  │
                    │  ├───────────────┤  │
                    │  │ Sync Engine   │  │
                    │  ├───────────────┤  │
                    │  │ Webhook Proc  │  │
                    │  └───────────────┘  │
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │   Reacher Core API  │
                    │  - Verification     │
                    │  - Enrichment       │
                    │  - Bulk Jobs        │
                    └─────────────────────┘
```

#### 4.4.4 Salesforce Integration Details

**Authentication**: OAuth 2.0 with refresh tokens

**Objects Supported**:
- Lead
- Contact
- Account (for domain enrichment)

**Custom Fields Created**:
```
Email_Verification_Status__c    (Picklist: Safe, Risky, Invalid, Unknown, Pending)
Email_Verified_At__c            (DateTime)
Email_Is_Disposable__c          (Checkbox)
Email_Is_Role_Account__c        (Checkbox)
Email_MX_Provider__c            (Text)
Enrichment_Job_Title__c         (Text)
Enrichment_LinkedIn_URL__c      (URL)
Enrichment_Company_Size__c      (Text)
```

**Triggers**:
1. After Insert on Lead/Contact → Queue verification
2. Manual "Verify Email" button → Immediate verification
3. Scheduled job → Bulk re-verification

**Apex Classes**:
```apex
// Verification trigger
trigger LeadEmailVerification on Lead (after insert, after update) {
    ReacherVerificationService.queueVerification(Trigger.new);
}

// Verification service
public class ReacherVerificationService {
    @future(callout=true)
    public static void queueVerification(List<Lead> leads) {
        // Call Reacher API
        // Update Lead records with results
    }
}
```

#### 4.4.5 HubSpot Integration Details

**Authentication**: OAuth 2.0

**Features**:
1. **Workflow Action**: "Verify Email with Reacher"
   - Usable in any workflow
   - Branch based on verification result

2. **Property Sync**:
   - `reacher_email_status` (enumeration)
   - `reacher_verified_at` (datetime)
   - `reacher_is_disposable` (boolean)
   - `reacher_enrichment_title` (string)

3. **Contact Import Integration**:
   - Verify on import
   - Show verification status in import preview

**API Endpoints for HubSpot**:
```yaml
POST /v2/integrations/hubspot/install
  # OAuth callback, creates connected app

POST /v2/integrations/hubspot/webhook
  # Receives CRM events (contact created, etc.)

GET /v2/integrations/hubspot/properties
  # Returns property definitions for sync
```

---

### 4.5 Compliance & Audit Logging

#### 4.5.1 Problem Statement

Enterprise customers require:
- GDPR/CCPA compliance documentation
- Immutable audit trails
- Data retention policies
- Right to deletion support
- SOC 2 compatible logging

Currently, Reacher has no built-in compliance features.

#### 4.5.2 Proposed Solution

1. **Audit Logging**
   - Immutable event log for all operations
   - Who did what, when, from where
   - API access logs with request/response hashes
   - User action logs (login, export, delete)

2. **Data Retention**
   - Configurable retention periods
   - Automatic data expiration
   - Manual purge capabilities

3. **GDPR/CCPA Support**
   - Data subject access requests (DSAR)
   - Right to deletion implementation
   - Consent tracking
   - Data processing records

4. **Compliance Reporting**
   - Audit log exports
   - Compliance dashboards
   - Scheduled compliance reports

#### 4.5.3 Database Schema

```sql
-- Immutable audit log (append-only)
CREATE TABLE audit_log (
    id BIGSERIAL PRIMARY KEY,

    -- Event identification
    event_id UUID NOT NULL DEFAULT gen_random_uuid(),
    event_type VARCHAR(100) NOT NULL,
    event_category VARCHAR(50) NOT NULL, -- 'api', 'user', 'system', 'compliance'

    -- Actor
    organization_id UUID,
    user_id UUID,
    api_key_id UUID,
    actor_type VARCHAR(20) NOT NULL, -- 'user', 'api_key', 'system', 'integration'
    actor_ip INET,
    actor_user_agent TEXT,

    -- Resource
    resource_type VARCHAR(50), -- 'email', 'list', 'user', 'api_key', etc.
    resource_id VARCHAR(255),

    -- Action details
    action VARCHAR(50) NOT NULL, -- 'create', 'read', 'update', 'delete', 'verify', 'export'
    action_detail JSONB, -- Additional context

    -- Request/Response (hashed for integrity)
    request_hash VARCHAR(64), -- SHA-256 of request body
    response_hash VARCHAR(64), -- SHA-256 of response body

    -- Integrity
    previous_hash VARCHAR(64), -- Hash of previous log entry (blockchain-style)
    entry_hash VARCHAR(64) NOT NULL, -- Hash of this entry

    -- Timestamp (immutable)
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Prevent updates/deletes
CREATE RULE audit_log_no_update AS ON UPDATE TO audit_log DO INSTEAD NOTHING;
CREATE RULE audit_log_no_delete AS ON DELETE TO audit_log DO INSTEAD NOTHING;

-- Data retention policies
CREATE TABLE data_retention_policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID REFERENCES organizations(id),

    -- Policy settings
    verification_results_days INTEGER DEFAULT 365,
    enrichment_data_days INTEGER DEFAULT 180,
    audit_logs_days INTEGER DEFAULT 2555, -- 7 years
    usage_events_days INTEGER DEFAULT 365,

    -- Auto-purge
    auto_purge_enabled BOOLEAN DEFAULT false,
    last_purge_at TIMESTAMPTZ,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- GDPR/CCPA requests
CREATE TABLE data_subject_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),

    -- Request details
    request_type VARCHAR(50) NOT NULL, -- 'access', 'deletion', 'portability', 'rectification'
    subject_email VARCHAR(255) NOT NULL,
    subject_identity_verified BOOLEAN DEFAULT false,

    -- Status
    status VARCHAR(50) DEFAULT 'pending', -- pending, processing, completed, rejected
    status_reason TEXT,

    -- Processing
    processed_by UUID REFERENCES users(id),
    processed_at TIMESTAMPTZ,

    -- Data collected (for access requests)
    collected_data JSONB,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    due_at TIMESTAMPTZ DEFAULT NOW() + INTERVAL '30 days',
    completed_at TIMESTAMPTZ
);

-- Consent records
CREATE TABLE consent_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),

    -- Subject
    email VARCHAR(255) NOT NULL,

    -- Consent details
    consent_type VARCHAR(50) NOT NULL, -- 'verification', 'enrichment', 'marketing'
    consent_given BOOLEAN NOT NULL,
    consent_method VARCHAR(50), -- 'api', 'form', 'import', 'integration'
    consent_source TEXT, -- URL or identifier of where consent was given

    -- Legal basis
    legal_basis VARCHAR(50), -- 'consent', 'legitimate_interest', 'contract', 'legal_obligation'

    -- Timestamps
    given_at TIMESTAMPTZ NOT NULL,
    expires_at TIMESTAMPTZ,
    withdrawn_at TIMESTAMPTZ,

    -- Audit
    recorded_at TIMESTAMPTZ DEFAULT NOW(),
    recorded_by UUID REFERENCES users(id)
);
```

#### 4.5.4 API Endpoints

```yaml
# Audit logs
GET /v2/audit-logs
  Query: event_type, start_date, end_date, actor_id, resource_type, limit, offset
  Response: Paginated audit log entries

GET /v2/audit-logs/export
  Query: format=json|csv, start_date, end_date
  Response: Downloadable file

# Data retention
GET  /v2/compliance/retention-policy
PUT  /v2/compliance/retention-policy
POST /v2/compliance/purge           # Manual purge trigger

# GDPR/CCPA
POST /v2/compliance/dsar            # Create data subject request
GET  /v2/compliance/dsar/{id}       # Get request status
GET  /v2/compliance/dsar/{id}/data  # Download collected data

# Consent
POST /v2/compliance/consent         # Record consent
GET  /v2/compliance/consent/{email} # Get consent status for email
DELETE /v2/compliance/consent/{id}  # Withdraw consent
```

---

### 4.6 Spam Trap & Honeypot Detection

#### 4.6.1 Problem Statement

Sending to spam traps causes:
- IP/domain blacklisting
- Sender reputation damage
- Deliverability drops (sometimes permanent)

Current verification detects invalid emails but not spam traps.

#### 4.6.2 Proposed Solution

1. **Known Spam Trap Database**
   - Integration with spam trap data providers
   - Community-reported spam traps
   - Historical bounce pattern analysis

2. **Heuristic Detection**
   - Email age estimation
   - Engagement pattern analysis
   - Domain reputation scoring

3. **Risk Scoring**
   - Spam trap probability score
   - Honeypot likelihood
   - Overall risk assessment

#### 4.6.3 Data Sources

| Source | Type | Coverage | Cost |
|--------|------|----------|------|
| Spamhaus | Blocklist API | Global | Enterprise license |
| Kickbox | Spam trap data | 100M+ | Per-query |
| Internal | Bounce analysis | Own data | Infrastructure |
| Community | User reports | Crowdsourced | Free |

#### 4.6.4 Detection Algorithm

```rust
/// Spam trap detection result
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SpamTrapAnalysis {
    /// Is this a known spam trap?
    pub is_known_trap: bool,
    /// Spam trap type if known
    pub trap_type: Option<SpamTrapType>,
    /// Probability this is a spam trap (0.0 - 1.0)
    pub trap_probability: f32,
    /// Risk factors detected
    pub risk_factors: Vec<RiskFactor>,
    /// Recommendation
    pub recommendation: SpamTrapRecommendation,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum SpamTrapType {
    /// Pristine trap - never used by real person
    Pristine,
    /// Recycled trap - abandoned email repurposed
    Recycled,
    /// Typo trap - common misspellings
    Typo,
    /// Honeypot - hidden form field
    Honeypot,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RiskFactor {
    pub factor: String,
    pub weight: f32,
    pub description: String,
}

// Example risk factors:
// - "no_mx_history": Domain MX records very new
// - "generic_pattern": Email matches known trap patterns (e.g., test@, spam@)
// - "no_social_presence": No LinkedIn/social profiles found
// - "blocklist_adjacent": Domain/IP near known blocklisted entities
// - "suspicious_registration": Domain registered recently with privacy guard
```

#### 4.6.5 API Response Addition

```json
{
  "email": "test@suspicious-domain.com",
  "is_reachable": "risky",
  "spam_trap": {
    "is_known_trap": false,
    "trap_probability": 0.73,
    "trap_type": null,
    "risk_factors": [
      {
        "factor": "generic_pattern",
        "weight": 0.3,
        "description": "Email matches common spam trap patterns"
      },
      {
        "factor": "new_domain",
        "weight": 0.25,
        "description": "Domain registered less than 6 months ago"
      },
      {
        "factor": "no_social_presence",
        "weight": 0.18,
        "description": "No LinkedIn or social profiles found"
      }
    ],
    "recommendation": "avoid"
  }
}
```

---

### 4.7 Scheduled Re-verification

#### 4.7.1 Problem Statement

Email lists degrade over time:
- Employees change jobs (5-10% annually)
- Email accounts get deactivated
- Domains expire or change MX records
- Catch-all settings change

One-time verification becomes stale within months.

#### 4.7.2 Proposed Solution

1. **Scheduled Jobs**
   - Daily, weekly, monthly re-verification
   - Custom cron schedules
   - Smart scheduling (only re-verify emails older than X days)

2. **Change Detection**
   - Track status changes over time
   - Alert on significant degradation
   - Identify newly invalid emails

3. **Monitoring Dashboard**
   - List health scores
   - Trend visualization
   - Proactive alerts

#### 4.7.3 Database Schema

```sql
-- Verification schedules
CREATE TABLE verification_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    list_id UUID REFERENCES email_lists(id), -- NULL = all lists

    -- Schedule
    schedule_type VARCHAR(20) NOT NULL, -- 'once', 'daily', 'weekly', 'monthly', 'cron'
    cron_expression VARCHAR(100), -- For custom schedules
    timezone VARCHAR(50) DEFAULT 'UTC',

    -- Filters
    only_if_older_than_days INTEGER DEFAULT 30,
    only_status TEXT[], -- Only re-verify these statuses

    -- Status
    is_active BOOLEAN DEFAULT true,
    last_run_at TIMESTAMPTZ,
    next_run_at TIMESTAMPTZ,

    -- Notifications
    notify_on_completion BOOLEAN DEFAULT true,
    notify_on_degradation BOOLEAN DEFAULT true,
    degradation_threshold DECIMAL(3,2) DEFAULT 0.05, -- 5% drop triggers alert
    notification_emails TEXT[],
    notification_webhook_url TEXT,

    -- Timestamps
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Verification history for trend tracking
CREATE TABLE email_verification_history (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    organization_id UUID NOT NULL,

    -- Result
    reachability VARCHAR(20) NOT NULL,
    is_disposable BOOLEAN,
    is_role_account BOOLEAN,
    smtp_response TEXT,

    -- Change tracking
    previous_reachability VARCHAR(20),
    status_changed BOOLEAN DEFAULT false,

    -- Timestamps
    verified_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_verification_history_email ON email_verification_history(email, verified_at DESC);
CREATE INDEX idx_verification_history_org ON email_verification_history(organization_id, verified_at DESC);
```

#### 4.7.4 API Endpoints

```yaml
# Schedules
POST   /v2/schedules                  # Create schedule
GET    /v2/schedules                  # List schedules
GET    /v2/schedules/{id}             # Get schedule
PATCH  /v2/schedules/{id}             # Update schedule
DELETE /v2/schedules/{id}             # Delete schedule
POST   /v2/schedules/{id}/run         # Trigger immediate run

# History & Trends
GET /v2/lists/{id}/history
  Query: email, start_date, end_date
  Response: Verification history with changes

GET /v2/lists/{id}/health
  Response: {
    "health_score": 0.85,
    "trend": "stable",
    "safe_percentage": 78,
    "degradation_rate": 0.02,
    "last_verified_at": "2024-12-01T00:00:00Z",
    "emails_changed_since_last": 42
  }
```

---

### 4.8 Bulk Domain Verification

#### 4.8.1 Problem Statement

Customers need to validate domains, not just individual emails:
- Is this domain real and receiving mail?
- Does it have proper email authentication (SPF/DKIM/DMARC)?
- Is the domain reputable or suspicious?
- What's the WHOIS information?

#### 4.8.2 Proposed Solution

1. **DNS Analysis**
   - MX record validation
   - SPF record parsing
   - DKIM selector discovery
   - DMARC policy extraction

2. **WHOIS Lookup**
   - Registrar information
   - Registration date
   - Expiration date
   - Privacy protection status

3. **Reputation Scoring**
   - Blocklist checking
   - Age-based scoring
   - SSL certificate validation
   - Web presence verification

#### 4.8.3 Data Model

```rust
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DomainVerification {
    pub domain: String,

    // DNS Records
    pub mx_records: Vec<MxRecord>,
    pub has_valid_mx: bool,
    pub spf: Option<SpfRecord>,
    pub dmarc: Option<DmarcRecord>,
    pub dkim_selectors: Vec<String>,

    // WHOIS
    pub whois: Option<WhoisData>,

    // Reputation
    pub reputation_score: f32,
    pub is_blocklisted: bool,
    pub blocklist_entries: Vec<BlocklistEntry>,

    // Web presence
    pub has_website: bool,
    pub ssl_valid: bool,
    pub ssl_expires_at: Option<DateTime<Utc>>,

    // Classification
    pub domain_type: DomainType,
    pub risk_level: RiskLevel,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SpfRecord {
    pub raw: String,
    pub is_valid: bool,
    pub mechanisms: Vec<String>,
    pub includes: Vec<String>,
    pub policy: String, // "pass", "softfail", "fail", "neutral"
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DmarcRecord {
    pub raw: String,
    pub policy: String, // "none", "quarantine", "reject"
    pub subdomain_policy: Option<String>,
    pub percentage: i32,
    pub rua: Vec<String>, // Aggregate report addresses
    pub ruf: Vec<String>, // Forensic report addresses
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct WhoisData {
    pub registrar: Option<String>,
    pub created_at: Option<DateTime<Utc>>,
    pub updated_at: Option<DateTime<Utc>>,
    pub expires_at: Option<DateTime<Utc>>,
    pub registrant_org: Option<String>,
    pub registrant_country: Option<String>,
    pub privacy_protected: bool,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum DomainType {
    Corporate,
    FreeProvider,
    Disposable,
    Educational,
    Government,
    Unknown,
}
```

#### 4.8.4 API Endpoints

```yaml
# Single domain
POST /v2/domain/verify
{
  "domain": "acme.com",
  "include": ["dns", "whois", "reputation", "web"]
}

# Bulk domain
POST /v2/domain/verify/bulk
{
  "domains": ["acme.com", "corp.io", ...],
  "include": ["dns", "whois"],
  "webhook_url": "https://..."
}
```

---

### 4.9 Sender Reputation Scoring

#### 4.9.1 Problem Statement

Knowing an email is valid isn't enough - customers need to know:
- Can I safely send to this domain from my IP/domain?
- What's the recipient domain's spam filtering level?
- What's my current sender reputation?

#### 4.9.2 Proposed Solution

1. **Recipient Domain Analysis**
   - Spam filter aggressiveness estimation
   - Historical delivery rates (aggregated)
   - Known email security vendors

2. **Sender Reputation Check**
   - IP blocklist status
   - Domain blocklist status
   - Authentication status (SPF/DKIM/DMARC)

3. **Deliverability Prediction**
   - Likelihood of inbox vs spam
   - Recommended warm-up period
   - Risk factors

#### 4.9.3 Data Model

```rust
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SenderReputation {
    // Sender details
    pub sending_domain: String,
    pub sending_ip: Option<IpAddr>,

    // Reputation scores (0-100)
    pub domain_reputation: i32,
    pub ip_reputation: Option<i32>,
    pub overall_score: i32,

    // Authentication
    pub spf_status: AuthStatus,
    pub dkim_status: AuthStatus,
    pub dmarc_status: AuthStatus,

    // Blocklist status
    pub blocklists: Vec<BlocklistCheck>,
    pub is_blocklisted: bool,

    // Recommendations
    pub recommendations: Vec<ReputationRecommendation>,
    pub estimated_inbox_rate: f32,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RecipientDomainAnalysis {
    pub domain: String,

    // Security posture
    pub email_security_vendor: Option<String>, // "Proofpoint", "Mimecast", etc.
    pub spam_filter_aggressiveness: AggressivenessLevel,

    // Historical data (aggregated, anonymized)
    pub historical_delivery_rate: Option<f32>,
    pub historical_bounce_rate: Option<f32>,

    // Recommendations
    pub sending_recommendations: Vec<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum AggressivenessLevel {
    Low,      // Most emails get through
    Medium,   // Standard filtering
    High,     // Strict filtering (enterprise)
    VeryHigh, // Very strict (government, financial)
}
```

---

### 4.10 Advanced Segmentation & Scoring Engine

#### 4.10.1 Problem Statement

Customers want to prioritize outreach based on multiple factors:
- Email quality (safe vs risky)
- Data freshness
- Enrichment completeness
- Engagement potential

Currently, only basic reachability status is provided.

#### 4.10.2 Proposed Solution

1. **Multi-Factor Scoring**
   - Deliverability score
   - Data quality score
   - Engagement potential score
   - Overall lead score

2. **Custom Scoring Rules**
   - User-defined scoring formulas
   - Weight adjustments
   - Threshold configuration

3. **Segmentation**
   - Dynamic segments based on scores
   - Saved segment definitions
   - Segment-based exports

#### 4.10.3 Data Model

```rust
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct EmailScore {
    pub email: String,

    // Component scores (0-100)
    pub deliverability_score: i32,
    pub data_quality_score: i32,
    pub engagement_potential_score: i32,

    // Overall score (weighted average)
    pub overall_score: i32,
    pub score_grade: ScoreGrade, // A, B, C, D, F

    // Score factors
    pub positive_factors: Vec<ScoreFactor>,
    pub negative_factors: Vec<ScoreFactor>,

    // Recommendations
    pub priority: Priority,
    pub recommended_action: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ScoreFactor {
    pub name: String,
    pub impact: i32, // Positive or negative
    pub description: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ScoringRule {
    pub id: Uuid,
    pub name: String,
    pub description: String,

    // Weights
    pub deliverability_weight: f32,
    pub data_quality_weight: f32,
    pub engagement_weight: f32,

    // Custom factors
    pub custom_factors: Vec<CustomFactor>,

    // Thresholds
    pub grade_thresholds: GradeThresholds,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CustomFactor {
    pub name: String,
    pub condition: ScoreCondition,
    pub impact: i32,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum ScoreCondition {
    HasEnrichment { field: String },
    ReachabilityIs { status: String },
    DomainAge { min_days: i32 },
    HasSocialProfile,
    IsB2B,
    TitleContains { keywords: Vec<String> },
    CompanySizeMin { min: i32 },
}
```

#### 4.10.4 API Endpoints

```yaml
# Scoring
POST /v2/score
{
  "email": "john@acme.com",
  "scoring_rule_id": "rule_123"  # Optional, uses default if not provided
}

POST /v2/score/bulk
{
  "list_id": "list_abc",
  "scoring_rule_id": "rule_123"
}

# Scoring rules
GET    /v2/scoring-rules
POST   /v2/scoring-rules
GET    /v2/scoring-rules/{id}
PUT    /v2/scoring-rules/{id}
DELETE /v2/scoring-rules/{id}

# Segments
GET    /v2/segments
POST   /v2/segments
GET    /v2/segments/{id}
PUT    /v2/segments/{id}
DELETE /v2/segments/{id}
GET    /v2/segments/{id}/emails  # Get emails matching segment
```

---

## 5. System Architecture

### 5.1 High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              Client Layer                                    │
│  ┌───────────────┐  ┌───────────────┐  ┌───────────────┐  ┌──────────────┐ │
│  │   Dashboard   │  │    REST API   │  │  CRM Plugins  │  │     CLI      │ │
│  │    (React)    │  │   Clients     │  │ (SF, HS, PD)  │  │              │ │
│  └───────┬───────┘  └───────┬───────┘  └───────┬───────┘  └──────┬───────┘ │
└──────────┼──────────────────┼──────────────────┼─────────────────┼─────────┘
           │                  │                  │                 │
           └──────────────────┴────────┬─────────┴─────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              API Gateway                                     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │    Auth     │  │    Rate     │  │   Request   │  │      Routing        │ │
│  │  (JWT/API)  │  │   Limiting  │  │   Logging   │  │   (v0/v1/v2)        │ │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────────────┘ │
└─────────────────────────────────────────┬───────────────────────────────────┘
                                          │
                    ┌─────────────────────┼─────────────────────┐
                    │                     │                     │
                    ▼                     ▼                     ▼
┌───────────────────────────┐ ┌─────────────────────┐ ┌─────────────────────┐
│    Verification Service   │ │  Enrichment Service │ │  Analytics Service  │
│  ┌─────────────────────┐  │ │ ┌─────────────────┐ │ │ ┌─────────────────┐ │
│  │  Syntax Validator   │  │ │ │  Apollo Client  │ │ │ │  Usage Tracker  │ │
│  │  MX Resolver        │  │ │ │  RocketReach    │ │ │ │  Billing Calc   │ │
│  │  SMTP Connector     │  │ │ │  Web Scraper    │ │ │ │  Report Gen     │ │
│  │  Spam Trap Checker  │  │ │ │  Data Merger    │ │ │ │  Alert Engine   │ │
│  └─────────────────────┘  │ │ └─────────────────┘ │ │ └─────────────────┘ │
└─────────────┬─────────────┘ └──────────┬──────────┘ └──────────┬──────────┘
              │                          │                       │
              └──────────────────────────┼───────────────────────┘
                                         │
                                         ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                           Worker Layer (RabbitMQ)                           │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │ Verify Queue│  │Enrich Queue │  │ Bulk Queue  │  │  Scheduled Queue    │ │
│  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────────┬──────────┘ │
│         │                │                │                    │            │
│         ▼                ▼                ▼                    ▼            │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │
│  │  Workers    │  │  Workers    │  │  Workers    │  │  Scheduler Workers  │ │
│  │  (N pods)   │  │  (N pods)   │  │  (N pods)   │  │  (cron jobs)        │ │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────────────┘ │
└─────────────────────────────────────────┬───────────────────────────────────┘
                                          │
                                          ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                              Data Layer                                      │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────────────────┐  │
│  │   PostgreSQL    │  │     Redis       │  │      Object Storage         │  │
│  │  - Results      │  │  - Cache        │  │  - CSV uploads              │  │
│  │  - Users        │  │  - Sessions     │  │  - Export files             │  │
│  │  - Audit logs   │  │  - Rate limits  │  │  - Audit archives           │  │
│  └─────────────────┘  └─────────────────┘  └─────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────────┘
                                          │
                                          ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                          External Services                                   │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────────┐  │
│  │  SMTP    │ │   DNS    │ │  Apollo  │ │ Stripe   │ │  CRM APIs        │  │
│  │ Servers  │ │ Resolvers│ │   API    │ │ Billing  │ │ (SF, HS, PD)     │  │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘ └──────────────────┘  │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 Technology Stack

| Layer | Technology | Rationale |
|-------|------------|-----------|
| Frontend | React 18, TypeScript, Vite | Modern, fast, large ecosystem |
| API Gateway | Rust (Warp) | Performance, existing codebase |
| Services | Rust | Consistency, performance |
| Queue | RabbitMQ | Existing, proven, reliable |
| Database | PostgreSQL 15 | Existing, mature, extensible |
| Cache | Redis | Sessions, rate limiting, caching |
| Storage | S3-compatible | File uploads, exports |
| Auth | JWT + API Keys | Flexible, stateless |
| Billing | Stripe | Industry standard |
| Monitoring | Prometheus + Grafana | Observability |
| Logging | Structured JSON → ELK | Searchable logs |

---

## 6. Database Schema Changes

### 6.1 New Tables Summary

| Table | Purpose | Feature |
|-------|---------|---------|
| `organizations` | Multi-tenancy | All |
| `users` | User accounts | Dashboard |
| `email_lists` | List management | 4.1 |
| `email_list_items` | List emails | 4.1 |
| `person_enrichment` | Person data cache | 4.2 |
| `company_enrichment` | Company data cache | 4.2 |
| `api_keys` | API key management | 4.3 |
| `usage_events` | Usage tracking | 4.3 |
| `usage_daily` | Aggregated usage | 4.3 |
| `subscriptions` | Billing state | 4.3 |
| `integrations` | CRM connections | 4.4 |
| `audit_log` | Immutable audit trail | 4.5 |
| `data_retention_policies` | Retention config | 4.5 |
| `data_subject_requests` | GDPR requests | 4.5 |
| `consent_records` | Consent tracking | 4.5 |
| `spam_trap_database` | Known spam traps | 4.6 |
| `verification_schedules` | Re-verification jobs | 4.7 |
| `email_verification_history` | Historical results | 4.7 |
| `domain_verifications` | Domain check cache | 4.8 |
| `sender_reputations` | Reputation data | 4.9 |
| `scoring_rules` | Custom scoring | 4.10 |
| `segments` | List segments | 4.10 |

### 6.2 Migration Strategy

1. **Phase 1**: Add tables without breaking changes
2. **Phase 2**: Migrate existing data to new schema
3. **Phase 3**: Add foreign keys and constraints
4. **Phase 4**: Deprecate old columns/tables

All migrations will be reversible and tested in staging.

---

## 7. API Design

### 7.1 API Versioning

- **v0**: Legacy synchronous API (deprecated but maintained)
- **v1**: Current async worker-based API
- **v2**: New unified API with all features

### 7.2 Authentication

```yaml
# API Key (existing)
Headers:
  x-reacher-secret: rch_xxxxxxxxxxxx

# JWT (new, for dashboard)
Headers:
  Authorization: Bearer eyJhbGciOiJIUzI1NiIs...

# OAuth (new, for CRM integrations)
# Standard OAuth 2.0 flow
```

### 7.3 Rate Limiting

```yaml
# Response headers
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1640995200

# 429 response
{
  "error": "rate_limit_exceeded",
  "message": "Rate limit exceeded. Try again in 60 seconds.",
  "retry_after": 60
}
```

### 7.4 Error Handling

```yaml
# Standard error response
{
  "error": {
    "code": "validation_error",
    "message": "Invalid email format",
    "details": {
      "field": "email",
      "value": "not-an-email",
      "constraint": "email_format"
    }
  },
  "request_id": "req_abc123"
}
```

---

## 8. Security Considerations

### 8.1 Data Security

| Concern | Mitigation |
|---------|------------|
| PII Storage | Encryption at rest (AES-256) |
| Data in Transit | TLS 1.3 only |
| API Key Security | Hashed storage, prefix-only display |
| Password Storage | Argon2id hashing |
| Session Security | Secure, HttpOnly, SameSite cookies |

### 8.2 Access Control

| Resource | Access Control |
|----------|---------------|
| Organization data | Organization membership |
| User data | Self or admin |
| API keys | Organization scope |
| Audit logs | Read-only, admin only |
| Billing | Organization owner only |

### 8.3 Compliance

| Standard | Status |
|----------|--------|
| GDPR | Compliant (with 4.5) |
| CCPA | Compliant (with 4.5) |
| SOC 2 Type II | Target (requires audit) |
| HIPAA | Not applicable |

---

## 9. Implementation Roadmap

### Phase 1: Foundation (Weeks 1-12)

| Week | Deliverable |
|------|-------------|
| 1-2 | Organizations, users, auth infrastructure |
| 3-4 | API key management, usage tracking |
| 5-6 | Dashboard shell, list upload |
| 7-8 | Dashboard results view, export |
| 9-10 | Billing integration (Stripe) |
| 11-12 | Spam trap detection (basic) |

**Milestone**: Dashboard MVP + Usage Analytics

### Phase 2: Enrichment (Weeks 13-20)

| Week | Deliverable |
|------|-------------|
| 13-14 | Enrichment data model, Apollo integration |
| 15-16 | Enrichment API, caching layer |
| 17-18 | Dashboard enrichment display |
| 19-20 | Bulk enrichment, export integration |

**Milestone**: Enrichment MVP

### Phase 3: Enterprise (Weeks 21-28)

| Week | Deliverable |
|------|-------------|
| 21-22 | Audit logging infrastructure |
| 23-24 | GDPR/CCPA compliance features |
| 25-26 | Scheduled re-verification |
| 27-28 | Domain verification |

**Milestone**: Enterprise Compliance

### Phase 4: Integrations (Weeks 29-36)

| Week | Deliverable |
|------|-------------|
| 29-30 | Salesforce integration |
| 31-32 | HubSpot integration |
| 33-34 | Pipedrive integration |
| 35-36 | Integration testing, documentation |

**Milestone**: CRM Integrations

### Phase 5: Advanced (Weeks 37-44)

| Week | Deliverable |
|------|-------------|
| 37-38 | Sender reputation scoring |
| 39-40 | Advanced segmentation |
| 41-42 | Scoring engine |
| 43-44 | Polish, optimization, launch prep |

**Milestone**: Full Platform Launch

---

## 10. Success Metrics

### 10.1 Product Metrics

| Metric | Current | 6-Month Target | 12-Month Target |
|--------|---------|----------------|-----------------|
| Monthly Active Users | ~100 | 1,000 | 10,000 |
| API Requests/Month | ~1M | 10M | 100M |
| Avg. Response Time | 2.5s | 2.0s | 1.5s |
| Uptime | 99.5% | 99.9% | 99.95% |

### 10.2 Business Metrics

| Metric | Current | 6-Month Target | 12-Month Target |
|--------|---------|----------------|-----------------|
| Paying Customers | 0 | 100 | 1,000 |
| Monthly Recurring Revenue | $0 | $10,000 | $100,000 |
| Average Revenue Per User | $0 | $50 | $100 |
| Customer Acquisition Cost | N/A | <$100 | <$50 |
| Net Revenue Retention | N/A | 100% | 120% |

### 10.3 Quality Metrics

| Metric | Target |
|--------|--------|
| Verification Accuracy | >98% |
| Enrichment Match Rate | >70% |
| Customer Support Response | <4 hours |
| Bug Resolution Time | <48 hours |
| NPS Score | >50 |

---

## 11. Open Questions

### 11.1 Technical

1. **Enrichment Provider**: Apollo vs RocketReach vs build own scraper?
2. **Frontend Hosting**: Same domain as API or separate?
3. **Multi-region**: Single region initially or multi-region from start?
4. **Database Scaling**: Partitioning strategy for high-volume tables?

### 11.2 Product

1. **Pricing Model**: Credit-based vs subscription vs hybrid?
2. **Free Tier Limits**: How generous to drive adoption?
3. **Enterprise Features**: Which features are enterprise-only?
4. **Open Source**: Which features remain in open-source core?

### 11.3 Business

1. **Go-to-Market**: Self-serve vs sales-led vs hybrid?
2. **Partnerships**: CRM marketplace partnerships worth pursuing?
3. **Competition**: How to differentiate from Hunter/Apollo?

---

## 12. Appendices

### Appendix A: Competitive Feature Matrix

| Feature | Reacher (Current) | Reacher (Proposed) | Hunter | ZeroBounce | Apollo |
|---------|-------------------|-------------------|--------|------------|--------|
| Email Verification | ✅ | ✅ | ✅ | ✅ | ✅ |
| Bulk Verification | ✅ | ✅ | ✅ | ✅ | ✅ |
| Dashboard | ❌ | ✅ | ✅ | ✅ | ✅ |
| Contact Enrichment | ❌ | ✅ | ✅ | ✅ | ✅ |
| Company Enrichment | ❌ | ✅ | ✅ | ❌ | ✅ |
| CRM Integrations | ❌ | ✅ | ✅ | ✅ | ✅ |
| Spam Trap Detection | ❌ | ✅ | ❌ | ✅ | ❌ |
| Scheduled Verification | ❌ | ✅ | ❌ | ✅ | ❌ |
| Compliance/Audit | ❌ | ✅ | ❌ | ❌ | ❌ |
| Self-Hosted | ✅ | ✅ | ❌ | ❌ | ❌ |
| Open Source | ✅ | ✅ (core) | ❌ | ❌ | ❌ |
| API Access | ✅ | ✅ | ✅ | ✅ | ✅ |

### Appendix B: Glossary

| Term | Definition |
|------|------------|
| **Catch-all** | Email server that accepts all addresses for a domain |
| **DKIM** | DomainKeys Identified Mail - email authentication |
| **DMARC** | Domain-based Message Authentication |
| **DSAR** | Data Subject Access Request (GDPR) |
| **MX Record** | Mail Exchanger DNS record |
| **Pristine Trap** | Spam trap email never used by real person |
| **RDNS** | Reverse DNS lookup |
| **Recycled Trap** | Abandoned email repurposed as spam trap |
| **SPF** | Sender Policy Framework - email authentication |
| **SMTP** | Simple Mail Transfer Protocol |

### Appendix C: References

1. Reacher Core Documentation: `/docs/`
2. Existing API Specification: `/backend/openapi.json`
3. Current Database Schema: `/backend/migrations/`
4. Configuration Reference: `/backend/backend_config.toml`

---

## Document History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 0.1 | 2024-12-03 | Oppulence Engineering | Initial draft |

---

**End of RFC-001**
