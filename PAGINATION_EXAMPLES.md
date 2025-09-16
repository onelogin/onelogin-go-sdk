# App Users Pagination Examples

This document demonstrates how to use the new pagination features for the `GetAppUsers` method.

## Overview

The OneLogin Go SDK now supports pagination for retrieving app users, allowing you to handle applications with more than 100 users. The implementation maintains backward compatibility:

1. `GetAppUsers(appID int)` - **Original method** (fully backward compatible) 
2. `GetAppUsersWithPagination(appID int, queryParams *AppUserQuery)` - **New method** with pagination parameters
3. `GetAppUsersWithPaginationAndContext(ctx context.Context, appID int, queryParams *AppUserQuery)` - **New method** with context support

## Basic Usage

### 1. Backward Compatible (Original Method - No Changes Required)

```go
// This works exactly as before - returns up to 100 users
// NO CODE CHANGES NEEDED for existing applications
users, err := sdk.GetAppUsers(123)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}

// users is still interface{} - same as before
fmt.Printf("Retrieved users: %+v\n", users)
```

### 2. With Pagination Parameters (New Method)

```go
// Create query with pagination parameters
query := &models.AppUserQuery{
    Limit: "50",  // Get 50 users per page
    Page:  "2",   // Get page 2
}

result, err := sdk.GetAppUsersWithPagination(123, query)
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

### 3. With Context Support (New Method)

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

query := &models.AppUserQuery{
    Limit: "50",
    Page:  "1",
}

result, err := sdk.GetAppUsersWithPaginationAndContext(ctx, 123, query)
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

result, err := sdk.GetAppUsersWithPagination(123, query)
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
    
    result, err := sdk.GetAppUsersWithPaginationAndContext(ctx, 123, query)
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

The new pagination methods return a `*PagedResponse` with:

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

result, err := sdk.GetAppUsersWithPagination(123, query)
if err != nil {
    // Handle validation error with context
    fmt.Printf("Parameter validation failed: %v\n", err)
}
```

## Migration Guide

### ✅ **No Breaking Changes - Existing Code Works Unchanged**

**Existing code (no changes needed):**
```go
users, err := sdk.GetAppUsers(appID)
// This continues to work exactly as before
```

**To add pagination (new functionality):**
```go
// Option 1: Add pagination support
query := &models.AppUserQuery{Limit: "50", Page: "1"}
result, err := sdk.GetAppUsersWithPagination(appID, query)
users := result.Data

// Option 2: Add context support
ctx := context.Background()
result, err := sdk.GetAppUsersWithPaginationAndContext(ctx, appID, query)
users := result.Data
```

## Backward Compatibility Guarantee

- ✅ **Original `GetAppUsers(appID int)` method unchanged**
- ✅ **Same return type: `(interface{}, error)`**
- ✅ **Same behavior: returns up to 100 users**
- ✅ **Existing code requires no modifications**
- ✅ **New functionality available through new methods**

This approach ensures existing applications continue to work while providing new pagination capabilities through additional methods.