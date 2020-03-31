package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	usExpectedBaseURL = "https://api.us.onelogin.com"
	euExpectedBaseURL = "https://api.eu.onelogin.com"
	shadow            = "https://oapi.onelogin-shadow01.com"
)

func TestClientBaseURL(t *testing.T) {
	defaultClientId := "test"
	defaultClientSecret := "test"
	defaultTimeout := 0

	tests := map[string]struct {
		region          string
		expectedBaseURL string
		url             string
	}{
		"us region": {
			region:          USRegion,
			expectedBaseURL: usExpectedBaseURL,
		},
		"eu region": {
			region:          EURegion,
			expectedBaseURL: euExpectedBaseURL,
		},
		"url given, no region": {
			url:             shadow,
			expectedBaseURL: shadow,
		},
		"url and region given": {
			region:          USRegion,
			url:             euExpectedBaseURL,
			expectedBaseURL: usExpectedBaseURL,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := APIClientConfig{
				ClientID:     defaultClientId,
				ClientSecret: defaultClientSecret,
				Region:       test.region,
				Url:          test.url,
				Timeout:      defaultTimeout,
			}
			config, err := config.Validate()
			client := NewClient(config)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedBaseURL, client.baseURL)
		})
	}
}

func TestNewClient(t *testing.T) {
	clientID := "test"
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
			config := APIClientConfig{
				ClientID:     test.clientID,
				ClientSecret: test.clientSecret,
				Region:       test.region,
				Timeout:      test.timeout,
			}
			config, err := config.Validate()
			assert.Nil(t, err)
			cl := NewClient(config)

			assert.Equal(t, test.expectedRegion, cl.region)
			assert.Equal(t, test.expectedBaseURL, cl.baseURL)
			assert.Equal(t, test.expectedClientID, cl.clientID)
			assert.Equal(t, test.expectedClientSecret, cl.clientSecret)
			assert.NotNil(t, cl.client)
			assert.Equal(t, test.expectedTimeout, cl.client.Timeout)
			assert.NotNil(t, cl.Services)
			assert.NotNil(t, cl.Services.AppsV2)
			assert.NotNil(t, cl.Services.AuthV2)
		})
	}
}
