# OSVarUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Content** | Pointer to **string** |  | [optional] 
**Osusers** | Pointer to **[]string** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 

## Methods

### NewOSVarUpdate

`func NewOSVarUpdate(id string, ) *OSVarUpdate`

NewOSVarUpdate instantiates a new OSVarUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVarUpdateWithDefaults

`func NewOSVarUpdateWithDefaults() *OSVarUpdate`

NewOSVarUpdateWithDefaults instantiates a new OSVarUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OSVarUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSVarUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSVarUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *OSVarUpdate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OSVarUpdate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OSVarUpdate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OSVarUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetOsusers

`func (o *OSVarUpdate) GetOsusers() []string`

GetOsusers returns the Osusers field if non-nil, zero value otherwise.

### GetOsusersOk

`func (o *OSVarUpdate) GetOsusersOk() (*[]string, bool)`

GetOsusersOk returns a tuple with the Osusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsusers

`func (o *OSVarUpdate) SetOsusers(v []string)`

SetOsusers sets Osusers field to given value.

### HasOsusers

`func (o *OSVarUpdate) HasOsusers() bool`

HasOsusers returns a boolean if a field has been set.

### GetGlobal

`func (o *OSVarUpdate) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *OSVarUpdate) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *OSVarUpdate) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *OSVarUpdate) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


