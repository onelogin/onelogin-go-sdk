# SendInviteLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Set to the user email address to generate an invite link. The value is case sensitive. | [optional] 
**PersonalEmail** | Pointer to **string** | To send an invite email to a different address than the one provided in email, provide it here. The invite link is sent to this address instead. | [optional] 

## Methods

### NewSendInviteLinkRequest

`func NewSendInviteLinkRequest() *SendInviteLinkRequest`

NewSendInviteLinkRequest instantiates a new SendInviteLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInviteLinkRequestWithDefaults

`func NewSendInviteLinkRequestWithDefaults() *SendInviteLinkRequest`

NewSendInviteLinkRequestWithDefaults instantiates a new SendInviteLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendInviteLinkRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendInviteLinkRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendInviteLinkRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SendInviteLinkRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPersonalEmail

`func (o *SendInviteLinkRequest) GetPersonalEmail() string`

GetPersonalEmail returns the PersonalEmail field if non-nil, zero value otherwise.

### GetPersonalEmailOk

`func (o *SendInviteLinkRequest) GetPersonalEmailOk() (*string, bool)`

GetPersonalEmailOk returns a tuple with the PersonalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEmail

`func (o *SendInviteLinkRequest) SetPersonalEmail(v string)`

SetPersonalEmail sets PersonalEmail field to given value.

### HasPersonalEmail

`func (o *SendInviteLinkRequest) HasPersonalEmail() bool`

HasPersonalEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


