# \VigilanceAIApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskRule**](VigilanceAIApi.md#CreateRiskRule) | **Post** /api/2/risk/rules | Create Rule
[**DeleteRiskRule**](VigilanceAIApi.md#DeleteRiskRule) | **Delete** /api/2/risk/rules/{rule_id} | Delete Rule
[**GetRiskRule**](VigilanceAIApi.md#GetRiskRule) | **Get** /api/2/risk/rules/{rule_id} | get Risk Rule
[**GetRiskScore**](VigilanceAIApi.md#GetRiskScore) | **Post** /api/2/risk/verify | Get a Risk Score
[**GetRiskScores**](VigilanceAIApi.md#GetRiskScores) | **Get** /api/2/risk/scores | Get Score Summary
[**ListRiskRules**](VigilanceAIApi.md#ListRiskRules) | **Get** /api/2/risk/rules | List Rules
[**TrackRiskEvent**](VigilanceAIApi.md#TrackRiskEvent) | **Post** /api/2/risk/events | Track an Event
[**UpdateRiskRule**](VigilanceAIApi.md#UpdateRiskRule) | **Put** /api/2/risk/rules/{rule_id} | Update Rule



## CreateRiskRule

> RiskRule CreateRiskRule(ctx).RiskRule(riskRule).Execute()

Create Rule



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
    riskRule := *openapiclient.NewRiskRule() // RiskRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.CreateRiskRule(context.Background()).RiskRule(riskRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.CreateRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskRule`: RiskRule
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.CreateRiskRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskRule** | [**RiskRule**](RiskRule.md) |  | 

### Return type

[**RiskRule**](RiskRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskRule

> DeleteRiskRule(ctx, ruleId).Execute()

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
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.DeleteRiskRule(context.Background(), ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.DeleteRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskRuleRequest struct via the builder pattern


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


## GetRiskRule

> RiskRule GetRiskRule(ctx, ruleId).Execute()

get Risk Rule



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
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.GetRiskRule(context.Background(), ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.GetRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskRule`: RiskRule
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.GetRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskRule**](RiskRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskScore

> GetRiskScore200Response GetRiskScore(ctx).GetRiskScoreRequest(getRiskScoreRequest).Before(before).After(after).Execute()

Get a Risk Score



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
    getRiskScoreRequest := *openapiclient.NewGetRiskScoreRequest("Ip_example", "UserAgent_example", *openapiclient.NewRiskUser("Id_example")) // GetRiskScoreRequest | 
    before := "before_example" // string | Optional ISO8601 formatted date string. Defaults to current date. Maximum date is 90 days ago. (optional)
    after := "after_example" // string | Optional ISO8601 formatted date string. Defaults to 30 days ago. Maximum date is 90 days ago. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.GetRiskScore(context.Background()).GetRiskScoreRequest(getRiskScoreRequest).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.GetRiskScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskScore`: GetRiskScore200Response
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.GetRiskScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getRiskScoreRequest** | [**GetRiskScoreRequest**](GetRiskScoreRequest.md) |  | 
 **before** | **string** | Optional ISO8601 formatted date string. Defaults to current date. Maximum date is 90 days ago. | 
 **after** | **string** | Optional ISO8601 formatted date string. Defaults to 30 days ago. Maximum date is 90 days ago. | 

### Return type

[**GetRiskScore200Response**](GetRiskScore200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskScores

> GetRiskScores200Response GetRiskScores(ctx).Execute()

Get Score Summary



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
    resp, r, err := apiClient.VigilanceAIApi.GetRiskScores(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.GetRiskScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskScores`: GetRiskScores200Response
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.GetRiskScores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskScoresRequest struct via the builder pattern


### Return type

[**GetRiskScores200Response**](GetRiskScores200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRiskRules

> []RiskRule ListRiskRules(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.ListRiskRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.ListRiskRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRiskRules`: []RiskRule
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.ListRiskRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRiskRulesRequest struct via the builder pattern


### Return type

[**[]RiskRule**](RiskRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackRiskEvent

> TrackRiskEvent(ctx).TrackRiskEventRequest(trackRiskEventRequest).Execute()

Track an Event



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
    trackRiskEventRequest := *openapiclient.NewTrackRiskEventRequest("Verb_example", "Ip_example", "UserAgent_example", *openapiclient.NewRiskUser("Id_example")) // TrackRiskEventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.TrackRiskEvent(context.Background()).TrackRiskEventRequest(trackRiskEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.TrackRiskEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackRiskEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackRiskEventRequest** | [**TrackRiskEventRequest**](TrackRiskEventRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskRule

> RiskRule UpdateRiskRule(ctx, ruleId).UpdateRiskRuleRequest(updateRiskRuleRequest).Execute()

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
    ruleId := "ruleId_example" // string | 
    updateRiskRuleRequest := *openapiclient.NewUpdateRiskRuleRequest() // UpdateRiskRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VigilanceAIApi.UpdateRiskRule(context.Background(), ruleId).UpdateRiskRuleRequest(updateRiskRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VigilanceAIApi.UpdateRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskRule`: RiskRule
    fmt.Fprintf(os.Stdout, "Response from `VigilanceAIApi.UpdateRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRiskRuleRequest** | [**UpdateRiskRuleRequest**](UpdateRiskRuleRequest.md) |  | 

### Return type

[**RiskRule**](RiskRule.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

