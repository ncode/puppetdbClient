package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
	"net/url"
)

func main() {
	client := puppetdbClient.New("http://localhost:8080/", nil)

	// A blank query string - just an example
	var values url.Values
	values = map[string][]string{
		"query": []string{"[\"=\",\"certname\",\"foobar\"]"},
	}
	queryString := values.Encode()

	response, _ := client.QueryReports(queryString)
	fmt.Printf("Reports: %v\n", response)
}
