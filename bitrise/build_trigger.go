package bitrise

import (
	"encoding/json"
)

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
	Branch        string       `json:"branch"`
	Tag           string       `json:"tag"`
	CommitHash    string       `json:"commit_hash"`
	CommitMessage string       `json:"commit_message"`
	WorkflowId    string       `json:"workflow_id"`
	Environment []Environments `json:"environments"`
}
