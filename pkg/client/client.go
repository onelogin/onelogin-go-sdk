// Package client provides intergation with api calls.
package client

import (
	"net/http"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
)

// APIClient is used to communicate with the available api services.
type APIClient struct {
	clientID     string
	clientSecret string
	region       string
	token        string
	baseURL      string
	client       *http.Client
	Services     *Services
}

// Services contains all the available api services.
type Services struct {
	AppsV2               *services.AppsV2
	AuthV2               *services.AuthV2
	SessionLoginTokensV1 *services.SessionLoginTokenV1
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

	authV2Service := services.NewAuthV2(&services.AuthConfigV2{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		BaseURL:      cfg.Url,
		Client:       httpClient,
	})

	appV2Service := services.NewAppsV2(&services.AppsV2Config{
		BaseURL: cfg.Url,
		Client:  httpClient,
		Auth:    authV2Service,
	})

	sessionLoginTokenV1Service := services.NewSessionLoginTokenV1(&services.SessionLoginTokenV1Config{
		BaseURL: cfg.Url,
		Client:  httpClient,
		Auth:    authV2Service,
	})

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
