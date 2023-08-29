# VirtualAliasUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Destinations** | Pointer to **[]string** |  | [optional] 
**Forwards** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualAliasUpdate

`func NewVirtualAliasUpdate(id string, ) *VirtualAliasUpdate`

NewVirtualAliasUpdate instantiates a new VirtualAliasUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAliasUpdateWithDefaults

`func NewVirtualAliasUpdateWithDefaults() *VirtualAliasUpdate`

NewVirtualAliasUpdateWithDefaults instantiates a new VirtualAliasUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualAliasUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualAliasUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualAliasUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *VirtualAliasUpdate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VirtualAliasUpdate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VirtualAliasUpdate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VirtualAliasUpdate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestinations

`func (o *VirtualAliasUpdate) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *VirtualAliasUpdate) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *VirtualAliasUpdate) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *VirtualAliasUpdate) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetForwards

`func (o *VirtualAliasUpdate) GetForwards() []string`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *VirtualAliasUpdate) GetForwardsOk() (*[]string, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *VirtualAliasUpdate) SetForwards(v []string)`

SetForwards sets Forwards field to given value.

### HasForwards

`func (o *VirtualAliasUpdate) HasForwards() bool`

HasForwards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


