# GetEnrolledFactors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**GetEnrolledFactors200ResponseData**](GetEnrolledFactors200ResponseData.md) |  | [optional] 

## Methods

### NewGetEnrolledFactors200Response

`func NewGetEnrolledFactors200Response() *GetEnrolledFactors200Response`

NewGetEnrolledFactors200Response instantiates a new GetEnrolledFactors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnrolledFactors200ResponseWithDefaults

`func NewGetEnrolledFactors200ResponseWithDefaults() *GetEnrolledFactors200Response`

NewGetEnrolledFactors200ResponseWithDefaults instantiates a new GetEnrolledFactors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEnrolledFactors200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEnrolledFactors200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEnrolledFactors200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEnrolledFactors200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetEnrolledFactors200Response) GetData() GetEnrolledFactors200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEnrolledFactors200Response) GetDataOk() (*GetEnrolledFactors200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEnrolledFactors200Response) SetData(v GetEnrolledFactors200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetEnrolledFactors200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


