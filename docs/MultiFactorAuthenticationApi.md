# \MultiFactorAuthenticationApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceVerification**](MultiFactorAuthenticationApi.md#CreateDeviceVerification) | **Post** /api/2/mfa/users/{user_id}/verifications | Create Device Verification
[**CreateFactorRegistration**](MultiFactorAuthenticationApi.md#CreateFactorRegistration) | **Post** /api/2/mfa/users/{user_id}/registrations | Create Factor Registration
[**DeleteEnrolledFactor**](MultiFactorAuthenticationApi.md#DeleteEnrolledFactor) | **Delete** /api/2/mfa/users/{user_id}/devices/{device_id} | Delete Enrolled Factor
[**GenerateOTP**](MultiFactorAuthenticationApi.md#GenerateOTP) | **Post** /api/2/mfa/users/{user_id}/mfa_token | Generate MFA token
[**GetAuthFactors**](MultiFactorAuthenticationApi.md#GetAuthFactors) | **Get** /api/2/mfa/users/{user_id}/factors | Get User Factors
[**GetAuthenticationDevices**](MultiFactorAuthenticationApi.md#GetAuthenticationDevices) | **Get** /api/2/mfa/users/{user_id}/devices | Get User Devices
[**GetUserRegistration**](MultiFactorAuthenticationApi.md#GetUserRegistration) | **Get** /api/2/mfa/users/{user_id}/registrations/{registration_id} | Get User Registration
[**GetUserVerification**](MultiFactorAuthenticationApi.md#GetUserVerification) | **Get** /api/2/mfa/users/{user_id}/verifications/{verification_id} | Get User Verification
[**VerifyUserRegistration**](MultiFactorAuthenticationApi.md#VerifyUserRegistration) | **Put** /api/2/mfa/users/{user_id}/registrations/{registration_id} | Verify User Registration
[**VerifyUserVerification**](MultiFactorAuthenticationApi.md#VerifyUserVerification) | **Put** /api/2/mfa/users/{user_id}/verifications/{verification_id} | Verify User Verification



## CreateDeviceVerification

> CreateDeviceVerification201Response CreateDeviceVerification(ctx, userId).ContentType(contentType).CreateDeviceVerificationRequest(createDeviceVerificationRequest).Execute()

Create Device Verification



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
    contentType := "application/json" // string |  (optional) (default to "application/json")
    createDeviceVerificationRequest := *openapiclient.NewCreateDeviceVerificationRequest(int32(58959)) // CreateDeviceVerificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.CreateDeviceVerification(context.Background(), userId).ContentType(contentType).CreateDeviceVerificationRequest(createDeviceVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.CreateDeviceVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceVerification`: CreateDeviceVerification201Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.CreateDeviceVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **createDeviceVerificationRequest** | [**CreateDeviceVerificationRequest**](CreateDeviceVerificationRequest.md) |  | 

### Return type

[**CreateDeviceVerification201Response**](CreateDeviceVerification201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFactorRegistration

> CreateFactorRegistration201Response CreateFactorRegistration(ctx, userId).ContentType(contentType).CreateFactorRegistrationRequest(createFactorRegistrationRequest).Execute()

Create Factor Registration



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
    contentType := "application/json" // string |  (optional) (default to "application/json")
    createFactorRegistrationRequest := *openapiclient.NewCreateFactorRegistrationRequest(int32(58959), "OneLogin SMS") // CreateFactorRegistrationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.CreateFactorRegistration(context.Background(), userId).ContentType(contentType).CreateFactorRegistrationRequest(createFactorRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.CreateFactorRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFactorRegistration`: CreateFactorRegistration201Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.CreateFactorRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFactorRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **createFactorRegistrationRequest** | [**CreateFactorRegistrationRequest**](CreateFactorRegistrationRequest.md) |  | 

### Return type

[**CreateFactorRegistration201Response**](CreateFactorRegistration201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnrolledFactor

> DeleteEnrolledFactor(ctx, userId, deviceId).Execute()

Delete Enrolled Factor



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
    resp, r, err := apiClient.MultiFactorAuthenticationApi.DeleteEnrolledFactor(context.Background(), userId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.DeleteEnrolledFactor``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEnrolledFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOTP

> GenerateOTP201Response GenerateOTP(ctx, userId).ContentType(contentType).GenerateOTPRequest(generateOTPRequest).Execute()

Generate MFA token



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
    contentType := "application/json" // string |  (optional) (default to "application/json")
    generateOTPRequest := *openapiclient.NewGenerateOTPRequest() // GenerateOTPRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.GenerateOTP(context.Background(), userId).ContentType(contentType).GenerateOTPRequest(generateOTPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.GenerateOTP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOTP`: GenerateOTP201Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.GenerateOTP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **generateOTPRequest** | [**GenerateOTPRequest**](GenerateOTPRequest.md) |  | 

### Return type

[**GenerateOTP201Response**](GenerateOTP201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthFactors

> GetAuthFactors200Response GetAuthFactors(ctx, userId).Execute()

Get User Factors



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
    resp, r, err := apiClient.MultiFactorAuthenticationApi.GetAuthFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.GetAuthFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthFactors`: GetAuthFactors200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.GetAuthFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAuthFactors200Response**](GetAuthFactors200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationDevices

> []GetAuthenticationDevices200ResponseInner GetAuthenticationDevices(ctx, userId).Execute()

Get User Devices



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
    resp, r, err := apiClient.MultiFactorAuthenticationApi.GetAuthenticationDevices(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.GetAuthenticationDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationDevices`: []GetAuthenticationDevices200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.GetAuthenticationDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetAuthenticationDevices200ResponseInner**](GetAuthenticationDevices200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRegistration

> map[string]interface{} GetUserRegistration(ctx, userId, registrationId).Execute()

Get User Registration



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
    registrationId := "<UUID>" // string | The id of a registration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.GetUserRegistration(context.Background(), userId, registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.GetUserRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRegistration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.GetUserRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**registrationId** | **string** | The id of a registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserVerification

> GetUserVerification200Response GetUserVerification(ctx, userId, verificationId).Execute()

Get User Verification



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
    verificationId := "<UUID>" // string | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.GetUserVerification(context.Background(), userId, verificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.GetUserVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserVerification`: GetUserVerification200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.GetUserVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**verificationId** | **string** | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserVerification200Response**](GetUserVerification200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUserRegistration

> VerifyUserRegistration200Response VerifyUserRegistration(ctx, userId, registrationId).ContentType(contentType).VerifyUserRegistrationRequest(verifyUserRegistrationRequest).Execute()

Verify User Registration



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
    registrationId := "<UUID>" // string | The id of a registration
    contentType := "application/json" // string |  (optional) (default to "application/json")
    verifyUserRegistrationRequest := *openapiclient.NewVerifyUserRegistrationRequest() // VerifyUserRegistrationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.VerifyUserRegistration(context.Background(), userId, registrationId).ContentType(contentType).VerifyUserRegistrationRequest(verifyUserRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.VerifyUserRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyUserRegistration`: VerifyUserRegistration200Response
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.VerifyUserRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**registrationId** | **string** | The id of a registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **verifyUserRegistrationRequest** | [**VerifyUserRegistrationRequest**](VerifyUserRegistrationRequest.md) |  | 

### Return type

[**VerifyUserRegistration200Response**](VerifyUserRegistration200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUserVerification

> Error VerifyUserVerification(ctx, userId, verificationId).ContentType(contentType).VerifyUserVerificationRequest(verifyUserVerificationRequest).Execute()

Verify User Verification



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
    verificationId := "<UUID>" // string | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call.
    contentType := "application/json" // string |  (optional) (default to "application/json")
    verifyUserVerificationRequest := *openapiclient.NewVerifyUserVerificationRequest() // VerifyUserVerificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultiFactorAuthenticationApi.VerifyUserVerification(context.Background(), userId, verificationId).ContentType(contentType).VerifyUserVerificationRequest(verifyUserVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultiFactorAuthenticationApi.VerifyUserVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyUserVerification`: Error
    fmt.Fprintf(os.Stdout, "Response from `MultiFactorAuthenticationApi.VerifyUserVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 
**verificationId** | **string** | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **verifyUserVerificationRequest** | [**VerifyUserVerificationRequest**](VerifyUserVerificationRequest.md) |  | 

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

