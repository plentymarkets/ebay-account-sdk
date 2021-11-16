# \SalesTaxApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceSalesTax**](SalesTaxApi.md#CreateOrReplaceSalesTax) | **Put** /sales_tax/{countryCode}/{jurisdictionId} | 
[**DeleteSalesTax**](SalesTaxApi.md#DeleteSalesTax) | **Delete** /sales_tax/{countryCode}/{jurisdictionId} | 
[**GetSalesTax**](SalesTaxApi.md#GetSalesTax) | **Get** /sales_tax/{countryCode}/{jurisdictionId} | 
[**GetSalesTaxes**](SalesTaxApi.md#GetSalesTaxes) | **Get** /sales_tax | 



## CreateOrReplaceSalesTax

> CreateOrReplaceSalesTax(ctx, countryCode, jurisdictionId).SalesTaxBase(salesTaxBase).Execute()





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
    countryCode := "countryCode_example" // string | This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country for which you want to create tax table entry.
    jurisdictionId := "jurisdictionId_example" // string | This path parameter specifies the ID of the sales-tax jurisdiction for the table entry you want to create.
    salesTaxBase := *openapiclient.NewSalesTaxBase() // SalesTaxBase | A container that describes the how the sales tax is calculated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SalesTaxApi.CreateOrReplaceSalesTax(context.Background(), countryCode, jurisdictionId).SalesTaxBase(salesTaxBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesTaxApi.CreateOrReplaceSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | This path parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code for the country for which you want to create tax table entry. | 
**jurisdictionId** | **string** | This path parameter specifies the ID of the sales-tax jurisdiction for the table entry you want to create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceSalesTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **salesTaxBase** | [**SalesTaxBase**](SalesTaxBase.md) | A container that describes the how the sales tax is calculated. | 

### Return type

 (empty response body)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSalesTax

> DeleteSalesTax(ctx, countryCode, jurisdictionId).Execute()





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
    countryCode := "countryCode_example" // string | This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country whose tax table entry you want to delete.
    jurisdictionId := "jurisdictionId_example" // string | This path parameter specifies the ID of the sales tax jurisdiction whose table entry you want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SalesTaxApi.DeleteSalesTax(context.Background(), countryCode, jurisdictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesTaxApi.DeleteSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | This path parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code for the country whose tax table entry you want to delete. | 
**jurisdictionId** | **string** | This path parameter specifies the ID of the sales tax jurisdiction whose table entry you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesTaxRequest struct via the builder pattern


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


## GetSalesTax

> SalesTax GetSalesTax(ctx, countryCode, jurisdictionId).Execute()





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
    countryCode := "countryCode_example" // string | This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country whose tax table you want to retrieve.
    jurisdictionId := "jurisdictionId_example" // string | This path parameter specifies the ID of the sales tax jurisdiction for the tax table entry you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SalesTaxApi.GetSalesTax(context.Background(), countryCode, jurisdictionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesTaxApi.GetSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSalesTax`: SalesTax
    fmt.Fprintf(os.Stdout, "Response from `SalesTaxApi.GetSalesTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | This path parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code for the country whose tax table you want to retrieve. | 
**jurisdictionId** | **string** | This path parameter specifies the ID of the sales tax jurisdiction for the tax table entry you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SalesTax**](SalesTax.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesTaxes

> SalesTaxes GetSalesTaxes(ctx).CountryCode(countryCode).Execute()





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
    countryCode := "countryCode_example" // string | This path parameter specifies the two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> code for the country whose tax table you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SalesTaxApi.GetSalesTaxes(context.Background()).CountryCode(countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesTaxApi.GetSalesTaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSalesTaxes`: SalesTaxes
    fmt.Fprintf(os.Stdout, "Response from `SalesTaxApi.GetSalesTaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | This path parameter specifies the two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; code for the country whose tax table you want to retrieve. For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum | 

### Return type

[**SalesTaxes**](SalesTaxes.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

