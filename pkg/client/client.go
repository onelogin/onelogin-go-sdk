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
	US_REGION = "us"
	EU_REGION = "eu"
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
	AppsV1 *services.AppsV1
}

type ApiClientConfig struct {
	TimeoutInSeconds int
	ClientId         string
	ClientSecret     string
	Region           string
}

// New takes in params needed for the new api client, and
// return the new api client with services.
func New(cfg *ApiClientConfig) *ApiClient {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	baseUrl := setBaseUrl(cfg.Region)

	return &ApiClient{
		clientId:     cfg.ClientId,
		clientSecret: cfg.ClientSecret,
		region:       cfg.Region,
		baseUrl:      baseUrl,
		client:       httpClient,
		Services: &Services{
			AppsV1: services.NewAppsV1(baseUrl, httpClient),
		},
	}
}

// setBaseUrl takes in a region to set the proper url,
// and returns the url for the provided region, with "us" by default.
func setBaseUrl(region string) string {
	regionToUse := strings.ToLower(region)

	if regionToUse != EU_REGION {
		regionToUse = US_REGION
	}

	return fmt.Sprintf("https://api.%s.onelogin.com", regionToUse)
}
