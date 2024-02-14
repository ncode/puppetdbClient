package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
	"os"
)

func main() {
	puppetdbURL := os.Getenv("PUPPETDB_URL")
	if puppetdbURL == "" {
		puppetdbURL = "http://localhost:8080/"
	}
	client := puppetdbClient.New(puppetdbURL, nil)

	// Query catalog
	catResponse, err := client.QueryCatalogs("foobar")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Catalog Name: %v\n", catResponse.Data.Name)
	fmt.Printf("Catalog Version: %v\n", catResponse.Data.Version)
	fmt.Printf("Catalog Transaction UUID: %v\n", catResponse.Data.TransactionUuid)
}
