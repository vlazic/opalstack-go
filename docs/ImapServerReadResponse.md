# ImapServerReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] [default to "imap1.us.opalstack.com"]
**Type** | Pointer to **string** |  | [optional] [default to "imap"]

## Methods

### NewImapServerReadResponse

`func NewImapServerReadResponse(id string, ) *ImapServerReadResponse`

NewImapServerReadResponse instantiates a new ImapServerReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapServerReadResponseWithDefaults

`func NewImapServerReadResponseWithDefaults() *ImapServerReadResponse`

NewImapServerReadResponseWithDefaults instantiates a new ImapServerReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImapServerReadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImapServerReadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImapServerReadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *ImapServerReadResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ImapServerReadResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ImapServerReadResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ImapServerReadResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetType

`func (o *ImapServerReadResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImapServerReadResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImapServerReadResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImapServerReadResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


