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

func TestRoleMembershipSerialization(t *testing.T) {
	name := "Test Role"

	tests := map[string]struct {
		role     models.Role
		expected string
	}{
		// The bug behind issue #114: a caller renaming a role sent empty
		// membership arrays it never set, and the API removed every app, user
		// and admin on that role.
		"omits memberships the caller never set": {
			role:     models.Role{Name: &name},
			expected: `{"name":"Test Role"}`,
		},
		"sends an explicitly empty array, which removes all memberships": {
			role: models.Role{
				Name:   &name,
				Users:  []int32{},
				Admins: []int32{},
				Apps:   []int32{},
			},
			expected: `{"admins":[],"apps":[],"name":"Test Role","users":[]}`,
		},
		"sends populated arrays, which replace the memberships": {
			role: models.Role{
				Name:  &name,
				Users: []int32{1, 2},
				Apps:  []int32{3},
			},
			expected: `{"apps":[3],"name":"Test Role","users":[1,2]}`,
		},
		"carries one membership change without disturbing the others": {
			role:     models.Role{Apps: []int32{7}},
			expected: `{"apps":[7]}`,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			// Marshalled as both a value and a pointer: the receiver is a value
			// so that json.Marshal cannot treat the two differently.
			for _, subject := range []interface{}{test.role, &test.role} {
				jsonData, err := json.Marshal(subject)
				if err != nil {
					t.Fatalf("Failed to marshal role: %v", err)
				}
				if string(jsonData) != test.expected {
					t.Errorf("marshalling %T\n got: %s\nwant: %s", subject, jsonData, test.expected)
				}
			}
		})
	}
}
