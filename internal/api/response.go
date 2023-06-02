package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}

// NewResponse - creates new API response
func NewResponse(response *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var parsedBody interface{}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: response.StatusCode,
		Body:       parsedBody,
		Headers:    response.Header,
	}, nil
}
