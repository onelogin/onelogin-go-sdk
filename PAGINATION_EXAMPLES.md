# App Users Pagination Examples

This document demonstrates how to use the new pagination features for the `GetAppUsers` method.

## Overview

The OneLogin Go SDK now supports pagination for retrieving app users, allowing you to handle applications with more than 100 users. Two simple methods are available:

1. `GetAppUsers(appID int, queryParams *AppUserQuery)` - Main method with optional pagination parameters
2. `GetAppUsersWithContext(ctx context.Context, appID int, queryParams *AppUserQuery)` - With context support

Both methods return a `*PagedResponse` that includes both the user data and pagination metadata.

## Basic Usage

### 1. Simple Usage (No Pagination)

```go
// Get users without pagination (pass nil for queryParams)
result, err := sdk.GetAppUsers(123, nil)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}

// Access the users data
users := result.Data
fmt.Printf("Retrieved %d users\n", len(users.([]interface{})))
```

### 2. With Pagination Parameters

```go
// Create query with pagination parameters
query := &models.AppUserQuery{
    Limit: "50",  // Get 50 users per page
    Page:  "2",   // Get page 2
}

result, err := sdk.GetAppUsers(123, query)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}

// Access the users data
users := result.Data

// Access pagination metadata
fmt.Printf("Current page: %d\n", result.Pagination.CurrentPage)
fmt.Printf("Total pages: %d\n", result.Pagination.TotalPages)
fmt.Printf("Total count: %d\n", result.Pagination.TotalCount)
fmt.Printf("Next page cursor: %s\n", result.Pagination.AfterCursor)
```

### 3. With Context Support

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

query := &models.AppUserQuery{
    Limit: "50",
    Page:  "1",
}

result, err := sdk.GetAppUsersWithContext(ctx, 123, query)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}

// Same result structure as above
users := result.Data
pagination := result.Pagination
```

## Advanced Usage

### Cursor-Based Pagination

For better performance with large datasets, use cursor-based pagination:

```go
query := &models.AppUserQuery{
    Limit:  "50",
    Cursor: "your_cursor_here",
}

result, err := sdk.GetAppUsers(123, query)
```

### Iterating Through All Pages

```go
var allUsers []interface{}
cursor := ""
ctx := context.Background()

for {
    query := &models.AppUserQuery{
        Limit:  "50",
        Cursor: cursor,
    }
    
    result, err := sdk.GetAppUsersWithContext(ctx, 123, query)
    if err != nil {
        log.Fatalf("Failed to get page: %v", err)
    }
    
    // Add users from this page
    if users, ok := result.Data.([]interface{}); ok {
        allUsers = append(allUsers, users...)
    }
    
    // Check if there's a next page
    if result.Pagination.AfterCursor == "" {
        break // No more pages
    }
    
    cursor = result.Pagination.AfterCursor
}

fmt.Printf("Retrieved %d total users\n", len(allUsers))
```

## AppUserQuery Parameters

The `AppUserQuery` struct supports the following parameters:

- `Limit` (string): Number of results per page (e.g., "50", "100")
- `Page` (string): Page number to retrieve (e.g., "1", "2", "3")
- `Cursor` (string): Cursor for cursor-based pagination

## PagedResponse Structure

All methods return a `*PagedResponse` with:

```go
type PagedResponse struct {
    Data       interface{}    `json:"data"`        // The actual user data
    Pagination PaginationInfo `json:"pagination"`  // Pagination metadata
}

type PaginationInfo struct {
    Cursor       string `json:"cursor,omitempty"`
    AfterCursor  string `json:"after_cursor,omitempty"`
    BeforeCursor string `json:"before_cursor,omitempty"`
    TotalPages   int    `json:"total_pages,omitempty"`
    CurrentPage  int    `json:"current_page,omitempty"`
    TotalCount   int    `json:"total_count,omitempty"`
}
```

## Error Handling

The methods include improved error handling with context:

```go
query := &models.AppUserQuery{
    Limit: "invalid_value",
}

result, err := sdk.GetAppUsers(123, query)
if err != nil {
    // Handle validation error with context
    fmt.Printf("Parameter validation failed: %v\n", err)
}
```

## Migration Guide

If you're using an older version:

**Before (5 different methods):**
```go
// Multiple confusing methods
users1, err := sdk.GetAppUsers(appID)
users2, err := sdk.GetAppUsersWithQuery(appID, query)
users3, err := sdk.GetAppUsersWithContext(ctx, appID, query)
result4, err := sdk.GetAppUsersWithPagination(appID, query)
result5, err := sdk.GetAppUsersWithPaginationWithContext(ctx, appID, query)
```

**After (2 simple methods):**
```go
// Simple, unified API
result, err := sdk.GetAppUsers(appID, query)           // nil query for no pagination
result, err := sdk.GetAppUsersWithContext(ctx, appID, query)  // with context support

// Always get both data and pagination info
users := result.Data
pagination := result.Pagination
```

## Breaking Changes

⚠️ **This is a breaking change from previous versions:**

- The method signature has changed to return `*PagedResponse` instead of `interface{}`
- Query parameters are now optional (pass `nil` for no pagination)
- All pagination information is always included in the response