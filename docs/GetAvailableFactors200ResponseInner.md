# GetAvailableFactors200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | Pointer to **int32** | Identifier for the factor which will be used for user enrollment | [optional] 
**Name** | Pointer to **string** | Authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**AuthFactorName** | Pointer to **string** | Internal use only | [optional] 

## Methods

### NewGetAvailableFactors200ResponseInner

`func NewGetAvailableFactors200ResponseInner() *GetAvailableFactors200ResponseInner`

NewGetAvailableFactors200ResponseInner instantiates a new GetAvailableFactors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAvailableFactors200ResponseInnerWithDefaults

`func NewGetAvailableFactors200ResponseInnerWithDefaults() *GetAvailableFactors200ResponseInner`

NewGetAvailableFactors200ResponseInnerWithDefaults instantiates a new GetAvailableFactors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *GetAvailableFactors200ResponseInner) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *GetAvailableFactors200ResponseInner) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *GetAvailableFactors200ResponseInner) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.

### HasFactorId

`func (o *GetAvailableFactors200ResponseInner) HasFactorId() bool`

HasFactorId returns a boolean if a field has been set.

### GetName

`func (o *GetAvailableFactors200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAvailableFactors200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAvailableFactors200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAvailableFactors200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *GetAvailableFactors200ResponseInner) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *GetAvailableFactors200ResponseInner) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *GetAvailableFactors200ResponseInner) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *GetAvailableFactors200ResponseInner) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


