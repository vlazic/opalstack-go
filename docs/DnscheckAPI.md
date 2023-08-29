# \DnscheckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnscheckLeChallengeCheck**](DnscheckAPI.md#DnscheckLeChallengeCheck) | **Post** /api/v1/dnscheck/le_challenge_check/ | 



## DnscheckLeChallengeCheck

> LEChallengeCheckResponse DnscheckLeChallengeCheck(ctx).LEChallengeCheck(lEChallengeCheck).Execute()



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
    lEChallengeCheck := *openapiclient.NewLEChallengeCheck([]string{"Domains_example"}) // LEChallengeCheck | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnscheckAPI.DnscheckLeChallengeCheck(context.Background()).LEChallengeCheck(lEChallengeCheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnscheckAPI.DnscheckLeChallengeCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnscheckLeChallengeCheck`: LEChallengeCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `DnscheckAPI.DnscheckLeChallengeCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnscheckLeChallengeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lEChallengeCheck** | [**LEChallengeCheck**](LEChallengeCheck.md) |  | 

### Return type

[**LEChallengeCheckResponse**](LEChallengeCheckResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

