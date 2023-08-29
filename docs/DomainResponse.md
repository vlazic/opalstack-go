# DomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] [default to "www.example.com"]
**DkimRecord** | **string** |  | 
**IsValidHostname** | **bool** |  | 

## Methods

### NewDomainResponse

`func NewDomainResponse(id string, state StateEnum, ready bool, dkimRecord string, isValidHostname bool, ) *DomainResponse`

NewDomainResponse instantiates a new DomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithDefaults

`func NewDomainResponseWithDefaults() *DomainResponse`

NewDomainResponseWithDefaults instantiates a new DomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *DomainResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *DomainResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *DomainResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *DomainResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *DomainResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDkimRecord

`func (o *DomainResponse) GetDkimRecord() string`

GetDkimRecord returns the DkimRecord field if non-nil, zero value otherwise.

### GetDkimRecordOk

`func (o *DomainResponse) GetDkimRecordOk() (*string, bool)`

GetDkimRecordOk returns a tuple with the DkimRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimRecord

`func (o *DomainResponse) SetDkimRecord(v string)`

SetDkimRecord sets DkimRecord field to given value.


### GetIsValidHostname

`func (o *DomainResponse) GetIsValidHostname() bool`

GetIsValidHostname returns the IsValidHostname field if non-nil, zero value otherwise.

### GetIsValidHostnameOk

`func (o *DomainResponse) GetIsValidHostnameOk() (*bool, bool)`

GetIsValidHostnameOk returns a tuple with the IsValidHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidHostname

`func (o *DomainResponse) SetIsValidHostname(v bool)`

SetIsValidHostname sets IsValidHostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


