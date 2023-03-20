# \RolesApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleAdmins**](RolesApi.md#AddRoleAdmins) | **Post** /api/2/roles/{role_id}/admins | Add Role Admins
[**AddRoleUsers**](RolesApi.md#AddRoleUsers) | **Post** /api/2/roles/{role_id}/users | Add Role Users
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /api/2/roles | Create Role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /api/2/roles/{role_id} | Delete Role by ID
[**GetRole**](RolesApi.md#GetRole) | **Get** /api/2/roles/{role_id} | Get Role by ID
[**GetRoleAdmins**](RolesApi.md#GetRoleAdmins) | **Get** /api/2/roles/{role_id}/admins | Get Role Admins
[**GetRoleApps**](RolesApi.md#GetRoleApps) | **Get** /api/2/roles/{role_id}/apps | Get all Apps assigned to Role
[**GetRoleById**](RolesApi.md#GetRoleById) | **Get** /api/1/roles/{role_id} | Get Role by ID
[**GetRoleByName**](RolesApi.md#GetRoleByName) | **Get** /api/1/roles | Get Role by Name
[**GetRoleUsers**](RolesApi.md#GetRoleUsers) | **Get** /api/2/roles/{role_id}/users | Get Role Users
[**ListRoles**](RolesApi.md#ListRoles) | **Get** /api/2/roles | List Roles
[**RemoveRoleAdmins**](RolesApi.md#RemoveRoleAdmins) | **Delete** /api/2/roles/{role_id}/admins | Remove Role Admins
[**RemoveRoleUsers**](RolesApi.md#RemoveRoleUsers) | **Delete** /api/2/roles/{role_id}/users | Remove Role Users
[**SetRoleApps**](RolesApi.md#SetRoleApps) | **Put** /api/2/roles/{role_id}/apps | Set Role Apps
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /api/2/roles/{role_id} | Update Role



## AddRoleAdmins

> []CreateRole201ResponseInner AddRoleAdmins(ctx, roleId).RequestBody(requestBody).Execute()

Add Role Admins



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.AddRoleAdmins(context.Background(), roleId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleAdmins`: []CreateRole201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddRoleAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int32** |  | 

### Return type

[**[]CreateRole201ResponseInner**](CreateRole201ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleUsers

> []CreateRole201ResponseInner AddRoleUsers(ctx, roleId).RequestBody(requestBody).Execute()

Add Role Users



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.AddRoleUsers(context.Background(), roleId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleUsers`: []CreateRole201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.AddRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int32** |  | 

### Return type

[**[]CreateRole201ResponseInner**](CreateRole201ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> []CreateRole201ResponseInner CreateRole(ctx).Role(role).Execute()

Create Role



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
    role := *openapiclient.NewRole("Name_example") // Role |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateRole(context.Background()).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: []CreateRole201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**[]CreateRole201ResponseInner**](CreateRole201ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Delete Role by ID



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.DeleteRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetRole

> Role GetRole(ctx, roleId).Execute()

Get Role by ID



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRole(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAdmins

> []User GetRoleAdmins(ctx, roleId).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()

Get Role Admins



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | Allows you to filter on first name, last name, username, and email address. (optional)
    includeUnassigned := true // bool | Optional. Defaults to false. Include users that aren’t assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleAdmins(context.Background(), roleId).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleAdmins`: []User
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **bool** | Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleApps

> []GetRoleApps200ResponseInner GetRoleApps(ctx, roleId).Limit(limit).Page(page).Cursor(cursor).Assigned(assigned).Execute()

Get all Apps assigned to Role



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    assigned := true // bool | Optional. Defaults to true. Returns all apps not yet assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleApps(context.Background(), roleId).Limit(limit).Page(page).Cursor(cursor).Assigned(assigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleApps`: []GetRoleApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **assigned** | **bool** | Optional. Defaults to true. Returns all apps not yet assigned to the role. | 

### Return type

[**[]GetRoleApps200ResponseInner**](GetRoleApps200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleById

> GetRoleById200Response GetRoleById(ctx, roleId).Execute()

Get Role by ID



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleById(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleById`: GetRoleById200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRoleById200Response**](GetRoleById200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleByName

> GetRoleByName200Response GetRoleByName(ctx).Name(name).Execute()

Get Role by Name



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
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleByName(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleByName`: GetRoleByName200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**GetRoleByName200Response**](GetRoleByName200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleUsers

> []User GetRoleUsers(ctx, roleId).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()

Get Role Users



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | Allows you to filter on first name, last name, username, and email address. (optional)
    includeUnassigned := true // bool | Optional. Defaults to false. Include users that aren’t assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRoleUsers(context.Background(), roleId).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **bool** | Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []Role ListRoles(ctx).AppId(appId).Limit(limit).Page(page).Cursor(cursor).RoleName(roleName).AppName(appName).Fields(fields).Execute()

List Roles



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
    appId := int32(56) // int32 | 
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    roleName := "roleName_example" // string | Optional. Filters by role name. (optional)
    appName := "appName_example" // string | Optional. Returns roles that contain this app name. (optional)
    fields := "fields_example" // string | Optional. Comma delimited list of fields to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRoles(context.Background()).AppId(appId).Limit(limit).Page(page).Cursor(cursor).RoleName(roleName).AppName(appName).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **int32** |  | 
 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **roleName** | **string** | Optional. Filters by role name. | 
 **appName** | **string** | Optional. Returns roles that contain this app name. | 
 **fields** | **string** | Optional. Comma delimited list of fields to return. | 

### Return type

[**[]Role**](Role.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleAdmins

> RemoveRoleAdmins(ctx, roleId).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()

Remove Role Admins



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    removeRoleUsersRequest := *openapiclient.NewRemoveRoleUsersRequest() // RemoveRoleUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RemoveRoleAdmins(context.Background(), roleId).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemoveRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeRoleUsersRequest** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleUsers

> RemoveRoleUsers(ctx, roleId).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()

Remove Role Users



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    removeRoleUsersRequest := *openapiclient.NewRemoveRoleUsersRequest() // RemoveRoleUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RemoveRoleUsers(context.Background(), roleId).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemoveRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeRoleUsersRequest** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleApps

> []CreateRole201ResponseInner SetRoleApps(ctx, roleId).RequestBody(requestBody).Execute()

Set Role Apps



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.SetRoleApps(context.Background(), roleId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.SetRoleApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRoleApps`: []CreateRole201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.SetRoleApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int32** |  | 

### Return type

[**[]CreateRole201ResponseInner**](CreateRole201ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole200Response UpdateRole(ctx, roleId).Role(role).Execute()

Update Role



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
    roleId := "roleId_example" // string | Set to the id of the role you want to return.
    role := *openapiclient.NewRole("Name_example") // Role |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateRole(context.Background(), roleId).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: UpdateRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**UpdateRole200Response**](UpdateRole200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

