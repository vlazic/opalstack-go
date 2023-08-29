# QuarantinedMailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | [**QuarantinedMailStateEnum**](QuarantinedMailStateEnum.md) |  | 
**Ready** | **bool** |  | 
**SmtpServer** | **string** |  | 
**QueueId** | **string** |  | 
**ArrivalTime** | **time.Time** |  | 
**MessageSizeBytes** | **int32** |  | 
**ClientAddress** | **string** |  | 
**SaslUsername** | **string** |  | 
**Sender** | **string** |  | 
**Recipients** | **[]string** |  | 
**SpamStatus** | **bool** |  | 
**SpamScore** | **float32** |  | 
**Subject** | **string** |  | 
**Body** | **string** |  | 
**Train** | **bool** |  | 

## Methods

### NewQuarantinedMailResponse

`func NewQuarantinedMailResponse(id string, state QuarantinedMailStateEnum, ready bool, smtpServer string, queueId string, arrivalTime time.Time, messageSizeBytes int32, clientAddress string, saslUsername string, sender string, recipients []string, spamStatus bool, spamScore float32, subject string, body string, train bool, ) *QuarantinedMailResponse`

NewQuarantinedMailResponse instantiates a new QuarantinedMailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantinedMailResponseWithDefaults

`func NewQuarantinedMailResponseWithDefaults() *QuarantinedMailResponse`

NewQuarantinedMailResponseWithDefaults instantiates a new QuarantinedMailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuarantinedMailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuarantinedMailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuarantinedMailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *QuarantinedMailResponse) GetState() QuarantinedMailStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuarantinedMailResponse) GetStateOk() (*QuarantinedMailStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuarantinedMailResponse) SetState(v QuarantinedMailStateEnum)`

SetState sets State field to given value.


### GetReady

`func (o *QuarantinedMailResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *QuarantinedMailResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *QuarantinedMailResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetSmtpServer

`func (o *QuarantinedMailResponse) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *QuarantinedMailResponse) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *QuarantinedMailResponse) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.


### GetQueueId

`func (o *QuarantinedMailResponse) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *QuarantinedMailResponse) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *QuarantinedMailResponse) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.


### GetArrivalTime

`func (o *QuarantinedMailResponse) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *QuarantinedMailResponse) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *QuarantinedMailResponse) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetMessageSizeBytes

`func (o *QuarantinedMailResponse) GetMessageSizeBytes() int32`

GetMessageSizeBytes returns the MessageSizeBytes field if non-nil, zero value otherwise.

### GetMessageSizeBytesOk

`func (o *QuarantinedMailResponse) GetMessageSizeBytesOk() (*int32, bool)`

GetMessageSizeBytesOk returns a tuple with the MessageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSizeBytes

`func (o *QuarantinedMailResponse) SetMessageSizeBytes(v int32)`

SetMessageSizeBytes sets MessageSizeBytes field to given value.


### GetClientAddress

`func (o *QuarantinedMailResponse) GetClientAddress() string`

GetClientAddress returns the ClientAddress field if non-nil, zero value otherwise.

### GetClientAddressOk

`func (o *QuarantinedMailResponse) GetClientAddressOk() (*string, bool)`

GetClientAddressOk returns a tuple with the ClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAddress

`func (o *QuarantinedMailResponse) SetClientAddress(v string)`

SetClientAddress sets ClientAddress field to given value.


### GetSaslUsername

`func (o *QuarantinedMailResponse) GetSaslUsername() string`

GetSaslUsername returns the SaslUsername field if non-nil, zero value otherwise.

### GetSaslUsernameOk

`func (o *QuarantinedMailResponse) GetSaslUsernameOk() (*string, bool)`

GetSaslUsernameOk returns a tuple with the SaslUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaslUsername

`func (o *QuarantinedMailResponse) SetSaslUsername(v string)`

SetSaslUsername sets SaslUsername field to given value.


### GetSender

`func (o *QuarantinedMailResponse) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *QuarantinedMailResponse) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *QuarantinedMailResponse) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipients

`func (o *QuarantinedMailResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *QuarantinedMailResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *QuarantinedMailResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetSpamStatus

`func (o *QuarantinedMailResponse) GetSpamStatus() bool`

GetSpamStatus returns the SpamStatus field if non-nil, zero value otherwise.

### GetSpamStatusOk

`func (o *QuarantinedMailResponse) GetSpamStatusOk() (*bool, bool)`

GetSpamStatusOk returns a tuple with the SpamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamStatus

`func (o *QuarantinedMailResponse) SetSpamStatus(v bool)`

SetSpamStatus sets SpamStatus field to given value.


### GetSpamScore

`func (o *QuarantinedMailResponse) GetSpamScore() float32`

GetSpamScore returns the SpamScore field if non-nil, zero value otherwise.

### GetSpamScoreOk

`func (o *QuarantinedMailResponse) GetSpamScoreOk() (*float32, bool)`

GetSpamScoreOk returns a tuple with the SpamScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamScore

`func (o *QuarantinedMailResponse) SetSpamScore(v float32)`

SetSpamScore sets SpamScore field to given value.


### GetSubject

`func (o *QuarantinedMailResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *QuarantinedMailResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *QuarantinedMailResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *QuarantinedMailResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *QuarantinedMailResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *QuarantinedMailResponse) SetBody(v string)`

SetBody sets Body field to given value.


### GetTrain

`func (o *QuarantinedMailResponse) GetTrain() bool`

GetTrain returns the Train field if non-nil, zero value otherwise.

### GetTrainOk

`func (o *QuarantinedMailResponse) GetTrainOk() (*bool, bool)`

GetTrainOk returns a tuple with the Train field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrain

`func (o *QuarantinedMailResponse) SetTrain(v bool)`

SetTrain sets Train field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


