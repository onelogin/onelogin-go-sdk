# VerFactorSaml200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Provides the SAML assertion. | [optional] 
**Message** | Pointer to **string** | Plain text description describing the outcome of the response. | [optional] 

## Methods

### NewVerFactorSaml200Response

`func NewVerFactorSaml200Response() *VerFactorSaml200Response`

NewVerFactorSaml200Response instantiates a new VerFactorSaml200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerFactorSaml200ResponseWithDefaults

`func NewVerFactorSaml200ResponseWithDefaults() *VerFactorSaml200Response`

NewVerFactorSaml200ResponseWithDefaults instantiates a new VerFactorSaml200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VerFactorSaml200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VerFactorSaml200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VerFactorSaml200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *VerFactorSaml200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *VerFactorSaml200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VerFactorSaml200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VerFactorSaml200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VerFactorSaml200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


