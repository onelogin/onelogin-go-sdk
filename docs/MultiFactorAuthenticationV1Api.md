# \MultiFactorAuthenticationV1Api

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateMfaFactors**](MultiFactorAuthenticationV1Api.md#ActivateMfaFactors) | **Post** /api/1/users/{user_id}/otp_devices/{device_id}/trigger | Activate a Factor
[**EnrollMfaFactor**](MultiFactorAuthenticationV1Api.md#EnrollMfaFactor) | **Post** /api/1/users/{user_id}/otp_devices | Enroll a Factor
[**GenerateMFAtoken**](MultiFactorAuthenticationV1Api.md#GenerateMFAtoken) | **Post** /api/1/users/{user_id}/mfa_token | Generate Temp MFA Token
[**GetEnrolledFactors**](MultiFactorAuthenticationV1Api.md#GetEnrolledFactors) | **Get** /api/1/users/{user_id}/otp_devices | Get Enrolled Factors
[**GetMFAFactors**](MultiFactorAuthenticationV1Api.md#GetMFAFactors) | **Get** /api/1/users/{user_id}/auth_factor | Get Available Factors
[**RemoveMfaFactors**](MultiFactorAuthenticationV1Api.md#RemoveMfaFactors) | **Delete** /api/1/users/{user_id}/otp_devices/{device_id} | Remove an Enrolled Factor
[**VerifyMfaFactor**](MultiFactorAuthenticationV1Api.md#VerifyMfaFactor) | **Post** /api/1/users/{user_id}/otp_devices/{device_id}/verify | Verify a Factor



## ActivateMfaFactors

> GetEnrolledFactors200Response ActivateMfaFactors(ctx, userId, deviceId).ActivateMfaFactorsRequest(activateMfaFactorsRequest).Execute()

Activate a Factor



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
    deviceId := "deviceId_example" // string | 
    activateMfaFactorsRequest := *openapiclient.NewActivateMfaFactorsRequest() // ActivateMfaFactorsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.ActivateMfaFactors(context.Background(), userId, deviceId).ActivateMfaFactorsRequest(activateMfaFactorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.ActivateMfaFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateMfaFactors`: GetEnrolledFactors200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.ActivateMfaFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateMfaFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activateMfaFactorsRequest** | [**ActivateMfaFactorsRequest**](ActivateMfaFactorsRequest.md) |  | 

### Return type

[**GetEnrolledFactors200Response**](GetEnrolledFactors200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollMfaFactor

> EnrollMfaFactor200Response EnrollMfaFactor(ctx, userId).OtpDevice(otpDevice).Execute()

Enroll a Factor



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
    otpDevice := *openapiclient.NewOtpDevice(int32(16282), "Rich's Phone", "+1xxxxxxxxxx") // OtpDevice |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.EnrollMfaFactor(context.Background(), userId).OtpDevice(otpDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.EnrollMfaFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollMfaFactor`: EnrollMfaFactor200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.EnrollMfaFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollMfaFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **otpDevice** | [**OtpDevice**](OtpDevice.md) |  | 

### Return type

[**EnrollMfaFactor200Response**](EnrollMfaFactor200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateMFAtoken

> GenerateMFAtoken200Response GenerateMFAtoken(ctx, userId).GenerateMFAtokenRequest(generateMFAtokenRequest).Execute()

Generate Temp MFA Token



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
    generateMFAtokenRequest := *openapiclient.NewGenerateMFAtokenRequest() // GenerateMFAtokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.GenerateMFAtoken(context.Background(), userId).GenerateMFAtokenRequest(generateMFAtokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.GenerateMFAtoken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateMFAtoken`: GenerateMFAtoken200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.GenerateMFAtoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMFAtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateMFAtokenRequest** | [**GenerateMFAtokenRequest**](GenerateMFAtokenRequest.md) |  | 

### Return type

[**GenerateMFAtoken200Response**](GenerateMFAtoken200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrolledFactors

> GetEnrolledFactors200Response GetEnrolledFactors(ctx, userId).Execute()

Get Enrolled Factors



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
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.GetEnrolledFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.GetEnrolledFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrolledFactors`: GetEnrolledFactors200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.GetEnrolledFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrolledFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnrolledFactors200Response**](GetEnrolledFactors200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAFactors

> GetMFAFactors200Response GetMFAFactors(ctx, userId).Execute()

Get Available Factors



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
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.GetMFAFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.GetMFAFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAFactors`: GetMFAFactors200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.GetMFAFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMFAFactors200Response**](GetMFAFactors200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMfaFactors

> RemoveMfaFactors(ctx, userId, deviceId).Execute()

Remove an Enrolled Factor



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
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.RemoveMfaFactors(context.Background(), userId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.RemoveMfaFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMfaFactorsRequest struct via the builder pattern


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


## VerifyMfaFactor

> Error VerifyMfaFactor(ctx, userId, deviceId).VerifyMfaFactorRequest(verifyMfaFactorRequest).Execute()

Verify a Factor



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
    deviceId := "deviceId_example" // string | 
    verifyMfaFactorRequest := *openapiclient.NewVerifyMfaFactorRequest() // VerifyMfaFactorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationV1Api.VerifyMfaFactor(context.Background(), userId, deviceId).VerifyMfaFactorRequest(verifyMfaFactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationV1Api.VerifyMfaFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMfaFactor`: Error
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationV1Api.VerifyMfaFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMfaFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verifyMfaFactorRequest** | [**VerifyMfaFactorRequest**](VerifyMfaFactorRequest.md) |  | 

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

