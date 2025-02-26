package main

import (
	"fmt"
	"log"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	// Set environment variables for OneLogin API credentials
	os.Setenv("ONELOGIN_CLIENT_ID", "client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "client_secret")
	os.Setenv("ONELOGIN_SUBDOMAIN", "your-api-subdomain")

	// Initialize the client
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create and get users
	email := "jane@example.com"
	newUser := models.User{
		Email:     email,
		Firstname: "Jane",
		Lastname:  "Doe",
	}

	result, err := client.CreateUser(newUser)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	createdUser := result.(models.User)
	fmt.Printf("Created user with ID: %d\n", createdUser.ID)

	// Get users
	userQuery := &models.UserQuery{
		Email: &email,
	}
	result, err = client.GetUsers(userQuery)
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}
	users := result.([]models.User)
	fmt.Printf("Found %d users\n", len(users))

	// Get apps
	appQuery := &models.AppQuery{}
	result, err = client.GetApps(appQuery)
	if err != nil {
		log.Fatalf("Failed to get apps: %v", err)
	}
	apps := result.([]models.App)
	fmt.Printf("Found %d apps\n", len(apps))

	// Print first user details if any exist
	if len(users) > 0 {
		user := users[0]
		fmt.Printf("\nFirst user details:\n")
		fmt.Printf("ID: %d\n", user.ID)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Username: %s\n", user.Username)
	}
}
