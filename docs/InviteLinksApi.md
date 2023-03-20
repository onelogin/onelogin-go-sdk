# \InviteLinksApi

All URIs are relative to *https://your-api-subdomain.onelogin.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInviteLink**](InviteLinksApi.md#GetInviteLink) | **Post** /api/1/invites/get_invite_link | Generate Invite Link
[**SendInviteLink**](InviteLinksApi.md#SendInviteLink) | **Post** /api/1/invites/send_invite_link | Send  Invite Link



## GetInviteLink

> GetInviteLink200Response GetInviteLink(ctx).GetInviteLinkRequest(getInviteLinkRequest).Execute()

Generate Invite Link



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
    getInviteLinkRequest := *openapiclient.NewGetInviteLinkRequest() // GetInviteLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InviteLinksApi.GetInviteLink(context.Background()).GetInviteLinkRequest(getInviteLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InviteLinksApi.GetInviteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInviteLink`: GetInviteLink200Response
    fmt.Fprintf(os.Stdout, "Response from `InviteLinksApi.GetInviteLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInviteLinkRequest** | [**GetInviteLinkRequest**](GetInviteLinkRequest.md) |  | 

### Return type

[**GetInviteLink200Response**](GetInviteLink200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInviteLink

> SendInviteLink200Response SendInviteLink(ctx).SendInviteLinkRequest(sendInviteLinkRequest).Execute()

Send  Invite Link



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
    sendInviteLinkRequest := *openapiclient.NewSendInviteLinkRequest() // SendInviteLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InviteLinksApi.SendInviteLink(context.Background()).SendInviteLinkRequest(sendInviteLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InviteLinksApi.SendInviteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInviteLink`: SendInviteLink200Response
    fmt.Fprintf(os.Stdout, "Response from `InviteLinksApi.SendInviteLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendInviteLinkRequest** | [**SendInviteLinkRequest**](SendInviteLinkRequest.md) |  | 

### Return type

[**SendInviteLink200Response**](SendInviteLink200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

