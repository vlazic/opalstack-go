# MailUsageMailUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ImapServer** | [**MailUsageImapServer**](MailUsageImapServer.md) |  | 

## Methods

### NewMailUsageMailUser

`func NewMailUsageMailUser(id string, name string, imapServer MailUsageImapServer, ) *MailUsageMailUser`

NewMailUsageMailUser instantiates a new MailUsageMailUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUsageMailUserWithDefaults

`func NewMailUsageMailUserWithDefaults() *MailUsageMailUser`

NewMailUsageMailUserWithDefaults instantiates a new MailUsageMailUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MailUsageMailUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailUsageMailUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailUsageMailUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MailUsageMailUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MailUsageMailUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MailUsageMailUser) SetName(v string)`

SetName sets Name field to given value.


### GetImapServer

`func (o *MailUsageMailUser) GetImapServer() MailUsageImapServer`

GetImapServer returns the ImapServer field if non-nil, zero value otherwise.

### GetImapServerOk

`func (o *MailUsageMailUser) GetImapServerOk() (*MailUsageImapServer, bool)`

GetImapServerOk returns a tuple with the ImapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapServer

`func (o *MailUsageMailUser) SetImapServer(v MailUsageImapServer)`

SetImapServer sets ImapServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


