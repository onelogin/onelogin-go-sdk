package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	clientID := "test"
	clientSecret := "test"
	usExpectedBaseURL := "https://api.us.onelogin.com"
	euExpectedBaseURL := "https://api.eu.onelogin.com"

	tests := map[string]struct {
		region               string
		timeout              int
		clientID             string
		clientSecret         string
		expectedRegion       string
		expectedBaseURL      string
		expectedTimeout      time.Duration
		expectedClientID     string
		expectedClientSecret string
	}{
		"has the configured timeout": {
			timeout:              2,
			region:               USRegion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       USRegion,
			expectedBaseURL:      usExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
		"has the region provided baseUrl": {
			timeout:              2,
			region:               EURegion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       EURegion,
			expectedBaseURL:      euExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := &APIClientConfig{
				ClientID:     test.clientID,
				ClientSecret: test.clientSecret,
				Region:       test.region,
				Timeout:      test.timeout,
			}

			cl, err := NewClient(config)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedRegion, cl.region)
			assert.Equal(t, test.expectedBaseURL, cl.baseURL)
			assert.Equal(t, test.expectedClientID, cl.clientID)
			assert.Equal(t, test.expectedClientSecret, cl.clientSecret)
			assert.NotNil(t, cl.client)
			assert.Equal(t, test.expectedTimeout, cl.client.Timeout)
			assert.NotNil(t, cl.Services)
			assert.NotNil(t, cl.Services.AppsV2)
		})
	}
}

func TestInvalidNewClient(t *testing.T) {
	invalidClientID := ""
	clientSecret := "test"
	tests := map[string]struct {
		region               string
		timeout              int
		clientID             string
		clientSecret         string
		expectedRegion       string
		expectedBaseURL      string
		expectedTimeout      time.Duration
		expectedClientID     string
		expectedClientSecret string
	}{
		"invalid config due to missing client id": {
			region:       USRegion,
			clientID:     invalidClientID,
			clientSecret: clientSecret,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := &APIClientConfig{
				ClientID:     test.clientID,
				ClientSecret: test.clientSecret,
				Region:       test.region,
			}
			cl, err := NewClient(config)
			assert.NotNil(t, err)
			assert.Equal(t, &APIClient{}, cl)
		})
	}
}
