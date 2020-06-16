package olhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

type TestResource struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
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
func authFailThrough() (*http.Response, error) {
	json := `{"type":"Unauthorized"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{StatusCode: 401, Body: r}, nil
}

// just query for the first one
func getOneByID(id int) (*http.Response, error) {
	resource := available[id]
	out, _ := json.Marshal(resource)
	r := ioutil.NopCloser(bytes.NewReader([]byte(out)))
	return &http.Response{StatusCode: 200, Body: r}, nil
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
		expectedError   error
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
		"HTTP Service Errors": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedOut:   TestResource{},
			expectedError: customerrors.OneloginErrorWrapper(resourceRequestuestContext, errors.New("service fail")),
		},
		"HTTP Service returns Non-Json response": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedOut:   TestResource{},
			expectedError: errors.New("Unable to unpack response"),
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
						if test.expectedError != nil {
							if test.expectedError.Error() == "error: context: [ol http service], error_message: [service fail]" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("service fail")))
								return &http.Response{StatusCode: 400, Body: r}, nil
							}
							if fmt.Sprintf("%s", test.expectedError) == "Unable to unpack response" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("[")))
								return &http.Response{StatusCode: 200, Body: r}, nil
							}
						}
						if test.resourceRequest.Payload != nil {
							tpl := test.resourceRequest.Payload.(TestResource)
							tpl.ID = "1"
							j, _ := json.Marshal(tpl)
							r := ioutil.NopCloser(bytes.NewReader(j))
							return &http.Response{StatusCode: 200, Body: r}, nil
						}
						r := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			actual := TestResource{}
			err := svc.Create(test.resourceRequest, &actual)
			if test.expectedError == nil {
				assert.Equal(t, test.expectedOut, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := map[string]struct {
		resourceRequest  OLHTTPRequest
		expectedReadOut  TestResource
		expectedQueryOut []TestResource
		expectedError    error
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
		"HTTP Service Errors": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: customerrors.OneloginErrorWrapper(resourceRequestuestContext, errors.New("service fail")),
		},
		"HTTP Service returns Non-Json response": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: errors.New("Unable to unpack response"),
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
						if test.expectedError != nil {
							if test.expectedError.Error() == "error: context: [ol http service], error_message: [service fail]" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("service fail")))
								return &http.Response{StatusCode: 400, Body: r}, nil
							}
							if fmt.Sprintf("%s", test.expectedError) == "Unable to unpack response" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("[")))
								return &http.Response{StatusCode: 200, Body: r}, nil
							}
						}
						parts := strings.Split(req.URL.Path, "/")
						idstr := parts[len(parts)-1]
						id, err := strconv.Atoi(idstr)
						if err == nil {
							return getOneByID(id)
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
			actual := []TestResource{}
			err := svc.Read(test.resourceRequest, &actual)
			if test.expectedQueryOut != nil {
				assert.Equal(t, test.expectedQueryOut, actual)
				assert.Nil(t, err)
			} else if test.expectedError == nil {
				actual := TestResource{}
				err := svc.Read(test.resourceRequest, &actual)
				assert.Equal(t, test.expectedReadOut, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedError, err)
			}

		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		resourceRequest  OLHTTPRequest
		originalResource TestResource
		expectedOut      TestResource
		expectedError    error
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
		"HTTP Service Errors": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: customerrors.OneloginErrorWrapper(resourceRequestuestContext, errors.New("service fail")),
		},
		"HTTP Service returns Non-Json response": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: errors.New("Unable to unpack response"),
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
						if test.expectedError != nil {
							if test.expectedError.Error() == "error: context: [ol http service], error_message: [service fail]" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("service fail")))
								return &http.Response{StatusCode: 400, Body: r}, nil
							}
							if fmt.Sprintf("%s", test.expectedError) == "Unable to unpack response" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("[")))
								return &http.Response{StatusCode: 200, Body: r}, nil
							}
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
			actual := TestResource{}
			err := svc.Update(test.resourceRequest, &actual)
			if test.expectedError == nil {
				assert.Equal(t, test.expectedOut, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		resourceRequest OLHTTPRequest
		expectedOut     TestResource
		expectedError   error
	}{
		"Destroy": {
			resourceRequest: OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			expectedOut:     TestResource{},
		},
		"HTTP Service Errors": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: customerrors.OneloginErrorWrapper(resourceRequestuestContext, errors.New("service fail")),
		},
		"HTTP Service returns Non-Json response": {
			resourceRequest: OLHTTPRequest{
				AuthMethod: "bearer",
				Headers:    map[string]string{"Content-Type": "application/json"},
			},
			expectedError: errors.New("Unable to unpack response"),
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
						if test.expectedError != nil {
							if test.expectedError.Error() == "error: context: [ol http service], error_message: [service fail]" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("service fail")))
								return &http.Response{StatusCode: 400, Body: r}, nil
							}
							if fmt.Sprintf("%s", test.expectedError) == "Unable to unpack response" {
								r := ioutil.NopCloser(bytes.NewReader([]byte("[")))
								return &http.Response{StatusCode: 200, Body: r}, nil
							}
						}
						r := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
						return &http.Response{StatusCode: 200, Body: r}, nil
					},
				},
			})
			actual := TestResource{}
			err := svc.Destroy(test.resourceRequest, &actual)
			if test.expectedError == nil {
				assert.Equal(t, test.expectedOut, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestExecuteHTTP(t *testing.T) {
	tests := map[string]struct {
		resourceRequest  OLHTTPRequest
		expectedError    error
		expectedOut      TestResource
		badRequest       bool
		serviceDown      bool
		staleToken       bool
		refreshTokenFail bool
		badCredential    bool
	}{
		"attempts the request": {
			resourceRequest: OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			staleToken:      false,
			expectedOut:     TestResource{Name: "name", ID: "1"},
		},
		"attempts bearer token refresh and retries the http request if auth fails and auth method is bearer token": {
			resourceRequest:  OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			staleToken:       true,
			refreshTokenFail: false,
			expectedOut:      TestResource{Name: "name", ID: "1"},
		},
		"failed bearer token refresh gives unauthorized error": {
			resourceRequest:  OLHTTPRequest{AuthMethod: "bearer", URL: "test.com/test_resources/1"},
			staleToken:       true,
			refreshTokenFail: true,
			expectedError:    customerrors.OneloginErrorWrapper("ol http service", errors.New("unauthorized")),
		},
		"unauthorized api request returns an error": {
			resourceRequest: OLHTTPRequest{AuthMethod: "basic", URL: "test.com/test_resources/1"},
			badCredential:   true,
			expectedError:   customerrors.OneloginErrorWrapper("ol http service", errors.New("unauthorized")),
		},
		"bad api request returns an error": {
			resourceRequest: OLHTTPRequest{AuthMethod: "basic", URL: "test.com/test_resources/1"},
			badRequest:      true,
			expectedError:   customerrors.OneloginErrorWrapper("ol http service", errors.New("bad request")),
		},
		"service down returns an error": {
			resourceRequest: OLHTTPRequest{AuthMethod: "basic", URL: "test.com/test_resources/1"},
			serviceDown:     true,
			expectedError:   customerrors.OneloginErrorWrapper("ol http service", errors.New("unable to connect")),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(services.HTTPServiceConfig{
				Client: MockClient{
					DoFunc: func(req *http.Request) (*http.Response, error) {
						if req.URL.Path == "/auth/oauth2/v2/token" {
							if test.refreshTokenFail {
								return authFailThrough()
							}
							return authPassThrough()
						}
						if test.staleToken {
							test.staleToken = test.refreshTokenFail
							return authFailThrough()
						}
						if test.badCredential {
							return authFailThrough()
						}
						if test.badRequest {
							r := ioutil.NopCloser(bytes.NewReader([]byte(`bad request`)))
							return &http.Response{StatusCode: 400, Body: r}, nil
						}
						if test.serviceDown {
							r := ioutil.NopCloser(bytes.NewReader([]byte(`unable to connect`)))
							return &http.Response{StatusCode: 500, Body: r}, nil
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
