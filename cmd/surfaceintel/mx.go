/*
Project:     SurfaceIntel
File:        mx.go
Description: DNS MX record lookup and reporting.
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
 * lookupMXRecords() retrieves the mail exchange records for a domain
 * It returns the records or an error if the DNS lookup fails
 * The lookup is limited to five seconds
 */
func lookupMXRecords(domain string) ([]*net.MX, error) {
	// Create a context to limit DNS lookup to five seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query domain's MX records
	records, err := net.DefaultResolver.LookupMX(ctx, domain)
	if err != nil {
		return nil, err
	}

	// Sort found records, with lower values preferred by mail servers
	sort.Slice(records, func(i, j int) bool {
		if records[i].Pref == records[j].Pref {
			return records[i].Host < records[j].Host
		}
		return records[i].Pref < records[j].Pref
	})

	// Return the sorted records.
	return records, nil
}

/*
 * printMXReport() prints the sorted MX records returned by DNS lookup
 */
func printMXReport(records []*net.MX) {
	// Mark the beginning of the MX results section
	fmt.Println("\n[*] MX Records")

	// Explain when the lookup returned no records
	if len(records) == 0 {
		fmt.Println("[-] No MX records found.")
		return
	}

	// Display each mail server and its preference value
	for _, record := range records {
		fmt.Printf("- Host: %s | Preference: %d\n", record.Host, record.Pref)
	}

	// Display the total number of records found
	fmt.Printf("[*] Total MX records: %d\n", len(records))
}
