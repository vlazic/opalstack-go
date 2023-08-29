# \MariauserAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariauserCreate**](MariauserAPI.md#MariauserCreate) | **Post** /api/v1/mariauser/create/ | 
[**MariauserDelete**](MariauserAPI.md#MariauserDelete) | **Post** /api/v1/mariauser/delete/ | 
[**MariauserList**](MariauserAPI.md#MariauserList) | **Get** /api/v1/mariauser/list/ | 
[**MariauserRead**](MariauserAPI.md#MariauserRead) | **Get** /api/v1/mariauser/read/{uuid} | 
[**MariauserUpdate**](MariauserAPI.md#MariauserUpdate) | **Post** /api/v1/mariauser/update/ | 



## MariauserCreate

> []MariaUserResponse MariauserCreate(ctx).MariaUserCreate(mariaUserCreate).Execute()



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
    mariaUserCreate := []openapiclient.MariaUserCreate{*openapiclient.NewMariaUserCreate("Server_example", "Name_example")} // []MariaUserCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MariauserAPI.MariauserCreate(context.Background()).MariaUserCreate(mariaUserCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariauserAPI.MariauserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariauserCreate`: []MariaUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MariauserAPI.MariauserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariauserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaUserCreate** | [**[]MariaUserCreate**](MariaUserCreate.md) |  | 

### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariauserDelete

> MariauserDelete(ctx).MariaUserRead(mariaUserRead).Execute()



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
    mariaUserRead := []openapiclient.MariaUserRead{*openapiclient.NewMariaUserRead("Id_example")} // []MariaUserRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MariauserAPI.MariauserDelete(context.Background()).MariaUserRead(mariaUserRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariauserAPI.MariauserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariauserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaUserRead** | [**[]MariaUserRead**](MariaUserRead.md) |  | 

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


## MariauserList

> []MariaUserResponse MariauserList(ctx).Execute()



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
    resp, r, err := apiClient.MariauserAPI.MariauserList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariauserAPI.MariauserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariauserList`: []MariaUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MariauserAPI.MariauserList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMariauserListRequest struct via the builder pattern


### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariauserRead

> MariaUserResponse MariauserRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.MariauserAPI.MariauserRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariauserAPI.MariauserRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariauserRead`: MariaUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MariauserAPI.MariauserRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariauserReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariauserUpdate

> []MariaUserResponse MariauserUpdate(ctx).MariaUserUpdate(mariaUserUpdate).Execute()



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
    mariaUserUpdate := []openapiclient.MariaUserUpdate{*openapiclient.NewMariaUserUpdate("Id_example")} // []MariaUserUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MariauserAPI.MariauserUpdate(context.Background()).MariaUserUpdate(mariaUserUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariauserAPI.MariauserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariauserUpdate`: []MariaUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MariauserAPI.MariauserUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariauserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaUserUpdate** | [**[]MariaUserUpdate**](MariaUserUpdate.md) |  | 

### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

