# GenerateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Set to client_credentials. | [default to "client_credentials"]

## Methods

### NewGenerateTokenRequest

`func NewGenerateTokenRequest(grantType string, ) *GenerateTokenRequest`

NewGenerateTokenRequest instantiates a new GenerateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateTokenRequestWithDefaults

`func NewGenerateTokenRequestWithDefaults() *GenerateTokenRequest`

NewGenerateTokenRequestWithDefaults instantiates a new GenerateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GenerateTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GenerateTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GenerateTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


