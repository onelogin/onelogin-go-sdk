# RevokeTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Set to the access token you want to revoke. This access token must have been generated using the client_id and client_secret provided in the Authorization header. | 

## Methods

### NewRevokeTokensRequest

`func NewRevokeTokensRequest(accessToken string, ) *RevokeTokensRequest`

NewRevokeTokensRequest instantiates a new RevokeTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeTokensRequestWithDefaults

`func NewRevokeTokensRequestWithDefaults() *RevokeTokensRequest`

NewRevokeTokensRequestWithDefaults instantiates a new RevokeTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RevokeTokensRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RevokeTokensRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RevokeTokensRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


