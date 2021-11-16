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

// Error A container that defines the elements of error and warning messages.
type Error struct {
	// The category type for this error or warning. It is a string that can have one of three values:<ul><li><code>Application</code>: Indicates an exception or error occurred in the application code or at runtime. Examples include catching an exception in a service's business logic, system failures, or request errors from a dependency.</li><li><code>Business</code>: Used when your service or a dependent service refused to continue processing on the resource because of a business rule violation such as \"Seller does not ship item to Antarctica\" or \"Buyer ineligible to purchase an alcoholic item\". Business errors are not syntactical input errors.</li><li><code>Request</code>: Used when there is anything wrong with the request, such as authentication, syntactical errors, rate limiting or missing headers, bad HTTP header values, and so on.</li></ul>
	Category *string `json:"category,omitempty"`
	// Name of the domain ,or primary system, of the service or application where the error occurred.
	Domain *string `json:"domain,omitempty"`
	// A positive integer that uniquely identifies the specific error condition that occurred. Your application can use error codes as identifiers in your customized error-handling algorithms.
	ErrorId *int32 `json:"errorId,omitempty"`
	// Identifies specific request elements associated with the error, if any. inputRefId's response is format specific. For JSON, use <i>JSONPath</i> notation.
	InputRefIds *[]string `json:"inputRefIds,omitempty"`
	// A more detailed explanation of the error than given in the <code>message</code> error field.
	LongMessage *string `json:"longMessage,omitempty"`
	// Information on how to correct the problem, in the end user's terms and language where applicable. Its value is at most 50 characters long. If applicable, the value is localized in the end user's requested locale.
	Message *string `json:"message,omitempty"`
	// Identifies specific response elements associated with the error, if any. Path format is the same as <code>inputRefId</code>.
	OutputRefIds *[]string `json:"outputRefIds,omitempty"`
	// This optional list of name/value pairs that contain context-specific <code>ErrorParameter</code> objects, with each item in the list being a parameter (or input field name) that caused an error condition. Each <code>ErrorParameter</code> object consists of two fields, a <code>name</code> and a <code>value</code>.
	Parameters *[]ErrorParameter `json:"parameters,omitempty"`
	// If present, indicates the subsystem in which the error occurred.
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Error) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Error) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Error) SetCategory(v string) {
	o.Category = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Error) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Error) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Error) SetDomain(v string) {
	o.Domain = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *Error) GetErrorId() int32 {
	if o == nil || o.ErrorId == nil {
		var ret int32
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorIdOk() (*int32, bool) {
	if o == nil || o.ErrorId == nil {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *Error) HasErrorId() bool {
	if o != nil && o.ErrorId != nil {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given int32 and assigns it to the ErrorId field.
func (o *Error) SetErrorId(v int32) {
	o.ErrorId = &v
}

// GetInputRefIds returns the InputRefIds field value if set, zero value otherwise.
func (o *Error) GetInputRefIds() []string {
	if o == nil || o.InputRefIds == nil {
		var ret []string
		return ret
	}
	return *o.InputRefIds
}

// GetInputRefIdsOk returns a tuple with the InputRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetInputRefIdsOk() (*[]string, bool) {
	if o == nil || o.InputRefIds == nil {
		return nil, false
	}
	return o.InputRefIds, true
}

// HasInputRefIds returns a boolean if a field has been set.
func (o *Error) HasInputRefIds() bool {
	if o != nil && o.InputRefIds != nil {
		return true
	}

	return false
}

// SetInputRefIds gets a reference to the given []string and assigns it to the InputRefIds field.
func (o *Error) SetInputRefIds(v []string) {
	o.InputRefIds = &v
}

// GetLongMessage returns the LongMessage field value if set, zero value otherwise.
func (o *Error) GetLongMessage() string {
	if o == nil || o.LongMessage == nil {
		var ret string
		return ret
	}
	return *o.LongMessage
}

// GetLongMessageOk returns a tuple with the LongMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetLongMessageOk() (*string, bool) {
	if o == nil || o.LongMessage == nil {
		return nil, false
	}
	return o.LongMessage, true
}

// HasLongMessage returns a boolean if a field has been set.
func (o *Error) HasLongMessage() bool {
	if o != nil && o.LongMessage != nil {
		return true
	}

	return false
}

// SetLongMessage gets a reference to the given string and assigns it to the LongMessage field.
func (o *Error) SetLongMessage(v string) {
	o.LongMessage = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Error) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Error) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Error) SetMessage(v string) {
	o.Message = &v
}

// GetOutputRefIds returns the OutputRefIds field value if set, zero value otherwise.
func (o *Error) GetOutputRefIds() []string {
	if o == nil || o.OutputRefIds == nil {
		var ret []string
		return ret
	}
	return *o.OutputRefIds
}

// GetOutputRefIdsOk returns a tuple with the OutputRefIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetOutputRefIdsOk() (*[]string, bool) {
	if o == nil || o.OutputRefIds == nil {
		return nil, false
	}
	return o.OutputRefIds, true
}

// HasOutputRefIds returns a boolean if a field has been set.
func (o *Error) HasOutputRefIds() bool {
	if o != nil && o.OutputRefIds != nil {
		return true
	}

	return false
}

// SetOutputRefIds gets a reference to the given []string and assigns it to the OutputRefIds field.
func (o *Error) SetOutputRefIds(v []string) {
	o.OutputRefIds = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Error) GetParameters() []ErrorParameter {
	if o == nil || o.Parameters == nil {
		var ret []ErrorParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetParametersOk() (*[]ErrorParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Error) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ErrorParameter and assigns it to the Parameters field.
func (o *Error) SetParameters(v []ErrorParameter) {
	o.Parameters = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Error) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *Error) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Error) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.ErrorId != nil {
		toSerialize["errorId"] = o.ErrorId
	}
	if o.InputRefIds != nil {
		toSerialize["inputRefIds"] = o.InputRefIds
	}
	if o.LongMessage != nil {
		toSerialize["longMessage"] = o.LongMessage
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.OutputRefIds != nil {
		toSerialize["outputRefIds"] = o.OutputRefIds
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


