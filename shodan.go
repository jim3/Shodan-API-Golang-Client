package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// --------------------------------------------------------

type Response struct {
	IP           string     `json:"ip_str"`
	Country      string     `json:"country_code"`
	Organization string     `json:"org"`
	Tags         []string   `json:"tags"`
	Domains      []string   `json:"domains"`
	Data         []DataItem `json:"data"` // Array of DataItem objects
}

// Host information endpoint `/shodan/host/{ip}`
func (r *Response) LookupIP(ipAddr string) (Response, error) {
	APIKEY := os.Getenv("SHODAN_API_KEY")
	if APIKEY == "" {
		return Response{}, fmt.Errorf("APIKEY environment variable not set")
	}

	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ipAddr, APIKEY)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return Response{}, fmt.Errorf("error unmarshalling json data: %v", err)
	}

	return *r, nil
}

type DataItem struct {
	IP        int64    `json:"ip"`
	Port      int      `json:"port"`
	Transport string   `json:"transport"`
	Hash      int64    `json:"hash"`
	Tags      []string `json:"tags"`
	Cloud     Cloud    `json:"cloud"`
	Location  Location `json:"location"`
}

type Cloud struct {
	Region   string  `json:"region"`
	Service  *string `json:"service"`
	Provider string  `json:"provider"`
}

type Location struct {
	City        string  `json:"city"`
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
		return Search{}, fmt.Errorf("error reading response body: %v", err)
	}

	var result Search
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Search{}, fmt.Errorf("error unmarshalling json data: %v", err)
	}

	return result, nil
}
