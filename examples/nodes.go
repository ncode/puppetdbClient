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
		"query": []string{""},
	}
	queryString := values.Encode()

	response, _ := server.QueryNodes(queryString)
	fmt.Printf("Nodes: %v\n", response)
}
