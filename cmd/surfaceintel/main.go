// Copyright © 2026 Daniel Bretschneider
// SPDX-License-Identifier: MIT
//
// Author: Daniel Bretschneider <daniel@bretschneider.cc>
// Purpose: Provide utilities for Intel Surface device interactions.
// Created: 2026-10-06
// Last Updated: 2026-10-06
// Location: cmd/surfaceintel/main.go
//
// Package surfaceintel provides the main entry point for the SurfaceIntel tool,
// which facilitates communication and management of Intel Surface devices.
package main

import (
	"SurfaceIntel/internal/crtsh"
	"fmt"
	"os"
)

// version holds the current semantic version of the application.
// Format: MAJOR.MINOR.PATCH (https://semver.org/)
const version = "0.1.0"

// main is the entry point of the SurfaceIntel application.
// It initializes and executes the primary program logic for interacting
// with Intel Surface devices. As a simple hello-world starter, it outputs
// the current version and confirmation message.
//
// This function serves as the starting point when running:
//
//	go run ./cmd/surfaceintel/
//
// or executing the compiled binary directly.
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: surfaceintel scan <domain>")
		os.Exit(1)
	}

	command := os.Args[1]
	target := os.Args[2]

	if command != "scan" {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	fmt.Printf("SurfaceIntel v%s\n", version)
	fmt.Printf("Target: %s\n\n", target)

	fmt.Println("[Certificate Transparency]")

	domains, err := crtsh.Discover(target)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, domain := range domains {
		fmt.Printf("  %s\n", domain)
	}

	fmt.Printf("\n[Summary]\n")
	fmt.Printf("  Domains discovered: %d\n", len(domains))
}
