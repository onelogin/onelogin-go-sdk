# CreateDeviceVerification201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int32** | Specifies the factor to be verified. | [optional] 
**DisplayName** | Pointer to **string** | A name for the users device | [optional] 
**ExpiresAt** | Pointer to **string** | A short lived token that is required to Verify the Factor. This token expires based on the expires_in parameter passed in. | [optional] 
**RedirectTo** | Pointer to **string** | Only applies to Email MagicLink factor. Redirects MagicLink success page to specified URL after 2 seconds. Email must already be configured by the user. | [optional] 
**UserDisplayName** | Pointer to **string** | Authentication factor display name assigned by users when they register the device. | [optional] 
**Id** | Pointer to **string** | Registration identifier. | [optional] 
**TypeDisplayName** | Pointer to **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] 
**AuthFactorName** | Pointer to **string** | Authentication factor name, as it appears to administrators in OneLogin. | [optional] 

## Methods

### NewCreateDeviceVerification201Response

`func NewCreateDeviceVerification201Response() *CreateDeviceVerification201Response`

NewCreateDeviceVerification201Response instantiates a new CreateDeviceVerification201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceVerification201ResponseWithDefaults

`func NewCreateDeviceVerification201ResponseWithDefaults() *CreateDeviceVerification201Response`

NewCreateDeviceVerification201ResponseWithDefaults instantiates a new CreateDeviceVerification201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *CreateDeviceVerification201Response) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CreateDeviceVerification201Response) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CreateDeviceVerification201Response) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *CreateDeviceVerification201Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateDeviceVerification201Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateDeviceVerification201Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateDeviceVerification201Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateDeviceVerification201Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateDeviceVerification201Response) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateDeviceVerification201Response) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateDeviceVerification201Response) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateDeviceVerification201Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRedirectTo

`func (o *CreateDeviceVerification201Response) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *CreateDeviceVerification201Response) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *CreateDeviceVerification201Response) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *CreateDeviceVerification201Response) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *CreateDeviceVerification201Response) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *CreateDeviceVerification201Response) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *CreateDeviceVerification201Response) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *CreateDeviceVerification201Response) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetId

`func (o *CreateDeviceVerification201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateDeviceVerification201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateDeviceVerification201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateDeviceVerification201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *CreateDeviceVerification201Response) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *CreateDeviceVerification201Response) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *CreateDeviceVerification201Response) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *CreateDeviceVerification201Response) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *CreateDeviceVerification201Response) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *CreateDeviceVerification201Response) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *CreateDeviceVerification201Response) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *CreateDeviceVerification201Response) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


