# OSVarResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Content** | **string** |  | 
**Osusers** | **[]string** |  | 
**Global** | **bool** |  | 

## Methods

### NewOSVarResponse

`func NewOSVarResponse(id string, state StateEnum, ready bool, name string, content string, osusers []string, global bool, ) *OSVarResponse`

NewOSVarResponse instantiates a new OSVarResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVarResponseWithDefaults

`func NewOSVarResponseWithDefaults() *OSVarResponse`

NewOSVarResponseWithDefaults instantiates a new OSVarResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OSVarResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSVarResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSVarResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *OSVarResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OSVarResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OSVarResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *OSVarResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *OSVarResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *OSVarResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *OSVarResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSVarResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSVarResponse) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *OSVarResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OSVarResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OSVarResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetOsusers

`func (o *OSVarResponse) GetOsusers() []string`

GetOsusers returns the Osusers field if non-nil, zero value otherwise.

### GetOsusersOk

`func (o *OSVarResponse) GetOsusersOk() (*[]string, bool)`

GetOsusersOk returns a tuple with the Osusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsusers

`func (o *OSVarResponse) SetOsusers(v []string)`

SetOsusers sets Osusers field to given value.


### GetGlobal

`func (o *OSVarResponse) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *OSVarResponse) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *OSVarResponse) SetGlobal(v bool)`

SetGlobal sets Global field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


