# GetAuthFactors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | Pointer to **int32** | Identifier for the factor which will be used for user enrollment | [optional] 
**Name** | Pointer to **string** | Authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**AuthFactorName** | Pointer to **string** | Internal use only | [optional] 

## Methods

### NewGetAuthFactors200Response

`func NewGetAuthFactors200Response() *GetAuthFactors200Response`

NewGetAuthFactors200Response instantiates a new GetAuthFactors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthFactors200ResponseWithDefaults

`func NewGetAuthFactors200ResponseWithDefaults() *GetAuthFactors200Response`

NewGetAuthFactors200ResponseWithDefaults instantiates a new GetAuthFactors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *GetAuthFactors200Response) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *GetAuthFactors200Response) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *GetAuthFactors200Response) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.

### HasFactorId

`func (o *GetAuthFactors200Response) HasFactorId() bool`

HasFactorId returns a boolean if a field has been set.

### GetName

`func (o *GetAuthFactors200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAuthFactors200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAuthFactors200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAuthFactors200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *GetAuthFactors200Response) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *GetAuthFactors200Response) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *GetAuthFactors200Response) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *GetAuthFactors200Response) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


