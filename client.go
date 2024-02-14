package puppetdbClient

import "net/http"

// Client Representation of a PuppetDB server instance.
type Client struct {
	ServerUrl  string
	httpClient *http.Client
}

// New Create a new instance of a Client, which can be used to perform operations on the PuppetDB instance.
func New(baseUrl string, client *http.Client) Client {
	if client == nil {
		client = &http.Client{}
	}
	return Client{baseUrl, client}
}
