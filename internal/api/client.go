package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	olerror "github.com/onelogin/onelogin-go-sdk/internal/error"
)

type Client struct {
	HttpClient HTTPClient
	Auth       *authentication.Authenticator
	OLdomain   string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Authenticator interface {
	GetToken() (*string, error)
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		Auth:       authentication.NewAuthenticator(),
		OLdomain:   fmt.Sprintf("https://%s.onelogin.com", os.Getenv("ONELOGIN_SUBDOMAIN")),
	}
}
func (c *Client) Get(path string, queryParams map[string]string) ([]byte, error) {
	req, err := c.newRequest(http.MethodGet, path, queryParams, nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

func (c *Client) Post(path string, queryParams map[string]string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, path, queryParams, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}
func (c *Client) Delete(path string, queryParams map[string]string) ([]byte, error) {
	req, err := c.newRequest(http.MethodDelete, path, queryParams, nil)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req)
}
func (c *Client) Put(path string, queryParams map[string]string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPut, path, queryParams, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)

}

func (c *Client) newRequest(method, path string, queryParams map[string]string, bodyReader *bytes.Reader) (*http.Request, error) {

	req, err := http.NewRequest(method, c.OLdomain, bodyReader)
	if err != nil {
		return nil, err
	}
	tk, err := c.Auth.GetToken()
	if err != nil {
		return nil, olerror.NewAuthenticationError("Token Generation/Retrieval Error")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tk))
	req.Header.Set("Content-Type", "application/json")

	// Add query parameters
	query := req.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (c *Client) sendRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("%b", respBody)
		apiError := olerror.NewAPIError(message, resp.StatusCode)
		return nil, apiError
	}

	return respBody, nil
}
