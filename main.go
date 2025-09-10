package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Command line flags
	ip := flag.String("iplookup", "", "IP address to look up")
	query := flag.String("search", "", "Host search query")
	flag.Parse()

	// Return host information `/shodan/host/{ip}`
	if *ip != "" {
		var resp Response
		result, err := resp.LookupIP(*ip)
		if err != nil {
			log.Fatalf("IP lookup failed: %v", err)
		}

		// Print Banner Information
		for _, item := range result.Data {
			fmt.Printf("Port: %d, Transport: %s, Location: %s\n", item.Port, item.Transport, item.Location.City)
		}
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
