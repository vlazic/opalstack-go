# ApplicationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Osuser** | **string** |  | 
**Type** | [**AppTypeEnum**](AppTypeEnum.md) |  | 
**InstallerUrl** | Pointer to **NullableString** |  | [optional] 
**Json** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApplicationCreate

`func NewApplicationCreate(name string, osuser string, type_ AppTypeEnum, ) *ApplicationCreate`

NewApplicationCreate instantiates a new ApplicationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCreateWithDefaults

`func NewApplicationCreateWithDefaults() *ApplicationCreate`

NewApplicationCreateWithDefaults instantiates a new ApplicationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOsuser

`func (o *ApplicationCreate) GetOsuser() string`

GetOsuser returns the Osuser field if non-nil, zero value otherwise.

### GetOsuserOk

`func (o *ApplicationCreate) GetOsuserOk() (*string, bool)`

GetOsuserOk returns a tuple with the Osuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuser

`func (o *ApplicationCreate) SetOsuser(v string)`

SetOsuser sets Osuser field to given value.


### GetType

`func (o *ApplicationCreate) GetType() AppTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCreate) GetTypeOk() (*AppTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCreate) SetType(v AppTypeEnum)`

SetType sets Type field to given value.


### GetInstallerUrl

`func (o *ApplicationCreate) GetInstallerUrl() string`

GetInstallerUrl returns the InstallerUrl field if non-nil, zero value otherwise.

### GetInstallerUrlOk

`func (o *ApplicationCreate) GetInstallerUrlOk() (*string, bool)`

GetInstallerUrlOk returns a tuple with the InstallerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerUrl

`func (o *ApplicationCreate) SetInstallerUrl(v string)`

SetInstallerUrl sets InstallerUrl field to given value.

### HasInstallerUrl

`func (o *ApplicationCreate) HasInstallerUrl() bool`

HasInstallerUrl returns a boolean if a field has been set.

### SetInstallerUrlNil

`func (o *ApplicationCreate) SetInstallerUrlNil(b bool)`

 SetInstallerUrlNil sets the value for InstallerUrl to be an explicit nil

### UnsetInstallerUrl
`func (o *ApplicationCreate) UnsetInstallerUrl()`

UnsetInstallerUrl ensures that no value is present for InstallerUrl, not even an explicit nil
### GetJson

`func (o *ApplicationCreate) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ApplicationCreate) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ApplicationCreate) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.

### HasJson

`func (o *ApplicationCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### SetJsonNil

`func (o *ApplicationCreate) SetJsonNil(b bool)`

 SetJsonNil sets the value for Json to be an explicit nil

### UnsetJson
`func (o *ApplicationCreate) UnsetJson()`

UnsetJson ensures that no value is present for Json, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


