// Package client provides intergation with api calls.
package client

import (
	"net/http"
	"time"

	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/apps"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/session_login_tokens"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/user_mappings"
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
	HTTPService          olhttp.OLHTTPService
	AppsV2               apps.V2Service
	UserMappingsV2       usermappings.V2Service
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

	repo := olhttp.New(services.HTTPServiceConfig{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		BaseURL:      cfg.Url,
		Client:       httpClient,
	})

	appV2Service := apps.New(repo, cfg.Url)
	userMappingsV2Service := usermappings.New(repo, cfg.Url)
	sessionLoginTokenV1Service := sessionlogintokens.New(repo, cfg.Url)

	return &APIClient{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		region:       cfg.Region,
		baseURL:      cfg.Url,
		client:       httpClient,
		Services: &Services{
			HTTPService:          *repo,
			AppsV2:               appV2Service,
			UserMappingsV2:       userMappingsV2Service,
			SessionLoginTokensV1: sessionLoginTokenV1Service,
		},
	}, nil
}
