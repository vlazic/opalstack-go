# WebUsagePsqlDBResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Psqldb** | [**WebUsagePsqlDB**](WebUsagePsqlDB.md) |  | 
**DiskUsageMb** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWebUsagePsqlDBResponse

`func NewWebUsagePsqlDBResponse(psqldb WebUsagePsqlDB, diskUsageMb int32, updatedAt time.Time, ) *WebUsagePsqlDBResponse`

NewWebUsagePsqlDBResponse instantiates a new WebUsagePsqlDBResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebUsagePsqlDBResponseWithDefaults

`func NewWebUsagePsqlDBResponseWithDefaults() *WebUsagePsqlDBResponse`

NewWebUsagePsqlDBResponseWithDefaults instantiates a new WebUsagePsqlDBResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPsqldb

`func (o *WebUsagePsqlDBResponse) GetPsqldb() WebUsagePsqlDB`

GetPsqldb returns the Psqldb field if non-nil, zero value otherwise.

### GetPsqldbOk

`func (o *WebUsagePsqlDBResponse) GetPsqldbOk() (*WebUsagePsqlDB, bool)`

GetPsqldbOk returns a tuple with the Psqldb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsqldb

`func (o *WebUsagePsqlDBResponse) SetPsqldb(v WebUsagePsqlDB)`

SetPsqldb sets Psqldb field to given value.


### GetDiskUsageMb

`func (o *WebUsagePsqlDBResponse) GetDiskUsageMb() int32`

GetDiskUsageMb returns the DiskUsageMb field if non-nil, zero value otherwise.

### GetDiskUsageMbOk

`func (o *WebUsagePsqlDBResponse) GetDiskUsageMbOk() (*int32, bool)`

GetDiskUsageMbOk returns a tuple with the DiskUsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsageMb

`func (o *WebUsagePsqlDBResponse) SetDiskUsageMb(v int32)`

SetDiskUsageMb sets DiskUsageMb field to given value.


### GetUpdatedAt

`func (o *WebUsagePsqlDBResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebUsagePsqlDBResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebUsagePsqlDBResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


