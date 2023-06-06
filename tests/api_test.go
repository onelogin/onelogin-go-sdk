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

	_, err := client.Get("/test", nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// Other tests go here.
