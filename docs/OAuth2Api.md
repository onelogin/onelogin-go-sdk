# \OAuth2Api

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateToken**](OAuth2Api.md#GenerateToken) | **Post** /auth/oauth2/v2/token | Generate Token
[**GetRateLimit**](OAuth2Api.md#GetRateLimit) | **Get** /auth/rate_limit | Get Rate Limit
[**RevokeTokens**](OAuth2Api.md#RevokeTokens) | **Post** /auth/oauth2/revoke | Revoke Tokens



## GenerateToken

> OauthToken GenerateToken(ctx).GenerateTokenRequest(generateTokenRequest).ContentType(contentType).Execute()

Generate Token



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
    generateTokenRequest := *openapiclient.NewGenerateTokenRequest("GrantType_example") // GenerateTokenRequest | Request Body to Generate OAuth Token
    contentType := "application/json" // string |  (optional) (default to "application/json")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2Api.GenerateToken(context.Background()).GenerateTokenRequest(generateTokenRequest).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2Api.GenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateToken`: OauthToken
    fmt.Fprintf(os.Stdout, "Response from `OAuth2Api.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateTokenRequest** | [**GenerateTokenRequest**](GenerateTokenRequest.md) | Request Body to Generate OAuth Token | 
 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**OauthToken**](OauthToken.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimit

> GetRateLimit200Response GetRateLimit(ctx).Execute()

Get Rate Limit



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
    resp, r, err := apiClient.OAuth2Api.GetRateLimit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2Api.GetRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimit`: GetRateLimit200Response
    fmt.Fprintf(os.Stdout, "Response from `OAuth2Api.GetRateLimit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitRequest struct via the builder pattern


### Return type

[**GetRateLimit200Response**](GetRateLimit200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeTokens

> Error RevokeTokens(ctx).ContentType(contentType).RevokeTokensRequest(revokeTokensRequest).Execute()

Revoke Tokens



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
    revokeTokensRequest := *openapiclient.NewRevokeTokensRequest("xx508xx63817x752xx74004x30705xx92x58349x5x78f5xx34x8x614xxxx1451") // RevokeTokensRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2Api.RevokeTokens(context.Background()).ContentType(contentType).RevokeTokensRequest(revokeTokensRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2Api.RevokeTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeTokens`: Error
    fmt.Fprintf(os.Stdout, "Response from `OAuth2Api.RevokeTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **revokeTokensRequest** | [**RevokeTokensRequest**](RevokeTokensRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

