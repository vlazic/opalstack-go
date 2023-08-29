# \TokenAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenCreate**](TokenAPI.md#TokenCreate) | **Post** /api/v1/token/create/ | 
[**TokenDelete**](TokenAPI.md#TokenDelete) | **Post** /api/v1/token/delete/ | 
[**TokenList**](TokenAPI.md#TokenList) | **Get** /api/v1/token/list/ | 
[**TokenRead**](TokenAPI.md#TokenRead) | **Get** /api/v1/token/read/{key} | 
[**TokenUpdate**](TokenAPI.md#TokenUpdate) | **Post** /api/v1/token/update/ | 



## TokenCreate

> []TokenResponse TokenCreate(ctx).TokenCreate(tokenCreate).Execute()



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
    tokenCreate := []openapiclient.TokenCreate{*openapiclient.NewTokenCreate()} // []TokenCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenAPI.TokenCreate(context.Background()).TokenCreate(tokenCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.TokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenCreate`: []TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenAPI.TokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreate** | [**[]TokenCreate**](TokenCreate.md) |  | 

### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenDelete

> TokenDelete(ctx).TokenRead(tokenRead).Execute()



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
    tokenRead := []openapiclient.TokenRead{*openapiclient.NewTokenRead("Key_example")} // []TokenRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TokenAPI.TokenDelete(context.Background()).TokenRead(tokenRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.TokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRead** | [**[]TokenRead**](TokenRead.md) |  | 

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


## TokenList

> []TokenResponse TokenList(ctx).Execute()



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
    resp, r, err := apiClient.TokenAPI.TokenList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.TokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenList`: []TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenAPI.TokenList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenListRequest struct via the builder pattern


### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenRead

> TokenResponse TokenRead(ctx, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenAPI.TokenRead(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.TokenRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenRead`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenAPI.TokenRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokenReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenUpdate

> []TokenResponse TokenUpdate(ctx).TokenUpdate(tokenUpdate).Execute()



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
    tokenUpdate := []openapiclient.TokenUpdate{*openapiclient.NewTokenUpdate("Key_example")} // []TokenUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenAPI.TokenUpdate(context.Background()).TokenUpdate(tokenUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.TokenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenUpdate`: []TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenAPI.TokenUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenUpdate** | [**[]TokenUpdate**](TokenUpdate.md) |  | 

### Return type

[**[]TokenResponse**](TokenResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

