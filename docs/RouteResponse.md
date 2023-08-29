# RouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**App** | **string** |  | 
**Uri** | Pointer to **string** |  | [optional] [default to "/"]

## Methods

### NewRouteResponse

`func NewRouteResponse(id string, app string, ) *RouteResponse`

NewRouteResponse instantiates a new RouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteResponseWithDefaults

`func NewRouteResponseWithDefaults() *RouteResponse`

NewRouteResponseWithDefaults instantiates a new RouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApp

`func (o *RouteResponse) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *RouteResponse) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *RouteResponse) SetApp(v string)`

SetApp sets App field to given value.


### GetUri

`func (o *RouteResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RouteResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RouteResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RouteResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


