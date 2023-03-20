# \AppsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsApi.md#CreateApp) | **Post** /api/2/apps | Create App
[**DeleteApp**](AppsApi.md#DeleteApp) | **Delete** /api/2/apps/{app_id} | Delete App
[**DeleteAppParameter**](AppsApi.md#DeleteAppParameter) | **Delete** /api/2/apps/{app_id}/parameters/{parameter_id} | Delete Parameter from App
[**GetApp**](AppsApi.md#GetApp) | **Get** /api/2/apps/{app_id} | Get App
[**GetAppUsers**](AppsApi.md#GetAppUsers) | **Get** /api/2/apps/{app_id}/users | Get App Users
[**ListApps**](AppsApi.md#ListApps) | **Get** /api/2/apps | List Apps
[**ListConnectors**](AppsApi.md#ListConnectors) | **Get** /api/2/connectors | List Connectors
[**UpdateApp**](AppsApi.md#UpdateApp) | **Put** /api/2/apps/{app_id} | Update App



## CreateApp

> CreateApp200Response CreateApp(ctx).ContentType(contentType).CreateAppRequest(createAppRequest).Execute()

Create App



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
    createAppRequest := openapiclient.createApp_request{GenericApp: openapiclient.NewGenericApp()} // CreateAppRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.CreateApp(context.Background()).ContentType(contentType).CreateAppRequest(createAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: CreateApp200Response
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **createAppRequest** | [**CreateAppRequest**](CreateAppRequest.md) |  | 

### Return type

[**CreateApp200Response**](CreateApp200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, appId).Execute()

Delete App



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.DeleteApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## DeleteAppParameter

> DeleteAppParameter(ctx, appId, parameterId).Execute()

Delete Parameter from App



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
    parameterId := "parameterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.DeleteAppParameter(context.Background(), appId, parameterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.DeleteAppParameter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**parameterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppParameterRequest struct via the builder pattern


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


## GetApp

> GenericApp GetApp(ctx, appId).Execute()

Get App



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetApp(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: GenericApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericApp**](GenericApp.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppUsers

> []User GetAppUsers(ctx, appId).Execute()

Get App Users



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetAppUsers(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListApps

> []GenericApp ListApps(ctx).Execute()

List Apps



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
    resp, r, err := apiClient.AppsApi.ListApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: []GenericApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


### Return type

[**[]GenericApp**](GenericApp.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> Connector ListConnectors(ctx).Name(name).Execute()

List Connectors



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
    resp, r, err := apiClient.AppsApi.ListConnectors(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectors`: Connector
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**Connector**](Connector.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> GenericApp UpdateApp(ctx, appId).RequestBody(requestBody).Execute()

Update App



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.UpdateApp(context.Background(), appId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UpdateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApp`: GenericApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**GenericApp**](GenericApp.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

