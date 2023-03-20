# \UsersV2Api

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser2**](UsersV2Api.md#CreateUser2) | **Post** /api/2/users | Create User
[**DeleteUser2**](UsersV2Api.md#DeleteUser2) | **Delete** /api/2/users/{user_id} | Delete User
[**GetUser2**](UsersV2Api.md#GetUser2) | **Get** /api/2/users/{user_id} | Get User
[**GetUserApps2**](UsersV2Api.md#GetUserApps2) | **Get** /api/2/users/{user_id}/apps | Get User Apps
[**ListUsers2**](UsersV2Api.md#ListUsers2) | **Get** /api/2/users | List Users
[**UpdateUser2**](UsersV2Api.md#UpdateUser2) | **Put** /api/2/users/{user_id} | Update User



## CreateUser2

> User CreateUser2(ctx).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()

Create User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mappings := "mappings_example" // string | Controls how mappings will be applied to the user on creation. Defaults to async. (optional)
    validatePolicy := true // bool | Will passwords validate against the User Policy? Defaults to true. (optional)
    user := *openapiclient.NewUser() // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.CreateUser2(context.Background()).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.CreateUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser2`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV2Api.CreateUser2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser2

> DeleteUser2(ctx, userId).Execute()

Delete User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := int32(56) // int32 | Set to the id of the user that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.DeleteUser2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.DeleteUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser2

> User GetUser2(ctx, userId).Execute()

Get User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := int32(56) // int32 | Set to the id of the user that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.GetUser2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.GetUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser2`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV2Api.GetUser2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApps2

> []GetUserApps200ResponseInner GetUserApps2(ctx, userId).IgnoreVisibility(ignoreVisibility).Execute()

Get User Apps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := int32(56) // int32 | Set to the id of the user that you want to return.
    ignoreVisibility := true // bool | Defaults to `false`. When `true` will show all apps that are assigned to a user regardless of their portal visibility setting. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.GetUserApps2(context.Background(), userId).IgnoreVisibility(ignoreVisibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.GetUserApps2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserApps2`: []GetUserApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersV2Api.GetUserApps2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApps2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ignoreVisibility** | **bool** | Defaults to &#x60;false&#x60;. When &#x60;true&#x60; will show all apps that are assigned to a user regardless of their portal visibility setting. | [default to false]

### Return type

[**[]GetUserApps200ResponseInner**](GetUserApps200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers2

> []User ListUsers2(ctx).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()

List Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    createdSince := "createdSince_example" // string | An ISO8601 timestamp value that returns all users created after a given date & time. (optional)
    createdUntil := "createdUntil_example" // string | An ISO8601 timestamp value that returns all users created before a given date & time. (optional)
    updatedSince := "updatedSince_example" // string | An ISO8601 timestamp value that returns all users updated after a given date & time. (optional)
    updatedUntil := "updatedUntil_example" // string | An ISO8601 timestamp value that returns all users updated before a given date & time. (optional)
    lastLoginSince := "lastLoginSince_example" // string | An ISO8601 timestamp value that returns all users that logged in after a given date & time. (optional)
    lastLoginUntil := "lastLoginUntil_example" // string | An ISO8601 timestamp value that returns all users that logged in before a given date & time. (optional)
    firstname := "firstname_example" // string | The first name of the user (optional)
    lastname := "lastname_example" // string | The last name of the user (optional)
    email := "email_example" // string | The email address of the user (optional)
    username := "username_example" // string | The username for the user (optional)
    samaccountname := "samaccountname_example" // string | The AD login name for the user (optional)
    directoryId := int32(56) // int32 |  (optional)
    externalId := "externalId_example" // string | An external identifier that has been set on the user (optional)
    userIds := "userIds_example" // string | A comma separated list of OneLogin User IDs (optional)
    customAttributesAttributeName := "customAttributesAttributeName_example" // string | The short name of a custom attribute. Note that the attribute name is prefixed with custom_attributes. (optional)
    fields := "fields_example" // string | Optional. Comma delimited list of fields to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.ListUsers2(context.Background()).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.ListUsers2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers2`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersV2Api.ListUsers2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsers2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **createdSince** | **string** | An ISO8601 timestamp value that returns all users created after a given date &amp; time. | 
 **createdUntil** | **string** | An ISO8601 timestamp value that returns all users created before a given date &amp; time. | 
 **updatedSince** | **string** | An ISO8601 timestamp value that returns all users updated after a given date &amp; time. | 
 **updatedUntil** | **string** | An ISO8601 timestamp value that returns all users updated before a given date &amp; time. | 
 **lastLoginSince** | **string** | An ISO8601 timestamp value that returns all users that logged in after a given date &amp; time. | 
 **lastLoginUntil** | **string** | An ISO8601 timestamp value that returns all users that logged in before a given date &amp; time. | 
 **firstname** | **string** | The first name of the user | 
 **lastname** | **string** | The last name of the user | 
 **email** | **string** | The email address of the user | 
 **username** | **string** | The username for the user | 
 **samaccountname** | **string** | The AD login name for the user | 
 **directoryId** | **int32** |  | 
 **externalId** | **string** | An external identifier that has been set on the user | 
 **userIds** | **string** | A comma separated list of OneLogin User IDs | 
 **customAttributesAttributeName** | **string** | The short name of a custom attribute. Note that the attribute name is prefixed with custom_attributes. | 
 **fields** | **string** | Optional. Comma delimited list of fields to return. | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser2

> User UpdateUser2(ctx, userId).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()

Update User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := int32(56) // int32 | Set to the id of the user that you want to return.
    mappings := "mappings_example" // string | Controls how mappings will be applied to the user on creation. Defaults to async. (optional)
    validatePolicy := true // bool | Will passwords validate against the User Policy? Defaults to true. (optional)
    user := *openapiclient.NewUser() // User |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV2Api.UpdateUser2(context.Background(), userId).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV2Api.UpdateUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser2`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV2Api.UpdateUser2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

