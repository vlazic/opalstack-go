# \MailuserAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MailuserCreate**](MailuserAPI.md#MailuserCreate) | **Post** /api/v1/mailuser/create/ | 
[**MailuserDelete**](MailuserAPI.md#MailuserDelete) | **Post** /api/v1/mailuser/delete/ | 
[**MailuserList**](MailuserAPI.md#MailuserList) | **Get** /api/v1/mailuser/list/ | 
[**MailuserRead**](MailuserAPI.md#MailuserRead) | **Get** /api/v1/mailuser/read/{uuid} | 
[**MailuserUpdate**](MailuserAPI.md#MailuserUpdate) | **Post** /api/v1/mailuser/update/ | 
[**MailuserUpdatePublic**](MailuserAPI.md#MailuserUpdatePublic) | **Post** /api/v1/mailuser/update_public/ | 



## MailuserCreate

> []MailUserResponse MailuserCreate(ctx).MailUserCreate(mailUserCreate).Execute()



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
    mailUserCreate := []openapiclient.MailUserCreate{*openapiclient.NewMailUserCreate("ImapServer_example", "Name_example")} // []MailUserCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailuserAPI.MailuserCreate(context.Background()).MailUserCreate(mailUserCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MailuserCreate`: []MailUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MailuserAPI.MailuserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMailuserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mailUserCreate** | [**[]MailUserCreate**](MailUserCreate.md) |  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailuserDelete

> MailuserDelete(ctx).MailUserDelete(mailUserDelete).Execute()



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
    mailUserDelete := []openapiclient.MailUserDelete{*openapiclient.NewMailUserDelete("Id_example")} // []MailUserDelete | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailuserAPI.MailuserDelete(context.Background()).MailUserDelete(mailUserDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMailuserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mailUserDelete** | [**[]MailUserDelete**](MailUserDelete.md) |  | 

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


## MailuserList

> []MailUserResponse MailuserList(ctx).Execute()



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
    resp, r, err := apiClient.MailuserAPI.MailuserList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MailuserList`: []MailUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MailuserAPI.MailuserList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMailuserListRequest struct via the builder pattern


### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailuserRead

> MailUserResponse MailuserRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.MailuserAPI.MailuserRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MailuserRead`: MailUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MailuserAPI.MailuserRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMailuserReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailuserUpdate

> []MailUserResponse MailuserUpdate(ctx).MailUserUpdate(mailUserUpdate).Execute()



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
    mailUserUpdate := []openapiclient.MailUserUpdate{*openapiclient.NewMailUserUpdate("Id_example")} // []MailUserUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailuserAPI.MailuserUpdate(context.Background()).MailUserUpdate(mailUserUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MailuserUpdate`: []MailUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MailuserAPI.MailuserUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMailuserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mailUserUpdate** | [**[]MailUserUpdate**](MailUserUpdate.md) |  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MailuserUpdatePublic

> []MailUserResponse MailuserUpdatePublic(ctx).MailUserUpdatePublic(mailUserUpdatePublic).Execute()



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
    mailUserUpdatePublic := []openapiclient.MailUserUpdatePublic{*openapiclient.NewMailUserUpdatePublic("Name_example", "CurrentPassword_example")} // []MailUserUpdatePublic | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailuserAPI.MailuserUpdatePublic(context.Background()).MailUserUpdatePublic(mailUserUpdatePublic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailuserAPI.MailuserUpdatePublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MailuserUpdatePublic`: []MailUserResponse
    fmt.Fprintf(os.Stdout, "Response from `MailuserAPI.MailuserUpdatePublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMailuserUpdatePublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mailUserUpdatePublic** | [**[]MailUserUpdatePublic**](MailUserUpdatePublic.md) |  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

