# GetRoleByName200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Pagination** | Pointer to [**GetRoleByName200ResponsePagination**](GetRoleByName200ResponsePagination.md) |  | [optional] 
**Data** | Pointer to [**[]GetRoleByName200ResponseDataInner**](GetRoleByName200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetRoleByName200Response

`func NewGetRoleByName200Response() *GetRoleByName200Response`

NewGetRoleByName200Response instantiates a new GetRoleByName200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleByName200ResponseWithDefaults

`func NewGetRoleByName200ResponseWithDefaults() *GetRoleByName200Response`

NewGetRoleByName200ResponseWithDefaults instantiates a new GetRoleByName200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetRoleByName200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRoleByName200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRoleByName200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRoleByName200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *GetRoleByName200Response) GetPagination() GetRoleByName200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetRoleByName200Response) GetPaginationOk() (*GetRoleByName200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetRoleByName200Response) SetPagination(v GetRoleByName200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetRoleByName200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetRoleByName200Response) GetData() []GetRoleByName200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRoleByName200Response) GetDataOk() (*[]GetRoleByName200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRoleByName200Response) SetData(v []GetRoleByName200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRoleByName200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


