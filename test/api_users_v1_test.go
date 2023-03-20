/*
OneLogin API

Testing UsersV1ApiService

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

func Test_onelogin_UsersV1ApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersV1ApiService AddRolesToUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.AddRolesToUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService CreateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersV1Api.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		httpRes, err := apiClient.UsersV1Api.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService GetCustomAttributes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersV1Api.GetCustomAttributes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService GetUserApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.GetUserApps(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService GetUserById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.GetUserById(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService GetUserRoles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.GetUserRoles(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService ListUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersV1Api.ListUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService LockAccountUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.LockAccountUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService LogOutUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.LogOutUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService RemoveUserRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.RemoveUserRole(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService SetUserState", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.SetUserState(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService UpdatePasswordInsecure", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.UpdatePasswordInsecure(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService UpdatePasswordSecure", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.UpdatePasswordSecure(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersV1ApiService UpdateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId int32

		resp, httpRes, err := apiClient.UsersV1Api.UpdateUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}