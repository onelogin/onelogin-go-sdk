package olhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

type TestResource struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type TestResourceQuery struct {
	Name  string `json:"name"`
	Limit string `json:"limit"`
}

var available = []TestResource{
	TestResource{ID: "0", Name: "name0"},
	TestResource{ID: "1", Name: "name1"},
	TestResource{ID: "2", Name: "name2"},
	TestResource{ID: "3", Name: "name3"},
}

func authPassThrough() (*http.Response, error) {
	json := `{"access_token":"token"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{StatusCode: 200, Body: r}, nil
}

// just query for the first one
func getOneByID() TestResource {
	return available[0]
}

// get n items starting at offset o
func queryNResources(n int, o int) []TestResource {
	var out []TestResource
	for i := o; i < n+o; i++ {
		out = append(out, available[i])
	}
	return out
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		resourceRequest OLHTTPRequest
		expectedOut     TestResource
	}{
		"Create with payload": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				URL:        "test.com/test_resources",
				Headers:    map[string]string{"Content-Type": "application/json"},
				Payload:    TestResource{Name: "test"},
			},
			expectedOut: TestResource{Name: "test", ID: "1"},
		},
		"Create without payload": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedOut: TestResource{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							return authPassThrough()
						}
						if test.resourceRequest.Payload != nil {
							tpl := test.resourceRequest.Payload.(TestResource)
							tpl.ID = "1"
							j, _ := json.Marshal(tpl)
							r := ioutil.NopCloser(bytes.NewReader([]byte(j)))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}
						r := ioutil.NopCloser(bytes.NewReader([]byte{}))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			actual, err := svc.Create(test.resourceRequest)
			var actualOut TestResource
			json.Unmarshal(actual, &actualOut)
			assert.Equal(t, test.expectedOut, actualOut)
			assert.Nil(t, err)
		})
	}
}

func TestRead(t *testing.T) {
	tests := map[string]struct {
		resourceRequest  OLHTTPRequest
		expectedReadOut  TestResource
		expectedQueryOut []TestResource
	}{
		"Get n resources according to limit": {
			resourceRequest: OLHTTPRequest{
				Payload:    TestResourceQuery{Limit: "2"},
				AuthMethod: "bearer",
				URL:        "test.com/test_resources",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedQueryOut: []TestResource{
				TestResource{Name: "name0", ID: "0"},
				TestResource{Name: "name1", ID: "1"},
			},
		},
		"Read one by giving resource id in URL": {
			resourceRequest: OLHTTPRequest{
				Payload:    TestResourceQuery{},
				AuthMethod: "bearer",
				URL:        "test.com/test_resources/0",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedReadOut: TestResource{Name: "name0", ID: "0"},
		},
		"Read all if no limit, id or query given": {
			resourceRequest: OLHTTPRequest{
				Payload:    TestResourceQuery{},
				AuthMethod: "bearer",
				URL:        "test.com/test_resources",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedQueryOut: []TestResource{
				TestResource{Name: "name0", ID: "0"},
				TestResource{Name: "name1", ID: "1"},
				TestResource{Name: "name2", ID: "2"},
				TestResource{Name: "name3", ID: "3"},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							return authPassThrough()
						}

						if req.URL.Path == "test.com/test_resources/0" {
							resource := getOneByID()
							out, _ := json.Marshal(resource)
							r := ioutil.NopCloser(bytes.NewReader([]byte(out)))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}

						tpl := test.resourceRequest.Payload.(TestResourceQuery)
						if tpl.Limit != "" {
							n, _ := strconv.Atoi(tpl.Limit)
							resources := queryNResources(n, 0)
							out, _ := json.Marshal(resources)
							r := ioutil.NopCloser(bytes.NewReader([]byte(out)))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}

						resources := queryNResources(2, 0)
						if req.URL.Query().Get("Cursor") != "" {
							moreResources := queryNResources(2, 2)
							resources = append(resources, moreResources...)
							out, _ := json.Marshal(resources)
							r := ioutil.NopCloser(bytes.NewReader([]byte(out)))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}
						out, _ := json.Marshal(resources)
						r := ioutil.NopCloser(bytes.NewReader([]byte(out)))
						return &http.Response{StatusCode: 200, Body: r, Header: http.Header{"After-Cursor": []string{"asdf"}}}, nil
					},
				},
			})

			if test.expectedQueryOut != nil {
				actual, err := svc.Read(test.resourceRequest)

				var actualOut []TestResource
				err = json.Unmarshal(actual, &actualOut)
				if err != nil {
					fmt.Println("ASDF", name, err)
				}

				assert.Equal(t, test.expectedQueryOut, actualOut)
				assert.Nil(t, err)
			} else {
				actual, err := svc.Read(test.resourceRequest)
				var actualOut TestResource
				json.Unmarshal(actual, &actualOut)
				assert.Equal(t, test.expectedReadOut, actualOut)
				assert.Nil(t, err)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		resourceRequest  OLHTTPRequest
		originalResource TestResource
		expectedOut      TestResource
	}{
		"Update with payload": {
			originalResource: TestResource{Name: "unchanged", ID: "1"},
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				URL:        "test.com/test_resources/1",
				Headers:    map[string]string{"Content-Type": "application/json"},
				Payload:    TestResource{Name: "changed", ID: "1"},
			},
			expectedOut: TestResource{Name: "changed", ID: "1"},
		},
		"Update without payload": {
			originalResource: TestResource{Name: "unchanged", ID: "1"},
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				URL:        "test.com/test_resources/1",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedOut: TestResource{Name: "unchanged", ID: "1"},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							return authPassThrough()
						}
						if test.resourceRequest.Payload != nil {
							tpl := test.resourceRequest.Payload.(TestResource)
							tpl.Name = "changed"
							j, _ := json.Marshal(tpl)
							r := ioutil.NopCloser(bytes.NewReader([]byte(j)))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}
						tpl := test.originalResource
						j, _ := json.Marshal(tpl)
						r := ioutil.NopCloser(bytes.NewReader([]byte(j)))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			actual, err := svc.Update(test.resourceRequest)
			var actualOut TestResource
			json.Unmarshal(actual, &actualOut)
			assert.Equal(t, test.expectedOut, actualOut)
			assert.Nil(t, err)
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		resourceRequest OLHTTPRequest
		expectedOut     TestResource
	}{
		"Destroy": {
			resourceRequest: OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			expectedOut:     TestResource{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							return authPassThrough()
						}
						r := ioutil.NopCloser(bytes.NewReader([]byte(``)))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			actual, err := svc.Destroy(test.resourceRequest)
			var actualOut TestResource
			json.Unmarshal(actual, &actualOut)
			assert.Equal(t, test.expectedOut, actualOut)
			assert.Nil(t, err)
		})
	}
}

func TestExecuteHTTP(t *testing.T) {
	tests := map[string]struct {
		resourceRequest OLHTTPRequest
		expectedError   error
		expectedOut     TestResource
		staleToken      bool
		badCredential   bool
	}{
		"attempts the request": {
			resourceRequest: OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			staleToken:      false,
			expectedOut:     TestResource{Name: "name", ID: "1"},
		},
		"attempts bearer token refresh and retries the http request if auth fails and auth method is bearer token": {
			resourceRequest: OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			staleToken:      true,
			expectedOut:     TestResource{Name: "name", ID: "1"},
		},
		"unauthorized api request returns an error": {
			resourceRequest: OLHTTPRequest{AuthMethod: "basic", URL: "test.com/test_resources/1"},
			badCredential:   true,
			expectedOut:     TestResource{},
			expectedError:   customerrors.OneloginErrorWrapper("ol request", errors.New("unauthorized")),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							return authPassThrough()
						}
						if test.staleToken {
							test.staleToken = false
							r := ioutil.NopCloser(bytes.NewReader([]byte(``)))
							return &http.Response{StatusCode: 401, Body: r}, errors.New("unauthorized")
						}
						if test.badCredential {
							r := ioutil.NopCloser(bytes.NewReader([]byte(``)))
							return &http.Response{StatusCode: 401, Body: r}, errors.New("unauthorized")
						}
						j, _ := json.Marshal(test.expectedOut)
						r := ioutil.NopCloser(bytes.NewReader([]byte(j)))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			req, _ := http.NewRequest(http.MethodGet, test.resourceRequest.URL, nil)
			_, data, err := svc.executeHTTP(req, test.resourceRequest)
			if test.expectedError == nil {
				assert.Nil(t, err)
				var tr = TestResource{}
				json.Unmarshal(data, &tr)
				assert.Equal(t, test.expectedOut, tr)
			} else {
				assert.Equal(t, test.expectedError, err)
			}

		})
	}
}

func TestAttachHeaders(t *testing.T) {
	tests := map[string]struct {
		req             *http.Request
		olReq           OLHTTPRequest
		service         OLHTTPService
		expectedHeaders http.Header
	}{
		"it attaches basic auth headers": {
			req: &http.Request{
				Header: http.Header{},
			},
			olReq: OLHTTPRequest{
				AuthMethod: "basic",
				Headers: map[string]string{
					"Content-type": "application/json",
				},
			},
			service: OLHTTPService{
				Config: services.HTTPServiceConfig{
					ClientID:     "test",
					ClientSecret: "sec",
				},
				ClientCredential: ClientCredential{
					AccessToken: oltypes.String("token"),
				},
			},
			expectedHeaders: http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Basic dGVzdDpzZWM="},
			},
		},
		"it attaches bearer auth headers": {
			req: &http.Request{
				Header: http.Header{},
			},
			olReq: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers: map[string]string{
					"Content-type": "application/json",
				},
			},
			service: OLHTTPService{
				Config: services.HTTPServiceConfig{
					ClientID:     "test",
					ClientSecret: "sec",
				},
				ClientCredential: ClientCredential{
					AccessToken: oltypes.String("token"),
				},
			},
			expectedHeaders: http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer token"},
			},
		},
		"it attaches headers without auth": {
			req: &http.Request{
				Header: http.Header{},
			},
			olReq: OLHTTPRequest{
				Headers: map[string]string{
					"Content-type": "application/json",
				},
			},
			service: OLHTTPService{
				Config: services.HTTPServiceConfig{
					ClientID:     "test",
					ClientSecret: "sec",
				},
				ClientCredential: ClientCredential{
					AccessToken: oltypes.String("token"),
				},
			},
			expectedHeaders: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.service.attachHeaders(test.req, test.olReq)
			assert.Equal(t, test.expectedHeaders, test.req.Header)
		})
	}
}

func TestAuthorize(t *testing.T) {
	tests := map[string]struct {
		service, expectedServiceState OLHTTPService
	}{
		"it memoizes the access token on the service": {
			service: OLHTTPService{
				Config: services.HTTPServiceConfig{
					Client: &MockClient{
						DoFunc: func(*http.Request) (*http.Response, error) {
							return authPassThrough()
						},
					},
				},
			},
			expectedServiceState: OLHTTPService{
				ClientCredential: ClientCredential{
					AccessToken: oltypes.String("token"),
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.service.authorize()
			assert.Nil(t, err)
			assert.Equal(t, *test.expectedServiceState.ClientCredential.AccessToken, *test.service.ClientCredential.AccessToken)
		})
	}
}

func TestSetBearerToken(t *testing.T) {
	tests := map[string]struct {
		service, expectedServiceState OLHTTPService
	}{
		"it memoizes the access token from the api on the service": {
			service: OLHTTPService{
				Config: services.HTTPServiceConfig{
					Client: &MockClient{
						DoFunc: func(*http.Request) (*http.Response, error) {
							return authPassThrough()
						},
					},
				},
			},
			expectedServiceState: OLHTTPService{
				ClientCredential: ClientCredential{
					AccessToken: oltypes.String("token"),
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setBearerToken(&test.service)
			assert.Equal(t, *test.expectedServiceState.ClientCredential.AccessToken, *test.service.ClientCredential.AccessToken)
		})
	}
}
