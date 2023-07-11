package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientIDValidation(t *testing.T) {
	invalidClientID := ""
	validClientID := "test"
	clientSecret := "test"

	t.Run("invalidClientID", func(t *testing.T) {
		config := &APIClientConfig{
			ClientID:     invalidClientID,
			ClientSecret: clientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Initialize()
		assert.NotNil(t, err)
	})

	t.Run("validClientID", func(t *testing.T) {
		config := &APIClientConfig{
			ClientID:     validClientID,
			ClientSecret: clientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Initialize()
		assert.Nil(t, err)
		assert.Equal(t, config.ClientID, validClientID)
	})
}

func TestClientSecretValidation(t *testing.T) {
	invalidClientSecret := ""
	validClientSecret := "test"
	clientID := "test"

	t.Run("invalidClientSecret", func(t *testing.T) {
		config := &APIClientConfig{
			ClientID:     clientID,
			ClientSecret: invalidClientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Initialize()
		assert.NotNil(t, err)
	})

	t.Run("validClientSecret", func(t *testing.T) {
		config := &APIClientConfig{
			ClientID:     clientID,
			ClientSecret: validClientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Initialize()
		assert.Nil(t, err)
		assert.Equal(t, config.ClientSecret, validClientSecret)
	})

}

func TestTimeoutBehavior(t *testing.T) {
	clientID := "test"
	clientSecret := "test"
	testTimeout := 2
	timeoutDefaultTrigger := 0

	tests := map[string]struct {
		timeout         int
		expectedTimeout int
	}{
		"has the provided timeout": {
			timeout:         testTimeout,
			expectedTimeout: testTimeout,
		},
		"uses the expected default timeout": {
			timeout:         timeoutDefaultTrigger,
			expectedTimeout: DefaultTimeout,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := &APIClientConfig{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Region:       USRegion,
				Timeout:      test.timeout,
			}
			config, err := config.Initialize()
			assert.Nil(t, err)
			assert.Equal(t, test.expectedTimeout, config.Timeout)
		})
	}
}

func TestRegionValidation(t *testing.T) {
	notSupporttedRegion := "Test"
	tests := map[string]struct {
		region             string
		expectedValidation bool
	}{
		"us region": {
			region:             USRegion,
			expectedValidation: true,
		},
		"eu region": {
			region:             EURegion,
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

func TestUrlConstruction(t *testing.T) {
	clientID := "test"
	clientSecret := "test"
	usexpectedUrl := "https://api.us.onelogin.com"
	euexpectedUrl := "https://api.eu.onelogin.com"
	shadow := "https://oapi.onelogin-shadow01.com"

	tests := map[string]struct {
		region      string
		expectedUrl string
		url         string
	}{
		"region given, no url honors region": {
			region:      USRegion,
			expectedUrl: usexpectedUrl,
		},
		"url given, no region honors url": {
			url:         shadow,
			expectedUrl: shadow,
		},
		"url and region given honors url given": {
			region:      USRegion,
			url:         euexpectedUrl,
			expectedUrl: euexpectedUrl,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := &APIClientConfig{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Region:       test.region,
				Url:          test.url,
			}
			config, err := config.Initialize()
			assert.Nil(t, err)
			assert.Equal(t, test.expectedUrl, config.Url)
		})
	}
}

func TestUrlRegionValidation(t *testing.T) {
	clientID := "test"
	clientSecret := "test"

	tests := map[string]struct {
		expectedUrl string
	}{
		"no region given, no url returns error": {
			expectedUrl: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			config := &APIClientConfig{
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
			config, err := config.Initialize()
			assert.NotNil(t, err)
			assert.Equal(t, test.expectedUrl, "")
		})
	}
}
