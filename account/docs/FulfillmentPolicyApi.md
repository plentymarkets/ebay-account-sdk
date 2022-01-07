# \FulfillmentPolicyApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFulfillmentPolicy**](FulfillmentPolicyApi.md#CreateFulfillmentPolicy) | **Post** /fulfillment_policy/ | 
[**DeleteFulfillmentPolicy**](FulfillmentPolicyApi.md#DeleteFulfillmentPolicy) | **Delete** /fulfillment_policy/{fulfillmentPolicyId} | 
[**GetFulfillmentPolicies**](FulfillmentPolicyApi.md#GetFulfillmentPolicies) | **Get** /fulfillment_policy | 
[**GetFulfillmentPolicy**](FulfillmentPolicyApi.md#GetFulfillmentPolicy) | **Get** /fulfillment_policy/{fulfillmentPolicyId} | 
[**GetFulfillmentPolicyByName**](FulfillmentPolicyApi.md#GetFulfillmentPolicyByName) | **Get** /fulfillment_policy/get_by_policy_name | 
[**UpdateFulfillmentPolicy**](FulfillmentPolicyApi.md#UpdateFulfillmentPolicy) | **Put** /fulfillment_policy/{fulfillmentPolicyId} | 



## CreateFulfillmentPolicy

> SetFulfillmentPolicyResponse CreateFulfillmentPolicy(ctx).FulfillmentPolicyRequest(fulfillmentPolicyRequest).Execute()





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
    fulfillmentPolicyRequest := *openapiclient.NewFulfillmentPolicyRequest() // FulfillmentPolicyRequest | Request to create a seller account fulfillment policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FulfillmentPolicyApi.CreateFulfillmentPolicy(context.Background()).FulfillmentPolicyRequest(fulfillmentPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.CreateFulfillmentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFulfillmentPolicy`: SetFulfillmentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `FulfillmentPolicyApi.CreateFulfillmentPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFulfillmentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fulfillmentPolicyRequest** | [**FulfillmentPolicyRequest**](FulfillmentPolicyRequest.md) | Request to create a seller account fulfillment policy. | 

### Return type

[**SetFulfillmentPolicyResponse**](SetFulfillmentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFulfillmentPolicy

> DeleteFulfillmentPolicy(ctx, fulfillmentPolicyId).Execute()





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
    fulfillmentPolicyId := "fulfillmentPolicyId_example" // string | This path parameter specifies the ID of the fulfillment policy to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FulfillmentPolicyApi.DeleteFulfillmentPolicy(context.Background(), fulfillmentPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.DeleteFulfillmentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fulfillmentPolicyId** | **string** | This path parameter specifies the ID of the fulfillment policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFulfillmentPolicyRequest struct via the builder pattern


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


## GetFulfillmentPolicies

> FulfillmentPolicyResponse GetFulfillmentPolicies(ctx).MarketplaceId(marketplaceId).Execute()





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
    resp, r, err := api_client.FulfillmentPolicyApi.GetFulfillmentPolicies(context.Background()).MarketplaceId(marketplaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.GetFulfillmentPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFulfillmentPolicies`: FulfillmentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `FulfillmentPolicyApi.GetFulfillmentPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the eBay marketplace of the policies you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 

### Return type

[**FulfillmentPolicyResponse**](FulfillmentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFulfillmentPolicy

> FulfillmentPolicy GetFulfillmentPolicy(ctx, fulfillmentPolicyId).Execute()





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
    fulfillmentPolicyId := "fulfillmentPolicyId_example" // string | This path parameter specifies the ID of the fulfillment policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FulfillmentPolicyApi.GetFulfillmentPolicy(context.Background(), fulfillmentPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.GetFulfillmentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFulfillmentPolicy`: FulfillmentPolicy
    fmt.Fprintf(os.Stdout, "Response from `FulfillmentPolicyApi.GetFulfillmentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fulfillmentPolicyId** | **string** | This path parameter specifies the ID of the fulfillment policy you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FulfillmentPolicy**](FulfillmentPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFulfillmentPolicyByName

> FulfillmentPolicy GetFulfillmentPolicyByName(ctx).MarketplaceId(marketplaceId).Name(name).Execute()





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
    name := "name_example" // string | This query parameter specifies the user-defined name of the fulfillment policy you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FulfillmentPolicyApi.GetFulfillmentPolicyByName(context.Background()).MarketplaceId(marketplaceId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.GetFulfillmentPolicyByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFulfillmentPolicyByName`: FulfillmentPolicy
    fmt.Fprintf(os.Stdout, "Response from `FulfillmentPolicyApi.GetFulfillmentPolicyByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentPolicyByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | This query parameter specifies the eBay marketplace of the policy you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum | 
 **name** | **string** | This query parameter specifies the user-defined name of the fulfillment policy you want to retrieve. | 

### Return type

[**FulfillmentPolicy**](FulfillmentPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFulfillmentPolicy

> SetFulfillmentPolicyResponse UpdateFulfillmentPolicy(ctx, fulfillmentPolicyId).FulfillmentPolicyRequest(fulfillmentPolicyRequest).Execute()





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
    fulfillmentPolicyId := "fulfillmentPolicyId_example" // string | This path parameter specifies the ID of the fulfillment policy you want to update.
    fulfillmentPolicyRequest := *openapiclient.NewFulfillmentPolicyRequest() // FulfillmentPolicyRequest | Fulfillment policy request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FulfillmentPolicyApi.UpdateFulfillmentPolicy(context.Background(), fulfillmentPolicyId).FulfillmentPolicyRequest(fulfillmentPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FulfillmentPolicyApi.UpdateFulfillmentPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFulfillmentPolicy`: SetFulfillmentPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `FulfillmentPolicyApi.UpdateFulfillmentPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fulfillmentPolicyId** | **string** | This path parameter specifies the ID of the fulfillment policy you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFulfillmentPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fulfillmentPolicyRequest** | [**FulfillmentPolicyRequest**](FulfillmentPolicyRequest.md) | Fulfillment policy request | 

### Return type

[**SetFulfillmentPolicyResponse**](SetFulfillmentPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

