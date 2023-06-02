package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/onelogin"
)

func main() {

	// Create a new instance of the Onelogin SDK
	ol := onelogin.NewOneloginSDK()

	// Use the Onelogin SDK to make API calls

	token, err := ol.GetToken()
	if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
		return
	}

	fmt.Printf("Testing to see token: %s\n", token)
}
