# SetFulfillmentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | The &lt;b&gt;CategoryTypeEnum&lt;/b&gt; value to which this policy applies. This container is used to discern accounts that sell motor vehicles from those that do not.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction:&lt;/b&gt; Currently, each policy can be set to only one &lt;b&gt;categoryTypes&lt;/b&gt; value at a time. | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the fulfillment policy for internal use (this value is not displayed to end users).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 250 | [optional] 
**FreightShipping** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller offers freight shipping. Freight shipping can be used for large items over 150 lbs. | [optional] 
**FulfillmentPolicyId** | Pointer to **string** | A unique eBay-assigned ID for a fulfillment policy. This ID is generated when the policy is created. | [optional] 
**GlobalShipping** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller has opted-in to the &lt;i&gt;Global Shipping Program&lt;/i&gt; and eBay automatically sets the international shipping service options to &lt;code&gt;International Priority Shipping&lt;/code&gt;. &lt;br&gt;&lt;br&gt;If the value of &lt;b&gt;globalShipping&lt;/b&gt; is &lt;code&gt;false&lt;/code&gt;, the seller is responsible for specifying one or more international shipping service options if they want to ship internationally. | [optional] 
**HandlingTime** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**LocalPickup** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, no shipping is offered by this policy and the seller offers only local pickup of the item (normally from a non-business location). This option is most often used for customer-to-customer sales and if set, &lt;b&gt;costType&lt;/b&gt; should be set to &lt;code&gt;NOT_SPECIFIED&lt;/code&gt;. | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this fulfillment policy applies. If this value is not specified, value defaults to the seller&#39;s eBay registration site. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this fulfillment policy. Names must be unique for policies assigned to the same marketplace. &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 64 | [optional] 
**PickupDropOff** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller offers the \&quot;Click and Collect\&quot; option. &lt;br&gt;&lt;br&gt;Currently, \&quot;Click and Collect\&quot; is available only to large retail merchants the eBay AU and UK marketplaces. | [optional] 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | A list that defines the seller&#39;s shipping configurations for &lt;code&gt;DOMESTIC&lt;/code&gt; and &lt;code&gt;INTERNATIONAL&lt;/code&gt; order shipments. &lt;br&gt;&lt;br&gt;The list has a single element if the seller ships to only domestic locations. If the seller also ships internationally, a second element defines their international shipping options. &lt;br&gt;&lt;br&gt;Shipping options configure the high-level shipping settings that apply to orders, such as flat-rate or calculated shipping, and any rate tables the seller wants to associate with the shipping services. &lt;br&gt;&lt;br&gt;Each &lt;b&gt;shippingOption&lt;/b&gt; element has a &lt;b&gt;shippingServices&lt;/b&gt; container that defines the list of shipping services (domestic or international) offered with this fulfillment policy. | [optional] 
**ShipToLocations** | Pointer to [**RegionSet**](RegionSet.md) |  | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | A list of warnings that were generated during the processing of the request. This field normally returns empty, which indicates the request did not generate any warnings. | [optional] 

## Methods

### NewSetFulfillmentPolicyResponse

`func NewSetFulfillmentPolicyResponse() *SetFulfillmentPolicyResponse`

NewSetFulfillmentPolicyResponse instantiates a new SetFulfillmentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetFulfillmentPolicyResponseWithDefaults

`func NewSetFulfillmentPolicyResponseWithDefaults() *SetFulfillmentPolicyResponse`

NewSetFulfillmentPolicyResponseWithDefaults instantiates a new SetFulfillmentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *SetFulfillmentPolicyResponse) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *SetFulfillmentPolicyResponse) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *SetFulfillmentPolicyResponse) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *SetFulfillmentPolicyResponse) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDescription

