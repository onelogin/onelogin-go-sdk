# \APIAuthScopesApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScope**](APIAuthScopesApi.md#CreateScope) | **Post** /api/2/api_authorizations/{api_auth_id}/scopes | Create Api Auth Server Scope
[**DeleteScope**](APIAuthScopesApi.md#DeleteScope) | **Delete** /api/2/api_authorizations/{api_auth_id}/scopes/{scope_id} | Delete Api Auth Server Scope
[**GetScopes**](APIAuthScopesApi.md#GetScopes) | **Get** /api/2/api_authorizations/{api_auth_id}/scopes | Get Api Auth Server Scopes
[**UpdateScope**](APIAuthScopesApi.md#UpdateScope) | **Put** /api/2/api_authorizations/{api_auth_id}/scopes/{scope_id} | Update Api Auth Server Scope



## CreateScope

> AuthScope CreateScope(ctx, apiAuthId).ContentType(contentType).AuthScope(authScope).Execute()

Create Api Auth Server Scope



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
    authScope := *openapiclient.NewAuthScope() // AuthScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthScopesApi.CreateScope(context.Background(), apiAuthId).ContentType(contentType).AuthScope(authScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthScopesApi.CreateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScope`: AuthScope
    fmt.Fprintf(os.Stdout, "Response from `APIAuthScopesApi.CreateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authScope** | [**AuthScope**](AuthScope.md) |  | 

### Return type

[**AuthScope**](AuthScope.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScope

> DeleteScope(ctx, apiAuthId, scopeId).ContentType(contentType).Execute()

Delete Api Auth Server Scope



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
    scopeId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthScopesApi.DeleteScope(context.Background(), apiAuthId, scopeId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthScopesApi.DeleteScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**scopeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopeRequest struct via the builder pattern


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


## GetScopes

> []AuthScope GetScopes(ctx, apiAuthId).ContentType(contentType).Execute()

Get Api Auth Server Scopes



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
    resp, r, err := apiClient.APIAuthScopesApi.GetScopes(context.Background(), apiAuthId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthScopesApi.GetScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopes`: []AuthScope
    fmt.Fprintf(os.Stdout, "Response from `APIAuthScopesApi.GetScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**[]AuthScope**](AuthScope.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScope

> AuthId UpdateScope(ctx, apiAuthId, scopeId).ContentType(contentType).AuthScope(authScope).Execute()

Update Api Auth Server Scope



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
    scopeId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")
    authScope := *openapiclient.NewAuthScope() // AuthScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthScopesApi.UpdateScope(context.Background(), apiAuthId, scopeId).ContentType(contentType).AuthScope(authScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthScopesApi.UpdateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScope`: AuthId
    fmt.Fprintf(os.Stdout, "Response from `APIAuthScopesApi.UpdateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**scopeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authScope** | [**AuthScope**](AuthScope.md) |  | 

### Return type

[**AuthId**](AuthId.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

