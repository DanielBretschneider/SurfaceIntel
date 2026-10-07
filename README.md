# SurfaceIntel

> Passive-first OSINT-driven external attack surface intelligence in Go.

SurfaceIntel is a Go-based OSINT and external attack-surface intelligence platform designed to discover, normalize, enrich, and track an organization's publicly observable infrastructure.

This project is still in progress, currently working on v0.1.

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

Starting small so I don't overcomplicate things.

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

Start with PostgreSQL.

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

The important concept is an **observation**. Goal is to store what was observed, when it was first seen, when it was last seen, and its source. This enables historical intelligence and change detection.

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

After passive discovery, add lightweight HTTP/TLS enrichment - let's see....

- HTTP status and redirects
- TLS metadata
- security headers
- server headers
- common framework/application indicators

Keep fingerprints explainable and source-backed.

## Historical intelligence

Track:

```text
+ new subdomain
+ new IP
- disappeared asset
+ new certificate
~ technology changed
```

This turns a one-time reconnaissance tool into a continuous external attack-surface intelligence system ;)

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

Later, trying to integrate a local Qwen model through Ollama.

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

## Responsible use

SurfaceIntel is intended for assets you own or are authorized to assess.

The initial discovery pipeline should be passive-first. Aggressive scanning of third-party infrastructure should be avoided.
