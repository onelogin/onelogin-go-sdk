package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/internal/models"
	"github.com/onelogin/onelogin-go-sdk/pkg/onelogin"
)

func main() {
	ol, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println("Unable to initialize client:", err)
		return
	}

	userQuery := models.UserQuery{}
	userList, err := ol.GetUsers(&userQuery)
	if err != nil {
		fmt.Println("Failed to get user:", err)
		return
	}
	fmt.Println(userList)

}
