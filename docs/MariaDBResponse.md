# MariaDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Server** | **string** |  | 
**Charset** | [**MariaCharset**](MariaCharset.md) |  | 
**DbusersReadwrite** | **[]string** |  | 
**DbusersReadonly** | **[]string** |  | 

## Methods

### NewMariaDBResponse

`func NewMariaDBResponse(id string, state StateEnum, ready bool, name string, server string, charset MariaCharset, dbusersReadwrite []string, dbusersReadonly []string, ) *MariaDBResponse`

NewMariaDBResponse instantiates a new MariaDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariaDBResponseWithDefaults

`func NewMariaDBResponseWithDefaults() *MariaDBResponse`

NewMariaDBResponseWithDefaults instantiates a new MariaDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariaDBResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariaDBResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariaDBResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *MariaDBResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MariaDBResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MariaDBResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *MariaDBResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MariaDBResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MariaDBResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *MariaDBResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MariaDBResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MariaDBResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *MariaDBResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MariaDBResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MariaDBResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetCharset

`func (o *MariaDBResponse) GetCharset() MariaCharset`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *MariaDBResponse) GetCharsetOk() (*MariaCharset, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *MariaDBResponse) SetCharset(v MariaCharset)`

SetCharset sets Charset field to given value.


### GetDbusersReadwrite

`func (o *MariaDBResponse) GetDbusersReadwrite() []string`

GetDbusersReadwrite returns the DbusersReadwrite field if non-nil, zero value otherwise.

### GetDbusersReadwriteOk

`func (o *MariaDBResponse) GetDbusersReadwriteOk() (*[]string, bool)`

GetDbusersReadwriteOk returns a tuple with the DbusersReadwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadwrite

`func (o *MariaDBResponse) SetDbusersReadwrite(v []string)`

SetDbusersReadwrite sets DbusersReadwrite field to given value.


### GetDbusersReadonly

`func (o *MariaDBResponse) GetDbusersReadonly() []string`

GetDbusersReadonly returns the DbusersReadonly field if non-nil, zero value otherwise.

### GetDbusersReadonlyOk

`func (o *MariaDBResponse) GetDbusersReadonlyOk() (*[]string, bool)`

GetDbusersReadonlyOk returns a tuple with the DbusersReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbusersReadonly

`func (o *MariaDBResponse) SetDbusersReadonly(v []string)`

SetDbusersReadonly sets DbusersReadonly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


