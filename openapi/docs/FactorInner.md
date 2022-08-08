# FactorInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | MFA device identifier. | [optional] 
**Status** | Pointer to **string** | accepted : factor has been verified. pending: registered but has not been verified. | [optional] 
**Default** | Pointer to **bool** | True &#x3D; is user&#39;s default MFA device for OneLogin. | [optional] 
**AuthFactorName** | Pointer to **string** | \&quot;Official\&quot; authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**TypeDisplayName** | Pointer to **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] 
**UserDisplayName** | Pointer to **string** | Authentication factor display name assigned by users when they enroll the device. | [optional] 
**ExpiresAt** | Pointer to **string** | A short lived token that is required to Verify the Factor. This token expires based on the expires_in parameter passed in. | [optional] 
**FactorData** | Pointer to [**FactorInnerFactorData**](FactorInnerFactorData.md) |  | [optional] 

## Methods

### NewFactorInner

`func NewFactorInner() *FactorInner`

NewFactorInner instantiates a new FactorInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactorInnerWithDefaults

`func NewFactorInnerWithDefaults() *FactorInner`

NewFactorInnerWithDefaults instantiates a new FactorInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FactorInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FactorInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FactorInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FactorInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FactorInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FactorInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FactorInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FactorInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefault

`func (o *FactorInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FactorInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FactorInner) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FactorInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *FactorInner) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *FactorInner) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *FactorInner) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *FactorInner) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *FactorInner) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *FactorInner) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *FactorInner) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *FactorInner) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *FactorInner) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *FactorInner) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *FactorInner) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *FactorInner) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *FactorInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FactorInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FactorInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FactorInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorData

`func (o *FactorInner) GetFactorData() FactorInnerFactorData`

GetFactorData returns the FactorData field if non-nil, zero value otherwise.

### GetFactorDataOk

`func (o *FactorInner) GetFactorDataOk() (*FactorInnerFactorData, bool)`

GetFactorDataOk returns a tuple with the FactorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorData

`func (o *FactorInner) SetFactorData(v FactorInnerFactorData)`

SetFactorData sets FactorData field to given value.

### HasFactorData

`func (o *FactorInner) HasFactorData() bool`

HasFactorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


