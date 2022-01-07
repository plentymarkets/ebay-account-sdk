/*
Account API

The <b>Account API</b> gives sellers the ability to configure their eBay seller accounts, including the seller's policies (seller-defined custom policies and eBay business policies), opt in and out of eBay seller programs, configure sales tax tables, and get account information.  <br><br>For details on the availability of the methods in this API, see <a href=\"/api-docs/sell/account/overview.html#requirements\">Account API requirements and restrictions</a>.

API version: v1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// CustomPolicyRequest struct for CustomPolicyRequest
type CustomPolicyRequest struct {
	// Details of the seller's specific policy and terms for this policy.<br><br><b>Max length:</b> 15,000
	Description *string `json:"description,omitempty"`
	// Customer-facing label shown on View Item pages for items to which the policy applies. This seller-defined string is displayed as a system-generated hyperlink pointing to detailed policy information.<br><br><b>Max length:</b> 65
	Label *string `json:"label,omitempty"`
	// The seller-defined name for the custom policy. Names must be unique for policies assigned to the same seller, policy type, and eBay marketplace.<br /><span class=\"tablenote\"><strong>Note:</strong> This field is visible only to the seller. </span><br/><br/><b>Max length:</b> 65
	Name *string `json:"name,omitempty"`
}

// NewCustomPolicyRequest instantiates a new CustomPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPolicyRequest() *CustomPolicyRequest {
	this := CustomPolicyRequest{}
	return &this
}

// NewCustomPolicyRequestWithDefaults instantiates a new CustomPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPolicyRequestWithDefaults() *CustomPolicyRequest {
	this := CustomPolicyRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomPolicyRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomPolicyRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomPolicyRequest) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPolicyRequest) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomPolicyRequest) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CustomPolicyRequest) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomPolicyRequest) SetName(v string) {
	o.Name = &v
}

func (o CustomPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCustomPolicyRequest struct {
	value *CustomPolicyRequest
	isSet bool
}

func (v NullableCustomPolicyRequest) Get() *CustomPolicyRequest {
	return v.value
}

func (v *NullableCustomPolicyRequest) Set(val *CustomPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPolicyRequest(val *CustomPolicyRequest) *NullableCustomPolicyRequest {
	return &NullableCustomPolicyRequest{value: val, isSet: true}
}

func (v NullableCustomPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


