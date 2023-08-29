# MariaUserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** |  | 
**Name** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewMariaUserCreate

`func NewMariaUserCreate(server string, name string, ) *MariaUserCreate`

NewMariaUserCreate instantiates a new MariaUserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaUserCreateWithDefaults

`func NewMariaUserCreateWithDefaults() *MariaUserCreate`

NewMariaUserCreateWithDefaults instantiates a new MariaUserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *MariaUserCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MariaUserCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MariaUserCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetName

`func (o *MariaUserCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MariaUserCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MariaUserCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *MariaUserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MariaUserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MariaUserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MariaUserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetExternal

`func (o *MariaUserCreate) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *MariaUserCreate) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *MariaUserCreate) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *MariaUserCreate) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


