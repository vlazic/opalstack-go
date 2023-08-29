# CertCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Cert** | **string** |  | 
**Intermediates** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** |  | 

## Methods

### NewCertCreate

`func NewCertCreate(cert string, key string, ) *CertCreate`

NewCertCreate instantiates a new CertCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertCreateWithDefaults

`func NewCertCreateWithDefaults() *CertCreate`

NewCertCreateWithDefaults instantiates a new CertCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCert

`func (o *CertCreate) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *CertCreate) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *CertCreate) SetCert(v string)`

SetCert sets Cert field to given value.


### GetIntermediates

`func (o *CertCreate) GetIntermediates() string`

GetIntermediates returns the Intermediates field if non-nil, zero value otherwise.

### GetIntermediatesOk

`func (o *CertCreate) GetIntermediatesOk() (*string, bool)`

GetIntermediatesOk returns a tuple with the Intermediates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediates

`func (o *CertCreate) SetIntermediates(v string)`

SetIntermediates sets Intermediates field to given value.

### HasIntermediates

`func (o *CertCreate) HasIntermediates() bool`

HasIntermediates returns a boolean if a field has been set.

### SetIntermediatesNil

`func (o *CertCreate) SetIntermediatesNil(b bool)`

 SetIntermediatesNil sets the value for Intermediates to be an explicit nil

### UnsetIntermediates
`func (o *CertCreate) UnsetIntermediates()`

UnsetIntermediates ensures that no value is present for Intermediates, not even an explicit nil
### GetKey

`func (o *CertCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CertCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CertCreate) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


