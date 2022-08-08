# GetRiskScore200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | A risk score 0 is low risk and 100 is the highest risk level possible. | [optional] 
**Triggers** | Pointer to **[]string** | Triggers are indicators of some of the key items that influenced the risk score. | [optional] 

## Methods

### NewGetRiskScore200Response

`func NewGetRiskScore200Response() *GetRiskScore200Response`

NewGetRiskScore200Response instantiates a new GetRiskScore200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskScore200ResponseWithDefaults

`func NewGetRiskScore200ResponseWithDefaults() *GetRiskScore200Response`

NewGetRiskScore200ResponseWithDefaults instantiates a new GetRiskScore200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *GetRiskScore200Response) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetRiskScore200Response) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetRiskScore200Response) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *GetRiskScore200Response) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTriggers

`func (o *GetRiskScore200Response) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *GetRiskScore200Response) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *GetRiskScore200Response) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *GetRiskScore200Response) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


