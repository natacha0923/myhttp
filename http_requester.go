package main

import (
	"io/ioutil"
	"net/http"
)

type HTTPRequester struct{}

func NewHTTPRequester() *HTTPRequester {
	return &HTTPRequester{}
}

// SendRequest sends http request to the specified url.
// Returns response body or error if it is exists.
func (x *HTTPRequester) SendRequest(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
