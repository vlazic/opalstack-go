# SiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Server** | **string** |  | 
**Ip4** | **string** |  | 
**Ip6** | **string** |  | 
**Disabled** | **bool** |  | 
**Domains** | **[]string** |  | 
**Routes** | [**[]RouteResponse**](RouteResponse.md) |  | 
**GenerateLe** | **bool** |  | 
**Cert** | **string** |  | 
**Redirect** | **bool** |  | 
**LeHttpChallengeTokens** | **[]string** |  | 

## Methods

### NewSiteResponse

`func NewSiteResponse(id string, state StateEnum, ready bool, name string, server string, ip4 string, ip6 string, disabled bool, domains []string, routes []RouteResponse, generateLe bool, cert string, redirect bool, leHttpChallengeTokens []string, ) *SiteResponse`

NewSiteResponse instantiates a new SiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteResponseWithDefaults

`func NewSiteResponseWithDefaults() *SiteResponse`

NewSiteResponseWithDefaults instantiates a new SiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *SiteResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SiteResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SiteResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *SiteResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *SiteResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *SiteResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *SiteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *SiteResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *SiteResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *SiteResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetIp4

`func (o *SiteResponse) GetIp4() string`

GetIp4 returns the Ip4 field if non-nil, zero value otherwise.

### GetIp4Ok

`func (o *SiteResponse) GetIp4Ok() (*string, bool)`

GetIp4Ok returns a tuple with the Ip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp4

`func (o *SiteResponse) SetIp4(v string)`

SetIp4 sets Ip4 field to given value.


### GetIp6

`func (o *SiteResponse) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *SiteResponse) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *SiteResponse) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.


### GetDisabled

`func (o *SiteResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SiteResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SiteResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetDomains

`func (o *SiteResponse) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SiteResponse) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SiteResponse) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetRoutes

`func (o *SiteResponse) GetRoutes() []RouteResponse`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SiteResponse) GetRoutesOk() (*[]RouteResponse, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SiteResponse) SetRoutes(v []RouteResponse)`

SetRoutes sets Routes field to given value.


### GetGenerateLe

`func (o *SiteResponse) GetGenerateLe() bool`

GetGenerateLe returns the GenerateLe field if non-nil, zero value otherwise.

### GetGenerateLeOk

`func (o *SiteResponse) GetGenerateLeOk() (*bool, bool)`

GetGenerateLeOk returns a tuple with the GenerateLe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLe

`func (o *SiteResponse) SetGenerateLe(v bool)`

SetGenerateLe sets GenerateLe field to given value.


### GetCert

`func (o *SiteResponse) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SiteResponse) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *SiteResponse) SetCert(v string)`

SetCert sets Cert field to given value.


### GetRedirect

`func (o *SiteResponse) GetRedirect() bool`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SiteResponse) GetRedirectOk() (*bool, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SiteResponse) SetRedirect(v bool)`

SetRedirect sets Redirect field to given value.


### GetLeHttpChallengeTokens

`func (o *SiteResponse) GetLeHttpChallengeTokens() []string`

GetLeHttpChallengeTokens returns the LeHttpChallengeTokens field if non-nil, zero value otherwise.

### GetLeHttpChallengeTokensOk

`func (o *SiteResponse) GetLeHttpChallengeTokensOk() (*[]string, bool)`

GetLeHttpChallengeTokensOk returns a tuple with the LeHttpChallengeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeHttpChallengeTokens

`func (o *SiteResponse) SetLeHttpChallengeTokens(v []string)`

SetLeHttpChallengeTokens sets LeHttpChallengeTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


