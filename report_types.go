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
