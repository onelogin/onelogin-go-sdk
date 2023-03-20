# GetGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**[]Group**](Group.md) |  | [optional] 

## Methods

### NewGetGroups200Response

`func NewGetGroups200Response() *GetGroups200Response`

NewGetGroups200Response instantiates a new GetGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroups200ResponseWithDefaults

`func NewGetGroups200ResponseWithDefaults() *GetGroups200Response`

NewGetGroups200ResponseWithDefaults instantiates a new GetGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetGroups200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGroups200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGroups200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGroups200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetGroups200Response) GetData() []Group`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetGroups200Response) GetDataOk() (*[]Group, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetGroups200Response) SetData(v []Group)`

SetData sets Data field to given value.

### HasData

`func (o *GetGroups200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


