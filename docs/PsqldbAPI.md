# \PsqldbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PsqldbCreate**](PsqldbAPI.md#PsqldbCreate) | **Post** /api/v1/psqldb/create/ | 
[**PsqldbDelete**](PsqldbAPI.md#PsqldbDelete) | **Post** /api/v1/psqldb/delete/ | 
[**PsqldbList**](PsqldbAPI.md#PsqldbList) | **Get** /api/v1/psqldb/list/ | 
[**PsqldbRead**](PsqldbAPI.md#PsqldbRead) | **Get** /api/v1/psqldb/read/{uuid} | 
[**PsqldbUpdate**](PsqldbAPI.md#PsqldbUpdate) | **Post** /api/v1/psqldb/update/ | 



## PsqldbCreate

> []PsqlDBResponse PsqldbCreate(ctx).PsqlDBCreate(psqlDBCreate).Execute()



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
    psqlDBCreate := []openapiclient.PsqlDBCreate{*openapiclient.NewPsqlDBCreate("Name_example", "Server_example")} // []PsqlDBCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqldbAPI.PsqldbCreate(context.Background()).PsqlDBCreate(psqlDBCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqldbAPI.PsqldbCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqldbCreate`: []PsqlDBResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqldbAPI.PsqldbCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqldbCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlDBCreate** | [**[]PsqlDBCreate**](PsqlDBCreate.md) |  | 

### Return type

[**[]PsqlDBResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqldbDelete

> PsqldbDelete(ctx).PsqlDBRead(psqlDBRead).Execute()



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
    psqlDBRead := []openapiclient.PsqlDBRead{*openapiclient.NewPsqlDBRead("Id_example")} // []PsqlDBRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PsqldbAPI.PsqldbDelete(context.Background()).PsqlDBRead(psqlDBRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqldbAPI.PsqldbDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqldbDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlDBRead** | [**[]PsqlDBRead**](PsqlDBRead.md) |  | 

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


## PsqldbList

> []PsqlDBResponse PsqldbList(ctx).Execute()



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
    resp, r, err := apiClient.PsqldbAPI.PsqldbList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqldbAPI.PsqldbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqldbList`: []PsqlDBResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqldbAPI.PsqldbList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPsqldbListRequest struct via the builder pattern


### Return type

[**[]PsqlDBResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqldbRead

> PsqlDBResponse PsqldbRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.PsqldbAPI.PsqldbRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqldbAPI.PsqldbRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqldbRead`: PsqlDBResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqldbAPI.PsqldbRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPsqldbReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PsqlDBResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PsqldbUpdate

> []PsqlDBResponse PsqldbUpdate(ctx).PsqlDBUpdate(psqlDBUpdate).Execute()



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
    psqlDBUpdate := []openapiclient.PsqlDBUpdate{*openapiclient.NewPsqlDBUpdate("Id_example")} // []PsqlDBUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PsqldbAPI.PsqldbUpdate(context.Background()).PsqlDBUpdate(psqlDBUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PsqldbAPI.PsqldbUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PsqldbUpdate`: []PsqlDBResponse
    fmt.Fprintf(os.Stdout, "Response from `PsqldbAPI.PsqldbUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPsqldbUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **psqlDBUpdate** | [**[]PsqlDBUpdate**](PsqlDBUpdate.md) |  | 

### Return type

[**[]PsqlDBResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

