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

	appQuery := models.AppQuery{}
	appResponse, err := ol.GetApps(&appQuery)
	if err != nil {
		fmt.Println("Failed to get app list:", err)
		return
	}
	appList, ok := appResponse.([]string) // or whatever the correct type should be
	if !ok {
		fmt.Println("Failed to assert appResponse as type []string")
		return
	}
	fmt.Println(appList)

}
