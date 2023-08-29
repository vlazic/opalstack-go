# DNSRecordCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Type** | [**DNSRecordTypeEnum**](DNSRecordTypeEnum.md) |  | 
**Content** | **string** |  | 
**Priority** | Pointer to **int32** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewDNSRecordCreate

`func NewDNSRecordCreate(domain string, type_ DNSRecordTypeEnum, content string, ) *DNSRecordCreate`

NewDNSRecordCreate instantiates a new DNSRecordCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordCreateWithDefaults

`func NewDNSRecordCreateWithDefaults() *DNSRecordCreate`

NewDNSRecordCreateWithDefaults instantiates a new DNSRecordCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DNSRecordCreate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DNSRecordCreate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DNSRecordCreate) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetType

`func (o *DNSRecordCreate) GetType() DNSRecordTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecordCreate) GetTypeOk() (*DNSRecordTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecordCreate) SetType(v DNSRecordTypeEnum)`

SetType sets Type field to given value.


### GetContent

`func (o *DNSRecordCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DNSRecordCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DNSRecordCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetPriority

`func (o *DNSRecordCreate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordCreate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordCreate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *DNSRecordCreate) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordCreate) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordCreate) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecordCreate) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


