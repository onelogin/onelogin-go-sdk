// Package client provides intergation with api calls.
package client

import (
	"errors"
	"fmt"
)

// constants for the client config.
const (
	USRegion        = "us"
	EURegion        = "eu"
	BaseURLTemplate = "https://api.%s.onelogin.com"
	DefaultTimeout  = 5
)

var (
	errRegion            = errors.New("region is missing or unsupported")
	errClientIDEmpty     = errors.New("client_id is missing")
	errClientSecretEmpty = errors.New("client_secret is missing")
)

// APIClientConfig is the configuration for the APIClient.
type APIClientConfig struct {
	Timeout      int
	ClientID     string
	ClientSecret string
	Region       string
	Url          string
	BaseURL      string
}

func (cfg *APIClientConfig) Validate() (*APIClientConfig, error) {

	// Validate clientID
	if len(cfg.ClientID) == 0 {
		return cfg, errClientIDEmpty
	}

	// Validate clientSecret
	if len(cfg.ClientSecret) == 0 {
		return cfg, errClientSecretEmpty
	}
	// Validate the region if no url given
	if !isSupportedRegion(cfg.Region) && len(cfg.Url) == 0 {
		return cfg, errRegion
	}

	// Create the BaseURL from the region and template unless a BaseURL is given
	if len(cfg.Url) == 0 {
		cfg.BaseURL = fmt.Sprintf(BaseURLTemplate, cfg.Region)
	} else {
		cfg.BaseURL = cfg.Url
	}

	// Validate the timeout
	if cfg.Timeout == 0 {
		cfg.Timeout = DefaultTimeout
	}

	return cfg, nil
}

func isSupportedRegion(region string) bool {
	return region == EURegion || region == USRegion
}
