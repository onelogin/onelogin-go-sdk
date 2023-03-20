# \BrandingServiceSMTPApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmailSettings**](BrandingServiceSMTPApi.md#DeleteEmailSettings) | **Delete** /api/2/branding/email_settings | Delete Custom Email Settings
[**GetEmailSettings**](BrandingServiceSMTPApi.md#GetEmailSettings) | **Get** /api/2/branding/email_settings | Get Email Settings
[**UpdateEmailSettings**](BrandingServiceSMTPApi.md#UpdateEmailSettings) | **Put** /api/2/branding/email_settings | Update Email Settings



## DeleteEmailSettings

> AltErr DeleteEmailSettings(ctx).Execute()

Delete Custom Email Settings



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
    resp, r, err := apiClient.BrandingServiceSMTPApi.DeleteEmailSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceSMTPApi.DeleteEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEmailSettings`: AltErr
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceSMTPApi.DeleteEmailSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailSettingsRequest struct via the builder pattern


### Return type

[**AltErr**](AltErr.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSettings

> GetEmailSettings200Response GetEmailSettings(ctx).Execute()

Get Email Settings



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
    resp, r, err := apiClient.BrandingServiceSMTPApi.GetEmailSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceSMTPApi.GetEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSettings`: GetEmailSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceSMTPApi.GetEmailSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSettingsRequest struct via the builder pattern


### Return type

[**GetEmailSettings200Response**](GetEmailSettings200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailSettings

> AltErr UpdateEmailSettings(ctx).EmailConfig(emailConfig).Execute()

Update Email Settings



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
    emailConfig := *openapiclient.NewEmailConfig("smtp.sendgrid.net", "email@example.com", "example.com") // EmailConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceSMTPApi.UpdateEmailSettings(context.Background()).EmailConfig(emailConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceSMTPApi.UpdateEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailSettings`: AltErr
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceSMTPApi.UpdateEmailSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailConfig** | [**EmailConfig**](EmailConfig.md) |  | 

### Return type

[**AltErr**](AltErr.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

