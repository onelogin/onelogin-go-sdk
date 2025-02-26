package main

import (
	"testing"

	"onelogin-test/mocks"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserOperations(t *testing.T) {
	// Create mock client
	mockClient := new(mocks.MockClient)

	// Setup test data
	email := "test@example.com"
	testUser := models.User{
		ID:        1,
		Email:     email,
		Username:  "testuser",
		Firstname: "Test",
		Lastname:  "User",
	}

	// Setup mock expectations
	mockClient.On("CreateUser", mock.AnythingOfType("models.User")).Return(testUser, nil)
	mockClient.On("GetUsers", mock.AnythingOfType("*models.UserQuery")).Return([]models.User{testUser}, nil)

	// Test user creation
	createdUser, err := mockClient.CreateUser(models.User{
		Firstname: "Test",
		Lastname:  "User",
		Email:     email,
	})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, testUser.Email, createdUser.Email)

	// Test user query
	query := &models.UserQuery{
		Email: &email,
	}
	users, err := mockClient.GetUsers(query)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, testUser.ID, users[0].ID)
}

func TestAppOperations(t *testing.T) {
	// Create mock client
	mockClient := new(mocks.MockClient)

	// Setup test data
	id := int32(1)
	name := "Test App"
	testApp := models.App{
		ID:   &id,
		Name: &name,
	}

	// Setup mock expectations
	mockClient.On("GetApps", mock.AnythingOfType("*models.AppQuery")).Return([]models.App{testApp}, nil)

	// Test app query
	query := &models.AppQuery{}
	apps, err := mockClient.GetApps(query)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, apps, 1)
	assert.Equal(t, *testApp.ID, *apps[0].ID)
}
