// Package client provides intergation with api calls.
package client

import (
	"net/http"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/internal/services/apps"
	"github.com/onelogin/onelogin-go-sdk/internal/services/client_credentials"
	"github.com/onelogin/onelogin-go-sdk/internal/services/session_login_tokens"
)

// APIClient is used to communicate with the available api services.
type APIClient struct {
	clientID     string
	clientSecret string
	region       string
	baseURL      string
	client       *http.Client
	Services     *Services
}

// Services contains all the available api services.
type Services struct {
	AppsV2               apps.V2Service
	AuthV2               clientcredentials.V2Service
	SessionLoginTokensV1 sessionlogintokens.V1Service
}

// NewClient uses the config to generate the api client with services attached, and returns
// the new api client.
func NewClient(cfg *APIClientConfig) (*APIClient, error) {
	cfg, err := cfg.Initialize()
	if err != nil {
		return &APIClient{}, err
	}

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(cfg.Timeout),
	}

	authV2Service := clientcredentials.New(&services.AuthServiceConfig{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		BaseURL:      cfg.Url,
	})

	apiServiceConfig := services.APIServiceConfig{
		BaseURL: cfg.Url,
		Client:  httpClient,
		Auth:    &authV2Service,
	}

	appV2Service := apps.New(&apiServiceConfig)
	sessionLoginTokenV1Service := sessionlogintokens.New(&apiServiceConfig)

	return &APIClient{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		region:       cfg.Region,
		baseURL:      cfg.Url,
		client:       httpClient,
		Services: &Services{
			AppsV2:               appV2Service,
			AuthV2:               authV2Service,
			SessionLoginTokensV1: sessionLoginTokenV1Service,
		},
	}, nil
}
