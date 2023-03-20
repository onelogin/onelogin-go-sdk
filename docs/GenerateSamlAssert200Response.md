# GenerateSamlAssert200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGenerateSamlAssert200Response

`func NewGenerateSamlAssert200Response() *GenerateSamlAssert200Response`

NewGenerateSamlAssert200Response instantiates a new GenerateSamlAssert200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSamlAssert200ResponseWithDefaults

`func NewGenerateSamlAssert200ResponseWithDefaults() *GenerateSamlAssert200Response`

NewGenerateSamlAssert200ResponseWithDefaults instantiates a new GenerateSamlAssert200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerateSamlAssert200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateSamlAssert200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerateSamlAssert200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerateSamlAssert200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GenerateSamlAssert200Response) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GenerateSamlAssert200Response) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GenerateSamlAssert200Response) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GenerateSamlAssert200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


