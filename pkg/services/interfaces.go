package services

import "net/http"

// Authenticator is an interface.
type Authenticator interface {
	Authorize() (string, error)
	ReAuthorize() (string, error)
}

type Repository interface {
	Create(r interface{}) ([]byte, error)
	Read(r interface{}) ([]byte, error)
	Update(r interface{}) ([]byte, error)
	Destroy(r interface{}) ([]byte, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ResourceServiceConfig struct {
	Repository Repository
}

type HTTPServiceConfig struct {
	ClientID, ClientSecret, BaseURL string
	Client                          HTTPClient
}
