# GetCustomAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to **[][]string** | Provides a list of custom attribute fields (also known as custom user fields) that are available for your account. The values returned correspond to the values you provided in the Shortname field when you defined the custom user field. For details about defining custom user fields, see Custom User Fields. | [optional] 

## Methods

### NewGetCustomAttributes200Response

`func NewGetCustomAttributes200Response() *GetCustomAttributes200Response`

NewGetCustomAttributes200Response instantiates a new GetCustomAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomAttributes200ResponseWithDefaults

`func NewGetCustomAttributes200ResponseWithDefaults() *GetCustomAttributes200Response`

NewGetCustomAttributes200ResponseWithDefaults instantiates a new GetCustomAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCustomAttributes200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCustomAttributes200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCustomAttributes200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCustomAttributes200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetCustomAttributes200Response) GetData() [][]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomAttributes200Response) GetDataOk() (*[][]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomAttributes200Response) SetData(v [][]string)`

SetData sets Data field to given value.

### HasData

`func (o *GetCustomAttributes200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


