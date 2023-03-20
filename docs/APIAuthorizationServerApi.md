# \APIAuthorizationServerApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthServer**](APIAuthorizationServerApi.md#CreateAuthServer) | **Post** /api/2/api_authorizations | Create Api Auth Server
[**DeleteAuthServer**](APIAuthorizationServerApi.md#DeleteAuthServer) | **Delete** /api/2/api_authorizations/{api_auth_id} | Delete Api Auth Server
[**GetAuthServer**](APIAuthorizationServerApi.md#GetAuthServer) | **Get** /api/2/api_authorizations/{api_auth_id} | Get Api Auth Server
[**ListAuthServers**](APIAuthorizationServerApi.md#ListAuthServers) | **Get** /api/2/api_authorizations | List Api Auth Servers
[**UpdateAuthServer**](APIAuthorizationServerApi.md#UpdateAuthServer) | **Put** /api/2/api_authorizations/{api_auth_id} | Update Api Auth Server



## CreateAuthServer

> AuthServer CreateAuthServer(ctx).ContentType(contentType).AuthServer(authServer).Execute()

Create Api Auth Server



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
    contentType := "application/json" // string |  (optional) (default to "application/json")
    authServer := *openapiclient.NewAuthServer("Contacts API", "API manages contacts", *openapiclient.NewAuthServerConfiguration([]string{"Audiences_example"}, "https://example.com/contacts")) // AuthServer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthorizationServerApi.CreateAuthServer(context.Background()).ContentType(contentType).AuthServer(authServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthorizationServerApi.CreateAuthServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthServer`: AuthServer
    fmt.Fprintf(os.Stdout, "Response from `APIAuthorizationServerApi.CreateAuthServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authServer** | [**AuthServer**](AuthServer.md) |  | 

### Return type

[**AuthServer**](AuthServer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthServer

> DeleteAuthServer(ctx, apiAuthId).ContentType(contentType).Execute()

Delete Api Auth Server



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
    apiAuthId := "apiAuthId_example" // string | 
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthorizationServerApi.DeleteAuthServer(context.Background(), apiAuthId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthorizationServerApi.DeleteAuthServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]

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


## GetAuthServer

> AuthServer GetAuthServer(ctx, apiAuthId).ContentType(contentType).Execute()

Get Api Auth Server



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
    apiAuthId := "apiAuthId_example" // string | 
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthorizationServerApi.GetAuthServer(context.Background(), apiAuthId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthorizationServerApi.GetAuthServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthServer`: AuthServer
    fmt.Fprintf(os.Stdout, "Response from `APIAuthorizationServerApi.GetAuthServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**AuthServer**](AuthServer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthServers

> []AuthServer ListAuthServers(ctx).Execute()

List Api Auth Servers



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
    resp, r, err := apiClient.APIAuthorizationServerApi.ListAuthServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthorizationServerApi.ListAuthServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthServers`: []AuthServer
    fmt.Fprintf(os.Stdout, "Response from `APIAuthorizationServerApi.ListAuthServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthServersRequest struct via the builder pattern


### Return type

[**[]AuthServer**](AuthServer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthServer

> AuthServer UpdateAuthServer(ctx, apiAuthId).ContentType(contentType).AuthServer(authServer).Execute()

Update Api Auth Server



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
    apiAuthId := "apiAuthId_example" // string | 
    contentType := "application/json" // string |  (optional) (default to "application/json")
    authServer := *openapiclient.NewAuthServer("Contacts API", "API manages contacts", *openapiclient.NewAuthServerConfiguration([]string{"Audiences_example"}, "https://example.com/contacts")) // AuthServer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthorizationServerApi.UpdateAuthServer(context.Background(), apiAuthId).ContentType(contentType).AuthServer(authServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthorizationServerApi.UpdateAuthServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthServer`: AuthServer
    fmt.Fprintf(os.Stdout, "Response from `APIAuthorizationServerApi.UpdateAuthServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authServer** | [**AuthServer**](AuthServer.md) |  | 

### Return type

[**AuthServer**](AuthServer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

