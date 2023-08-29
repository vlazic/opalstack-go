# SharedCertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Cert** | **string** |  | 
**Intermediates** | **string** |  | 
**ExpDate** | **time.Time** |  | 
**ListedDomains** | **[]string** |  | 
**IsOpalstackLetsencrypt** | Pointer to **bool** |  | [optional] [default to true]
**IsOpalstackSharedCert** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewSharedCertResponse

`func NewSharedCertResponse(id string, name string, cert string, intermediates string, expDate time.Time, listedDomains []string, ) *SharedCertResponse`

NewSharedCertResponse instantiates a new SharedCertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedCertResponseWithDefaults

`func NewSharedCertResponseWithDefaults() *SharedCertResponse`

NewSharedCertResponseWithDefaults instantiates a new SharedCertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedCertResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedCertResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedCertResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SharedCertResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedCertResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedCertResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCert

`func (o *SharedCertResponse) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *SharedCertResponse) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *SharedCertResponse) SetCert(v string)`

SetCert sets Cert field to given value.


### GetIntermediates

`func (o *SharedCertResponse) GetIntermediates() string`

GetIntermediates returns the Intermediates field if non-nil, zero value otherwise.

### GetIntermediatesOk

`func (o *SharedCertResponse) GetIntermediatesOk() (*string, bool)`

GetIntermediatesOk returns a tuple with the Intermediates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediates

`func (o *SharedCertResponse) SetIntermediates(v string)`

SetIntermediates sets Intermediates field to given value.


### GetExpDate

`func (o *SharedCertResponse) GetExpDate() time.Time`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *SharedCertResponse) GetExpDateOk() (*time.Time, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *SharedCertResponse) SetExpDate(v time.Time)`

SetExpDate sets ExpDate field to given value.


### GetListedDomains

`func (o *SharedCertResponse) GetListedDomains() []string`

GetListedDomains returns the ListedDomains field if non-nil, zero value otherwise.

### GetListedDomainsOk

`func (o *SharedCertResponse) GetListedDomainsOk() (*[]string, bool)`

GetListedDomainsOk returns a tuple with the ListedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedDomains

`func (o *SharedCertResponse) SetListedDomains(v []string)`

SetListedDomains sets ListedDomains field to given value.


### GetIsOpalstackLetsencrypt

`func (o *SharedCertResponse) GetIsOpalstackLetsencrypt() bool`

GetIsOpalstackLetsencrypt returns the IsOpalstackLetsencrypt field if non-nil, zero value otherwise.

### GetIsOpalstackLetsencryptOk

`func (o *SharedCertResponse) GetIsOpalstackLetsencryptOk() (*bool, bool)`

GetIsOpalstackLetsencryptOk returns a tuple with the IsOpalstackLetsencrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpalstackLetsencrypt

`func (o *SharedCertResponse) SetIsOpalstackLetsencrypt(v bool)`

SetIsOpalstackLetsencrypt sets IsOpalstackLetsencrypt field to given value.

### HasIsOpalstackLetsencrypt

`func (o *SharedCertResponse) HasIsOpalstackLetsencrypt() bool`

HasIsOpalstackLetsencrypt returns a boolean if a field has been set.

### GetIsOpalstackSharedCert

`func (o *SharedCertResponse) GetIsOpalstackSharedCert() bool`

GetIsOpalstackSharedCert returns the IsOpalstackSharedCert field if non-nil, zero value otherwise.

### GetIsOpalstackSharedCertOk

`func (o *SharedCertResponse) GetIsOpalstackSharedCertOk() (*bool, bool)`

GetIsOpalstackSharedCertOk returns a tuple with the IsOpalstackSharedCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpalstackSharedCert

`func (o *SharedCertResponse) SetIsOpalstackSharedCert(v bool)`

SetIsOpalstackSharedCert sets IsOpalstackSharedCert field to given value.

### HasIsOpalstackSharedCert

`func (o *SharedCertResponse) HasIsOpalstackSharedCert() bool`

HasIsOpalstackSharedCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


