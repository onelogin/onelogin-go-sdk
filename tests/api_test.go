package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/api"
	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
)

type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestNewRequest(t *testing.T) {
	httpmock := MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			resp := http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			}
			return &resp, nil
		},
	}

	authenticator := authentication.NewAuthenticator()

	client := api.Client{
		HttpClient: &httpmock,
		Auth:       authenticator,
		OLdomain:   "http://localhost",
	}
	test := "/test"
	resp, err := client.Get(&test, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGet(t *testing.T) {
	// Create a mock HTTP client
	httpmock := MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			// Modify the response based on the test case
			if req.URL.Path == "/test" && req.Method == "GET" {
				// Successful case
				resp := http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
				}
				return &resp, nil
			} else if req.URL.Path == "/test" && req.Method == "POST" {
				// Unsuccessful case - incorrect method
				resp := http.Response{
					StatusCode: http.StatusMethodNotAllowed,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`Method not allowed`)),
				}
				return &resp, nil
			} else {
				// Unsuccessful case - unexpected request
				resp := http.Response{
					StatusCode: http.StatusNotFound,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`Not found`)),
				}
				return &resp, nil
			}
		},
	}

	// Create an API client with the mock HTTP client
	authenticator := authentication.NewAuthenticator()
	client := api.Client{
		HttpClient: &httpmock,
		Auth:       authenticator,
		OLdomain:   "http://localhost",
	}

	// Test successful GET request
	test := "/test"
	resp, err := client.Get(&test, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// Test unsuccessful POST request
	resp, err = client.Post(&test, nil)
	if err == nil {
		t.Error("Expected error, but got no error")
	}
	if resp != nil && resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, but got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}
	uhoh := "/unexpected"
	// Test unsuccessful GET request with unexpected path
	resp, err = client.Get(&uhoh, nil)
	if err == nil {
		t.Error("Expected error, but got no error")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, resp.StatusCode)
	}
}
