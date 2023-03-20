# EnforcementPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireSitewideAuthentication** | Pointer to **bool** | Require user authentication to access any resource protected by this enforcement point. | [optional] 
**Conditions** | Pointer to **string** | If access is conditional, the conditions that must evaluate to true to allow access to a resource. For example, to require the user must be authenticated and have either the role Admin or User | [optional] 
**SessionExpiryFixed** | Pointer to [**ClockCounter**](ClockCounter.md) |  | [optional] 
**SessionExpiryInactivity** | Pointer to [**ClockCounter**](ClockCounter.md) |  | [optional] 
**Permissions** | Pointer to **string** | Specify to always &#x60;allow&#x60;, &#x60;deny&#x60; access to resources, of if access is &#x60;conditional&#x60;. | [optional] 
**Token** | Pointer to **string** | Can only be set on create. Access Gateway Token. | [optional] [readonly] 
**Target** | Pointer to **NullableString** | A fully-qualified URL to the internal application including scheme, authority and path. The target host authority must be an IP address, not a hostname. | [optional] 
**Resources** | Pointer to [**[]EnforcementPointResourcesInner**](EnforcementPointResourcesInner.md) | Array of resource objects | [optional] 
**ContextRoot** | Pointer to **string** | The root path to the application, often the name of the application. Can be any name, path or just a slash (“/”). The context root uniquely identifies the application within the enforcement point. | [optional] 
**UseTargetHostHeader** | Pointer to **bool** | Use the target host header as opposed to the original gateway or upstream host header. | [optional] 
**Vhost** | Pointer to **NullableString** | A comma-delimited list of one or more virtual hosts that map to applications assigned to the enforcement point. A VHOST may be a host name or an IP address. VHOST distinguish between applications that are at the same context root. | [optional] 
**LandingPage** | Pointer to **NullableString** | The location within the context root to which the browser will be redirected for IdP-initiated single sign-on. For example, the landing page might be an index page in the context root such as index.html or default.aspx. The landing page cannot begin with a slash and must use valid URL characters. | [optional] 
**CaseSensitive** | Pointer to **bool** | The URL path evaluation is case insensitive by default. Resources hosted on web servers such as Apache, NGINX and Java EE are case sensitive paths. Web servers such as Microsoft IIS are not case-sensitive. | [optional] 

## Methods

### NewEnforcementPoint

`func NewEnforcementPoint() *EnforcementPoint`

NewEnforcementPoint instantiates a new EnforcementPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnforcementPointWithDefaults

`func NewEnforcementPointWithDefaults() *EnforcementPoint`

NewEnforcementPointWithDefaults instantiates a new EnforcementPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireSitewideAuthentication

`func (o *EnforcementPoint) GetRequireSitewideAuthentication() bool`

GetRequireSitewideAuthentication returns the RequireSitewideAuthentication field if non-nil, zero value otherwise.

### GetRequireSitewideAuthenticationOk

`func (o *EnforcementPoint) GetRequireSitewideAuthenticationOk() (*bool, bool)`

GetRequireSitewideAuthenticationOk returns a tuple with the RequireSitewideAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSitewideAuthentication

`func (o *EnforcementPoint) SetRequireSitewideAuthentication(v bool)`

SetRequireSitewideAuthentication sets RequireSitewideAuthentication field to given value.

### HasRequireSitewideAuthentication

`func (o *EnforcementPoint) HasRequireSitewideAuthentication() bool`

HasRequireSitewideAuthentication returns a boolean if a field has been set.

### GetConditions

