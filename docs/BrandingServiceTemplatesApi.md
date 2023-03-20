# \BrandingServiceTemplatesApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageTemplate**](BrandingServiceTemplatesApi.md#CreateMessageTemplate) | **Post** /api/2/branding/brands/{brand_id}/templates | Create Message Template
[**DeleteMessageTemplate**](BrandingServiceTemplatesApi.md#DeleteMessageTemplate) | **Delete** /api/2/branding/brands/{brand_id}/templates/{template_id} | Delete Message Template
[**GetMasterByType**](BrandingServiceTemplatesApi.md#GetMasterByType) | **Get** /api/2/branding/brands/master/templates/{template_type} | Get Master Template by Type
[**GetMessageTemplateById**](BrandingServiceTemplatesApi.md#GetMessageTemplateById) | **Get** /api/2/branding/brands/{brand_id}/templates/{template_id} | Get Message Template
[**GetTemplateByLocale**](BrandingServiceTemplatesApi.md#GetTemplateByLocale) | **Get** /api/2/branding/brands/{brand_id}/templates/{template_type}/{locale} | Get Template by Type &amp; Locale
[**ListMessageTemplates**](BrandingServiceTemplatesApi.md#ListMessageTemplates) | **Get** /api/2/branding/brands/{brand_id}/templates | List Message Templates
[**UpdateMessageTemplateById**](BrandingServiceTemplatesApi.md#UpdateMessageTemplateById) | **Put** /api/2/branding/brands/{brand_id}/templates/{template_id} | Update Message Template
[**UpdateTemplateByLocale**](BrandingServiceTemplatesApi.md#UpdateTemplateByLocale) | **Put** /api/2/branding/brands/{brand_id}/templates/{template_type}/{locale} | Update Template by Type &amp; Locale



## CreateMessageTemplate

> MessageTemplate CreateMessageTemplate(ctx, brandId).Locale(locale).MessageTemplate(messageTemplate).Execute()

Create Message Template



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    locale := "en" // string | The 2 character language locale for the template. e.g. en = English, es = Spanish (optional)
    messageTemplate := *openapiclient.NewMessageTemplate("Type_example", "en", openapiclient.message_template_template{MessageTemplateTemplateOneOf: openapiclient.NewMessageTemplateTemplateOneOf("Email MFA App Verification Code", "<html><head></head><body><p>Here is the code: {{otp_code}}</p></body></html>", "Here is the code: {{otp_code}}")}) // MessageTemplate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.CreateMessageTemplate(context.Background(), brandId).Locale(locale).MessageTemplate(messageTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.CreateMessageTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessageTemplate`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.CreateMessageTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish | 
 **messageTemplate** | [**MessageTemplate**](MessageTemplate.md) |  | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessageTemplate

> DeleteMessageTemplate(ctx, brandId, templateId).Execute()

Delete Message Template



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    templateId := int32(25) // int32 | Unique identifier for the template to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.DeleteMessageTemplate(context.Background(), brandId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.DeleteMessageTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 
**templateId** | **int32** | Unique identifier for the template to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageTemplateRequest struct via the builder pattern


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


## GetMasterByType

> MessageTemplate GetMasterByType(ctx, templateType).Execute()

Get Master Template by Type



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
    templateType := "email_template" // string | The message template type to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.GetMasterByType(context.Background(), templateType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.GetMasterByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMasterByType`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.GetMasterByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateType** | **string** | The message template type to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageTemplateById

> MessageTemplate GetMessageTemplateById(ctx, brandId, templateId).Execute()

Get Message Template



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    templateId := int32(25) // int32 | Unique identifier for the template to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.GetMessageTemplateById(context.Background(), brandId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.GetMessageTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageTemplateById`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.GetMessageTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 
**templateId** | **int32** | Unique identifier for the template to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateByLocale

> MessageTemplate GetTemplateByLocale(ctx, brandId, templateType, locale).Execute()

Get Template by Type & Locale



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    templateType := "email_template" // string | The message template type to return.
    locale := "en" // string | The 2 character language locale for the template. e.g. en = English, es = Spanish

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.GetTemplateByLocale(context.Background(), brandId, templateType, locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.GetTemplateByLocale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateByLocale`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.GetTemplateByLocale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 
**templateType** | **string** | The message template type to return. | 
**locale** | **string** | The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateByLocaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageTemplates

> []ListMessageTemplates200ResponseInner ListMessageTemplates(ctx, brandId).Locale(locale).Execute()

List Message Templates



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    locale := "en" // string | The 2 character language locale for the template. e.g. en = English, es = Spanish (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.ListMessageTemplates(context.Background(), brandId).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.ListMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessageTemplates`: []ListMessageTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.ListMessageTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMessageTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish | 

### Return type

[**[]ListMessageTemplates200ResponseInner**](ListMessageTemplates200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageTemplateById

> MessageTemplate UpdateMessageTemplateById(ctx, brandId, templateId).Execute()

Update Message Template



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    templateId := int32(25) // int32 | Unique identifier for the template to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.UpdateMessageTemplateById(context.Background(), brandId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.UpdateMessageTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessageTemplateById`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.UpdateMessageTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 
**templateId** | **int32** | Unique identifier for the template to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplateByLocale

> MessageTemplate UpdateTemplateByLocale(ctx, brandId, templateType, locale).Execute()

Update Template by Type & Locale



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
    brandId := int32(9) // int32 | Unique identifier for the branding object.
    templateType := "email_template" // string | The message template type to return.
    locale := "en" // string | The 2 character language locale for the template. e.g. en = English, es = Spanish

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceTemplatesApi.UpdateTemplateByLocale(context.Background(), brandId, templateType, locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceTemplatesApi.UpdateTemplateByLocale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplateByLocale`: MessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceTemplatesApi.UpdateTemplateByLocale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 
**templateType** | **string** | The message template type to return. | 
**locale** | **string** | The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateByLocaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

