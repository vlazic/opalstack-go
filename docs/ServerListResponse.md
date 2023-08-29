# ServerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebServers** | [**[]WebServerReadResponse**](WebServerReadResponse.md) |  | 
**ImapServers** | [**[]ImapServerReadResponse**](ImapServerReadResponse.md) |  | 
**SmtpServers** | [**[]SmtpServerReadResponse**](SmtpServerReadResponse.md) |  | 

## Methods

### NewServerListResponse

`func NewServerListResponse(webServers []WebServerReadResponse, imapServers []ImapServerReadResponse, smtpServers []SmtpServerReadResponse, ) *ServerListResponse`

NewServerListResponse instantiates a new ServerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerListResponseWithDefaults

`func NewServerListResponseWithDefaults() *ServerListResponse`

NewServerListResponseWithDefaults instantiates a new ServerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebServers

`func (o *ServerListResponse) GetWebServers() []WebServerReadResponse`

GetWebServers returns the WebServers field if non-nil, zero value otherwise.

### GetWebServersOk

`func (o *ServerListResponse) GetWebServersOk() (*[]WebServerReadResponse, bool)`

GetWebServersOk returns a tuple with the WebServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServers

`func (o *ServerListResponse) SetWebServers(v []WebServerReadResponse)`

SetWebServers sets WebServers field to given value.


### GetImapServers

`func (o *ServerListResponse) GetImapServers() []ImapServerReadResponse`

GetImapServers returns the ImapServers field if non-nil, zero value otherwise.

### GetImapServersOk

`func (o *ServerListResponse) GetImapServersOk() (*[]ImapServerReadResponse, bool)`

GetImapServersOk returns a tuple with the ImapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapServers

`func (o *ServerListResponse) SetImapServers(v []ImapServerReadResponse)`

SetImapServers sets ImapServers field to given value.


### GetSmtpServers

`func (o *ServerListResponse) GetSmtpServers() []SmtpServerReadResponse`

GetSmtpServers returns the SmtpServers field if non-nil, zero value otherwise.

### GetSmtpServersOk

`func (o *ServerListResponse) GetSmtpServersOk() (*[]SmtpServerReadResponse, bool)`

GetSmtpServersOk returns a tuple with the SmtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServers

`func (o *ServerListResponse) SetSmtpServers(v []SmtpServerReadResponse)`

SetSmtpServers sets SmtpServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


