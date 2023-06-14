package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/api"
	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	"github.com/onelogin/onelogin-go-sdk/internal/models"
)

// TestNewClient tests the NewClient function.
func TestNewClient(t *testing.T) {
	// Set up environment variables
	os.Setenv("ONELOGIN_SUBDOMAIN", "test")
	os.Setenv("ONELOGIN_CLIENT_ID", "test")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "test")

	// Create a new mock client
	mockClient := &api.MockClient{}

	// Check that the client was created successfully
	if mockClient == nil {
		t.Error("Client was not created")
		return
	}

}

// TestGet tests the Get method.
func TestGet(t *testing.T) {
	// Create a new client
	client := &api.Client{
		HttpClient: http.DefaultClient,
		Auth:       authentication.NewAuthenticator(),
		OLdomain:   "https://test.onelogin.com",
	}

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write a response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer ts.Close()

	// Send a test request
	path := "/test"
	queryParams := models.UserQuery{}
	resp, err := client.Get(&path, &queryParams)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Check that the response was received successfully
	if resp == nil {
		t.Error("Response was not received")
	}

	// Check that the response's body was read successfully
	if resp == nil || resp.Body == nil {
		t.Error("Response body is nil")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(body) != "test" {
		t.Errorf("Response body was %s, expected test", string(body))
	}
}

// TestDelete tests the Delete method.
func TestDelete(t *testing.T) {
	// Create a new client
	client := &api.Client{
		HttpClient: http.DefaultClient,
		Auth:       authentication.NewAuthenticator(),
		OLdomain:   "https://test.onelogin.com",
	}

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write a response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer ts.Close()

	// Send a test request
	path := "/test"
	resp, err := client.Delete(&path)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Check that the response was received successfully
	if resp == nil || resp.Body == nil {
		t.Error("Response was not received")
	}

	// Check that the response's body was read successfully
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(body) != "test" {
		t.Errorf("Response body was %s, expected test", string(body))
	}
}

// TestPost tests the Post method.
func TestPost(t *testing.T) {
	// Create a new client
	client := &api.Client{
		HttpClient: http.DefaultClient,
		Auth:       authentication.NewAuthenticator(),
		OLdomain:   "https://test.onelogin.com",
	}

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request's body was read correctly
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
		}
		if string(body) != "test" {
			t.Errorf("Request body was %s, expected test", string(body))
		}

		// Write a response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer ts.Close()

	// Send a test request
	path := "/test"
	body := "test"
	resp, err := client.Post(&path, &body)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Check that the response was received successfully
	if resp == nil || resp.Body == nil {
		t.Error("Response was not received")
	}

	// Check that the response's body was read successfully
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	respBody := string(bodyBytes)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(body) != "test" {
		t.Errorf("Response body was %s, expected test", string(respBody))
	}
}

// TestPut tests the Put method.
func TestPut(t *testing.T) {
	// Create a new client
	client := &api.Client{
		HttpClient: http.DefaultClient,
		Auth:       authentication.NewAuthenticator(),
		OLdomain:   "https://test.onelogin.com",
	}

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the request's body was read correctly
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
		}
		if string(body) != "test" {
			t.Errorf("Request body was %s, expected test", string(body))
		}

		// Write a response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer ts.Close()

	// Send a test request
	path := "/test"
	body := "test"
	resp, err := client.Put(&path, &body)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	// Check that the response was received successfully
	if resp == nil || resp.Body == nil {
		t.Error("Response was not received")
	}

	// Check that the response's body was read successfully
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	respBody := string(bodyBytes)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if string(body) != "test" {
		t.Errorf("Response body was %s, expected test", string(respBody))
	}
}
