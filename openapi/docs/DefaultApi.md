# \DefaultApi

All URIs are relative to *https://onelogininc.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](DefaultApi.md#ActivateFactor) | **Post** /api/2/mfa/users/{user_id}/verifications | 
[**AddAccessTokenClaim**](DefaultApi.md#AddAccessTokenClaim) | **Post** /api/2/api_authorizations/{id}/claims | 
[**AddClientApp**](DefaultApi.md#AddClientApp) | **Post** /api/2/api_authorizations/{id}/clients | 
[**AddRoleAdmins**](DefaultApi.md#AddRoleAdmins) | **Post** /api/2/roles/{role_id}/admins | 
[**AddRoleUsers**](DefaultApi.md#AddRoleUsers) | **Post** /api/2/roles/{role_id}/users | 
[**AddScope**](DefaultApi.md#AddScope) | **Post** /api/2/api_authorizations/{id}/scopes | 
[**BulkMappingSort**](DefaultApi.md#BulkMappingSort) | **Put** /api/2/apps/mappings/sort | 
[**BulkSort**](DefaultApi.md#BulkSort) | **Put** /api/2/apps/{app_id}/rules/sort | 
[**CreateApp**](DefaultApi.md#CreateApp) | **Post** /api/2/apps | 
[**CreateAuthorizationServer**](DefaultApi.md#CreateAuthorizationServer) | **Post** /api/2/api_authorizations | 
[**CreateEnvironmentVariable**](DefaultApi.md#CreateEnvironmentVariable) | **Post** /api/2/hooks/envs | 
[**CreateHook**](DefaultApi.md#CreateHook) | **Post** /api/2/hooks | 
[**CreateMapping**](DefaultApi.md#CreateMapping) | **Post** /api/2/mappings | 
[**CreateRiskRule**](DefaultApi.md#CreateRiskRule) | **Post** /api/2/risk/rules | 
[**CreateRoles**](DefaultApi.md#CreateRoles) | **Post** /api/2/roles | 
[**CreateRule**](DefaultApi.md#CreateRule) | **Post** /api/2/apps/{app_id}/rules | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /api/2/users | 
[**DeleteAccessTokenClaim**](DefaultApi.md#DeleteAccessTokenClaim) | **Delete** /api/2/api_authorizations/{id}/claims/{claim_id} | 
[**DeleteApp**](DefaultApi.md#DeleteApp) | **Delete** /api/2/apps/{app_id} | 
[**DeleteAppParameter**](DefaultApi.md#DeleteAppParameter) | **Delete** /api/2/apps/{app_id}/parameters/{parameter_id} | 
[**DeleteAuthorizationServer**](DefaultApi.md#DeleteAuthorizationServer) | **Delete** /api/2/api_authorizations/{id} | 
[**DeleteEnvironmentVariable**](DefaultApi.md#DeleteEnvironmentVariable) | **Delete** /api/2/hooks/envs/{envvar_id} | 
[**DeleteFactor**](DefaultApi.md#DeleteFactor) | **Delete** /api/2/mfa/users/{user_id}/devices/{device_id} | 
[**DeleteHook**](DefaultApi.md#DeleteHook) | **Delete** /api/2/hooks/{hook_id} | 
[**DeleteMapping**](DefaultApi.md#DeleteMapping) | **Delete** /api/2/mappings/{mapping_id} | 
[**DeleteRiskRule**](DefaultApi.md#DeleteRiskRule) | **Delete** /api/2/risk/rules/{risk_rule_id} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /api/2/roles/{role_id} | 
[**DeleteRule**](DefaultApi.md#DeleteRule) | **Delete** /api/2/apps/{app_id}/rules/{rule_id} | 
[**DeleteScope**](DefaultApi.md#DeleteScope) | **Delete** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /api/2/users/{user_id} | 
[**DryRunMapping**](DefaultApi.md#DryRunMapping) | **Post** /api/2/mappings/{mapping_id}/dryrun | 
[**EnrollFactor**](DefaultApi.md#EnrollFactor) | **Post** /api/2/mfa/users/{user_id}/registrations | 
[**GenerateMfaToken**](DefaultApi.md#GenerateMfaToken) | **Post** /api/2/mfs/users/{user_id}/mfa_token | 
[**GenerateSamlAssertion**](DefaultApi.md#GenerateSamlAssertion) | **Post** /api/2/saml_assertion | 
[**GenerateToken**](DefaultApi.md#GenerateToken) | **Post** /auth/oauth2/v2/token | 
[**GetApp**](DefaultApi.md#GetApp) | **Get** /api/2/apps/{app_id} | 
[**GetAuthorizationServer**](DefaultApi.md#GetAuthorizationServer) | **Get** /api/2/api_authorizations/{id} | 
[**GetAvailableFactors**](DefaultApi.md#GetAvailableFactors) | **Get** /api/2/mfa/users/{user_id}/factors | 
[**GetClientApps**](DefaultApi.md#GetClientApps) | **Get** /api/2/api_authorizations/{id}/clients | 
[**GetEnrolledFactors**](DefaultApi.md#GetEnrolledFactors) | **Get** /api/2/mfa/users/{user_id}/devices | 
[**GetEnvironmentVariable**](DefaultApi.md#GetEnvironmentVariable) | **Get** /api/2/hooks/envs/{envvar_id} | 
[**GetHook**](DefaultApi.md#GetHook) | **Get** /api/2/hooks/{hook_id} | 
[**GetLogs**](DefaultApi.md#GetLogs) | **Get** /api/2/hooks/{hook_id}/logs | 
[**GetMapping**](DefaultApi.md#GetMapping) | **Get** /api/2/mappings/{mapping_id} | 
[**GetRateLimit**](DefaultApi.md#GetRateLimit) | **Get** /auth/rate_limit | 
[**GetRiskRule**](DefaultApi.md#GetRiskRule) | **Get** /api/2/risk/rules/{risk_rule_id} | 
[**GetRiskScore**](DefaultApi.md#GetRiskScore) | **Post** /api/2/risk/verify | 
[**GetRole**](DefaultApi.md#GetRole) | **Get** /api/2/roles/{role_id} | 
[**GetRoleAdmins**](DefaultApi.md#GetRoleAdmins) | **Get** /api/2/roles/{role_id}/admins | 
[**GetRoleApps**](DefaultApi.md#GetRoleApps) | **Get** /api/2/roles/{role_id}/apps | 
[**GetRoleUsers**](DefaultApi.md#GetRoleUsers) | **Get** /api/2/roles/{role_id}/users | 
[**GetRule**](DefaultApi.md#GetRule) | **Get** /api/2/apps/{app_id}/rules/{rule_id} | 
[**GetScoreInsights**](DefaultApi.md#GetScoreInsights) | **Get** /api/2/risk/scores | 
[**GetUser**](DefaultApi.md#GetUser) | **Get** /api/2/users/{user_id} | 
[**GetUserApps**](DefaultApi.md#GetUserApps) | **Get** /api/2/users/{user_id}/apps | 
[**ListAccessTokenClaims**](DefaultApi.md#ListAccessTokenClaims) | **Get** /api/2/api_authorizations/{id}/claims | 
[**ListActionValues**](DefaultApi.md#ListActionValues) | **Get** /api/2/apps/{app_id}/rules/actions/{actuion_value}/values | 
[**ListActions**](DefaultApi.md#ListActions) | **Get** /api/2/apps/{app_id}/rules/actions | 
[**ListAppUsers**](DefaultApi.md#ListAppUsers) | **Get** /api/2/apps/{app_id}/users | 
[**ListApps**](DefaultApi.md#ListApps) | **Get** /api/2/apps | 
[**ListAuthorizationServers**](DefaultApi.md#ListAuthorizationServers) | **Get** /api/2/api_authorizations | 
[**ListConditionOperators**](DefaultApi.md#ListConditionOperators) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/operators | 
[**ListConditionValues**](DefaultApi.md#ListConditionValues) | **Get** /api/2/apps/{app_id}/rules/conditions/{condition_value}/values | 
[**ListConditions**](DefaultApi.md#ListConditions) | **Get** /api/2/apps/{app_id}/rules/conditions | 
[**ListConnectors**](DefaultApi.md#ListConnectors) | **Get** /api/2/connectors | 
[**ListEnvironmentVariables**](DefaultApi.md#ListEnvironmentVariables) | **Get** /api/2/hooks/envs | 
[**ListHooks**](DefaultApi.md#ListHooks) | **Get** /api/2/hooks | 
[**ListMappingActionValues**](DefaultApi.md#ListMappingActionValues) | **Get** /api/2/apps/mappings/actions/{actuion_value}/values | 
[**ListMappingActions**](DefaultApi.md#ListMappingActions) | **Get** /api/2/apps/mappings/actions | 
[**ListMappingConditionOperators**](DefaultApi.md#ListMappingConditionOperators) | **Get** /api/2/apps/mappings/conditions/{condition_value}/operators | 
[**ListMappingConditionValues**](DefaultApi.md#ListMappingConditionValues) | **Get** /api/2/apps/mappings/conditions/{condition_value}/values | 
[**ListMappingConditions**](DefaultApi.md#ListMappingConditions) | **Get** /api/2/apps/mappings/conditions | 
[**ListMappings**](DefaultApi.md#ListMappings) | **Get** /api/2/mappings | 
[**ListRiskRules**](DefaultApi.md#ListRiskRules) | **Get** /api/2/risk/rules | 
[**ListRoles**](DefaultApi.md#ListRoles) | **Get** /api/2/roles | 
[**ListRules**](DefaultApi.md#ListRules) | **Get** /api/2/apps/{app_id}/rules | 
[**ListScopes**](DefaultApi.md#ListScopes) | **Get** /api/2/api_authorizations/{id}/scopes | 
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /api/2/users | 
[**RemoveClientApp**](DefaultApi.md#RemoveClientApp) | **Delete** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
[**RemoveRoleAdmins**](DefaultApi.md#RemoveRoleAdmins) | **Delete** /api/2/roles/{role_id}/admins | 
[**RemoveRoleUsers**](DefaultApi.md#RemoveRoleUsers) | **Delete** /api/2/roles/{role_id}/users | 
[**RevokeToken**](DefaultApi.md#RevokeToken) | **Post** /auth/oauth2/revoke | 
[**SetRoleApps**](DefaultApi.md#SetRoleApps) | **Put** /api/2/roles/{role_id}/apps | 
[**TrackEvent**](DefaultApi.md#TrackEvent) | **Post** /api/2/risk/events | 
[**UpdateAccessTokenClaim**](DefaultApi.md#UpdateAccessTokenClaim) | **Put** /api/2/api_authorizations/{id}/claims/{claim_id} | 
[**UpdateApp**](DefaultApi.md#UpdateApp) | **Put** /api/2/apps/{app_id} | 
[**UpdateAuthorizationServer**](DefaultApi.md#UpdateAuthorizationServer) | **Put** /api/2/api_authorizations/{id} | 
[**UpdateClientApp**](DefaultApi.md#UpdateClientApp) | **Put** /api/2/api_authorizations/{id}/clients/{client_app_id} | 
[**UpdateEnvironmentVariable**](DefaultApi.md#UpdateEnvironmentVariable) | **Put** /api/2/hooks/envs/{envvar_id} | 
[**UpdateHook**](DefaultApi.md#UpdateHook) | **Put** /api/2/hooks/{hook_id} | 
[**UpdateMapping**](DefaultApi.md#UpdateMapping) | **Put** /api/2/mappings/{mapping_id} | 
[**UpdateRiskRule**](DefaultApi.md#UpdateRiskRule) | **Put** /api/2/risk/rules/{risk_rule_id} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Put** /api/2/roles/{role_id} | 
[**UpdateRule**](DefaultApi.md#UpdateRule) | **Put** /api/2/apps/{app_id}/rules/{rule_id} | 
[**UpdateScope**](DefaultApi.md#UpdateScope) | **Put** /api/2/api_authorizations/{id}/scopes/{scope_id} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Put** /api/2/users/{user_id} | 
[**VerifyEnrollment**](DefaultApi.md#VerifyEnrollment) | **Put** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
[**VerifyEnrollmentVoiceProtect**](DefaultApi.md#VerifyEnrollmentVoiceProtect) | **Get** /api/2/mfa/users/{user_id}/registrations/{registration_id} | 
[**VerifyFactor**](DefaultApi.md#VerifyFactor) | **Put** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 
[**VerifyFactorSaml**](DefaultApi.md#VerifyFactorSaml) | **Post** /api/2/saml_assertion/verify_factor | 
[**VerifyFactorVoice**](DefaultApi.md#VerifyFactorVoice) | **Get** /api/2/mfa/users/{user_id}/verifications/{verification_id} | 



## ActivateFactor

> ActivateFactor(ctx, userId).Authorization(authorization).ActivateFactorRequest(activateFactorRequest).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    activateFactorRequest := *openapiclient.NewActivateFactorRequest() // ActivateFactorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ActivateFactor(context.Background(), userId).Authorization(authorization).ActivateFactorRequest(activateFactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ActivateFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **activateFactorRequest** | [**ActivateFactorRequest**](ActivateFactorRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessTokenClaim

> Id AddAccessTokenClaim(ctx, id).Authorization(authorization).AddAccessTokenClaimRequest(addAccessTokenClaimRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    addAccessTokenClaimRequest := *openapiclient.NewAddAccessTokenClaimRequest() // AddAccessTokenClaimRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddAccessTokenClaim(context.Background(), id).Authorization(authorization).AddAccessTokenClaimRequest(addAccessTokenClaimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddAccessTokenClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessTokenClaim`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddAccessTokenClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessTokenClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **addAccessTokenClaimRequest** | [**AddAccessTokenClaimRequest**](AddAccessTokenClaimRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClientApp

> ClientApp AddClientApp(ctx, id).Authorization(authorization).AddClientAppRequest(addClientAppRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    addClientAppRequest := *openapiclient.NewAddClientAppRequest() // AddClientAppRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddClientApp(context.Background(), id).Authorization(authorization).AddClientAppRequest(addClientAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClientApp`: ClientApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **addClientAppRequest** | [**AddClientAppRequest**](AddClientAppRequest.md) |  | 

### Return type

[**ClientApp**](ClientApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleAdmins

> []AddRoleUsers200ResponseInner AddRoleAdmins(ctx, roleId).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddRoleAdmins(context.Background(), roleId).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleAdmins`: []AddRoleUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddRoleAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **requestBody** | **[]int32** |  | 

### Return type

[**[]AddRoleUsers200ResponseInner**](AddRoleUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleUsers

> []AddRoleUsers200ResponseInner AddRoleUsers(ctx, roleId).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddRoleUsers(context.Background(), roleId).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRoleUsers`: []AddRoleUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **requestBody** | **[]int32** |  | 

### Return type

[**[]AddRoleUsers200ResponseInner**](AddRoleUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddScope

> Id AddScope(ctx, id).Authorization(authorization).AddScopeRequest(addScopeRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    addScopeRequest := *openapiclient.NewAddScopeRequest() // AddScopeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddScope(context.Background(), id).Authorization(authorization).AddScopeRequest(addScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScope`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **addScopeRequest** | [**AddScopeRequest**](AddScopeRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkMappingSort

> []int32 BulkMappingSort(ctx).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    requestBody := []int32{int32(123)} // []int32 | The request body must contain an array of User Mapping IDs in the desired order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.BulkMappingSort(context.Background()).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BulkMappingSort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkMappingSort`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BulkMappingSort`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkMappingSortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **requestBody** | **[]int32** | The request body must contain an array of User Mapping IDs in the desired order. | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSort

> []int32 BulkSort(ctx, appId).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    requestBody := []int32{int32(123)} // []int32 | The request body must contain an array of App Rule IDs in the desired order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.BulkSort(context.Background(), appId).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BulkSort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkSort`: []int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BulkSort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkSortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **requestBody** | **[]int32** | The request body must contain an array of App Rule IDs in the desired order. | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApp

> Schema CreateApp(ctx).Authorization(authorization).Schema(schema).Execute()



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
    authorization := "authorization_example" // string | 
    schema := *openapiclient.NewSchema() // Schema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateApp(context.Background()).Authorization(authorization).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationServer

> Id CreateAuthorizationServer(ctx).Authorization(authorization).CreateAuthorizationServerRequest(createAuthorizationServerRequest).Execute()



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
    authorization := "authorization_example" // string | 
    createAuthorizationServerRequest := *openapiclient.NewCreateAuthorizationServerRequest() // CreateAuthorizationServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAuthorizationServer(context.Background()).Authorization(authorization).CreateAuthorizationServerRequest(createAuthorizationServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServer`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **createAuthorizationServerRequest** | [**CreateAuthorizationServerRequest**](CreateAuthorizationServerRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironmentVariable

> Envvar CreateEnvironmentVariable(ctx).Authorization(authorization).CreateEnvironmentVariableRequest(createEnvironmentVariableRequest).Execute()



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
    authorization := "authorization_example" // string | 
    createEnvironmentVariableRequest := *openapiclient.NewCreateEnvironmentVariableRequest("Name_example", "Value_example") // CreateEnvironmentVariableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEnvironmentVariable(context.Background()).Authorization(authorization).CreateEnvironmentVariableRequest(createEnvironmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentVariable`: Envvar
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEnvironmentVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **createEnvironmentVariableRequest** | [**CreateEnvironmentVariableRequest**](CreateEnvironmentVariableRequest.md) |  | 

### Return type

[**Envvar**](Envvar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHook

> CreateHook(ctx).Authorization(authorization).Hook(hook).Execute()



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
    authorization := "authorization_example" // string | 
    hook := *openapiclient.NewHook("Type_example", false, int32(123), []string{"EnvVars_example"}, "Runtime_example", int32(123), map[string]interface{}(123), "Function_example") // Hook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateHook(context.Background()).Authorization(authorization).Hook(hook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **hook** | [**Hook**](Hook.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMapping

> int32 CreateMapping(ctx).Authorization(authorization).Mapping(mapping).Execute()



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
    authorization := "authorization_example" // string | 
    mapping := *openapiclient.NewMapping("Name_example", false, "Match_example", int32(123), []openapiclient.Action{*openapiclient.NewAction()}) // Mapping | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMapping(context.Background()).Authorization(authorization).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMapping`: int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **mapping** | [**Mapping**](Mapping.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRiskRule

> CreateRiskRule(ctx).Authorization(authorization).RiskRule(riskRule).Execute()



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
    authorization := "authorization_example" // string | 
    riskRule := *openapiclient.NewRiskRule() // RiskRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateRiskRule(context.Background()).Authorization(authorization).RiskRule(riskRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **riskRule** | [**RiskRule**](RiskRule.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoles

> []CreateRoles201ResponseInner CreateRoles(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateRoles(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoles`: []CreateRoles201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]CreateRoles201ResponseInner**](CreateRoles201ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> RuleId CreateRule(ctx, appId).Authorization(authorization).Rule(rule).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    rule := *openapiclient.NewRule() // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateRule(context.Background(), appId).Authorization(authorization).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: RuleId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **rule** | [**Rule**](Rule.md) |  | 

### Return type

[**RuleId**](RuleId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).Authorization(authorization).User(user).Mappings(mappings).ValidatePolicy(validatePolicy).Execute()



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
    authorization := "authorization_example" // string | 
    user := *openapiclient.NewUser() // User | 
    mappings := "mappings_example" // string | Controls how mappings will be applied to the user on creation. Defaults to async. (optional)
    validatePolicy := true // bool | Will passwords validate against the User Policy? Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateUser(context.Background()).Authorization(authorization).User(user).Mappings(mappings).ValidatePolicy(validatePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **user** | [**User**](User.md) |  | 
 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessTokenClaim

> DeleteAccessTokenClaim(ctx, id, claimId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    claimId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAccessTokenClaim(context.Background(), id, claimId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAccessTokenClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**claimId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessTokenClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, appId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteApp(context.Background(), appId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppParameter

> DeleteAppParameter(ctx, appId, parameterId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    parameterId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAppParameter(context.Background(), appId, parameterId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAppParameter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**parameterId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppParameterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthorizationServer

> DeleteAuthorizationServer(ctx, id).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAuthorizationServer(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentVariable

> DeleteEnvironmentVariable(ctx, envvarId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteEnvironmentVariable(context.Background(), envvarId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEnvironmentVariable``: %v\n", err)
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
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFactor

> DeleteFactor(ctx, userId, deviceId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    deviceId := int32(56) // int32 | Set to the device_id of the MFA device.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteFactor(context.Background(), userId, deviceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 
**deviceId** | **int32** | Set to the device_id of the MFA device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHook

> DeleteHook(ctx, hookId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteHook(context.Background(), hookId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteHook``: %v\n", err)
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
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMapping

> DeleteMapping(ctx, mappingId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    mappingId := int32(56) // int32 | The id of the user mapping to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteMapping(context.Background(), mappingId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMapping``: %v\n", err)
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
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskRule

> RiskRule DeleteRiskRule(ctx, riskRuleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    riskRuleId := "riskRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRiskRule(context.Background(), riskRuleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRiskRule`: RiskRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**RiskRule**](RiskRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRole(context.Background(), roleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, appId, ruleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    ruleId := int32(56) // int32 | The id of the app rule to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRule(context.Background(), appId, ruleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **int32** | The id of the app rule to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScope

> DeleteScope(ctx, id, scopeId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    scopeId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteScope(context.Background(), id, scopeId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**scopeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteUser(context.Background(), userId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryRunMapping

> []map[string]interface{} DryRunMapping(ctx, mappingId).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    mappingId := int32(56) // int32 | The id of the user mapping to locate.
    requestBody := []int32{int32(123)} // []int32 | Request body is a list of user IDs tested against the mapping conditions to verify that the mapping would be applied

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DryRunMapping(context.Background(), mappingId).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DryRunMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DryRunMapping`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DryRunMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **int32** | The id of the user mapping to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDryRunMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **requestBody** | **[]int32** | Request body is a list of user IDs tested against the mapping conditions to verify that the mapping would be applied | 

### Return type

**[]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollFactor

> [][]FactorInner EnrollFactor(ctx, userId).Authorization(authorization).EnrollFactorRequest(enrollFactorRequest).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    enrollFactorRequest := *openapiclient.NewEnrollFactorRequest(int32(123), "DisplayName_example") // EnrollFactorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EnrollFactor(context.Background(), userId).Authorization(authorization).EnrollFactorRequest(enrollFactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollFactor`: [][]FactorInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **enrollFactorRequest** | [**EnrollFactorRequest**](EnrollFactorRequest.md) |  | 

### Return type

[**[][]FactorInner**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateMfaToken

> GenerateMfaToken200Response GenerateMfaToken(ctx).Authorization(authorization).GenerateMfaTokenRequest(generateMfaTokenRequest).Execute()



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
    authorization := "authorization_example" // string | 
    generateMfaTokenRequest := *openapiclient.NewGenerateMfaTokenRequest() // GenerateMfaTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GenerateMfaToken(context.Background()).Authorization(authorization).GenerateMfaTokenRequest(generateMfaTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GenerateMfaToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateMfaToken`: GenerateMfaToken200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GenerateMfaToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMfaTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **generateMfaTokenRequest** | [**GenerateMfaTokenRequest**](GenerateMfaTokenRequest.md) |  | 

### Return type

[**GenerateMfaToken200Response**](GenerateMfaToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSamlAssertion

> GenerateSamlAssertion(ctx).Authorization(authorization).GenerateSamlAssertionRequest(generateSamlAssertionRequest).Execute()



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
    authorization := "authorization_example" // string | 
    generateSamlAssertionRequest := *openapiclient.NewGenerateSamlAssertionRequest("UsernameOrEmail_example", "Password_example", "AppId_example", "Subdomain_example") // GenerateSamlAssertionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GenerateSamlAssertion(context.Background()).Authorization(authorization).GenerateSamlAssertionRequest(generateSamlAssertionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GenerateSamlAssertion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSamlAssertionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **generateSamlAssertionRequest** | [**GenerateSamlAssertionRequest**](GenerateSamlAssertionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateToken

> GenerateToken200Response GenerateToken(ctx).Authorization(authorization).GenerateTokenRequest(generateTokenRequest).Execute()



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
    authorization := "authorization_example" // string | 
    generateTokenRequest := *openapiclient.NewGenerateTokenRequest() // GenerateTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GenerateToken(context.Background()).Authorization(authorization).GenerateTokenRequest(generateTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateToken`: GenerateToken200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **generateTokenRequest** | [**GenerateTokenRequest**](GenerateTokenRequest.md) |  | 

### Return type

[**GenerateToken200Response**](GenerateToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> Schema GetApp(ctx, appId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetApp(context.Background(), appId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationServer

> GetAuthorizationServer200Response GetAuthorizationServer(ctx, id).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAuthorizationServer(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServer`: GetAuthorizationServer200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**GetAuthorizationServer200Response**](GetAuthorizationServer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableFactors

> []GetAvailableFactors200ResponseInner GetAvailableFactors(ctx, userId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAvailableFactors(context.Background(), userId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvailableFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableFactors`: []GetAvailableFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAvailableFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]GetAvailableFactors200ResponseInner**](GetAvailableFactors200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientApps

> []GetClientApps200ResponseInner GetClientApps(ctx, id).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClientApps(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClientApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientApps`: []GetClientApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClientApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]GetClientApps200ResponseInner**](GetClientApps200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrolledFactors

> []Device GetEnrolledFactors(ctx, userId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEnrolledFactors(context.Background(), userId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnrolledFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnrolledFactors`: []Device
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnrolledFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrolledFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentVariable

> Envvar GetEnvironmentVariable(ctx, envvarId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEnvironmentVariable(context.Background(), envvarId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentVariable`: Envvar
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironmentVariable`: %v\n", resp)
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
 **authorization** | **string** |  | 


### Return type

[**Envvar**](Envvar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHook

> Hook GetHook(ctx, hookId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetHook(context.Background(), hookId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHook`: %v\n", resp)
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
 **authorization** | **string** |  | 


### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> []Log GetLogs(ctx, hookId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).RequestId(requestId).CorrelationId(correlationId).Execute()



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
    authorization := "authorization_example" // string | 
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    requestId := "requestId_example" // string | Returns logs that contain this request_id. (optional)
    correlationId := "correlationId_example" // string | Returns logs that contain this correlation_id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetLogs(context.Background(), hookId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).RequestId(requestId).CorrelationId(correlationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: []Log
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLogs`: %v\n", resp)
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
 **authorization** | **string** |  | 

 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **requestId** | **string** | Returns logs that contain this request_id. | 
 **correlationId** | **string** | Returns logs that contain this correlation_id. | 

### Return type

[**[]Log**](Log.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapping

> Mapping GetMapping(ctx, mappingId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    mappingId := int32(56) // int32 | The id of the user mapping to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetMapping(context.Background(), mappingId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapping`: Mapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMapping`: %v\n", resp)
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
 **authorization** | **string** |  | 


### Return type

[**Mapping**](Mapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateLimit

> GetRateLimit200Response GetRateLimit(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRateLimit(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateLimit`: GetRateLimit200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRateLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetRateLimit200Response**](GetRateLimit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskRule

> GetRiskRule(ctx, riskRuleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    riskRuleId := "riskRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRiskRule(context.Background(), riskRuleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskScore

> GetRiskScore200Response GetRiskScore(ctx).Authorization(authorization).GetRiskScoreRequest(getRiskScoreRequest).Execute()



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
    authorization := "authorization_example" // string | 
    getRiskScoreRequest := *openapiclient.NewGetRiskScoreRequest("Ip_example", "UserAgent_example", *openapiclient.NewRiskUser("Id_example")) // GetRiskScoreRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRiskScore(context.Background()).Authorization(authorization).GetRiskScoreRequest(getRiskScoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRiskScore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskScore`: GetRiskScore200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRiskScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **getRiskScoreRequest** | [**GetRiskScoreRequest**](GetRiskScoreRequest.md) |  | 

### Return type

[**GetRiskScore200Response**](GetRiskScore200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRole(context.Background(), roleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAdmins

> []Schema1 GetRoleAdmins(ctx, roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | Allows you to filter on first name, last name, username, and email address. (optional)
    includeUnassigned := true // bool | Optional. Defaults to false. Include users that aren’t assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoleAdmins(context.Background(), roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleAdmins`: []Schema1
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoleAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **bool** | Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]Schema1**](Schema1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleApps

> []Schema GetRoleApps(ctx, roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Assigned(assigned).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    assigned := true // bool | Optional. Defaults to true. Returns all apps not yet assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoleApps(context.Background(), roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Assigned(assigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoleApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleApps`: []Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoleApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **assigned** | **bool** | Optional. Defaults to true. Returns all apps not yet assigned to the role. | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleUsers

> []Schema1 GetRoleUsers(ctx, roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | Allows you to filter on first name, last name, username, and email address. (optional)
    includeUnassigned := true // bool | Optional. Defaults to false. Include users that aren’t assigned to the role. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoleUsers(context.Background(), roleId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).IncludeUnassigned(includeUnassigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleUsers`: []Schema1
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | Allows you to filter on first name, last name, username, and email address. | 
 **includeUnassigned** | **bool** | Optional. Defaults to false. Include users that aren’t assigned to the role. | 

### Return type

[**[]Schema1**](Schema1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> Rule GetRule(ctx, appId, ruleId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    ruleId := int32(56) // int32 | The id of the app rule to locate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRule(context.Background(), appId, ruleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: Rule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **int32** | The id of the app rule to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScoreInsights

> GetScoreInsights200Response GetScoreInsights(ctx).Authorization(authorization).Before(before).After(after).Execute()



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
    authorization := "authorization_example" // string | 
    before := "before_example" // string | Optional ISO8601 formatted date string. Defaults to current date. Maximum date is 90 days ago. (optional)
    after := "after_example" // string | Optional ISO8601 formatted date string. Defaults to 30 days ago. Maximum date is 90 days ago. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetScoreInsights(context.Background()).Authorization(authorization).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetScoreInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScoreInsights`: GetScoreInsights200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetScoreInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScoreInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **before** | **string** | Optional ISO8601 formatted date string. Defaults to current date. Maximum date is 90 days ago. | 
 **after** | **string** | Optional ISO8601 formatted date string. Defaults to 30 days ago. Maximum date is 90 days ago. | 

### Return type

[**GetScoreInsights200Response**](GetScoreInsights200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user that you want to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUser(context.Background(), userId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApps

> []GetUserApps200ResponseInner GetUserApps(ctx, userId).Authorization(authorization).IgnoreVisibility(ignoreVisibility).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user that you want to return.
    ignoreVisibility := true // bool | Defaults to `false`. When `true` will show all apps that are assigned to a user regardless of their portal visibility setting. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserApps(context.Background(), userId).Authorization(authorization).IgnoreVisibility(ignoreVisibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserApps`: []GetUserApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **ignoreVisibility** | **bool** | Defaults to &#x60;false&#x60;. When &#x60;true&#x60; will show all apps that are assigned to a user regardless of their portal visibility setting. | 

### Return type

[**[]GetUserApps200ResponseInner**](GetUserApps200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessTokenClaims

> []ListAccessTokenClaims200ResponseInner ListAccessTokenClaims(ctx, id).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAccessTokenClaims(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccessTokenClaims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessTokenClaims`: []ListAccessTokenClaims200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccessTokenClaims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessTokenClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]ListAccessTokenClaims200ResponseInner**](ListAccessTokenClaims200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActionValues

> []ListConditionValues200ResponseInner ListActionValues(ctx, appId, actionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    actionValue := "actionValue_example" // string | The value for the selected action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListActionValues(context.Background(), appId, actionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListActionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActionValues`: []ListConditionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListActionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**actionValue** | **string** | The value for the selected action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**[]ListConditionValues200ResponseInner**](ListConditionValues200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActions

> []ListActions200ResponseInner ListActions(ctx, appId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListActions(context.Background(), appId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActions`: []ListActions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListActions`: %v\n", resp)
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
 **authorization** | **string** |  | 


### Return type

[**[]ListActions200ResponseInner**](ListActions200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppUsers

> []ListAppUsers200ResponseInner ListAppUsers(ctx, appId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAppUsers(context.Background(), appId).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppUsers`: []ListAppUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]ListAppUsers200ResponseInner**](ListAppUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> []Schema ListApps(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).ConnectorId(connectorId).AuthMethod(authMethod).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | The name or partial name of the app to search for. When using a partial name you must append a wildcard `*` (optional)
    connectorId := int32(56) // int32 | Returns all apps based off a specific connector. See List Connectors for a complete list of Connector IDs. (optional)
    authMethod := openapiclient.auth_method(0) // AuthMethod | Returns all apps based of a given type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListApps(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).ConnectorId(connectorId).AuthMethod(authMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: []Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | The name or partial name of the app to search for. When using a partial name you must append a wildcard &#x60;*&#x60; | 
 **connectorId** | **int32** | Returns all apps based off a specific connector. See List Connectors for a complete list of Connector IDs. | 
 **authMethod** | [**AuthMethod**](AuthMethod.md) | Returns all apps based of a given type. | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationServers

> []ListAuthorizationServers200ResponseInner ListAuthorizationServers(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAuthorizationServers(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAuthorizationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServers`: []ListAuthorizationServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAuthorizationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]ListAuthorizationServers200ResponseInner**](ListAuthorizationServers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditionOperators

> []ListConditionOperators200ResponseInner ListConditionOperators(ctx, appId, conditionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    conditionValue := "conditionValue_example" // string | The value for the selected condition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConditionOperators(context.Background(), appId, conditionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConditionOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditionOperators`: []ListConditionOperators200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConditionOperators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**conditionValue** | **string** | The value for the selected condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**[]ListConditionOperators200ResponseInner**](ListConditionOperators200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditionValues

> []ListConditionValues200ResponseInner ListConditionValues(ctx, appId, conditionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    conditionValue := "conditionValue_example" // string | The value for the selected condition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConditionValues(context.Background(), appId, conditionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConditionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditionValues`: []ListConditionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConditionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**conditionValue** | **string** | The value for the selected condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**[]ListConditionValues200ResponseInner**](ListConditionValues200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditions

> []ListConditions200ResponseInner ListConditions(ctx, appId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConditions(context.Background(), appId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditions`: []ListConditions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConditions`: %v\n", resp)
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
 **authorization** | **string** |  | 


### Return type

[**[]ListConditions200ResponseInner**](ListConditions200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> []Connector ListConnectors(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).AuthMethod(authMethod).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | The name or partial name of the connector to search for. When using a partial name you must append a wildcard `*` (optional)
    authMethod := openapiclient.auth_method(0) // AuthMethod | Returns all connectors of a given type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListConnectors(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).AuthMethod(authMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectors`: []Connector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | The name or partial name of the connector to search for. When using a partial name you must append a wildcard &#x60;*&#x60; | 
 **authMethod** | [**AuthMethod**](AuthMethod.md) | Returns all connectors of a given type. | 

### Return type

[**[]Connector**](Connector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentVariables

> []Envvar ListEnvironmentVariables(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEnvironmentVariables(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEnvironmentVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentVariables`: []Envvar
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]Envvar**](Envvar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHooks

> []Hook ListHooks(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListHooks(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListHooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHooks`: []Hook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 

### Return type

[**[]Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingActionValues

> []ListConditionValues200ResponseInner ListMappingActionValues(ctx, actionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    actionValue := "actionValue_example" // string | The value for the selected action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappingActionValues(context.Background(), actionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappingActionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingActionValues`: []ListConditionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappingActionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**actionValue** | **string** | The value for the selected action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingActionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]ListConditionValues200ResponseInner**](ListConditionValues200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingActions

> []ListActions200ResponseInner ListMappingActions(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappingActions(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappingActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingActions`: []ListActions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappingActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]ListActions200ResponseInner**](ListActions200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingConditionOperators

> []ListMappingConditionOperators200ResponseInner ListMappingConditionOperators(ctx, conditionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    conditionValue := "conditionValue_example" // string | The value for the selected condition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappingConditionOperators(context.Background(), conditionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappingConditionOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingConditionOperators`: []ListMappingConditionOperators200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappingConditionOperators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionValue** | **string** | The value for the selected condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingConditionOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]ListMappingConditionOperators200ResponseInner**](ListMappingConditionOperators200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingConditionValues

> []ListConditionValues200ResponseInner ListMappingConditionValues(ctx, conditionValue).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    conditionValue := "conditionValue_example" // string | The value for the selected condition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappingConditionValues(context.Background(), conditionValue).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappingConditionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingConditionValues`: []ListConditionValues200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappingConditionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionValue** | **string** | The value for the selected condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingConditionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]ListConditionValues200ResponseInner**](ListConditionValues200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappingConditions

> []ListMappingConditions200ResponseInner ListMappingConditions(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappingConditions(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappingConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappingConditions`: []ListMappingConditions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappingConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]ListMappingConditions200ResponseInner**](ListMappingConditions200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappings

> []Mapping ListMappings(ctx).Authorization(authorization).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()



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
    authorization := "authorization_example" // string | 
    enabled := true // bool | Defaults to true. When set to `false` will return all disabled mappings. (optional) (default to true)
    hasCondition := "has_condition=has_role:123456,status:1" // string | Filters Mappings based on their Conditions. (optional)
    hasConditionType := "hasConditionType_example" // string | Filters Mappings based on their condition types. (optional)
    hasAction := "has_action=set_groups:123456,set_usertype:*" // string | Filters Mappings based on their Actions. (optional)
    hasActionType := "hasActionType_example" // string | Filters Mappings based on their action types. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMappings(context.Background()).Authorization(authorization).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappings`: []Mapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **enabled** | **bool** | Defaults to true. When set to &#x60;false&#x60; will return all disabled mappings. | [default to true]
 **hasCondition** | **string** | Filters Mappings based on their Conditions. | 
 **hasConditionType** | **string** | Filters Mappings based on their condition types. | 
 **hasAction** | **string** | Filters Mappings based on their Actions. | 
 **hasActionType** | **string** | Filters Mappings based on their action types. | 

### Return type

[**[]Mapping**](Mapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRiskRules

> ListRiskRules(ctx).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRiskRules(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRiskRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRiskRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []Role ListRoles(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).AppId(appId).Fields(fields).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    name := "name_example" // string | Optional. Filters by role name. (optional)
    appId := "appId_example" // string | Optional. Returns roles that contain this app name. (optional)
    fields := "fields_example" // string | Optional. Comma delimited list of fields to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRoles(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).Name(name).AppId(appId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **name** | **string** | Optional. Filters by role name. | 
 **appId** | **string** | Optional. Returns roles that contain this app name. | 
 **fields** | **string** | Optional. Comma delimited list of fields to return. | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> []Rule ListRules(ctx, appId).Authorization(authorization).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    enabled := true // bool | Defaults to true. When set to `false` will return all disabled rules. (optional)
    hasCondition := "hasCondition_example" // string | Filters Rules based on their Conditions. (optional)
    hasConditionType := "hasConditionType_example" // string | Filters Rules based on their condition types. (optional)
    hasAction := "hasAction_example" // string | Filters Rules based on their Actions. (optional)
    hasActionType := "hasActionType_example" // string | Filters Rules based on their action types. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRules(context.Background(), appId).Authorization(authorization).Enabled(enabled).HasCondition(hasCondition).HasConditionType(hasConditionType).HasAction(hasAction).HasActionType(hasActionType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRules`: []Rule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **enabled** | **bool** | Defaults to true. When set to &#x60;false&#x60; will return all disabled rules. | 
 **hasCondition** | **string** | Filters Rules based on their Conditions. | 
 **hasConditionType** | **string** | Filters Rules based on their condition types. | 
 **hasAction** | **string** | Filters Rules based on their Actions. | 
 **hasActionType** | **string** | Filters Rules based on their action types. | 

### Return type

[**[]Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScopes

> []ListScopes200ResponseInner ListScopes(ctx, id).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListScopes(context.Background(), id).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScopes`: []ListScopes200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


### Return type

[**[]ListScopes200ResponseInner**](ListScopes200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).AppId(appId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()



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
    authorization := "authorization_example" // string | 
    limit := int32(56) // int32 | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. (optional)
    page := int32(56) // int32 | The page number of results to return. (optional)
    cursor := "cursor_example" // string | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. (optional)
    createdSince := "createdSince_example" // string | An ISO8601 timestamp value that returns all users created after a given date & time. (optional)
    createdUntil := "createdUntil_example" // string | An ISO8601 timestamp value that returns all users created before a given date & time. (optional)
    updatedSince := "updatedSince_example" // string | An ISO8601 timestamp value that returns all users updated after a given date & time. (optional)
    updatedUntil := "updatedUntil_example" // string | An ISO8601 timestamp value that returns all users updated before a given date & time. (optional)
    lastLoginSince := "lastLoginSince_example" // string | An ISO8601 timestamp value that returns all users that logged in after a given date & time. (optional)
    lastLoginUntil := "lastLoginUntil_example" // string |  (optional)
    firstname := "firstname_example" // string | The first name of the user (optional)
    lastname := "lastname_example" // string | The last name of the user (optional)
    email := "email_example" // string | The email address of the user (optional)
    username := "username_example" // string | The username for the user (optional)
    samaccountname := "samaccountname_example" // string | The AD login name for the user (optional)
    directoryId := "directoryId_example" // string | The ID in OneLogin of the Directory that the user belongs to (optional)
    externalId := "externalId_example" // string | An external identifier that has been set on the user (optional)
    appId := "appId_example" // string | The ID of a OneLogin Application (optional)
    userIds := "userIds_example" // string | A comma separated list of OneLogin User IDs (optional)
    customAttributesAttributeName := "customAttributesAttributeName_example" // string | The short name of a custom attribute. Note that the attribute name is prefixed with custom_attributes. (optional)
    fields := "fields_example" // string | A comma separated list user attributes to return. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListUsers(context.Background()).Authorization(authorization).Limit(limit).Page(page).Cursor(cursor).CreatedSince(createdSince).CreatedUntil(createdUntil).UpdatedSince(updatedSince).UpdatedUntil(updatedUntil).LastLoginSince(lastLoginSince).LastLoginUntil(lastLoginUntil).Firstname(firstname).Lastname(lastname).Email(email).Username(username).Samaccountname(samaccountname).DirectoryId(directoryId).ExternalId(externalId).AppId(appId).UserIds(userIds).CustomAttributesAttributeName(customAttributesAttributeName).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **limit** | **int32** | The total number of items returned per page. The maximum limit varies between endpoints, see the relevant endpoint documentation for the specific limit. | 
 **page** | **int32** | The page number of results to return. | 
 **cursor** | **string** | Set to the value extracted from Before-Cursor or After-Cursor headers to return the previous or next page. | 
 **createdSince** | **string** | An ISO8601 timestamp value that returns all users created after a given date &amp; time. | 
 **createdUntil** | **string** | An ISO8601 timestamp value that returns all users created before a given date &amp; time. | 
 **updatedSince** | **string** | An ISO8601 timestamp value that returns all users updated after a given date &amp; time. | 
 **updatedUntil** | **string** | An ISO8601 timestamp value that returns all users updated before a given date &amp; time. | 
 **lastLoginSince** | **string** | An ISO8601 timestamp value that returns all users that logged in after a given date &amp; time. | 
 **lastLoginUntil** | **string** |  | 
 **firstname** | **string** | The first name of the user | 
 **lastname** | **string** | The last name of the user | 
 **email** | **string** | The email address of the user | 
 **username** | **string** | The username for the user | 
 **samaccountname** | **string** | The AD login name for the user | 
 **directoryId** | **string** | The ID in OneLogin of the Directory that the user belongs to | 
 **externalId** | **string** | An external identifier that has been set on the user | 
 **appId** | **string** | The ID of a OneLogin Application | 
 **userIds** | **string** | A comma separated list of OneLogin User IDs | 
 **customAttributesAttributeName** | **string** | The short name of a custom attribute. Note that the attribute name is prefixed with custom_attributes. | 
 **fields** | **string** | A comma separated list user attributes to return. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveClientApp

> RemoveClientApp(ctx, id, clientAppId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    clientAppId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RemoveClientApp(context.Background(), id, clientAppId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**clientAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleAdmins

> RemoveRoleAdmins(ctx, roleId).Authorization(authorization).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    removeRoleUsersRequest := *openapiclient.NewRemoveRoleUsersRequest() // RemoveRoleUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RemoveRoleAdmins(context.Background(), roleId).Authorization(authorization).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveRoleAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **removeRoleUsersRequest** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleUsers

> RemoveRoleUsers(ctx, roleId).Authorization(authorization).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    removeRoleUsersRequest := *openapiclient.NewRemoveRoleUsersRequest() // RemoveRoleUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RemoveRoleUsers(context.Background(), roleId).Authorization(authorization).RemoveRoleUsersRequest(removeRoleUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RemoveRoleUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **removeRoleUsersRequest** | [**RemoveRoleUsersRequest**](RemoveRoleUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

> GenerateToken400Response RevokeToken(ctx).Authorization(authorization).RevokeTokenRequest(revokeTokenRequest).Execute()



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
    authorization := "authorization_example" // string | 
    revokeTokenRequest := *openapiclient.NewRevokeTokenRequest() // RevokeTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RevokeToken(context.Background()).Authorization(authorization).RevokeTokenRequest(revokeTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeToken`: GenerateToken400Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RevokeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **revokeTokenRequest** | [**RevokeTokenRequest**](RevokeTokenRequest.md) |  | 

### Return type

[**GenerateToken400Response**](GenerateToken400Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleApps

> []SetRoleApps200ResponseInner SetRoleApps(ctx, roleId).Authorization(authorization).RequestBody(requestBody).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    requestBody := []int32{int32(123)} // []int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SetRoleApps(context.Background(), roleId).Authorization(authorization).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetRoleApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRoleApps`: []SetRoleApps200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetRoleApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **requestBody** | **[]int32** |  | 

### Return type

[**[]SetRoleApps200ResponseInner**](SetRoleApps200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEvent

> TrackEvent(ctx).Authorization(authorization).TrackEventRequest(trackEventRequest).Execute()



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
    authorization := "authorization_example" // string | 
    trackEventRequest := *openapiclient.NewTrackEventRequest("Verb_example", "Ip_example", "UserAgent_example", *openapiclient.NewRiskUser("Id_example")) // TrackEventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrackEvent(context.Background()).Authorization(authorization).TrackEventRequest(trackEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrackEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **trackEventRequest** | [**TrackEventRequest**](TrackEventRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessTokenClaim

> Id UpdateAccessTokenClaim(ctx, id, claimId).Authorization(authorization).AddAccessTokenClaimRequest(addAccessTokenClaimRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    claimId := int32(56) // int32 | 
    addAccessTokenClaimRequest := *openapiclient.NewAddAccessTokenClaimRequest() // AddAccessTokenClaimRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAccessTokenClaim(context.Background(), id, claimId).Authorization(authorization).AddAccessTokenClaimRequest(addAccessTokenClaimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccessTokenClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessTokenClaim`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccessTokenClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**claimId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessTokenClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **addAccessTokenClaimRequest** | [**AddAccessTokenClaimRequest**](AddAccessTokenClaimRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> Schema UpdateApp(ctx, appId).Authorization(authorization).Schema(schema).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    schema := *openapiclient.NewSchema() // Schema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateApp(context.Background(), appId).Authorization(authorization).Schema(schema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApp`: Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **schema** | [**Schema**](Schema.md) |  | 

### Return type

[**Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationServer

> Id UpdateAuthorizationServer(ctx, id).Authorization(authorization).CreateAuthorizationServerRequest(createAuthorizationServerRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    createAuthorizationServerRequest := *openapiclient.NewCreateAuthorizationServerRequest() // CreateAuthorizationServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAuthorizationServer(context.Background(), id).Authorization(authorization).CreateAuthorizationServerRequest(createAuthorizationServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorizationServer`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **createAuthorizationServerRequest** | [**CreateAuthorizationServerRequest**](CreateAuthorizationServerRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientApp

> ClientApp UpdateClientApp(ctx, id, clientAppId).Authorization(authorization).UpdateClientAppRequest(updateClientAppRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    clientAppId := int32(56) // int32 | 
    updateClientAppRequest := *openapiclient.NewUpdateClientAppRequest() // UpdateClientAppRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateClientApp(context.Background(), id, clientAppId).Authorization(authorization).UpdateClientAppRequest(updateClientAppRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateClientApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientApp`: ClientApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateClientApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**clientAppId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **updateClientAppRequest** | [**UpdateClientAppRequest**](UpdateClientAppRequest.md) |  | 

### Return type

[**ClientApp**](ClientApp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentVariable

> Envvar UpdateEnvironmentVariable(ctx, envvarId).Authorization(authorization).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()



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
    authorization := "authorization_example" // string | 
    envvarId := "envvarId_example" // string | Set to the id of the Hook Environment Variable that you want to fetch.
    updateEnvironmentVariableRequest := *openapiclient.NewUpdateEnvironmentVariableRequest("Value_example") // UpdateEnvironmentVariableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEnvironmentVariable(context.Background(), envvarId).Authorization(authorization).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEnvironmentVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentVariable`: Envvar
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEnvironmentVariable`: %v\n", resp)
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
 **authorization** | **string** |  | 

 **updateEnvironmentVariableRequest** | [**UpdateEnvironmentVariableRequest**](UpdateEnvironmentVariableRequest.md) |  | 

### Return type

[**Envvar**](Envvar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHook

> Hook UpdateHook(ctx, hookId).Authorization(authorization).Hook(hook).Execute()



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
    authorization := "authorization_example" // string | 
    hookId := "hookId_example" // string | Set to the id of the Hook that you want to return.
    hook := *openapiclient.NewHook("Type_example", false, int32(123), []string{"EnvVars_example"}, "Runtime_example", int32(123), map[string]interface{}(123), "Function_example") // Hook | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateHook(context.Background(), hookId).Authorization(authorization).Hook(hook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateHook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHook`: Hook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateHook`: %v\n", resp)
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
 **authorization** | **string** |  | 

 **hook** | [**Hook**](Hook.md) |  | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapping

> int32 UpdateMapping(ctx, mappingId).Authorization(authorization).Mapping(mapping).Execute()



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
    authorization := "authorization_example" // string | 
    mappingId := int32(56) // int32 | The id of the user mapping to locate.
    mapping := *openapiclient.NewMapping("Name_example", false, "Match_example", int32(123), []openapiclient.Action{*openapiclient.NewAction()}) // Mapping | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMapping(context.Background(), mappingId).Authorization(authorization).Mapping(mapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapping`: int32
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMapping`: %v\n", resp)
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
 **authorization** | **string** |  | 

 **mapping** | [**Mapping**](Mapping.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskRule

> RiskRule UpdateRiskRule(ctx, riskRuleId).Authorization(authorization).RiskRule(riskRule).Execute()



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
    authorization := "authorization_example" // string | 
    riskRuleId := "riskRuleId_example" // string | 
    riskRule := *openapiclient.NewRiskRule() // RiskRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateRiskRule(context.Background(), riskRuleId).Authorization(authorization).RiskRule(riskRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRiskRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskRule`: RiskRule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**riskRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **riskRule** | [**RiskRule**](RiskRule.md) |  | 

### Return type

[**RiskRule**](RiskRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole200Response UpdateRole(ctx, roleId).Authorization(authorization).Role(role).Execute()



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
    authorization := "authorization_example" // string | 
    roleId := int32(56) // int32 | Set to the id of the role you want to return.
    role := *openapiclient.NewRole("Name_example") // Role | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateRole(context.Background(), roleId).Authorization(authorization).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: UpdateRole200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | Set to the id of the role you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **role** | [**Role**](Role.md) |  | 

### Return type

[**UpdateRole200Response**](UpdateRole200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> RuleId UpdateRule(ctx, appId, ruleId).Authorization(authorization).Rule(rule).Execute()



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
    authorization := "authorization_example" // string | 
    appId := int32(56) // int32 | 
    ruleId := int32(56) // int32 | The id of the app rule to locate.
    rule := *openapiclient.NewRule() // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateRule(context.Background(), appId, ruleId).Authorization(authorization).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: RuleId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 
**ruleId** | **int32** | The id of the app rule to locate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **rule** | [**Rule**](Rule.md) |  | 

### Return type

[**RuleId**](RuleId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScope

> Id UpdateScope(ctx, id, scopeId).Authorization(authorization).AddScopeRequest(addScopeRequest).Execute()



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
    authorization := "authorization_example" // string | 
    id := int32(56) // int32 | 
    scopeId := int32(56) // int32 | 
    addScopeRequest := *openapiclient.NewAddScopeRequest() // AddScopeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateScope(context.Background(), id, scopeId).Authorization(authorization).AddScopeRequest(addScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScope`: Id
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**scopeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **addScopeRequest** | [**AddScopeRequest**](AddScopeRequest.md) |  | 

### Return type

[**Id**](Id.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).Authorization(authorization).User(user).Mappings(mappings).ValidatePolicy(validatePolicy).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user that you want to return.
    user := *openapiclient.NewUser() // User | 
    mappings := "mappings_example" // string | Controls how mappings will be applied to the user on creation. Defaults to async. (optional)
    validatePolicy := true // bool | Will passwords validate against the User Policy? Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateUser(context.Background(), userId).Authorization(authorization).User(user).Mappings(mappings).ValidatePolicy(validatePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user that you want to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

 **user** | [**User**](User.md) |  | 
 **mappings** | **string** | Controls how mappings will be applied to the user on creation. Defaults to async. | 
 **validatePolicy** | **bool** | Will passwords validate against the User Policy? Defaults to true. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEnrollment

> []Registration VerifyEnrollment(ctx, userId, registrationId).Authorization(authorization).VerifyEnrollmentRequest(verifyEnrollmentRequest).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    registrationId := int32(56) // int32 | Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor.
    verifyEnrollmentRequest := *openapiclient.NewVerifyEnrollmentRequest() // VerifyEnrollmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VerifyEnrollment(context.Background(), userId, registrationId).Authorization(authorization).VerifyEnrollmentRequest(verifyEnrollmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEnrollment`: []Registration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VerifyEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 
**registrationId** | **int32** | Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **verifyEnrollmentRequest** | [**VerifyEnrollmentRequest**](VerifyEnrollmentRequest.md) |  | 

### Return type

[**[]Registration**](Registration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEnrollmentVoiceProtect

> []Registration VerifyEnrollmentVoiceProtect(ctx, userId, registrationId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    registrationId := int32(56) // int32 | Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VerifyEnrollmentVoiceProtect(context.Background(), userId, registrationId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyEnrollmentVoiceProtect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEnrollmentVoiceProtect`: []Registration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VerifyEnrollmentVoiceProtect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 
**registrationId** | **int32** | Set to the uuid of the registration. This was included in the response as part of the initial request in Enroll Factor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEnrollmentVoiceProtectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**[]Registration**](Registration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFactor

> Status2 VerifyFactor(ctx, userId, verificationId).Authorization(authorization).VerifyFactorRequest(verifyFactorRequest).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    verificationId := int32(56) // int32 | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call.
    verifyFactorRequest := *openapiclient.NewVerifyFactorRequest() // VerifyFactorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VerifyFactor(context.Background(), userId, verificationId).Authorization(authorization).VerifyFactorRequest(verifyFactorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyFactor`: Status2
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VerifyFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 
**verificationId** | **int32** | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 


 **verifyFactorRequest** | [**VerifyFactorRequest**](VerifyFactorRequest.md) |  | 

### Return type

[**Status2**](Status2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFactorSaml

> VerifyFactorSaml200Response VerifyFactorSaml(ctx).Authorization(authorization).VerifyFactorSamlRequest(verifyFactorSamlRequest).Execute()



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
    authorization := "authorization_example" // string | 
    verifyFactorSamlRequest := *openapiclient.NewVerifyFactorSamlRequest("AppId_example", "DeviceId_example", "StateToken_example") // VerifyFactorSamlRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VerifyFactorSaml(context.Background()).Authorization(authorization).VerifyFactorSamlRequest(verifyFactorSamlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyFactorSaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyFactorSaml`: VerifyFactorSaml200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VerifyFactorSaml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **verifyFactorSamlRequest** | [**VerifyFactorSamlRequest**](VerifyFactorSamlRequest.md) |  | 

### Return type

[**VerifyFactorSaml200Response**](VerifyFactorSaml200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFactorVoice

> []VerifyFactorVoice200ResponseInner VerifyFactorVoice(ctx, userId, verificationId).Authorization(authorization).Execute()



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
    authorization := "authorization_example" // string | 
    userId := int32(56) // int32 | Set to the id of the user.
    verificationId := int32(56) // int32 | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VerifyFactorVoice(context.Background(), userId, verificationId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyFactorVoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyFactorVoice`: []VerifyFactorVoice200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VerifyFactorVoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | Set to the id of the user. | 
**verificationId** | **int32** | The verification_id is returned on activation of the factor or you can get the device_id using the Activate Factor API call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorVoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**[]VerifyFactorVoice200ResponseInner**](VerifyFactorVoice200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

