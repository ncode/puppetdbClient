package puppetdbClient

import "time"

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

// FactsPath struct representing a PuppetDB facts path object.
type FactsPath struct {
	Path  []string `json:"path"`
	Type  string   `json:"type"`
	Depth int      `json:"depth"`
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

type Report struct {
	CatalogUuid          string         `json:"catalog_uuid"`
	ReceiveTime          time.Time      `json:"receive_time"`
	Producer             string         `json:"producer"`
	Hash                 string         `json:"hash"`
	TransactionUuid      string         `json:"transaction_uuid"`
	PuppetVersion        string         `json:"puppet_version"`
	Noop                 bool           `json:"noop"`
	CorrectiveChange     string         `json:"corrective_change"`
	Logs                 ReportLogs     `json:"logs"`
	ReportFormat         int            `json:"report_format"`
	StartTime            time.Time      `json:"start_time"`
	ProducerTimestamp    time.Time      `json:"producer_timestamp"`
	Type                 string         `json:"type"`
	CachedCatalogStatus  string         `json:"cached_catalog_status"`
	EndTime              time.Time      `json:"end_time"`
	ResourceEvents       ResourceEvents `json:"resource_events"`
	Status               string         `json:"status"`
	ConfigurationVersion string         `json:"configuration_version"`
	Environment          string         `json:"environment"`
	CodeId               string         `json:"code_id"`
	NoopPending          bool           `json:"noop_pending"`
	Certname             string         `json:"certname"`
	Metrics              Metrics        `json:"metrics"`
	JobId                string         `json:"job_id"`
}

// Metrics struct representing the metrics of a PuppetDB report object.
type Metrics struct {
	Data []MetricData `json:"data"`
	Href string       `json:"href"`
}

// MetricData struct representing the data of a PuppetDB report object.
type MetricData struct {
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Category string  `json:"category"`
}

// ReportLogs struct representing the logs of a PuppetDB report object.
type ReportLogs struct {
	Data []ReportData `json:"data"`
	Href string       `json:"href"`
}

// ReportData struct representing the data of a PuppetDB report object.
type ReportData struct {
	File    string    `json:"file"`
	Line    int       `json:"line"`
	Tags    []string  `json:"tags"`
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Source  string    `json:"source"`
	Message string    `json:"message"`
}

// ResourceEvents struct representing the resource events of a PuppetDB report object.
type ResourceEvents struct {
	Data []ResourceEvent `json:"data"`
	Href string          `json:"href"`
}

// ResourceEvent A representation of a resource event from a report.
type ResourceEvent struct {
	NewValue        string    `json:"new_value"`
	Property        string    `json:"property"`
	Name            string    `json:"name"`
	File            string    `json:"file"`
	OldValue        string    `json:"old_value"`
	Line            int       `json:"line"`
	ResourceType    string    `json:"resource_type"`
	Status          string    `json:"status"`
	ResourceTitle   string    `json:"resource_title"`
	Timestamp       time.Time `json:"timestamp"`
	ContainmentPath []string  `json:"containment_path"`
	Message         string    `json:"message"`
}

// Event A representation of an event from a report.
type Event struct {
	NewValue             string    `json:"new_value"`
	Report               string    `json:"report"`
	CorrectiveChange     bool      `json:"corrective_change"`
	RunStartTime         time.Time `json:"run_start_time"`
	Property             string    `json:"property"`
	Name                 string    `json:"name"`
	File                 string    `json:"file"`
	OldValue             string    `json:"old_value"`
	ContainingClass      string    `json:"containing_class"`
	Line                 int       `json:"line"`
	ResourceType         string    `json:"resource_type"`
	Status               string    `json:"status"`
	ConfigurationVersion string    `json:"configuration_version"`
	ResourceTitle        string    `json:"resource_title"`
	Environment          string    `json:"environment"`
	Timestamp            time.Time `json:"timestamp"`
	RunEndTime           time.Time `json:"run_end_time"`
	ReportReceiveTime    time.Time `json:"report_receive_time"`
	ContainmentPath      []string  `json:"containment_path"`
	Certname             string    `json:"certname"`
	Message              string    `json:"message"`
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
	Certname                string `json:"certname"`
	Deactivated             string `json:"deactivated"`
	CatalogTimestamp        string `json:"catalog_timestamp"`
	FactsTimestamp          string `json:"facts_timestamp"`
	ReportTimestamp         string `json:"report_timestamp"`
	CatalogEnvironment      string `json:"catalog_environment"`
	FactsEnvironment        string `json:"facts_environment"`
	ReportEnvironment       string `json:"report_environment"`
	LatestReportStatus      string `json:"latest_report_status"`
	LatestReportNoop        bool   `json:"latest_report_noop"`
	LatestReportNoopPending bool   `json:"latest_report_noop_pending"`
	LatestReportHash        string `json:"latest_report_hash"`
	LatestReportJobId       string `json:"latest_report_job_id"`
}

// Inventory struct representing a PuppetDB inventory object.
type Inventory struct {
	Certname    string            `json:"certname"`
	Facts       map[string]string `json:"facts"`
	Trusted     map[string]string `json:"trusted"`
	Environment string            `json:"environment"`
	Timestamp   string            `json:"timestamp"`
}
