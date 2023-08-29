# SigninResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redirect** | Pointer to **string** |  | [optional] [default to "/"]
**Token** | **string** |  | 

## Methods

### NewSigninResponse

`func NewSigninResponse(token string, ) *SigninResponse`

NewSigninResponse instantiates a new SigninResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigninResponseWithDefaults

`func NewSigninResponseWithDefaults() *SigninResponse`

NewSigninResponseWithDefaults instantiates a new SigninResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirect

`func (o *SigninResponse) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SigninResponse) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SigninResponse) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *SigninResponse) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetToken

`func (o *SigninResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SigninResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SigninResponse) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


