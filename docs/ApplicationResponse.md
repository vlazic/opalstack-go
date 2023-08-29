# ApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**Server** | **string** |  | 
**Osuser** | **string** |  | 
**Type** | [**AppTypeEnum**](AppTypeEnum.md) |  | 
**Port** | Pointer to **int32** |  | [optional] [default to 8000]
**InstallerUrl** | Pointer to **string** |  | [optional] [default to "https://raw.githubusercontent.com/opalstack/installers/master/core/django/install.py"]
**Json** | [**ApplicationResponseJson**](ApplicationResponseJson.md) |  | 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(id string, state StateEnum, ready bool, name string, server string, osuser string, type_ AppTypeEnum, json ApplicationResponseJson, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ApplicationResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *ApplicationResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ApplicationResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ApplicationResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *ApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *ApplicationResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ApplicationResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ApplicationResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetOsuser

`func (o *ApplicationResponse) GetOsuser() string`

GetOsuser returns the Osuser field if non-nil, zero value otherwise.

### GetOsuserOk

`func (o *ApplicationResponse) GetOsuserOk() (*string, bool)`

GetOsuserOk returns a tuple with the Osuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuser

`func (o *ApplicationResponse) SetOsuser(v string)`

SetOsuser sets Osuser field to given value.


### GetType

`func (o *ApplicationResponse) GetType() AppTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponse) GetTypeOk() (*AppTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponse) SetType(v AppTypeEnum)`

SetType sets Type field to given value.


### GetPort

`func (o *ApplicationResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApplicationResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApplicationResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApplicationResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetInstallerUrl

`func (o *ApplicationResponse) GetInstallerUrl() string`

GetInstallerUrl returns the InstallerUrl field if non-nil, zero value otherwise.

### GetInstallerUrlOk

`func (o *ApplicationResponse) GetInstallerUrlOk() (*string, bool)`

GetInstallerUrlOk returns a tuple with the InstallerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerUrl

`func (o *ApplicationResponse) SetInstallerUrl(v string)`

SetInstallerUrl sets InstallerUrl field to given value.

### HasInstallerUrl

`func (o *ApplicationResponse) HasInstallerUrl() bool`

HasInstallerUrl returns a boolean if a field has been set.

### GetJson

`func (o *ApplicationResponse) GetJson() ApplicationResponseJson`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ApplicationResponse) GetJsonOk() (*ApplicationResponseJson, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ApplicationResponse) SetJson(v ApplicationResponseJson)`

SetJson sets Json field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


