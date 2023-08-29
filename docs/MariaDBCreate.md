# MariaDBCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Server** | **string** |  | 
**Charset** | Pointer to [**MariaCharset**](MariaCharset.md) |  | [optional] 
**DbusersReadwrite** | Pointer to **[]string** |  | [optional] 
**DbusersReadonly** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMariaDBCreate

`func NewMariaDBCreate(name string, server string, ) *MariaDBCreate`

NewMariaDBCreate instantiates a new MariaDBCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaDBCreateWithDefaults

`func NewMariaDBCreateWithDefaults() *MariaDBCreate`

NewMariaDBCreateWithDefaults instantiates a new MariaDBCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MariaDBCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MariaDBCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MariaDBCreate) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *MariaDBCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MariaDBCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MariaDBCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetCharset

`func (o *MariaDBCreate) GetCharset() MariaCharset`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *MariaDBCreate) GetCharsetOk() (*MariaCharset, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *MariaDBCreate) SetCharset(v MariaCharset)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *MariaDBCreate) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetDbusersReadwrite

`func (o *MariaDBCreate) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *MariaDBCreate) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *MariaDBCreate) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.

### HasDbusersReadwrite

`func (o *MariaDBCreate) HasDbusersReadwrite() bool`

HasDbusersReadwrite returns a boolean if a field has been set.

### GetDbusersReadonly

`func (o *MariaDBCreate) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *MariaDBCreate) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *MariaDBCreate) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.

### HasDbusersReadonly

`func (o *MariaDBCreate) HasDbusersReadonly() bool`

HasDbusersReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


