# \PrivilegeApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivileges**](PrivilegeApi.md#GetPrivileges) | **Get** /privilege | 



## GetPrivileges

> SellingPrivileges GetPrivileges(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivilegeApi.GetPrivileges(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeApi.GetPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivileges`: SellingPrivileges
    fmt.Fprintf(os.Stdout, "Response from `PrivilegeApi.GetPrivileges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


### Return type

[**SellingPrivileges**](SellingPrivileges.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

