package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Command line flags
	ip := flag.String("iplookup", "", "IP address to look up")
	query := flag.String("search", "", "Host search query")
	// TODO: Add more flags as needed
	flag.Parse()

	// Return host information using the `-iplookup` flag
	if *ip != "" {
		// instance of the response
		var resp HostInfo

		err := resp.LookupIP(*ip)
		if err != nil {
			log.Fatalf("IP lookup failed: %v", err)
		}
		// Print out basic host info in a pretty table-like format
		fmt.Printf("%-15s %-20s %-20s\n", "IP", "City", "Ports")
		fmt.Printf("%-15s %-20s %-20s\n",
			resp.IP, resp.City, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(resp.Ports)), ", "), "[]"))
	}

	// Return query results `shodan/host/search`
	if *query != "" {
		var s Search
		r, err := s.HostSearch(*query)
		if err != nil {
			log.Fatalf("Host search failed: %v", err)
		}
		fmt.Println(r.Matches)

		for i, v := range r.Matches {
			fmt.Println(v.Port)
			if i >= 4 {
				break
			}
		}
	}
}
