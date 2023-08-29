# MailUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailusers** | [**[]MailUsageMailUserResponse**](MailUsageMailUserResponse.md) |  | 

## Methods

### NewMailUsageResponse

`func NewMailUsageResponse(mailusers []MailUsageMailUserResponse, ) *MailUsageResponse`

NewMailUsageResponse instantiates a new MailUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailUsageResponseWithDefaults

`func NewMailUsageResponseWithDefaults() *MailUsageResponse`

NewMailUsageResponseWithDefaults instantiates a new MailUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailusers

`func (o *MailUsageResponse) GetMailusers() []MailUsageMailUserResponse`

GetMailusers returns the Mailusers field if non-nil, zero value otherwise.

### GetMailusersOk

`func (o *MailUsageResponse) GetMailusersOk() (*[]MailUsageMailUserResponse, bool)`

GetMailusersOk returns a tuple with the Mailusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailusers

`func (o *MailUsageResponse) SetMailusers(v []MailUsageMailUserResponse)`

SetMailusers sets Mailusers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


