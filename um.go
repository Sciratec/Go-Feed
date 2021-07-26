package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type SearchRes struct {
	Results []struct {
		IndexedAt time.Time `json:"indexedAt"`
		Task      struct {
			Visibility string    `json:"visibility"`
			Method     string    `json:"method"`
			Domain     string    `json:"domain"`
			Time       time.Time `json:"time"`
			UUID       string    `json:"uuid"`
			URL        string    `json:"url"`
			Tags       []string  `json:"tags"`
		} `json:"task"`
		Stats struct {
			UniqIPs           int `json:"uniqIPs"`
			ConsoleMsgs       int `json:"consoleMsgs"`
			UniqCountries     int `json:"uniqCountries"`
			DataLength        int `json:"dataLength"`
			EncodedDataLength int `json:"encodedDataLength"`
			Requests          int `json:"requests"`
		} `json:"stats"`
		Page struct {
			Country  string `json:"country"`
			Server   string `json:"server"`
			Domain   string `json:"domain"`
			IP       string `json:"ip"`
			MimeType string `json:"mimeType"`
			Asnname  string `json:"asnname"`
			Asn      string `json:"asn"`
			URL      string `json:"url"`
			Ptr      string `json:"ptr"`
			Status   string `json:"status"`
		} `json:"page"`
		ID         string        `json:"_id"`
		Sort       []interface{} `json:"sort"`
		Result     string        `json:"result"`
		Screenshot string        `json:"screenshot"`
	} `json:"results"`
}

func um(sb []byte) {
	sr := SearchRes{}
	json.Unmarshal(sb, &sr)
	sr.collectUrls()
}

func (s SearchRes) collectUrls() {
	// Printing for now, will use this to look into these results to pivot with
	for counter := range s.Results {
		fmt.Printf("%s\n", s.Results[counter].Result)
	}
}
