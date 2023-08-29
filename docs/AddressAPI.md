# \AddressAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressCreate**](AddressAPI.md#AddressCreate) | **Post** /api/v1/address/create/ | 
[**AddressDelete**](AddressAPI.md#AddressDelete) | **Post** /api/v1/address/delete/ | 
[**AddressList**](AddressAPI.md#AddressList) | **Get** /api/v1/address/list/ | 
[**AddressRead**](AddressAPI.md#AddressRead) | **Get** /api/v1/address/read/{uuid} | 
[**AddressUpdate**](AddressAPI.md#AddressUpdate) | **Post** /api/v1/address/update/ | 



## AddressCreate

> []VirtualAliasResponse AddressCreate(ctx).VirtualAliasCreate(virtualAliasCreate).Execute()



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
    virtualAliasCreate := []openapiclient.VirtualAliasCreate{*openapiclient.NewVirtualAliasCreate("Source_example", []string{"Forwards_example"})} // []VirtualAliasCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressAPI.AddressCreate(context.Background()).VirtualAliasCreate(virtualAliasCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressCreate`: []VirtualAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAliasCreate** | [**[]VirtualAliasCreate**](VirtualAliasCreate.md) |  | 

### Return type

[**[]VirtualAliasResponse**](VirtualAliasResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressDelete

> AddressDelete(ctx).VirtualAliasRead(virtualAliasRead).Execute()



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
    virtualAliasRead := []openapiclient.VirtualAliasRead{*openapiclient.NewVirtualAliasRead("Id_example")} // []VirtualAliasRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressAPI.AddressDelete(context.Background()).VirtualAliasRead(virtualAliasRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAliasRead** | [**[]VirtualAliasRead**](VirtualAliasRead.md) |  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressList

> []VirtualAliasResponse AddressList(ctx).Execute()



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
    resp, r, err := apiClient.AddressAPI.AddressList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressList`: []VirtualAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddressListRequest struct via the builder pattern


### Return type

[**[]VirtualAliasResponse**](VirtualAliasResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressRead

> VirtualAliasResponse AddressRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.AddressAPI.AddressRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressRead`: VirtualAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualAliasResponse**](VirtualAliasResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressUpdate

> []VirtualAliasResponse AddressUpdate(ctx).VirtualAliasUpdate(virtualAliasUpdate).Execute()



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
    virtualAliasUpdate := []openapiclient.VirtualAliasUpdate{*openapiclient.NewVirtualAliasUpdate("Id_example")} // []VirtualAliasUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressAPI.AddressUpdate(context.Background()).VirtualAliasUpdate(virtualAliasUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressUpdate`: []VirtualAliasResponse
    fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAliasUpdate** | [**[]VirtualAliasUpdate**](VirtualAliasUpdate.md) |  | 

### Return type

[**[]VirtualAliasResponse**](VirtualAliasResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

