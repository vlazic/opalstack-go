# WebserverUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] [default to "opal1.opalstack.com"]
**DiskTotal** | **int32** |  | 
**DiskUsed** | **int32** |  | 
**DiskAvailable** | **int32** |  | 
**RssTotal** | **int32** |  | 
**RssUsed** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "W"]

## Methods

### NewWebserverUsageResponse

`func NewWebserverUsageResponse(diskTotal int32, diskUsed int32, diskAvailable int32, rssTotal int32, rssUsed int32, ) *WebserverUsageResponse`

NewWebserverUsageResponse instantiates a new WebserverUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebserverUsageResponseWithDefaults

`func NewWebserverUsageResponseWithDefaults() *WebserverUsageResponse`

NewWebserverUsageResponseWithDefaults instantiates a new WebserverUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *WebserverUsageResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *WebserverUsageResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *WebserverUsageResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *WebserverUsageResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDiskTotal

`func (o *WebserverUsageResponse) GetDiskTotal() int32`

GetDiskTotal returns the DiskTotal field if non-nil, zero value otherwise.

### GetDiskTotalOk

`func (o *WebserverUsageResponse) GetDiskTotalOk() (*int32, bool)`

GetDiskTotalOk returns a tuple with the DiskTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTotal

`func (o *WebserverUsageResponse) SetDiskTotal(v int32)`

SetDiskTotal sets DiskTotal field to given value.


### GetDiskUsed

`func (o *WebserverUsageResponse) GetDiskUsed() int32`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *WebserverUsageResponse) GetDiskUsedOk() (*int32, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *WebserverUsageResponse) SetDiskUsed(v int32)`

SetDiskUsed sets DiskUsed field to given value.


### GetDiskAvailable

`func (o *WebserverUsageResponse) GetDiskAvailable() int32`

GetDiskAvailable returns the DiskAvailable field if non-nil, zero value otherwise.

### GetDiskAvailableOk

`func (o *WebserverUsageResponse) GetDiskAvailableOk() (*int32, bool)`

GetDiskAvailableOk returns a tuple with the DiskAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAvailable

`func (o *WebserverUsageResponse) SetDiskAvailable(v int32)`

SetDiskAvailable sets DiskAvailable field to given value.


### GetRssTotal

`func (o *WebserverUsageResponse) GetRssTotal() int32`

GetRssTotal returns the RssTotal field if non-nil, zero value otherwise.

### GetRssTotalOk

`func (o *WebserverUsageResponse) GetRssTotalOk() (*int32, bool)`

GetRssTotalOk returns a tuple with the RssTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssTotal

`func (o *WebserverUsageResponse) SetRssTotal(v int32)`

SetRssTotal sets RssTotal field to given value.


### GetRssUsed

`func (o *WebserverUsageResponse) GetRssUsed() int32`

GetRssUsed returns the RssUsed field if non-nil, zero value otherwise.

### GetRssUsedOk

`func (o *WebserverUsageResponse) GetRssUsedOk() (*int32, bool)`

GetRssUsedOk returns a tuple with the RssUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssUsed

`func (o *WebserverUsageResponse) SetRssUsed(v int32)`

SetRssUsed sets RssUsed field to given value.


### GetType

`func (o *WebserverUsageResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebserverUsageResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebserverUsageResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebserverUsageResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


