// Package client provides intergation with api calls.
package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
)

// constants for the client.
const (
	BaseURLTemplate = "https://api.%s.onelogin.com"
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
	AppsV2 *services.AppsV2
	AuthV2 *services.AuthV2
}

// New uses the config to generate the api client with services attached, and returns
// the new api client.
func NewClient(cfg APIClientConfig) *APIClient {
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(cfg.timeout),
	}

	baseURL := fmt.Sprintf(BaseURLTemplate, cfg.region)

	authV2Service := services.NewAuthV2(&services.AuthConfigV2{
		ClientID:     cfg.clientID,
		ClientSecret: cfg.clientSecret,
		BaseURL:      baseURL,
		Client:       httpClient,
	})

	appv2Service := services.NewAppsV2(&services.AppsV2Config{
		BaseURL: baseURL,
		Client:  httpClient,
		Auth:    authV2Service,
	})

	return &APIClient{
		clientID:     cfg.clientID,
		clientSecret: cfg.clientSecret,
		region:       cfg.region,
		baseURL:      baseURL,
		client:       httpClient,
		Services: &Services{
			AppsV2: appv2Service,
			AuthV2: authV2Service,
		},
	}
}
