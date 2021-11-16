# FulfillmentPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | The &lt;b&gt;CategoryTypeEnum&lt;/b&gt; value to which this policy applies. Used to discern accounts that sell motor vehicles from those that don&#39;t. (Currently, each policy can be set to only one &lt;b&gt;categoryTypes&lt;/b&gt; value at a time.) | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the fulfillment policy for internal use (this value is not displayed to end users).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 250 | [optional] 
**FreightShipping** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller offers freight shipping.  Freight shipping can be used for large items over 150 lbs.&lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: false | [optional] 
**GlobalShipping** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller has opted-in to the eBay &lt;a href&#x3D;\&quot;http://pages.ebay.com/help/sell/shipping-globally.html\&quot;&gt;Global Shipping Program&lt;/a&gt; and that they use that service for their international shipments. Setting this value automatically sets the international shipping service for the policy to &lt;code&gt;International Priority Shipping&lt;/code&gt; and the buyer does not need to set any other shipping services for their INTERNATIONAL shipping options (unless they sell items not covered by the Global Shipping Program). &lt;br&gt;&lt;br&gt;If this value is set to &lt;code&gt;false&lt;/code&gt;, the seller is responsible for manually specifying the international shipping services, as described in &lt;a href&#x3D;\&quot;https://developer.ebay.com/api-docs/sell/static/seller-accounts/ht_shipping-worldwide.html\&quot;&gt;Setting up worldwide shipping&lt;/a&gt;. &lt;br&gt;&lt;br&gt;To opt-in to the Global Shipping Program, log in to eBay and navigate to &lt;b&gt;My Account &gt; Site Preferences &gt; Shipping preferences&lt;/b&gt;.  &lt;p&gt;&lt;b&gt;Default&lt;/b&gt;: false&lt;/p&gt; | [optional] 
**HandlingTime** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**LocalPickup** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, no shipping is offered by this policy and the seller offers only local pickup of the item (normally from a non-business location). This option is most often used for customer-to-customer sales and if set, &lt;b&gt;costType&lt;/b&gt; should be set to &lt;code&gt;NOT_SPECIFIED&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: &lt;code&gt;false&lt;/code&gt; | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this fulfillment policy applies. If this value is not specified, value defaults to the seller&#39;s eBay registration site. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this fulfillment policy. Names must be unique for policies assigned to the same marketplace. &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 64 | [optional] 
**PickupDropOff** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller offers the \&quot;Click and Collect\&quot; feature. Click and Collect is supported by the Inventory API, and it can be used with Add/Revise/Relist calls.  &lt;br /&gt;&lt;br /&gt;To enable \&quot;Click and Collect\&quot;, a seller (1) must be eligible for Click and Collect and (2) must set this boolean field to &#39;true&#39;. Currently, Click and Collect is available to only large retail merchants selling in the eBay AU and UK marketplaces.  &lt;br /&gt;&lt;br /&gt;In addition to setting this field, the merchant must also do the following to enable the \&quot;Click and Collect\&quot; option on a listing: &lt;ul&gt;&lt;li&gt;Have inventory for the product at one or more physical stores tied to the merchant&#39;s account. &lt;br /&gt;Sellers can use the &lt;b&gt;createInventoryLocaion&lt;/b&gt; method in the Inventory API to associate physical stores to their account and they can then can add inventory to specific store locations.&lt;/li&gt;&lt;li&gt;Set an immediate payment requirement on the item. The immediate payment feature requires the seller to: &lt;ul&gt;&lt;li&gt;Set the &lt;b&gt;immediatePay&lt;/b&gt; flag in the payment policy to &#39;true&#39;.&lt;/li&gt;&lt;li&gt;Have a valid store location with a complete street address.&lt;/li&gt;&lt;/ul&gt;&lt;/li&gt;&lt;/ul&gt;&lt;br /&gt;&lt;br /&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note&lt;/b&gt;: Payments that are not handled by eBay managed payments have the following additional requirements:&lt;ul&gt;&lt;li&gt;Include only one &lt;b&gt;paymentMethods&lt;/b&gt; field in the payment policy and set its value to &lt;code&gt;PAYPAL&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;Include a valid PayPal payment address in the &lt;b&gt;recipientAccountReference.referenceId&lt;/b&gt; field of the payment policy.&lt;/li&gt;&lt;/ul&gt;&lt;/span&gt; &lt;br /&gt;&lt;br /&gt;When a merchant successfully lists an item with Click and Collect, prospective buyers within a reasonable distance from one of the merchant&#39;s stores (that has stock available) will see the \&quot;Available for Click and Collect\&quot; option on the listing, along with information on the closest store that has the item.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Default&lt;/b&gt;: false | [optional] 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | A list that defines the seller&#39;s shipping configurations for DOMESTIC and INTERNATIONAL order shipments. &lt;p&gt;&lt;b&gt;shippingOptions&lt;/b&gt; is a list with a single element if the seller ships to only domestic locations. If the seller also ships internationally, the list contains a second element that defines their international shipping options.&lt;/p&gt; &lt;p&gt;Shipping options configure the high-level shipping settings that apply to orders, such as flat-rate or calculated shipping, any rate tables the seller wants to associate with the shipping services, plus other details (such as the &lt;b&gt;shippingServices&lt;/b&gt; offered for domestic or international shipments).&lt;/p&gt; | [optional] 
**ShipToLocations** | Pointer to [**RegionSet**](RegionSet.md) |  | [optional] 

