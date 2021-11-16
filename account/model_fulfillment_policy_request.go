/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (the Fulfillment Policy, Payment Policy, and Return Policy), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// FulfillmentPolicyRequest <p>This root container defines a seller's fulfillment policy for a specific marketplace and category type. Used when creating or updating a fulfillment policy, <b>fulfillmentPolicyRequest</b> encapsulates a seller's terms for fulfilling an order and includes the shipping carriers and services used for shipment and time the seller takes to ship an order. While each seller must define at least one fulfillment policy for every marketplace into which they sell, sellers can define multiple fulfillment policies for a single marketplace by specifying different configurations for the unique policies.</p> <p>A successful call returns a <b>fulfillmentPolicyId</b>, plus the <b>Location</b> response header contains the URI to the resource.</p>  <p>Policy instructions can be localized by providing a locale in the <code>Content-Language</code> HTTP request header. For example: <code>Content-Language: de-DE</code>.</p>   <p class=\"tablenote\"><b>Tip: </b>For more on using business policies, see <a href=\"/api-docs/sell/static/seller-accounts/business-policies.html\">eBay business policies</a>.</p>
type FulfillmentPolicyRequest struct {
	// The <b>CategoryTypeEnum</b> value to which this policy applies. Used to discern accounts that sell motor vehicles from those that don't. (Currently, each policy can be set to only one <b>categoryTypes</b> value at a time.)
	CategoryTypes *[]CategoryType `json:"categoryTypes,omitempty"`
	// An optional seller-defined description of the fulfillment policy for internal use (this value is not displayed to end users).  <br><br><b>Max length</b>: 250
	Description *string `json:"description,omitempty"`
	// If set to <code>true</code>, the seller offers freight shipping.  Freight shipping can be used for large items over 150 lbs.<br><br><b>Default</b>: false
	FreightShipping *bool `json:"freightShipping,omitempty"`
	// If set to <code>true</code>, the seller has opted-in to the eBay <a href=\"http://pages.ebay.com/help/sell/shipping-globally.html\">Global Shipping Program</a> and that they use that service for their international shipments. Setting this value automatically sets the international shipping service for the policy to <code>International Priority Shipping</code> and the buyer does not need to set any other shipping services for their INTERNATIONAL shipping options (unless they sell items not covered by the Global Shipping Program). <br><br>If this value is set to <code>false</code>, the seller is responsible for manually specifying the international shipping services, as described in <a href=\"https://developer.ebay.com/api-docs/sell/static/seller-accounts/ht_shipping-worldwide.html\">Setting up worldwide shipping</a>. <br><br>To opt-in to the Global Shipping Program, log in to eBay and navigate to <b>My Account > Site Preferences > Shipping preferences</b>.  <p><b>Default</b>: false</p>
	GlobalShipping *bool `json:"globalShipping,omitempty"`
	HandlingTime *TimeDuration `json:"handlingTime,omitempty"`
	// If set to <code>true</code>, no shipping is offered by this policy and the seller offers only local pickup of the item (normally from a non-business location). This option is most often used for customer-to-customer sales and if set, <b>costType</b> should be set to <code>NOT_SPECIFIED</code>.  <br><br><b>Default</b>: <code>false</code>
	LocalPickup *bool `json:"localPickup,omitempty"`
	// The ID of the eBay marketplace to which this fulfillment policy applies. If this value is not specified, value defaults to the seller's eBay registration site. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum'>eBay API documentation</a>
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// A user-defined name for this fulfillment policy. Names must be unique for policies assigned to the same marketplace. <br><br><b>Max length</b>: 64
	Name *string `json:"name,omitempty"`
	// If set to <code>true</code>, the seller offers the \"Click and Collect\" feature. Click and Collect is supported by the Inventory API, and it can be used with Add/Revise/Relist calls.  <br /><br />To enable \"Click and Collect\", a seller (1) must be eligible for Click and Collect and (2) must set this boolean field to 'true'. Currently, Click and Collect is available to only large retail merchants selling in the eBay AU and UK marketplaces.  <br /><br />In addition to setting this field, the merchant must also do the following to enable the \"Click and Collect\" option on a listing: <ul><li>Have inventory for the product at one or more physical stores tied to the merchant's account. <br />Sellers can use the <b>createInventoryLocaion</b> method in the Inventory API to associate physical stores to their account and they can then can add inventory to specific store locations.</li><li>Set an immediate payment requirement on the item. The immediate payment feature requires the seller to: <ul><li>Set the <b>immediatePay</b> flag in the payment policy to 'true'.</li><li>Have a valid store location with a complete street address.</li></ul></li></ul><br /><br /> <span class=\"tablenote\"><b>Note</b>: Payments that are not handled by eBay managed payments have the following additional requirements:<ul><li>Include only one <b>paymentMethods</b> field in the payment policy and set its value to <code>PAYPAL</code>.</li><li>Include a valid PayPal payment address in the <b>recipientAccountReference.referenceId</b> field of the payment policy.</li></ul></span> <br /><br />When a merchant successfully lists an item with Click and Collect, prospective buyers within a reasonable distance from one of the merchant's stores (that has stock available) will see the \"Available for Click and Collect\" option on the listing, along with information on the closest store that has the item.<br /><br /><b>Default</b>: false
	PickupDropOff *bool `json:"pickupDropOff,omitempty"`
	// A list that defines the seller's shipping configurations for DOMESTIC and INTERNATIONAL order shipments. <p><b>shippingOptions</b> is a list with a single element if the seller ships to only domestic locations. If the seller also ships internationally, the list contains a second element that defines their international shipping options.</p> <p>Shipping options configure the high-level shipping settings that apply to orders, such as flat-rate or calculated shipping, any rate tables the seller wants to associate with the shipping services, plus other details (such as the <b>shippingServices</b> offered for domestic or international shipments).</p>
	ShippingOptions *[]ShippingOption `json:"shippingOptions,omitempty"`
	ShipToLocations *RegionSet `json:"shipToLocations,omitempty"`
}

