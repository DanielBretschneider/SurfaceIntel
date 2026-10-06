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
	endpoint := "https://crt.sh/" + url.Values{
		"q":      []string{query},  // Search query for domain pattern
		"output": []string{"json"}, // Request response in JSON format
	}.Encode()

	// Send HTTP GET request to crt.sh API
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("request crt.sh: %w", err)
	}
	defer resp.Body.Close()

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

	// Use a map to track already-seen domains for deduplication
	seen := make(map[string]struct{})
	var domains []string

	// Iterate through each certificate record returned by crt.sh
	for _, record := range records {
		// Split potentially multi-value NameValue field (domains separated by newlines)
		for _, value := range strings.Split(record.NameValue, "\n") {
			// Normalize the domain and skip empty results
			domain := normalizeDomain(value)

			if domain == "" {
				continue
			}

			// Skip if already processed (deduplication)
			if _, exists := seen[domain]; exists {
				continue
			}

			// Mark as seen and add to results
			seen[domain] = struct{}{}
			domains = append(domains, domain)
		}
	}

	return domains, nil
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
