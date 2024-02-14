package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ncode/puppetdbClient"
)

func main() {
	puppetdbURL := os.Getenv("PUPPETDB_URL")
	if puppetdbURL == "" {
		puppetdbURL = "http://localhost:8080/"
	}
	client := puppetdbClient.New(puppetdbURL, nil)

	// A blank query string - just an example
	var values url.Values
	values = map[string][]string{
		"query":        {"[\"=\",\"latest_report?\",\"true\"]"},
		"summarize_by": {"certname"},
	}
	queryString := values.Encode()

	response, err := client.QueryEventCounts(queryString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Event Counts: %v\n", response)
}
