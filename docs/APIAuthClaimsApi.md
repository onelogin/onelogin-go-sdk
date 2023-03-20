# \APIAuthClaimsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthClaim**](APIAuthClaimsApi.md#CreateAuthClaim) | **Post** /api/2/api_authorizations/{api_auth_id}/claims | Create Api Auth Server Claim
[**DeleteAuthClaim**](APIAuthClaimsApi.md#DeleteAuthClaim) | **Delete** /api/2/api_authorizations/{api_auth_id}/claims/{claim_id} | Delete Api Auth Server Claim
[**GetAuthclaims**](APIAuthClaimsApi.md#GetAuthclaims) | **Get** /api/2/api_authorizations/{api_auth_id}/claims | Get Api Auth Server claims
[**UpdateClaim**](APIAuthClaimsApi.md#UpdateClaim) | **Put** /api/2/api_authorizations/{api_auth_id}/claims/{claim_id} | Update Api Auth Server Claim



## CreateAuthClaim

> int32 CreateAuthClaim(ctx, apiAuthId).ContentType(contentType).AuthClaim(authClaim).Execute()

Create Api Auth Server Claim



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
    authClaim := *openapiclient.NewAuthClaim("email_address") // AuthClaim |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClaimsApi.CreateAuthClaim(context.Background(), apiAuthId).ContentType(contentType).AuthClaim(authClaim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClaimsApi.CreateAuthClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthClaim`: int32
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClaimsApi.CreateAuthClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authClaim** | [**AuthClaim**](AuthClaim.md) |  | 

### Return type

**int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthClaim

> DeleteAuthClaim(ctx, apiAuthId, claimId).ContentType(contentType).Execute()

Delete Api Auth Server Claim



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
    claimId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClaimsApi.DeleteAuthClaim(context.Background(), apiAuthId, claimId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClaimsApi.DeleteAuthClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**claimId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthClaimRequest struct via the builder pattern


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


## GetAuthclaims

> []TokenClaim GetAuthclaims(ctx, apiAuthId).ContentType(contentType).Execute()

Get Api Auth Server claims



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
    resp, r, err := apiClient.APIAuthClaimsApi.GetAuthclaims(context.Background(), apiAuthId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClaimsApi.GetAuthclaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthclaims`: []TokenClaim
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClaimsApi.GetAuthclaims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthclaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**[]TokenClaim**](TokenClaim.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClaim

> AuthId UpdateClaim(ctx, apiAuthId, claimId).ContentType(contentType).AuthClaim(authClaim).Execute()

Update Api Auth Server Claim



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
    claimId := int32(56) // int32 | 
    contentType := "application/json" // string |  (optional) (default to "application/json")
    authClaim := *openapiclient.NewAuthClaim("email_address") // AuthClaim |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIAuthClaimsApi.UpdateClaim(context.Background(), apiAuthId, claimId).ContentType(contentType).AuthClaim(authClaim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIAuthClaimsApi.UpdateClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClaim`: AuthId
    fmt.Fprintf(os.Stdout, "Response from `APIAuthClaimsApi.UpdateClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiAuthId** | **string** |  | 
**claimId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **authClaim** | [**AuthClaim**](AuthClaim.md) |  | 

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

