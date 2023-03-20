# GenerateMFAtoken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaToken** | Pointer to **string** | Token can function as a temporary MFA token. It can be used to authenticate for any app when valid. | [optional] 
**Resuable** | Pointer to **bool** | true indcates the token can be used multiple times, until it expires. false indicates the token is invalid after a single use or once it expires. Defaults to false. | [optional] 
**ExpiresAt** | Pointer to **string** | Defines the expiration time and date for the token. Format is UTC time. | [optional] 

## Methods

### NewGenerateMFAtoken200Response

`func NewGenerateMFAtoken200Response() *GenerateMFAtoken200Response`

NewGenerateMFAtoken200Response instantiates a new GenerateMFAtoken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateMFAtoken200ResponseWithDefaults

`func NewGenerateMFAtoken200ResponseWithDefaults() *GenerateMFAtoken200Response`

NewGenerateMFAtoken200ResponseWithDefaults instantiates a new GenerateMFAtoken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaToken

`func (o *GenerateMFAtoken200Response) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *GenerateMFAtoken200Response) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *GenerateMFAtoken200Response) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *GenerateMFAtoken200Response) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.

### GetResuable

`func (o *GenerateMFAtoken200Response) GetResuable() bool`

GetResuable returns the Resuable field if non-nil, zero value otherwise.

### GetResuableOk

`func (o *GenerateMFAtoken200Response) GetResuableOk() (*bool, bool)`

GetResuableOk returns a tuple with the Resuable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResuable

`func (o *GenerateMFAtoken200Response) SetResuable(v bool)`

SetResuable sets Resuable field to given value.

### HasResuable

`func (o *GenerateMFAtoken200Response) HasResuable() bool`

HasResuable returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GenerateMFAtoken200Response) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GenerateMFAtoken200Response) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GenerateMFAtoken200Response) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GenerateMFAtoken200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


