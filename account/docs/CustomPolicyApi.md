# \CustomPolicyApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomPolicy**](CustomPolicyApi.md#CreateCustomPolicy) | **Post** /custom_policy/ | 
[**GetCustomPolicies**](CustomPolicyApi.md#GetCustomPolicies) | **Get** /custom_policy/ | 
[**GetCustomPolicy**](CustomPolicyApi.md#GetCustomPolicy) | **Get** /custom_policy/{custom_policy_id} | 
[**UpdateCustomPolicy**](CustomPolicyApi.md#UpdateCustomPolicy) | **Put** /custom_policy/{custom_policy_id} | 



## CreateCustomPolicy

> map[string]interface{} CreateCustomPolicy(ctx).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).CustomPolicyCreateRequest(customPolicyCreateRequest).Execute()





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
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the <a href=\"/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> type definition.<br/> <br/> <span class=\"tablenote\"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span>
    customPolicyCreateRequest := *openapiclient.NewCustomPolicyCreateRequest() // CustomPolicyCreateRequest | Request to create a new Custom Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPolicyApi.CreateCustomPolicy(context.Background()).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).CustomPolicyCreateRequest(customPolicyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPolicyApi.CreateCustomPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomPolicyApi.CreateCustomPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEBAYCMARKETPLACEID** | **string** | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt; | 
 **customPolicyCreateRequest** | [**CustomPolicyCreateRequest**](CustomPolicyCreateRequest.md) | Request to create a new Custom Policy. | 

### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPolicies

> CustomPolicyResponse GetCustomPolicies(ctx).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).PolicyTypes(policyTypes).Execute()





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
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the <a href=\"/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> type definition.<br/> <br/> <span class=\"tablenote\"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span>
    policyTypes := "policyTypes_example" // string | This query parameter specifies the type of custom policies to be returned.<br /><br />Multiple policy types may be requested in a single call by providing a comma-delimited set of all policy types to be returned.<br/><br/><span class=\"tablenote\"><strong>Note:</strong> Omitting this query parameter from a request will also return policies of all policy types.</span><br/><br/>Two Custom Policy types are supported: <ul><li>Product Compliance (PRODUCT_COMPLIANCE)</li> <li>Takeback (TAKE_BACK)</li></ul> (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPolicyApi.GetCustomPolicies(context.Background()).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).PolicyTypes(policyTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPolicyApi.GetCustomPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomPolicies`: CustomPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomPolicyApi.GetCustomPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEBAYCMARKETPLACEID** | **string** | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt; | 
 **policyTypes** | **string** | This query parameter specifies the type of custom policies to be returned.&lt;br /&gt;&lt;br /&gt;Multiple policy types may be requested in a single call by providing a comma-delimited set of all policy types to be returned.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; Omitting this query parameter from a request will also return policies of all policy types.&lt;/span&gt;&lt;br/&gt;&lt;br/&gt;Two Custom Policy types are supported: &lt;ul&gt;&lt;li&gt;Product Compliance (PRODUCT_COMPLIANCE)&lt;/li&gt; &lt;li&gt;Takeback (TAKE_BACK)&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**CustomPolicyResponse**](CustomPolicyResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPolicy

> CustomPolicy GetCustomPolicy(ctx, customPolicyId).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()





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
    customPolicyId := "customPolicyId_example" // string | This path parameter is the unique custom policy identifier for the policy to be returned.<br/><br/><span class=\"tablenote\"><strong>Note:</strong> This value is automatically assigned by the system when the policy is created.</span>
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the <a href=\"/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> type definition.<br/> <br/> <span class=\"tablenote\"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span>

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPolicyApi.GetCustomPolicy(context.Background(), customPolicyId).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPolicyApi.GetCustomPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomPolicy`: CustomPolicy
    fmt.Fprintf(os.Stdout, "Response from `CustomPolicyApi.GetCustomPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customPolicyId** | **string** | This path parameter is the unique custom policy identifier for the policy to be returned.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This value is automatically assigned by the system when the policy is created.&lt;/span&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEBAYCMARKETPLACEID** | **string** | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt; | 

### Return type

[**CustomPolicy**](CustomPolicy.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomPolicy

> UpdateCustomPolicy(ctx, customPolicyId).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).CustomPolicyRequest(customPolicyRequest).Execute()





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
    customPolicyId := "customPolicyId_example" // string | This path parameter is the unique custom policy identifier for the policy to be returned.<br/><br/><span class=\"tablenote\"><strong>Note:</strong> This value is automatically assigned by the system when the policy is created.</span>
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the <a href=\"/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\" target=\"_blank\">MarketplaceIdEnum</a> type definition.<br/> <br/> <span class=\"tablenote\"><strong>Note:</strong> The following eBay marketplaces support Custom Policies: <ul><li>Germany (EBAY_DE)</li> <li>Canada (EBAY_CA)</li> <li>Australia (EBAY_AU)</li> <li>United States (EBAY_US)</li> <li>France (EBAY_FR)</li></ul></span>
    customPolicyRequest := *openapiclient.NewCustomPolicyRequest() // CustomPolicyRequest | Request to update a current custom policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomPolicyApi.UpdateCustomPolicy(context.Background(), customPolicyId).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).CustomPolicyRequest(customPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPolicyApi.UpdateCustomPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customPolicyId** | **string** | This path parameter is the unique custom policy identifier for the policy to be returned.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This value is automatically assigned by the system when the policy is created.&lt;/span&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEBAYCMARKETPLACEID** | **string** | This header parameter specifies the eBay markeplace for the custom policy that is being created. Supported values for this header can be found in the &lt;a href&#x3D;\&quot;/api-docs/sell/compliance/types/bas:MarketplaceIdEnum\&quot; target&#x3D;\&quot;_blank\&quot;&gt;MarketplaceIdEnum&lt;/a&gt; type definition.&lt;br/&gt; &lt;br/&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; The following eBay marketplaces support Custom Policies: &lt;ul&gt;&lt;li&gt;Germany (EBAY_DE)&lt;/li&gt; &lt;li&gt;Canada (EBAY_CA)&lt;/li&gt; &lt;li&gt;Australia (EBAY_AU)&lt;/li&gt; &lt;li&gt;United States (EBAY_US)&lt;/li&gt; &lt;li&gt;France (EBAY_FR)&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt; | 
 **customPolicyRequest** | [**CustomPolicyRequest**](CustomPolicyRequest.md) | Request to update a current custom policy. | 

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

