# \PrivilegesApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivilegeToRole**](PrivilegesApi.md#AddPrivilegeToRole) | **Post** /api/1/privileges/{privilege_id}/roles | Assign a Privilege to Roles
[**AssignUsersToPrivilege**](PrivilegesApi.md#AssignUsersToPrivilege) | **Post** /api/1/privileges/{privilege_id}/users | Assign Users to a Privilege
[**CreatePrivilege**](PrivilegesApi.md#CreatePrivilege) | **Post** /api/1/privileges | Create a Privilege
[**DeletePrivilege**](PrivilegesApi.md#DeletePrivilege) | **Delete** /api/1/privileges/{privilege_id} | Delete a Privilege
[**DeleteRoleFromPrivilege**](PrivilegesApi.md#DeleteRoleFromPrivilege) | **Delete** /api/1/privileges/{privilege_id}/roles/{role_id} | Remove a Privilege from a Role
[**GetAssignedUser**](PrivilegesApi.md#GetAssignedUser) | **Get** /api/1/privileges/{privilege_id}/users | Get Users assigned to a Privilege
[**GetPrivilege**](PrivilegesApi.md#GetPrivilege) | **Get** /api/1/privileges/{privilege_id} | Get a Privilege
[**ListPrivelegeRoles**](PrivilegesApi.md#ListPrivelegeRoles) | **Get** /api/1/privileges/{privilege_id}/roles | Get Roles assigned to Privilege
[**ListPriveleges**](PrivilegesApi.md#ListPriveleges) | **Get** /api/1/privileges | List Privileges
[**RemoveUserFromPrivilege**](PrivilegesApi.md#RemoveUserFromPrivilege) | **Delete** /api/1/privileges/{privilege_id}/users/{user_id} | Remove a Privilege from Users
[**UpdatePrivilege**](PrivilegesApi.md#UpdatePrivilege) | **Put** /api/1/privileges/{privilege_id} | Update a Privilege



## AddPrivilegeToRole

> AddPrivilegeToRole201Response AddPrivilegeToRole(ctx, privilegeId).AddPrivilegeToRoleRequest(addPrivilegeToRoleRequest).Execute()

Assign a Privilege to Roles



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
    privilegeId := "privilegeId_example" // string | 
    addPrivilegeToRoleRequest := *openapiclient.NewAddPrivilegeToRoleRequest() // AddPrivilegeToRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.AddPrivilegeToRole(context.Background(), privilegeId).AddPrivilegeToRoleRequest(addPrivilegeToRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.AddPrivilegeToRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrivilegeToRole`: AddPrivilegeToRole201Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.AddPrivilegeToRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivilegeToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPrivilegeToRoleRequest** | [**AddPrivilegeToRoleRequest**](AddPrivilegeToRoleRequest.md) |  | 

### Return type

[**AddPrivilegeToRole201Response**](AddPrivilegeToRole201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignUsersToPrivilege

> AddPrivilegeToRole201Response AssignUsersToPrivilege(ctx, privilegeId).AssignUsersToPrivilegeRequest(assignUsersToPrivilegeRequest).Execute()

Assign Users to a Privilege



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
    privilegeId := "privilegeId_example" // string | 
    assignUsersToPrivilegeRequest := *openapiclient.NewAssignUsersToPrivilegeRequest() // AssignUsersToPrivilegeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.AssignUsersToPrivilege(context.Background(), privilegeId).AssignUsersToPrivilegeRequest(assignUsersToPrivilegeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.AssignUsersToPrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignUsersToPrivilege`: AddPrivilegeToRole201Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.AssignUsersToPrivilege`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUsersToPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignUsersToPrivilegeRequest** | [**AssignUsersToPrivilegeRequest**](AssignUsersToPrivilegeRequest.md) |  | 

### Return type

[**AddPrivilegeToRole201Response**](AddPrivilegeToRole201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege

> CreatePrivilege200Response CreatePrivilege(ctx).Privilege(privilege).Execute()

Create a Privilege



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
    privilege := *openapiclient.NewPrivilege("User Administrator", *openapiclient.NewPrivilegePrivilege()) // Privilege |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.CreatePrivilege(context.Background()).Privilege(privilege).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.CreatePrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivilege`: CreatePrivilege200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.CreatePrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privilege** | [**Privilege**](Privilege.md) |  | 

### Return type

[**CreatePrivilege200Response**](CreatePrivilege200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivilege

> DeletePrivilege(ctx, privilegeId).Execute()

Delete a Privilege



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
    privilegeId := "privilegeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.DeletePrivilege(context.Background(), privilegeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.DeletePrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivilegeRequest struct via the builder pattern


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


## DeleteRoleFromPrivilege

> DeleteRoleFromPrivilege(ctx, privilegeId, roleId).Execute()

Remove a Privilege from a Role



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
    privilegeId := "privilegeId_example" // string | 
    roleId := "roleId_example" // string | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.DeleteRoleFromPrivilege(context.Background(), privilegeId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.DeleteRoleFromPrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 
**roleId** | **string** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleFromPrivilegeRequest struct via the builder pattern


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


## GetAssignedUser

> GetAssignedUser200Response GetAssignedUser(ctx, privilegeId).Execute()

Get Users assigned to a Privilege



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
    privilegeId := "privilegeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.GetAssignedUser(context.Background(), privilegeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.GetAssignedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignedUser`: GetAssignedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.GetAssignedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssignedUser200Response**](GetAssignedUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivilege

> Privilege GetPrivilege(ctx, privilegeId).Execute()

Get a Privilege



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
    privilegeId := "privilegeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.GetPrivilege(context.Background(), privilegeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.GetPrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivilege`: Privilege
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.GetPrivilege`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Privilege**](Privilege.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivelegeRoles

> ListPrivelegeRoles200Response ListPrivelegeRoles(ctx, privilegeId).Execute()

Get Roles assigned to Privilege



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
    privilegeId := "privilegeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.ListPrivelegeRoles(context.Background(), privilegeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.ListPrivelegeRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivelegeRoles`: ListPrivelegeRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.ListPrivelegeRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivelegeRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPrivelegeRoles200Response**](ListPrivelegeRoles200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPriveleges

> []Privilege ListPriveleges(ctx).Execute()

List Privileges



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
    resp, r, err := apiClient.PrivilegesApi.ListPriveleges(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.ListPriveleges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPriveleges`: []Privilege
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.ListPriveleges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivelegesRequest struct via the builder pattern


### Return type

[**[]Privilege**](Privilege.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromPrivilege

> RemoveUserFromPrivilege(ctx, privilegeId, userId).Execute()

Remove a Privilege from Users



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
    privilegeId := "privilegeId_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.RemoveUserFromPrivilege(context.Background(), privilegeId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.RemoveUserFromPrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromPrivilegeRequest struct via the builder pattern


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


## UpdatePrivilege

> UpdatePrivilege200Response UpdatePrivilege(ctx, privilegeId).Privilege(privilege).Execute()

Update a Privilege



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
    privilegeId := "privilegeId_example" // string | 
    privilege := *openapiclient.NewPrivilege("User Administrator", *openapiclient.NewPrivilegePrivilege()) // Privilege |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivilegesApi.UpdatePrivilege(context.Background(), privilegeId).Privilege(privilege).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.UpdatePrivilege``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrivilege`: UpdatePrivilege200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.UpdatePrivilege`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privilege** | [**Privilege**](Privilege.md) |  | 

### Return type

[**UpdatePrivilege200Response**](UpdatePrivilege200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

