# Locale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | locale string | [optional] 
**IsDefault** | Pointer to **bool** | indicator if language is default | [optional] 

## Methods

### NewLocale

`func NewLocale() *Locale`

NewLocale instantiates a new Locale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocaleWithDefaults

`func NewLocaleWithDefaults() *Locale`

NewLocaleWithDefaults instantiates a new Locale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *Locale) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Locale) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Locale) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Locale) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIsDefault

`func (o *Locale) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Locale) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Locale) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Locale) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


