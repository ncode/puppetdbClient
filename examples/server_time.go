package main

import (
	"fmt"
	"github.com/ncode/puppetdb-client-go"
)

func main() {
	server := puppetdb.New("http://localhost:8080/")
	response, _ := server.QueryServerTime()
	fmt.Printf("Server Time: %v\n", response.ServerTime)
}
