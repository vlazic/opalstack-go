# MailUserUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Procmailrc** | Pointer to **string** |  | [optional] 
**AutoresponderEnable** | Pointer to **bool** |  | [optional] 
**AutoresponderSubject** | Pointer to **string** |  | [optional] 
**AutoresponderMessage** | Pointer to **string** |  | [optional] 
**AutoresponderNoreply** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMailUserUpdate

`func NewMailUserUpdate(id string, ) *MailUserUpdate`

NewMailUserUpdate instantiates a new MailUserUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUserUpdateWithDefaults

`func NewMailUserUpdateWithDefaults() *MailUserUpdate`

NewMailUserUpdateWithDefaults instantiates a new MailUserUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MailUserUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailUserUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailUserUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *MailUserUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailUserUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailUserUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MailUserUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProcmailrc

`func (o *MailUserUpdate) GetProcmailrc() string`

GetProcmailrc returns the Procmailrc field if non-nil, zero value otherwise.

### GetProcmailrcOk

`func (o *MailUserUpdate) GetProcmailrcOk() (*string, bool)`

GetProcmailrcOk returns a tuple with the Procmailrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcmailrc

`func (o *MailUserUpdate) SetProcmailrc(v string)`

SetProcmailrc sets Procmailrc field to given value.

### HasProcmailrc

`func (o *MailUserUpdate) HasProcmailrc() bool`

HasProcmailrc returns a boolean if a field has been set.

### GetAutoresponderEnable

`func (o *MailUserUpdate) GetAutoresponderEnable() bool`

GetAutoresponderEnable returns the AutoresponderEnable field if non-nil, zero value otherwise.

### GetAutoresponderEnableOk

`func (o *MailUserUpdate) GetAutoresponderEnableOk() (*bool, bool)`

GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderEnable

`func (o *MailUserUpdate) SetAutoresponderEnable(v bool)`

SetAutoresponderEnable sets AutoresponderEnable field to given value.

### HasAutoresponderEnable

`func (o *MailUserUpdate) HasAutoresponderEnable() bool`

HasAutoresponderEnable returns a boolean if a field has been set.

### GetAutoresponderSubject

`func (o *MailUserUpdate) GetAutoresponderSubject() string`

GetAutoresponderSubject returns the AutoresponderSubject field if non-nil, zero value otherwise.

### GetAutoresponderSubjectOk

`func (o *MailUserUpdate) GetAutoresponderSubjectOk() (*string, bool)`

GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderSubject

`func (o *MailUserUpdate) SetAutoresponderSubject(v string)`

SetAutoresponderSubject sets AutoresponderSubject field to given value.

### HasAutoresponderSubject

`func (o *MailUserUpdate) HasAutoresponderSubject() bool`

HasAutoresponderSubject returns a boolean if a field has been set.

### GetAutoresponderMessage

`func (o *MailUserUpdate) GetAutoresponderMessage() string`

GetAutoresponderMessage returns the AutoresponderMessage field if non-nil, zero value otherwise.

### GetAutoresponderMessageOk

`func (o *MailUserUpdate) GetAutoresponderMessageOk() (*string, bool)`

GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderMessage

`func (o *MailUserUpdate) SetAutoresponderMessage(v string)`

SetAutoresponderMessage sets AutoresponderMessage field to given value.

### HasAutoresponderMessage

`func (o *MailUserUpdate) HasAutoresponderMessage() bool`

HasAutoresponderMessage returns a boolean if a field has been set.

### GetAutoresponderNoreply

`func (o *MailUserUpdate) GetAutoresponderNoreply() []string`

GetAutoresponderNoreply returns the AutoresponderNoreply field if non-nil, zero value otherwise.

### GetAutoresponderNoreplyOk

`func (o *MailUserUpdate) GetAutoresponderNoreplyOk() (*[]string, bool)`

GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderNoreply

`func (o *MailUserUpdate) SetAutoresponderNoreply(v []string)`

SetAutoresponderNoreply sets AutoresponderNoreply field to given value.

### HasAutoresponderNoreply

`func (o *MailUserUpdate) HasAutoresponderNoreply() bool`

HasAutoresponderNoreply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


