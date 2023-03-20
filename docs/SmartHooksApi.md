# \SmartHooksApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentVariable**](SmartHooksApi.md#CreateEnvironmentVariable) | **Post** /api/2/hooks/envs | Create Environment Variable
[**CreateHook**](SmartHooksApi.md#CreateHook) | **Post** /api/2/hooks | Create Smart Hook
[**DeleteEnvironmentVariable**](SmartHooksApi.md#DeleteEnvironmentVariable) | **Delete** /api/2/hooks/envs/{envvar_id} | Delete Environment Variable
[**DeleteHook**](SmartHooksApi.md#DeleteHook) | **Delete** /api/2/hooks/{hook_id} | Delete Smart Hook by ID
[**GetEnvironmentVariable**](SmartHooksApi.md#GetEnvironmentVariable) | **Get** /api/2/hooks/envs/{envvar_id} | Get Environment Variable
[**GetHook**](SmartHooksApi.md#GetHook) | **Get** /api/2/hooks/{hook_id} | Get Smart Hook by ID
[**GetLogs**](SmartHooksApi.md#GetLogs) | **Get** /api/2/hooks/{hook_id}/logs | Get Smart Hook Logs
[**ListEnvironmentVariables**](SmartHooksApi.md#ListEnvironmentVariables) | **Get** /api/2/hooks/envs | List Environment Variables
[**ListHooks**](SmartHooksApi.md#ListHooks) | **Get** /api/2/hooks | List all Smart Hooks
[**UpdateEnvironmentVariable**](SmartHooksApi.md#UpdateEnvironmentVariable) | **Put** /api/2/hooks/envs/{envvar_id} | Update Environment Variable
[**UpdateHook**](SmartHooksApi.md#UpdateHook) | **Put** /api/2/hooks/{hook_id} | Update Smart Hook by ID



## CreateEnvironmentVariable

> HookEnvvar CreateEnvironmentVariable(ctx).HookEnvvar(hookEnvvar).Execute()

Create Environment Variable



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
    hookEnvvar := *openapiclient.NewHookEnvvar("Name_example", "Value_example") // HookEnvvar | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.CreateEnvironmentVariable(context.Background()).HookEnvvar(hookEnvvar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.CreateEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentVariable`: HookEnvvar
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.CreateEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hookEnvvar** | [**HookEnvvar**](HookEnvvar.md) |  | 

### Return type

[**HookEnvvar**](HookEnvvar.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHook

> Hook CreateHook(ctx).Hook(hook).Execute()

Create Smart Hook



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
    hook := *openapiclient.NewHook("Type_example", false, int32(123), []string{"EnvVars_example"}, "Runtime_example", int32(123), map[string]string{"key": "Inner_example"}, "Function_example") // Hook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.CreateHook(context.Background()).Hook(hook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.CreateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.CreateHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hook** | [**Hook**](Hook.md) |  | 

### Return type

[**Hook**](Hook.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentVariable

> DeleteEnvironmentVariable(ctx, envvarId).Execute()

Delete Environment Variable



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
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.DeleteEnvironmentVariable(context.Background(), envvarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.DeleteEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envvarId** | **string** | Set to the id of the Hook Environment Variable that you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentVariableRequest struct via the builder pattern


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


## DeleteHook

> DeleteHook(ctx, hookId).Execute()

Delete Smart Hook by ID



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
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.DeleteHook(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.DeleteHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Set to the id of the Hook that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHookRequest struct via the builder pattern


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


## GetEnvironmentVariable

> HookEnvvar GetEnvironmentVariable(ctx, envvarId).Execute()

Get Environment Variable



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
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.GetEnvironmentVariable(context.Background(), envvarId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.GetEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentVariable`: HookEnvvar
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.GetEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envvarId** | **string** | Set to the id of the Hook Environment Variable that you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HookEnvvar**](HookEnvvar.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHook

> Hook GetHook(ctx, hookId).Execute()

Get Smart Hook by ID



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
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.GetHook(context.Background(), hookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.GetHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.GetHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Set to the id of the Hook that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Hook**](Hook.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> []HookLog GetLogs(ctx, hookId).Limit(limit).Page(page).Cursor(cursor).RequestId(requestId).CorrelationId(correlationId).Execute()

Get Smart Hook Logs



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
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    requestId := "requestId_example" // string | Returns logs that contain this request_id. (optional)
    correlationId := "correlationId_example" // string | Returns logs that contain this correlation_id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.GetLogs(context.Background(), hookId).Limit(limit).Page(page).Cursor(cursor).RequestId(requestId).CorrelationId(correlationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: []HookLog
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Set to the id of the Hook that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **requestId** | **string** | Returns logs that contain this request_id. | 
 **correlationId** | **string** | Returns logs that contain this correlation_id. | 

### Return type

[**[]HookLog**](HookLog.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentVariables

> []HookEnvvar ListEnvironmentVariables(ctx).Limit(limit).Page(page).Cursor(cursor).Execute()

List Environment Variables



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
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.ListEnvironmentVariables(context.Background()).Limit(limit).Page(page).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.ListEnvironmentVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentVariables`: []HookEnvvar
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.ListEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]HookEnvvar**](HookEnvvar.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHooks

> []Hook ListHooks(ctx).Limit(limit).Page(page).Cursor(cursor).Execute()

List all Smart Hooks



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
    limit := int32(56) // int32 | How many items to return at one time (max 100) (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.ListHooks(context.Background()).Limit(limit).Page(page).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.ListHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHooks`: []Hook
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.ListHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many items to return at one time (max 100) | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]Hook**](Hook.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentVariable

> HookEnvvar UpdateEnvironmentVariable(ctx, envvarId).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()

Update Environment Variable



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
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.
    updateEnvironmentVariableRequest := *openapiclient.NewUpdateEnvironmentVariableRequest("Value_example") // UpdateEnvironmentVariableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.UpdateEnvironmentVariable(context.Background(), envvarId).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.UpdateEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentVariable`: HookEnvvar
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.UpdateEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envvarId** | **string** | Set to the id of the Hook Environment Variable that you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvironmentVariableRequest** | [**UpdateEnvironmentVariableRequest**](UpdateEnvironmentVariableRequest.md) |  | 

### Return type

[**HookEnvvar**](HookEnvvar.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHook

> Hook UpdateHook(ctx, hookId).Hook(hook).Execute()

Update Smart Hook by ID



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
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.
    hook := *openapiclient.NewHook("Type_example", false, int32(123), []string{"EnvVars_example"}, "Runtime_example", int32(123), map[string]string{"key": "Inner_example"}, "Function_example") // Hook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartHooksApi.UpdateHook(context.Background(), hookId).Hook(hook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartHooksApi.UpdateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `SmartHooksApi.UpdateHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookId** | **string** | Set to the id of the Hook that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hook** | [**Hook**](Hook.md) |  | 

### Return type

[**Hook**](Hook.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

