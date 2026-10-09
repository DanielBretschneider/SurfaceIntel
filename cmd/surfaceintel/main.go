/*
Project:     SurfaceIntel
File:        main.go
Description: Main entry point for the SurfaceIntel CLI.
Author:      Daniel Bretschneider, <daniel@bretschneider.cc>
Created:     2026-10-09
Version:     0.1.0
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// main is the entry point of the Surfaceintel CLI application.
// Execution starts here when the program is launched.
func main() {
	// Print app name and its purpose to the terminal.
	fmt.Println("SurfaceIntel v0.1.0 - Domain Intelligence")

	// check whether the user supplied a domain argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: surfaceintel <domain>")
		return
	}

	// Read the first argument after the program name
	domain := os.Args[1]

	if !isValidDomain(domain) {
		fmt.Println("[-] Error: invalid domain format.")
		fmt.Println("[-] Please provide a domain seuch as example.com")
		return
	}

	// Display the supplied argument
	fmt.Println("[*] Target domain: ", domain)
}

// isValidDomain checks whether the supplied string looks like a domain
func isValidDomain(domain string) bool {
	// a domain must contain at least one dot
	if !strings.Contains(domain, ".") {
		return false
	}

	// Reject spaces, URL schemes, and paths for now
	// TODO
	if strings.ContainsAny(domain, " /:\\") {
		return false
	}

	// basic checks passed
	return true
}
