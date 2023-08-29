# \DomainAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainCreate**](DomainAPI.md#DomainCreate) | **Post** /api/v1/domain/create/ | 
[**DomainDelete**](DomainAPI.md#DomainDelete) | **Post** /api/v1/domain/delete/ | 
[**DomainList**](DomainAPI.md#DomainList) | **Get** /api/v1/domain/list/ | 
[**DomainRead**](DomainAPI.md#DomainRead) | **Get** /api/v1/domain/read/{uuid} | 



## DomainCreate

> []DomainResponse DomainCreate(ctx).DomainCreate(domainCreate).Execute()



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
    domainCreate := []openapiclient.DomainCreate{*openapiclient.NewDomainCreate("Name_example")} // []DomainCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainCreate(context.Background()).DomainCreate(domainCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainCreate`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainCreate** | [**[]DomainCreate**](DomainCreate.md) |  | 

### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDelete

> DomainDelete(ctx).DomainDelete(domainDelete).Execute()



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
    domainDelete := []openapiclient.DomainDelete{*openapiclient.NewDomainDelete("Id_example")} // []DomainDelete | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainAPI.DomainDelete(context.Background()).DomainDelete(domainDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainDelete** | [**[]DomainDelete**](DomainDelete.md) |  | 

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


## DomainList

> []DomainResponse DomainList(ctx).Execute()



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
    resp, r, err := apiClient.DomainAPI.DomainList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainList`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDomainListRequest struct via the builder pattern


### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainRead

> DomainResponse DomainRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.DomainAPI.DomainRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainRead`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