## Methods

### NewFulfillmentPolicyRequest

`func NewFulfillmentPolicyRequest() *FulfillmentPolicyRequest`

NewFulfillmentPolicyRequest instantiates a new FulfillmentPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentPolicyRequestWithDefaults

`func NewFulfillmentPolicyRequestWithDefaults() *FulfillmentPolicyRequest`

NewFulfillmentPolicyRequestWithDefaults instantiates a new FulfillmentPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *FulfillmentPolicyRequest) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *FulfillmentPolicyRequest) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *FulfillmentPolicyRequest) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *FulfillmentPolicyRequest) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDescription

`func (o *FulfillmentPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FulfillmentPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FulfillmentPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FulfillmentPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFreightShipping

`func (o *FulfillmentPolicyRequest) GetFreightShipping() bool`

GetFreightShipping returns the FreightShipping field if non-nil, zero value otherwise.

### GetFreightShippingOk

`func (o *FulfillmentPolicyRequest) GetFreightShippingOk() (*bool, bool)`

GetFreightShippingOk returns a tuple with the FreightShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightShipping

`func (o *FulfillmentPolicyRequest) SetFreightShipping(v bool)`

SetFreightShipping sets FreightShipping field to given value.

### HasFreightShipping

`func (o *FulfillmentPolicyRequest) HasFreightShipping() bool`

HasFreightShipping returns a boolean if a field has been set.

### GetGlobalShipping

`func (o *FulfillmentPolicyRequest) GetGlobalShipping() bool`

GetGlobalShipping returns the GlobalShipping field if non-nil, zero value otherwise.

### GetGlobalShippingOk

`func (o *FulfillmentPolicyRequest) GetGlobalShippingOk() (*bool, bool)`

GetGlobalShippingOk returns a tuple with the GlobalShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalShipping

`func (o *FulfillmentPolicyRequest) SetGlobalShipping(v bool)`

SetGlobalShipping sets GlobalShipping field to given value.

### HasGlobalShipping

`func (o *FulfillmentPolicyRequest) HasGlobalShipping() bool`

HasGlobalShipping returns a boolean if a field has been set.

### GetHandlingTime

`func (o *FulfillmentPolicyRequest) GetHandlingTime() TimeDuration`

GetHandlingTime returns the HandlingTime field if non-nil, zero value otherwise.

### GetHandlingTimeOk

`func (o *FulfillmentPolicyRequest) GetHandlingTimeOk() (*TimeDuration, bool)`

GetHandlingTimeOk returns a tuple with the HandlingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTime

`func (o *FulfillmentPolicyRequest) SetHandlingTime(v TimeDuration)`

SetHandlingTime sets HandlingTime field to given value.

### HasHandlingTime

`func (o *FulfillmentPolicyRequest) HasHandlingTime() bool`

HasHandlingTime returns a boolean if a field has been set.

### GetLocalPickup

`func (o *FulfillmentPolicyRequest) GetLocalPickup() bool`

GetLocalPickup returns the LocalPickup field if non-nil, zero value otherwise.

### GetLocalPickupOk

`func (o *FulfillmentPolicyRequest) GetLocalPickupOk() (*bool, bool)`

GetLocalPickupOk returns a tuple with the LocalPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPickup

`func (o *FulfillmentPolicyRequest) SetLocalPickup(v bool)`

SetLocalPickup sets LocalPickup field to given value.

### HasLocalPickup

`func (o *FulfillmentPolicyRequest) HasLocalPickup() bool`

HasLocalPickup returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *FulfillmentPolicyRequest) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *FulfillmentPolicyRequest) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *FulfillmentPolicyRequest) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *FulfillmentPolicyRequest) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *FulfillmentPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FulfillmentPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FulfillmentPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FulfillmentPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPickupDropOff

`func (o *FulfillmentPolicyRequest) GetPickupDropOff() bool`

GetPickupDropOff returns the PickupDropOff field if non-nil, zero value otherwise.

### GetPickupDropOffOk

`func (o *FulfillmentPolicyRequest) GetPickupDropOffOk() (*bool, bool)`

GetPickupDropOffOk returns a tuple with the PickupDropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDropOff

`func (o *FulfillmentPolicyRequest) SetPickupDropOff(v bool)`

SetPickupDropOff sets PickupDropOff field to given value.

### HasPickupDropOff

`func (o *FulfillmentPolicyRequest) HasPickupDropOff() bool`

HasPickupDropOff returns a boolean if a field has been set.

### GetShippingOptions

`func (o *FulfillmentPolicyRequest) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *FulfillmentPolicyRequest) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *FulfillmentPolicyRequest) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *FulfillmentPolicyRequest) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetShipToLocations

`func (o *FulfillmentPolicyRequest) GetShipToLocations() RegionSet`

GetShipToLocations returns the ShipToLocations field if non-nil, zero value otherwise.

### GetShipToLocationsOk

`func (o *FulfillmentPolicyRequest) GetShipToLocationsOk() (*RegionSet, bool)`

GetShipToLocationsOk returns a tuple with the ShipToLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToLocations

`func (o *FulfillmentPolicyRequest) SetShipToLocations(v RegionSet)`

SetShipToLocations sets ShipToLocations field to given value.

### HasShipToLocations

`func (o *FulfillmentPolicyRequest) HasShipToLocations() bool`

HasShipToLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


