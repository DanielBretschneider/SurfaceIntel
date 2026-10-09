# SurfaceIntel

**Domain Intelligence & Passive OSINT CLI**

SurfaceIntel is a command-line tool written in Go, designed to collect and organize publicly available information about internet domains.

The project is being developed on CachyOS Linux, with a focus on learning Go, clean software architecture, and practical OSINT techniques.

The long-term goal is to build a modular, reliable, and easy-to-use tool for domain intelligence and security research.

## Project Status

* **Version:** 0.1.0
* **Status:** Early development
* **Language:** Go
* **Platform:** Linux (CachyOS)

### Currently Implemented

* [x] Initial Go project structure
* [x] Basic executable entry point
* [x] Documented main function
* [x] Initial terminal output

### Planned Features

* [ ] Command-line argument parsing
* [ ] Domain input validation
* [ ] DNS record lookups (A, AAAA, MX, NS, TXT, CAA)
* [ ] Domain registration information via RDAP
* [ ] Passive subdomain discovery via Certificate Transparency
* [ ] HTTP response metadata and security headers
* [ ] Structured JSON reports
* [ ] Error handling and configurable timeouts
* [ ] Unit tests and expanded documentation

*Planned features have not yet been implemented.*

## Technology Stack

| Technology          | Purpose                   |
| ------------------- | ------------------------- |
| Go                  | Main programming language |
| Git                 | Version control           |
| CachyOS Linux       | Development environment   |
| Go standard library | Initial implementation    |

External dependencies will be introduced only when they provide a clear benefit.

## Project Structure

```text
SurfaceIntel/
├── cmd/
│   └── surfaceintel/
│       └── main.go
├── go.mod
├── README.md
├── README.txt
├── .gitignore
└── .git/
```

The project structure will evolve as additional modules and features are introduced.

## Getting Started

### Prerequisites

* Go
* Git

### Run from Source

Clone the repository if you do not already have it locally:

```bash
git clone <repository-url>
cd SurfaceIntel
```

Run the application from the project root:

```bash
go run ./cmd/surfaceintel
```

### Build an Executable

```bash
go build -o surfaceintel ./cmd/surfaceintel
```

Run the compiled application:

```bash
./surfaceintel
```

The current version prints the application name and purpose to the terminal. Domain intelligence features will be added incrementally.

## Development Principles

SurfaceIntel is being developed with the following principles:

* **Learn by building:** Implement small, understandable features.
* **Keep it modular:** Separate responsibilities as the application grows.
* **Document the code:** Explain important functions and design decisions.
* **Prefer simplicity:** Use Go's standard library where practical.
* **Handle errors explicitly:** Make failures understandable and predictable.
* **Test incrementally:** Add tests as functionality is introduced.
* **Support automation:** Provide structured output alongside human-readable results.
* **Respect service limits:** Follow the terms and rate limits of public data sources.
* **Use responsibly:** Perform active checks only against systems for which authorization exists.

## Roadmap

The development roadmap is divided into small milestones.

1. **Project foundations** — Initialize the Go module and establish the application entry point.
2. **CLI arguments** — Accept a domain as a command-line argument.
3. **Domain validation** — Validate and normalize domain input.
4. **DNS intelligence** — Retrieve and display DNS records.
5. **HTTP metadata** — Inspect basic website response information.
6. **Registration intelligence** — Retrieve available RDAP data.
7. **Passive discovery** — Identify domain names from public Certificate Transparency data.
8. **Reporting** — Export results as JSON.
9. **Quality and release** — Add tests, improve documentation, and prepare release builds.

Each milestone will be implemented and tested incrementally.

## Disclaimer

SurfaceIntel is intended for legitimate OSINT, security research, and educational purposes.

Users are responsible for complying with applicable laws, service terms, and authorization requirements when assessing domains and infrastructure.

## Author and Contributions

**Author:** Daniel Bretschneider and local LLM (Ollama Qewn 3.5 22B)

SurfaceIntel is an educational project under active development. Features, documentation, and architecture may change as the project evolves.

## License

MIT
