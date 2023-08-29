# ApplicationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**AppTypeEnum**](AppTypeEnum.md) |  | [optional] 
**Json** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApplicationUpdate

`func NewApplicationUpdate(id string, ) *ApplicationUpdate`

NewApplicationUpdate instantiates a new ApplicationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateWithDefaults

`func NewApplicationUpdateWithDefaults() *ApplicationUpdate`

NewApplicationUpdateWithDefaults instantiates a new ApplicationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplicationUpdate) GetType() AppTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationUpdate) GetTypeOk() (*AppTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationUpdate) SetType(v AppTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJson

`func (o *ApplicationUpdate) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ApplicationUpdate) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ApplicationUpdate) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.

### HasJson

`func (o *ApplicationUpdate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### SetJsonNil

`func (o *ApplicationUpdate) SetJsonNil(b bool)`

 SetJsonNil sets the value for Json to be an explicit nil

### UnsetJson
`func (o *ApplicationUpdate) UnsetJson()`

UnsetJson ensures that no value is present for Json, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


