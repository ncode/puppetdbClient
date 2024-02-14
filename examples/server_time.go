package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	server := puppetdbClient.New("http://localhost:8080/", nil)
	response, _ := server.QueryServerTime()
	fmt.Printf("Client Time: %v\n", response.ServerTime)
}
