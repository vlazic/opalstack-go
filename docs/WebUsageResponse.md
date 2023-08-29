# WebUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Osusers** | [**[]WebUsageOSUserResponse**](WebUsageOSUserResponse.md) |  | 
**Mariadbs** | [**[]WebUsageMariaDBResponse**](WebUsageMariaDBResponse.md) |  | 
**Psqldbs** | [**[]WebUsagePsqlDBResponse**](WebUsagePsqlDBResponse.md) |  | 

## Methods

### NewWebUsageResponse

`func NewWebUsageResponse(osusers []WebUsageOSUserResponse, mariadbs []WebUsageMariaDBResponse, psqldbs []WebUsagePsqlDBResponse, ) *WebUsageResponse`

NewWebUsageResponse instantiates a new WebUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebUsageResponseWithDefaults

`func NewWebUsageResponseWithDefaults() *WebUsageResponse`

NewWebUsageResponseWithDefaults instantiates a new WebUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsusers

`func (o *WebUsageResponse) GetOsusers() []WebUsageOSUserResponse`

GetOsusers returns the Osusers field if non-nil, zero value otherwise.

### GetOsusersOk

`func (o *WebUsageResponse) GetOsusersOk() (*[]WebUsageOSUserResponse, bool)`

GetOsusersOk returns a tuple with the Osusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsusers

`func (o *WebUsageResponse) SetOsusers(v []WebUsageOSUserResponse)`

SetOsusers sets Osusers field to given value.


### GetMariadbs

`func (o *WebUsageResponse) GetMariadbs() []WebUsageMariaDBResponse`

GetMariadbs returns the Mariadbs field if non-nil, zero value otherwise.

### GetMariadbsOk

`func (o *WebUsageResponse) GetMariadbsOk() (*[]WebUsageMariaDBResponse, bool)`

GetMariadbsOk returns a tuple with the Mariadbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMariadbs

`func (o *WebUsageResponse) SetMariadbs(v []WebUsageMariaDBResponse)`

SetMariadbs sets Mariadbs field to given value.


### GetPsqldbs

`func (o *WebUsageResponse) GetPsqldbs() []WebUsagePsqlDBResponse`

GetPsqldbs returns the Psqldbs field if non-nil, zero value otherwise.

### GetPsqldbsOk

`func (o *WebUsageResponse) GetPsqldbsOk() (*[]WebUsagePsqlDBResponse, bool)`

GetPsqldbsOk returns a tuple with the Psqldbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsqldbs

`func (o *WebUsageResponse) SetPsqldbs(v []WebUsagePsqlDBResponse)`

SetPsqldbs sets Psqldbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


