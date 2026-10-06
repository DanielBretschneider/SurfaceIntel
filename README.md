# SurfaceIntel

> Passive-first OSINT-driven external attack surface intelligence in Go.

SurfaceIntel is a Go-based OSINT and external attack-surface intelligence platform designed to discover, normalize, enrich, and track an organization's publicly observable infrastructure.

## Core idea

Instead of building another generic port scanner, SurfaceIntel turns distributed public observations into an intelligence picture:

```text
Target Domain
     |
     +--> DNS discovery
     +--> Certificate Transparency
     +--> RDAP / registration data
     +--> IP / ASN enrichment
     |
     v
Normalization & Entity Resolution
     |
     v
PostgreSQL
     |
     +--> Technology Detection
     +--> Exposure Scoring
     +--> Historical Change Detection
     |
     v
REST API / Web Dashboard
     |
     v
Optional Local Intelligence Report
```

## MVP

Start deliberately small:

```bash
surfaceintel scan example.com
```

The first version should discover:

- DNS records: A, AAAA, CNAME, MX, NS, TXT
- Certificate Transparency names
- IP addresses
- ASN information
- First-seen / last-seen timestamps

Persist normalized data in PostgreSQL.

## Roadmap

### v0.1
- Go CLI
- DNS collector
- Certificate Transparency collector
- IP resolution
- PostgreSQL
- basic normalization

### v0.2
- RDAP
- ASN enrichment
- concurrent collectors
- HTTP metadata
- TLS metadata

### v0.3
- technology fingerprinting
- Exposure Score
- first-seen / last-seen tracking
- change detection

### v0.4
- REST API
- Go + HTMX dashboard
- Docker Compose
- tests and CI

### v0.5
- threat-intelligence enrichment
- local Qwen/Ollama reporting
- evidence/source tracking
- polished case study

## Suggested architecture

```text
surfaceintel/
+-- cmd/surfaceintel/main.go
+-- internal/
|   +-- dns/
|   +-- crtsh/
|   +-- rdap/
|   +-- asn/
|   +-- enrichment/
|   +-- discovery/
|   +-- database/
|   +-- scoring/
|   +-- api/
+-- web/
+-- migrations/
+-- docs/
+-- Dockerfile
+-- docker-compose.yml
+-- go.mod
+-- README.md
```

## Database model

Start with PostgreSQL rather than Neo4j.

Suggested tables:

```text
organizations
domains
ip_addresses
domain_ips
certificates
certificate_domains
asns
technologies
assets
observations
```

The important concept is an **observation**. Store what was observed, when it was first seen, when it was last seen, and its source. This enables historical intelligence and change detection.

## Concurrent discovery

Go is a natural fit because independent OSINT sources have different latency characteristics.

```go
type DiscoveryResult struct {
    Domains       []Domain
    IPAddresses   []IP
    Certificates  []Certificate
    ASNs          []ASN
}
```

A later `Scan()` implementation can run independent collectors concurrently and normalize their results before persistence.

## Technology detection

After passive discovery, add lightweight HTTP/TLS enrichment:

- HTTP status and redirects
- TLS metadata
- security headers
- server headers
- common framework/application indicators

Keep fingerprints explainable and source-backed.

## Exposure Score

Use an explicitly documented **Exposure Score**, not a vulnerability score.

Potential inputs:

```text
Exposure
Technology
Environment
Age
Certificate state
Infrastructure context
Intelligence indicators
```

The score should explain why an asset was ranked.

## Historical intelligence

Track:

```text
+ new subdomain
+ new IP
- disappeared asset
+ new certificate
~ technology changed
```

This turns a one-time reconnaissance tool into a continuous external attack-surface intelligence system.

## Web UI

A Go + HTMX dashboard is a good initial choice.

Suggested pages:

- Overview
- Assets
- Asset detail
- Certificates
- Infrastructure
- Technologies
- Recent changes
- Intelligence reports

## Local AI

Later, integrate a local Qwen model through Ollama.

The LLM receives structured observations and produces an intelligence report containing:

- Executive Summary
- Key Observations
- Infrastructure Changes
- Interesting Assets
- Confidence Assessment
- Recommended Investigation Areas

Keep the distinction explicit:

```text
OBSERVED
INFERRED
UNKNOWN
```

The LLM should never be treated as the source of truth.

## Responsible use

SurfaceIntel is intended for assets you own or are authorized to assess.

The initial discovery pipeline should be passive-first. Avoid aggressive scanning of third-party infrastructure.

For public demonstrations, use your own domains, a lab environment, deliberately provided targets, or appropriate test domains.

## Portfolio angle

SurfaceIntel demonstrates:

- Cybersecurity
- OSINT methodology
- Threat Intelligence
- Go engineering
- Concurrent programming
- Data modeling
- PostgreSQL
- Web development
- Historical analysis
- Responsible security practices
- Optional local AI integration

## Suggested LinkedIn positioning

> I built SurfaceIntel, a passive-first OSINT platform in Go for mapping and tracking an organization's external attack surface.
>
> It correlates DNS, Certificate Transparency, IP/ASN and HTTP observations into a historical asset model and highlights changes over time.
>
> The goal wasn't to build another port scanner. I wanted to explore how raw OSINT data can be transformed into actionable security intelligence.
>
> Next steps: technology fingerprinting, change detection, and local LLM-assisted intelligence reporting.