// NewFulfillmentPolicyRequest instantiates a new FulfillmentPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentPolicyRequest() *FulfillmentPolicyRequest {
	this := FulfillmentPolicyRequest{}
	return &this
}

// NewFulfillmentPolicyRequestWithDefaults instantiates a new FulfillmentPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentPolicyRequestWithDefaults() *FulfillmentPolicyRequest {
	this := FulfillmentPolicyRequest{}
	return &this
}

// GetCategoryTypes returns the CategoryTypes field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetCategoryTypes() []CategoryType {
	if o == nil || o.CategoryTypes == nil {
		var ret []CategoryType
		return ret
	}
	return *o.CategoryTypes
}

// GetCategoryTypesOk returns a tuple with the CategoryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetCategoryTypesOk() (*[]CategoryType, bool) {
	if o == nil || o.CategoryTypes == nil {
		return nil, false
	}
	return o.CategoryTypes, true
}

// HasCategoryTypes returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasCategoryTypes() bool {
	if o != nil && o.CategoryTypes != nil {
		return true
	}

	return false
}

// SetCategoryTypes gets a reference to the given []CategoryType and assigns it to the CategoryTypes field.
func (o *FulfillmentPolicyRequest) SetCategoryTypes(v []CategoryType) {
	o.CategoryTypes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FulfillmentPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFreightShipping returns the FreightShipping field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetFreightShipping() bool {
	if o == nil || o.FreightShipping == nil {
		var ret bool
		return ret
	}
	return *o.FreightShipping
}

// GetFreightShippingOk returns a tuple with the FreightShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetFreightShippingOk() (*bool, bool) {
	if o == nil || o.FreightShipping == nil {
		return nil, false
	}
	return o.FreightShipping, true
}

// HasFreightShipping returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasFreightShipping() bool {
	if o != nil && o.FreightShipping != nil {
		return true
	}

	return false
}

// SetFreightShipping gets a reference to the given bool and assigns it to the FreightShipping field.
func (o *FulfillmentPolicyRequest) SetFreightShipping(v bool) {
	o.FreightShipping = &v
}

// GetGlobalShipping returns the GlobalShipping field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetGlobalShipping() bool {
	if o == nil || o.GlobalShipping == nil {
		var ret bool
		return ret
	}
	return *o.GlobalShipping
}

// GetGlobalShippingOk returns a tuple with the GlobalShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetGlobalShippingOk() (*bool, bool) {
	if o == nil || o.GlobalShipping == nil {
		return nil, false
	}
	return o.GlobalShipping, true
}

// HasGlobalShipping returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasGlobalShipping() bool {
	if o != nil && o.GlobalShipping != nil {
		return true
	}

	return false
}

// SetGlobalShipping gets a reference to the given bool and assigns it to the GlobalShipping field.
func (o *FulfillmentPolicyRequest) SetGlobalShipping(v bool) {
	o.GlobalShipping = &v
}

// GetHandlingTime returns the HandlingTime field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetHandlingTime() TimeDuration {
	if o == nil || o.HandlingTime == nil {
		var ret TimeDuration
		return ret
	}
	return *o.HandlingTime
}

// GetHandlingTimeOk returns a tuple with the HandlingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetHandlingTimeOk() (*TimeDuration, bool) {
	if o == nil || o.HandlingTime == nil {
		return nil, false
	}
	return o.HandlingTime, true
}

// HasHandlingTime returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasHandlingTime() bool {
	if o != nil && o.HandlingTime != nil {
		return true
	}

	return false
}

// SetHandlingTime gets a reference to the given TimeDuration and assigns it to the HandlingTime field.
func (o *FulfillmentPolicyRequest) SetHandlingTime(v TimeDuration) {
	o.HandlingTime = &v
}

// GetLocalPickup returns the LocalPickup field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetLocalPickup() bool {
	if o == nil || o.LocalPickup == nil {
		var ret bool
		return ret
	}
	return *o.LocalPickup
}

