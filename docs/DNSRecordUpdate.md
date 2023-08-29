# DNSRecordUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DNSRecordTypeEnum**](DNSRecordTypeEnum.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewDNSRecordUpdate

`func NewDNSRecordUpdate(id string, ) *DNSRecordUpdate`

NewDNSRecordUpdate instantiates a new DNSRecordUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordUpdateWithDefaults

`func NewDNSRecordUpdateWithDefaults() *DNSRecordUpdate`

NewDNSRecordUpdateWithDefaults instantiates a new DNSRecordUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DNSRecordUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSRecordUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSRecordUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *DNSRecordUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DNSRecordUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DNSRecordUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DNSRecordUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetType

`func (o *DNSRecordUpdate) GetType() DNSRecordTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecordUpdate) GetTypeOk() (*DNSRecordTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecordUpdate) SetType(v DNSRecordTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *DNSRecordUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *DNSRecordUpdate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DNSRecordUpdate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DNSRecordUpdate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DNSRecordUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPriority

`func (o *DNSRecordUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *DNSRecordUpdate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordUpdate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordUpdate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecordUpdate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


