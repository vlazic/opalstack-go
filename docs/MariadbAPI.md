# \MariadbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCreate**](MariadbAPI.md#MariadbCreate) | **Post** /api/v1/mariadb/create/ | 
[**MariadbDelete**](MariadbAPI.md#MariadbDelete) | **Post** /api/v1/mariadb/delete/ | 
[**MariadbList**](MariadbAPI.md#MariadbList) | **Get** /api/v1/mariadb/list/ | 
[**MariadbRead**](MariadbAPI.md#MariadbRead) | **Get** /api/v1/mariadb/read/{uuid} | 
[**MariadbUpdate**](MariadbAPI.md#MariadbUpdate) | **Post** /api/v1/mariadb/update/ | 



## MariadbCreate

> []MariaDBResponse MariadbCreate(ctx).MariaDBCreate(mariaDBCreate).Execute()



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
    mariaDBCreate := []openapiclient.MariaDBCreate{*openapiclient.NewMariaDBCreate("Name_example", "Server_example")} // []MariaDBCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MariadbAPI.MariadbCreate(context.Background()).MariaDBCreate(mariaDBCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariadbAPI.MariadbCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariadbCreate`: []MariaDBResponse
    fmt.Fprintf(os.Stdout, "Response from `MariadbAPI.MariadbCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaDBCreate** | [**[]MariaDBCreate**](MariaDBCreate.md) |  | 

### Return type

[**[]MariaDBResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariadbDelete

> MariadbDelete(ctx).MariaDBRead(mariaDBRead).Execute()



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
    mariaDBRead := []openapiclient.MariaDBRead{*openapiclient.NewMariaDBRead("Id_example")} // []MariaDBRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MariadbAPI.MariadbDelete(context.Background()).MariaDBRead(mariaDBRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariadbAPI.MariadbDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaDBRead** | [**[]MariaDBRead**](MariaDBRead.md) |  | 

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


## MariadbList

> []MariaDBResponse MariadbList(ctx).Execute()



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
    resp, r, err := apiClient.MariadbAPI.MariadbList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariadbAPI.MariadbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariadbList`: []MariaDBResponse
    fmt.Fprintf(os.Stdout, "Response from `MariadbAPI.MariadbList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListRequest struct via the builder pattern


### Return type

[**[]MariaDBResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariadbRead

> MariaDBResponse MariadbRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.MariadbAPI.MariadbRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariadbAPI.MariadbRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariadbRead`: MariaDBResponse
    fmt.Fprintf(os.Stdout, "Response from `MariadbAPI.MariadbRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MariaDBResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariadbUpdate

> []MariaDBResponse MariadbUpdate(ctx).MariaDBUpdate(mariaDBUpdate).Execute()



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
    mariaDBUpdate := []openapiclient.MariaDBUpdate{*openapiclient.NewMariaDBUpdate("Id_example")} // []MariaDBUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MariadbAPI.MariadbUpdate(context.Background()).MariaDBUpdate(mariaDBUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MariadbAPI.MariadbUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MariadbUpdate`: []MariaDBResponse
    fmt.Fprintf(os.Stdout, "Response from `MariadbAPI.MariadbUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariaDBUpdate** | [**[]MariaDBUpdate**](MariaDBUpdate.md) |  | 

### Return type

[**[]MariaDBResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

