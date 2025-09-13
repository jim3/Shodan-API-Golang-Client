package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HostInfo struct {
	IP           string        `json:"ip_str"`
	Ports        []int         `json:"ports"`
	CountryName  string        `json:"country_name"`
	City         string        `json:"city"`
	Region       string        `json:"region_code"`
	Domains      []string      `json:"domains"`
	HostNames    []string      `json:"hostnames"`
	Organization string        `json:"org"`
	ISP          string        `json:"isp"`
	LastUpdated  string        `json:"last_update"`
	Tags         []string      `json:"tags"`
	Data         []HostDetails `json:"data"` // specific service information
}

// Host information endpoint `/shodan/host/{ip}`
func (r *HostInfo) LookupIP(ipAddr string) error {
	APIKEY := os.Getenv("SHODAN_API_KEY")
	if APIKEY == "" {
		return fmt.Errorf("APIKEY environment variable not set")
	}

	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ipAddr, APIKEY)
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading HostInfo body: %v", err)
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return fmt.Errorf("error unmarshalling json data: %v", err)
	}
	return nil
}

type HostDetails struct {
	IP        int64    `json:"ip"`
	Port      int      `json:"port"`      // Port number
	Transport string   `json:"transport"` // Transport protocol (e.g., "tcp", "udp")
	Hash      int64    `json:"hash"`
	Tags      []string `json:"tags"`
	Cloud     Cloud    `json:"cloud"`
	Location  Location `json:"location"` // Location
}

type Cloud struct {
	Region   string  `json:"region"`
	Service  *string `json:"service"`
	Provider string  `json:"provider"`
}

type Location struct {
	City        string  `json:"city"` // Location.City
	RegionCode  string  `json:"region_code"`
	AreaCode    *string `json:"area_code"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
}

// --------------------------------------------------------

type Search struct {
	Matches []QueryResult `json:"matches"`
}

type QueryResult struct {
	Port int `json:"port"`
	IP   int `json:"ip"`
}

// search query endpoint `shodan/host/search`
func (s *Search) HostSearch(q string) (Search, error) {
	APIKEY := os.Getenv("SHODAN_API_KEY")
	if APIKEY == "" {
		return Search{}, fmt.Errorf("APIKEY environment variable not set")
	}

	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&query=%s", APIKEY, q)
	res, err := http.Get(URL)
	if err != nil {
		return Search{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Search{}, fmt.Errorf("error reading HostInfo body: %v", err)
	}

	var result Search
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Search{}, fmt.Errorf("error unmarshalling json data: %v", err)
	}

	return result, nil
}
