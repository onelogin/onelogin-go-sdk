package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	notSupporttedRegion = "Test"
)

func TestClientIDValidation(t *testing.T) {
	invalidClientID := ""
	validClientID := "test"
	clientSecret := "test"

	t.Run("invalidClientID", func(t *testing.T) {
		config := APIClientConfig{
			ClientID:     invalidClientID,
			ClientSecret: clientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Validate()
		assert.NotNil(t, err)
	})

	t.Run("validClientID", func(t *testing.T) {
		config := APIClientConfig{
			ClientID:     validClientID,
			ClientSecret: clientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Validate()
		assert.Nil(t, err)
		assert.Equal(t, config.ClientID, validClientID)
	})
}

func TestClientSecretValidation(t *testing.T) {
	invalidClientSecret := ""
	validClientSecret := "test"
	clientID := "test"

	t.Run("invalidClientSecret", func(t *testing.T) {
		config := APIClientConfig{
			ClientID:     clientID,
			ClientSecret: invalidClientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Validate()
		assert.NotNil(t, err)
	})

	t.Run("validClientSecret", func(t *testing.T) {
		config := APIClientConfig{
			ClientID:     clientID,
			ClientSecret: validClientSecret,
			Region:       USRegion,
			Timeout:      DefaultTimeout,
		}
		config, err := config.Validate()
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
			config := APIClientConfig{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Region:       USRegion,
				Timeout:      test.timeout,
			}
			config, err := config.Validate()
			assert.Nil(t, err)
			assert.Equal(t, test.expectedTimeout, config.Timeout)
		})
	}
}

func TestRegionValidation(t *testing.T) {
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