/*
Project:     SurfaceIntel
File:        ns.go
Description: DNS name server lookup and reporting.
Author:      Daniel Bretschneider, <daniel@bretschneider.cc>
Created:     2026-10-09
Version:     0.1.0
*/

package main

import (
	"context"
	"fmt"
	"net"
	"sort"
	"time"
)

/*
 * lookupNSRecords() retrieves the name server records for a domain.
 * It returns the records or an error if the DNS lookup fails
 * The lookup is limited to five seconds
 */
func lookupNSRecords(domain string) ([]*net.NS, error) {
	// Create Context that limits the DNS lookup to five seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query the domains DNS name server records
	records, err := net.DefaultResolver.LookupNS(ctx, domain)
	if err != nil {
		return nil, err
	}

	// Sort the records alphabetically for consistent output
	sort.Slice(records, func(i, j int) bool {
		return records[i].Host < records[j].Host
	})

	// return the sorted records
	return records, nil
}

/*
 * printNSReport displays the name server records returned by the DNS lookup
 */
func printNSReport(records []*net.NS) {
	// Mark beginning of the name server results section
	fmt.Println("\n[*] NS Records")

	// Explain when the lookup returned no results
	if len(records) == 0 {
		fmt.Println("[-] No NS records found.")
		return
	}

	// Display each discovered name server
	for _, record := range records {
		fmt.Println("-", record.Host)
	}

	// Display the total number of records found
	fmt.Printf("[*] Total NS records: %d\n", len(records))
}
