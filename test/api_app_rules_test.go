/*
OneLogin API

Testing AppRulesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package onelogin

import (
	"context"
	"testing"

	openapiclient "github.com/onelogin/onelogin-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_onelogin_AppRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppRulesApiService CreateAppRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.AppRulesApi.CreateAppRule(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService DeleteRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleId string

		httpRes, err := apiClient.AppRulesApi.DeleteRule(context.Background(), appId, ruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService GetAppRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleId string

		resp, httpRes, err := apiClient.AppRulesApi.GetAppRule(context.Background(), appId, ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListActionValies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleActionValue string

		resp, httpRes, err := apiClient.AppRulesApi.ListActionValies(context.Background(), appId, ruleActionValue).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListActions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.AppRulesApi.ListActions(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListAppRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.AppRulesApi.ListAppRules(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListConditionOperators", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleConditionValue string

		resp, httpRes, err := apiClient.AppRulesApi.ListConditionOperators(context.Background(), appId, ruleConditionValue).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListConditionValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleConditionValue string

		resp, httpRes, err := apiClient.AppRulesApi.ListConditionValues(context.Background(), appId, ruleConditionValue).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService ListConditions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.AppRulesApi.ListConditions(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService SortAppRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.AppRulesApi.SortAppRules(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppRulesApiService UpdateAppRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var ruleId string

		resp, httpRes, err := apiClient.AppRulesApi.UpdateAppRule(context.Background(), appId, ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}