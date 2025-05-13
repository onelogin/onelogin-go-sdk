package main

import (
	"fmt"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
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

	// Simple test - get connectors
	fmt.Println("Testing connection to OneLogin API...")
	connectors, err := sdk.ListConnectors()
	if err != nil {
		fmt.Printf("Error listing connectors: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success! Connected to OneLogin API.")
	fmt.Printf("Connectors data type: %T\n", connectors)

	// Test the token generation
	token, err := sdk.GetToken()
	if err != nil {
		fmt.Printf("Error getting token: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully retrieved token!")
	fmt.Printf("Token length: %d\n", len(token))
}
