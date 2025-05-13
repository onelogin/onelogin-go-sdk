package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/api"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// Mock the HTTP client for testing
func setupMockClient() (*api.Client, *MockHttpClient) {
	mockHttpClient := &MockHttpClient{}
	client := createMockClient()
	client.HttpClient = mockHttpClient
	return client, mockHttpClient
}

// Test the User Mappings feature
func TestUserMappings(t *testing.T) {
	client, mockHttpClient := setupMockClient()

	// Create a OneLogin SDK with the mock client
	sdk := &onelogin.OneloginSDK{Client: client}

	// Test listing user mappings
	t.Run("List User Mappings", func(t *testing.T) {
		// Mock response for listing user mappings
		mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
			// Check the request path
			if req.URL.Path != "/api/2/mappings" {
				t.Fatalf("Expected path to be /api/2/mappings, got %s", req.URL.Path)
			}

			// Return a mock response
			mockResponse := `[
				{
					"id": 123,
					"name": "Test Mapping",
					"match": "all",
					"enabled": true,
					"position": 1,
					"conditions": [
						{
							"source": "email",
							"operator": "=",
							"value": "test@example.com"
						}
					],
					"actions": [
						{
							"action": "set_status",
							"value": ["1"]
						}
					]
				}
			]`
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		}

		// Call the function that uses the mocked client
		mappings, err := sdk.ListUserMappings(nil)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the response
		if len(mappings) != 1 {
			t.Fatalf("Expected 1 mapping, got %d", len(mappings))
		}

		if *mappings[0].ID != 123 {
			t.Fatalf("Expected mapping ID to be 123, got %d", *mappings[0].ID)
		}

		if *mappings[0].Name != "Test Mapping" {
			t.Fatalf("Expected mapping name to be 'Test Mapping', got %s", *mappings[0].Name)
		}
	})

	// Test getting a single user mapping
	t.Run("Get User Mapping", func(t *testing.T) {
		mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
			// Check the request path
			if req.URL.Path != "/api/2/mappings/123" {
				t.Fatalf("Expected path to be /api/2/mappings/123, got %s", req.URL.Path)
			}

			// Return a mock response
			mockResponse := `{
				"id": 123,
				"name": "Test Mapping",
				"match": "all",
				"enabled": true,
				"position": 1,
				"conditions": [
					{
						"source": "email",
						"operator": "=",
						"value": "test@example.com"
					}
				],
				"actions": [
					{
						"action": "set_status",
						"value": ["1"]
					}
				]
			}`
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		}

		// Call the function that uses the mocked client
		mapping, err := sdk.GetUserMapping(123)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the response
		if *mapping.ID != 123 {
			t.Fatalf("Expected mapping ID to be 123, got %d", *mapping.ID)
		}

		if *mapping.Name != "Test Mapping" {
			t.Fatalf("Expected mapping name to be 'Test Mapping', got %s", *mapping.Name)
		}
	})

	// Test creating a user mapping
	t.Run("Create User Mapping", func(t *testing.T) {
		mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
			// Check the request path
			if req.URL.Path != "/api/2/mappings" {
				t.Fatalf("Expected path to be /api/2/mappings, got %s", req.URL.Path)
			}

			// Check the request method
			if req.Method != "POST" {
				t.Fatalf("Expected method to be POST, got %s", req.Method)
			}

			// Return a mock response
			mockResponse := `{
				"id": 456,
				"name": "New Mapping",
				"match": "all",
				"enabled": true,
				"position": 1,
				"conditions": [
					{
						"source": "email",
						"operator": "=",
						"value": "new@example.com"
					}
				],
				"actions": [
					{
						"action": "set_status",
						"value": ["1"]
					}
				]
			}`
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		}

		// Create a user mapping
		name := "New Mapping"
		match := "all"
		enabled := true
		position := int32(1)
		sourceCondition := "email"
		equalsOperator := "="
		valueCondition := "new@example.com"
		actionName := "set_status"
		actionValue := []string{"1"}

		newMapping := models.UserMapping{
			Name:     &name,
			Match:    &match,
			Enabled:  &enabled,
			Position: &position,
			Conditions: []models.UserMappingConditions{
				{
					Source:   &sourceCondition,
					Operator: &equalsOperator,
					Value:    &valueCondition,
				},
			},
			Actions: []models.UserMappingActions{
				{
					Action: &actionName,
					Value:  actionValue,
				},
			},
		}

		// Call the function that uses the mocked client
		createdMapping, err := sdk.CreateUserMapping(newMapping)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the response
		if *createdMapping.ID != 456 {
			t.Fatalf("Expected mapping ID to be 456, got %d", *createdMapping.ID)
		}

		if *createdMapping.Name != "New Mapping" {
			t.Fatalf("Expected mapping name to be 'New Mapping', got %s", *createdMapping.Name)
		}
	})

	// Test updating a user mapping with just ID response
	t.Run("Update User Mapping with ID Response", func(t *testing.T) {
		// First request to update returns just the ID
		firstCallDone := false
		mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
			if !firstCallDone {
				// First call - PUT request
				firstCallDone = true
				// Check the request path
				if req.URL.Path != "/api/2/mappings/789" {
					t.Fatalf("Expected path to be /api/2/mappings/789, got %s", req.URL.Path)
				}

				// Check the request method
				if req.Method != "PUT" {
					t.Fatalf("Expected method to be PUT, got %s", req.Method)
				}

				// Return a mock response with just ID
				mockResponse := `{"id": 789}`
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponse)),
				}, nil
			} else {
				// Second call - GET request to fetch the full object
				if req.URL.Path != "/api/2/mappings/789" {
					t.Fatalf("Expected path to be /api/2/mappings/789, got %s", req.URL.Path)
				}

				// Check the request method
				if req.Method != "GET" {
					t.Fatalf("Expected method to be GET, got %s", req.Method)
				}

				// Return a mock response with the full object
				mockResponse := `{
					"id": 789,
					"name": "Updated Mapping",
					"match": "all",
					"enabled": true,
					"position": 1,
					"conditions": [
						{
							"source": "email",
							"operator": "=",
							"value": "updated@example.com"
						}
					],
					"actions": [
						{
							"action": "set_status",
							"value": ["1"]
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponse)),
				}, nil
			}
		}

		// Create a user mapping
		name := "Updated Mapping"
		match := "all"
		enabled := true
		position := int32(1)
		sourceCondition := "email"
		equalsOperator := "="
		valueCondition := "updated@example.com"
		actionName := "set_status"
		actionValue := []string{"1"}

		updateMapping := models.UserMapping{
			Name:     &name,
			Match:    &match,
			Enabled:  &enabled,
			Position: &position,
			Conditions: []models.UserMappingConditions{
				{
					Source:   &sourceCondition,
					Operator: &equalsOperator,
					Value:    &valueCondition,
				},
			},
			Actions: []models.UserMappingActions{
				{
					Action: &actionName,
					Value:  actionValue,
				},
			},
		}

		// Call the function that uses the mocked client
		updatedMapping, err := sdk.UpdateUserMapping(789, updateMapping)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the response
		if *updatedMapping.ID != 789 {
			t.Fatalf("Expected mapping ID to be 789, got %d", *updatedMapping.ID)
		}

		if *updatedMapping.Name != "Updated Mapping" {
			t.Fatalf("Expected mapping name to be 'Updated Mapping', got %s", *updatedMapping.Name)
		}
	})

	// Test deleting a user mapping
	t.Run("Delete User Mapping", func(t *testing.T) {
		mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
			// Check the request path
			if req.URL.Path != "/api/2/mappings/123" {
				t.Fatalf("Expected path to be /api/2/mappings/123, got %s", req.URL.Path)
			}

			// Check the request method
			if req.Method != "DELETE" {
				t.Fatalf("Expected method to be DELETE, got %s", req.Method)
			}

			// Return a mock response
			return &http.Response{
				StatusCode: 204,
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			}, nil
		}

		// Call the function that uses the mocked client
		err := sdk.DeleteUserMapping(123)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})
}