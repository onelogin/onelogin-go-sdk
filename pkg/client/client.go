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

const (
	US_REGION       = "us"
	EU_REGION       = "eu"
	DEFAULT_TIMEOUT = 5
)

type ApiClient struct {
	clientId     string
	clientSecret string
	region       string
	token        string
	baseUrl      string
	client       *http.Client
	Services     *Services
}

type Services struct {
	AppsV2 *services.AppsV2
	AuthV2 *services.AuthV2
}

type ApiClientConfig struct {
	TimeoutInSeconds int
	ClientId         string
	ClientSecret     string
	Region           string
}

// New uses the config to generate the api client with services attached, and returns
// the new api client.
func New(cfg *ApiClientConfig) *ApiClient {
	timeout := cfg.TimeoutInSeconds
	regionToUse := cfg.Region

	if cfg.TimeoutInSeconds == 0 {
		timeout = DEFAULT_TIMEOUT
	}

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	isValidRegion := isSupportedRegion(regionToUse)

	if !isValidRegion {
		regionToUse = getDefaultRegion()
	}

	baseUrl := setBaseUrl(regionToUse)

	authV2Service := services.NewAuthV2(&services.AuthConfigV2{
		ClientId:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		BaseUrl:      baseUrl,
		Client:       httpClient,
	})

	appv2Service := services.NewAppsV2(&services.AppsV2Config{
		BaseUrl: baseUrl,
		Client:  httpClient,
		Auth:    authV2Service,
	})

	return &ApiClient{
		clientId:     cfg.ClientId,
		clientSecret: cfg.ClientSecret,
		region:       regionToUse,
		baseUrl:      baseUrl,
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
	case US_REGION:
		return true
	case EU_REGION:
		return true
	default:
		return false
	}
}

// getDefaultRegion grabs the default region, and returns it.
func getDefaultRegion() string {
	return US_REGION
}

// setBaseUrl generates the proper base url based on region, if supported,and returns
// the base url for the provided region, or 'us' by default.
func setBaseUrl(region string) string {
	regionToUse := strings.ToLower(region)
	return fmt.Sprintf("https://api.%s.onelogin.com", regionToUse)
}
