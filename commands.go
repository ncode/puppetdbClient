package puppetdb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
Generic command submission support, for submitting commands to a PuppetDB instance.

This is ordinarily not used, instead its recommended to use the various direct
functions instead.

More detail here: http://docs.puppetlabs.com/puppetdb/latest/api/commands.html
*/
func (server *Server) SubmitCommand(command string, version int, payload interface{}) (*CommandResponse, error) {
	baseUrl := server.BaseUrl
	commandsUrl := strings.Join([]string{baseUrl, "v3/commands"}, "")

	commandObject := CommandObject{command, version, payload}
	commandJson, err := json.Marshal(commandObject)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("payload", string(commandJson[:]))

	req, err := http.NewRequest("POST", commandsUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyRC, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var commandResponse CommandResponse
	err = json.Unmarshal(bodyRC, &commandResponse)
	if err != nil {
		return nil, err
	}

	return &commandResponse, nil
}

// DeactivateNode Deactivate a node in PuppetDB.
func (server *Server) DeactivateNode(certname string) (commandResponse *CommandResponse, err error) {
	certnameJson, err := json.Marshal(certname)
	if err != nil {
		return commandResponse, err
	}

	return server.SubmitCommand("deactivate node", 1, string(certnameJson[:]))
}
