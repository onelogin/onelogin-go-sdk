# GetRiskScores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scores** | Pointer to [**GetRiskScores200ResponseScores**](GetRiskScores200ResponseScores.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetRiskScores200Response

`func NewGetRiskScores200Response() *GetRiskScores200Response`

NewGetRiskScores200Response instantiates a new GetRiskScores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskScores200ResponseWithDefaults

`func NewGetRiskScores200ResponseWithDefaults() *GetRiskScores200Response`

NewGetRiskScores200ResponseWithDefaults instantiates a new GetRiskScores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScores

`func (o *GetRiskScores200Response) GetScores() GetRiskScores200ResponseScores`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *GetRiskScores200Response) GetScoresOk() (*GetRiskScores200ResponseScores, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *GetRiskScores200Response) SetScores(v GetRiskScores200ResponseScores)`

SetScores sets Scores field to given value.

### HasScores

`func (o *GetRiskScores200Response) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetTotal

`func (o *GetRiskScores200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRiskScores200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRiskScores200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetRiskScores200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


