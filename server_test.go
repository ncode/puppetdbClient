package puppetdbClient

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	baseUrl := "http://localhost:8081"
	client := &http.Client{}

	server := New(baseUrl, client)

	if server.BaseUrl != baseUrl {
		t.Errorf("Expected BaseUrl to be %s, but got %s", baseUrl, server.BaseUrl)
	}

	if server.Client != client {
		t.Errorf("Expected Client to be %+v, but got %+v", client, server.Client)
	}

	// Test case when client is nil
	server = New(baseUrl, nil)

	if server.BaseUrl != baseUrl {
		t.Errorf("Expected BaseUrl to be %s, but got %s", baseUrl, server.BaseUrl)
	}

	if server.Client == nil {
		t.Errorf("Expected Client not to be nil")
	}
}
