package main

import (
	"io/ioutil"
	"net/http"
)

func getSearchRes() []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://urlscan.io/api/v1/search/?q=date:>now-1h", nil)
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "$APIKEY") // Provide own key

	resp, err := client.Do(req)
	checkErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return body
}
