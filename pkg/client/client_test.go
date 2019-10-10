package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	usExpectedBaseUrl   = "https://api.us.onelogin.com"
	euExpectedBaseUrl   = "https://api.eu.onelogin.com"
	notSupporttedRegion = "Test"
)

func TestSetBaseUrl(t *testing.T) {
	tests := map[string]struct {
		region          string
		expectedBaseUrl string
	}{
		"us region": {
			region:          US_REGION,
			expectedBaseUrl: usExpectedBaseUrl,
		},
		"eu region": {
			region:          EU_REGION,
			expectedBaseUrl: euExpectedBaseUrl,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			baseUrl := setBaseUrl(test.region)
			assert.Equal(t, test.expectedBaseUrl, baseUrl)
		})
	}
}

func TestIsSupportedRegion(t *testing.T) {
	tests := map[string]struct {
		region             string
		expectedValidation bool
	}{
		"us region": {
			region:             US_REGION,
			expectedValidation: true,
		},
		"eu region": {
			region:             EU_REGION,
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
	assert.Equal(t, US_REGION, defRegion)
}

func TestNew(t *testing.T) {
	clientId := "test"
	clientSecret := "test"

	tests := map[string]struct {
		region               string
		timeout              int
		clientId             string
		clientSecret         string
		expectedRegion       string
		expectedBaseUrl      string
		expectedTimeout      time.Duration
		expectedClientId     string
		expectedClientSecret string
	}{
		"creates the expected default timeout": {
			timeout:              0,
			region:               US_REGION,
			clientId:             clientId,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(DEFAULT_TIMEOUT),
			expectedRegion:       US_REGION,
			expectedBaseUrl:      usExpectedBaseUrl,
			expectedClientId:     clientId,
			expectedClientSecret: clientSecret,
		},
		"has the provided timeout": {
			timeout:              2,
			region:               US_REGION,
			clientId:             clientId,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       US_REGION,
			expectedBaseUrl:      usExpectedBaseUrl,
			expectedClientId:     clientId,
			expectedClientSecret: clientSecret,
		},
		"has the region provided baseUrl": {
			timeout:              2,
			region:               EU_REGION,
			clientId:             clientId,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       EU_REGION,
			expectedBaseUrl:      euExpectedBaseUrl,
			expectedClientId:     clientId,
			expectedClientSecret: clientSecret,
		},
		"has the default baseurl and region": {
			timeout:              2,
			region:               notSupporttedRegion,
			clientId:             clientId,
			clientSecret:         clientSecret,
			expectedTimeout:      time.Second * time.Duration(2),
			expectedRegion:       US_REGION,
			expectedBaseUrl:      usExpectedBaseUrl,
			expectedClientId:     clientId,
			expectedClientSecret: clientSecret,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cl := New(&ApiClientConfig{
				Region:           test.region,
				TimeoutInSeconds: test.timeout,
				ClientId:         test.clientId,
				ClientSecret:     test.clientSecret,
			})

			assert.Equal(t, test.expectedRegion, cl.region)
			assert.Equal(t, test.expectedBaseUrl, cl.baseUrl)
			assert.Equal(t, test.expectedClientId, cl.clientId)
			assert.Equal(t, test.expectedClientSecret, cl.clientSecret)
			assert.NotNil(t, cl.client)
			assert.Equal(t, test.expectedTimeout, cl.client.Timeout)
			assert.NotNil(t, cl.Services)
			assert.NotNil(t, cl.Services.AppsV2)
			assert.NotNil(t, cl.Services.AuthV2)
		})
	}
}
