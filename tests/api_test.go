package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/api"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/authentication"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type MockAuthenticator struct {
	GetTokenFunc         func() (string, error)
	NewAuthenticatorFunc func() *authentication.Authenticator
}

func (m *MockAuthenticator) GetToken() (string, error) {
	return m.GetTokenFunc()
}

func (m *MockAuthenticator) NewAuthenticator() *authentication.Authenticator {
	return &authentication.Authenticator{}
}

func createMockClient() *api.Client {
	mockClient := &MockHttpClient{}
	mockAuth := &MockAuthenticator{}

	mockAuth.GetTokenFunc = func() (string, error) {
		return "mockToken", nil
	}

	auth := authentication.NewAuthenticator("test")
	client := &api.Client{
		HttpClient: mockClient,
		Auth:       auth,
		OLdomain:   "https://api.onelogin.com",
	}

	return client
}

func TestClientGet(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
		}
		return response, nil
	}

	resp, err := client.Get(new(string), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `{"key":"value"}` {
		t.Fatalf("Expected `{\"key\":\"value\"}`, got %s", string(body))
	}
}

func TestClientPost(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(bytes.NewBufferString(`{"result":"created"}`)),
		}
		return response, nil
	}

	resp, err := client.Post(new(string), map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `{"result":"created"}` {
		t.Fatalf("Expected `{\"result\":\"created\"}`, got %s", string(body))
	}
}

// ... Additional tests for Delete, DeleteWithBody, Put

func TestClientDelete(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 204,
			Body:       io.NopCloser(bytes.NewBufferString(``)),
		}
		return response, nil
	}

	resp, err := client.Delete(new(string))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `` {
		t.Fatalf("Expected ``, got %s", string(body))
	}
}

func TestClientDeleteWithBody(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 204,
			Body:       io.NopCloser(bytes.NewBufferString(``)),
		}
		return response, nil
	}

	resp, err := client.DeleteWithBody(new(string), map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `` {
		t.Fatalf("Expected ``, got %s", string(body))
	}
}

func TestRoleEmptyArraysSerialization(t *testing.T) {
	// Create role with empty arrays
	role := &models.Role{
		Name:   new(string),
		Users:  []int32{}, // Empty array
		Admins: []int32{}, // Empty array
		Apps:   []int32{}, // Empty array
	}
	*role.Name = "Test Role"

	// Serialize to JSON
	jsonData, err := json.Marshal(role)
	if err != nil {
		t.Fatalf("Failed to marshal role: %v", err)
	}

	// Verify that empty arrays are included in JSON
	jsonStr := string(jsonData)
	t.Logf("Serialized JSON: %s", jsonStr)

	// Check for empty arrays in the JSON
	if !bytes.Contains(jsonData, []byte(`"users":[]`)) {
		t.Error("JSON should contain empty users array")
	}
	if !bytes.Contains(jsonData, []byte(`"admins":[]`)) {
		t.Error("JSON should contain empty admins array")
	}
	if !bytes.Contains(jsonData, []byte(`"apps":[]`)) {
		t.Error("JSON should contain empty apps array")
	}

	// Also check with nil slices
	roleWithNil := &models.Role{
		Name: role.Name,
		// All arrays are nil by default
	}

	jsonData, err = json.Marshal(roleWithNil)
	if err != nil {
		t.Fatalf("Failed to marshal role with nil arrays: %v", err)
	}

	jsonStr = string(jsonData)
	t.Logf("Serialized JSON with nil arrays: %s", jsonStr)

	// Check for empty arrays in the JSON
	if !bytes.Contains(jsonData, []byte(`"users":[]`)) {
		t.Error("JSON should contain empty users array for nil slice")
	}
	if !bytes.Contains(jsonData, []byte(`"admins":[]`)) {
		t.Error("JSON should contain empty admins array for nil slice")
	}
	if !bytes.Contains(jsonData, []byte(`"apps":[]`)) {
		t.Error("JSON should contain empty apps array for nil slice")
	}
}
