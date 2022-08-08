# GetRateLimit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Data** | Pointer to [**GetRateLimit200ResponseData**](GetRateLimit200ResponseData.md) |  | [optional] 

## Methods

### NewGetRateLimit200Response

`func NewGetRateLimit200Response() *GetRateLimit200Response`

NewGetRateLimit200Response instantiates a new GetRateLimit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRateLimit200ResponseWithDefaults

`func NewGetRateLimit200ResponseWithDefaults() *GetRateLimit200Response`

NewGetRateLimit200ResponseWithDefaults instantiates a new GetRateLimit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetRateLimit200Response) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRateLimit200Response) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRateLimit200Response) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRateLimit200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetRateLimit200Response) GetData() GetRateLimit200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRateLimit200Response) GetDataOk() (*GetRateLimit200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRateLimit200Response) SetData(v GetRateLimit200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetRateLimit200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


