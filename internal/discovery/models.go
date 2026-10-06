// Copyright © 2026 Daniel Bretschneider
// SPDX-License-Identifier: MIT
//
// Author: Daniel Bretschneider <daniel@bretschneider.cc>
// Purpose: Define data models for network discovery results.
// Created: 2026-10-06
// Last Updated: 2026-10-06
// Package discovery provides types and functions for discovering network resources.
package discovery

import "time"

// Domain represents a discovered domain name with metadata about its origin and lifecycle timestamps.
type Domain struct {
	Name      string
	Source    string
	FirstSeen time.Time
	LastSeen  time.Time
}

// IPAddress represents a discovered IP address with version information and discovery source tracking.
type IPAddress struct {
	IPAddress string
	Version   int
	Source    string
	FirstSeen time.Time
	LastSeen  time.Time
}

// Certificate represents an X.509 certificate with relevant metadata for discovery and tracking purposes.
type Certificate struct {
	SerialNumber string
	Issuer       string
	NotBefore    time.Time
	NotAfter     time.Time
	Domains      []string
}

// DiscoveryResult aggregates all discovered network resources from a single discovery operation.
type DiscoveryResult struct {
	Domains      []Domain
	IPAdresses   []IPAddress
	Certificates []Certificate
}
