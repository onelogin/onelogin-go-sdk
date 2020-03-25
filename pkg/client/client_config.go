// Package client provides intergation with api calls.
package client

import (
	"errors"
)

// constants for the client.
const (
	USRegion       = "us"
	EURegion       = "eu"
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
}

func NewConfig(clientID string, clientSecret string, region string, timeout int) (APIClientConfig, error) {
	var config APIClientConfig

	// Validate clientID
	if len(clientID) == 0 {
		return config, errClientIDEmpty
	}

	// Validate clientSecret
	if len(clientSecret) == 0 {
		return config, errClientSecretEmpty
	}
	// validate the region
	if !isSupportedRegion(region) {
		return config, errRegion
	}

	// Validate the timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// Build the new client config
	config = APIClientConfig{
		timeout:      timeout,
		clientID:     clientID,
		clientSecret: clientSecret,
		region:       region,
	}

	return config, nil
}

func isSupportedRegion(region string) bool {
	return region == EURegion || region == USRegion
}
