# SiteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Ip4** | Pointer to **NullableString** |  | [optional] 
**Ip6** | Pointer to **NullableString** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Cert** | Pointer to **NullableString** |  | [optional] 
**Redirect** | Pointer to **bool** | Automatically redirect to https:// | [optional] 
**GenerateLe** | Pointer to **bool** | Automatically issue Lets Encrypt certificate for the domains on this site? | [optional] 
**LeHttpChallengeTokens** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**PrimaryDomain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSiteUpdate

`func NewSiteUpdate(id string, ) *SiteUpdate`

NewSiteUpdate instantiates a new SiteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteUpdateWithDefaults

`func NewSiteUpdateWithDefaults() *SiteUpdate`

NewSiteUpdateWithDefaults instantiates a new SiteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SiteUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp4

`func (o *SiteUpdate) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *SiteUpdate) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *SiteUpdate) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.

### HasIp4

`func (o *SiteUpdate) HasIp4() bool`

HasIp4 returns a boolean if a field has been set.

### SetIp4Nil

`func (o *SiteUpdate) SetIp4Nil(b bool)`

 SetIp4Nil sets the value for Ip4 to be an explicit nil

### UnsetIp4
`func (o *SiteUpdate) UnsetIp4()`

UnsetIp4 ensures that no value is present for Ip4, not even an explicit nil
### GetIp6

`func (o *SiteUpdate) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *SiteUpdate) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *SiteUpdate) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *SiteUpdate) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### SetIp6Nil

`func (o *SiteUpdate) SetIp6Nil(b bool)`

 SetIp6Nil sets the value for Ip6 to be an explicit nil

### UnsetIp6
`func (o *SiteUpdate) UnsetIp6()`

UnsetIp6 ensures that no value is present for Ip6, not even an explicit nil
### GetDomains

`func (o *SiteUpdate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SiteUpdate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SiteUpdate) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *SiteUpdate) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetRoutes

`func (o *SiteUpdate) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SiteUpdate) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SiteUpdate) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *SiteUpdate) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetCert

`func (o *SiteUpdate) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SiteUpdate) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *SiteUpdate) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *SiteUpdate) HasCert() bool`

HasCert returns a boolean if a field has been set.

### SetCertNil

`func (o *SiteUpdate) SetCertNil(b bool)`

 SetCertNil sets the value for Cert to be an explicit nil

### UnsetCert
`func (o *SiteUpdate) UnsetCert()`

UnsetCert ensures that no value is present for Cert, not even an explicit nil
### GetRedirect

`func (o *SiteUpdate) GetRedirect() bool`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SiteUpdate) GetRedirectOk() (*bool, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SiteUpdate) SetRedirect(v bool)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *SiteUpdate) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetGenerateLe

`func (o *SiteUpdate) GetGenerateLe() bool`

GetGenerateLe returns the GenerateLe field if non-nil, zero value otherwise.

### GetGenerateLeOk

`func (o *SiteUpdate) GetGenerateLeOk() (*bool, bool)`

GetGenerateLeOk returns a tuple with the GenerateLe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLe

`func (o *SiteUpdate) SetGenerateLe(v bool)`

SetGenerateLe sets GenerateLe field to given value.

### HasGenerateLe

`func (o *SiteUpdate) HasGenerateLe() bool`

HasGenerateLe returns a boolean if a field has been set.

### GetLeHttpChallengeTokens

`func (o *SiteUpdate) GetLeHttpChallengeTokens() []string`

GetLeHttpChallengeTokens returns the LeHttpChallengeTokens field if non-nil, zero value otherwise.

### GetLeHttpChallengeTokensOk

`func (o *SiteUpdate) GetLeHttpChallengeTokensOk() (*[]string, bool)`

GetLeHttpChallengeTokensOk returns a tuple with the LeHttpChallengeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeHttpChallengeTokens

`func (o *SiteUpdate) SetLeHttpChallengeTokens(v []string)`

SetLeHttpChallengeTokens sets LeHttpChallengeTokens field to given value.

### HasLeHttpChallengeTokens

`func (o *SiteUpdate) HasLeHttpChallengeTokens() bool`

HasLeHttpChallengeTokens returns a boolean if a field has been set.

### GetDisabled

`func (o *SiteUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SiteUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SiteUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SiteUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPrimaryDomain

`func (o *SiteUpdate) GetPrimaryDomain() string`

GetPrimaryDomain returns the PrimaryDomain field if non-nil, zero value otherwise.

### GetPrimaryDomainOk

`func (o *SiteUpdate) GetPrimaryDomainOk() (*string, bool)`

GetPrimaryDomainOk returns a tuple with the PrimaryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDomain

`func (o *SiteUpdate) SetPrimaryDomain(v string)`

SetPrimaryDomain sets PrimaryDomain field to given value.

### HasPrimaryDomain

`func (o *SiteUpdate) HasPrimaryDomain() bool`

HasPrimaryDomain returns a boolean if a field has been set.

### SetPrimaryDomainNil

`func (o *SiteUpdate) SetPrimaryDomainNil(b bool)`

 SetPrimaryDomainNil sets the value for PrimaryDomain to be an explicit nil

### UnsetPrimaryDomain
`func (o *SiteUpdate) UnsetPrimaryDomain()`

UnsetPrimaryDomain ensures that no value is present for PrimaryDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


