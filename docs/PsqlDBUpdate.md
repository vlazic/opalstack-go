# PsqlDBUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DbusersReadwrite** | Pointer to **[]string** |  | [optional] 
**DbusersReadonly** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPsqlDBUpdate

`func NewPsqlDBUpdate(id string, ) *PsqlDBUpdate`

NewPsqlDBUpdate instantiates a new PsqlDBUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsqlDBUpdateWithDefaults

`func NewPsqlDBUpdateWithDefaults() *PsqlDBUpdate`

NewPsqlDBUpdateWithDefaults instantiates a new PsqlDBUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PsqlDBUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PsqlDBUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PsqlDBUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetDbusersReadwrite

`func (o *PsqlDBUpdate) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *PsqlDBUpdate) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *PsqlDBUpdate) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.

### HasDbusersReadwrite

`func (o *PsqlDBUpdate) HasDbusersReadwrite() bool`

HasDbusersReadwrite returns a boolean if a field has been set.

### GetDbusersReadonly

`func (o *PsqlDBUpdate) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *PsqlDBUpdate) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *PsqlDBUpdate) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.

### HasDbusersReadonly

`func (o *PsqlDBUpdate) HasDbusersReadonly() bool`

HasDbusersReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


