# \DnsrecordAPI

All URIs are relative to *https://my.opalstack.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsrecordCreate**](DnsrecordAPI.md#DnsrecordCreate) | **Post** /api/v1/dnsrecord/create/ | 
[**DnsrecordDelete**](DnsrecordAPI.md#DnsrecordDelete) | **Post** /api/v1/dnsrecord/delete/ | 
[**DnsrecordList**](DnsrecordAPI.md#DnsrecordList) | **Get** /api/v1/dnsrecord/list/ | 
[**DnsrecordRead**](DnsrecordAPI.md#DnsrecordRead) | **Get** /api/v1/dnsrecord/read/{uuid} | 
[**DnsrecordUpdate**](DnsrecordAPI.md#DnsrecordUpdate) | **Post** /api/v1/dnsrecord/update/ | 



## DnsrecordCreate

> []DNSRecordResponse DnsrecordCreate(ctx).DNSRecordCreate(dNSRecordCreate).Execute()



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
    dNSRecordCreate := []openapiclient.DNSRecordCreate{*openapiclient.NewDNSRecordCreate("Domain_example", openapiclient.DNSRecordTypeEnum("A"), "Content_example")} // []DNSRecordCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsrecordAPI.DnsrecordCreate(context.Background()).DNSRecordCreate(dNSRecordCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsrecordAPI.DnsrecordCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsrecordCreate`: []DNSRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsrecordAPI.DnsrecordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsrecordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dNSRecordCreate** | [**[]DNSRecordCreate**](DNSRecordCreate.md) |  | 

### Return type

[**[]DNSRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsrecordDelete

> DnsrecordDelete(ctx).DNSRecordRead(dNSRecordRead).Execute()



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
    dNSRecordRead := []openapiclient.DNSRecordRead{*openapiclient.NewDNSRecordRead("Id_example")} // []DNSRecordRead | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnsrecordAPI.DnsrecordDelete(context.Background()).DNSRecordRead(dNSRecordRead).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsrecordAPI.DnsrecordDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsrecordDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dNSRecordRead** | [**[]DNSRecordRead**](DNSRecordRead.md) |  | 

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


## DnsrecordList

> []DNSRecordResponse DnsrecordList(ctx).Execute()



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
    resp, r, err := apiClient.DnsrecordAPI.DnsrecordList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsrecordAPI.DnsrecordList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsrecordList`: []DNSRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsrecordAPI.DnsrecordList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDnsrecordListRequest struct via the builder pattern


### Return type

[**[]DNSRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsrecordRead

> DNSRecordResponse DnsrecordRead(ctx, uuid).Execute()



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
    resp, r, err := apiClient.DnsrecordAPI.DnsrecordRead(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsrecordAPI.DnsrecordRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsrecordRead`: DNSRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsrecordAPI.DnsrecordRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsrecordReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DNSRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsrecordUpdate

> []DNSRecordResponse DnsrecordUpdate(ctx).DNSRecordUpdate(dNSRecordUpdate).Execute()



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
    dNSRecordUpdate := []openapiclient.DNSRecordUpdate{*openapiclient.NewDNSRecordUpdate("Id_example")} // []DNSRecordUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnsrecordAPI.DnsrecordUpdate(context.Background()).DNSRecordUpdate(dNSRecordUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsrecordAPI.DnsrecordUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsrecordUpdate`: []DNSRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DnsrecordAPI.DnsrecordUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsrecordUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dNSRecordUpdate** | [**[]DNSRecordUpdate**](DNSRecordUpdate.md) |  | 

### Return type

[**[]DNSRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

