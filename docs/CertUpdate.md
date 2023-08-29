# CertUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Intermediates** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewCertUpdate

`func NewCertUpdate(id string, ) *CertUpdate`

NewCertUpdate instantiates a new CertUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertUpdateWithDefaults

`func NewCertUpdateWithDefaults() *CertUpdate`

NewCertUpdateWithDefaults instantiates a new CertUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CertUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCert

`func (o *CertUpdate) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *CertUpdate) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *CertUpdate) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *CertUpdate) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetIntermediates

`func (o *CertUpdate) GetIntermediates() string`

GetIntermediates returns the Intermediates field if non-nil, zero value otherwise.

### GetIntermediatesOk

`func (o *CertUpdate) GetIntermediatesOk() (*string, bool)`

GetIntermediatesOk returns a tuple with the Intermediates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediates

`func (o *CertUpdate) SetIntermediates(v string)`

SetIntermediates sets Intermediates field to given value.

### HasIntermediates

`func (o *CertUpdate) HasIntermediates() bool`

HasIntermediates returns a boolean if a field has been set.

### SetIntermediatesNil

`func (o *CertUpdate) SetIntermediatesNil(b bool)`

 SetIntermediatesNil sets the value for Intermediates to be an explicit nil

### UnsetIntermediates
`func (o *CertUpdate) UnsetIntermediates()`

UnsetIntermediates ensures that no value is present for Intermediates, not even an explicit nil
### GetKey

`func (o *CertUpdate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CertUpdate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CertUpdate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CertUpdate) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


