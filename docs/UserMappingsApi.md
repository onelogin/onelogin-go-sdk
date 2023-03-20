# \UserMappingsApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMapping**](UserMappingsApi.md#CreateMapping) | **Post** /api/2/mappings | Create Mapping
[**DeleteMapping**](UserMappingsApi.md#DeleteMapping) | **Delete** /api/2/mappings/{mapping_id} | Delete Mapping
[**GetMapping**](UserMappingsApi.md#GetMapping) | **Get** /api/2/mappings/{mapping_id} | Get Mapping
[**ListMappingActionValues**](UserMappingsApi.md#ListMappingActionValues) | **Get** /api/2/mappings/actions/{mapping_action_value}/values | List Actions Values
[**ListMappingConditions**](UserMappingsApi.md#ListMappingConditions) | **Get** /api/2/mappings/conditions | List Conditions
[**ListMappingConditionsOperators**](UserMappingsApi.md#ListMappingConditionsOperators) | **Get** /api/2/mappings/conditions/{mapping_condition_value}/operators | List Conditions Operators
[**ListMappingContionValues**](UserMappingsApi.md#ListMappingContionValues) | **Get** /api/2/mappings/conditions/{mapping_condition_value}/values | List Conditions Values
[**ListMappings**](UserMappingsApi.md#ListMappings) | **Get** /api/2/mappings | List Mappings
[**ListMappingsActions**](UserMappingsApi.md#ListMappingsActions) | **Get** /api/2/mappings/actions | List Actions
[**SortMappings**](UserMappingsApi.md#SortMappings) | **Put** /api/2/mappings/sort | Bulk Sort
[**UpdateMapping**](UserMappingsApi.md#UpdateMapping) | **Put** /api/2/mappings/{mapping_id} | Update Mapping



## CreateMapping

> []Mapping CreateMapping(ctx).ContentType(contentType).Mapping(mapping).Execute()

Create Mapping



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
    mapping := *openapiclient.NewMapping("Name_example", false, "Match_example", int32(123), []openapiclient.Condition{*openapiclient.NewCondition()}, []openapiclient.ActionObj{*openapiclient.NewActionObj()}) // Mapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.CreateMapping(context.Background()).ContentType(contentType).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.CreateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMapping`: []Mapping
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.CreateMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **mapping** | [**Mapping**](Mapping.md) |  | 

### Return type

[**[]Mapping**](Mapping.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMapping

> DeleteMapping(ctx, mappingId).Execute()

Delete Mapping



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
    mappingId := int32(56) // int32 | The id of the user mapping to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.DeleteMapping(context.Background(), mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.DeleteMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **int32** | The id of the user mapping to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMappingRequest struct via the builder pattern


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


## GetMapping

> Mapping GetMapping(ctx, mappingId).Execute()

Get Mapping



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
    mappingId := int32(56) // int32 | The id of the user mapping to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.GetMapping(context.Background(), mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: Mapping
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.GetMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **int32** | The id of the user mapping to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mapping**](Mapping.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingActionValues

> []ListMappingActionValues200ResponseInner ListMappingActionValues(ctx, mappingActionValue).Execute()

List Actions Values



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
    mappingActionValue := "mappingActionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.ListMappingActionValues(context.Background(), mappingActionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappingActionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingActionValues`: []ListMappingActionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappingActionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingActionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingActionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMappingActionValues200ResponseInner**](ListMappingActionValues200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingConditions

> ListMappingConditions200Response ListMappingConditions(ctx).Execute()

List Conditions



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
    resp, r, err := apiClient.UserMappingsApi.ListMappingConditions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappingConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingConditions`: ListMappingConditions200Response
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappingConditions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingConditionsRequest struct via the builder pattern


### Return type

[**ListMappingConditions200Response**](ListMappingConditions200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingConditionsOperators

> []ListMappingConditionsOperators200ResponseInner ListMappingConditionsOperators(ctx, mappingConditionValue).Execute()

List Conditions Operators



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
    mappingConditionValue := "mappingConditionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.ListMappingConditionsOperators(context.Background(), mappingConditionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappingConditionsOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingConditionsOperators`: []ListMappingConditionsOperators200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappingConditionsOperators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingConditionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingConditionsOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMappingConditionsOperators200ResponseInner**](ListMappingConditionsOperators200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingContionValues

> []ListMappingContionValues200ResponseInner ListMappingContionValues(ctx, mappingConditionValue).Execute()

List Conditions Values



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
    mappingConditionValue := "mappingConditionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.ListMappingContionValues(context.Background(), mappingConditionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappingContionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingContionValues`: []ListMappingContionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappingContionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingConditionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingContionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMappingContionValues200ResponseInner**](ListMappingContionValues200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappings

> []Mapping ListMappings(ctx).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()

List Mappings



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
    enabled := true // bool | Defaults to true. When set to `false` will return all disabled mappings. (optional) (default to true)
    hasCondition := "has_condition=has_role:123456" // string | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition=has_role:123456 Multiple filters. has_condition=has_role:123456,status:1 Wildcard for conditions. has_condition=*:123456 Wildcard for condition values. has_condition=has_role:* (optional)
    hasConditionType := "hasConditionType_example" // string | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition=has_role:123456 Multiple filters. has_condition=has_role:123456,status:1 Wildcard for conditions. has_condition=*:123456 Wildcard for condition values. has_condition=has_role:* (optional)
    hasAction := "has_action=set_groups:123456,set_usertype:*" // string | Filters Rules based on their Actions. Values formatted as :, where name is the Action to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_action=set_licenses:123456 Multiple filters. has_action=set_groups:123456,set_usertype:* Wildcard for actions. has_action=*:123456 Wildcard for action values. has_action=set_userprincipalname:* (optional)
    hasActionType := "hasActionType_example" // string | Filters Rules based on their action types. Allowed values are: builtin - actions that involve standard attributes custom - actions that involve custom attributes none - no actions are defined For example: Find Rules with no actions has_action_type=none (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.ListMappings(context.Background()).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappings`: []Mapping
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | Defaults to true. When set to &#x60;false&#x60; will return all disabled mappings. | [default to true]
 **hasCondition** | **string** | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition&#x3D;has_role:123456 Multiple filters. has_condition&#x3D;has_role:123456,status:1 Wildcard for conditions. has_condition&#x3D;*:123456 Wildcard for condition values. has_condition&#x3D;has_role:* | 
 **hasConditionType** | **string** | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition&#x3D;has_role:123456 Multiple filters. has_condition&#x3D;has_role:123456,status:1 Wildcard for conditions. has_condition&#x3D;*:123456 Wildcard for condition values. has_condition&#x3D;has_role:* | 
 **hasAction** | **string** | Filters Rules based on their Actions. Values formatted as :, where name is the Action to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_action&#x3D;set_licenses:123456 Multiple filters. has_action&#x3D;set_groups:123456,set_usertype:* Wildcard for actions. has_action&#x3D;*:123456 Wildcard for action values. has_action&#x3D;set_userprincipalname:* | 
 **hasActionType** | **string** | Filters Rules based on their action types. Allowed values are: builtin - actions that involve standard attributes custom - actions that involve custom attributes none - no actions are defined For example: Find Rules with no actions has_action_type&#x3D;none | 

### Return type

[**[]Mapping**](Mapping.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingsActions

> []ListMappingsActions200ResponseInner ListMappingsActions(ctx).Execute()

List Actions



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
    resp, r, err := apiClient.UserMappingsApi.ListMappingsActions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.ListMappingsActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingsActions`: []ListMappingsActions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.ListMappingsActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsActionsRequest struct via the builder pattern


### Return type

[**[]ListMappingsActions200ResponseInner**](ListMappingsActions200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SortMappings

> []int32 SortMappings(ctx).RequestBody(requestBody).Execute()

Bulk Sort



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
    requestBody := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.SortMappings(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.SortMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SortMappings`: []int32
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.SortMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSortMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** |  | 

### Return type

**[]int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapping

> int32 UpdateMapping(ctx, mappingId).ContentType(contentType).Body(body).Execute()

Update Mapping



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
    mappingId := int32(56) // int32 | The id of the user mapping to locate.
    contentType := "application/json" // string |  (optional) (default to "application/json")
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingsApi.UpdateMapping(context.Background(), mappingId).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingsApi.UpdateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapping`: int32
    fmt.Fprintf(os.Stdout, "Response from `UserMappingsApi.UpdateMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **int32** | The id of the user mapping to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | [default to &quot;application/json&quot;]
 **body** | **map[string]interface{}** |  | 

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

