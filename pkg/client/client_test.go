package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetBaseUrWithUsRegion(t *testing.T) {
	region := US_REGION
	expectedBaseUrl := "https://api.us.onelogin.com"

	cl := New(&ApiClientConfig{
		Region: region,
	})
	assert.Equal(t, region, cl.region)
	assert.Equal(t, expectedBaseUrl, cl.baseUrl)
}

func TestSetBaseUrWithEuRegion(t *testing.T) {
	region := EU_REGION
	expectedBaseUrl := "https://api.eu.onelogin.com"

	cl := New(&ApiClientConfig{
		Region: region,
	})
	assert.Equal(t, region, cl.region)
	assert.Equal(t, expectedBaseUrl, cl.baseUrl)
}

func TestSetBaseUrWithNotSupportedRegion(t *testing.T) {
	region := "es"
	expectedBaseUrl := "https://api.us.onelogin.com"

	cl := New(&ApiClientConfig{
		Region: region,
	})
	assert.Equal(t, region, cl.region)
	assert.Equal(t, expectedBaseUrl, cl.baseUrl)
}

func TestNewProducesAllServices(t *testing.T) {
	region := US_REGION

	cl := New(&ApiClientConfig{
		Region: region,
	})
	assert.NotNil(t, cl.Services)
	assert.NotNil(t, cl.Services.AppsV1)
}

func TestNewProducesProperDefualtTimeout(t *testing.T) {
	region := US_REGION

	cl := New(&ApiClientConfig{
		Region: region,
	})
	assert.Equal(t, (time.Second * DEFAULT_TIMEOUT), cl.client.Timeout)
}

func TestNewProducesProperProvidedTimeout(t *testing.T) {
	region := US_REGION
	timeoutNumb := 10
	cl := New(&ApiClientConfig{
		Region:           region,
		TimeoutInSeconds: timeoutNumb,
	})
	assert.NotNil(t, cl.Services)
	assert.Equal(t, (time.Second * 10), cl.client.Timeout)
}
