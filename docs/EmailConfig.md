# EmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Email Settings server address | 
**UseTls** | Pointer to **bool** | Use TLS | [optional] [default to true]
**From** | **string** | The From email address in the message. | 
**Domain** | **string** | Domain of the From address. | 
**UserName** | Pointer to **string** | The user name of the account to authenticate with the Email Settings server. | [optional] 
**Password** | Pointer to **string** | The password of the account to authenticate with the Email Settings server. | [optional] 
**Port** | Pointer to **int32** | Defaults to 25. | [optional] [default to 25]

## Methods

### NewEmailConfig

`func NewEmailConfig(address string, from string, domain string, ) *EmailConfig`

NewEmailConfig instantiates a new EmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigWithDefaults

`func NewEmailConfigWithDefaults() *EmailConfig`

NewEmailConfigWithDefaults instantiates a new EmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmailConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailConfig) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetUseTls

`func (o *EmailConfig) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *EmailConfig) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *EmailConfig) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *EmailConfig) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetFrom

`func (o *EmailConfig) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailConfig) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailConfig) SetFrom(v string)`

SetFrom sets From field to given value.


### GetDomain

`func (o *EmailConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUserName

`func (o *EmailConfig) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EmailConfig) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EmailConfig) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *EmailConfig) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *EmailConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmailConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmailConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EmailConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *EmailConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


