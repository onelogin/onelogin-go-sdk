# TokenClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID of the claim. | [optional] 
**Label** | Pointer to **string** | The UI label for the claims. | [optional] 
**UserAttributeMappings** | Pointer to **string** | A user attribute to map values from. | [optional] 
**UserAttributeMacros** | Pointer to **string** | When &#x60;user_attribute_mappings&#x60; is set to &#x60;_macro_&#x60; this macro will be used to assign the claims value. | [optional] 
**AttributeTransformations** | Pointer to **string** | The type of transformation to perform on multi valued attributes. | [optional] 
**SkipIfBlank** | Pointer to **bool** | not used | [optional] 
**Values** | Pointer to **[]string** | Relates to Rules/Entitlements. Not supported yet. | [optional] 
**DefaultValues** | Pointer to **string** | Relates to Rules/Entitlements. Not supported yet. | [optional] 
**ProvisionedEntitlements** | Pointer to **bool** | Relates to Rules/Entitlements. Not supported yet. | [optional] 

## Methods

### NewTokenClaim

`func NewTokenClaim() *TokenClaim`

NewTokenClaim instantiates a new TokenClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenClaimWithDefaults

`func NewTokenClaimWithDefaults() *TokenClaim`

NewTokenClaimWithDefaults instantiates a new TokenClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenClaim) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenClaim) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenClaim) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TokenClaim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TokenClaim) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TokenClaim) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TokenClaim) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TokenClaim) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUserAttributeMappings

`func (o *TokenClaim) GetUserAttributeMappings() string`

GetUserAttributeMappings returns the UserAttributeMappings field if non-nil, zero value otherwise.

### GetUserAttributeMappingsOk

`func (o *TokenClaim) GetUserAttributeMappingsOk() (*string, bool)`

GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMappings

`func (o *TokenClaim) SetUserAttributeMappings(v string)`

SetUserAttributeMappings sets UserAttributeMappings field to given value.

### HasUserAttributeMappings

`func (o *TokenClaim) HasUserAttributeMappings() bool`

HasUserAttributeMappings returns a boolean if a field has been set.

### GetUserAttributeMacros

`func (o *TokenClaim) GetUserAttributeMacros() string`

GetUserAttributeMacros returns the UserAttributeMacros field if non-nil, zero value otherwise.

### GetUserAttributeMacrosOk

`func (o *TokenClaim) GetUserAttributeMacrosOk() (*string, bool)`

GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMacros

`func (o *TokenClaim) SetUserAttributeMacros(v string)`

SetUserAttributeMacros sets UserAttributeMacros field to given value.

### HasUserAttributeMacros

`func (o *TokenClaim) HasUserAttributeMacros() bool`

HasUserAttributeMacros returns a boolean if a field has been set.

### GetAttributeTransformations

`func (o *TokenClaim) GetAttributeTransformations() string`

GetAttributeTransformations returns the AttributeTransformations field if non-nil, zero value otherwise.

### GetAttributeTransformationsOk

`func (o *TokenClaim) GetAttributeTransformationsOk() (*string, bool)`

GetAttributeTransformationsOk returns a tuple with the AttributeTransformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeTransformations

`func (o *TokenClaim) SetAttributeTransformations(v string)`

SetAttributeTransformations sets AttributeTransformations field to given value.

### HasAttributeTransformations

`func (o *TokenClaim) HasAttributeTransformations() bool`

HasAttributeTransformations returns a boolean if a field has been set.

### GetSkipIfBlank

`func (o *TokenClaim) GetSkipIfBlank() bool`

GetSkipIfBlank returns the SkipIfBlank field if non-nil, zero value otherwise.

### GetSkipIfBlankOk

`func (o *TokenClaim) GetSkipIfBlankOk() (*bool, bool)`

GetSkipIfBlankOk returns a tuple with the SkipIfBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfBlank

`func (o *TokenClaim) SetSkipIfBlank(v bool)`

SetSkipIfBlank sets SkipIfBlank field to given value.

### HasSkipIfBlank

`func (o *TokenClaim) HasSkipIfBlank() bool`

HasSkipIfBlank returns a boolean if a field has been set.

### GetValues

`func (o *TokenClaim) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TokenClaim) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TokenClaim) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *TokenClaim) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDefaultValues

`func (o *TokenClaim) GetDefaultValues() string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *TokenClaim) GetDefaultValuesOk() (*string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *TokenClaim) SetDefaultValues(v string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *TokenClaim) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetProvisionedEntitlements

`func (o *TokenClaim) GetProvisionedEntitlements() bool`

GetProvisionedEntitlements returns the ProvisionedEntitlements field if non-nil, zero value otherwise.

### GetProvisionedEntitlementsOk

`func (o *TokenClaim) GetProvisionedEntitlementsOk() (*bool, bool)`

GetProvisionedEntitlementsOk returns a tuple with the ProvisionedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedEntitlements

`func (o *TokenClaim) SetProvisionedEntitlements(v bool)`

SetProvisionedEntitlements sets ProvisionedEntitlements field to given value.

### HasProvisionedEntitlements

`func (o *TokenClaim) HasProvisionedEntitlements() bool`

HasProvisionedEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


