# ConfigurationSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignatureAlgorithm** | **string** | One of the following:   - SHA-1   - SHA-256   - SHA-348   - SHA-512 | 
**CertificateId** | **int32** | When creating apps the default certificate will be used unless the &#x60;certificate_id&#x60; attribute is applied in the &#x60;configuration&#x60; object. | 

## Methods

### NewConfigurationSaml

`func NewConfigurationSaml(signatureAlgorithm string, certificateId int32, ) *ConfigurationSaml`

NewConfigurationSaml instantiates a new ConfigurationSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationSamlWithDefaults

`func NewConfigurationSamlWithDefaults() *ConfigurationSaml`

NewConfigurationSamlWithDefaults instantiates a new ConfigurationSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatureAlgorithm

`func (o *ConfigurationSaml) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *ConfigurationSaml) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *ConfigurationSaml) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetCertificateId

`func (o *ConfigurationSaml) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ConfigurationSaml) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ConfigurationSaml) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


