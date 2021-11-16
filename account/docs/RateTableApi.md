# \RateTableApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRateTables**](RateTableApi.md#GetRateTables) | **Get** /rate_table | 



## GetRateTables

> RateTableResponse GetRateTables(ctx).CountryCode(countryCode).Execute()





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
    countryCode := "countryCode_example" // string | This query parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code of country for which you want shipping-rate table information. If you do not specify a county code, the request returns all the seller-defined rate tables. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateTableApi.GetRateTables(context.Background()).CountryCode(countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateTableApi.GetRateTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRateTables`: RateTableResponse
    fmt.Fprintf(os.Stdout, "Response from `RateTableApi.GetRateTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRateTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | This query parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code of country for which you want shipping-rate table information. If you do not specify a county code, the request returns all the seller-defined rate tables. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum | 

### Return type

[**RateTableResponse**](RateTableResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

