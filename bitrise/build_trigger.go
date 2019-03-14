package bitrise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestParams struct {
	HookInfo    HookInfo    `json:"hook_info"`
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

func TriggerBuild(appSlug string, hi *HookInfo, bp BuildParams) (*http.Request, error) {
	token := hi.BuildTriggerToken
	if token == "" {
		return nil, fmt.Errorf("hook_info: Build trigger token required.")
	}

	if appSlug == "" {
		return nil, fmt.Errorf("app_slug: App slug required.")
	}

	requestParams := RequestParams{
		HookInfo:    hi,
		BuildParams: bp,
	}

	req, err := http.NewRequest(
		"POST",
		"https://app.bitrise.io/app/"+appSlug+"/build/start.json",
		bytes.NewBuffer([]byte(string(requestParams))),
	)

	if err != nil {
		fmt.Errorf(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client()
	resp, err := checkResponse(client.Do(req))
	if err != nil {
		fmt.Errorf(err)
	}
}
