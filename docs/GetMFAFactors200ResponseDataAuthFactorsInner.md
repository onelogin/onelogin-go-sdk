# GetMFAFactors200ResponseDataAuthFactorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Official authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**FactorId** | Pointer to **int32** | Identifier for the factor which will be used for user enrollment | [optional] 

## Methods

### NewGetMFAFactors200ResponseDataAuthFactorsInner

`func NewGetMFAFactors200ResponseDataAuthFactorsInner() *GetMFAFactors200ResponseDataAuthFactorsInner`

NewGetMFAFactors200ResponseDataAuthFactorsInner instantiates a new GetMFAFactors200ResponseDataAuthFactorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMFAFactors200ResponseDataAuthFactorsInnerWithDefaults

`func NewGetMFAFactors200ResponseDataAuthFactorsInnerWithDefaults() *GetMFAFactors200ResponseDataAuthFactorsInner`

NewGetMFAFactors200ResponseDataAuthFactorsInnerWithDefaults instantiates a new GetMFAFactors200ResponseDataAuthFactorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFactorId

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.

### HasFactorId

`func (o *GetMFAFactors200ResponseDataAuthFactorsInner) HasFactorId() bool`

HasFactorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


