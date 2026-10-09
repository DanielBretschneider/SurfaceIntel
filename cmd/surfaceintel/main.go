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
	"net"
	"os"
	"sort"
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
	fmt.Println("\n[*] Target domain: ", domain)

	// Look up the IP addresses associate with the domain
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("[-] DNS lookup failed: ", err)
		return
	}

	// define address counter for v4 and v6 addresses
	ipv4Count := 0
	ipv6Count := 0

	// Store v4 and v6 address seperately
	var ipv4Addresses []string
	var ipv6Addresses []string

	// Display the IP addresses return by the lookup
	fmt.Printf("\n[*] Found %d IP address(es):", len(ips))
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4Count++
			ipv4Addresses = append(ipv4Addresses, ip.String())
		} else {
			ipv6Count++
			ipv6Addresses = append(ipv6Addresses, ip.String())
		}
	}

	// sort addresses for consistent output
	sort.Strings(ipv4Addresses)
	sort.Strings(ipv6Addresses)

	// Display IPv4 addresses.
	fmt.Println()
	fmt.Printf("[+] IPv4 Addresses (%d):\n", len(ipv4Addresses))
	for _, ip := range ipv4Addresses {
		fmt.Println("-", ip)
	}

	// Display IPv6 addresses.
	fmt.Println()
	fmt.Printf("[+] IPv6 Addresses (%d):\n", len(ipv6Addresses))
	for _, ip := range ipv6Addresses {
		fmt.Println("-", ip)
	}

	// Display a summary of the DNS lookup results.
	fmt.Println()
	fmt.Println("[*] DNS Summary")
	fmt.Println("IPv4 addresses:", ipv4Count)
	fmt.Println("IPv6 addresses:", ipv6Count)
	fmt.Println("Total addresses:", len(ips))
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
