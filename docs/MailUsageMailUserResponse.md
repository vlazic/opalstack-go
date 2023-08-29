# MailUsageMailUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailuser** | [**MailUsageMailUser**](MailUsageMailUser.md) |  | 
**DiskUsageMb** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMailUsageMailUserResponse

`func NewMailUsageMailUserResponse(mailuser MailUsageMailUser, diskUsageMb int32, updatedAt time.Time, ) *MailUsageMailUserResponse`

NewMailUsageMailUserResponse instantiates a new MailUsageMailUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUsageMailUserResponseWithDefaults

`func NewMailUsageMailUserResponseWithDefaults() *MailUsageMailUserResponse`

NewMailUsageMailUserResponseWithDefaults instantiates a new MailUsageMailUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailuser

`func (o *MailUsageMailUserResponse) GetMailuser() MailUsageMailUser`

GetMailuser returns the Mailuser field if non-nil, zero value otherwise.

### GetMailuserOk

`func (o *MailUsageMailUserResponse) GetMailuserOk() (*MailUsageMailUser, bool)`

GetMailuserOk returns a tuple with the Mailuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailuser

`func (o *MailUsageMailUserResponse) SetMailuser(v MailUsageMailUser)`

SetMailuser sets Mailuser field to given value.


### GetDiskUsageMb

`func (o *MailUsageMailUserResponse) GetDiskUsageMb() int32`

GetDiskUsageMb returns the DiskUsageMb field if non-nil, zero value otherwise.

### GetDiskUsageMbOk

`func (o *MailUsageMailUserResponse) GetDiskUsageMbOk() (*int32, bool)`

GetDiskUsageMbOk returns a tuple with the DiskUsageMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsageMb

`func (o *MailUsageMailUserResponse) SetDiskUsageMb(v int32)`

SetDiskUsageMb sets DiskUsageMb field to given value.


### GetUpdatedAt

`func (o *MailUsageMailUserResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MailUsageMailUserResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MailUsageMailUserResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


