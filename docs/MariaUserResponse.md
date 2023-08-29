# MariaUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Server** | **string** |  | 
**External** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMariaUserResponse

`func NewMariaUserResponse(id string, state StateEnum, ready bool, name string, server string, ) *MariaUserResponse`

NewMariaUserResponse instantiates a new MariaUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaUserResponseWithDefaults

`func NewMariaUserResponseWithDefaults() *MariaUserResponse`

NewMariaUserResponseWithDefaults instantiates a new MariaUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariaUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariaUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariaUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *MariaUserResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MariaUserResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MariaUserResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *MariaUserResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MariaUserResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MariaUserResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *MariaUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MariaUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MariaUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *MariaUserResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MariaUserResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MariaUserResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetExternal

`func (o *MariaUserResponse) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *MariaUserResponse) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *MariaUserResponse) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *MariaUserResponse) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


