# SsoSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataUrl** | Pointer to **string** |  | [optional] 
**AcsUrl** | Pointer to **string** |  | [optional] 
**SlsUrl** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to [**SsoSamlCertificate**](SsoSamlCertificate.md) |  | [optional] 

## Methods

### NewSsoSaml

`func NewSsoSaml() *SsoSaml`

NewSsoSaml instantiates a new SsoSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSamlWithDefaults

`func NewSsoSamlWithDefaults() *SsoSaml`

NewSsoSamlWithDefaults instantiates a new SsoSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataUrl

`func (o *SsoSaml) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *SsoSaml) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *SsoSaml) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *SsoSaml) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetAcsUrl

`func (o *SsoSaml) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *SsoSaml) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *SsoSaml) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *SsoSaml) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetSlsUrl

`func (o *SsoSaml) GetSlsUrl() string`

GetSlsUrl returns the SlsUrl field if non-nil, zero value otherwise.

### GetSlsUrlOk

`func (o *SsoSaml) GetSlsUrlOk() (*string, bool)`

GetSlsUrlOk returns a tuple with the SlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlsUrl

`func (o *SsoSaml) SetSlsUrl(v string)`

SetSlsUrl sets SlsUrl field to given value.

### HasSlsUrl

`func (o *SsoSaml) HasSlsUrl() bool`

HasSlsUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *SsoSaml) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SsoSaml) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SsoSaml) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SsoSaml) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetCertificate

`func (o *SsoSaml) GetCertificate() SsoSamlCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SsoSaml) GetCertificateOk() (*SsoSamlCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SsoSaml) SetCertificate(v SsoSamlCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SsoSaml) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


