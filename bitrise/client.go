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

type Client struct {
	AppSlug    string
	HttpClient http.Client
}

func NewClient(appSlug string) *Client {
	return &Client{
		AppSlug: appSlug,
	}
}

func checkResponse(response *http.Response, err error) (*http.Response, error) {
	switch response.StatusCode {
	case 200, 201:
		return response, nil
	default:
		var reader io.Reader = response.Body
		reader = io.TeeReader(reader, os.Stderr)

		var errorMsg ErrorMsg
		err = json.NewDecoder(reader).Decode(&errorMsg)
		// TODO: Same message format
		if err != nil {
			return nil, fmt.Errorf("Status code: %s", response.Status)
		} else {
			return nil, fmt.Errorf(errorMsg.Message)
		}
	}
}
