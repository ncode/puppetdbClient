package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	client := puppetdbClient.New("http://localhost:8080/", nil)
	response, _ := client.QueryFactNames()
	fmt.Printf("Fact Names: %v\n", response)
}
