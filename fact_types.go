package puppetdb

// FactsWireFormat struct representing the wire format of a PuppetDB facts object.
type FactsWireFormat struct {
	// Certificate name of node to replace facts for
	Name string `json:"name"`
	// A map of fact key/value pairs
	Values map[string]string `json:"values"`
}

// Fact struct representing a PuppetDB fact object.
type Fact struct {
	CertName string `json:"certname"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

// FactContent struct representing a PuppetDB fact content object.
type FactContent struct {
	CertName    string   `json:"certname"`
	Environment string   `json:"environment"`
	Name        string   `json:"name"`
	Value       string   `json:"value"`
	Path        []string `json:"path"`
}
