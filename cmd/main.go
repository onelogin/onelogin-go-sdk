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
	appList, err := ol.GetApps(&appQuery)
	if err != nil {
		fmt.Println("Failed to get app list:", err)
		return
	}
	fmt.Println("App List:", appList)
	description := "Test App Description"
	newApp := models.App{Name: "Go SDK Demo App", Description: &description, ConnectorID: 2}
	createdApp, err := ol.CreateApp(&newApp)
	if err != nil {
		fmt.Println("Failed to create app:", err)
		return
	}
	fmt.Println("Created App:", createdApp)

}
