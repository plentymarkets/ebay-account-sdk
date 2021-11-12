# ShippingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalShippingCost** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**BuyerResponsibleForPickup** | Pointer to **bool** | This field is only applicable to vehicle categories on eBay Motors (US and Canada). If set to &lt;code&gt;true&lt;/code&gt;, the buyer is responsible for picking up the vehicle. Otherwise, the seller should specify the vehicle pickup arrangements in the item description. &lt;br&gt;&lt;br&gt;The seller cannot modify this flag if the vehicle has bids or if the listing ends within 12 hours. &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: false | [optional] 
**BuyerResponsibleForShipping** | Pointer to **bool** | This field is applicable for only items listed in vehicle categories on eBay Motors (US and Canada). If set to &lt;code&gt;true&lt;/code&gt;, the buyer is responsible for the shipment of the vehicle. Otherwise, the seller should specify the vehicle shipping arrangements in the item description. &lt;br&gt;&lt;br&gt;The seller cannot modify this flag if the vehicle has bids or if the listing ends within 12 hours. &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: false | [optional] 
**CashOnDeliveryFee** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**FreeShipping** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller offers free shipping to the buyer. This field can only be included and set to &#39;true&#39; for the first domestic shipping service option specified in the &lt;b&gt;shippingServices&lt;/b&gt; container (it is ignored if set for subsequent shipping services). The first specified shipping service option has a &lt;b&gt;sortOrder&lt;/b&gt; value of &lt;code&gt;1&lt;/code&gt; or (if the sortOrderId field is not used) it is the shipping service option that&#39;s specified first in the &lt;b&gt;shippingServices&lt;/b&gt; container. | [optional] 
**ShippingCarrierCode** | Pointer to **string** | The shipping carrier, such as &#39;USPS&#39;, &#39;FedEx&#39;, &#39;UPS&#39;, and so on. | [optional] 
**ShippingCost** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**ShippingServiceCode** | Pointer to **string** | The shipping service that the shipping carrier uses to ship an item. For example, an overnight, two-day delivery, or other type of service. For details on configuring shipping services, see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/ht_shipping-free.html#shippingServices\&quot;&gt;Setting the shipping carrier and shipping service values&lt;/a&gt;. | [optional] 
**ShipToLocations** | Pointer to [**RegionSet**](RegionSet.md) |  | [optional] 
**SortOrder** | Pointer to **int32** | This integer value controls the order that this shipping service option appears in the View Item and Checkout pages, as related to the other specified shipping service options. &lt;br&gt;&lt;br&gt;Sellers can specify up to four domestic shipping services (in four separate &lt;b&gt;shippingService&lt;/b&gt; containers), so valid values are 1, 2, 3, and 4. A shipping service option with a &lt;b&gt;sortOrder&lt;/b&gt; value of &#39;1&#39; appears at the top of View Item and Checkout pages. Conversely, a shipping service option with a &lt;b&gt;sortOrder&lt;/b&gt; value of &#39;4&#39; appears at the bottom of the list. &lt;br&gt;&lt;br&gt;Sellers can specify up to five international shipping services (in five separate &lt;b&gt;shippingService&lt;/b&gt; containers, so valid values for international shipping services are 1, 2, 3, 4, and 5. Similarly to domestic shipping service options, the &lt;b&gt;sortOrder&lt;/b&gt; value of a international shipping service option controls the placement of that shipping service option in the View Item and Checkout pages. Set up different domestic and international services by configuring two &lt;b&gt;shippingOptions&lt;/b&gt; containers, where you set &lt;b&gt;shippingOptions.optionType&lt;/b&gt; to either &lt;code&gt;DOMESTIC&lt;/code&gt; or &lt;code&gt;INTERNATIONAL&lt;/code&gt; to indicate the area supported by the listed shipping services. &lt;br&gt;&lt;br&gt;If the &lt;b&gt;sortOrder&lt;/b&gt; field is not supplied, the order of domestic and international shipping service options is determined by the order in which they are listed in the API call. &lt;br&gt;&lt;br&gt;&lt;b&gt;Min&lt;/b&gt;: 1. &lt;b&gt;Max&lt;/b&gt;: 4 (for domestic shipping service) or 5 (for international shipping service). | [optional] 
**Surcharge** | Pointer to [**Amount**](Amount.md) |  | [optional] 

## Methods

### NewShippingService

`func NewShippingService() *ShippingService`

NewShippingService instantiates a new ShippingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingServiceWithDefaults

`func NewShippingServiceWithDefaults() *ShippingService`

NewShippingServiceWithDefaults instantiates a new ShippingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalShippingCost

`func (o *ShippingService) GetAdditionalShippingCost() Amount`

GetAdditionalShippingCost returns the AdditionalShippingCost field if non-nil, zero value otherwise.

### GetAdditionalShippingCostOk

`func (o *ShippingService) GetAdditionalShippingCostOk() (*Amount, bool)`

GetAdditionalShippingCostOk returns a tuple with the AdditionalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalShippingCost

`func (o *ShippingService) SetAdditionalShippingCost(v Amount)`

SetAdditionalShippingCost sets AdditionalShippingCost field to given value.

### HasAdditionalShippingCost

`func (o *ShippingService) HasAdditionalShippingCost() bool`

HasAdditionalShippingCost returns a boolean if a field has been set.

### GetBuyerResponsibleForPickup

`func (o *ShippingService) GetBuyerResponsibleForPickup() bool`

GetBuyerResponsibleForPickup returns the BuyerResponsibleForPickup field if non-nil, zero value otherwise.

