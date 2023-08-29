# MariaDBUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DbusersReadwrite** | Pointer to **[]string** |  | [optional] 
**DbusersReadonly** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMariaDBUpdate

`func NewMariaDBUpdate(id string, ) *MariaDBUpdate`

NewMariaDBUpdate instantiates a new MariaDBUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaDBUpdateWithDefaults

`func NewMariaDBUpdateWithDefaults() *MariaDBUpdate`

NewMariaDBUpdateWithDefaults instantiates a new MariaDBUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariaDBUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariaDBUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariaDBUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetDbusersReadwrite

`func (o *MariaDBUpdate) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *MariaDBUpdate) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *MariaDBUpdate) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.

### HasDbusersReadwrite

`func (o *MariaDBUpdate) HasDbusersReadwrite() bool`

HasDbusersReadwrite returns a boolean if a field has been set.

### GetDbusersReadonly

`func (o *MariaDBUpdate) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *MariaDBUpdate) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *MariaDBUpdate) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.

### HasDbusersReadonly

`func (o *MariaDBUpdate) HasDbusersReadonly() bool`

HasDbusersReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


