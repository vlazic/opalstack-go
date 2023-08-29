# VirtualAliasCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Destinations** | Pointer to **[]string** |  | [optional] 
**Forwards** | **[]string** |  | 

## Methods

### NewVirtualAliasCreate

`func NewVirtualAliasCreate(source string, forwards []string, ) *VirtualAliasCreate`

NewVirtualAliasCreate instantiates a new VirtualAliasCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAliasCreateWithDefaults

`func NewVirtualAliasCreateWithDefaults() *VirtualAliasCreate`

NewVirtualAliasCreateWithDefaults instantiates a new VirtualAliasCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *VirtualAliasCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VirtualAliasCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VirtualAliasCreate) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestinations

`func (o *VirtualAliasCreate) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *VirtualAliasCreate) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *VirtualAliasCreate) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *VirtualAliasCreate) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetForwards

`func (o *VirtualAliasCreate) GetForwards() []string`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *VirtualAliasCreate) GetForwardsOk() (*[]string, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *VirtualAliasCreate) SetForwards(v []string)`

SetForwards sets Forwards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


