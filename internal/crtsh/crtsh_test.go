// Copyright © 2026 Daniel Bretschneider
// SPDX-License-Identifier: MIT
//
// Author: Daniel Bretschneider <daniel@bretschneider.cc>
// Purpose: Unit tests for domain normalization functionality.
// Created: 2026-10-06
// Last Updated: 2026-10-06
// Location: internal/crtsh/crtsh_test.go
//
// Package crtsh contains unit tests for certificate-based domain discovery functionality.
package crtsh

import "testing"

func TestNormalizeDomain(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "example.com",
			expected: "example.com",
		},
		{
			input:    "WWW.EXAMPLE.COM",
			expected: "www.example.com",
		},
		{
			input:    "*.example.com",
			expected: "example.com",
		},
		{
			input:    "  api.example.com  ",
			expected: "api.example.com",
		},
	}

	for _, tt := range tests {
		result := normalizeDomain(tt.input)

		if result != tt.expected {
			t.Errorf(
				"normalizeDomain(%q) = %q, want %q",
				tt.input,
				result,
				tt.expected,
			)
		}
	}
}
