# GetEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Pagination** | Pointer to [**GetEvents200ResponsePagination**](GetEvents200ResponsePagination.md) |  | [optional] 
**Data** | Pointer to [**[]Event**](Event.md) |  | [optional] 

## Methods

### NewGetEvents200Response

`func NewGetEvents200Response() *GetEvents200Response`

NewGetEvents200Response instantiates a new GetEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseWithDefaults

`func NewGetEvents200ResponseWithDefaults() *GetEvents200Response`

NewGetEvents200ResponseWithDefaults instantiates a new GetEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEvents200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEvents200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEvents200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEvents200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *GetEvents200Response) GetPagination() GetEvents200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetEvents200Response) GetPaginationOk() (*GetEvents200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetEvents200Response) SetPagination(v GetEvents200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetEvents200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetEvents200Response) GetData() []Event`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEvents200Response) GetDataOk() (*[]Event, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEvents200Response) SetData(v []Event)`

SetData sets Data field to given value.

### HasData

`func (o *GetEvents200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


