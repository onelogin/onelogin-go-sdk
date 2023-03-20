# \SAMLAssertionsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateSamlAssert**](SAMLAssertionsApi.md#GenerateSamlAssert) | **Post** /api/1/saml_assertion | Generate SAML Assertion
[**GenerateSamlAssert2**](SAMLAssertionsApi.md#GenerateSamlAssert2) | **Post** /api/2/saml_assertion | Generate SAML Assertion
[**VerFactorSaml**](SAMLAssertionsApi.md#VerFactorSaml) | **Post** /api/1/saml_assertion/verify_factor | Verify Factor SAML
[**VerFactorSaml2**](SAMLAssertionsApi.md#VerFactorSaml2) | **Post** /api/2/saml_assertion/verify_factor | Verify Factor SAML



## GenerateSamlAssert

> GenerateSamlAssert200Response GenerateSamlAssert(ctx).ContentType(contentType).SamlAssert(samlAssert).Execute()

Generate SAML Assertion



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
    samlAssert := *openapiclient.NewSamlAssert("UsernameOrEmail_example", "Password_example", "AppId_example", "Subdomain_example") // SamlAssert |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SAMLAssertionsApi.GenerateSamlAssert(context.Background()).ContentType(contentType).SamlAssert(samlAssert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAssertionsApi.GenerateSamlAssert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSamlAssert`: GenerateSamlAssert200Response
    fmt.Fprintf(os.Stdout, "Response from `SAMLAssertionsApi.GenerateSamlAssert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSamlAssertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **samlAssert** | [**SamlAssert**](SamlAssert.md) |  | 

### Return type

[**GenerateSamlAssert200Response**](GenerateSamlAssert200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSamlAssert2

> GenerateSamlAssert200Response GenerateSamlAssert2(ctx).ContentType(contentType).SamlAssert(samlAssert).Execute()

Generate SAML Assertion



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
    samlAssert := *openapiclient.NewSamlAssert("UsernameOrEmail_example", "Password_example", "AppId_example", "Subdomain_example") // SamlAssert |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SAMLAssertionsApi.GenerateSamlAssert2(context.Background()).ContentType(contentType).SamlAssert(samlAssert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAssertionsApi.GenerateSamlAssert2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSamlAssert2`: GenerateSamlAssert200Response
    fmt.Fprintf(os.Stdout, "Response from `SAMLAssertionsApi.GenerateSamlAssert2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSamlAssert2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **samlAssert** | [**SamlAssert**](SamlAssert.md) |  | 

### Return type

[**GenerateSamlAssert200Response**](GenerateSamlAssert200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerFactorSaml

> VerFactorSaml200Response VerFactorSaml(ctx).ContentType(contentType).SamlFactor(samlFactor).Execute()

Verify Factor SAML



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
    samlFactor := *openapiclient.NewSamlFactor("1657651", "1657651", "11x0x1x16x1x3259xxxx0x59xx6xxxx670x61x45xxxxx") // SamlFactor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SAMLAssertionsApi.VerFactorSaml(context.Background()).ContentType(contentType).SamlFactor(samlFactor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAssertionsApi.VerFactorSaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerFactorSaml`: VerFactorSaml200Response
    fmt.Fprintf(os.Stdout, "Response from `SAMLAssertionsApi.VerFactorSaml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerFactorSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **samlFactor** | [**SamlFactor**](SamlFactor.md) |  | 

### Return type

[**VerFactorSaml200Response**](VerFactorSaml200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerFactorSaml2

> VerFactorSaml200Response VerFactorSaml2(ctx).ContentType(contentType).SamlFactor(samlFactor).Execute()

Verify Factor SAML



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
    samlFactor := *openapiclient.NewSamlFactor("1657651", "1657651", "11x0x1x16x1x3259xxxx0x59xx6xxxx670x61x45xxxxx") // SamlFactor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SAMLAssertionsApi.VerFactorSaml2(context.Background()).ContentType(contentType).SamlFactor(samlFactor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAssertionsApi.VerFactorSaml2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerFactorSaml2`: VerFactorSaml200Response
    fmt.Fprintf(os.Stdout, "Response from `SAMLAssertionsApi.VerFactorSaml2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerFactorSaml2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **samlFactor** | [**SamlFactor**](SamlFactor.md) |  | 

### Return type

[**VerFactorSaml200Response**](VerFactorSaml200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

