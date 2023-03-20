# AltErr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | HTTP error code https://developer.mozilla.org/en-US/docs/Web/HTTP/Status | [optional] 
**Name** | Pointer to **string** | Error Code Name | [optional] 
**Message** | Pointer to **string** | Description of Error | [optional] 

## Methods

### NewAltErr

`func NewAltErr() *AltErr`

NewAltErr instantiates a new AltErr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAltErrWithDefaults

`func NewAltErrWithDefaults() *AltErr`

NewAltErrWithDefaults instantiates a new AltErr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *AltErr) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AltErr) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AltErr) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AltErr) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetName

`func (o *AltErr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AltErr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AltErr) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AltErr) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *AltErr) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AltErr) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AltErr) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AltErr) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


