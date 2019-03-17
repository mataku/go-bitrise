package bitrise

import (
	"testing"
)

func TestTriggerBuild_noBuildTriggerTokenPassed(t *testing.T) {
	client := NewClient("appSlug")
	hookInfo := HookInfo{}
	buildParams := BuildParams{}
	res, err := client.TriggerBuild(hookInfo, buildParams)
	if err == nil {
		t.Fatalf("err is nil!")
	}

	if res != nil {
		t.Fatalf("response is not nil! Status: %s", res.Status)
	}
}
