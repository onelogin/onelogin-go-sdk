# GetRoleById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**[]GetRoleById200ResponseDataInner**](GetRoleById200ResponseDataInner.md) |  | [optional] 

## Methods

### NewGetRoleById200Response

`func NewGetRoleById200Response() *GetRoleById200Response`

NewGetRoleById200Response instantiates a new GetRoleById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleById200ResponseWithDefaults

`func NewGetRoleById200ResponseWithDefaults() *GetRoleById200Response`

NewGetRoleById200ResponseWithDefaults instantiates a new GetRoleById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetRoleById200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRoleById200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRoleById200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRoleById200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetRoleById200Response) GetData() []GetRoleById200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRoleById200Response) GetDataOk() (*[]GetRoleById200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRoleById200Response) SetData(v []GetRoleById200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRoleById200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


