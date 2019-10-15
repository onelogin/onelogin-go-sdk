/*
 * Onelogin api client.
 */
package client

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
)

// constants for the client.
const (
	USregion       = "us"
	EUregion       = "eu"
	DefaultTimeout = 5
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

// APIClientConfig is the configuration for the APIClient.
type APIClientConfig struct {
	TimeoutInSeconds int
	ClientID         string
	ClientSecret     string
	Region           string
}

// New uses the config to generate the api client with services attached, and returns
// the new api client.
func New(cfg *APIClientConfig) *APIClient {
	timeout := cfg.TimeoutInSeconds
	regionToUse := cfg.Region

	if cfg.TimeoutInSeconds == 0 {
		timeout = DefaultTimeout
	}

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	isValidRegion := isSupportedRegion(regionToUse)

	if !isValidRegion {
		regionToUse = getDefaultRegion()
	}

	baseURL := setBaseUrl(regionToUse)

	authV2Service := services.NewAuthV2(&services.AuthConfigV2{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		BaseURL:      baseURL,
		Client:       httpClient,
	})

	appv2Service := services.NewAppsV2(&services.AppsV2Config{
		BaseURL: baseURL,
		Client:  httpClient,
		Auth:    authV2Service,
	})

	return &APIClient{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		region:       regionToUse,
		baseURL:      baseURL,
		client:       httpClient,
		Services: &Services{
			AppsV2: appv2Service,
			AuthV2: authV2Service,
		},
	}
}

// isSupportedRegion validates whether a region is supported, and returns
// a boolean indicating whether it is or not.
func isSupportedRegion(region string) bool {
	switch strings.ToLower(region) {
	case USregion:
		return true
	case EUregion:
		return true
	default:
		return false
	}
}

// getDefaultRegion grabs the default region, and returns it.
func getDefaultRegion() string {
	return USregion
}

// setBaseUrl generates the proper base url based on region, if supported,and returns
// the base url for the provided region, or 'us' by default.
func setBaseUrl(region string) string {
	regionToUse := strings.ToLower(region)
	return fmt.Sprintf("https://api.%s.onelogin.com", regionToUse)
}
