package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	server := puppetdbClient.New("https://localhost:8080/", nil)

	// Query catalog
	catResponse, _ := server.QueryCatalogs("foobar")
	fmt.Printf("Catalog Name: %v\n", catResponse.Data.Name)
	fmt.Printf("Catalog Version: %v\n", catResponse.Data.Version)
	fmt.Printf("Catalog Transaction UUID: %v\n", catResponse.Data.TransactionUuid)
}
