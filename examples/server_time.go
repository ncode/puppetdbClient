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

	response, err := client.QueryServerTime()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client Time: %v\n", response.ServerTime)
}
