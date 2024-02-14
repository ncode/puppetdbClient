package puppetdbClient

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	baseUrl := "http://localhost:8081"
	client := &http.Client{}

	server := New(baseUrl, client)

	if server.ServerUrl != baseUrl {
		t.Errorf("Expected ServerUrl to be %s, but got %s", baseUrl, server.ServerUrl)
	}

	if server.httpClient != client {
		t.Errorf("Expected httpClient to be %+v, but got %+v", client, server.httpClient)
	}

	// Test case when client is nil
	server = New(baseUrl, nil)

	if server.ServerUrl != baseUrl {
		t.Errorf("Expected ServerUrl to be %s, but got %s", baseUrl, server.ServerUrl)
	}

	if server.httpClient == nil {
		t.Errorf("Expected httpClient not to be nil")
	}
}
