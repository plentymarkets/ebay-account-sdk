# CustomerServiceMetricsFilterCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerServiceMetricType** | Pointer to **string** | An enumeration value that specifies the customer service metric that eBay tracks to measure seller performance. See CustomerServiceMetricTypeEnum for values. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:CustomerServiceMetricTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**EvaluationMarketplaceId** | Pointer to **string** | An enumeration value that specifies the eBay marketplace where the evaluation occurs. See MarketplaceIdEnum for values. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/bas:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ListingCategories** | Pointer to **[]string** | A list of listing category IDs on which the service metric is measured. A seller can use one or more L1 (top-level) eBay categories to get metrics specific to those L1 categories. The Category IDs for each L1 category are required. Category ID values for L1 categories can be retrieved using the Taxonomy API. Note: Pass this attribute to narrow down your filter results for the ITEM_NOT_AS_DESCRIBED customerServiceMetricType. Supported categories include: primary(L1) category Id | [optional] 
**ShippingRegions** | Pointer to **[]string** | A list of shipping region enumeration values on which the service metric is measured. This comma delimited array allows the seller to customize the report to focus on domestic or international shipping. Note: Pass this attribute to narrow down your filter results for the ITEM_NOT_RECEIVED customerServiceMetricType. Supported categories include: primary(L1) category IdSee ShippingRegionTypeEnum for values | [optional] 

## Methods

### NewCustomerServiceMetricsFilterCriteria

`func NewCustomerServiceMetricsFilterCriteria() *CustomerServiceMetricsFilterCriteria`

NewCustomerServiceMetricsFilterCriteria instantiates a new CustomerServiceMetricsFilterCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceMetricsFilterCriteriaWithDefaults

`func NewCustomerServiceMetricsFilterCriteriaWithDefaults() *CustomerServiceMetricsFilterCriteria`

NewCustomerServiceMetricsFilterCriteriaWithDefaults instantiates a new CustomerServiceMetricsFilterCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerServiceMetricType

`func (o *CustomerServiceMetricsFilterCriteria) GetCustomerServiceMetricType() string`

GetCustomerServiceMetricType returns the CustomerServiceMetricType field if non-nil, zero value otherwise.

### GetCustomerServiceMetricTypeOk

`func (o *CustomerServiceMetricsFilterCriteria) GetCustomerServiceMetricTypeOk() (*string, bool)`

GetCustomerServiceMetricTypeOk returns a tuple with the CustomerServiceMetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServiceMetricType

`func (o *CustomerServiceMetricsFilterCriteria) SetCustomerServiceMetricType(v string)`

SetCustomerServiceMetricType sets CustomerServiceMetricType field to given value.

### HasCustomerServiceMetricType

`func (o *CustomerServiceMetricsFilterCriteria) HasCustomerServiceMetricType() bool`

HasCustomerServiceMetricType returns a boolean if a field has been set.

### GetEvaluationMarketplaceId

`func (o *CustomerServiceMetricsFilterCriteria) GetEvaluationMarketplaceId() string`

GetEvaluationMarketplaceId returns the EvaluationMarketplaceId field if non-nil, zero value otherwise.

### GetEvaluationMarketplaceIdOk

`func (o *CustomerServiceMetricsFilterCriteria) GetEvaluationMarketplaceIdOk() (*string, bool)`

GetEvaluationMarketplaceIdOk returns a tuple with the EvaluationMarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMarketplaceId

`func (o *CustomerServiceMetricsFilterCriteria) SetEvaluationMarketplaceId(v string)`

SetEvaluationMarketplaceId sets EvaluationMarketplaceId field to given value.

### HasEvaluationMarketplaceId

`func (o *CustomerServiceMetricsFilterCriteria) HasEvaluationMarketplaceId() bool`

HasEvaluationMarketplaceId returns a boolean if a field has been set.

### GetListingCategories

`func (o *CustomerServiceMetricsFilterCriteria) GetListingCategories() []string`

GetListingCategories returns the ListingCategories field if non-nil, zero value otherwise.

### GetListingCategoriesOk

`func (o *CustomerServiceMetricsFilterCriteria) GetListingCategoriesOk() (*[]string, bool)`

GetListingCategoriesOk returns a tuple with the ListingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingCategories

`func (o *CustomerServiceMetricsFilterCriteria) SetListingCategories(v []string)`

SetListingCategories sets ListingCategories field to given value.

### HasListingCategories

`func (o *CustomerServiceMetricsFilterCriteria) HasListingCategories() bool`

HasListingCategories returns a boolean if a field has been set.

### GetShippingRegions

`func (o *CustomerServiceMetricsFilterCriteria) GetShippingRegions() []string`

GetShippingRegions returns the ShippingRegions field if non-nil, zero value otherwise.

### GetShippingRegionsOk

`func (o *CustomerServiceMetricsFilterCriteria) GetShippingRegionsOk() (*[]string, bool)`

GetShippingRegionsOk returns a tuple with the ShippingRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingRegions

`func (o *CustomerServiceMetricsFilterCriteria) SetShippingRegions(v []string)`

SetShippingRegions sets ShippingRegions field to given value.

### HasShippingRegions

`func (o *CustomerServiceMetricsFilterCriteria) HasShippingRegions() bool`

HasShippingRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


