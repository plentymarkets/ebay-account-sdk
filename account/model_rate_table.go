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

// RateTable A complex type that contains information pertaining to a shipping rate table.
type RateTable struct {
	// A two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" title=\"https://www.iso.org\" target=\"_blank\">ISO 3166</a> country code representing the eBay marketplace where an item is listed. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum'>eBay API documentation</a>
	CountryCode *string `json:"countryCode,omitempty"`
	// The region covered by the shipping rate table, either DOMESTIC or INTERNATIONAL. <br><br>DOMESTIC indicates that the shipping rate table applies to regions within the country where an item is listed (the <i>source</i> country) while INTERNATIONAL indicates that the shipping rate table applies to regions outside of the country where an item is listed. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/sell/account/types/api:ShippingOptionTypeEnum'>eBay API documentation</a>
	Locality *string `json:"locality,omitempty"`
	// The user-defined name for the shipping rate table. Sellers can access Seller Hub (or <b>My eBay > Account > Site Preferences > Shipping preferences</b>) to create and assign names to their shipping rate tables.
	Name *string `json:"name,omitempty"`
	// A unique eBay-assigned ID for a seller's shipping rate table. Call <b>getRateTables</b> to retrieve the seller's current rate table IDs.
	RateTableId *string `json:"rateTableId,omitempty"`
}

// NewRateTable instantiates a new RateTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateTable() *RateTable {
	this := RateTable{}
	return &this
}

// NewRateTableWithDefaults instantiates a new RateTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateTableWithDefaults() *RateTable {
	this := RateTable{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *RateTable) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateTable) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RateTable) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *RateTable) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *RateTable) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateTable) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *RateTable) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *RateTable) SetLocality(v string) {
	o.Locality = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RateTable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateTable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RateTable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RateTable) SetName(v string) {
	o.Name = &v
}

// GetRateTableId returns the RateTableId field value if set, zero value otherwise.
func (o *RateTable) GetRateTableId() string {
	if o == nil || o.RateTableId == nil {
		var ret string
		return ret
	}
	return *o.RateTableId
}

// GetRateTableIdOk returns a tuple with the RateTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateTable) GetRateTableIdOk() (*string, bool) {
	if o == nil || o.RateTableId == nil {
		return nil, false
	}
	return o.RateTableId, true
}

// HasRateTableId returns a boolean if a field has been set.
func (o *RateTable) HasRateTableId() bool {
	if o != nil && o.RateTableId != nil {
		return true
	}

	return false
}

// SetRateTableId gets a reference to the given string and assigns it to the RateTableId field.
func (o *RateTable) SetRateTableId(v string) {
	o.RateTableId = &v
}

func (o RateTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RateTableId != nil {
		toSerialize["rateTableId"] = o.RateTableId
	}
	return json.Marshal(toSerialize)
}

type NullableRateTable struct {
	value *RateTable
	isSet bool
}

func (v NullableRateTable) Get() *RateTable {
	return v.value
}

func (v *NullableRateTable) Set(val *RateTable) {
	v.value = val
	v.isSet = true
}

func (v NullableRateTable) IsSet() bool {
	return v.isSet
}

func (v *NullableRateTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateTable(val *RateTable) *NullableRateTable {
	return &NullableRateTable{value: val, isSet: true}
}

func (v NullableRateTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


