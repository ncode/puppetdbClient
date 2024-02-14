package puppetdbClient

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
	TransactionUuid string            `json:"transaction_uuid"`
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
	Certname string `json:"certname"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

// FactContent struct representing a PuppetDB fact content object.
type FactContent struct {
	Certname    string   `json:"certname"`
	Environment string   `json:"environment"`
	Name        string   `json:"name"`
	Value       string   `json:"value"`
	Path        []string `json:"path"`
}

// ServerTime struct representing the server time.
type ServerTime struct {
	ServerTime string `json:"server_time"`
}

// Version struct representing the version of the PuppetDB server.
type Version struct {
	Version string `json:"version"`
}

// ReportWireFormat A representation of a report in wire format.
type ReportWireFormat struct {
	Certname             string          `json:"certname"`
	PuppetVersion        string          `json:"puppet_version"`
	ReportFormat         int             `json:"report_format"`
	ConfigurationVersion string          `json:"configuration_version"`
	StartTime            string          `json:"start_time"`
	EndTime              string          `json:"end_time"`
	ResourceEvents       []ResourceEvent `json:"resource_events"`
	TransactionUuid      string          `json:"transaction_uuid"`
}

// Report A representation of a report
type Report struct {
	Certname             string `json:"certname"`
	PuppetVersion        string `json:"puppet_version"`
	ReportFormat         int    `json:"report_format"`
	ConfigurationVersion string `json:"configuration_version"`
	StartTime            string `json:"start_time"`
	EndTime              string `json:"end_time"`
	TransactionUuid      string `json:"transaction_uuid"`
	Hash                 string `json:"hash"`
}

// ResourceEvent A representation of a resource event from a report.
type ResourceEvent struct {
	ResourceType    string   `json:"resource_type"`
	ResourceTitle   string   `json:"resource_title"`
	Property        string   `json:"property"`
	Timestamp       string   `json:"timestamp"`
	Status          string   `json:"status"`
	OldValue        string   `json:"old_value"`
	NewValue        string   `json:"new_value"`
	Message         string   `json:"message"`
	File            string   `json:"file"`
	Line            int      `json:"line"`
	ContainmentPath []string `json:"containment_path"`
}

// Event A representation of an event from a report.
type Event struct {
	Certname          string   `json:"certname"`
	Report            string   `json:"report"`
	RunStartTime      string   `json:"run_start_time"`
	RunEndTime        string   `json:"run_end_time"`
	ReportReceiveTime string   `json:"report_receive_time"`
	LatestReport      bool     `json:"latest_report?"`
	ResourceType      string   `json:"resource_type"`
	ResourceTitle     string   `json:"resource_title"`
	Property          string   `json:"property"`
	Timestamp         string   `json:"timestamp"`
	Status            string   `json:"status"`
	OldValue          string   `json:"old_value"`
	NewValue          string   `json:"new_value"`
	Message           string   `json:"message"`
	File              string   `json:"file"`
	Line              int      `json:"line"`
	ContainmentPath   []string `json:"containment_path"`
}

// EventCounts A representation of event counts.
type EventCounts struct {
	SubjectType string       `json:"subject_type"`
	Subject     SubjectField `json:"subject"`
	Failures    int          `json:"failures"`
	Successes   int          `json:"successes"`
	Noops       int          `json:"noops"`
	Skips       int          `json:"skips"`
}

type SubjectField struct {
	Title string `json:"title"`
}

// AggregateEventCounts A representation of aggregate event counts.
type AggregateEventCounts struct {
	Failures  string `json:"failures"`
	Successes string `json:"successes"`
	Noops     string `json:"noops"`
	Skips     string `json:"skips"`
	Total     string `json:"total"`
}

// Node Representation of a PuppetDB node.
type Node struct {
	Name             string `json:"name"`
	Deactivated      string `json:"deactivated"`
	CatalogTimestamp string `json:"catalog_timestamp"`
	FactsTimestamp   string `json:"facts_timestamp"`
	ReportTimestamp  string `json:"report_timestamp"`
}

// Inventory struct representing a PuppetDB inventory object.
type Inventory struct {
	Certname    string            `json:"certname"`
	Facts       map[string]string `json:"facts"`
	Trusted     map[string]string `json:"trusted"`
	Environment string            `json:"environment"`
	Timestamp   string            `json:"timestamp"`
}
