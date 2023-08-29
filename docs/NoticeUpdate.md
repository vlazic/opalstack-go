# NoticeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**NoticeTypeEnum**](NoticeTypeEnum.md) |  | [optional] 
**Content** | **string** |  | 

## Methods

### NewNoticeUpdate

`func NewNoticeUpdate(id string, content string, ) *NoticeUpdate`

NewNoticeUpdate instantiates a new NoticeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoticeUpdateWithDefaults

`func NewNoticeUpdateWithDefaults() *NoticeUpdate`

NewNoticeUpdateWithDefaults instantiates a new NoticeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoticeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoticeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoticeUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NoticeUpdate) GetType() NoticeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoticeUpdate) GetTypeOk() (*NoticeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoticeUpdate) SetType(v NoticeTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *NoticeUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *NoticeUpdate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoticeUpdate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoticeUpdate) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


