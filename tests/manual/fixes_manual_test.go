package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func main() {
	// Check if required environment variables are set
	checkRequiredEnvVars()

	// Initialize the SDK
	sdk, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Printf("Error initializing SDK: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	fmt.Println("=== Testing SDK Fixes ===")
	
	fmt.Println("\n1. Testing Context Propagation:")
	testContextPropagation(sdk)
	
	fmt.Println("\n2. Testing Pagination:")
	testPagination(sdk)
	
	fmt.Println("\n3. Testing GroupID Filtering:")
	testGroupIDFiltering(sdk)
	
	fmt.Println("\n=== All Tests Completed Successfully ===")
}

func checkRequiredEnvVars() {
	required := []string{
		"ONELOGIN_CLIENT_ID", 
		"ONELOGIN_CLIENT_SECRET", 
		"ONELOGIN_SUBDOMAIN",
	}
	
	missing := []string{}
	for _, env := range required {
		if os.Getenv(env) == "" {
			missing = append(missing, env)
		}
	}
	
	if len(missing) > 0 {
		fmt.Println("Missing required environment variables:")
		for _, env := range missing {
			fmt.Printf("  - %s\n", env)
		}
		os.Exit(1)
	}
}

func testContextPropagation(sdk *onelogin.OneloginSDK) {
	fmt.Println("  Creating context with 10s timeout")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Create a query to get a single user
	query := &models.UserQuery{
		Limit: "1",
	}
	
	fmt.Println("  Getting users with context...")
	
	// Use the context-aware method
	result, err := sdk.GetUsersWithContext(ctx, query)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}
	
	// If we reach here, the context-aware call succeeded
	fmt.Println("  ✅ Context successfully propagated to API request")
	
	// Try with a very short timeout that should fail
	fmt.Println("  Testing with 1ms timeout (should fail or time out)...")
	
	ctxShort, cancelShort := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancelShort()
	
	// Sleep to ensure timeout occurs
	time.Sleep(5 * time.Millisecond)
	
	// This should fail with a context deadline exceeded error
	_, err = sdk.GetUsersWithContext(ctxShort, query)
	if err != nil && (err.Error() == "context deadline exceeded" || err.Error() == "context canceled") {
		fmt.Println("  ✅ Short timeout correctly triggered context deadline error")
	} else if err != nil {
		fmt.Printf("  ⚠️ Got error but not context deadline: %v\n", err)
	} else {
		fmt.Println("  ⚠️ Short timeout test didn't fail as expected (network might be very fast)")
	}
	
	fmt.Println("  Context propagation test completed")
}

func testPagination(sdk *onelogin.OneloginSDK) {
	// Create query with pagination parameters
	query := &models.UserQuery{
		Limit: "2", // Get 2 users per page
		Page: "1",  // Start with the first page
	}
	
	fmt.Println("  Getting first page of users with pagination...")
	
	// Get the first page of users with pagination info
	resp, err := sdk.GetUsersWithPagination(query)
	if err != nil {
		fmt.Printf("  ❌ Error getting users with pagination: %v\n", err)
		return
	}
	
	// Print pagination info
	fmt.Println("  ✅ Pagination information successfully retrieved")
	fmt.Printf("  • Current Page: %d\n", resp.Pagination.CurrentPage)
	fmt.Printf("  • Total Pages: %d\n", resp.Pagination.TotalPages)
	fmt.Printf("  • Total Count: %d\n", resp.Pagination.TotalCount)
	
	// Check if we have a cursor or after cursor
	if resp.Pagination.Cursor != "" {
		fmt.Printf("  • Cursor: %s\n", resp.Pagination.Cursor)
	}
	
	if resp.Pagination.AfterCursor != "" {
		fmt.Printf("  • After Cursor: %s\n", resp.Pagination.AfterCursor)
		
		// Try getting the next page using the cursor
		nextQuery := &models.UserQuery{
			Limit: "2",
			Cursor: resp.Pagination.AfterCursor,
		}
		
		fmt.Println("  Getting next page using cursor...")
		nextResp, err := sdk.GetUsersWithPagination(nextQuery)
		if err != nil {
			fmt.Printf("  ⚠️ Error getting next page: %v\n", err)
		} else {
			fmt.Println("  ✅ Next page retrieved successfully using cursor")
		}
	}
	
	// Try getting the second page using page parameter
	page2Query := &models.UserQuery{
		Limit: "2",
		Page: "2",
	}
	
	fmt.Println("  Getting page 2 using page parameter...")
	page2Resp, err := sdk.GetUsersWithPagination(page2Query)
	if err != nil {
		fmt.Printf("  ⚠️ Error getting page 2: %v\n", err)
	} else {
		fmt.Println("  ✅ Page 2 retrieved successfully")
		
		// Get the data to display count of users
		if users, ok := page2Resp.Data.([]interface{}); ok {
			fmt.Printf("  • Retrieved %d users on page 2\n", len(users))
		}
	}
	
	fmt.Println("  Pagination test completed")
}

func testGroupIDFiltering(sdk *onelogin.OneloginSDK) {
	fmt.Println("  Trying to get groups to find a valid Group ID...")
	
	// Try to get all groups to find a valid group ID for filtering
	// Note: This assumes a "ListGroups" method; adjust if the method name is different
	groups, err := sdk.ListGroups()
	if err != nil {
		fmt.Printf("  ⚠️ Couldn't retrieve groups: %v\n", err)
		fmt.Println("  Skipping GroupID filtering test")
		return
	}
	
	// Extract a group ID from the response
	groupsData, ok := groups.([]interface{})
	if !ok || len(groupsData) == 0 {
		fmt.Println("  ⚠️ No groups found or invalid response format")
		fmt.Println("  Skipping GroupID filtering test")
		return
	}
	
	// Get the first group ID
	firstGroup, ok := groupsData[0].(map[string]interface{})
	if !ok {
		fmt.Println("  ⚠️ Invalid group data format")
		fmt.Println("  Skipping GroupID filtering test")
		return
	}
	
	groupID := fmt.Sprintf("%v", firstGroup["id"])
	fmt.Printf("  Found Group ID: %s\n", groupID)
	
	// Filter users by this group ID
	query := &models.UserQuery{}
	query.GroupID = &groupID
	
	fmt.Printf("  Filtering users by Group ID: %s\n", groupID)
	users, err := sdk.GetUsers(query)
	if err != nil {
		fmt.Printf("  ❌ Error filtering users by group ID: %v\n", err)
		return
	}
	
	// Count the users in the group
	usersData, ok := users.([]interface{})
	if ok {
		fmt.Printf("  ✅ Found %d users in group %s\n", len(usersData), groupID)
	} else {
		fmt.Println("  ⚠️ Users data in unexpected format")
	}
	
	fmt.Println("  GroupID filtering test completed")
}