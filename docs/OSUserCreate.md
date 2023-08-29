# OSUserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewOSUserCreate

`func NewOSUserCreate(server string, ) *OSUserCreate`

NewOSUserCreate instantiates a new OSUserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSUserCreateWithDefaults

`func NewOSUserCreateWithDefaults() *OSUserCreate`

NewOSUserCreateWithDefaults instantiates a new OSUserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *OSUserCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OSUserCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OSUserCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetName

`func (o *OSUserCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSUserCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSUserCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSUserCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *OSUserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OSUserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OSUserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OSUserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


