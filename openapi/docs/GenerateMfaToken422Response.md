# GenerateMfaToken422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**GenerateMfaToken422ResponseDetails**](GenerateMfaToken422ResponseDetails.md) |  | [optional] 

## Methods

### NewGenerateMfaToken422Response

`func NewGenerateMfaToken422Response() *GenerateMfaToken422Response`

NewGenerateMfaToken422Response instantiates a new GenerateMfaToken422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateMfaToken422ResponseWithDefaults

`func NewGenerateMfaToken422ResponseWithDefaults() *GenerateMfaToken422Response`

NewGenerateMfaToken422ResponseWithDefaults instantiates a new GenerateMfaToken422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GenerateMfaToken422Response) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GenerateMfaToken422Response) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GenerateMfaToken422Response) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GenerateMfaToken422Response) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetName

`func (o *GenerateMfaToken422Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateMfaToken422Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateMfaToken422Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateMfaToken422Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *GenerateMfaToken422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GenerateMfaToken422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GenerateMfaToken422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GenerateMfaToken422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *GenerateMfaToken422Response) GetDetails() GenerateMfaToken422ResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GenerateMfaToken422Response) GetDetailsOk() (*GenerateMfaToken422ResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GenerateMfaToken422Response) SetDetails(v GenerateMfaToken422ResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GenerateMfaToken422Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


