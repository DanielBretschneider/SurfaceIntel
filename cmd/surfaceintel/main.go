/*
Project:     SurfaceIntel
File:        main.go
Description: Main entry point and core functions for the SurfaceIntel CLI.
Author:      Daniel Bretschneider, <daniel@bretschneider.cc>
Created:     2026-10-09
Version:     0.1.0
*/

package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

// main is the entry point of the SurfaceIntel CLI application.
// It coordinates input validation, DNS resolution, and reporting.
func main() {
	// Print the application name and its purpose.
	fmt.Println("SurfaceIntel v0.1.0 - Domain Intelligence")

	// Check whether the user supplied a domain argument.
	if len(os.Args) < 2 {
		fmt.Println("Usage: surfaceintel <domain>")
		return
	}

	// Read the first argument after the program name.
	domain := os.Args[1]

	// Validate the supplied domain before performing any lookups.
	if !isValidDomain(domain) {
		fmt.Println("[-] Error: invalid domain format.")
		fmt.Println("[-] Please provide a domain such as example.com.")
		return
	}

	// Display the target domain.
	fmt.Printf("\n[*] Target domain: %s\n", domain)

	// Resolve the domain to IP addresses.
	ips, err := resolveIPAddresses(domain)
	if err != nil {
		fmt.Println("[-] DNS lookup failed:", err)
		return
	}

	// Display the DNS results and their summary.
	printDNSReport(ips)
}

// isValidDomain checks whether the supplied string passes our basic
// domain format checks. It does not verify whether the domain exists.
func isValidDomain(domain string) bool {
	// A domain must contain at least one dot.
	if !strings.Contains(domain, ".") {
		return false
	}

	// Reject spaces, URL schemes, paths, and other unsupported characters.
	if strings.ContainsAny(domain, " /:\\") {
		return false
	}

	// All basic checks passed.
	return true
}

// resolveIPAddresses looks up the IP addresses associated with a domain.
// It returns the resolved addresses or an error if the lookup fails.
// The lookup is limited to five seconds.
func resolveIPAddresses(domain string) ([]net.IP, error) {
	// Create a context that limits the DNS lookup to five seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Look up the IP addresses associated with the domain.
	ipAddresses, err := net.DefaultResolver.LookupIPAddr(ctx, domain)
	if err != nil {
		return nil, err
	}

	// Convert the lookup results into a slice of IP addresses.
	ips := make([]net.IP, 0, len(ipAddresses))
	for _, address := range ipAddresses {
		ips = append(ips, address.IP)
	}

	// Return the resolved addresses and indicate success.
	return ips, nil
}

// printDNSReport separates, sorts, and displays IPv4 and IPv6 addresses.
// It also prints a summary containing the number of addresses found.
func printDNSReport(ips []net.IP) {
	// Store IPv4 and IPv6 addresses separately.
	var ipv4Addresses []string
	var ipv6Addresses []string

	// Identify and group each IP address.
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4Addresses = append(ipv4Addresses, ip.String())
		} else {
			ipv6Addresses = append(ipv6Addresses, ip.String())
		}
	}

	// Sort the address lists for consistent textual output.
	sort.Strings(ipv4Addresses)
	sort.Strings(ipv6Addresses)

	// Mark the beginning of the DNS results section.
	fmt.Println("\n[*] DNS Results")

	// Display the total number of addresses returned by the lookup.
	fmt.Printf("\n[*] Found %d IP address(es):\n", len(ips))

	// Display IPv4 addresses.
	fmt.Printf("\n[+] IPv4 Addresses (%d):\n", len(ipv4Addresses))
	for _, ip := range ipv4Addresses {
		fmt.Println("-", ip)
	}

	// Display IPv6 addresses.
	fmt.Printf("\n[+] IPv6 Addresses (%d):\n", len(ipv6Addresses))
	for _, ip := range ipv6Addresses {
		fmt.Println("-", ip)
	}

	// Display a summary of the DNS lookup results.
	fmt.Println("\n[*] DNS Summary")
	fmt.Println("IPv4 addresses:", len(ipv4Addresses))
	fmt.Println("IPv6 addresses:", len(ipv6Addresses))
	fmt.Println("Total addresses:", len(ips))
}
