package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	server := puppetdbClient.New("https://localhost:8080/", nil)
	response, _ := server.QueryServerTime()
	fmt.Printf("Server Time: %v\n", response.ServerTime)
}
