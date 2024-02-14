package puppetdbClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Query will query the PuppetDB instance with the given URL.
func (server *Client) Query(url string) ([]byte, error) {
	baseUrl := server.ServerUrl
	queryUrl := strings.Join([]string{baseUrl, url}, "")

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := server.httpClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed status code: %d body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// QueryVersion will query the PuppetDB instance version end-point.
func (server *Client) QueryVersion() (version *Version, err error) {
	body, err := server.Query("pdb/meta/v1/version")
	if err != nil {
		return version, err
	}

	version = &Version{}
	err = json.Unmarshal(body, version)
	if err != nil {
		return nil, err
	}

	return version, err
}

// QueryServerTime will query the PuppetDB instance server-time end-point.
func (server *Client) QueryServerTime() (serverTime *ServerTime, err error) {
	body, err := server.Query("pdb/meta/v1/server-time")
	if err != nil {
		return serverTime, err
	}

	serverTime = &ServerTime{}
	err = json.Unmarshal(body, serverTime)
	if err != nil {
		return nil, err
	}

	return serverTime, err
}

// QueryFactNames will query the PuppetDB instance fact-names end-point.S
func (server *Client) QueryFactNames() (factNames []string, err error) {
	body, err := server.Query("pdb/query/v4/fact-names")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &factNames)
	if err != nil {
		return nil, err
	}

	return factNames, err
}

// QueryFactPaths will query the PuppetDB instance fact-paths end-point.
func (server *Client) QueryFactPaths() (factPaths []string, err error) {
	body, err := server.Query("pdb/query/v4/fact-paths")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &factPaths)
	if err != nil {
		return nil, err

	}

	return factPaths, err
}

// QueryFactContents will query the PuppetDB instance fact-contents end-point.
func (server *Client) QueryFactContents(queryString string) (factContents []FactContent, err error) {
	url := fmt.Sprintf("pdb/query/v4/fact-contents?%v", queryString)

	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &factContents)
	if err != nil {
		return nil, err
	}

	return factContents, err
}

// QueryCatalogs will query the PuppetDB instance catalogs end-point for a given agent
func (server *Client) QueryCatalogs(certname string) (catalog *CatalogWireFormat, err error) {
	url := fmt.Sprintf("pdb/query/v4/catalogs/%v", certname)
	body, err := server.Query(url)
	if err != nil {
		return catalog, err
	}

	catalog = &CatalogWireFormat{}
	err = json.Unmarshal(body, catalog)
	if err != nil {
		return nil, err
	}

	return catalog, err
}

// QueryFacts will query the PuppetDB instance facts end-point.
func (server *Client) QueryFacts(queryString string) (facts []Fact, err error) {
	url := fmt.Sprintf("pdb/query/v4/facts?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, err
}

// QueryFactsByName will query the PuppetDB instance facts end-point for a given fact name.
func (server *Client) QueryFactsByName(name string, queryString string) (facts []Fact, err error) {
	url := fmt.Sprintf("pdb/query/v4/facts/%v?%v", name, queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, err
}

// QueryFactsByNameValue will query the PuppetDB instance facts end-point for a given fact name and value.
func (server *Client) QueryFactsByNameValue(name string, value string, queryString string) (facts []Fact, err error) {
	url := fmt.Sprintf("pdb/query/v4/facts/%v/%v?%v", name, value, queryString)

	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, err
}

// QueryResources will query the PuppetDB instance resources end-point.
func (server *Client) QueryResources(queryString string) (catalogResource []CatalogResource, err error) {
	url := fmt.Sprintf("pdb/query/v4/resources?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &catalogResource)
	if err != nil {
		return nil, err
	}

	return catalogResource, err
}

// QueryNodes will query the PuppetDB instance nodes end-point.
func (server *Client) QueryNodes(queryString string) (nodes []Node, err error) {
	url := fmt.Sprintf("pdb/query/v4/nodes?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, err
}

// QueryReports will query the PuppetDB instance reports end-point.
func (server *Client) QueryReports(queryString string) (reports []Report, err error) {
	url := fmt.Sprintf("pdb/query/v4/reports?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &reports)
	if err != nil {
		return nil, err
	}

	return reports, err
}

// QueryEvents will query the PuppetDB instance events end-point.
func (server *Client) QueryEvents(queryString string) (events []Event, err error) {
	url := fmt.Sprintf("pdb/query/v4/events?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}

	return events, err
}

// QueryEventCounts will query the PuppetDB instance event-counts end-point.
func (server *Client) QueryEventCounts(queryString string) (counts []EventCounts, err error) {
	url := fmt.Sprintf("pdb/query/v4/event-counts?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &counts)
	if err != nil {
		return nil, err
	}

	return counts, err
}

// QueryAggregateEventCounts will query the PuppetDB instance aggregate-event-counts end-point.
func (server *Client) QueryAggregateEventCounts(queryString string) (aggregateEventCounts *AggregateEventCounts, err error) {
	url := fmt.Sprintf("pdb/query/v4/aggregate-event-counts?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	aggregateEventCounts = &AggregateEventCounts{}
	err = json.Unmarshal(body, aggregateEventCounts)
	if err != nil {
		return nil, err
	}

	return aggregateEventCounts, err
}

// QueryInventory will query the PuppetDB instance inventory end-point.
func (server *Client) QueryInventory(queryString string) (inventory []Inventory, err error) {
	url := fmt.Sprintf("pdb/query/v4/inventory?%v", queryString)
	body, err := server.Query(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &inventory)
	if err != nil {
		return nil, err
	}

	return inventory, err
}
