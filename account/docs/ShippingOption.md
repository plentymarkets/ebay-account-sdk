# ShippingOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostType** | Pointer to **string** | Defines whether the shipping cost is &lt;code&gt;FLAT_RATE&lt;/code&gt; (the same rate for all buyers), &lt;code&gt;CALCULATED&lt;/code&gt; (the shipping rate varies by the ship-to location and size and weight of the package, as defined by the item), or &lt;code&gt;NOT_SPECIFIED&lt;/code&gt; (for use with local pickup).  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if &lt;/i&gt; the policy offers shipping options using a &lt;b&gt;shippingOptions&lt;/b&gt; container. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ShippingCostTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**InsuranceFee** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**InsuranceOffered** | Pointer to **bool** | This field has been deprecated. &lt;br&gt;&lt;br&gt;Shipping insurance is offered only via a shipping carrier&#39;s shipping services and is no longer available via eBay shipping policies. | [optional] 
**OptionType** | Pointer to **string** | Use this field to set the &lt;b&gt;ShippingOption&lt;/b&gt; element to either &lt;code&gt;DOMESTIC&lt;/code&gt; or &lt;code&gt;INTERNATIONAL&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if &lt;/i&gt; the policy offers shipping options using a &lt;b&gt;shippingOptions&lt;/b&gt; container. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ShippingOptionTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PackageHandlingCost** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**RateTableId** | Pointer to **string** | A unique eBay-assigned ID associated with a user-created shipping rate table. The &lt;b&gt;locality&lt;/b&gt; of a shipping rate table can be either &lt;code&gt;DOMESTIC&lt;/code&gt; or &lt;code&gt;INTERNATIONAL&lt;/code&gt; and you must ensure the value specified in this field references a shipping rate table that matches the type specified in the &lt;b&gt;shippingOptions.optionType&lt;/b&gt; field. If you mismatch the types, eBay responds with a &lt;code&gt;20403&lt;/code&gt; error. &lt;br&gt;&lt;br&gt;Call &lt;a href&#x3D;\&quot;/api-docs/sell/account/resources/rate_table/methods/getRateTables\&quot;&gt;getRateTable&lt;/a&gt; to retrieve information (including &lt;b&gt;rateTableId&lt;/b&gt; values) on the rate tables configured by a seller. For information on creating rate tables, see &lt;a href&#x3D;\&quot;http://pages.ebay.com/help/pay/shipping-costs.html#tables\&quot;&gt;Using shipping rate tables&lt;/a&gt;. | [optional] 
**ShippingServices** | Pointer to [**[]ShippingService**](ShippingService.md) | Contains a list of shipping services offered for either DOMESTIC or INTERNATIONAL shipments. &lt;br&gt;&lt;br&gt;Sellers can specify up to four domestic shipping services and up to five international shipping services by using separate &lt;b&gt;shippingService&lt;/b&gt; containers for each. Note that if the seller is opted in to the Global Shipping Program, they can specify only four other international shipping services, regardless of whether or not Global Shipping is offered as one of the services.  &lt;br&gt;&lt;br&gt; See &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/ht_shipping-setting-shipping-carrier-and-service-values.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;How to set up shipping carrier and shipping service values&lt;/a&gt;. &lt;br /&gt;&lt;br /&gt;To use the eBay standard envelope service (eSE), see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/using-the-ebay-standard-envelope-service.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Using eBay standard envelope (eSE) service&lt;/a&gt;.&lt;br /&gt;&lt;br /&gt;&lt;i&gt;Required if &lt;/i&gt; the policy offers shipping options using a &lt;b&gt;shippingOptions&lt;/b&gt; container. | [optional] 

## Methods

### NewShippingOption

`func NewShippingOption() *ShippingOption`

NewShippingOption instantiates a new ShippingOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingOptionWithDefaults

`func NewShippingOptionWithDefaults() *ShippingOption`

NewShippingOptionWithDefaults instantiates a new ShippingOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostType

