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

// Programs A list of the supported seller programs.
type Programs struct {
	// A list of seller programs.
	Programs *[]Program `json:"programs,omitempty"`
}

// NewPrograms instantiates a new Programs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrograms() *Programs {
	this := Programs{}
	return &this
}

// NewProgramsWithDefaults instantiates a new Programs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramsWithDefaults() *Programs {
	this := Programs{}
	return &this
}

// GetPrograms returns the Programs field value if set, zero value otherwise.
func (o *Programs) GetPrograms() []Program {
	if o == nil || o.Programs == nil {
		var ret []Program
		return ret
	}
	return *o.Programs
}

// GetProgramsOk returns a tuple with the Programs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Programs) GetProgramsOk() (*[]Program, bool) {
	if o == nil || o.Programs == nil {
		return nil, false
	}
	return o.Programs, true
}

// HasPrograms returns a boolean if a field has been set.
func (o *Programs) HasPrograms() bool {
	if o != nil && o.Programs != nil {
		return true
	}

	return false
}

// SetPrograms gets a reference to the given []Program and assigns it to the Programs field.
func (o *Programs) SetPrograms(v []Program) {
	o.Programs = &v
}

func (o Programs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Programs != nil {
		toSerialize["programs"] = o.Programs
	}
	return json.Marshal(toSerialize)
}

type NullablePrograms struct {
	value *Programs
	isSet bool
}

func (v NullablePrograms) Get() *Programs {
	return v.value
}

func (v *NullablePrograms) Set(val *Programs) {
	v.value = val
	v.isSet = true
}

func (v NullablePrograms) IsSet() bool {
	return v.isSet
}

func (v *NullablePrograms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrograms(val *Programs) *NullablePrograms {
	return &NullablePrograms{value: val, isSet: true}
}

func (v NullablePrograms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrograms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


