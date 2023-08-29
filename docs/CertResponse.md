# CertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Cert** | **string** |  | 
**Intermediates** | **string** |  | 
**Key** | **string** |  | 
**ExpDate** | **time.Time** |  | 
**ListedDomains** | **[]string** |  | 
**IsOpalstackLetsencrypt** | Pointer to **bool** |  | [optional] [default to true]
**IsOpalstackSharedCert** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCertResponse

`func NewCertResponse(id string, name string, cert string, intermediates string, key string, expDate time.Time, listedDomains []string, ) *CertResponse`

NewCertResponse instantiates a new CertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertResponseWithDefaults

`func NewCertResponseWithDefaults() *CertResponse`

NewCertResponseWithDefaults instantiates a new CertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CertResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCert

`func (o *CertResponse) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *CertResponse) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *CertResponse) SetCert(v string)`

SetCert sets Cert field to given value.


### GetIntermediates

`func (o *CertResponse) GetIntermediates() string`

GetIntermediates returns the Intermediates field if non-nil, zero value otherwise.

### GetIntermediatesOk

`func (o *CertResponse) GetIntermediatesOk() (*string, bool)`

GetIntermediatesOk returns a tuple with the Intermediates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediates

`func (o *CertResponse) SetIntermediates(v string)`

SetIntermediates sets Intermediates field to given value.


### GetKey

`func (o *CertResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CertResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CertResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetExpDate

`func (o *CertResponse) GetExpDate() time.Time`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *CertResponse) GetExpDateOk() (*time.Time, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *CertResponse) SetExpDate(v time.Time)`

SetExpDate sets ExpDate field to given value.


### GetListedDomains

`func (o *CertResponse) GetListedDomains() []string`

GetListedDomains returns the ListedDomains field if non-nil, zero value otherwise.

### GetListedDomainsOk

`func (o *CertResponse) GetListedDomainsOk() (*[]string, bool)`

GetListedDomainsOk returns a tuple with the ListedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedDomains

`func (o *CertResponse) SetListedDomains(v []string)`

SetListedDomains sets ListedDomains field to given value.


### GetIsOpalstackLetsencrypt

`func (o *CertResponse) GetIsOpalstackLetsencrypt() bool`

GetIsOpalstackLetsencrypt returns the IsOpalstackLetsencrypt field if non-nil, zero value otherwise.

### GetIsOpalstackLetsencryptOk

`func (o *CertResponse) GetIsOpalstackLetsencryptOk() (*bool, bool)`

GetIsOpalstackLetsencryptOk returns a tuple with the IsOpalstackLetsencrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpalstackLetsencrypt

`func (o *CertResponse) SetIsOpalstackLetsencrypt(v bool)`

SetIsOpalstackLetsencrypt sets IsOpalstackLetsencrypt field to given value.

### HasIsOpalstackLetsencrypt

`func (o *CertResponse) HasIsOpalstackLetsencrypt() bool`

HasIsOpalstackLetsencrypt returns a boolean if a field has been set.

### GetIsOpalstackSharedCert

`func (o *CertResponse) GetIsOpalstackSharedCert() bool`

GetIsOpalstackSharedCert returns the IsOpalstackSharedCert field if non-nil, zero value otherwise.

### GetIsOpalstackSharedCertOk

`func (o *CertResponse) GetIsOpalstackSharedCertOk() (*bool, bool)`

GetIsOpalstackSharedCertOk returns a tuple with the IsOpalstackSharedCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpalstackSharedCert

`func (o *CertResponse) SetIsOpalstackSharedCert(v bool)`

SetIsOpalstackSharedCert sets IsOpalstackSharedCert field to given value.

### HasIsOpalstackSharedCert

`func (o *CertResponse) HasIsOpalstackSharedCert() bool`

HasIsOpalstackSharedCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


