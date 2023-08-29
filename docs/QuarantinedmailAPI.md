# \QuarantinedmailAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuarantinedmailDelete**](QuarantinedmailAPI.md#QuarantinedmailDelete) | **Post** /api/v1/quarantinedmail/delete/ | 
[**QuarantinedmailList**](QuarantinedmailAPI.md#QuarantinedmailList) | **Get** /api/v1/quarantinedmail/list/ | 
[**QuarantinedmailRead**](QuarantinedmailAPI.md#QuarantinedmailRead) | **Get** /api/v1/quarantinedmail/read/{uuid} | 
[**QuarantinedmailSubmit**](QuarantinedmailAPI.md#QuarantinedmailSubmit) | **Post** /api/v1/quarantinedmail/submit/ | 



## QuarantinedmailDelete

> QuarantinedmailDelete(ctx).QuarantinedMailRead(quarantinedMailRead).Execute()



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
    quarantinedMailRead := []openapiclient.QuarantinedMailRead{*openapiclient.NewQuarantinedMailRead("Id_example")} // []QuarantinedMailRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.QuarantinedmailAPI.QuarantinedmailDelete(context.Background()).QuarantinedMailRead(quarantinedMailRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantinedmailAPI.QuarantinedmailDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuarantinedmailDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quarantinedMailRead** | [**[]QuarantinedMailRead**](QuarantinedMailRead.md) |  | 

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


## QuarantinedmailList

> []QuarantinedMailResponse QuarantinedmailList(ctx).Execute()



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
    resp, r, err := apiClient.QuarantinedmailAPI.QuarantinedmailList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantinedmailAPI.QuarantinedmailList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuarantinedmailList`: []QuarantinedMailResponse
    fmt.Fprintf(os.Stdout, "Response from `QuarantinedmailAPI.QuarantinedmailList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuarantinedmailListRequest struct via the builder pattern


### Return type

[**[]QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuarantinedmailRead

> QuarantinedMailResponse QuarantinedmailRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.QuarantinedmailAPI.QuarantinedmailRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantinedmailAPI.QuarantinedmailRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuarantinedmailRead`: QuarantinedMailResponse
    fmt.Fprintf(os.Stdout, "Response from `QuarantinedmailAPI.QuarantinedmailRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuarantinedmailReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuarantinedmailSubmit

> []QuarantinedMailResponse QuarantinedmailSubmit(ctx).QuarantinedMailUpdate(quarantinedMailUpdate).Execute()



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
    quarantinedMailUpdate := []openapiclient.QuarantinedMailUpdate{*openapiclient.NewQuarantinedMailUpdate("Id_example")} // []QuarantinedMailUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuarantinedmailAPI.QuarantinedmailSubmit(context.Background()).QuarantinedMailUpdate(quarantinedMailUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuarantinedmailAPI.QuarantinedmailSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuarantinedmailSubmit`: []QuarantinedMailResponse
    fmt.Fprintf(os.Stdout, "Response from `QuarantinedmailAPI.QuarantinedmailSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuarantinedmailSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quarantinedMailUpdate** | [**[]QuarantinedMailUpdate**](QuarantinedMailUpdate.md) |  | 

### Return type

[**[]QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

