# BrandReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Brand’s unique ID in OneLogin. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the brand is enabled or not. | [optional] 
**Name** | Pointer to **string** | Brand name for humans. This isn’t related to subdomains. | [optional] 

## Methods

### NewBrandReq

`func NewBrandReq() *BrandReq`

NewBrandReq instantiates a new BrandReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandReqWithDefaults

`func NewBrandReqWithDefaults() *BrandReq`

NewBrandReqWithDefaults instantiates a new BrandReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrandReq) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandReq) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandReq) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BrandReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *BrandReq) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BrandReq) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BrandReq) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BrandReq) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *BrandReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandReq) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


