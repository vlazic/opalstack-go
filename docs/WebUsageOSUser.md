# WebUsageOSUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Server** | [**WebUsageWebServer**](WebUsageWebServer.md) |  | 

## Methods

### NewWebUsageOSUser

`func NewWebUsageOSUser(id string, name string, server WebUsageWebServer, ) *WebUsageOSUser`

NewWebUsageOSUser instantiates a new WebUsageOSUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebUsageOSUserWithDefaults

`func NewWebUsageOSUserWithDefaults() *WebUsageOSUser`

NewWebUsageOSUserWithDefaults instantiates a new WebUsageOSUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebUsageOSUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebUsageOSUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebUsageOSUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebUsageOSUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebUsageOSUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebUsageOSUser) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *WebUsageOSUser) GetServer() WebUsageWebServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *WebUsageOSUser) GetServerOk() (*WebUsageWebServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *WebUsageOSUser) SetServer(v WebUsageWebServer)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


