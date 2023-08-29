# PsqlUserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewPsqlUserUpdate

`func NewPsqlUserUpdate(id string, ) *PsqlUserUpdate`

NewPsqlUserUpdate instantiates a new PsqlUserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsqlUserUpdateWithDefaults

`func NewPsqlUserUpdateWithDefaults() *PsqlUserUpdate`

NewPsqlUserUpdateWithDefaults instantiates a new PsqlUserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PsqlUserUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PsqlUserUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PsqlUserUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *PsqlUserUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PsqlUserUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PsqlUserUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PsqlUserUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetExternal

`func (o *PsqlUserUpdate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *PsqlUserUpdate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *PsqlUserUpdate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *PsqlUserUpdate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


