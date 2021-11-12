# \PaymentsProgramApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentsProgram**](PaymentsProgramApi.md#GetPaymentsProgram) | **Get** /payments_program/{marketplace_id}/{payments_program_type} | 



## GetPaymentsProgram

> PaymentsProgramResponse GetPaymentsProgram(ctx, marketplaceId, paymentsProgramType).Execute()





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
    marketplaceId := "marketplaceId_example" // string | This path parameter specifies the eBay marketplace of the payments program for which you want to retrieve the seller's status.
    paymentsProgramType := "paymentsProgramType_example" // string | This path parameter specifies the payments program whose status is returned by the call.  <br><br>Currently the only supported payments program is <code>EBAY_PAYMENTS</code>. For details on the program, see <a href=\"https://pages.ebay.com/payment/2.0/terms.html\" target=\"_blank\">Payments Terms of Use</a>.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsProgramApi.GetPaymentsProgram(context.Background(), marketplaceId, paymentsProgramType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsProgramApi.GetPaymentsProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsProgram`: PaymentsProgramResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentsProgramApi.GetPaymentsProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketplaceId** | **string** | This path parameter specifies the eBay marketplace of the payments program for which you want to retrieve the seller&#39;s status. | 
**paymentsProgramType** | **string** | This path parameter specifies the payments program whose status is returned by the call.  &lt;br&gt;&lt;br&gt;Currently the only supported payments program is &lt;code&gt;EBAY_PAYMENTS&lt;/code&gt;. For details on the program, see &lt;a href&#x3D;\&quot;https://pages.ebay.com/payment/2.0/terms.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Payments Terms of Use&lt;/a&gt;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentsProgramResponse**](PaymentsProgramResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

