# \OsuserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OsuserCreate**](OsuserAPI.md#OsuserCreate) | **Post** /api/v1/osuser/create/ | 
[**OsuserDelete**](OsuserAPI.md#OsuserDelete) | **Post** /api/v1/osuser/delete/ | 
[**OsuserList**](OsuserAPI.md#OsuserList) | **Get** /api/v1/osuser/list/ | 
[**OsuserRead**](OsuserAPI.md#OsuserRead) | **Get** /api/v1/osuser/read/{uuid} | 
[**OsuserUpdate**](OsuserAPI.md#OsuserUpdate) | **Post** /api/v1/osuser/update/ | 



## OsuserCreate

> []OSUserResponse OsuserCreate(ctx).OSUserCreate(oSUserCreate).Execute()



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
    oSUserCreate := []openapiclient.OSUserCreate{*openapiclient.NewOSUserCreate("Server_example")} // []OSUserCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsuserAPI.OsuserCreate(context.Background()).OSUserCreate(oSUserCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsuserAPI.OsuserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsuserCreate`: []OSUserResponse
    fmt.Fprintf(os.Stdout, "Response from `OsuserAPI.OsuserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsuserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSUserCreate** | [**[]OSUserCreate**](OSUserCreate.md) |  | 

### Return type

[**[]OSUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsuserDelete

> OsuserDelete(ctx).OSUserRead(oSUserRead).Execute()



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
    oSUserRead := []openapiclient.OSUserRead{*openapiclient.NewOSUserRead("Id_example")} // []OSUserRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OsuserAPI.OsuserDelete(context.Background()).OSUserRead(oSUserRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsuserAPI.OsuserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsuserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSUserRead** | [**[]OSUserRead**](OSUserRead.md) |  | 

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


## OsuserList

> []OSUserResponse OsuserList(ctx).Execute()



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
    resp, r, err := apiClient.OsuserAPI.OsuserList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsuserAPI.OsuserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsuserList`: []OSUserResponse
    fmt.Fprintf(os.Stdout, "Response from `OsuserAPI.OsuserList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOsuserListRequest struct via the builder pattern


### Return type

[**[]OSUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsuserRead

> OSUserResponse OsuserRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.OsuserAPI.OsuserRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsuserAPI.OsuserRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsuserRead`: OSUserResponse
    fmt.Fprintf(os.Stdout, "Response from `OsuserAPI.OsuserRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOsuserReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OSUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OsuserUpdate

> []OSUserResponse OsuserUpdate(ctx).OSUserUpdate(oSUserUpdate).Execute()



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
    oSUserUpdate := []openapiclient.OSUserUpdate{*openapiclient.NewOSUserUpdate("Id_example")} // []OSUserUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsuserAPI.OsuserUpdate(context.Background()).OSUserUpdate(oSUserUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsuserAPI.OsuserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OsuserUpdate`: []OSUserResponse
    fmt.Fprintf(os.Stdout, "Response from `OsuserAPI.OsuserUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOsuserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oSUserUpdate** | [**[]OSUserUpdate**](OSUserUpdate.md) |  | 

### Return type

[**[]OSUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

