# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**StateEnum**](StateEnum.md) |  | 
**Ready** | **bool** |  | 
**Email** | **string** |  | 
**WebServers** | **[]string** |  | 
**ImapServers** | **[]string** |  | 
**SmtpServers** | **[]string** |  | 
**Ips** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**PaymentProcessor** | [**AccountResponsePaymentProcessorEnum**](AccountResponsePaymentProcessorEnum.md) |  | 
**UsageData** | [**UsageDataResponse**](UsageDataResponse.md) |  | 

## Methods

### NewAccountResponse

`func NewAccountResponse(id string, state StateEnum, ready bool, email string, webServers []string, imapServers []string, smtpServers []string, ips []string, createdAt time.Time, paymentProcessor AccountResponsePaymentProcessorEnum, usageData UsageDataResponse, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *AccountResponse) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountResponse) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountResponse) SetState(v StateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *AccountResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AccountResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AccountResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetEmail

`func (o *AccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetWebServers

`func (o *AccountResponse) GetWebServers() []string`

GetWebServers returns the WebServers field if non-nil, zero value otherwise.

### GetWebServersOk

`func (o *AccountResponse) GetWebServersOk() (*[]string, bool)`

GetWebServersOk returns a tuple with the WebServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServers

`func (o *AccountResponse) SetWebServers(v []string)`

SetWebServers sets WebServers field to given value.


### GetImapServers

`func (o *AccountResponse) GetImapServers() []string`

GetImapServers returns the ImapServers field if non-nil, zero value otherwise.

### GetImapServersOk

`func (o *AccountResponse) GetImapServersOk() (*[]string, bool)`

GetImapServersOk returns a tuple with the ImapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapServers

`func (o *AccountResponse) SetImapServers(v []string)`

SetImapServers sets ImapServers field to given value.


### GetSmtpServers

`func (o *AccountResponse) GetSmtpServers() []string`

GetSmtpServers returns the SmtpServers field if non-nil, zero value otherwise.

### GetSmtpServersOk

`func (o *AccountResponse) GetSmtpServersOk() (*[]string, bool)`

GetSmtpServersOk returns a tuple with the SmtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServers

`func (o *AccountResponse) SetSmtpServers(v []string)`

SetSmtpServers sets SmtpServers field to given value.


### GetIps

`func (o *AccountResponse) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *AccountResponse) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *AccountResponse) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetCreatedAt

`func (o *AccountResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPaymentProcessor

`func (o *AccountResponse) GetPaymentProcessor() AccountResponsePaymentProcessorEnum`

GetPaymentProcessor returns the PaymentProcessor field if non-nil, zero value otherwise.

### GetPaymentProcessorOk

`func (o *AccountResponse) GetPaymentProcessorOk() (*AccountResponsePaymentProcessorEnum, bool)`

GetPaymentProcessorOk returns a tuple with the PaymentProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessor

`func (o *AccountResponse) SetPaymentProcessor(v AccountResponsePaymentProcessorEnum)`

SetPaymentProcessor sets PaymentProcessor field to given value.


### GetUsageData

`func (o *AccountResponse) GetUsageData() UsageDataResponse`

GetUsageData returns the UsageData field if non-nil, zero value otherwise.

### GetUsageDataOk

`func (o *AccountResponse) GetUsageDataOk() (*UsageDataResponse, bool)`

GetUsageDataOk returns a tuple with the UsageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageData

`func (o *AccountResponse) SetUsageData(v UsageDataResponse)`

SetUsageData sets UsageData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