`func (o *EnforcementPoint) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EnforcementPoint) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EnforcementPoint) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *EnforcementPoint) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSessionExpiryFixed

`func (o *EnforcementPoint) GetSessionExpiryFixed() ClockCounter`

GetSessionExpiryFixed returns the SessionExpiryFixed field if non-nil, zero value otherwise.

### GetSessionExpiryFixedOk

`func (o *EnforcementPoint) GetSessionExpiryFixedOk() (*ClockCounter, bool)`

GetSessionExpiryFixedOk returns a tuple with the SessionExpiryFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiryFixed

`func (o *EnforcementPoint) SetSessionExpiryFixed(v ClockCounter)`

SetSessionExpiryFixed sets SessionExpiryFixed field to given value.

### HasSessionExpiryFixed

`func (o *EnforcementPoint) HasSessionExpiryFixed() bool`

HasSessionExpiryFixed returns a boolean if a field has been set.

### GetSessionExpiryInactivity

`func (o *EnforcementPoint) GetSessionExpiryInactivity() ClockCounter`

GetSessionExpiryInactivity returns the SessionExpiryInactivity field if non-nil, zero value otherwise.

### GetSessionExpiryInactivityOk

`func (o *EnforcementPoint) GetSessionExpiryInactivityOk() (*ClockCounter, bool)`

GetSessionExpiryInactivityOk returns a tuple with the SessionExpiryInactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiryInactivity

`func (o *EnforcementPoint) SetSessionExpiryInactivity(v ClockCounter)`

SetSessionExpiryInactivity sets SessionExpiryInactivity field to given value.

### HasSessionExpiryInactivity

`func (o *EnforcementPoint) HasSessionExpiryInactivity() bool`

HasSessionExpiryInactivity returns a boolean if a field has been set.

### GetPermissions

`func (o *EnforcementPoint) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EnforcementPoint) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EnforcementPoint) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EnforcementPoint) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetToken

`func (o *EnforcementPoint) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EnforcementPoint) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EnforcementPoint) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EnforcementPoint) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTarget

`func (o *EnforcementPoint) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EnforcementPoint) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EnforcementPoint) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EnforcementPoint) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *EnforcementPoint) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *EnforcementPoint) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetResources

`func (o *EnforcementPoint) GetResources() []EnforcementPointResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EnforcementPoint) GetResourcesOk() (*[]EnforcementPointResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EnforcementPoint) SetResources(v []EnforcementPointResourcesInner)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EnforcementPoint) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetContextRoot

`func (o *EnforcementPoint) GetContextRoot() string`

GetContextRoot returns the ContextRoot field if non-nil, zero value otherwise.

### GetContextRootOk

`func (o *EnforcementPoint) GetContextRootOk() (*string, bool)`

GetContextRootOk returns a tuple with the ContextRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRoot

`func (o *EnforcementPoint) SetContextRoot(v string)`

SetContextRoot sets ContextRoot field to given value.

### HasContextRoot

`func (o *EnforcementPoint) HasContextRoot() bool`

HasContextRoot returns a boolean if a field has been set.

### GetUseTargetHostHeader

`func (o *EnforcementPoint) GetUseTargetHostHeader() bool`

GetUseTargetHostHeader returns the UseTargetHostHeader field if non-nil, zero value otherwise.

### GetUseTargetHostHeaderOk

`func (o *EnforcementPoint) GetUseTargetHostHeaderOk() (*bool, bool)`

GetUseTargetHostHeaderOk returns a tuple with the UseTargetHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTargetHostHeader

`func (o *EnforcementPoint) SetUseTargetHostHeader(v bool)`

SetUseTargetHostHeader sets UseTargetHostHeader field to given value.

### HasUseTargetHostHeader

`func (o *EnforcementPoint) HasUseTargetHostHeader() bool`

HasUseTargetHostHeader returns a boolean if a field has been set.

### GetVhost

`func (o *EnforcementPoint) GetVhost() string`

GetVhost returns the Vhost field if non-nil, zero value otherwise.

### GetVhostOk

`func (o *EnforcementPoint) GetVhostOk() (*string, bool)`

GetVhostOk returns a tuple with the Vhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhost

`func (o *EnforcementPoint) SetVhost(v string)`

SetVhost sets Vhost field to given value.

### HasVhost

`func (o *EnforcementPoint) HasVhost() bool`

HasVhost returns a boolean if a field has been set.

### SetVhostNil

`func (o *EnforcementPoint) SetVhostNil(b bool)`

 SetVhostNil sets the value for Vhost to be an explicit nil

### UnsetVhost
`func (o *EnforcementPoint) UnsetVhost()`

UnsetVhost ensures that no value is present for Vhost, not even an explicit nil
### GetLandingPage

`func (o *EnforcementPoint) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *EnforcementPoint) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *EnforcementPoint) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *EnforcementPoint) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### SetLandingPageNil

`func (o *EnforcementPoint) SetLandingPageNil(b bool)`

 SetLandingPageNil sets the value for LandingPage to be an explicit nil

### UnsetLandingPage
`func (o *EnforcementPoint) UnsetLandingPage()`

UnsetLandingPage ensures that no value is present for LandingPage, not even an explicit nil
### GetCaseSensitive

`func (o *EnforcementPoint) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *EnforcementPoint) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *EnforcementPoint) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *EnforcementPoint) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


