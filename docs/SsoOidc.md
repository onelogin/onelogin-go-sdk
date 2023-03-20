# SsoOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | OIDC: The OpenId Connect Client Id.  Note that client_secret is only returned after Creating an App | [optional] 

## Methods

### NewSsoOidc

`func NewSsoOidc() *SsoOidc`

NewSsoOidc instantiates a new SsoOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoOidcWithDefaults

`func NewSsoOidcWithDefaults() *SsoOidc`

NewSsoOidcWithDefaults instantiates a new SsoOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SsoOidc) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SsoOidc) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SsoOidc) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SsoOidc) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


