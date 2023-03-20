# \AppRulesApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppRule**](AppRulesApi.md#CreateAppRule) | **Post** /api/2/apps/{app_id}/rules | 
[**DeleteRule**](AppRulesApi.md#DeleteRule) | **Delete** /api/2/apps/{app_id}/rules/{rule_id} | Delete Rule
[**GetAppRule**](AppRulesApi.md#GetAppRule) | **Get** /api/2/apps/{app_id}/rules/{rule_id} | Get Rule
[**ListActionValies**](AppRulesApi.md#ListActionValies) | **Get** /api/2/apps/{app_id}/rules/actions/{rule_action_value}/values | List Actions Values
[**ListActions**](AppRulesApi.md#ListActions) | **Get** /api/2/apps/{app_id}/rules/actions | List Actions
[**ListAppRules**](AppRulesApi.md#ListAppRules) | **Get** /api/2/apps/{app_id}/rules | List Rules
[**ListConditionOperators**](AppRulesApi.md#ListConditionOperators) | **Get** /api/2/apps/{app_id}/rules/conditions/{rule_condition_value}/operators | List Conditions Operators
[**ListConditionValues**](AppRulesApi.md#ListConditionValues) | **Get** /api/2/apps/{app_id}/rules/conditions/{rule_condition_value}/values | List Conditions Values
[**ListConditions**](AppRulesApi.md#ListConditions) | **Get** /api/2/apps/{app_id}/rules/conditions | List Conditions
[**SortAppRules**](AppRulesApi.md#SortAppRules) | **Put** /api/2/apps/{app_id}/rules/sort | Bulk Sort
[**UpdateAppRule**](AppRulesApi.md#UpdateAppRule) | **Put** /api/2/apps/{app_id}/rules/{rule_id} | Update Rule



## CreateAppRule

> AppRule CreateAppRule(ctx, appId).AppRule(appRule).Execute()





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
    appId := int32(56) // int32 | 
    appRule := *openapiclient.NewAppRule() // AppRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.CreateAppRule(context.Background(), appId).AppRule(appRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.CreateAppRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppRule`: AppRule
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.CreateAppRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appRule** | [**AppRule**](AppRule.md) |  | 

### Return type

[**AppRule**](AppRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, appId, ruleId).Execute()

Delete Rule



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
    appId := int32(56) // int32 | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.DeleteRule(context.Background(), appId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetAppRule

> AppRule GetAppRule(ctx, appId, ruleId).Execute()

Get Rule



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
    appId := int32(56) // int32 | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.GetAppRule(context.Background(), appId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.GetAppRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppRule`: AppRule
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.GetAppRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppRule**](AppRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionValies

> []RuleAction ListActionValies(ctx, appId, ruleActionValue).Execute()

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
    appId := int32(56) // int32 | 
    ruleActionValue := "ruleActionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListActionValies(context.Background(), appId, ruleActionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListActionValies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActionValies`: []RuleAction
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListActionValies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleActionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActionValiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleAction**](RuleAction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActions

> []RuleAction ListActions(ctx, appId).Execute()

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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListActions(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActions`: []RuleAction
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RuleAction**](RuleAction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppRules

> []AppRule ListAppRules(ctx, appId).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Enabled(enabled).Execute()

List Rules



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
    appId := int32(56) // int32 | 
    hasCondition := "has_condition=has_role:123456" // string | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition=has_role:123456 Multiple filters. has_condition=has_role:123456,status:1 Wildcard for conditions. has_condition=*:123456 Wildcard for condition values. has_condition=has_role:* (optional)
    hasConditionType := "hasConditionType_example" // string | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition=has_role:123456 Multiple filters. has_condition=has_role:123456,status:1 Wildcard for conditions. has_condition=*:123456 Wildcard for condition values. has_condition=has_role:* (optional)
    hasAction := "has_action=set_groups:123456,set_usertype:*" // string | Filters Rules based on their Actions. Values formatted as :, where name is the Action to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_action=set_licenses:123456 Multiple filters. has_action=set_groups:123456,set_usertype:* Wildcard for actions. has_action=*:123456 Wildcard for action values. has_action=set_userprincipalname:* (optional)
    hasActionType := "hasActionType_example" // string | Filters Rules based on their action types. Allowed values are: builtin - actions that involve standard attributes custom - actions that involve custom attributes none - no actions are defined For example: Find Rules with no actions has_action_type=none (optional)
    enabled := true // bool | Defaults to true. When set to `false` will return all disabled mappings. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListAppRules(context.Background(), appId).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListAppRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppRules`: []AppRule
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListAppRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hasCondition** | **string** | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition&#x3D;has_role:123456 Multiple filters. has_condition&#x3D;has_role:123456,status:1 Wildcard for conditions. has_condition&#x3D;*:123456 Wildcard for condition values. has_condition&#x3D;has_role:* | 
 **hasConditionType** | **string** | Filters Rules based on their Conditions. Values formatted as :, where name is the Condition to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_condition&#x3D;has_role:123456 Multiple filters. has_condition&#x3D;has_role:123456,status:1 Wildcard for conditions. has_condition&#x3D;*:123456 Wildcard for condition values. has_condition&#x3D;has_role:* | 
 **hasAction** | **string** | Filters Rules based on their Actions. Values formatted as :, where name is the Action to look for, and value is the value to find. Multiple filters can be declared by using a comma delimited list. Wildcards are supported in both the name and value fields. For example: Single filter. has_action&#x3D;set_licenses:123456 Multiple filters. has_action&#x3D;set_groups:123456,set_usertype:* Wildcard for actions. has_action&#x3D;*:123456 Wildcard for action values. has_action&#x3D;set_userprincipalname:* | 
 **hasActionType** | **string** | Filters Rules based on their action types. Allowed values are: builtin - actions that involve standard attributes custom - actions that involve custom attributes none - no actions are defined For example: Find Rules with no actions has_action_type&#x3D;none | 
 **enabled** | **bool** | Defaults to true. When set to &#x60;false&#x60; will return all disabled mappings. | [default to true]

### Return type

[**[]AppRule**](AppRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditionOperators

> []RuleCondition ListConditionOperators(ctx, appId, ruleConditionValue).Execute()

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
    appId := int32(56) // int32 | 
    ruleConditionValue := "ruleConditionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListConditionOperators(context.Background(), appId, ruleConditionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListConditionOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditionOperators`: []RuleCondition
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListConditionOperators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleConditionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleCondition**](RuleCondition.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditionValues

> RuleCondition ListConditionValues(ctx, appId, ruleConditionValue).Execute()

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
    appId := int32(56) // int32 | 
    ruleConditionValue := "ruleConditionValue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListConditionValues(context.Background(), appId, ruleConditionValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListConditionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditionValues`: RuleCondition
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListConditionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleConditionValue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RuleCondition**](RuleCondition.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditions

> []ListConditions200ResponseInner ListConditions(ctx, appId).Execute()

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
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.ListConditions(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.ListConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditions`: []ListConditions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.ListConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListConditions200ResponseInner**](ListConditions200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SortAppRules

> []int32 SortAppRules(ctx, appId).RequestBody(requestBody).Execute()

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
    appId := int32(56) // int32 | 
    requestBody := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.SortAppRules(context.Background(), appId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.SortAppRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SortAppRules`: []int32
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.SortAppRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSortAppRulesRequest struct via the builder pattern


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


## UpdateAppRule

> AppRule UpdateAppRule(ctx, appId, ruleId).AppRule(appRule).Execute()

Update Rule



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
    appId := int32(56) // int32 | 
    ruleId := "ruleId_example" // string | 
    appRule := *openapiclient.NewAppRule() // AppRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppRulesApi.UpdateAppRule(context.Background(), appId, ruleId).AppRule(appRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppRulesApi.UpdateAppRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppRule`: AppRule
    fmt.Fprintf(os.Stdout, "Response from `AppRulesApi.UpdateAppRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appRule** | [**AppRule**](AppRule.md) |  | 

### Return type

[**AppRule**](AppRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

