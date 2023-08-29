# MailUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Name** | **string** |  | 
**ImapServer** | **string** |  | 
**Procmailrc** | **string** |  | 
**AutoresponderEnable** | **string** |  | 
**AutoresponderSubject** | **string** |  | 
**AutoresponderMessage** | **string** |  | 
**AutoresponderNoreply** | **[]string** |  | 

## Methods

### NewMailUserResponse

`func NewMailUserResponse(id string, state StateEnum, ready bool, name string, imapServer string, procmailrc string, autoresponderEnable string, autoresponderSubject string, autoresponderMessage string, autoresponderNoreply []string, ) *MailUserResponse`

NewMailUserResponse instantiates a new MailUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUserResponseWithDefaults

`func NewMailUserResponseWithDefaults() *MailUserResponse`

NewMailUserResponseWithDefaults instantiates a new MailUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MailUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *MailUserResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MailUserResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MailUserResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *MailUserResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MailUserResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MailUserResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetName

`func (o *MailUserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailUserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailUserResponse) SetName(v string)`

SetName sets Name field to given value.


### GetImapServer

`func (o *MailUserResponse) GetImapServer() string`

GetImapServer returns the ImapServer field if non-nil, zero value otherwise.

### GetImapServerOk

`func (o *MailUserResponse) GetImapServerOk() (*string, bool)`

GetImapServerOk returns a tuple with the ImapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapServer

`func (o *MailUserResponse) SetImapServer(v string)`

SetImapServer sets ImapServer field to given value.


### GetProcmailrc

`func (o *MailUserResponse) GetProcmailrc() string`

GetProcmailrc returns the Procmailrc field if non-nil, zero value otherwise.

### GetProcmailrcOk

`func (o *MailUserResponse) GetProcmailrcOk() (*string, bool)`

GetProcmailrcOk returns a tuple with the Procmailrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcmailrc

`func (o *MailUserResponse) SetProcmailrc(v string)`

SetProcmailrc sets Procmailrc field to given value.


### GetAutoresponderEnable

`func (o *MailUserResponse) GetAutoresponderEnable() string`

GetAutoresponderEnable returns the AutoresponderEnable field if non-nil, zero value otherwise.

### GetAutoresponderEnableOk

`func (o *MailUserResponse) GetAutoresponderEnableOk() (*string, bool)`

GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderEnable

`func (o *MailUserResponse) SetAutoresponderEnable(v string)`

SetAutoresponderEnable sets AutoresponderEnable field to given value.


### GetAutoresponderSubject

`func (o *MailUserResponse) GetAutoresponderSubject() string`

GetAutoresponderSubject returns the AutoresponderSubject field if non-nil, zero value otherwise.

### GetAutoresponderSubjectOk

`func (o *MailUserResponse) GetAutoresponderSubjectOk() (*string, bool)`

GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderSubject

`func (o *MailUserResponse) SetAutoresponderSubject(v string)`

SetAutoresponderSubject sets AutoresponderSubject field to given value.


### GetAutoresponderMessage

`func (o *MailUserResponse) GetAutoresponderMessage() string`

GetAutoresponderMessage returns the AutoresponderMessage field if non-nil, zero value otherwise.

### GetAutoresponderMessageOk

`func (o *MailUserResponse) GetAutoresponderMessageOk() (*string, bool)`

GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderMessage

`func (o *MailUserResponse) SetAutoresponderMessage(v string)`

SetAutoresponderMessage sets AutoresponderMessage field to given value.


### GetAutoresponderNoreply

`func (o *MailUserResponse) GetAutoresponderNoreply() []string`

GetAutoresponderNoreply returns the AutoresponderNoreply field if non-nil, zero value otherwise.

### GetAutoresponderNoreplyOk

`func (o *MailUserResponse) GetAutoresponderNoreplyOk() (*[]string, bool)`

GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderNoreply

`func (o *MailUserResponse) SetAutoresponderNoreply(v []string)`

SetAutoresponderNoreply sets AutoresponderNoreply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


