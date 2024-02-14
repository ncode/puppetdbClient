package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	server := puppetdbClient.New("https://localhost:8080/", nil)
	response, _ := server.DeactivateNode("foobar")
	fmt.Printf("UUID: %v\n", response.Uuid)
}
