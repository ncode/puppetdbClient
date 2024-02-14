package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
	"net/url"
)

func main() {
	server := puppetdbClient.New("http://localhost:8080/", nil)

	// A blank query string - just an example
	var values url.Values
	values = map[string][]string{
		"query": []string{"[\"=\",\"certname\",\"foobar\"]"},
	}
	queryString := values.Encode()

	response, _ := server.QueryEvents(queryString)
	fmt.Printf("Events: %v\n", response)
}
