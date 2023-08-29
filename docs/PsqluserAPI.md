# \PsqluserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PsqluserCreate**](PsqluserAPI.md#PsqluserCreate) | **Post** /api/v1/psqluser/create/ | 
[**PsqluserDelete**](PsqluserAPI.md#PsqluserDelete) | **Post** /api/v1/psqluser/delete/ | 
[**PsqluserList**](PsqluserAPI.md#PsqluserList) | **Get** /api/v1/psqluser/list/ | 
[**PsqluserRead**](PsqluserAPI.md#PsqluserRead) | **Get** /api/v1/psqluser/read/{uuid} | 
[**PsqluserUpdate**](PsqluserAPI.md#PsqluserUpdate) | **Post** /api/v1/psqluser/update/ | 



## PsqluserCreate

> []PsqlUserResponse PsqluserCreate(ctx).PsqlUserCreate(psqlUserCreate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    psqlUserCreate := []openapiclient.PsqlUserCreate{*openapiclient.NewPsqlUserCreate("Server_example", "Name_example")} // []PsqlUserCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqluserAPI.PsqluserCreate(context.Background()).PsqlUserCreate(psqlUserCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqluserAPI.PsqluserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqluserCreate`: []PsqlUserResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqluserAPI.PsqluserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqluserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlUserCreate** | [**[]PsqlUserCreate**](PsqlUserCreate.md) |  | 

### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqluserDelete

> PsqluserDelete(ctx).PsqlUserRead(psqlUserRead).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    psqlUserRead := []openapiclient.PsqlUserRead{*openapiclient.NewPsqlUserRead("Id_example")} // []PsqlUserRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PsqluserAPI.PsqluserDelete(context.Background()).PsqlUserRead(psqlUserRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqluserAPI.PsqluserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqluserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlUserRead** | [**[]PsqlUserRead**](PsqlUserRead.md) |  | 

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


## PsqluserList

> []PsqlUserResponse PsqluserList(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqluserAPI.PsqluserList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqluserAPI.PsqluserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqluserList`: []PsqlUserResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqluserAPI.PsqluserList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPsqluserListRequest struct via the builder pattern


### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqluserRead

> PsqlUserResponse PsqluserRead(ctx, uuid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqluserAPI.PsqluserRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqluserAPI.PsqluserRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqluserRead`: PsqlUserResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqluserAPI.PsqluserRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPsqluserReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqluserUpdate

> []PsqlUserResponse PsqluserUpdate(ctx).PsqlUserUpdate(psqlUserUpdate).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    psqlUserUpdate := []openapiclient.PsqlUserUpdate{*openapiclient.NewPsqlUserUpdate("Id_example")} // []PsqlUserUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqluserAPI.PsqluserUpdate(context.Background()).PsqlUserUpdate(psqlUserUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqluserAPI.PsqluserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqluserUpdate`: []PsqlUserResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqluserAPI.PsqluserUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqluserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlUserUpdate** | [**[]PsqlUserUpdate**](PsqlUserUpdate.md) |  | 

### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

