# IPAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Ip** | **string** |  | 
**Server** | **string** |  | 
**Type** | [**IPAddressTypeEnum**](IPAddressTypeEnum.md) |  | 
**Primary** | **bool** |  | 

## Methods

### NewIPAddressResponse

`func NewIPAddressResponse(id string, ip string, server string, type_ IPAddressTypeEnum, primary bool, ) *IPAddressResponse`

NewIPAddressResponse instantiates a new IPAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressResponseWithDefaults

`func NewIPAddressResponseWithDefaults() *IPAddressResponse`

NewIPAddressResponseWithDefaults instantiates a new IPAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddressResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddressResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddressResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *IPAddressResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPAddressResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPAddressResponse) SetIp(v string)`

SetIp sets Ip field to given value.


### GetServer

`func (o *IPAddressResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IPAddressResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IPAddressResponse) SetServer(v string)`

SetServer sets Server field to given value.


### GetType

`func (o *IPAddressResponse) GetType() IPAddressTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAddressResponse) GetTypeOk() (*IPAddressTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAddressResponse) SetType(v IPAddressTypeEnum)`

SetType sets Type field to given value.


### GetPrimary

`func (o *IPAddressResponse) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *IPAddressResponse) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *IPAddressResponse) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


