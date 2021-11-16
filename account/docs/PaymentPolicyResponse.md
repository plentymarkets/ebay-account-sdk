# PaymentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | This field is for future use. | [optional] 
**Limit** | Pointer to **int32** | This field is for future use. | [optional] 
**Next** | Pointer to **string** | This field is for future use. | [optional] 
**Offset** | Pointer to **int32** | This field is for future use. | [optional] 
**Prev** | Pointer to **string** | This field is for future use. | [optional] 
**Total** | Pointer to **int32** | The total number of items retrieved in the result set.  &lt;br&gt;&lt;br&gt;If no items are found, this field is returned with a value of &lt;code&gt;0&lt;/code&gt;. | [optional] 
**PaymentPolicies** | Pointer to [**[]PaymentPolicy**](PaymentPolicy.md) | A list of the seller&#39;s payment policies. | [optional] 

## Methods

### NewPaymentPolicyResponse

`func NewPaymentPolicyResponse() *PaymentPolicyResponse`

NewPaymentPolicyResponse instantiates a new PaymentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPolicyResponseWithDefaults

`func NewPaymentPolicyResponseWithDefaults() *PaymentPolicyResponse`

NewPaymentPolicyResponseWithDefaults instantiates a new PaymentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PaymentPolicyResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PaymentPolicyResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PaymentPolicyResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PaymentPolicyResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *PaymentPolicyResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaymentPolicyResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaymentPolicyResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaymentPolicyResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *PaymentPolicyResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaymentPolicyResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaymentPolicyResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaymentPolicyResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *PaymentPolicyResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaymentPolicyResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaymentPolicyResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaymentPolicyResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *PaymentPolicyResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaymentPolicyResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaymentPolicyResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaymentPolicyResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetTotal

`func (o *PaymentPolicyResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaymentPolicyResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaymentPolicyResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaymentPolicyResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPaymentPolicies

`func (o *PaymentPolicyResponse) GetPaymentPolicies() []PaymentPolicy`

GetPaymentPolicies returns the PaymentPolicies field if non-nil, zero value otherwise.

### GetPaymentPoliciesOk

`func (o *PaymentPolicyResponse) GetPaymentPoliciesOk() (*[]PaymentPolicy, bool)`

GetPaymentPoliciesOk returns a tuple with the PaymentPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPolicies

`func (o *PaymentPolicyResponse) SetPaymentPolicies(v []PaymentPolicy)`

SetPaymentPolicies sets PaymentPolicies field to given value.

### HasPaymentPolicies

`func (o *PaymentPolicyResponse) HasPaymentPolicies() bool`

HasPaymentPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