`func (o *SetFulfillmentPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetFulfillmentPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetFulfillmentPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetFulfillmentPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFreightShipping

`func (o *SetFulfillmentPolicyResponse) GetFreightShipping() bool`

GetFreightShipping returns the FreightShipping field if non-nil, zero value otherwise.

### GetFreightShippingOk

`func (o *SetFulfillmentPolicyResponse) GetFreightShippingOk() (*bool, bool)`

GetFreightShippingOk returns a tuple with the FreightShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightShipping

`func (o *SetFulfillmentPolicyResponse) SetFreightShipping(v bool)`

SetFreightShipping sets FreightShipping field to given value.

### HasFreightShipping

`func (o *SetFulfillmentPolicyResponse) HasFreightShipping() bool`

HasFreightShipping returns a boolean if a field has been set.

### GetFulfillmentPolicyId

`func (o *SetFulfillmentPolicyResponse) GetFulfillmentPolicyId() string`

GetFulfillmentPolicyId returns the FulfillmentPolicyId field if non-nil, zero value otherwise.

### GetFulfillmentPolicyIdOk

`func (o *SetFulfillmentPolicyResponse) GetFulfillmentPolicyIdOk() (*string, bool)`

GetFulfillmentPolicyIdOk returns a tuple with the FulfillmentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentPolicyId

`func (o *SetFulfillmentPolicyResponse) SetFulfillmentPolicyId(v string)`

SetFulfillmentPolicyId sets FulfillmentPolicyId field to given value.

### HasFulfillmentPolicyId

`func (o *SetFulfillmentPolicyResponse) HasFulfillmentPolicyId() bool`

HasFulfillmentPolicyId returns a boolean if a field has been set.

### GetGlobalShipping

`func (o *SetFulfillmentPolicyResponse) GetGlobalShipping() bool`

GetGlobalShipping returns the GlobalShipping field if non-nil, zero value otherwise.

### GetGlobalShippingOk

`func (o *SetFulfillmentPolicyResponse) GetGlobalShippingOk() (*bool, bool)`

GetGlobalShippingOk returns a tuple with the GlobalShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalShipping

`func (o *SetFulfillmentPolicyResponse) SetGlobalShipping(v bool)`

SetGlobalShipping sets GlobalShipping field to given value.

### HasGlobalShipping

`func (o *SetFulfillmentPolicyResponse) HasGlobalShipping() bool`

HasGlobalShipping returns a boolean if a field has been set.

### GetHandlingTime

`func (o *SetFulfillmentPolicyResponse) GetHandlingTime() TimeDuration`

GetHandlingTime returns the HandlingTime field if non-nil, zero value otherwise.

### GetHandlingTimeOk

`func (o *SetFulfillmentPolicyResponse) GetHandlingTimeOk() (*TimeDuration, bool)`

GetHandlingTimeOk returns a tuple with the HandlingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTime

`func (o *SetFulfillmentPolicyResponse) SetHandlingTime(v TimeDuration)`

SetHandlingTime sets HandlingTime field to given value.

### HasHandlingTime

`func (o *SetFulfillmentPolicyResponse) HasHandlingTime() bool`

HasHandlingTime returns a boolean if a field has been set.

### GetLocalPickup

`func (o *SetFulfillmentPolicyResponse) GetLocalPickup() bool`

GetLocalPickup returns the LocalPickup field if non-nil, zero value otherwise.

### GetLocalPickupOk

`func (o *SetFulfillmentPolicyResponse) GetLocalPickupOk() (*bool, bool)`

GetLocalPickupOk returns a tuple with the LocalPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPickup

`func (o *SetFulfillmentPolicyResponse) SetLocalPickup(v bool)`

SetLocalPickup sets LocalPickup field to given value.

### HasLocalPickup

`func (o *SetFulfillmentPolicyResponse) HasLocalPickup() bool`

HasLocalPickup returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *SetFulfillmentPolicyResponse) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *SetFulfillmentPolicyResponse) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *SetFulfillmentPolicyResponse) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *SetFulfillmentPolicyResponse) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *SetFulfillmentPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetFulfillmentPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetFulfillmentPolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SetFulfillmentPolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPickupDropOff

`func (o *SetFulfillmentPolicyResponse) GetPickupDropOff() bool`

GetPickupDropOff returns the PickupDropOff field if non-nil, zero value otherwise.

### GetPickupDropOffOk

`func (o *SetFulfillmentPolicyResponse) GetPickupDropOffOk() (*bool, bool)`

GetPickupDropOffOk returns a tuple with the PickupDropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDropOff

`func (o *SetFulfillmentPolicyResponse) SetPickupDropOff(v bool)`

SetPickupDropOff sets PickupDropOff field to given value.

### HasPickupDropOff

`func (o *SetFulfillmentPolicyResponse) HasPickupDropOff() bool`

HasPickupDropOff returns a boolean if a field has been set.

### GetShippingOptions

`func (o *SetFulfillmentPolicyResponse) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *SetFulfillmentPolicyResponse) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *SetFulfillmentPolicyResponse) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *SetFulfillmentPolicyResponse) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetShipToLocations

`func (o *SetFulfillmentPolicyResponse) GetShipToLocations() RegionSet`

GetShipToLocations returns the ShipToLocations field if non-nil, zero value otherwise.

### GetShipToLocationsOk

`func (o *SetFulfillmentPolicyResponse) GetShipToLocationsOk() (*RegionSet, bool)`

GetShipToLocationsOk returns a tuple with the ShipToLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToLocations

`func (o *SetFulfillmentPolicyResponse) SetShipToLocations(v RegionSet)`

SetShipToLocations sets ShipToLocations field to given value.

### HasShipToLocations

`func (o *SetFulfillmentPolicyResponse) HasShipToLocations() bool`

HasShipToLocations returns a boolean if a field has been set.

### GetWarnings

`func (o *SetFulfillmentPolicyResponse) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SetFulfillmentPolicyResponse) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SetFulfillmentPolicyResponse) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SetFulfillmentPolicyResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


