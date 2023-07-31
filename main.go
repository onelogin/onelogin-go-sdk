package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	UserTwo := models.User{Firstname: "Mikhail", Lastname: "Beaverton", Email: "mb@example.com"}
	UserQueryOne := models.UserQuery{Email: &UserTwo.Email}

	Client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	Client.CreateUser(models.User{Firstname: "Jane", Lastname: "Pukalava", Email: "jp@example.com"})
	Client.CreateUser(UserTwo)
	Client.GetUsers(&UserQueryOne)
}
