/*
Project:     SurfaceIntel
File:        main_test.go
Description: Unit tests for domain validation.
Author:      Daniel Bretschneider
Created:     2026-10-09
Version:     0.1.0
Notes: 		 Run in project root: go test ./...
*/

package main

import "testing"

// TestIsValidDomain checks the basic domain validation rules.
func TestIsValidDomain(t *testing.T) {
	// Define inputs and their expected results.
	tests := []struct {
		name   string
		domain string
		want   bool
	}{
		{
			name:   "valid domain",
			domain: "example.com",
			want:   true,
		},
		{
			name:   "valid subdomain",
			domain: "sub.example.org",
			want:   true,
		},
		{
			name:   "missing dot",
			domain: "banana",
			want:   false,
		},
		{
			name:   "URL scheme",
			domain: "https://example.com",
			want:   false,
		},
		{
			name:   "domain with path",
			domain: "example.com/path",
			want:   false,
		},
		{
			name:   "domain with spaces",
			domain: "example .com",
			want:   false,
		}, // next cases should fail
		{
			name:   "domain starts with a dot",
			domain: ".com",
			want:   false,
		},
		{
			name:   "domain contains consecutive dots",
			domain: "example..com",
			want:   false,
		},
		{
			name:   "domain label starts with a hyphen",
			domain: "-example.com",
			want:   false,
		},
	}

	// Run each test case independently.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isValidDomain(test.domain)

			// Compare the actual result with the expected result.
			if got != test.want {
				t.Errorf(
					"isValidDomain(%q) = %t; want %t",
					test.domain,
					got,
					test.want,
				)
			}
		})
	}
}
