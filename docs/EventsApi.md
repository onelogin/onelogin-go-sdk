# \EventsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventById**](EventsApi.md#GetEventById) | **Get** /api/1/events/{event_id} | Get Event by ID
[**GetEventTypes**](EventsApi.md#GetEventTypes) | **Get** /api/1/events/types | Get Event Types
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /api/1/events | Get Events



## GetEventById

> GetEventById200Response GetEventById(ctx, eventId).Execute()

Get Event by ID



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
    eventId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventById(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventById`: GetEventById200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventById200Response**](GetEventById200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> GetEventTypes200Response GetEventTypes(ctx).ContentType(contentType).Execute()

Get Event Types



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventTypes(context.Background()).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventTypes`: GetEventTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]

### Return type

[**GetEventTypes200Response**](GetEventTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> GetEvents200Response GetEvents(ctx).EventTypeId(eventTypeId).ClientId(clientId).DirectoryId(directoryId).Id(id).CreatedAt(createdAt).Resolution(resolution).Since(since).Until(until).UserId(userId).Execute()

Get Events



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
    eventTypeId := []int32{int32(123)} // []int32 |  (optional)
    clientId := int32(56) // int32 |  (optional)
    directoryId := int32(56) // int32 |  (optional)
    id := int32(56) // int32 |  (optional)
    createdAt := "createdAt_example" // string |  (optional)
    resolution := "resolution_example" // string |  (optional)
    since := "since_example" // string |  (optional)
    until := "until_example" // string |  (optional)
    userId := int32(56) // int32 | Set to the id of the user that you want to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvents(context.Background()).EventTypeId(eventTypeId).ClientId(clientId).DirectoryId(directoryId).Id(id).CreatedAt(createdAt).Resolution(resolution).Since(since).Until(until).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: GetEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventTypeId** | **[]int32** |  | 
 **clientId** | **int32** |  | 
 **directoryId** | **int32** |  | 
 **id** | **int32** |  | 
 **createdAt** | **string** |  | 
 **resolution** | **string** |  | 
 **since** | **string** |  | 
 **until** | **string** |  | 
 **userId** | **int32** | Set to the id of the user that you want to return. | 

### Return type

[**GetEvents200Response**](GetEvents200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

