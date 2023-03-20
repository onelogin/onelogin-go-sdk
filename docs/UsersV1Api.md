# \UsersV1Api

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRolesToUser**](UsersV1Api.md#AddRolesToUser) | **Put** /api/1/users/{user_id}/add_roles | Add Roles for a User
[**CreateUser**](UsersV1Api.md#CreateUser) | **Post** /api/1/users | Create a User
[**DeleteUser**](UsersV1Api.md#DeleteUser) | **Delete** /api/1/users/{user_id} | Delete a User
[**GetCustomAttributes**](UsersV1Api.md#GetCustomAttributes) | **Get** /api/1/users/custom_attributes | Get Custom Attributes
[**GetUserApps**](UsersV1Api.md#GetUserApps) | **Get** /api/1/users/{user_id}/apps | Get Apps for a User
[**GetUserById**](UsersV1Api.md#GetUserById) | **Get** /api/1/users/{user_id} | Get User by ID
[**GetUserRoles**](UsersV1Api.md#GetUserRoles) | **Get** /api/1/users/{user_id}/roles | Get Roles for a User
[**ListUsers**](UsersV1Api.md#ListUsers) | **Get** /api/1/users | List Users
[**LockAccountUser**](UsersV1Api.md#LockAccountUser) | **Put** /api/1/users/{user_id}/lock_user | Lock User Account
[**LogOutUser**](UsersV1Api.md#LogOutUser) | **Put** /api/1/users/{user_id}/logout | Log User Out
[**RemoveUserRole**](UsersV1Api.md#RemoveUserRole) | **Put** /api/1/users/{user_id}/remove_roles | Remove Roles for a User
[**SetUserState**](UsersV1Api.md#SetUserState) | **Put** /api/1/users/{user_id}/set_state | Set User State
[**UpdatePasswordInsecure**](UsersV1Api.md#UpdatePasswordInsecure) | **Put** /api/1/users/set_password_clear_text/{user_id} | Set Password Using ID in Cleartext
[**UpdatePasswordSecure**](UsersV1Api.md#UpdatePasswordSecure) | **Put** /api/1/users/set_password_using_salt/{user_id} | Set Password Using ID and SHA-256 and Salt
[**UpdateUser**](UsersV1Api.md#UpdateUser) | **Put** /api/1/users/{user_id} | Update a User



## AddRolesToUser

> Error AddRolesToUser(ctx, userId).AddRolesToUserRequest(addRolesToUserRequest).Execute()

Add Roles for a User



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
    addRolesToUserRequest := *openapiclient.NewAddRolesToUserRequest([]int32{int32(123)}) // AddRolesToUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.AddRolesToUser(context.Background(), userId).AddRolesToUserRequest(addRolesToUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.AddRolesToUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRolesToUser`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.AddRolesToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRolesToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addRolesToUserRequest** | [**AddRolesToUserRequest**](AddRolesToUserRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()

Create a User



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
    resp, r, err := apiClient.UsersV1Api.CreateUser(context.Background()).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete a User



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
    resp, r, err := apiClient.UsersV1Api.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.DeleteUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetCustomAttributes

> GetCustomAttributes200Response GetCustomAttributes(ctx).Execute()

Get Custom Attributes



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.GetCustomAttributes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.GetCustomAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomAttributes`: GetCustomAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.GetCustomAttributes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAttributesRequest struct via the builder pattern


### Return type

[**GetCustomAttributes200Response**](GetCustomAttributes200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApps

> []GetUserApps200ResponseInner GetUserApps(ctx, userId).IgnoreVisibility(ignoreVisibility).Execute()

Get Apps for a User



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
    resp, r, err := apiClient.UsersV1Api.GetUserApps(context.Background(), userId).IgnoreVisibility(ignoreVisibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.GetUserApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserApps`: []GetUserApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.GetUserApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppsRequest struct via the builder pattern


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


## GetUserById

> User GetUserById(ctx, userId).Execute()

Get User by ID



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
    resp, r, err := apiClient.UsersV1Api.GetUserById(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.GetUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserById`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


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


## GetUserRoles

> GetUserRoles200Response GetUserRoles(ctx, userId).Execute()

Get Roles for a User



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
    resp, r, err := apiClient.UsersV1Api.GetUserRoles(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.GetUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRoles`: GetUserRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.GetUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserRoles200Response**](GetUserRoles200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()

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
    resp, r, err := apiClient.UsersV1Api.ListUsers(context.Background()).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


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


## LockAccountUser

> Error LockAccountUser(ctx, userId).LockAccountUserRequest(lockAccountUserRequest).Execute()

Lock User Account



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
    lockAccountUserRequest := *openapiclient.NewLockAccountUserRequest(int32(15)) // LockAccountUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.LockAccountUser(context.Background(), userId).LockAccountUserRequest(lockAccountUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.LockAccountUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockAccountUser`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.LockAccountUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockAccountUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockAccountUserRequest** | [**LockAccountUserRequest**](LockAccountUserRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOutUser

> Error LogOutUser(ctx, userId).Body(body).Execute()

Log User Out



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.LogOutUser(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.LogOutUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogOutUser`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.LogOutUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogOutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserRole

> Error RemoveUserRole(ctx, userId).RemoveUserRoleRequest(removeUserRoleRequest).Execute()

Remove Roles for a User



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
    removeUserRoleRequest := *openapiclient.NewRemoveUserRoleRequest([]openapiclient.RemoveUserRoleRequestRoleIdArrayInner{*openapiclient.NewRemoveUserRoleRequestRoleIdArrayInner()}) // RemoveUserRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.RemoveUserRole(context.Background(), userId).RemoveUserRoleRequest(removeUserRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.RemoveUserRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserRole`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.RemoveUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeUserRoleRequest** | [**RemoveUserRoleRequest**](RemoveUserRoleRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserState

> Error SetUserState(ctx, userId).SetUserStateRequest(setUserStateRequest).Execute()

Set User State



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
    setUserStateRequest := *openapiclient.NewSetUserStateRequest(int32(1)) // SetUserStateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.SetUserState(context.Background(), userId).SetUserStateRequest(setUserStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.SetUserState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserState`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.SetUserState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setUserStateRequest** | [**SetUserStateRequest**](SetUserStateRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordInsecure

> Error UpdatePasswordInsecure(ctx, userId).UpdatePasswordInsecureRequest(updatePasswordInsecureRequest).Execute()

Set Password Using ID in Cleartext



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
    updatePasswordInsecureRequest := *openapiclient.NewUpdatePasswordInsecureRequest("<password>", "<password_confirmation>") // UpdatePasswordInsecureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.UpdatePasswordInsecure(context.Background(), userId).UpdatePasswordInsecureRequest(updatePasswordInsecureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.UpdatePasswordInsecure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordInsecure`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.UpdatePasswordInsecure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordInsecureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePasswordInsecureRequest** | [**UpdatePasswordInsecureRequest**](UpdatePasswordInsecureRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordSecure

> Error UpdatePasswordSecure(ctx, userId).UpdatePasswordSecureRequest(updatePasswordSecureRequest).Execute()

Set Password Using ID and SHA-256 and Salt



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
    updatePasswordSecureRequest := *openapiclient.NewUpdatePasswordSecureRequest("xxxxx637aead4030a653f29dae62f1542d67484342c00627a65066e05c5f0", "xxxxx637aead4030a653f29dae62f1542d67484342c00627a65066e05c5f0", "salt+sha256") // UpdatePasswordSecureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersV1Api.UpdatePasswordSecure(context.Background(), userId).UpdatePasswordSecureRequest(updatePasswordSecureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.UpdatePasswordSecure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordSecure`: Error
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.UpdatePasswordSecure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordSecureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePasswordSecureRequest** | [**UpdatePasswordSecureRequest**](UpdatePasswordSecureRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()

Update a User



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
    resp, r, err := apiClient.UsersV1Api.UpdateUser(context.Background(), userId).Mappings(mappings).ValidatePolicy(validatePolicy).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

