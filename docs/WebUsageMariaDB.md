# WebUsageMariaDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Server** | [**WebUsageWebServer**](WebUsageWebServer.md) |  | 

## Methods

### NewWebUsageMariaDB

`func NewWebUsageMariaDB(id string, name string, server WebUsageWebServer, ) *WebUsageMariaDB`

NewWebUsageMariaDB instantiates a new WebUsageMariaDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebUsageMariaDBWithDefaults

`func NewWebUsageMariaDBWithDefaults() *WebUsageMariaDB`

NewWebUsageMariaDBWithDefaults instantiates a new WebUsageMariaDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebUsageMariaDB) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebUsageMariaDB) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebUsageMariaDB) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebUsageMariaDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebUsageMariaDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebUsageMariaDB) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *WebUsageMariaDB) GetServer() WebUsageWebServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *WebUsageMariaDB) GetServerOk() (*WebUsageWebServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *WebUsageMariaDB) SetServer(v WebUsageWebServer)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


