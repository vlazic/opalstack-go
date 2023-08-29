# \OsvarAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OsvarCreate**](OsvarAPI.md#OsvarCreate) | **Post** /api/v1/osvar/create/ | 
[**OsvarDelete**](OsvarAPI.md#OsvarDelete) | **Post** /api/v1/osvar/delete/ | 
[**OsvarList**](OsvarAPI.md#OsvarList) | **Get** /api/v1/osvar/list/ | 
[**OsvarRead**](OsvarAPI.md#OsvarRead) | **Get** /api/v1/osvar/read/{uuid} | 
[**OsvarUpdate**](OsvarAPI.md#OsvarUpdate) | **Post** /api/v1/osvar/update/ | 



## OsvarCreate

> []OSVarResponse OsvarCreate(ctx).OSVarCreate(oSVarCreate).Execute()



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
    oSVarCreate := []openapiclient.OSVarCreate{*openapiclient.NewOSVarCreate("Name_example")} // []OSVarCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsvarAPI.OsvarCreate(context.Background()).OSVarCreate(oSVarCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsvarAPI.OsvarCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsvarCreate`: []OSVarResponse
    fmt.Fprintf(os.Stdout, "Response from `OsvarAPI.OsvarCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsvarCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSVarCreate** | [**[]OSVarCreate**](OSVarCreate.md) |  | 

### Return type

[**[]OSVarResponse**](OSVarResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsvarDelete

> OsvarDelete(ctx).OSVarRead(oSVarRead).Execute()



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
    oSVarRead := []openapiclient.OSVarRead{*openapiclient.NewOSVarRead("Id_example")} // []OSVarRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OsvarAPI.OsvarDelete(context.Background()).OSVarRead(oSVarRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsvarAPI.OsvarDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsvarDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSVarRead** | [**[]OSVarRead**](OSVarRead.md) |  | 

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


## OsvarList

> []OSVarResponse OsvarList(ctx).Execute()



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
    resp, r, err := apiClient.OsvarAPI.OsvarList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsvarAPI.OsvarList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsvarList`: []OSVarResponse
    fmt.Fprintf(os.Stdout, "Response from `OsvarAPI.OsvarList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOsvarListRequest struct via the builder pattern


### Return type

[**[]OSVarResponse**](OSVarResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsvarRead

> OSVarResponse OsvarRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.OsvarAPI.OsvarRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsvarAPI.OsvarRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsvarRead`: OSVarResponse
    fmt.Fprintf(os.Stdout, "Response from `OsvarAPI.OsvarRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOsvarReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSVarResponse**](OSVarResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsvarUpdate

> []OSVarResponse OsvarUpdate(ctx).OSVarUpdate(oSVarUpdate).Execute()



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
    oSVarUpdate := []openapiclient.OSVarUpdate{*openapiclient.NewOSVarUpdate("Id_example")} // []OSVarUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsvarAPI.OsvarUpdate(context.Background()).OSVarUpdate(oSVarUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsvarAPI.OsvarUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsvarUpdate`: []OSVarResponse
    fmt.Fprintf(os.Stdout, "Response from `OsvarAPI.OsvarUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsvarUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSVarUpdate** | [**[]OSVarUpdate**](OSVarUpdate.md) |  | 

### Return type

[**[]OSVarResponse**](OSVarResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

