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
		"query": {"[\"extract\", [\"certname\", \"latest_report_hash\"], [\"and\", [\"not\", [\"=\", \"latest_report_hash\", null]], [\"=\", \"deactivated\", null]]]"},
		"limit": {"10"},
	}
	queryString := values.Encode()

	response, err := client.QueryNodes(queryString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Nodes: %v\n", response)
}
