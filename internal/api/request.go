package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Url     string
	Headers map[string]string
	Body    interface{}
}

// NewRequest - creates new API request
func NewRequest(url string, headers map[string]string, body interface{}) *Request {
	return &Request{
		Url:     url,
		Headers: headers,
		Body:    body,
	}
}

// Send - send the API request
func (r *Request) Send() (*http.Response, error) {
	bodyBytes, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.Url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	return client.Do(req)
}
