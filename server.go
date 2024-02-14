package puppetdb

import "net/http"

// Server Representation of a PuppetDB server instance.
type Server struct {
	BaseUrl string
	Client  *http.Client
}

// New Create a new instance of a Server, which can be used to perform operations on the PuppetDB instance.
func New(baseUrl string, client *http.Client) Server {
	if client == nil {
		client = &http.Client{}
	}
	return Server{baseUrl, client}
}
