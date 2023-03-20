# \BrandingServiceApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrand**](BrandingServiceApi.md#CreateBrand) | **Post** /api/2/branding/brands | Create Brand
[**DeleteBrand**](BrandingServiceApi.md#DeleteBrand) | **Delete** /api/2/branding/brands/{brand_id} | Delete Brand
[**GetBrand**](BrandingServiceApi.md#GetBrand) | **Get** /api/2/branding/brands/{brand_id} | Get Brand
[**GetBrandApps**](BrandingServiceApi.md#GetBrandApps) | **Get** /api/2/branding/brands/{brand_id}/apps | Get Brand Apps
[**ListBrands**](BrandingServiceApi.md#ListBrands) | **Get** /api/2/branding/brands | List Account Brands
[**UpdateBrand**](BrandingServiceApi.md#UpdateBrand) | **Put** /api/2/branding/brands/{brand_id} | Update Brand



## CreateBrand

> Brand CreateBrand(ctx).Brand(brand).Execute()

Create Brand



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
    brand := *openapiclient.NewBrand(int32(123), true, false, "#1298b4", "#b60012", "#beefed", int32(40), "You must register with the OneLogin Protect app in order to login", true, "ACME Username or Email", "To login, enter your ACME Username or Email. Reach out to help.desk@acme.org if you have trouble logging in.", "ACME Login Instructions", true, *openapiclient.NewBrandBackground(*openapiclient.NewBrandBackgroundUrls("Original_example", "Login_example", "Branding_example"), int32(123), "UpdatedAt_example", "ContentType_example"), *openapiclient.NewBrandLogo(*openapiclient.NewBrandLogoUrls("Original_example", "Login_example", "Navigation_example"), int32(123), "UpdatedAt_example", "ContentType_example")) // Brand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceApi.CreateBrand(context.Background()).Brand(brand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.CreateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceApi.CreateBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brand** | [**Brand**](Brand.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrand

> DeleteBrand(ctx, brandId).Execute()

Delete Brand



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceApi.DeleteBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.DeleteBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


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


## GetBrand

> Brand GetBrand(ctx, brandId).Execute()

Get Brand



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceApi.GetBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceApi.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Brand**](Brand.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandApps

> []BrandApp GetBrandApps(ctx, brandId).Execute()

Get Brand Apps



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceApi.GetBrandApps(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.GetBrandApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandApps`: []BrandApp
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceApi.GetBrandApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BrandApp**](BrandApp.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrands

> []BrandReq ListBrands(ctx).Execute()

List Account Brands



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
    resp, r, err := apiClient.BrandingServiceApi.ListBrands(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.ListBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrands`: []BrandReq
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceApi.ListBrands`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandsRequest struct via the builder pattern


### Return type

[**[]BrandReq**](BrandReq.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrand

> Brand UpdateBrand(ctx, brandId).RequestBrand(requestBrand).Execute()

Update Brand



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
    requestBrand := *openapiclient.NewRequestBrand("Acme Branding") // RequestBrand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingServiceApi.UpdateBrand(context.Background(), brandId).RequestBrand(requestBrand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingServiceApi.UpdateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `BrandingServiceApi.UpdateBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **int32** | Unique identifier for the branding object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBrand** | [**RequestBrand**](RequestBrand.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

