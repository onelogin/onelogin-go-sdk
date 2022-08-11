# ListAccessTokenClaims200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**UserAttributeMappings** | Pointer to **string** |  | [optional] 
**UserAttributeMacros** | Pointer to **string** |  | [optional] 
**AttributesTransformations** | Pointer to **string** |  | [optional] 
**SkipIfBlank** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**DefaultValues** | Pointer to **string** |  | [optional] 
**ProvisionedEntitlements** | Pointer to **bool** |  | [optional] 

## Methods

### NewListAccessTokenClaims200ResponseInner

`func NewListAccessTokenClaims200ResponseInner() *ListAccessTokenClaims200ResponseInner`

NewListAccessTokenClaims200ResponseInner instantiates a new ListAccessTokenClaims200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessTokenClaims200ResponseInnerWithDefaults

`func NewListAccessTokenClaims200ResponseInnerWithDefaults() *ListAccessTokenClaims200ResponseInner`

NewListAccessTokenClaims200ResponseInnerWithDefaults instantiates a new ListAccessTokenClaims200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAccessTokenClaims200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAccessTokenClaims200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAccessTokenClaims200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListAccessTokenClaims200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ListAccessTokenClaims200ResponseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListAccessTokenClaims200ResponseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListAccessTokenClaims200ResponseInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListAccessTokenClaims200ResponseInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUserAttributeMappings

`func (o *ListAccessTokenClaims200ResponseInner) GetUserAttributeMappings() string`

GetUserAttributeMappings returns the UserAttributeMappings field if non-nil, zero value otherwise.

### GetUserAttributeMappingsOk

`func (o *ListAccessTokenClaims200ResponseInner) GetUserAttributeMappingsOk() (*string, bool)`

GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMappings

`func (o *ListAccessTokenClaims200ResponseInner) SetUserAttributeMappings(v string)`

SetUserAttributeMappings sets UserAttributeMappings field to given value.

### HasUserAttributeMappings

`func (o *ListAccessTokenClaims200ResponseInner) HasUserAttributeMappings() bool`

HasUserAttributeMappings returns a boolean if a field has been set.

### GetUserAttributeMacros

`func (o *ListAccessTokenClaims200ResponseInner) GetUserAttributeMacros() string`

GetUserAttributeMacros returns the UserAttributeMacros field if non-nil, zero value otherwise.

### GetUserAttributeMacrosOk

`func (o *ListAccessTokenClaims200ResponseInner) GetUserAttributeMacrosOk() (*string, bool)`

GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMacros

`func (o *ListAccessTokenClaims200ResponseInner) SetUserAttributeMacros(v string)`

SetUserAttributeMacros sets UserAttributeMacros field to given value.

### HasUserAttributeMacros

`func (o *ListAccessTokenClaims200ResponseInner) HasUserAttributeMacros() bool`

HasUserAttributeMacros returns a boolean if a field has been set.

### GetAttributesTransformations

`func (o *ListAccessTokenClaims200ResponseInner) GetAttributesTransformations() string`

GetAttributesTransformations returns the AttributesTransformations field if non-nil, zero value otherwise.

### GetAttributesTransformationsOk

`func (o *ListAccessTokenClaims200ResponseInner) GetAttributesTransformationsOk() (*string, bool)`

GetAttributesTransformationsOk returns a tuple with the AttributesTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesTransformations

`func (o *ListAccessTokenClaims200ResponseInner) SetAttributesTransformations(v string)`

SetAttributesTransformations sets AttributesTransformations field to given value.

### HasAttributesTransformations

`func (o *ListAccessTokenClaims200ResponseInner) HasAttributesTransformations() bool`

HasAttributesTransformations returns a boolean if a field has been set.

### GetSkipIfBlank

`func (o *ListAccessTokenClaims200ResponseInner) GetSkipIfBlank() bool`

GetSkipIfBlank returns the SkipIfBlank field if non-nil, zero value otherwise.

### GetSkipIfBlankOk

`func (o *ListAccessTokenClaims200ResponseInner) GetSkipIfBlankOk() (*bool, bool)`

GetSkipIfBlankOk returns a tuple with the SkipIfBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfBlank

`func (o *ListAccessTokenClaims200ResponseInner) SetSkipIfBlank(v bool)`

SetSkipIfBlank sets SkipIfBlank field to given value.

### HasSkipIfBlank

`func (o *ListAccessTokenClaims200ResponseInner) HasSkipIfBlank() bool`

HasSkipIfBlank returns a boolean if a field has been set.

### GetValues

`func (o *ListAccessTokenClaims200ResponseInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListAccessTokenClaims200ResponseInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListAccessTokenClaims200ResponseInner) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ListAccessTokenClaims200ResponseInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDefaultValues

`func (o *ListAccessTokenClaims200ResponseInner) GetDefaultValues() string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *ListAccessTokenClaims200ResponseInner) GetDefaultValuesOk() (*string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *ListAccessTokenClaims200ResponseInner) SetDefaultValues(v string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *ListAccessTokenClaims200ResponseInner) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetProvisionedEntitlements

`func (o *ListAccessTokenClaims200ResponseInner) GetProvisionedEntitlements() bool`

GetProvisionedEntitlements returns the ProvisionedEntitlements field if non-nil, zero value otherwise.

### GetProvisionedEntitlementsOk

`func (o *ListAccessTokenClaims200ResponseInner) GetProvisionedEntitlementsOk() (*bool, bool)`

GetProvisionedEntitlementsOk returns a tuple with the ProvisionedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedEntitlements

`func (o *ListAccessTokenClaims200ResponseInner) SetProvisionedEntitlements(v bool)`

SetProvisionedEntitlements sets ProvisionedEntitlements field to given value.

### HasProvisionedEntitlements

`func (o *ListAccessTokenClaims200ResponseInner) HasProvisionedEntitlements() bool`

HasProvisionedEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


