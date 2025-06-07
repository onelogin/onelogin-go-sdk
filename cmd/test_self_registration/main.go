package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Create a new OneLogin SDK client
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		log.Fatalf("Error creating OneLogin SDK client: %v", err)
	}

	// Get all self-registration profiles
	query := &models.SelfRegistrationProfileQuery{
		Limit: "10",
		Page:  "1",
	}
	
	fmt.Println("Fetching self-registration profiles...")
	profiles, err := client.GetSelfRegistrationProfiles(query)
	if err != nil {
		log.Fatalf("Error getting self-registration profiles: %v", err)
	}

	// Pretty print the profiles
	profilesJSON, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling profiles to JSON: %v", err)
	}
	fmt.Println("Self-registration profiles:")
	fmt.Println(string(profilesJSON))

	// Create a new self-registration profile
	fmt.Println("\nCreating a new self-registration profile...")
	newProfile := models.SelfRegistrationProfile{
		Name:                 "Test Profile via SDK",
		URL:                  "test-profile-sdk-" + os.Getenv("USER"),
		Enabled:              true,
		Moderated:            false,
		DefaultRoleID:        0, // Use a valid role ID from your account
		Helptext:             "Welcome to our community. Please follow the guidelines.",
		ThankyouMessage:      "Thank you for joining!",
		DomainBlacklist:      "spam.com, banned.com",
		DomainWhitelist:      "example.com, trusted.com",
		DomainListStrategy:   models.DomainWhitelistStrategy,
		EmailVerificationType: models.EmailMagicLink,
	}

	result, err := client.CreateSelfRegistrationProfile(newProfile)
	if err != nil {
		log.Fatalf("Error creating self-registration profile: %v", err)
	}

	// Extract the ID from the result
	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling result to JSON: %v", err)
	}
	fmt.Println("Created self-registration profile:")
	fmt.Println(string(resultJSON))

	// Parse the result to get the ID
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		log.Fatalf("Error parsing result: unexpected type")
	}

	profileMap, ok := resultMap["self_registration_profile"].(map[string]interface{})
	if !ok {
		log.Fatalf("Error parsing result: self_registration_profile not found or not a map")
	}

	profileID := int(profileMap["id"].(float64))
	fmt.Printf("\nCreated profile ID: %d\n", profileID)

	// Get the profile by ID
	fmt.Printf("\nFetching profile with ID %d...\n", profileID)
	profile, err := client.GetSelfRegistrationProfile(profileID)
	if err != nil {
		log.Fatalf("Error getting self-registration profile: %v", err)
	}

	// Pretty print the profile
	profileJSON, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling profile to JSON: %v", err)
	}
	fmt.Println("Self-registration profile:")
	fmt.Println(string(profileJSON))

	// Update the profile
	fmt.Printf("\nUpdating profile with ID %d...\n", profileID)
	updatedProfile := models.SelfRegistrationProfile{
		Name:                 "Updated Test Profile via SDK",
		URL:                  "test-profile-sdk-" + os.Getenv("USER"),
		Enabled:              false,
		Moderated:            true,
		Helptext:             "Updated welcome message.",
		ThankyouMessage:      "Updated thank you message!",
		EmailVerificationType: models.EmailOTP,
	}

	updateResult, err := client.UpdateSelfRegistrationProfile(profileID, updatedProfile)
	if err != nil {
		log.Fatalf("Error updating self-registration profile: %v", err)
	}

	// Pretty print the update result
	updateResultJSON, err := json.MarshalIndent(updateResult, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling update result to JSON: %v", err)
	}
	fmt.Println("Updated self-registration profile:")
	fmt.Println(string(updateResultJSON))

	// Delete the profile
	fmt.Printf("\nDeleting profile with ID %d...\n", profileID)
	deleteResult, err := client.DeleteSelfRegistrationProfile(profileID)
	if err != nil {
		log.Fatalf("Error deleting self-registration profile: %v", err)
	}

	fmt.Println("Profile deleted successfully!")
	fmt.Printf("Delete result: %v\n", deleteResult)
}
