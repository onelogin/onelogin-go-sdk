# RiskUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the user. | 
**Name** | Pointer to **string** | A name for the user. | [optional] 
**Authenticated** | Pointer to **bool** | Indicates if the metadata supplied in this event should be considered as trusted for the user. | [optional] [default to false]

## Methods

### NewRiskUser

`func NewRiskUser(id string, ) *RiskUser`

NewRiskUser instantiates a new RiskUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskUserWithDefaults

`func NewRiskUserWithDefaults() *RiskUser`

NewRiskUserWithDefaults instantiates a new RiskUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RiskUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticated

`func (o *RiskUser) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *RiskUser) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *RiskUser) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *RiskUser) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


