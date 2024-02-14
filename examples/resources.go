package main

import (
	"fmt"
	"github.com/ncode/puppetdb-client-go"
	"net/url"
)

func main() {
	server := puppetdb.New("http://localhost:8080/")

	// A blank query string - just an example
        var values url.Values
        values = map[string][]string{
                "query":[]string{""},
        }
        queryString := values.Encode()

	response, _ := server.QueryResources(queryString)
	fmt.Printf("Resources: %v\n", response)
}
