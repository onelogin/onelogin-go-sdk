# App Users Pagination Examples

This document demonstrates how to use the new pagination features for the `GetAppUsers` method.

## Overview

The OneLogin Go SDK now supports pagination for retrieving app users, allowing you to handle applications with more than 100 users. Multiple methods are available:

1. `GetAppUsers(appID int)` - Original method (backward compatible)
2. `GetAppUsersWithQuery(appID int, queryParams *AppUserQuery)` - With pagination parameters
3. `GetAppUsersWithContext(ctx context.Context, appID int, queryParams *AppUserQuery)` - With context support
4. `GetAppUsersWithPagination(appID int, queryParams *AppUserQuery)` - With full pagination information
5. `GetAppUsersWithPaginationWithContext(ctx context.Context, appID int, queryParams *AppUserQuery)` - With context and full pagination

## Basic Usage

### 1. Backward Compatible (Original Method)

```go
// This works exactly as before - returns up to 100 users
users, err := sdk.GetAppUsers(123)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}
```

### 2. With Pagination Parameters

```go
// Create query with pagination parameters
query := &models.AppUserQuery{
    Limit: "50",  // Get 50 users per page
    Page:  "2",   // Get page 2
}

users, err := sdk.GetAppUsersWithQuery(123, query)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}
```

### 3. With Context Support

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

query := &models.AppUserQuery{
    Limit: "50",
    Page:  "1",
}

users, err := sdk.GetAppUsersWithContext(ctx, 123, query)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}
```

### 4. With Full Pagination Information

```go
query := &models.AppUserQuery{
    Limit: "50",
    Page:  "1",
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

## Advanced Usage

### Context with Timeout and Pagination

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

query := &models.AppUserQuery{
    Limit: "50",
    Page:  "1",
}

result, err := sdk.GetAppUsersWithPaginationWithContext(ctx, 123, query)
if err != nil {
    log.Fatalf("Failed to get app users: %v", err)
}
```

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
    
    result, err := sdk.GetAppUsersWithPaginationWithContext(ctx, 123, query)
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

## Error Handling

The methods include improved error handling with context:

```go
query := &models.AppUserQuery{
    Limit: "invalid_value",
}

// This will return a descriptive error due to invalid parameters
result, err := sdk.GetAppUsersWithPagination(123, query)
if err != nil {
    // Handle validation error with context
    fmt.Printf("Parameter validation failed: %v\n", err)
}
```

## Migration Guide

If you're currently using `GetAppUsers()`:

**Before:**
```go
users, err := sdk.GetAppUsers(appID)
```

**After (with pagination):**
```go
// Option 1: Keep existing behavior (no changes needed)
users, err := sdk.GetAppUsers(appID)

// Option 2: Add pagination support
query := &models.AppUserQuery{Limit: "50", Page: "1"}
users, err := sdk.GetAppUsersWithQuery(appID, query)

// Option 3: Add context support
ctx := context.Background()
users, err := sdk.GetAppUsersWithContext(ctx, appID, query)

// Option 4: Get pagination metadata
result, err := sdk.GetAppUsersWithPagination(appID, query)
users := result.Data
```

The original `GetAppUsers()` method remains unchanged for backward compatibility.