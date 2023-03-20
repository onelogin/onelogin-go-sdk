# GetMFAFactors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**GetMFAFactors200ResponseData**](GetMFAFactors200ResponseData.md) |  | [optional] 

## Methods

### NewGetMFAFactors200Response

`func NewGetMFAFactors200Response() *GetMFAFactors200Response`

NewGetMFAFactors200Response instantiates a new GetMFAFactors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMFAFactors200ResponseWithDefaults

`func NewGetMFAFactors200ResponseWithDefaults() *GetMFAFactors200Response`

NewGetMFAFactors200ResponseWithDefaults instantiates a new GetMFAFactors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetMFAFactors200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMFAFactors200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMFAFactors200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMFAFactors200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetMFAFactors200Response) GetData() GetMFAFactors200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMFAFactors200Response) GetDataOk() (*GetMFAFactors200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMFAFactors200Response) SetData(v GetMFAFactors200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetMFAFactors200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


