# WebUsageOSUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Osuser** | [**WebUsageOSUser**](WebUsageOSUser.md) |  | 
**MemoryUsageMb** | **int32** |  | 
**DiskUsageMb** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWebUsageOSUserResponse

`func NewWebUsageOSUserResponse(osuser WebUsageOSUser, memoryUsageMb int32, diskUsageMb int32, updatedAt time.Time, ) *WebUsageOSUserResponse`

NewWebUsageOSUserResponse instantiates a new WebUsageOSUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebUsageOSUserResponseWithDefaults

`func NewWebUsageOSUserResponseWithDefaults() *WebUsageOSUserResponse`

NewWebUsageOSUserResponseWithDefaults instantiates a new WebUsageOSUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsuser

`func (o *WebUsageOSUserResponse) GetOsuser() WebUsageOSUser`

GetOsuser returns the Osuser field if non-nil, zero value otherwise.

### GetOsuserOk

`func (o *WebUsageOSUserResponse) GetOsuserOk() (*WebUsageOSUser, bool)`

GetOsuserOk returns a tuple with the Osuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsuser

`func (o *WebUsageOSUserResponse) SetOsuser(v WebUsageOSUser)`

SetOsuser sets Osuser field to given value.


### GetMemoryUsageMb

`func (o *WebUsageOSUserResponse) GetMemoryUsageMb() int32`

GetMemoryUsageMb returns the MemoryUsageMb field if non-nil, zero value otherwise.

### GetMemoryUsageMbOk

`func (o *WebUsageOSUserResponse) GetMemoryUsageMbOk() (*int32, bool)`

GetMemoryUsageMbOk returns a tuple with the MemoryUsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageMb

`func (o *WebUsageOSUserResponse) SetMemoryUsageMb(v int32)`

SetMemoryUsageMb sets MemoryUsageMb field to given value.


### GetDiskUsageMb

`func (o *WebUsageOSUserResponse) GetDiskUsageMb() int32`

GetDiskUsageMb returns the DiskUsageMb field if non-nil, zero value otherwise.

### GetDiskUsageMbOk

`func (o *WebUsageOSUserResponse) GetDiskUsageMbOk() (*int32, bool)`

GetDiskUsageMbOk returns a tuple with the DiskUsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsageMb

`func (o *WebUsageOSUserResponse) SetDiskUsageMb(v int32)`

SetDiskUsageMb sets DiskUsageMb field to given value.


### GetUpdatedAt

`func (o *WebUsageOSUserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebUsageOSUserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebUsageOSUserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


