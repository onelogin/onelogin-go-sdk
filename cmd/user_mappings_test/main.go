package main

import (
	"fmt"
	"log"
	"os"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	// Initialize the SDK
	sdk, err := onelogin.NewOneloginSDK()
	if err != nil {
		log.Fatalf("Failed to create OneLogin client: %v", err)
	}

	// Get authentication token
	log.Println("Getting authentication token...")
	_, err = sdk.GetToken()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	log.Println("Authentication token retrieved successfully.")

	fmt.Println("=== User Mappings API Example ===")

	// Skip listing and just create a mapping
	fmt.Println("\n-- Creating a User Mapping --")
	name := "Test Mapping via SDK"
	match := "all"
	enabled := true
	position := int32(1)

	// Set up conditions
	sourceCondition := "email"
	equalsOperator := "="
	valueCondition := "test@example.com"

	// Set up actions
	actionName := "set_status"
	actionValue := []string{"1"}

	newMapping := models.UserMapping{
		Name:     &name,
		Match:    &match,
		Enabled:  &enabled,
		Position: &position,
		Conditions: []models.UserMappingConditions{
			{
				Source:   &sourceCondition,
				Operator: &equalsOperator,
				Value:    &valueCondition,
			},
		},
		Actions: []models.UserMappingActions{
			{
				Action: &actionName,
				Value:  actionValue,
			},
		},
	}

	// Set to true to keep the mapping for demonstration
	leaveMapping := true

	createdMapping, err := sdk.CreateUserMapping(newMapping)
	if err != nil {
		log.Printf("Failed to create user mapping: %v", err)
	} else {
		fmt.Printf("Created new mapping with ID: %d\n", *createdMapping.ID)

		// Get the created mapping
		fmt.Println("\n-- Getting a User Mapping --")
		retrievedMapping, err := sdk.GetUserMapping(*createdMapping.ID)
		if err != nil {
			log.Printf("Failed to get user mapping: %v", err)
		} else {
			fmt.Printf("Retrieved mapping: ID: %d, Name: %s\n", *retrievedMapping.ID, *retrievedMapping.Name)
		}

		// The update might fail due to API constraints, but we want to clean up regardless
		fmt.Println("\n-- Updating a User Mapping --")
		updatedName := "Updated Test Mapping"
		retrievedMapping.Name = &updatedName

		var mappingIDToDelete int32
		mappingIDToDelete = *retrievedMapping.ID

		updatedMapping, err := sdk.UpdateUserMapping(*retrievedMapping.ID, *retrievedMapping)
		if err != nil {
			log.Printf("Failed to update user mapping: %v", err)
			// Check if the mapping was actually updated despite the error
			checkMapping, checkErr := sdk.GetUserMapping(*retrievedMapping.ID)
			if checkErr == nil && *checkMapping.Name == "Updated Test Mapping" {
				log.Printf("Note: The mapping was actually updated successfully despite the error")
				fmt.Printf("Updated mapping: ID: %d, Name: %s\n", *checkMapping.ID, *checkMapping.Name)
				mappingIDToDelete = *checkMapping.ID
			}
		} else {
			fmt.Printf("Updated mapping: ID: %d, Name: %s\n", *updatedMapping.ID, *updatedMapping.Name)
			mappingIDToDelete = *updatedMapping.ID
		}

		// Only delete if not keeping the mapping
		if !leaveMapping {
			// Clean up - delete the mapping
			fmt.Println("\n-- Deleting a User Mapping --")
			err = sdk.DeleteUserMapping(mappingIDToDelete)
			if err != nil {
				log.Printf("Failed to delete user mapping: %v", err)
			} else {
				fmt.Printf("Deleted mapping with ID: %d\n", mappingIDToDelete)
			}
		} else {
			fmt.Println("\n-- Skipping deletion (leaving mapping for demonstration) --")
			fmt.Printf("Mapping with ID: %d was created successfully and left in place\n", mappingIDToDelete)
		}
	}

	fmt.Println("\n=== Example Complete ===")
	os.Exit(0)
}
