# ErrorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ErrorStatusErrorsInner**](ErrorStatusErrorsInner.md) |  | [optional] 

## Methods

### NewErrorStatus

`func NewErrorStatus() *ErrorStatus`

NewErrorStatus instantiates a new ErrorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorStatusWithDefaults

`func NewErrorStatusWithDefaults() *ErrorStatus`

NewErrorStatusWithDefaults instantiates a new ErrorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorStatus) GetErrors() []ErrorStatusErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorStatus) GetErrorsOk() (*[]ErrorStatusErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorStatus) SetErrors(v []ErrorStatusErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


