# \APIAuthClientAppsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClientApp**](APIAuthClientAppsApi.md#AddClientApp) | **Post** /api/2/api_authorizations/{api_auth_id}/clients | Add Client App
[**DeleteClientApp**](APIAuthClientAppsApi.md#DeleteClientApp) | **Delete** /api/2/api_authorizations/{api_auth_id}/clients/{client_app_id} | Remove Client App
[**ListClientApps**](APIAuthClientAppsApi.md#ListClientApps) | **Get** /api/2/api_authorizations/{api_auth_id}/clients | List Clients Apps
[**UpdateClientApp**](APIAuthClientAppsApi.md#UpdateClientApp) | **Put** /api/2/api_authorizations/{api_auth_id}/clients/{client_app_id} | Update Client App



## AddClientApp

> AddClientApp201Response AddClientApp(ctx, apiAuthId).ContentType(contentType).AddClientAppRequest(addClientAppRequest).Execute()

Add Client App



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
    addClientAppRequest := *openapiclient.NewAddClientAppRequest() // AddClientAppRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClientAppsApi.AddClientApp(context.Background(), apiAuthId).ContentType(contentType).AddClientAppRequest(addClientAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClientAppsApi.AddClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClientApp`: AddClientApp201Response
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClientAppsApi.AddClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **addClientAppRequest** | [**AddClientAppRequest**](AddClientAppRequest.md) |  | 

### Return type

[**AddClientApp201Response**](AddClientApp201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientApp

> AddClientApp201Response DeleteClientApp(ctx, apiAuthId, clientAppId).ContentType(contentType).Execute()

Remove Client App



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
    clientAppId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClientAppsApi.DeleteClientApp(context.Background(), apiAuthId, clientAppId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClientAppsApi.DeleteClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientApp`: AddClientApp201Response
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClientAppsApi.DeleteClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**clientAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**AddClientApp201Response**](AddClientApp201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClientApps

> ClientAppFull ListClientApps(ctx, apiAuthId).ContentType(contentType).Execute()

List Clients Apps



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
    resp, r, err := apiClient.APIAuthClientAppsApi.ListClientApps(context.Background(), apiAuthId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClientAppsApi.ListClientApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClientApps`: ClientAppFull
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClientAppsApi.ListClientApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClientAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**ClientAppFull**](ClientAppFull.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientApp

> AddClientApp201Response UpdateClientApp(ctx, apiAuthId, clientAppId).ContentType(contentType).UpdateClientAppRequest(updateClientAppRequest).Execute()

Update Client App



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
    clientAppId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")
    updateClientAppRequest := *openapiclient.NewUpdateClientAppRequest() // UpdateClientAppRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClientAppsApi.UpdateClientApp(context.Background(), apiAuthId, clientAppId).ContentType(contentType).UpdateClientAppRequest(updateClientAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClientAppsApi.UpdateClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientApp`: AddClientApp201Response
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClientAppsApi.UpdateClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**clientAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **updateClientAppRequest** | [**UpdateClientAppRequest**](UpdateClientAppRequest.md) |  | 

### Return type

[**AddClientApp201Response**](AddClientApp201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

