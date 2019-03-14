package bitrise

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ErrorMsg struct {
	Message string `json:"error_msg"`
}

const (
	bitriseApiEndpoint = "https://app.bitrise.io/app/"
)

func checkResponse(response *http.Response, err error) (*http.Response, error) {
	switch response.StatusCode {
	case 200, 201:
		return response, nil
	default:
		var reader io.Reader = response.Body
		reader = io.TeeReader(r, os.Stderr)

		var errorMsg ErrorMsg
		err = json.NewDecoder(reader).Decode(&errorMsg)
		if err != nil {
			return nil, fmt.Errorf(err)
		} else {
			return nil, fmt.Errorf("Status code: %s", response.Status)
		}
	}
}
