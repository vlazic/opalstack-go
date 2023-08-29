# \AppAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCreate**](AppAPI.md#AppCreate) | **Post** /api/v1/app/create/ | 
[**AppDelete**](AppAPI.md#AppDelete) | **Post** /api/v1/app/delete/ | 
[**AppFailed**](AppAPI.md#AppFailed) | **Post** /api/v1/app/failed/ | 
[**AppInstalled**](AppAPI.md#AppInstalled) | **Post** /api/v1/app/installed/ | 
[**AppList**](AppAPI.md#AppList) | **Get** /api/v1/app/list/ | 
[**AppRead**](AppAPI.md#AppRead) | **Get** /api/v1/app/read/{uuid} | 
[**AppUpdate**](AppAPI.md#AppUpdate) | **Post** /api/v1/app/update/ | 



## AppCreate

> []ApplicationResponse AppCreate(ctx).ApplicationCreate(applicationCreate).Execute()





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
    applicationCreate := []openapiclient.ApplicationCreate{*openapiclient.NewApplicationCreate("Name_example", "Osuser_example", openapiclient.AppTypeEnum("STA"))} // []ApplicationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAPI.AppCreate(context.Background()).ApplicationCreate(applicationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCreate`: []ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationCreate** | [**[]ApplicationCreate**](ApplicationCreate.md) |  | 

### Return type

[**[]ApplicationResponse**](ApplicationResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDelete

> AppDelete(ctx).ApplicationRead(applicationRead).Execute()



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
    applicationRead := []openapiclient.ApplicationRead{*openapiclient.NewApplicationRead("Id_example")} // []ApplicationRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppAPI.AppDelete(context.Background()).ApplicationRead(applicationRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationRead** | [**[]ApplicationRead**](ApplicationRead.md) |  | 

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


## AppFailed

> AppFailed(ctx).ApplicationInstalledRequest(applicationInstalledRequest).Execute()



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
    applicationInstalledRequest := []openapiclient.ApplicationInstalledRequest{*openapiclient.NewApplicationInstalledRequest("Id_example")} // []ApplicationInstalledRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppAPI.AppFailed(context.Background()).ApplicationInstalledRequest(applicationInstalledRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppFailed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationInstalledRequest** | [**[]ApplicationInstalledRequest**](ApplicationInstalledRequest.md) |  | 

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


## AppInstalled

> AppInstalled(ctx).ApplicationInstalledRequest(applicationInstalledRequest).Execute()



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
    applicationInstalledRequest := []openapiclient.ApplicationInstalledRequest{*openapiclient.NewApplicationInstalledRequest("Id_example")} // []ApplicationInstalledRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppAPI.AppInstalled(context.Background()).ApplicationInstalledRequest(applicationInstalledRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppInstalled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppInstalledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationInstalledRequest** | [**[]ApplicationInstalledRequest**](ApplicationInstalledRequest.md) |  | 

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


## AppList

> []ApplicationResponse AppList(ctx).Execute()



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
    resp, r, err := apiClient.AppAPI.AppList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppList`: []ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppListRequest struct via the builder pattern


### Return type

[**[]ApplicationResponse**](ApplicationResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRead

> ApplicationResponse AppRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.AppAPI.AppRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppRead`: ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdate

> []ApplicationResponse AppUpdate(ctx).ApplicationUpdate(applicationUpdate).Execute()



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
    applicationUpdate := []openapiclient.ApplicationUpdate{*openapiclient.NewApplicationUpdate("Id_example")} // []ApplicationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAPI.AppUpdate(context.Background()).ApplicationUpdate(applicationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppUpdate`: []ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationUpdate** | [**[]ApplicationUpdate**](ApplicationUpdate.md) |  | 

### Return type

[**[]ApplicationResponse**](ApplicationResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