### GetBuyerResponsibleForPickupOk

`func (o *ShippingService) GetBuyerResponsibleForPickupOk() (*bool, bool)`

GetBuyerResponsibleForPickupOk returns a tuple with the BuyerResponsibleForPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerResponsibleForPickup

`func (o *ShippingService) SetBuyerResponsibleForPickup(v bool)`

SetBuyerResponsibleForPickup sets BuyerResponsibleForPickup field to given value.

### HasBuyerResponsibleForPickup

`func (o *ShippingService) HasBuyerResponsibleForPickup() bool`

HasBuyerResponsibleForPickup returns a boolean if a field has been set.

### GetBuyerResponsibleForShipping

`func (o *ShippingService) GetBuyerResponsibleForShipping() bool`

GetBuyerResponsibleForShipping returns the BuyerResponsibleForShipping field if non-nil, zero value otherwise.

### GetBuyerResponsibleForShippingOk

`func (o *ShippingService) GetBuyerResponsibleForShippingOk() (*bool, bool)`

GetBuyerResponsibleForShippingOk returns a tuple with the BuyerResponsibleForShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerResponsibleForShipping

`func (o *ShippingService) SetBuyerResponsibleForShipping(v bool)`

SetBuyerResponsibleForShipping sets BuyerResponsibleForShipping field to given value.

### HasBuyerResponsibleForShipping

`func (o *ShippingService) HasBuyerResponsibleForShipping() bool`

HasBuyerResponsibleForShipping returns a boolean if a field has been set.

### GetCashOnDeliveryFee

`func (o *ShippingService) GetCashOnDeliveryFee() Amount`

GetCashOnDeliveryFee returns the CashOnDeliveryFee field if non-nil, zero value otherwise.

### GetCashOnDeliveryFeeOk

`func (o *ShippingService) GetCashOnDeliveryFeeOk() (*Amount, bool)`

GetCashOnDeliveryFeeOk returns a tuple with the CashOnDeliveryFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDeliveryFee

`func (o *ShippingService) SetCashOnDeliveryFee(v Amount)`

SetCashOnDeliveryFee sets CashOnDeliveryFee field to given value.

### HasCashOnDeliveryFee

`func (o *ShippingService) HasCashOnDeliveryFee() bool`

HasCashOnDeliveryFee returns a boolean if a field has been set.

### GetFreeShipping

`func (o *ShippingService) GetFreeShipping() bool`

GetFreeShipping returns the FreeShipping field if non-nil, zero value otherwise.

### GetFreeShippingOk

`func (o *ShippingService) GetFreeShippingOk() (*bool, bool)`

GetFreeShippingOk returns a tuple with the FreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeShipping

`func (o *ShippingService) SetFreeShipping(v bool)`

SetFreeShipping sets FreeShipping field to given value.

### HasFreeShipping

`func (o *ShippingService) HasFreeShipping() bool`

HasFreeShipping returns a boolean if a field has been set.

### GetShippingCarrierCode

`func (o *ShippingService) GetShippingCarrierCode() string`

GetShippingCarrierCode returns the ShippingCarrierCode field if non-nil, zero value otherwise.

### GetShippingCarrierCodeOk

`func (o *ShippingService) GetShippingCarrierCodeOk() (*string, bool)`

GetShippingCarrierCodeOk returns a tuple with the ShippingCarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierCode

`func (o *ShippingService) SetShippingCarrierCode(v string)`

SetShippingCarrierCode sets ShippingCarrierCode field to given value.

### HasShippingCarrierCode

`func (o *ShippingService) HasShippingCarrierCode() bool`

HasShippingCarrierCode returns a boolean if a field has been set.

### GetShippingCost

`func (o *ShippingService) GetShippingCost() Amount`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ShippingService) GetShippingCostOk() (*Amount, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ShippingService) SetShippingCost(v Amount)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ShippingService) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetShippingServiceCode

`func (o *ShippingService) GetShippingServiceCode() string`

GetShippingServiceCode returns the ShippingServiceCode field if non-nil, zero value otherwise.

### GetShippingServiceCodeOk

`func (o *ShippingService) GetShippingServiceCodeOk() (*string, bool)`

GetShippingServiceCodeOk returns a tuple with the ShippingServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingServiceCode

`func (o *ShippingService) SetShippingServiceCode(v string)`

SetShippingServiceCode sets ShippingServiceCode field to given value.

### HasShippingServiceCode

`func (o *ShippingService) HasShippingServiceCode() bool`

HasShippingServiceCode returns a boolean if a field has been set.

### GetShipToLocations

`func (o *ShippingService) GetShipToLocations() RegionSet`

GetShipToLocations returns the ShipToLocations field if non-nil, zero value otherwise.

### GetShipToLocationsOk

`func (o *ShippingService) GetShipToLocationsOk() (*RegionSet, bool)`

GetShipToLocationsOk returns a tuple with the ShipToLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToLocations

`func (o *ShippingService) SetShipToLocations(v RegionSet)`

SetShipToLocations sets ShipToLocations field to given value.

### HasShipToLocations

`func (o *ShippingService) HasShipToLocations() bool`

HasShipToLocations returns a boolean if a field has been set.

### GetSortOrder

`func (o *ShippingService) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ShippingService) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ShippingService) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ShippingService) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSurcharge

`func (o *ShippingService) GetSurcharge() Amount`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *ShippingService) GetSurchargeOk() (*Amount, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *ShippingService) SetSurcharge(v Amount)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *ShippingService) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


