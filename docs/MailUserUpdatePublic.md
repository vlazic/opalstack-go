# MailUserUpdatePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CurrentPassword** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Procmailrc** | Pointer to **string** |  | [optional] 
**AutoresponderEnable** | Pointer to **bool** |  | [optional] 
**AutoresponderSubject** | Pointer to **string** |  | [optional] 
**AutoresponderMessage** | Pointer to **string** |  | [optional] 
**AutoresponderNoreply** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMailUserUpdatePublic

`func NewMailUserUpdatePublic(name string, currentPassword string, ) *MailUserUpdatePublic`

NewMailUserUpdatePublic instantiates a new MailUserUpdatePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUserUpdatePublicWithDefaults

`func NewMailUserUpdatePublicWithDefaults() *MailUserUpdatePublic`

NewMailUserUpdatePublicWithDefaults instantiates a new MailUserUpdatePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MailUserUpdatePublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailUserUpdatePublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailUserUpdatePublic) SetName(v string)`

SetName sets Name field to given value.


### GetCurrentPassword

`func (o *MailUserUpdatePublic) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *MailUserUpdatePublic) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *MailUserUpdatePublic) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.


### GetPassword

`func (o *MailUserUpdatePublic) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MailUserUpdatePublic) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MailUserUpdatePublic) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MailUserUpdatePublic) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProcmailrc

`func (o *MailUserUpdatePublic) GetProcmailrc() string`

GetProcmailrc returns the Procmailrc field if non-nil, zero value otherwise.

### GetProcmailrcOk

`func (o *MailUserUpdatePublic) GetProcmailrcOk() (*string, bool)`

GetProcmailrcOk returns a tuple with the Procmailrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcmailrc

`func (o *MailUserUpdatePublic) SetProcmailrc(v string)`

SetProcmailrc sets Procmailrc field to given value.

### HasProcmailrc

`func (o *MailUserUpdatePublic) HasProcmailrc() bool`

HasProcmailrc returns a boolean if a field has been set.

### GetAutoresponderEnable

`func (o *MailUserUpdatePublic) GetAutoresponderEnable() bool`

GetAutoresponderEnable returns the AutoresponderEnable field if non-nil, zero value otherwise.

### GetAutoresponderEnableOk

`func (o *MailUserUpdatePublic) GetAutoresponderEnableOk() (*bool, bool)`

GetAutoresponderEnableOk returns a tuple with the AutoresponderEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderEnable

`func (o *MailUserUpdatePublic) SetAutoresponderEnable(v bool)`

SetAutoresponderEnable sets AutoresponderEnable field to given value.

### HasAutoresponderEnable

`func (o *MailUserUpdatePublic) HasAutoresponderEnable() bool`

HasAutoresponderEnable returns a boolean if a field has been set.

### GetAutoresponderSubject

`func (o *MailUserUpdatePublic) GetAutoresponderSubject() string`

GetAutoresponderSubject returns the AutoresponderSubject field if non-nil, zero value otherwise.

### GetAutoresponderSubjectOk

`func (o *MailUserUpdatePublic) GetAutoresponderSubjectOk() (*string, bool)`

GetAutoresponderSubjectOk returns a tuple with the AutoresponderSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderSubject

`func (o *MailUserUpdatePublic) SetAutoresponderSubject(v string)`

SetAutoresponderSubject sets AutoresponderSubject field to given value.

### HasAutoresponderSubject

`func (o *MailUserUpdatePublic) HasAutoresponderSubject() bool`

HasAutoresponderSubject returns a boolean if a field has been set.

### GetAutoresponderMessage

`func (o *MailUserUpdatePublic) GetAutoresponderMessage() string`

GetAutoresponderMessage returns the AutoresponderMessage field if non-nil, zero value otherwise.

### GetAutoresponderMessageOk

`func (o *MailUserUpdatePublic) GetAutoresponderMessageOk() (*string, bool)`

GetAutoresponderMessageOk returns a tuple with the AutoresponderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderMessage

`func (o *MailUserUpdatePublic) SetAutoresponderMessage(v string)`

SetAutoresponderMessage sets AutoresponderMessage field to given value.

### HasAutoresponderMessage

`func (o *MailUserUpdatePublic) HasAutoresponderMessage() bool`

HasAutoresponderMessage returns a boolean if a field has been set.

### GetAutoresponderNoreply

`func (o *MailUserUpdatePublic) GetAutoresponderNoreply() []string`

GetAutoresponderNoreply returns the AutoresponderNoreply field if non-nil, zero value otherwise.

### GetAutoresponderNoreplyOk

`func (o *MailUserUpdatePublic) GetAutoresponderNoreplyOk() (*[]string, bool)`

GetAutoresponderNoreplyOk returns a tuple with the AutoresponderNoreply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoresponderNoreply

`func (o *MailUserUpdatePublic) SetAutoresponderNoreply(v []string)`

SetAutoresponderNoreply sets AutoresponderNoreply field to given value.

### HasAutoresponderNoreply

`func (o *MailUserUpdatePublic) HasAutoresponderNoreply() bool`

HasAutoresponderNoreply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


