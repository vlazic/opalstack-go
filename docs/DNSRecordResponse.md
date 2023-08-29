# DNSRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Domain** | **string** |  | 
**Type** | [**DNSRecordTypeEnum**](DNSRecordTypeEnum.md) |  | 
**Content** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] [default to 10]
**Ttl** | Pointer to **int32** |  | [optional] [default to 3600]

## Methods

### NewDNSRecordResponse

`func NewDNSRecordResponse(id string, state StateEnum, ready bool, domain string, type_ DNSRecordTypeEnum, content string, ) *DNSRecordResponse`

NewDNSRecordResponse instantiates a new DNSRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordResponseWithDefaults

`func NewDNSRecordResponseWithDefaults() *DNSRecordResponse`

NewDNSRecordResponseWithDefaults instantiates a new DNSRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DNSRecordResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSRecordResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSRecordResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *DNSRecordResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DNSRecordResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DNSRecordResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *DNSRecordResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *DNSRecordResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *DNSRecordResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetDomain

`func (o *DNSRecordResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DNSRecordResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DNSRecordResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetType

`func (o *DNSRecordResponse) GetType() DNSRecordTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecordResponse) GetTypeOk() (*DNSRecordTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecordResponse) SetType(v DNSRecordTypeEnum)`

SetType sets Type field to given value.


### GetContent

`func (o *DNSRecordResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DNSRecordResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DNSRecordResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetPriority

`func (o *DNSRecordResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *DNSRecordResponse) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordResponse) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordResponse) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecordResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


