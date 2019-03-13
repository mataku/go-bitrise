package bitrise

import (
	"fmt"
	"net/http"
)

const (
	bitriseApiEndpoint = "https://app.bitrise.io/app/"
)

func checkResponse(response *http.Response, err error) (*http.Response, error) {
	switch response.StatusCode {
	case 200, 201:
		return response, nil
	}
}
