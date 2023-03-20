# GetEventTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**[]GetEventTypes200ResponseDataInner**](GetEventTypes200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetEventTypes200Response

`func NewGetEventTypes200Response() *GetEventTypes200Response`

NewGetEventTypes200Response instantiates a new GetEventTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventTypes200ResponseWithDefaults

`func NewGetEventTypes200ResponseWithDefaults() *GetEventTypes200Response`

NewGetEventTypes200ResponseWithDefaults instantiates a new GetEventTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEventTypes200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEventTypes200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEventTypes200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEventTypes200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetEventTypes200Response) GetData() []GetEventTypes200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventTypes200Response) GetDataOk() (*[]GetEventTypes200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventTypes200Response) SetData(v []GetEventTypes200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetEventTypes200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


