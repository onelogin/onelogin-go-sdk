# MessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**AccountId** | Pointer to **int32** |  | [optional] [readonly] 
**Type** | **string** | Template type that describes the source (sms, voice, email) and purpose (registration, invite, etc) | 
**Locale** | **string** | The 2 character language locale for the template. e.g. en &#x3D; English, es &#x3D; Spanish | 
**Template** | [**MessageTemplateTemplate**](MessageTemplateTemplate.md) |  | 
**TemplateClass** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | Last time template was updated | [optional] [readonly] 
**BrandId** | Pointer to **int32** | brand id number | [optional] [readonly] 

## Methods

### NewMessageTemplate

`func NewMessageTemplate(type_ string, locale string, template MessageTemplateTemplate, ) *MessageTemplate`

NewMessageTemplate instantiates a new MessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTemplateWithDefaults

`func NewMessageTemplateWithDefaults() *MessageTemplate`

NewMessageTemplateWithDefaults instantiates a new MessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MessageTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *MessageTemplate) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MessageTemplate) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MessageTemplate) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MessageTemplate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetType

`func (o *MessageTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetLocale

`func (o *MessageTemplate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MessageTemplate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MessageTemplate) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetTemplate

`func (o *MessageTemplate) GetTemplate() MessageTemplateTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *MessageTemplate) GetTemplateOk() (*MessageTemplateTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *MessageTemplate) SetTemplate(v MessageTemplateTemplate)`

SetTemplate sets Template field to given value.


### GetTemplateClass

`func (o *MessageTemplate) GetTemplateClass() string`

GetTemplateClass returns the TemplateClass field if non-nil, zero value otherwise.

### GetTemplateClassOk

`func (o *MessageTemplate) GetTemplateClassOk() (*string, bool)`

GetTemplateClassOk returns a tuple with the TemplateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateClass

`func (o *MessageTemplate) SetTemplateClass(v string)`

SetTemplateClass sets TemplateClass field to given value.

### HasTemplateClass

`func (o *MessageTemplate) HasTemplateClass() bool`

HasTemplateClass returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MessageTemplate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageTemplate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageTemplate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessageTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBrandId

`func (o *MessageTemplate) GetBrandId() int32`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *MessageTemplate) GetBrandIdOk() (*int32, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *MessageTemplate) SetBrandId(v int32)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *MessageTemplate) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


