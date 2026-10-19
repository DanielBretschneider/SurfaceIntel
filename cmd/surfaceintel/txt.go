/*
Project:     SurfaceIntel
File:        txt.go
Description: DNS TXT record lookup and reporting.
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
 * lookupTXTRecords() retrieves the TXT records for a domain.
 * It returns the records or an error if the DNS lookup fails.
 * The lookup is limited to five seconds.
 */
func lookupTXTRecords(domain string) ([]string, error) {
	// Create a context that limits the DNS lookup to five seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query the domain's TXT records.
	records, err := net.DefaultResolver.LookupTXT(ctx, domain)
	if err != nil {
		return nil, err
	}

	// Sort the records alphabetically for consistent output.
	sort.Strings(records)

	// Return the sorted records.
	return records, nil
}

/*
 * printTXTReport displays the TXT records returned by the DNS lookup.
 */
func printTXTReport(records []string) {
	// Mark the beginning of the TXT results section.
	fmt.Println("\n[*] TXT Records")

	// Explain when the lookup returned no results.
	if len(records) == 0 {
		fmt.Println("[-] No TXT records found.")
		return
	}

	// Display each discovered TXT record.
	for _, record := range records {
		fmt.Println("-", record)
	}

	// Display the total number of records found.
	fmt.Printf("[*] Total TXT records: %d\n", len(records))
}
