# PsqlDBCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Server** | **string** |  | 
**Charset** | Pointer to [**PsqlCharset**](PsqlCharset.md) |  | [optional] 
**DbusersReadwrite** | Pointer to **[]string** |  | [optional] 
**DbusersReadonly** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPsqlDBCreate

`func NewPsqlDBCreate(name string, server string, ) *PsqlDBCreate`

NewPsqlDBCreate instantiates a new PsqlDBCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsqlDBCreateWithDefaults

`func NewPsqlDBCreateWithDefaults() *PsqlDBCreate`

NewPsqlDBCreateWithDefaults instantiates a new PsqlDBCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PsqlDBCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PsqlDBCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PsqlDBCreate) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *PsqlDBCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PsqlDBCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PsqlDBCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetCharset

`func (o *PsqlDBCreate) GetCharset() PsqlCharset`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *PsqlDBCreate) GetCharsetOk() (*PsqlCharset, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *PsqlDBCreate) SetCharset(v PsqlCharset)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *PsqlDBCreate) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetDbusersReadwrite

`func (o *PsqlDBCreate) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *PsqlDBCreate) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *PsqlDBCreate) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.

### HasDbusersReadwrite

`func (o *PsqlDBCreate) HasDbusersReadwrite() bool`

HasDbusersReadwrite returns a boolean if a field has been set.

### GetDbusersReadonly

`func (o *PsqlDBCreate) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *PsqlDBCreate) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *PsqlDBCreate) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.

### HasDbusersReadonly

`func (o *PsqlDBCreate) HasDbusersReadonly() bool`

HasDbusersReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


