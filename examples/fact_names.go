package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	server := puppetdbClient.New("http://localhost:8080/", nil)
	response, _ := server.QueryFactNames()
	fmt.Printf("Fact Names: %v\n", response)
}
