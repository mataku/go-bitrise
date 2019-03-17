package bitrise

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCheckResponse_successful(t *testing.T) {
	successfulResponse := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
	}
	resp, err := checkResponse(&successfulResponse, nil)
	if err != nil {
		t.Fatalf("Failed! err is %s", err.Error())
	}

	if resp == nil {
		t.Fatalf("Failed! response is nil")
	}
}

func TestCheckResponse_failure(t *testing.T) {
	failedResponse := http.Response{
		Status:     "404 Not Found",
		StatusCode: 404,
		Body:       ioutil.NopCloser(strings.NewReader("{error_msg: fail}")),
	}
	_, err := checkResponse(&failedResponse, nil)
	if err == nil {
		t.Fatalf("Failed! err is nil..")
	}
}
