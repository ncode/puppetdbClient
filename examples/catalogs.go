package main

import (
	"fmt"
	"github.com/ncode/puppetdb-client-go"
)

func main() {
	server := puppetdb.New("http://localhost:8080/")

	// Query catalog
	catResponse, _ := server.QueryCatalogs("foobar")
	fmt.Printf("Catalog Name: %v\n", catResponse.Data.Name)
	fmt.Printf("Catalog Version: %v\n", catResponse.Data.Version)
	fmt.Printf("Catalog Transaction UUID: %v\n", catResponse.Data.TransactionUuid)
}
