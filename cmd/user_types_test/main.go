package main

import (
	"fmt"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	// Test the SDK changes related to user properties and custom attributes

	// Get credentials from environment variables
	clientID := os.Getenv("ONELOGIN_CLIENT_ID")
	clientSecret := os.Getenv("ONELOGIN_CLIENT_SECRET")
	subdomain := os.Getenv("ONELOGIN_SUBDOMAIN")

	if clientID == "" || clientSecret == "" || subdomain == "" {
		fmt.Println("Please set ONELOGIN_CLIENT_ID, ONELOGIN_CLIENT_SECRET, and ONELOGIN_SUBDOMAIN environment variables")
		os.Exit(1)
	}

	// Set environment variables for SDK initialization
	os.Setenv("ONELOGIN_CLIENT_ID", clientID)
	os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	os.Setenv("ONELOGIN_SUBDOMAIN", subdomain)

	// Initialize the SDK
	sdk, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Printf("Error initializing SDK: %v\n", err)
		os.Exit(1)
	}

	// Test 1: Create a user with the fixed properties
	fmt.Println("Test 1: Create a user with the fixed properties")
	user := models.User{
		Firstname:     "Test",
		Lastname:      "User",
		Email:         "testuser@example.com",
		Username:      "testuser@example.com",
		MemberOf:      []string{"Group1", "Group2"},
		ManagerADID:   12345,
		ExternalID:    "ext-id-123",
		RoleIDs:       []int32{1, 2, 3},
		CustomAttributes: map[string]interface{}{
			"department": "Engineering",
			"location":   "Remote",
		},
	}

	fmt.Printf("User to create: %+v\n", user)

	userResp, err := sdk.CreateUser(user)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		// Continue with other tests even if this fails
	} else {
		fmt.Printf("User created successfully: %+v\n", userResp)
		
		// Get the user ID from the response
		// Note: This is a simplified approach, in a real test you'd parse the response properly
		// userID := ... get from response
	}

	// Test 2: Get custom attributes
	fmt.Println("\nTest 2: Get custom attributes")
	customAttrs, err := sdk.GetCustomAttributes()
	if err != nil {
		fmt.Printf("Error getting custom attributes: %v\n", err)
	} else {
		fmt.Printf("Custom attributes: %+v\n", customAttrs)
	}

	// Test 3: Get users with query parameters
	fmt.Println("\nTest 3: Get users with query parameters")
	query := models.UserQuery{
		Limit: "10",
	}

	users, err := sdk.GetUsersModels(&query)
	if err != nil {
		fmt.Printf("Error getting users: %v\n", err)
	} else {
		fmt.Printf("Found %d users\n", len(users))
		// Print details of the first user to verify field types
		if len(users) > 0 {
			firstUser := users[0]
			fmt.Printf("First user details:\n")
			fmt.Printf("  ID: %d\n", firstUser.ID)
			fmt.Printf("  Name: %s %s\n", firstUser.Firstname, firstUser.Lastname)
			fmt.Printf("  Email: %s\n", firstUser.Email)
			fmt.Printf("  ExternalID type: %T, value: %v\n", firstUser.ExternalID, firstUser.ExternalID)
			fmt.Printf("  ManagerADID type: %T, value: %v\n", firstUser.ManagerADID, firstUser.ManagerADID)
			fmt.Printf("  ManagerUserID type: %T, value: %v\n", firstUser.ManagerUserID, firstUser.ManagerUserID)
			fmt.Printf("  MemberOf type: %T, value: %v\n", firstUser.MemberOf, firstUser.MemberOf)
			fmt.Printf("  RoleIDs type: %T, value: %v\n", firstUser.RoleIDs, firstUser.RoleIDs)
			fmt.Printf("  CustomAttributes type: %T, value: %v\n", firstUser.CustomAttributes, firstUser.CustomAttributes)
		}
	}

	fmt.Println("\nTests completed")
}