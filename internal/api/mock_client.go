package api

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
)

type MockClient struct {
	Auth     *authentication.MockAuthenticator // Authenticator for managing authentication
	OLdomain string                            // OneLogin domain
}

func (c *MockClient) Post(path *string, body interface{}) (*http.Response, error) {
	token, err := c.Auth.GetToken()
	if err != nil {
		return nil, err
	}

	// Use the token in the request
	bodyString, ok := body.(string)
	if !ok {
		return nil, errors.New("invalid request body")
	}
	req, err := http.NewRequest("POST", c.OLdomain+*path, bytes.NewBufferString(bodyString))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *MockClient) Put(path *string, body interface{}) (*http.Response, error) {
	token, err := c.Auth.GetToken()
	if err != nil {
		return nil, err
	}

	// Use the token in the request
	bodyString, ok := body.(string)
	if !ok {
		return nil, errors.New("invalid request body")
	}
	req, err := http.NewRequest("PUT", c.OLdomain+*path, bytes.NewBufferString(bodyString))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
