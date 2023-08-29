# UsageDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImapTotal** | **int32** |  | 
**ImapUsed** | **int32** |  | 
**ImapAvailable** | **int32** |  | 
**Webservers** | [**WebserverIdUsageResponse**](WebserverIdUsageResponse.md) |  | 

## Methods

### NewUsageDataResponse

`func NewUsageDataResponse(imapTotal int32, imapUsed int32, imapAvailable int32, webservers WebserverIdUsageResponse, ) *UsageDataResponse`

NewUsageDataResponse instantiates a new UsageDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDataResponseWithDefaults

`func NewUsageDataResponseWithDefaults() *UsageDataResponse`

NewUsageDataResponseWithDefaults instantiates a new UsageDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImapTotal

`func (o *UsageDataResponse) GetImapTotal() int32`

GetImapTotal returns the ImapTotal field if non-nil, zero value otherwise.

### GetImapTotalOk

`func (o *UsageDataResponse) GetImapTotalOk() (*int32, bool)`

GetImapTotalOk returns a tuple with the ImapTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapTotal

`func (o *UsageDataResponse) SetImapTotal(v int32)`

SetImapTotal sets ImapTotal field to given value.


### GetImapUsed

`func (o *UsageDataResponse) GetImapUsed() int32`

GetImapUsed returns the ImapUsed field if non-nil, zero value otherwise.

### GetImapUsedOk

`func (o *UsageDataResponse) GetImapUsedOk() (*int32, bool)`

GetImapUsedOk returns a tuple with the ImapUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapUsed

`func (o *UsageDataResponse) SetImapUsed(v int32)`

SetImapUsed sets ImapUsed field to given value.


### GetImapAvailable

`func (o *UsageDataResponse) GetImapAvailable() int32`

GetImapAvailable returns the ImapAvailable field if non-nil, zero value otherwise.

### GetImapAvailableOk

`func (o *UsageDataResponse) GetImapAvailableOk() (*int32, bool)`

GetImapAvailableOk returns a tuple with the ImapAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapAvailable

`func (o *UsageDataResponse) SetImapAvailable(v int32)`

SetImapAvailable sets ImapAvailable field to given value.


### GetWebservers

`func (o *UsageDataResponse) GetWebservers() WebserverIdUsageResponse`

GetWebservers returns the Webservers field if non-nil, zero value otherwise.

### GetWebserversOk

`func (o *UsageDataResponse) GetWebserversOk() (*WebserverIdUsageResponse, bool)`

GetWebserversOk returns a tuple with the Webservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebservers

`func (o *UsageDataResponse) SetWebservers(v WebserverIdUsageResponse)`

SetWebservers sets Webservers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


