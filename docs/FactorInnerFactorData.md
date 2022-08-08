# FactorInnerFactorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationToken** | Pointer to **string** | The token which can be used to verify the factor registration. | [optional] 
**TotpUrl** | Pointer to **string** | OTP Url that can be leveraged for any authenticator that supports the key uri format. | [optional] 

## Methods

### NewFactorInnerFactorData

`func NewFactorInnerFactorData() *FactorInnerFactorData`

NewFactorInnerFactorData instantiates a new FactorInnerFactorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactorInnerFactorDataWithDefaults

`func NewFactorInnerFactorDataWithDefaults() *FactorInnerFactorData`

NewFactorInnerFactorDataWithDefaults instantiates a new FactorInnerFactorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationToken

`func (o *FactorInnerFactorData) GetVerificationToken() string`

GetVerificationToken returns the VerificationToken field if non-nil, zero value otherwise.

### GetVerificationTokenOk

`func (o *FactorInnerFactorData) GetVerificationTokenOk() (*string, bool)`

GetVerificationTokenOk returns a tuple with the VerificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationToken

`func (o *FactorInnerFactorData) SetVerificationToken(v string)`

SetVerificationToken sets VerificationToken field to given value.

### HasVerificationToken

`func (o *FactorInnerFactorData) HasVerificationToken() bool`

HasVerificationToken returns a boolean if a field has been set.

### GetTotpUrl

`func (o *FactorInnerFactorData) GetTotpUrl() string`

GetTotpUrl returns the TotpUrl field if non-nil, zero value otherwise.

### GetTotpUrlOk

`func (o *FactorInnerFactorData) GetTotpUrlOk() (*string, bool)`

GetTotpUrlOk returns a tuple with the TotpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUrl

`func (o *FactorInnerFactorData) SetTotpUrl(v string)`

SetTotpUrl sets TotpUrl field to given value.

### HasTotpUrl

`func (o *FactorInnerFactorData) HasTotpUrl() bool`

HasTotpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


