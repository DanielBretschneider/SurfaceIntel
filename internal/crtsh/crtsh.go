// Copyright © 2026 Daniel Bretschneider
// SPDX-License-Identifier: MIT
//
// Author: Daniel Bretschneider <daniel@bretschneider.cc>
// Purpose: Interface with crt.sh API for certificate-based domain discovery.
// Created: 2026-10-06
// Last Updated: 2026-10-06
// Location: internal/crtsh/crtsh.go
//
// Package crtsh provides functionality to discover domain names by querying
// the crt.sh Certificate Transparency log API. This is commonly used in
// reconnaissance to find subdomains associated with a target domain.
//
// # Usage
//
// Use the Discover function to query crt.sh for domains related to a target:
//
//	domains, err := crtsh.Discover("example.com")
//	if err != nil {
//	    log.Fatalf("Discovery failed: %v", err)
//	}
//
// The returned slice contains unique domain names found in certificate data.
package crtsh

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// certificateRecord represents a single domain name record returned by the crt.sh API.
type certificateRecord struct {
	NameValue string `json:"name_value"`
}

// Discover queries the crt.sh Certificate Transparency log API to find all domain names
// associated with the given target domain. It searches for certificates containing the
// domain as part of their subject or SAN (Subject Alternative Name) entries.
//
// Parameters:
//   - domain: The target domain to search for (e.g., "example.com")
//
// Returns:
//   - []string: A slice of unique domain names discovered from crt.sh, including subdomains
//   - error: An error if the API request fails or response cannot be parsed
//
// Example:
//
//	domains, err := Discover("example.com")
//	// Returns: ["example.com", "www.example.com", "api.example.com"]
//
// The function performs the following steps:
// 1. Constructs a crt.sh query URL for wildcard domain matching
// 2. Fetches certificate data from crt.sh in JSON format
// 3. Parses and deduplicates returned domain names
// 4. Normalizes domains by trimming whitespace, lowercasing, and removing wildcards
func Discover(domain string) ([]string, error) {
	// Construct wildcard query pattern to match all subdomains (e.g., "%.example.com")
	query := fmt.Sprintf("%%.%s", domain)

	// Build crt.sh API endpoint with query parameters for JSON output
	endpoint := "https://crt.sh/?" + url.Values{
		"q":      []string{query},  // Search query for domain pattern
		"output": []string{"json"}, // Request response in JSON format
	}.Encode()

	// Create Client to build request for crt.sh (problem with http code 429)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Start creating request header
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create crt.sh request: %w", err)
	}

	// Set Request Header for crt.sh
	req.Header.Set("User-Agent", "SurfaceIntel/0.1.0")

	// Send HTTP GET request to crt.sh API
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request crt.sh: %w", err)
	}
	defer resp.Body.Close()

	// Check if too many requests status
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("crt.sh rate limit reached (HTTP 429")
	}

	// Validate HTTP response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("crt.sh returned HTTP %d", resp.StatusCode)
	}

	// Initialize slice to hold parsed certificate records
	var records []certificateRecord

	// Decode JSON response into certificate record slice
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		return nil, fmt.Errorf("decode crt.sh response: %w", err)
	}

	return parseRecords(records, domain), nil
}

func parseRecords(records []certificateRecord, target string) []string {
	target = normalizeDomain(target)

	seen := make(map[string]struct{})
	var domains []string

	for _, record := range records {
		for _, value := range strings.Split(record.NameValue, "\n") {
			domain := normalizeDomain(value)

			if domain == "" {
				continue
			}

			if !belongsToTarget(domain, target) {
				continue
			}

			if _, exists := seen[domain]; exists {
				continue
			}

			seen[domain] = struct{}{}
			domains = append(domains, domain)
		}
	}

	return domains
}

// normalizeDomain processes a raw domain string from crt.sh into a standardized format.
// It performs the following transformations:
// 1. Trims leading and trailing whitespace
// 2. Converts to lowercase for consistency
// 3. Removes wildcard prefixes (*.) for cleaner results
//
// Parameters:
//   - value: The raw domain string from crt.sh API response
//
// Returns:
//   - string: Normalized domain name ready for use, or empty string if invalid
//
// Examples:
//
//	normalizeDomain("  Example.COM  ") // Returns: "example.com"
//	normalizeDomain("*.api.example.com") // Returns: "api.example.com"
//	normalizeDomain("") // Returns: ""
func normalizeDomain(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	value = strings.TrimPrefix(value, "*.")

	return value
}

// This is necessary so "example.com.evil.com" which contains
// "example.com" is not counted to "example.com". It's not
// a valid subdomain. This func checks if that's the case.
// TODO: better comment
func belongsToTarget(domain, target string) bool {
	return domain == target ||
		strings.HasSuffix(domain, "."+target)
}
