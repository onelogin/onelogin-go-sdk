package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	usExpectedBaseURL   = "https://api.us.onelogin.com"
	euExpectedBaseURL   = "https://api.eu.onelogin.com"
	notSupporttedRegion = "Test"
)

func TestSetBaseUrl(t *testing.T) {
	tests := map[string]struct {
		region          string
		expectedBaseURL string
	}{
		"us region": {
			region:          USregion,
			expectedBaseURL: usExpectedBaseURL,
		},
		"eu region": {
			region:          EUregion,
			expectedBaseURL: euExpectedBaseURL,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			baseURL := setBaseURL(test.region)
			assert.Equal(t, test.expectedBaseURL, baseURL)
		})
	}
}

func TestIsSupportedRegion(t *testing.T) {
	tests := map[string]struct {
		region             string
		expectedValidation bool
	}{
		"us region": {
			region:             USregion,
			expectedValidation: true,
		},
		"eu region": {
			region:             EUregion,
			expectedValidation: true,
		},
		"not supported": {
			region:             notSupporttedRegion,
			expectedValidation: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isValid := isSupportedRegion(test.region)
			assert.Equal(t, test.expectedValidation, isValid)
		})
	}
}

func TestGetDefaultRegion(t *testing.T) {
	defRegion := getDefaultRegion()
	assert.Equal(t, USregion, defRegion)
}

func TestNew(t *testing.T) {
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
		"creates the expected default timeout": {
			timeout:              0,
			region:               USregion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(DefaultTimeout),
			expectedRegion:       USregion,
			expectedBaseURL:      usExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
		"has the provided timeout": {
			timeout:              2,
			region:               USregion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       USregion,
			expectedBaseURL:      usExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
		"has the region provided baseUrl": {
			timeout:              2,
			region:               EUregion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       EUregion,
			expectedBaseURL:      euExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
		"has the default baseurl and region": {
			timeout:              2,
			region:               notSupporttedRegion,
			clientID:             clientID,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       USregion,
			expectedBaseURL:      usExpectedBaseURL,
			expectedClientID:     clientID,
			expectedClientSecret: clientSecret,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cl := New(&APIClientConfig{
				Region:           test.region,
				TimeoutInSeconds: test.timeout,
				ClientID:         test.clientID,
				ClientSecret:     test.clientSecret,
			})

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
