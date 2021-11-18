# \CategoryTreeApi

All URIs are relative to *https://api.ebay.com/commerce/taxonomy/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchItemAspects**](CategoryTreeApi.md#FetchItemAspects) | **Get** /category_tree/{category_tree_id}/fetch_item_aspects | Get Aspects for All Leaf Categories in a Marketplace
[**GetCategorySubtree**](CategoryTreeApi.md#GetCategorySubtree) | **Get** /category_tree/{category_tree_id}/get_category_subtree | Get a Category Subtree
[**GetCategorySuggestions**](CategoryTreeApi.md#GetCategorySuggestions) | **Get** /category_tree/{category_tree_id}/get_category_suggestions | Get Suggested Categories
[**GetCategoryTree**](CategoryTreeApi.md#GetCategoryTree) | **Get** /category_tree/{category_tree_id} | Get a Category Tree
[**GetCompatibilityProperties**](CategoryTreeApi.md#GetCompatibilityProperties) | **Get** /category_tree/{category_tree_id}/get_compatibility_properties | Get Compatibility Properties
[**GetCompatibilityPropertyValues**](CategoryTreeApi.md#GetCompatibilityPropertyValues) | **Get** /category_tree/{category_tree_id}/get_compatibility_property_values | Get Compatibility Property Values
[**GetDefaultCategoryTreeId**](CategoryTreeApi.md#GetDefaultCategoryTreeId) | **Get** /get_default_category_tree_id | Get a Default Category Tree ID
[**GetItemAspectsForCategory**](CategoryTreeApi.md#GetItemAspectsForCategory) | **Get** /category_tree/{category_tree_id}/get_item_aspects_for_category | 



## FetchItemAspects

> GetCategoriesAspectResponse FetchItemAspects(ctx, categoryTreeId).Execute()

Get Aspects for All Leaf Categories in a Marketplace



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
    categoryTreeId := "categoryTreeId_example" // string | The unique identifier of the eBay category tree being requested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.FetchItemAspects(context.Background(), categoryTreeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.FetchItemAspects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchItemAspects`: GetCategoriesAspectResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.FetchItemAspects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | The unique identifier of the eBay category tree being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchItemAspectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCategoriesAspectResponse**](GetCategoriesAspectResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategorySubtree

> CategorySubtree GetCategorySubtree(ctx, categoryTreeId).CategoryId(categoryId).Execute()

Get a Category Subtree



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
    categoryId := "categoryId_example" // string | The unique identifier of the category at the top of the subtree being requested. Note: If the category_id submitted identifies the root node of the tree, this call returns an error. To retrieve the complete tree, use this value with the getCategoryTree call. If the category_id submitted identifies a leaf node of the tree, the call response will contain information about only that leaf node, which is a valid subtree.
    categoryTreeId := "categoryTreeId_example" // string | The unique identifier of the eBay category tree from which a category subtree is being requested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetCategorySubtree(context.Background(), categoryTreeId).CategoryId(categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetCategorySubtree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategorySubtree`: CategorySubtree
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetCategorySubtree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | The unique identifier of the eBay category tree from which a category subtree is being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySubtreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | The unique identifier of the category at the top of the subtree being requested. Note: If the category_id submitted identifies the root node of the tree, this call returns an error. To retrieve the complete tree, use this value with the getCategoryTree call. If the category_id submitted identifies a leaf node of the tree, the call response will contain information about only that leaf node, which is a valid subtree. | 


### Return type

[**CategorySubtree**](CategorySubtree.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategorySuggestions

> CategorySuggestionResponse GetCategorySuggestions(ctx, categoryTreeId).Q(q).Execute()

Get Suggested Categories



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
    categoryTreeId := "categoryTreeId_example" // string | The unique identifier of the eBay category tree for which suggested nodes are being requested.
    q := "q_example" // string | A quoted string that describes or characterizes the item being offered for sale. The string format is free form, and can contain any combination of phrases or keywords. eBay will parse the string and return suggested categories for the item.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetCategorySuggestions(context.Background(), categoryTreeId).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetCategorySuggestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategorySuggestions`: CategorySuggestionResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetCategorySuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | The unique identifier of the eBay category tree for which suggested nodes are being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategorySuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | A quoted string that describes or characterizes the item being offered for sale. The string format is free form, and can contain any combination of phrases or keywords. eBay will parse the string and return suggested categories for the item. | 

### Return type

[**CategorySuggestionResponse**](CategorySuggestionResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryTree

> CategoryTree GetCategoryTree(ctx, categoryTreeId).Execute()

Get a Category Tree



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
    categoryTreeId := "categoryTreeId_example" // string | The unique identifier of the eBay category tree being requested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetCategoryTree(context.Background(), categoryTreeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetCategoryTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryTree`: CategoryTree
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetCategoryTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | The unique identifier of the eBay category tree being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CategoryTree**](CategoryTree.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompatibilityProperties

> GetCompatibilityMetadataResponse GetCompatibilityProperties(ctx, categoryTreeId).CategoryId(categoryId).Execute()

Get Compatibility Properties



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
    categoryTreeId := "categoryTreeId_example" // string | This is the unique identifier of category tree. The following is the list of category_tree_id values and the eBay marketplaces that they represent. One of these ID values must be passed in as a path parameter, and the category_id value, that is passed in as query parameter, must be a valid eBay category on that eBay marketplace that supports parts compatibility for cars, trucks, or motorcyles. eBay US: 0 eBay Motors US: 100 eBay Canada: 2 eBay UK: 3 eBay Germany: 77 eBay Australia: 15 eBay France: 71 eBay Italy: 101 eBay Spain: 186
    categoryId := "categoryId_example" // string | The unique identifier of an eBay category. This eBay category must be a valid eBay category on the specified eBay marketplace, and the category must support parts compatibility for cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method of the Selling Metadata API can be used to retrieve all eBay categories for an eBay marketplace that supports parts compatibility cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method can also be used to see if one or more specific eBay categories support parts compatibility.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetCompatibilityProperties(context.Background(), categoryTreeId).CategoryId(categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetCompatibilityProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompatibilityProperties`: GetCompatibilityMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetCompatibilityProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | This is the unique identifier of category tree. The following is the list of category_tree_id values and the eBay marketplaces that they represent. One of these ID values must be passed in as a path parameter, and the category_id value, that is passed in as query parameter, must be a valid eBay category on that eBay marketplace that supports parts compatibility for cars, trucks, or motorcyles. eBay US: 0 eBay Motors US: 100 eBay Canada: 2 eBay UK: 3 eBay Germany: 77 eBay Australia: 15 eBay France: 71 eBay Italy: 101 eBay Spain: 186 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompatibilityPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryId** | **string** | The unique identifier of an eBay category. This eBay category must be a valid eBay category on the specified eBay marketplace, and the category must support parts compatibility for cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method of the Selling Metadata API can be used to retrieve all eBay categories for an eBay marketplace that supports parts compatibility cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method can also be used to see if one or more specific eBay categories support parts compatibility. | 

### Return type

[**GetCompatibilityMetadataResponse**](GetCompatibilityMetadataResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompatibilityPropertyValues

> GetCompatibilityPropertyValuesResponse GetCompatibilityPropertyValues(ctx, categoryTreeId).CompatibilityProperty(compatibilityProperty).CategoryId(categoryId).Filter(filter).Execute()

Get Compatibility Property Values



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
    categoryTreeId := "categoryTreeId_example" // string | This is the unique identifier of the category tree. The following is the list of category_tree_id values and the eBay marketplaces that they represent. One of these ID values must be passed in as a path parameter, and the category_id value, that is passed in as query parameter, must be a valid eBay category on that eBay marketplace that supports parts compatibility for cars, trucks, or motorcyles. eBay US: 0 eBay Motors US: 100 eBay Canada: 2 eBay UK: 3 eBay Germany: 77 eBay Australia: 15 eBay France: 71 eBay Italy: 101 eBay Spain: 186
    compatibilityProperty := "compatibilityProperty_example" // string | One compatible vehicle property applicable to the specified eBay marketplace and eBay category is specified in this required filter. Compatible vehicle properties are returned in the compatibilityProperties.name field of a getCompatibilityProperties response. For example, if you wanted to retrieve all vehicle trims for a 2018 Toyota Camry, you would set this filter as follows: compatibility_property=Trim; and then include the following three name/value filters through one filter parameter: filter=Year:2018,Make:Toyota,Model:Camry. So, putting this all together, your URI would look something like this: GET https://api.ebay.com/commerce/ taxonomy/v1/category_tree/100/ get_compatibility_property_values? category_id=6016&amp;compatibility_property=Trim &amp;filter=filter=Year:2018,Make:Toyota,Model:Camry
    categoryId := "categoryId_example" // string | The unique identifier of an eBay category. This eBay category must be a valid eBay category on the specified eBay marketplace, and the category must support parts compatibility for cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method of the Selling Metadata API can be used to retrieve all eBay categories for an eBay marketplace that supports parts compatibility cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method can also be used to see if one or more specific eBay categories support parts compatibility.
    filter := "filter_example" // string | One or more compatible vehicle property name/value pairs are passed in through this query parameter. The compatible vehicle property name and corresponding value are delimited with a colon (:), such as filter=Year:2018, and multiple compatible vehicle property name/value pairs are delimited with a comma (,). For example, if you wanted to retrieve all vehicle trims for a 2018 Toyota Camry, you would set the compatibility_property filter as follows: compatibility_property=Trim; and then include the following three name/value filters through one filter parameter: filter=Year:2018,Make:Toyota,Model:Camry. So, putting this all together, your URI would look something like this: GET https://api.ebay.com/commerce/ taxonomy/v1/category_tree/100/ get_compatibility_property_values? category_id=6016&amp;compatibility_property=Trim &amp;filter=filter=Year:2018,Make:Toyota,Model:Camry For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:ConstraintFilter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetCompatibilityPropertyValues(context.Background(), categoryTreeId).CompatibilityProperty(compatibilityProperty).CategoryId(categoryId).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetCompatibilityPropertyValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompatibilityPropertyValues`: GetCompatibilityPropertyValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetCompatibilityPropertyValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | This is the unique identifier of the category tree. The following is the list of category_tree_id values and the eBay marketplaces that they represent. One of these ID values must be passed in as a path parameter, and the category_id value, that is passed in as query parameter, must be a valid eBay category on that eBay marketplace that supports parts compatibility for cars, trucks, or motorcyles. eBay US: 0 eBay Motors US: 100 eBay Canada: 2 eBay UK: 3 eBay Germany: 77 eBay Australia: 15 eBay France: 71 eBay Italy: 101 eBay Spain: 186 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompatibilityPropertyValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **compatibilityProperty** | **string** | One compatible vehicle property applicable to the specified eBay marketplace and eBay category is specified in this required filter. Compatible vehicle properties are returned in the compatibilityProperties.name field of a getCompatibilityProperties response. For example, if you wanted to retrieve all vehicle trims for a 2018 Toyota Camry, you would set this filter as follows: compatibility_property&#x3D;Trim; and then include the following three name/value filters through one filter parameter: filter&#x3D;Year:2018,Make:Toyota,Model:Camry. So, putting this all together, your URI would look something like this: GET https://api.ebay.com/commerce/ taxonomy/v1/category_tree/100/ get_compatibility_property_values? category_id&#x3D;6016&amp;amp;compatibility_property&#x3D;Trim &amp;amp;filter&#x3D;filter&#x3D;Year:2018,Make:Toyota,Model:Camry | 
 **categoryId** | **string** | The unique identifier of an eBay category. This eBay category must be a valid eBay category on the specified eBay marketplace, and the category must support parts compatibility for cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method of the Selling Metadata API can be used to retrieve all eBay categories for an eBay marketplace that supports parts compatibility cars, trucks, or motorcyles. The getAutomotivePartsCompatibilityPolicies method can also be used to see if one or more specific eBay categories support parts compatibility. | 
 **filter** | **string** | One or more compatible vehicle property name/value pairs are passed in through this query parameter. The compatible vehicle property name and corresponding value are delimited with a colon (:), such as filter&#x3D;Year:2018, and multiple compatible vehicle property name/value pairs are delimited with a comma (,). For example, if you wanted to retrieve all vehicle trims for a 2018 Toyota Camry, you would set the compatibility_property filter as follows: compatibility_property&#x3D;Trim; and then include the following three name/value filters through one filter parameter: filter&#x3D;Year:2018,Make:Toyota,Model:Camry. So, putting this all together, your URI would look something like this: GET https://api.ebay.com/commerce/ taxonomy/v1/category_tree/100/ get_compatibility_property_values? category_id&#x3D;6016&amp;amp;compatibility_property&#x3D;Trim &amp;amp;filter&#x3D;filter&#x3D;Year:2018,Make:Toyota,Model:Camry For implementation help, refer to eBay API documentation at https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:ConstraintFilter | 

### Return type

[**GetCompatibilityPropertyValuesResponse**](GetCompatibilityPropertyValuesResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultCategoryTreeId

> BaseCategoryTree GetDefaultCategoryTreeId(ctx).MarketplaceId(marketplaceId).AcceptLanguage(acceptLanguage).Execute()

Get a Default Category Tree ID



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
    marketplaceId := "marketplaceId_example" // string | The ID of the eBay marketplace for which the category tree ID is being requested. For a list of supported marketplace IDs, see Marketplaces with Default Category Trees.
    acceptLanguage := "acceptLanguage_example" // string | A header used to indicate the natural language the seller prefers for the response. This specifies the language that the seller wants to use when the field values provided in the request body are displayed to consumers. Note: For details, see Accept-Language in HTTP request headers. Valid Values: For EBAY_CA in French: Accept-Language: fr-CA For EBAY_BE in French: Accept-Language: fr-BE (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetDefaultCategoryTreeId(context.Background()).MarketplaceId(marketplaceId).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetDefaultCategoryTreeId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultCategoryTreeId`: BaseCategoryTree
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetDefaultCategoryTreeId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultCategoryTreeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceId** | **string** | The ID of the eBay marketplace for which the category tree ID is being requested. For a list of supported marketplace IDs, see Marketplaces with Default Category Trees. | 
 **acceptLanguage** | **string** | A header used to indicate the natural language the seller prefers for the response. This specifies the language that the seller wants to use when the field values provided in the request body are displayed to consumers. Note: For details, see Accept-Language in HTTP request headers. Valid Values: For EBAY_CA in French: Accept-Language: fr-CA For EBAY_BE in French: Accept-Language: fr-BE | 

### Return type

[**BaseCategoryTree**](BaseCategoryTree.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemAspectsForCategory

> AspectMetadata GetItemAspectsForCategory(ctx, categoryTreeId).CategoryId(categoryId).Execute()





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
    categoryId := "categoryId_example" // string | The unique identifier of the leaf category for which aspects are being requested. Note: If the category_id submitted does not identify a leaf node of the tree, this call returns an error.
    categoryTreeId := "categoryTreeId_example" // string | The unique identifier of the eBay category tree from which the specified category's aspects are being requested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoryTreeApi.GetItemAspectsForCategory(context.Background(), categoryTreeId).CategoryId(categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryTreeApi.GetItemAspectsForCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemAspectsForCategory`: AspectMetadata
    fmt.Fprintf(os.Stdout, "Response from `CategoryTreeApi.GetItemAspectsForCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryTreeId** | **string** | The unique identifier of the eBay category tree from which the specified category&#39;s aspects are being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAspectsForCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | The unique identifier of the leaf category for which aspects are being requested. Note: If the category_id submitted does not identify a leaf node of the tree, this call returns an error. | 


### Return type

[**AspectMetadata**](AspectMetadata.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

