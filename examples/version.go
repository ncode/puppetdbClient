package main

import (
	"fmt"
	"github.com/ncode/puppetdb-client-go"
)

func main() {
	server := puppetdb.New("http://localhost:8080/")
	response, _ := server.QueryVersion()
	fmt.Printf("Version: %v\n", response.Version)
}