`func (o *ShippingOption) GetCostType() string`

GetCostType returns the CostType field if non-nil, zero value otherwise.

### GetCostTypeOk

`func (o *ShippingOption) GetCostTypeOk() (*string, bool)`

GetCostTypeOk returns a tuple with the CostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostType

`func (o *ShippingOption) SetCostType(v string)`

SetCostType sets CostType field to given value.

### HasCostType

`func (o *ShippingOption) HasCostType() bool`

HasCostType returns a boolean if a field has been set.

### GetInsuranceFee

`func (o *ShippingOption) GetInsuranceFee() Amount`

GetInsuranceFee returns the InsuranceFee field if non-nil, zero value otherwise.

### GetInsuranceFeeOk

`func (o *ShippingOption) GetInsuranceFeeOk() (*Amount, bool)`

GetInsuranceFeeOk returns a tuple with the InsuranceFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceFee

`func (o *ShippingOption) SetInsuranceFee(v Amount)`

SetInsuranceFee sets InsuranceFee field to given value.

### HasInsuranceFee

`func (o *ShippingOption) HasInsuranceFee() bool`

HasInsuranceFee returns a boolean if a field has been set.

### GetInsuranceOffered

`func (o *ShippingOption) GetInsuranceOffered() bool`

GetInsuranceOffered returns the InsuranceOffered field if non-nil, zero value otherwise.

### GetInsuranceOfferedOk

`func (o *ShippingOption) GetInsuranceOfferedOk() (*bool, bool)`

GetInsuranceOfferedOk returns a tuple with the InsuranceOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceOffered

`func (o *ShippingOption) SetInsuranceOffered(v bool)`

SetInsuranceOffered sets InsuranceOffered field to given value.

### HasInsuranceOffered

`func (o *ShippingOption) HasInsuranceOffered() bool`

HasInsuranceOffered returns a boolean if a field has been set.

### GetOptionType

`func (o *ShippingOption) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *ShippingOption) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *ShippingOption) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *ShippingOption) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetPackageHandlingCost

`func (o *ShippingOption) GetPackageHandlingCost() Amount`

GetPackageHandlingCost returns the PackageHandlingCost field if non-nil, zero value otherwise.

### GetPackageHandlingCostOk

`func (o *ShippingOption) GetPackageHandlingCostOk() (*Amount, bool)`

GetPackageHandlingCostOk returns a tuple with the PackageHandlingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageHandlingCost

`func (o *ShippingOption) SetPackageHandlingCost(v Amount)`

SetPackageHandlingCost sets PackageHandlingCost field to given value.

### HasPackageHandlingCost

`func (o *ShippingOption) HasPackageHandlingCost() bool`

HasPackageHandlingCost returns a boolean if a field has been set.

### GetRateTableId

`func (o *ShippingOption) GetRateTableId() string`

GetRateTableId returns the RateTableId field if non-nil, zero value otherwise.

### GetRateTableIdOk

`func (o *ShippingOption) GetRateTableIdOk() (*string, bool)`

GetRateTableIdOk returns a tuple with the RateTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTableId

`func (o *ShippingOption) SetRateTableId(v string)`

SetRateTableId sets RateTableId field to given value.

### HasRateTableId

`func (o *ShippingOption) HasRateTableId() bool`

HasRateTableId returns a boolean if a field has been set.

### GetShippingServices

`func (o *ShippingOption) GetShippingServices() []ShippingService`

GetShippingServices returns the ShippingServices field if non-nil, zero value otherwise.

### GetShippingServicesOk

`func (o *ShippingOption) GetShippingServicesOk() (*[]ShippingService, bool)`

GetShippingServicesOk returns a tuple with the ShippingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingServices

`func (o *ShippingOption) SetShippingServices(v []ShippingService)`

SetShippingServices sets ShippingServices field to given value.

### HasShippingServices

`func (o *ShippingOption) HasShippingServices() bool`

HasShippingServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


