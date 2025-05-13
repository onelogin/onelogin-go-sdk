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

	// Example 1: Create a custom attribute using the new helper function
	fmt.Println("Example 1: Creating a custom attribute")

	// Create with the new helper that properly structures the payload
	createResp, err := sdk.CreateCustomAttribute("Employee ID", "employee_id")
	if err != nil {
		fmt.Printf("Error creating custom attribute: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Custom attribute created successfully: %+v\n", createResp)

	// Extract the ID from the response for later use
	attr := createResp.(map[string]interface{})
	attrID := int(attr["id"].(float64))
	fmt.Printf("Created attribute ID: %d\n", attrID)

	// Example 2: Get all custom attributes
	fmt.Println("\nExample 2: Get custom attributes")
	customAttrs, err := sdk.GetCustomAttributes()
	if err != nil {
		fmt.Printf("Error getting custom attributes: %v\n", err)
	} else {
		fmt.Printf("Custom attributes: %+v\n", customAttrs)
	}

	// Example 3: Get custom attribute by ID
	fmt.Println("\nExample 3: Get custom attribute by ID")
	attrByID, err := sdk.GetCustomAttributeByID(attrID)
	if err != nil {
		fmt.Printf("Error getting custom attribute by ID: %v\n", err)
	} else {
		fmt.Printf("Custom attribute by ID: %+v\n", attrByID)
	}

	// Example 4: Update the custom attribute
	fmt.Println("\nExample 4: Update custom attribute")
	updateResp, err := sdk.UpdateCustomAttribute(attrID, "Employee Number", "employee_id")
	if err != nil {
		fmt.Printf("Error updating custom attribute: %v\n", err)
	} else {
		fmt.Printf("Custom attribute updated successfully: %+v\n", updateResp)
	}

	// Verify the update
	updatedAttr, err := sdk.GetCustomAttributeByID(attrID)
	if err != nil {
		fmt.Printf("Error getting updated custom attribute: %v\n", err)
	} else {
		fmt.Printf("Updated custom attribute: %+v\n", updatedAttr)
	}

	// Example 5: Delete the custom attribute
	fmt.Println("\nExample 5: Delete custom attribute")
	deleteResp, err := sdk.DeleteCustomAttributes(attrID)
	if err != nil {
		fmt.Printf("Error deleting custom attribute: %v\n", err)
	} else {
		fmt.Printf("Custom attribute deleted successfully: %+v\n", deleteResp)
	}

	// Verify the deletion
	fmt.Println("\nVerifying deletion - get all custom attributes")
	remainingAttrs, err := sdk.GetCustomAttributes()
	if err != nil {
		fmt.Printf("Error getting remaining custom attributes: %v\n", err)
	} else {
		fmt.Printf("Remaining custom attributes: %+v\n", remainingAttrs)

		// Check if our deleted attribute is still in the list
		deleted := true
		attrs := remainingAttrs.([]interface{})
		for _, a := range attrs {
			attr, ok := a.(map[string]interface{})
			if !ok {
				continue
			}

			if id, exists := attr["id"]; exists && int(id.(float64)) == attrID {
				deleted = false
				break
			}
		}

		if deleted {
			fmt.Println("✅ Attribute was successfully deleted!")
		} else {
			fmt.Println("❌ Attribute was NOT successfully deleted!")
		}
	}

	fmt.Println("\nExamples completed")
}
