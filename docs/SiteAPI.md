# \SiteAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiteCreate**](SiteAPI.md#SiteCreate) | **Post** /api/v1/site/create/ | 
[**SiteDelete**](SiteAPI.md#SiteDelete) | **Post** /api/v1/site/delete/ | 
[**SiteList**](SiteAPI.md#SiteList) | **Get** /api/v1/site/list/ | 
[**SiteRead**](SiteAPI.md#SiteRead) | **Get** /api/v1/site/read/{uuid} | 
[**SiteUpdate**](SiteAPI.md#SiteUpdate) | **Post** /api/v1/site/update/ | 



## SiteCreate

> []SiteResponse SiteCreate(ctx).SiteCreate(siteCreate).Execute()



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
    siteCreate := []openapiclient.SiteCreate{*openapiclient.NewSiteCreate([]string{"Domains_example"}, []openapiclient.Route{*openapiclient.NewRoute("App_example")})} // []SiteCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAPI.SiteCreate(context.Background()).SiteCreate(siteCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.SiteCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteCreate`: []SiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SiteAPI.SiteCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteCreate** | [**[]SiteCreate**](SiteCreate.md) |  | 

### Return type

[**[]SiteResponse**](SiteResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteDelete

> SiteDelete(ctx).SiteRead(siteRead).Execute()



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
    siteRead := []openapiclient.SiteRead{*openapiclient.NewSiteRead("Id_example")} // []SiteRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SiteAPI.SiteDelete(context.Background()).SiteRead(siteRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.SiteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteRead** | [**[]SiteRead**](SiteRead.md) |  | 

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


## SiteList

> []SiteResponse SiteList(ctx).Execute()



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
    resp, r, err := apiClient.SiteAPI.SiteList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.SiteList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteList`: []SiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SiteAPI.SiteList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSiteListRequest struct via the builder pattern


### Return type

[**[]SiteResponse**](SiteResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteRead

> SiteResponse SiteRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.SiteAPI.SiteRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.SiteRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteRead`: SiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SiteAPI.SiteRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteResponse**](SiteResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteUpdate

> []SiteResponse SiteUpdate(ctx).SiteUpdate(siteUpdate).Execute()



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
    siteUpdate := []openapiclient.SiteUpdate{*openapiclient.NewSiteUpdate("Id_example")} // []SiteUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteAPI.SiteUpdate(context.Background()).SiteUpdate(siteUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteAPI.SiteUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteUpdate`: []SiteResponse
    fmt.Fprintf(os.Stdout, "Response from `SiteAPI.SiteUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteUpdate** | [**[]SiteUpdate**](SiteUpdate.md) |  | 

### Return type

[**[]SiteResponse**](SiteResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

