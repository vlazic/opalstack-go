# \NoticeAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NoticeCreate**](NoticeAPI.md#NoticeCreate) | **Post** /api/v1/notice/create/ | 
[**NoticeDelete**](NoticeAPI.md#NoticeDelete) | **Post** /api/v1/notice/delete/ | 
[**NoticeList**](NoticeAPI.md#NoticeList) | **Get** /api/v1/notice/list/ | 
[**NoticeRead**](NoticeAPI.md#NoticeRead) | **Get** /api/v1/notice/read/{uuid} | 
[**NoticeUpdate**](NoticeAPI.md#NoticeUpdate) | **Post** /api/v1/notice/update/ | 



## NoticeCreate

> []NoticeResponse NoticeCreate(ctx).NoticeCreate(noticeCreate).Execute()



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
    noticeCreate := []openapiclient.NoticeCreate{*openapiclient.NewNoticeCreate("Content_example")} // []NoticeCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NoticeAPI.NoticeCreate(context.Background()).NoticeCreate(noticeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoticeAPI.NoticeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoticeCreate`: []NoticeResponse
    fmt.Fprintf(os.Stdout, "Response from `NoticeAPI.NoticeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNoticeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noticeCreate** | [**[]NoticeCreate**](NoticeCreate.md) |  | 

### Return type

[**[]NoticeResponse**](NoticeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoticeDelete

> NoticeDelete(ctx).NoticeRead(noticeRead).Execute()



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
    noticeRead := []openapiclient.NoticeRead{*openapiclient.NewNoticeRead("Id_example")} // []NoticeRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NoticeAPI.NoticeDelete(context.Background()).NoticeRead(noticeRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoticeAPI.NoticeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNoticeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noticeRead** | [**[]NoticeRead**](NoticeRead.md) |  | 

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


## NoticeList

> []NoticeResponse NoticeList(ctx).Execute()



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
    resp, r, err := apiClient.NoticeAPI.NoticeList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoticeAPI.NoticeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoticeList`: []NoticeResponse
    fmt.Fprintf(os.Stdout, "Response from `NoticeAPI.NoticeList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNoticeListRequest struct via the builder pattern


### Return type

[**[]NoticeResponse**](NoticeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoticeRead

> NoticeResponse NoticeRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.NoticeAPI.NoticeRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoticeAPI.NoticeRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoticeRead`: NoticeResponse
    fmt.Fprintf(os.Stdout, "Response from `NoticeAPI.NoticeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoticeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NoticeResponse**](NoticeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoticeUpdate

> []NoticeResponse NoticeUpdate(ctx).NoticeUpdate(noticeUpdate).Execute()



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
    noticeUpdate := []openapiclient.NoticeUpdate{*openapiclient.NewNoticeUpdate("Id_example", "Content_example")} // []NoticeUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NoticeAPI.NoticeUpdate(context.Background()).NoticeUpdate(noticeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NoticeAPI.NoticeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoticeUpdate`: []NoticeResponse
    fmt.Fprintf(os.Stdout, "Response from `NoticeAPI.NoticeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNoticeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noticeUpdate** | [**[]NoticeUpdate**](NoticeUpdate.md) |  | 

### Return type

[**[]NoticeResponse**](NoticeResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