// GetLocalPickupOk returns a tuple with the LocalPickup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetLocalPickupOk() (*bool, bool) {
	if o == nil || o.LocalPickup == nil {
		return nil, false
	}
	return o.LocalPickup, true
}

// HasLocalPickup returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasLocalPickup() bool {
	if o != nil && o.LocalPickup != nil {
		return true
	}

	return false
}

// SetLocalPickup gets a reference to the given bool and assigns it to the LocalPickup field.
func (o *FulfillmentPolicyRequest) SetLocalPickup(v bool) {
	o.LocalPickup = &v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetMarketplaceId() string {
	if o == nil || o.MarketplaceId == nil {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || o.MarketplaceId == nil {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasMarketplaceId() bool {
	if o != nil && o.MarketplaceId != nil {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *FulfillmentPolicyRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FulfillmentPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetPickupDropOff returns the PickupDropOff field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetPickupDropOff() bool {
	if o == nil || o.PickupDropOff == nil {
		var ret bool
		return ret
	}
	return *o.PickupDropOff
}

// GetPickupDropOffOk returns a tuple with the PickupDropOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetPickupDropOffOk() (*bool, bool) {
	if o == nil || o.PickupDropOff == nil {
		return nil, false
	}
	return o.PickupDropOff, true
}

// HasPickupDropOff returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasPickupDropOff() bool {
	if o != nil && o.PickupDropOff != nil {
		return true
	}

	return false
}

// SetPickupDropOff gets a reference to the given bool and assigns it to the PickupDropOff field.
func (o *FulfillmentPolicyRequest) SetPickupDropOff(v bool) {
	o.PickupDropOff = &v
}

// GetShippingOptions returns the ShippingOptions field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetShippingOptions() []ShippingOption {
	if o == nil || o.ShippingOptions == nil {
		var ret []ShippingOption
		return ret
	}
	return *o.ShippingOptions
}

// GetShippingOptionsOk returns a tuple with the ShippingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetShippingOptionsOk() (*[]ShippingOption, bool) {
	if o == nil || o.ShippingOptions == nil {
		return nil, false
	}
	return o.ShippingOptions, true
}

// HasShippingOptions returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasShippingOptions() bool {
	if o != nil && o.ShippingOptions != nil {
		return true
	}

	return false
}

// SetShippingOptions gets a reference to the given []ShippingOption and assigns it to the ShippingOptions field.
func (o *FulfillmentPolicyRequest) SetShippingOptions(v []ShippingOption) {
	o.ShippingOptions = &v
}

// GetShipToLocations returns the ShipToLocations field value if set, zero value otherwise.
func (o *FulfillmentPolicyRequest) GetShipToLocations() RegionSet {
	if o == nil || o.ShipToLocations == nil {
		var ret RegionSet
		return ret
	}
	return *o.ShipToLocations
}

// GetShipToLocationsOk returns a tuple with the ShipToLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPolicyRequest) GetShipToLocationsOk() (*RegionSet, bool) {
	if o == nil || o.ShipToLocations == nil {
		return nil, false
	}
	return o.ShipToLocations, true
}

// HasShipToLocations returns a boolean if a field has been set.
func (o *FulfillmentPolicyRequest) HasShipToLocations() bool {
	if o != nil && o.ShipToLocations != nil {
		return true
	}

	return false
}

// SetShipToLocations gets a reference to the given RegionSet and assigns it to the ShipToLocations field.
func (o *FulfillmentPolicyRequest) SetShipToLocations(v RegionSet) {
	o.ShipToLocations = &v
}

func (o FulfillmentPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryTypes != nil {
		toSerialize["categoryTypes"] = o.CategoryTypes
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FreightShipping != nil {
		toSerialize["freightShipping"] = o.FreightShipping
	}
	if o.GlobalShipping != nil {
		toSerialize["globalShipping"] = o.GlobalShipping
	}
	if o.HandlingTime != nil {
		toSerialize["handlingTime"] = o.HandlingTime
	}
	if o.LocalPickup != nil {
		toSerialize["localPickup"] = o.LocalPickup
	}
	if o.MarketplaceId != nil {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PickupDropOff != nil {
		toSerialize["pickupDropOff"] = o.PickupDropOff
	}
	if o.ShippingOptions != nil {
		toSerialize["shippingOptions"] = o.ShippingOptions
	}
	if o.ShipToLocations != nil {
		toSerialize["shipToLocations"] = o.ShipToLocations
	}
	return json.Marshal(toSerialize)
}

type NullableFulfillmentPolicyRequest struct {
	value *FulfillmentPolicyRequest
	isSet bool
}

func (v NullableFulfillmentPolicyRequest) Get() *FulfillmentPolicyRequest {
	return v.value
}

func (v *NullableFulfillmentPolicyRequest) Set(val *FulfillmentPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentPolicyRequest(val *FulfillmentPolicyRequest) *NullableFulfillmentPolicyRequest {
	return &NullableFulfillmentPolicyRequest{value: val, isSet: true}
}

func (v NullableFulfillmentPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


