package main

import (
	"fmt"
	"github.com/ncode/puppetdb-client-go"
)

func main() {
	server := puppetdb.New("http://localhost:8080/")

	facts := map[string]string{
		"foo": "bar",
	}

	response, _ := server.ReplaceFacts("foobar", facts)
	fmt.Printf("UUID: %v\n", response.Uuid)
}
