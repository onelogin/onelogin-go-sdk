package services

import "net/http"

// Repository is a thing that stores and retireves data using these methods. This
// could represent a HTTP backend service or a local database or file where business
// entities live such as apps or users
type Repository interface {
	Create(r interface{}, o interface{}) error
	Read(r interface{}, o interface{}) error
	Update(r interface{}, o interface{}) error
	Destroy(r interface{}, o interface{}) error
}

type SimpleQuery interface {
	Query(address string, o interface{}) error
}

type ResponseData struct {
	Data interface{}
}

// HTTPClient is a thing that implements the Do method. This enables mocking a
// HTTPService's Client for testing
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPServiceConfig Represents configuration information needed to work with an authenticated HTTP remote service
type HTTPServiceConfig struct {
	ClientID, ClientSecret, BaseURL string
	Client                          HTTPClient
}
