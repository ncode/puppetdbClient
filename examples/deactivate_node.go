package main

import (
	"fmt"
	"github.com/ncode/puppetdbClient"
)

func main() {
	client := puppetdbClient.New("http://localhost:8080/", nil)
	response, _ := client.DeactivateNode("foobar")
	fmt.Printf("UUID: %v\n", response.Uuid)
}
