# EnrollMfaFactor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Error**](Error.md) |  | [optional] 
**Data** | Pointer to [**[]GetEnrolledFactors200ResponseDataOtpDevicesInner**](GetEnrolledFactors200ResponseDataOtpDevicesInner.md) |  | [optional] 

## Methods

### NewEnrollMfaFactor200Response

`func NewEnrollMfaFactor200Response() *EnrollMfaFactor200Response`

NewEnrollMfaFactor200Response instantiates a new EnrollMfaFactor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollMfaFactor200ResponseWithDefaults

`func NewEnrollMfaFactor200ResponseWithDefaults() *EnrollMfaFactor200Response`

NewEnrollMfaFactor200ResponseWithDefaults instantiates a new EnrollMfaFactor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnrollMfaFactor200Response) GetStatus() Error`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrollMfaFactor200Response) GetStatusOk() (*Error, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrollMfaFactor200Response) SetStatus(v Error)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnrollMfaFactor200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *EnrollMfaFactor200Response) GetData() []GetEnrolledFactors200ResponseDataOtpDevicesInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnrollMfaFactor200Response) GetDataOk() (*[]GetEnrolledFactors200ResponseDataOtpDevicesInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnrollMfaFactor200Response) SetData(v []GetEnrolledFactors200ResponseDataOtpDevicesInner)`

SetData sets Data field to given value.

### HasData

`func (o *EnrollMfaFactor200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


