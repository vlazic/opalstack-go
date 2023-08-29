# SiteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ip4** | Pointer to **NullableString** |  | [optional] 
**Ip6** | Pointer to **NullableString** |  | [optional] 
**Domains** | **[]string** |  | 
**Routes** | [**[]Route**](Route.md) |  | 
**Cert** | Pointer to **NullableString** |  | [optional] 
**Redirect** | Pointer to **bool** | Automatically redirect to https:// | [optional] 
**GenerateLe** | Pointer to **bool** | Automatically issue Lets Encrypt certificate for the domains on this site? | [optional] 
**LeHttpChallengeTokens** | Pointer to **[]string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**PrimaryDomain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSiteCreate

`func NewSiteCreate(domains []string, routes []Route, ) *SiteCreate`

NewSiteCreate instantiates a new SiteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteCreateWithDefaults

`func NewSiteCreateWithDefaults() *SiteCreate`

NewSiteCreateWithDefaults instantiates a new SiteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SiteCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp4

`func (o *SiteCreate) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *SiteCreate) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *SiteCreate) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.

### HasIp4

`func (o *SiteCreate) HasIp4() bool`

HasIp4 returns a boolean if a field has been set.

### SetIp4Nil

`func (o *SiteCreate) SetIp4Nil(b bool)`

 SetIp4Nil sets the value for Ip4 to be an explicit nil

### UnsetIp4
`func (o *SiteCreate) UnsetIp4()`

UnsetIp4 ensures that no value is present for Ip4, not even an explicit nil
### GetIp6

`func (o *SiteCreate) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *SiteCreate) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *SiteCreate) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *SiteCreate) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### SetIp6Nil

`func (o *SiteCreate) SetIp6Nil(b bool)`

 SetIp6Nil sets the value for Ip6 to be an explicit nil

### UnsetIp6
`func (o *SiteCreate) UnsetIp6()`

UnsetIp6 ensures that no value is present for Ip6, not even an explicit nil
### GetDomains

`func (o *SiteCreate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SiteCreate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SiteCreate) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetRoutes

`func (o *SiteCreate) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SiteCreate) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SiteCreate) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.


### GetCert

`func (o *SiteCreate) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SiteCreate) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *SiteCreate) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *SiteCreate) HasCert() bool`

HasCert returns a boolean if a field has been set.

### SetCertNil

`func (o *SiteCreate) SetCertNil(b bool)`

 SetCertNil sets the value for Cert to be an explicit nil

### UnsetCert
`func (o *SiteCreate) UnsetCert()`

UnsetCert ensures that no value is present for Cert, not even an explicit nil
### GetRedirect

`func (o *SiteCreate) GetRedirect() bool`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SiteCreate) GetRedirectOk() (*bool, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SiteCreate) SetRedirect(v bool)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *SiteCreate) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetGenerateLe

`func (o *SiteCreate) GetGenerateLe() bool`

GetGenerateLe returns the GenerateLe field if non-nil, zero value otherwise.

### GetGenerateLeOk

`func (o *SiteCreate) GetGenerateLeOk() (*bool, bool)`

GetGenerateLeOk returns a tuple with the GenerateLe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLe

`func (o *SiteCreate) SetGenerateLe(v bool)`

SetGenerateLe sets GenerateLe field to given value.

### HasGenerateLe

`func (o *SiteCreate) HasGenerateLe() bool`

HasGenerateLe returns a boolean if a field has been set.

### GetLeHttpChallengeTokens

`func (o *SiteCreate) GetLeHttpChallengeTokens() []string`

GetLeHttpChallengeTokens returns the LeHttpChallengeTokens field if non-nil, zero value otherwise.

### GetLeHttpChallengeTokensOk

`func (o *SiteCreate) GetLeHttpChallengeTokensOk() (*[]string, bool)`

GetLeHttpChallengeTokensOk returns a tuple with the LeHttpChallengeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeHttpChallengeTokens

`func (o *SiteCreate) SetLeHttpChallengeTokens(v []string)`

SetLeHttpChallengeTokens sets LeHttpChallengeTokens field to given value.

### HasLeHttpChallengeTokens

`func (o *SiteCreate) HasLeHttpChallengeTokens() bool`

HasLeHttpChallengeTokens returns a boolean if a field has been set.

### GetDisabled

`func (o *SiteCreate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SiteCreate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SiteCreate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SiteCreate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetPrimaryDomain

`func (o *SiteCreate) GetPrimaryDomain() string`

GetPrimaryDomain returns the PrimaryDomain field if non-nil, zero value otherwise.

### GetPrimaryDomainOk

`func (o *SiteCreate) GetPrimaryDomainOk() (*string, bool)`

GetPrimaryDomainOk returns a tuple with the PrimaryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDomain

`func (o *SiteCreate) SetPrimaryDomain(v string)`

SetPrimaryDomain sets PrimaryDomain field to given value.

### HasPrimaryDomain

`func (o *SiteCreate) HasPrimaryDomain() bool`

HasPrimaryDomain returns a boolean if a field has been set.

### SetPrimaryDomainNil

`func (o *SiteCreate) SetPrimaryDomainNil(b bool)`

 SetPrimaryDomainNil sets the value for PrimaryDomain to be an explicit nil

### UnsetPrimaryDomain
`func (o *SiteCreate) UnsetPrimaryDomain()`

UnsetPrimaryDomain ensures that no value is present for PrimaryDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


