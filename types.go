package puppetdb

/*
Wire format representation of a catalog.

You probably want to take a look at the NewCatalogWireFormat function, as this
is the suggested way to create a new catalog wire format data structure from
scratch.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html
*/
type CatalogWireFormat struct {
	// Metadata for this catalog
	Metadata CatalogMetadata `json:"metadata"`
	// Data for this catalog
	Data CatalogData `json:"data"`
}

/*
Catalog metadata struct.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html#main-data-type-catalog
*/
type CatalogMetadata struct {
	// Catalog data API version
	ApiVersion int `json:"api_version"`
}

/*
Data for a catalog.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html#main-data-type-catalog
*/
type CatalogData struct {
	// Certificate name owning the catalog to be replaced
	Name string `json:"name"`
	// Version of the catalog
	Version string `json:"version"`
	// Unique identifier provided by client to marry catalogs with reports and other (future) objects
	TransactionUuid string `json:"transaction-uuid"`
	// Edges represented in this catalog
	Edges []CatalogEdge `json:"edges"`
	// Resources represented in this catalog
	Resources []CatalogResource `json:"resources"`
}

/*
A representation of an edge inside a catalog.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html#data-type-edge
*/
type CatalogEdge struct {
	// Source resource spec for this edge
	Source CatalogResourceSpec `json:"source"`
	// Target resource spec for this edge
	Target CatalogResourceSpec `json:"target"`
	// Relationship type
	Relationship string `json:"relationship"`
}

/*
This struct represents a catalog resource reference for use in edges.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html#data-type-resource-spec
*/
type CatalogResourceSpec struct {
	// The type of a catalog resource
	Type string `json:"type"`
	// The title of a catalog resource
	Title string `json:"title"`
}

/*
A collection of catalog resources
*/
type CatalogResources []CatalogResource

/*
A catalog resource.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/wire_format/catalog_format.html#data-type-resource
*/
type CatalogResource struct {
	// The type of a catalog resource
	Type string `json:"type"`
	// The title of a catalog resource
	Title string `json:"title"`
	// Aliases for this resource
	//Aliases []string `json:"aliases"`
	// Exported status
	Exported bool `json:"exported"`
	// Source file this resource appears in
	File string `json:"file"`
	// Line in the file this resource appears in
	Line int `json:"line"`
	// All tags applied to this resource
	Tags []string `json:"tags"`
	// A map containing a list of key/value pairs for each parameter of this resource
	Parameters map[string]string `json:"parameters"`
}

/*
Create a new catalog
*/
func NewCatalogWireFormat() CatalogWireFormat {
	metadata := CatalogMetadata{0}
	data := CatalogData{"", "", "", nil, nil}
	return CatalogWireFormat{metadata, data}
}
package puppetdb

/*
Top level struct representing a PuppetDB commands payload object.

See here for more details on the protocol: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
type CommandObject struct {
	// Command name, such as 'replace facts' or 'deactivate node'
	Command string `json:"command"`
	// Command version as an integer
	Version int `json:"version"`
	// Command payload, may contain different data types depending on command
	Payload interface{} `json:"payload"`
}

/*
Response to a commands submission request.

This struct contains the fields that are returned when a command was
successfully submitted. This does not indicate the command was processed,
just an acknowledgement it was received and will be processed in the future.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html#command-submission
*/
type CommandResponse struct {
	// A UUID returned by the server uniquely identifying a command submission
	Uuid string `json:"uuid"`
}
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
package puppetdb

/*
Response to servertime query end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/server-time.html#get-v3server-time
*/
type ServerTime struct {
	ServerTime string `json:"server-time"`
}

/*
Response to version query end-point.

More details here: http://docs.puppetlabs.com/puppetdb/latest/api/query/v3/version.html#get-v3version
*/
type Version struct {
	Version string `json:"version"`
}
package puppetdb

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
package puppetdb

// Node Representation of a PuppetDB node
type Node struct {
	Name             string `json:"name"`
	Deactivated      string `json:"deactivated"`
	CatalogTimestamp string `json:"catalog_timestamp"`
	FactsTimestamp   string `json:"facts_timestamp"`
	ReportTimestamp  string `json:"report_timestamp"`
}
