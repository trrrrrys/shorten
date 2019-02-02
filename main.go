package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	url string
)

// BITLYURL is bitly API base url.
const BITLYURL = "https://api-ssl.bitly.com/v3/shorten"

type bitlyResponse struct {
	Data Data `json:"data,required"`
}

type Data struct {
	URL string `json:"url,required"`
}

func run() int {
	bitlyURL := fmt.Sprintf("%s?access_token=%s&longUrl=%s", BITLYURL, os.Getenv("BITLY_TOKEN"), url)
	resp, err := http.Get(bitlyURL)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return 1
	}
	defer resp.Body.Close()
	var data bitlyResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	fmt.Println(data.Data.URL)
	return 0
}

func main() {
	url = os.Args[1]
	os.Exit(run())
}
