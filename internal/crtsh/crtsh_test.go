package crtsh

import (
	"reflect"
	"testing"
)

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

func TestBelongsToTarget(t *testing.T) {
	tests := []struct {
		domain string
		target string
		want   bool
	}{
		{"example.com", "example.com", true},
		{"www.example.com", "example.com", true},
		{"api.example.com", "example.com", true},
		{"dev.api.example.com", "example.com", true},

		{"example.com.evil.com", "example.com", false},
		{"evil-example.com", "example.com", false},
		{"example.org", "example.com", false},
	}

	for _, tt := range tests {
		result := belongsToTarget(tt.domain, tt.target)

		if result != tt.want {
			t.Errorf(
				"belongsToTarget(%q, %q) = %v, want %v",
				tt.domain,
				tt.target,
				result,
			)
		}
	}
}

func TestParseRecords(t *testing.T) {
	records := []certificateRecord{
		{
			NameValue: "example.com\nwww.example.com\n*.api.example.com",
		},
		{
			NameValue: "example.com\nwww.example.com",
		},
		{
			NameValue: "example.com.evil.com",
		},
		{
			NameValue: "evil-example.com",
		},
	}

	expected := []string{
		"example.com",
		"www.example.com",
		"api.example.com",
	}

	result := parseRecords(records, "example.com")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"parseRecords() = %v, want %v",
			result,
			expected,
		)
	}
}
