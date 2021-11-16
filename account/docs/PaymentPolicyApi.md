# \PaymentPolicyApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentPolicy**](PaymentPolicyApi.md#CreatePaymentPolicy) | **Post** /payment_policy | 
[**DeletePaymentPolicy**](PaymentPolicyApi.md#DeletePaymentPolicy) | **Delete** /payment_policy/{payment_policy_id} | 
[**GetPaymentPolicies**](PaymentPolicyApi.md#GetPaymentPolicies) | **Get** /payment_policy | 
[**GetPaymentPolicy**](PaymentPolicyApi.md#GetPaymentPolicy) | **Get** /payment_policy/{payment_policy_id} | 
[**GetPaymentPolicyByName**](PaymentPolicyApi.md#GetPaymentPolicyByName) | **Get** /payment_policy/get_by_policy_name | 
[**UpdatePaymentPolicy**](PaymentPolicyApi.md#UpdatePaymentPolicy) | **Put** /payment_policy/{payment_policy_id} | 



## CreatePaymentPolicy

> SetPaymentPolicyResponse CreatePaymentPolicy(ctx).PaymentPolicyRequest(paymentPolicyRequest).Execute()





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
    paymentPolicyRequest := *openapiclient.NewPaymentPolicyRequest() // PaymentPolicyRequest | Payment policy request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.CreatePaymentPolicy(context.Background()).PaymentPolicyRequest(paymentPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.CreatePaymentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentPolicy`: SetPaymentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentPolicyApi.CreatePaymentPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentPolicyRequest** | [**PaymentPolicyRequest**](PaymentPolicyRequest.md) | Payment policy request | 

### Return type

[**SetPaymentPolicyResponse**](SetPaymentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentPolicy

> DeletePaymentPolicy(ctx, paymentPolicyId).Execute()





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
    paymentPolicyId := "paymentPolicyId_example" // string | This path parameter specifies the ID of the payment policy you want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.DeletePaymentPolicy(context.Background(), paymentPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.DeletePaymentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPolicyId** | **string** | This path parameter specifies the ID of the payment policy you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentPolicyRequest struct via the builder pattern


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


## GetPaymentPolicies

> PaymentPolicyResponse GetPaymentPolicies(ctx).MarketplaceId(marketplaceId).Execute()





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
    marketplaceId := "marketplaceId_example" // string | This query parameter specifies the eBay marketplace of the policies you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.GetPaymentPolicies(context.Background()).MarketplaceId(marketplaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.GetPaymentPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentPolicies`: PaymentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentPolicyApi.GetPaymentPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the eBay marketplace of the policies you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 

### Return type

[**PaymentPolicyResponse**](PaymentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentPolicy

> PaymentPolicy GetPaymentPolicy(ctx, paymentPolicyId).Execute()





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
    paymentPolicyId := "paymentPolicyId_example" // string | This path parameter specifies the ID of the payment policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.GetPaymentPolicy(context.Background(), paymentPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.GetPaymentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentPolicy`: PaymentPolicy
    fmt.Fprintf(os.Stdout, "Response from `PaymentPolicyApi.GetPaymentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPolicyId** | **string** | This path parameter specifies the ID of the payment policy you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentPolicy**](PaymentPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentPolicyByName

> PaymentPolicy GetPaymentPolicyByName(ctx).MarketplaceId(marketplaceId).Name(name).Execute()





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
    marketplaceId := "marketplaceId_example" // string | This query parameter specifies the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum
    name := "name_example" // string | This query parameter specifies the user-defined name of the payment policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.GetPaymentPolicyByName(context.Background()).MarketplaceId(marketplaceId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.GetPaymentPolicyByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentPolicyByName`: PaymentPolicy
    fmt.Fprintf(os.Stdout, "Response from `PaymentPolicyApi.GetPaymentPolicyByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentPolicyByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 
 **name** | **string** | This query parameter specifies the user-defined name of the payment policy you want to retrieve. | 

### Return type

[**PaymentPolicy**](PaymentPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentPolicy

> SetPaymentPolicyResponse UpdatePaymentPolicy(ctx, paymentPolicyId).PaymentPolicyRequest(paymentPolicyRequest).Execute()





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
    paymentPolicyId := "paymentPolicyId_example" // string | This path parameter specifies the ID of the payment policy you want to update.
    paymentPolicyRequest := *openapiclient.NewPaymentPolicyRequest() // PaymentPolicyRequest | Payment policy request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentPolicyApi.UpdatePaymentPolicy(context.Background(), paymentPolicyId).PaymentPolicyRequest(paymentPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentPolicyApi.UpdatePaymentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentPolicy`: SetPaymentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentPolicyApi.UpdatePaymentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentPolicyId** | **string** | This path parameter specifies the ID of the payment policy you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentPolicyRequest** | [**PaymentPolicyRequest**](PaymentPolicyRequest.md) | Payment policy request | 

### Return type

[**SetPaymentPolicyResponse**](SetPaymentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

