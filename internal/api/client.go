package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	olerror "github.com/onelogin/onelogin-go-sdk/internal/error"
)

// Client represents the API client.
type Client struct {
	HttpClient HTTPClient                    // HTTPClient interface for making HTTP requests
	Auth       *authentication.Authenticator // Authenticator for managing authentication
	OLdomain   string                        // OneLogin domain
}

// HTTPClient is an interface that defines the Do method for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Authenticator is an interface that defines the GetToken method for retrieving authentication tokens.
type Authenticator interface {
	GetToken() (string, error)
}

// NewClient creates a new instance of the API client.
func NewClient() *Client {
	authenticator := authentication.NewAuthenticator()
	token, err := authenticator.GetToken()

	if err != nil || token == "" {
		_, err := authenticator.GenerateToken()
		if err != nil {
			// Handle error
			fmt.Printf("Failed to generate token: %s", err.Error())
			os.Exit(1)
		}
	}

	return &Client{
		HttpClient: http.DefaultClient,
		Auth:       authenticator,
		OLdomain:   fmt.Sprintf("https://%s.onelogin.com", os.Getenv("ONELOGIN_SUBDOMAIN")),
	}
}

// newRequest creates a new HTTP request with the specified method, path, query parameters, and request body.
func (c *Client) newRequest(method, path string, queryParams *map[string]string, body io.Reader) (*http.Request, error) {
	// Parse the OneLogin domain and path
	u, err := url.Parse(c.OLdomain + path)
	if err != nil {
		return nil, err
	}

	// Add query parameters to the URL
	query := u.Query()
	if queryParams != nil {
		for key, value := range *queryParams {
			query.Add(key, value)
		}
	}
	u.RawQuery = query.Encode()

	// Create a new HTTP request
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Get authentication token
	tk, err := c.Auth.GetToken()
	if err != nil {
		return nil, olerror.NewAuthenticationError("Access Token Retrieval Error")
	}

	// Set request headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tk))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// Get sends a GET request to the specified path with the given query parameters.
func (c *Client) Get(path string, queryParams *map[string]string) ([]byte, error) {
	req, err := c.newRequest(http.MethodGet, path, queryParams, http.NoBody)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

// Delete sends a DELETE request to the specified path with the given query parameters.
func (c *Client) Delete(path string, queryParams *map[string]string) ([]byte, error) {
	req, err := c.newRequest(http.MethodDelete, path, queryParams, http.NoBody)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

// Post sends a POST request to the specified path with the given query parameters and request body.
func (c *Client) Post(path string, queryParams *map[string]string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader

	if body != nil {
		// Convert request body to JSON
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := c.newRequest(http.MethodPost, path, queryParams, bodyReader)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

// Put sends a PUT request to the specified path with the given query parameters and request body.
func (c *Client) Put(path string, queryParams *map[string]string, body interface{}) ([]byte, error) {
	// Convert request body to JSON
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

// sendRequest sends the specified HTTP request and returns the response body.
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

	// Check for API errors
	if resp.StatusCode == http.StatusUnauthorized {
		// Regenerate the token and reattempt the request
		_, err := c.Auth.GenerateToken()
		if err != nil {
			return nil, olerror.NewAuthenticationError("Failed to refresh access token")
		}

		// Retry the request
		resp, err = c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		respBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("%b", respBody)
		apiError := olerror.NewAPIError(message, resp.StatusCode)
		return nil, apiError
	}

	return respBody, nil
}
