// Package client provides intergation with api calls.
package client

import (
	"errors"
)

// constants for the client.
const (
	USRegion       = "us"
	EURegion       = "eu"
	BaseUrl        = "https://api.us.onelogin.com"
	DefaultTimeout = 5
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
}

func (cfg APIClientConfig) Validate() (APIClientConfig, error) {

	// Validate clientID
	if len(cfg.ClientID) == 0 {
		return cfg, errClientIDEmpty
	}

	// Validate clientSecret
	if len(cfg.ClientSecret) == 0 {
		return cfg, errClientSecretEmpty
	}
	// validate the region if no url given
	if !isSupportedRegion(cfg.Region) && len(cfg.Url) == 0 {
		return cfg, errRegion
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
