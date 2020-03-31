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
	timeout      int
	clientID     string
	clientSecret string
	region       string
	url          string
}

func ValidateConfig(cfg APIClientConfig) (APIClientConfig, error) {

	// Validate clientID
	if len(cfg.clientID) == 0 {
		return cfg, errClientIDEmpty
	}

	// Validate clientSecret
	if len(cfg.clientSecret) == 0 {
		return cfg, errClientSecretEmpty
	}
	// validate the region if no url given
	if !isSupportedRegion(cfg.region) && len(cfg.url) == 0 {
		return cfg, errRegion
	}

	// Validate the timeout
	if cfg.timeout == 0 {
		cfg.timeout = DefaultTimeout
	}

	return cfg, nil
}

func isSupportedRegion(region string) bool {
	return region == EURegion || region == USRegion
}
