# \OnboardingApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentsProgramOnboarding**](OnboardingApi.md#GetPaymentsProgramOnboarding) | **Get** /payments_program/{marketplace_id}/{payments_program_type}/onboarding | 



## GetPaymentsProgramOnboarding

> PaymentsProgramOnboardingResponse GetPaymentsProgramOnboarding(ctx, marketplaceId, paymentsProgramType).Execute()





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
    marketplaceId := "marketplaceId_example" // string | The eBay marketplace ID associated with the onboarding status to retrieve. Only enums for marketplaces that support or will soon support eBay managed payments are allowed. Error 20408 is returned for any other eBay marketplace. No response payload is returned with this error.
    paymentsProgramType := "paymentsProgramType_example" // string | The type of payments program whose status is returned by the call. Presently, the only supported payments program is <code>EBAY_PAYMENTS</code>. For details on the program, see <a href=\"https://pages.ebay.com/payment/2.0/terms.html\" target=\"_blank\">Payments Terms of Use</a>. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OnboardingApi.GetPaymentsProgramOnboarding(context.Background(), marketplaceId, paymentsProgramType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.GetPaymentsProgramOnboarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsProgramOnboarding`: PaymentsProgramOnboardingResponse
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.GetPaymentsProgramOnboarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplaceId** | **string** | The eBay marketplace ID associated with the onboarding status to retrieve. Only enums for marketplaces that support or will soon support eBay managed payments are allowed. Error 20408 is returned for any other eBay marketplace. No response payload is returned with this error. | 
**paymentsProgramType** | **string** | The type of payments program whose status is returned by the call. Presently, the only supported payments program is &lt;code&gt;EBAY_PAYMENTS&lt;/code&gt;. For details on the program, see &lt;a href&#x3D;\&quot;https://pages.ebay.com/payment/2.0/terms.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Payments Terms of Use&lt;/a&gt;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsProgramOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentsProgramOnboardingResponse**](PaymentsProgramOnboardingResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

