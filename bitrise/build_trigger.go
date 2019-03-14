package bitrise

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	apiEndPoint = "https://app.bitrise.io/app/"
)

type RequestParams struct {
	Info        HookInfo    `json:"hook_info"`
	BuildParams BuildParams `json:"build_params"`
}

type HookInfo struct {
	Type              string `json:"type"`
	BuildTriggerToken string `json:"build_trigger_token"`
}

type Environments struct {
	MappedTo string `json:"mapped_to"`
	Value    string `json:"value"`
	IsExpand bool   `json:"is_expand"`
}

type BuildParams struct {
	Branch        string         `json:"branch"`
	Tag           string         `json:"tag"`
	CommitHash    string         `json:"commit_hash"`
	CommitMessage string         `json:"commit_message"`
	WorkflowId    string         `json:"workflow_id"`
	Environment   []Environments `json:"environments"`
}

func (c *Client) TriggerBuild(hi *HookInfo, bp BuildParams) (*http.Request, error) {
	token := hi.BuildTriggerToken
	if token == "" {
		return nil, fmt.Errorf("hook_info: Build trigger token required.")
	}

	requestParams := RequestParams{
		HookInfo:    hi,
		BuildParams: bp,
	}

	req, err := http.NewRequest(
		"POST",
		apiEndPoint+client.AppSlug+"/build/start.json",
		bytes.NewBuffer([]byte(string(requestParams))),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := checkResponse(c.HttpClient.Do(req))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
