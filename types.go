package puppetdb

// CatalogWireFormat struct representing the wire format of a PuppetDB catalog object.
type CatalogWireFormat struct {
	Metadata CatalogMetadata `json:"metadata"`
	Data     CatalogData     `json:"data"`
}

// CatalogMetadata struct representing the metadata of a PuppetDB catalog object.
type CatalogMetadata struct {
	ApiVersion int `json:"api_version"`
}

// CatalogData struct representing the data of a PuppetDB catalog object.
type CatalogData struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	TransactionUuid string            `json:"transaction-uuid"`
	Edges           []CatalogEdge     `json:"edges"`
	Resources       []CatalogResource `json:"resources"`
}

// CatalogEdge struct representing a catalog edge.
type CatalogEdge struct {
	Source       CatalogResourceSpec `json:"source"`
	Target       CatalogResourceSpec `json:"target"`
	Relationship string              `json:"relationship"`
}

// CatalogResourceSpec struct representing a catalog resource spec.
type CatalogResourceSpec struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

// CatalogResources A list of catalog resources.
type CatalogResources []CatalogResource

// CatalogResource struct representing a catalog resource.
type CatalogResource struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	//Aliases []string `json:"aliases"`
	Exported   bool              `json:"exported"`
	File       string            `json:"file"`
	Line       int               `json:"line"`
	Tags       []string          `json:"tags"`
	Parameters map[string]string `json:"parameters"`
}

// CommandObject struct representing a PuppetDB command object.
type CommandObject struct {
	Command string      `json:"command"`
	Version int         `json:"version"`
	Payload interface{} `json:"payload"`
}

// CommandResponse struct representing a PuppetDB command response object.
type CommandResponse struct {
	Uuid string `json:"uuid"`
}

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

// ServerTime struct representing the server time.
type ServerTime struct {
	ServerTime string `json:"server-time"`
}

// Version struct representing the version of the PuppetDB server.
type Version struct {
	Version string `json:"version"`
}

// ReportWireFormat A representation of a report in wire format.
type ReportWireFormat struct {
	Certname             string          `json:"certname"`
	PuppetVersion        string          `json:"puppet-version"`
	ReportFormat         int             `json:"report-format"`
	ConfigurationVersion string          `json:"configuration-version"`
	StartTime            string          `json:"start-time"`
	EndTime              string          `json:"end-time"`
	ResourceEvents       []ResourceEvent `json:"resource-events"`
	TransactionUuid      string          `json:"transaction-uuid"`
}

// Report A representation of a report
type Report struct {
	Certname             string `json:"certname"`
	PuppetVersion        string `json:"puppet-version"`
	ReportFormat         int    `json:"report-format"`
	ConfigurationVersion string `json:"configuration-version"`
	StartTime            string `json:"start-time"`
	EndTime              string `json:"end-time"`
	TransactionUuid      string `json:"transaction-uuid"`
	Hash                 string `json:"hash"`
}

// ResourceEvent A representation of a resource event from a report.
type ResourceEvent struct {
	ResourceType    string   `json:"resource-type"`
	ResourceTitle   string   `json:"resource-title"`
	Property        string   `json:"property"`
	Timestamp       string   `json:"timestamp"`
	Status          string   `json:"status"`
	OldValue        string   `json:"old-value"`
	NewValue        string   `json:"new-value"`
	Message         string   `json:"message"`
	File            string   `json:"file"`
	Line            int      `json:"line"`
	ContainmentPath []string `json:"containment-path"`
}

// Event A representation of an event from a report.
type Event struct {
	Certname          string   `json:"certname"`
	Report            string   `json:"report"`
	RunStartTime      string   `json:"run-start-time"`
	RunEndTime        string   `json:"run-end-time"`
	ReportReceiveTime string   `json:"report-receive-time"`
	LatestReport      bool     `json:"latest-report?"`
	ResourceType      string   `json:"resource-type"`
	ResourceTitle     string   `json:"resource-title"`
	Property          string   `json:"property"`
	Timestamp         string   `json:"timestamp"`
	Status            string   `json:"status"`
	OldValue          string   `json:"old-value"`
	NewValue          string   `json:"new-value"`
	Message           string   `json:"message"`
	File              string   `json:"file"`
	Line              int      `json:"line"`
	ContainmentPath   []string `json:"containment-path"`
}

// EventCounts A representation of event counts.
type EventCounts struct {
	SubjectType string `json:"subject-type"`
	Subject     string `json:"subject"`
	Failures    string `json:"failures"`
	Successes   string `json:"successes"`
	Noops       string `json:"noops"`
	Skips       string `json:"skips"`
}

// AggregateEventCounts A representation of aggregate event counts.
type AggregateEventCounts struct {
	Failures  string `json:"failures"`
	Successes string `json:"successes"`
	Noops     string `json:"noops"`
	Skips     string `json:"skips"`
	Total     string `json:"total"`
}

// Node Representation of a PuppetDB node
type Node struct {
	Name             string `json:"name"`
	Deactivated      string `json:"deactivated"`
	CatalogTimestamp string `json:"catalog_timestamp"`
	FactsTimestamp   string `json:"facts_timestamp"`
	ReportTimestamp  string `json:"report_timestamp"`
}
