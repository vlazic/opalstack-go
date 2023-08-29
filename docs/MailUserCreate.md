# MailUserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImapServer** | **string** |  | 
**Name** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Procmailrc** | Pointer to **string** |  | [optional] 
**AutoresponderEnable** | Pointer to **bool** |  | [optional] 
**AutoresponderSubject** | Pointer to **string** |  | [optional] 
**AutoresponderMessage** | Pointer to **string** |  | [optional] 
**AutoresponderNoreply** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMailUserCreate

`func NewMailUserCreate(imapServer string, name string, ) *MailUserCreate`

NewMailUserCreate instantiates a new MailUserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUserCreateWithDefaults

`func NewMailUserCreateWithDefaults() *MailUserCreate`

NewMailUserCreateWithDefaults instantiates a new MailUserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImapServer

`func (o *MailUserCreate) GetImapServer() string`

GetImapServer returns the ImapServer field if non-nil, zero value otherwise.

### GetImapServerOk

`func (o *MailUserCreate) GetImapServerOk() (*string, bool)`

GetImapServerOk returns a tuple with the ImapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapServer

`func (o *MailUserCreate) SetImapServer(v string)`

SetImapServer sets ImapServer field to given value.


### GetName

`func (o *MailUserCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailUserCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailUserCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *MailUserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailUserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailUserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MailUserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProcmailrc

`func (o *MailUserCreate) GetProcmailrc() string`

GetProcmailrc returns the Procmailrc field if non-nil, zero value otherwise.

### GetProcmailrcOk

`func (o *MailUserCreate) GetProcmailrcOk() (*string, bool)`

GetProcmailrcOk returns a tuple with the Procmailrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcmailrc

`func (o *MailUserCreate) SetProcmailrc(v string)`

SetProcmailrc sets Procmailrc field to given value.

### HasProcmailrc

`func (o *MailUserCreate) HasProcmailrc() bool`

HasProcmailrc returns a boolean if a field has been set.

### GetAutoresponderEnable

`func (o *MailUserCreate) GetAutoresponderEnable() bool`

GetAutoresponderEnable returns the AutoresponderEnable field if non-nil, zero value otherwise.

### GetAutoresponderEnableOk

`func (o *MailUserCreate) GetAutoresponderEnableOk() (*bool, bool)`

GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderEnable

`func (o *MailUserCreate) SetAutoresponderEnable(v bool)`

SetAutoresponderEnable sets AutoresponderEnable field to given value.

### HasAutoresponderEnable

`func (o *MailUserCreate) HasAutoresponderEnable() bool`

HasAutoresponderEnable returns a boolean if a field has been set.

### GetAutoresponderSubject

`func (o *MailUserCreate) GetAutoresponderSubject() string`

GetAutoresponderSubject returns the AutoresponderSubject field if non-nil, zero value otherwise.

### GetAutoresponderSubjectOk

`func (o *MailUserCreate) GetAutoresponderSubjectOk() (*string, bool)`

GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderSubject

`func (o *MailUserCreate) SetAutoresponderSubject(v string)`

SetAutoresponderSubject sets AutoresponderSubject field to given value.

### HasAutoresponderSubject

`func (o *MailUserCreate) HasAutoresponderSubject() bool`

HasAutoresponderSubject returns a boolean if a field has been set.

### GetAutoresponderMessage

`func (o *MailUserCreate) GetAutoresponderMessage() string`

GetAutoresponderMessage returns the AutoresponderMessage field if non-nil, zero value otherwise.

### GetAutoresponderMessageOk

`func (o *MailUserCreate) GetAutoresponderMessageOk() (*string, bool)`

GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderMessage

`func (o *MailUserCreate) SetAutoresponderMessage(v string)`

SetAutoresponderMessage sets AutoresponderMessage field to given value.

### HasAutoresponderMessage

`func (o *MailUserCreate) HasAutoresponderMessage() bool`

HasAutoresponderMessage returns a boolean if a field has been set.

### GetAutoresponderNoreply

`func (o *MailUserCreate) GetAutoresponderNoreply() []string`

GetAutoresponderNoreply returns the AutoresponderNoreply field if non-nil, zero value otherwise.

### GetAutoresponderNoreplyOk

`func (o *MailUserCreate) GetAutoresponderNoreplyOk() (*[]string, bool)`

GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderNoreply

`func (o *MailUserCreate) SetAutoresponderNoreply(v []string)`

SetAutoresponderNoreply sets AutoresponderNoreply field to given value.

### HasAutoresponderNoreply

`func (o *MailUserCreate) HasAutoresponderNoreply() bool`

HasAutoresponderNoreply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


