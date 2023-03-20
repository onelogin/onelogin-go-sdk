# \GroupsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupById**](GroupsApi.md#GetGroupById) | **Get** /api/1/groups/{group_id} | Get Group by ID
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /api/1/groups | Get Groups



## GetGroupById

> GetGroups200Response GetGroupById(ctx, groupId).Execute()

Get Group by ID



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupById(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupById`: GetGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGroups200Response**](GetGroups200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GetGroups200Response GetGroups(ctx).Execute()

Get Groups



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
    resp, r, err := apiClient.GroupsApi.GetGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: GetGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


### Return type

[**GetGroups200Response**](GetGroups200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

