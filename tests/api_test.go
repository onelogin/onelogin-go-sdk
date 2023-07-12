package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/api"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/authentication"
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

	auth := authentication.NewAuthenticator()
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
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
		}
		return response, nil
	}

	resp, err := client.Get(new(string), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != `{"key":"value"}` {
		t.Fatalf("Expected `{\"key\":\"value\"}`, got %s", string(body))
	}
}

func TestClientPost(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 201,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"result":"created"}`)),
		}
		return response, nil
	}

	resp, err := client.Post(new(string), map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
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
			Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
		}
		return response, nil
	}

	resp, err := client.Delete(new(string))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != `` {
		t.Fatalf("Expected ``, got %s", string(body))
	}
}

func TestClientDeleteWithBody(t *testing.T) {
	client := createMockClient()

	client.HttpClient.(*MockHttpClient).DoFunc = func(*http.Request) (*http.Response, error) {
		response := &http.Response{
			StatusCode: 204,
			Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
		}
		return response, nil
	}

	resp, err := client.DeleteWithBody(new(string), map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != `` {
		t.Fatalf("Expected ``, got %s", string(body))
	}
}
