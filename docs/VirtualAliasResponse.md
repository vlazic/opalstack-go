# VirtualAliasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Source** | **string** |  | 
**Destinations** | **[]string** |  | 
**Forwards** | **[]string** |  | 

## Methods

### NewVirtualAliasResponse

`func NewVirtualAliasResponse(id string, state StateEnum, ready bool, source string, destinations []string, forwards []string, ) *VirtualAliasResponse`

NewVirtualAliasResponse instantiates a new VirtualAliasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAliasResponseWithDefaults

`func NewVirtualAliasResponseWithDefaults() *VirtualAliasResponse`

NewVirtualAliasResponseWithDefaults instantiates a new VirtualAliasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualAliasResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualAliasResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualAliasResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *VirtualAliasResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualAliasResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VirtualAliasResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *VirtualAliasResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *VirtualAliasResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *VirtualAliasResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetSource

`func (o *VirtualAliasResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VirtualAliasResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VirtualAliasResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestinations

`func (o *VirtualAliasResponse) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *VirtualAliasResponse) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *VirtualAliasResponse) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetForwards

`func (o *VirtualAliasResponse) GetForwards() []string`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *VirtualAliasResponse) GetForwardsOk() (*[]string, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *VirtualAliasResponse) SetForwards(v []string)`

SetForwards sets Forwards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


