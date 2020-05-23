package olhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/utils"
)

const resourceRequestuestContext = "ol request"

var errInvalidRequestInput = errors.New("Invalid input for request creation")

type OLHTTPService struct {
	ErrorContext     string
	Config           services.HTTPServiceConfig
	ClientCredential ClientCredential
}

// New uses the cfg to generate the new auth service, and returns
// the created auth service for version 2.
func New(cfg services.HTTPServiceConfig) *OLHTTPService {
	return &OLHTTPService{
		Config:           cfg,
		ErrorContext:     resourceRequestuestContext,
		ClientCredential: ClientCredential{},
	}
}

// Create creates a new resource in the remote location over HTTP
func (svc OLHTTPService) Create(r interface{}) ([]byte, error) {
	resourceRequest := r.(OLHTTPRequest)
	var (
		req    *http.Request
		reqErr error
	)
	if resourceRequest.Payload != nil {
		bodyToSend, marshErr := json.Marshal(resourceRequest.Payload)
		if marshErr != nil {
			return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, marshErr)
		}
		req, reqErr = http.NewRequest(http.MethodPost, resourceRequest.URL, bytes.NewBuffer(bodyToSend))
	} else {
		req, reqErr = http.NewRequest(http.MethodPost, resourceRequest.URL, bytes.NewBuffer([]byte("")))
	}
	if reqErr != nil {
		return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, reqErr)
	}
	_, data, err := svc.executeHTTP(req, resourceRequest)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// Read retrieves a resource from remote over HTTP
func (svc OLHTTPService) Read(r interface{}) ([]byte, error) {
	resourceRequest := r.(OLHTTPRequest)
	req, reqErr := http.NewRequest(http.MethodGet, resourceRequest.URL, nil)
	if reqErr != nil {
		return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, reqErr)
	}

	if resourceRequest.Payload != nil {
		params := req.URL.Query()
		b, err := json.Marshal(resourceRequest.Payload)
		if err != nil {
			return nil, err
		}
		var m map[string]string
		if err = json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		for k, v := range m {
			if v != "" {
				params.Add(utils.ToSnakeCase(k), v)
			}
		}
		req.URL.RawQuery = params.Encode()
	}

	var (
		allData []byte
		next    string
	)
	resp, data, err := svc.executeHTTP(req, resourceRequest)
	if err != nil {
		return []byte{}, err
	}
	for {
		allData = append(allData, data...)
		next = resp.Header.Get("After-Cursor")
		if next == "" || err != nil {
			break
		}
		params := req.URL.Query()
		params.Add("Cursor", next)
		req.URL.RawQuery = params.Encode()
		resp, data, err = svc.executeHTTP(req, resourceRequest)
	}

	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// Update updates a resource in its remote location over HTTP
func (svc OLHTTPService) Update(r interface{}) ([]byte, error) {
	resourceRequest := r.(OLHTTPRequest)
	var (
		req    *http.Request
		reqErr error
	)
	if resourceRequest.Payload != nil {
		bodyToSend, marshErr := json.Marshal(resourceRequest.Payload)
		if marshErr != nil {
			return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, marshErr)
		}
		req, reqErr = http.NewRequest(http.MethodPut, resourceRequest.URL, bytes.NewBuffer(bodyToSend))
	} else {
		req, reqErr = http.NewRequest(http.MethodPut, resourceRequest.URL, bytes.NewBuffer([]byte("")))
	}
	if reqErr != nil {
		return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, reqErr)
	}
	_, data, err := svc.executeHTTP(req, resourceRequest)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// Destroy executes a HTTP destroy and removes the resource from its location in a remote
func (svc OLHTTPService) Destroy(r interface{}) ([]byte, error) {
	resourceRequest := r.(OLHTTPRequest)
	req, reqErr := http.NewRequest(http.MethodDelete, resourceRequest.URL, nil)
	if reqErr != nil {
		return nil, customerrors.OneloginErrorWrapper(resourceRequestuestContext, reqErr)
	}
	_, data, err := svc.executeHTTP(req, resourceRequest)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// attaches http request headers supplied by caller and auth headers depending
// on the request's auth type (e.g. bearer or basic)
func (svc *OLHTTPService) attachHeaders(req *http.Request, resourceRequest OLHTTPRequest) error {
	// set headers
	for key, val := range resourceRequest.Headers {
		req.Header.Set(key, val)
	}
	switch strings.ToLower(resourceRequest.AuthMethod) {
	case "bearer":
		if err := svc.authorize(); err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *svc.ClientCredential.AccessToken))
	case "basic":
		req.SetBasicAuth(svc.Config.ClientID, svc.Config.ClientSecret)
	default:
		// no auth headers
	}
	return nil
}

// executes the http request, initiates retry on expired bearer tokens and returns the
// response's byte array resource representation
func (svc *OLHTTPService) executeHTTP(req *http.Request, resourceRequest OLHTTPRequest) (*http.Response, []byte, error) {
	if err := svc.attachHeaders(req, resourceRequest); err != nil {
		return nil, nil, err
	}
	resp, err := svc.Config.Client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == 401 && resourceRequest.AuthMethod == "bearer" {
		if err := setBearerToken(svc); err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.executeHTTP(req, resourceRequest)
	}
	if err != nil {
		return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err)
	}
	return resp, responseData, nil
}

// requests a fresh access token
func (svc *OLHTTPService) mintBearerToken() (ClientCredential, error) {
	resp, err := svc.Create(OLHTTPRequest{
		URL:        fmt.Sprintf("%s/auth/oauth2/v2/token", svc.Config.BaseURL),
		Headers:    map[string]string{"Content-Type": "application/json"},
		Payload:    AuthBody{GrantType: "client_credentials"},
		AuthMethod: "basic",
	})
	if err != nil {
		return ClientCredential{}, err
	}

	var output ClientCredential
	if err = json.Unmarshal(resp, &output); err != nil {
		return ClientCredential{}, err
	}
	return output, nil
}

// initializes the service's access token memoization
func (svc *OLHTTPService) authorize() error {
	if (svc.ClientCredential == ClientCredential{}) {
		if err := setBearerToken(svc); err != nil {
			return err
		}
	}
	return nil
}

// force overwrite the service's memoized access token
func setBearerToken(svc *OLHTTPService) error {
	cred, err := svc.mintBearerToken()
	if err != nil {
		return err
	}
	svc.ClientCredential = cred
	return nil
}
