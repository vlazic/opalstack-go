# \LoginAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](LoginAPI.md#Login) | **Post** /api/v1/login/ | 



## Login

> SigninResponse Login(ctx).Signin(signin).Execute()



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
    signin := *openapiclient.NewSignin("Username_example", "Password_example") // Signin | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginAPI.Login(context.Background()).Signin(signin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: SigninResponse
    fmt.Fprintf(os.Stdout, "Response from `LoginAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signin** | [**Signin**](Signin.md) |  | 

### Return type

[**SigninResponse**](SigninResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

