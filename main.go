package main

import (
	"fmt"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {

	// Set environment variables for OneLogin API credentials
	os.Setenv("ONELOGIN_CLIENT_ID", "client_id")
	os.Setenv("ONELOGIN_CLIENT_SECRET", "client_secret")
	os.Setenv("ONELOGIN_SUBDOMAIN", "your-api-subdomain")

	// Create a user object with the specified attributes
	UserTwo := models.User{Firstname: "Mikhail", Lastname: "Beaverton", Email: "mb@example.com"}

	// Create a user query object with the specified email
	UserQueryOne := models.UserQuery{Email: &UserTwo.Email}

	// Create a new OneLogin SDK client
	Client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	// Create a new user with the specified attributes
	Client.CreateUser(models.User{Firstname: "Jane", Lastname: "Pukalava", Email: "jp@example.com"})

	// Create a new user with the specified attributes
	Client.CreateUser(UserTwo)

	// Get a list of users that match the specified query
	Client.GetUsers(&UserQueryOne)
}
