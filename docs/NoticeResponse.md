# NoticeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**NoticeTypeEnum**](NoticeTypeEnum.md) |  | 
**Content** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewNoticeResponse

`func NewNoticeResponse(id string, type_ NoticeTypeEnum, content string, createdAt time.Time, ) *NoticeResponse`

NewNoticeResponse instantiates a new NoticeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoticeResponseWithDefaults

`func NewNoticeResponseWithDefaults() *NoticeResponse`

NewNoticeResponseWithDefaults instantiates a new NoticeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoticeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoticeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoticeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NoticeResponse) GetType() NoticeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoticeResponse) GetTypeOk() (*NoticeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoticeResponse) SetType(v NoticeTypeEnum)`

SetType sets Type field to given value.


### GetContent

`func (o *NoticeResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoticeResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoticeResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *NoticeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NoticeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NoticeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


