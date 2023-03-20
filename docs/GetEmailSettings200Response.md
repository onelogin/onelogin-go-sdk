# GetEmailSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**Address** | **string** | Email Settings server address | 
**UseTls** | Pointer to **bool** | Use TLS | [optional] [default to true]
**From** | **string** | The From email address in the message. | 
**Domain** | **string** | Domain of the From address. | 
**UserName** | Pointer to **string** | The user name of the account to authenticate with the Email Settings server. | [optional] 
**Password** | Pointer to **string** | The password of the account to authenticate with the Email Settings server. | [optional] 
**Port** | Pointer to **int32** | Defaults to 25. | [optional] [default to 25]

## Methods

### NewGetEmailSettings200Response

`func NewGetEmailSettings200Response(address string, from string, domain string, ) *GetEmailSettings200Response`

NewGetEmailSettings200Response instantiates a new GetEmailSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailSettings200ResponseWithDefaults

`func NewGetEmailSettings200ResponseWithDefaults() *GetEmailSettings200Response`

NewGetEmailSettings200ResponseWithDefaults instantiates a new GetEmailSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetEmailSettings200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetEmailSettings200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetEmailSettings200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetEmailSettings200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAddress

`func (o *GetEmailSettings200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetEmailSettings200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetEmailSettings200Response) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetUseTls

`func (o *GetEmailSettings200Response) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *GetEmailSettings200Response) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *GetEmailSettings200Response) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *GetEmailSettings200Response) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetFrom

`func (o *GetEmailSettings200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetEmailSettings200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetEmailSettings200Response) SetFrom(v string)`

SetFrom sets From field to given value.


### GetDomain

`func (o *GetEmailSettings200Response) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetEmailSettings200Response) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetEmailSettings200Response) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUserName

`func (o *GetEmailSettings200Response) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetEmailSettings200Response) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetEmailSettings200Response) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetEmailSettings200Response) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *GetEmailSettings200Response) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetEmailSettings200Response) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetEmailSettings200Response) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetEmailSettings200Response) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *GetEmailSettings200Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetEmailSettings200Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetEmailSettings200Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetEmailSettings200Response) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


