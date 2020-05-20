package services

import "net/http"

// Authenticator is an interface.
type Authenticator interface {
	Authorize() (string, error)
	ReAuthorize() (string, error)
}

type Creator interface {
	Create(newItem interface{}) (http.Response, interface{}, error)
}

type Updater interface {
	Update(id int, newState interface{}) (http.Response, interface{}, error)
}

type Destroyer interface {
	Destroy(id int) (http.Response, error)
}

type Querier interface {
	Query(params interface{}) (http.Response, interface{}, error)
}

type Getter interface {
	GetOne(id int) (http.Response, interface{}, error)
}

type ResourceReader interface {
	Querier
	Getter
}

type ResourceWriter interface {
	Creator
	Updater
}

type CRUD interface {
	ResourceWriter
	ResourceReader
	Destroyer
}

type APIServiceConfig struct {
	BaseURL string
	Client  *http.Client
	Auth    Authenticator
}

type AuthServiceConfig struct {
	ClientID, ClientSecret, BaseURL string
	Client                          *http.Client
}
