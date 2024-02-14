package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
	"net/url"
	"os"
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
		"query": []string{""},
	}
	queryString := values.Encode()

	response, err := client.QueryFacts(queryString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fact Names: %v\n", response)
}
