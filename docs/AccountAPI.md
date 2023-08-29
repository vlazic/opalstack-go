# \AccountAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountList**](AccountAPI.md#AccountList) | **Get** /api/v1/account/list/ | 
[**AccountRead**](AccountAPI.md#AccountRead) | **Get** /api/v1/account/read/{uuid} | 
[**AccountReadCurrent**](AccountAPI.md#AccountReadCurrent) | **Get** /api/v1/account/info/ | 
[**AccountReadCurrent2**](AccountAPI.md#AccountReadCurrent2) | **Get** /api/v1/account/read/current/ | 
[**AccountUpdate**](AccountAPI.md#AccountUpdate) | **Post** /api/v1/account/update/ | 
[**AccountUpdateCurrent**](AccountAPI.md#AccountUpdateCurrent) | **Post** /api/v1/account/update/current/ | 



## AccountList

> []AccountResponse AccountList(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountList`: []AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountListRequest struct via the builder pattern


### Return type

[**[]AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountRead

> AccountResponse AccountRead(ctx, uuid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountRead`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountReadCurrent

> AccountResponse AccountReadCurrent(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountReadCurrent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountReadCurrent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountReadCurrent`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountReadCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountReadCurrentRequest struct via the builder pattern


### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountReadCurrent2

> AccountResponse AccountReadCurrent2(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountReadCurrent2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountReadCurrent2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountReadCurrent2`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountReadCurrent2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountReadCurrent2Request struct via the builder pattern


### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountUpdate

> []AccountResponse AccountUpdate(ctx).AccountUpdate(accountUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {
    accountUpdate := []openapiclient.AccountUpdate{*openapiclient.NewAccountUpdate("Id_example")} // []AccountUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountUpdate(context.Background()).AccountUpdate(accountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountUpdate`: []AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUpdate** | [**[]AccountUpdate**](AccountUpdate.md) |  | 

### Return type

[**[]AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountUpdateCurrent

> AccountResponse AccountUpdateCurrent(ctx).AccountUpdate(accountUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/vlazic/opalstack-go"
)

func main() {
    accountUpdate := *openapiclient.NewAccountUpdate("Id_example") // AccountUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AccountUpdateCurrent(context.Background()).AccountUpdate(accountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountUpdateCurrent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountUpdateCurrent`: AccountResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountUpdateCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountUpdateCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUpdate** | [**AccountUpdate**](AccountUpdate.md) |  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

