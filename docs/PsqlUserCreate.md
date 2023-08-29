# PsqlUserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** |  | 
**Name** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewPsqlUserCreate

`func NewPsqlUserCreate(server string, name string, ) *PsqlUserCreate`

NewPsqlUserCreate instantiates a new PsqlUserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsqlUserCreateWithDefaults

`func NewPsqlUserCreateWithDefaults() *PsqlUserCreate`

NewPsqlUserCreateWithDefaults instantiates a new PsqlUserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *PsqlUserCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PsqlUserCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PsqlUserCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetName

`func (o *PsqlUserCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PsqlUserCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PsqlUserCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *PsqlUserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PsqlUserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PsqlUserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PsqlUserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetExternal

`func (o *PsqlUserCreate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *PsqlUserCreate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *PsqlUserCreate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *PsqlUserCreate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


