# \ReturnPolicyApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReturnPolicy**](ReturnPolicyApi.md#CreateReturnPolicy) | **Post** /return_policy | 
[**DeleteReturnPolicy**](ReturnPolicyApi.md#DeleteReturnPolicy) | **Delete** /return_policy/{return_policy_id} | 
[**GetReturnPolicies**](ReturnPolicyApi.md#GetReturnPolicies) | **Get** /return_policy | 
[**GetReturnPolicy**](ReturnPolicyApi.md#GetReturnPolicy) | **Get** /return_policy/{return_policy_id} | 
[**GetReturnPolicyByName**](ReturnPolicyApi.md#GetReturnPolicyByName) | **Get** /return_policy/get_by_policy_name | 
[**UpdateReturnPolicy**](ReturnPolicyApi.md#UpdateReturnPolicy) | **Put** /return_policy/{return_policy_id} | 



## CreateReturnPolicy

> SetReturnPolicyResponse CreateReturnPolicy(ctx).ReturnPolicyRequest(returnPolicyRequest).Execute()





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
    returnPolicyRequest := *openapiclient.NewReturnPolicyRequest() // ReturnPolicyRequest | Return policy request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.CreateReturnPolicy(context.Background()).ReturnPolicyRequest(returnPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.CreateReturnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReturnPolicy`: SetReturnPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnPolicyApi.CreateReturnPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnPolicyRequest** | [**ReturnPolicyRequest**](ReturnPolicyRequest.md) | Return policy request | 

### Return type

[**SetReturnPolicyResponse**](SetReturnPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReturnPolicy

> DeleteReturnPolicy(ctx, returnPolicyId).Execute()





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
    returnPolicyId := "returnPolicyId_example" // string | This path parameter specifies the ID of the return policy you want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.DeleteReturnPolicy(context.Background(), returnPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.DeleteReturnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnPolicyId** | **string** | This path parameter specifies the ID of the return policy you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnPolicies

> ReturnPolicyResponse GetReturnPolicies(ctx).MarketplaceId(marketplaceId).Execute()





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
    marketplaceId := "marketplaceId_example" // string | This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.GetReturnPolicies(context.Background()).MarketplaceId(marketplaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.GetReturnPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnPolicies`: ReturnPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnPolicyApi.GetReturnPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 

### Return type

[**ReturnPolicyResponse**](ReturnPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnPolicy

> ReturnPolicy GetReturnPolicy(ctx, returnPolicyId).Execute()





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
    returnPolicyId := "returnPolicyId_example" // string | This path parameter specifies the of the return policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.GetReturnPolicy(context.Background(), returnPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.GetReturnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnPolicy`: ReturnPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReturnPolicyApi.GetReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnPolicyId** | **string** | This path parameter specifies the of the return policy you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReturnPolicy**](ReturnPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnPolicyByName

> ReturnPolicy GetReturnPolicyByName(ctx).MarketplaceId(marketplaceId).Name(name).Execute()





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
    marketplaceId := "marketplaceId_example" // string | This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
    name := "name_example" // string | This query parameter specifies the user-defined name of the return policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.GetReturnPolicyByName(context.Background()).MarketplaceId(marketplaceId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.GetReturnPolicyByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnPolicyByName`: ReturnPolicy
    fmt.Fprintf(os.Stdout, "Response from `ReturnPolicyApi.GetReturnPolicyByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnPolicyByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the ID of the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 
 **name** | **string** | This query parameter specifies the user-defined name of the return policy you want to retrieve. | 

### Return type

[**ReturnPolicy**](ReturnPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReturnPolicy

> SetReturnPolicyResponse UpdateReturnPolicy(ctx, returnPolicyId).ReturnPolicyRequest(returnPolicyRequest).Execute()





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
    returnPolicyId := "returnPolicyId_example" // string | This path parameter specifies the ID of the return policy you want to update.
    returnPolicyRequest := *openapiclient.NewReturnPolicyRequest() // ReturnPolicyRequest | Container for a return policy request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReturnPolicyApi.UpdateReturnPolicy(context.Background(), returnPolicyId).ReturnPolicyRequest(returnPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnPolicyApi.UpdateReturnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReturnPolicy`: SetReturnPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnPolicyApi.UpdateReturnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnPolicyId** | **string** | This path parameter specifies the ID of the return policy you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReturnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnPolicyRequest** | [**ReturnPolicyRequest**](ReturnPolicyRequest.md) | Container for a return policy request. | 

### Return type

[**SetReturnPolicyResponse**](SetReturnPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

