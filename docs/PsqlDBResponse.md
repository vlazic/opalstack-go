# PsqlDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Server** | **string** |  | 
**Charset** | [**PsqlCharset**](PsqlCharset.md) |  | 
**DbusersReadwrite** | **[]string** |  | 
**DbusersReadonly** | **[]string** |  | 

## Methods

### NewPsqlDBResponse

`func NewPsqlDBResponse(id string, state StateEnum, ready bool, name string, server string, charset PsqlCharset, dbusersReadwrite []string, dbusersReadonly []string, ) *PsqlDBResponse`

NewPsqlDBResponse instantiates a new PsqlDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPsqlDBResponseWithDefaults

`func NewPsqlDBResponseWithDefaults() *PsqlDBResponse`

NewPsqlDBResponseWithDefaults instantiates a new PsqlDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PsqlDBResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PsqlDBResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PsqlDBResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *PsqlDBResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PsqlDBResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PsqlDBResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *PsqlDBResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *PsqlDBResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *PsqlDBResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *PsqlDBResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PsqlDBResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PsqlDBResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *PsqlDBResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PsqlDBResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PsqlDBResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetCharset

`func (o *PsqlDBResponse) GetCharset() PsqlCharset`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *PsqlDBResponse) GetCharsetOk() (*PsqlCharset, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *PsqlDBResponse) SetCharset(v PsqlCharset)`

SetCharset sets Charset field to given value.


### GetDbusersReadwrite

`func (o *PsqlDBResponse) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *PsqlDBResponse) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *PsqlDBResponse) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.


### GetDbusersReadonly

`func (o *PsqlDBResponse) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *PsqlDBResponse) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *PsqlDBResponse) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


