# OSVarCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Content** | Pointer to **string** |  | [optional] 
**Osusers** | Pointer to **[]string** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 

## Methods

### NewOSVarCreate

`func NewOSVarCreate(name string, ) *OSVarCreate`

NewOSVarCreate instantiates a new OSVarCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVarCreateWithDefaults

`func NewOSVarCreateWithDefaults() *OSVarCreate`

NewOSVarCreateWithDefaults instantiates a new OSVarCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OSVarCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSVarCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSVarCreate) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *OSVarCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OSVarCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OSVarCreate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OSVarCreate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetOsusers

`func (o *OSVarCreate) GetOsusers() []string`

GetOsusers returns the Osusers field if non-nil, zero value otherwise.

### GetOsusersOk

`func (o *OSVarCreate) GetOsusersOk() (*[]string, bool)`

GetOsusersOk returns a tuple with the Osusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsusers

`func (o *OSVarCreate) SetOsusers(v []string)`

SetOsusers sets Osusers field to given value.

### HasOsusers

`func (o *OSVarCreate) HasOsusers() bool`

HasOsusers returns a boolean if a field has been set.

### GetGlobal

`func (o *OSVarCreate) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *OSVarCreate) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *OSVarCreate) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *OSVarCreate) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


