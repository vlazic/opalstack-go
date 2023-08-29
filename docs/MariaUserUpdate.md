# MariaUserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewMariaUserUpdate

`func NewMariaUserUpdate(id string, ) *MariaUserUpdate`

NewMariaUserUpdate instantiates a new MariaUserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaUserUpdateWithDefaults

`func NewMariaUserUpdateWithDefaults() *MariaUserUpdate`

NewMariaUserUpdateWithDefaults instantiates a new MariaUserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariaUserUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariaUserUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariaUserUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *MariaUserUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MariaUserUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MariaUserUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MariaUserUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetExternal

`func (o *MariaUserUpdate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *MariaUserUpdate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *MariaUserUpdate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *MariaUserUpdate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


