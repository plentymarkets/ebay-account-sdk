# FulfillmentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentPolicies** | Pointer to [**[]FulfillmentPolicy**](FulfillmentPolicy.md) | A list of the seller&#39;s fulfillment policies. | [optional] 
**Href** | Pointer to **string** | This field is for future use. | [optional] 
**Limit** | Pointer to **int32** | This field is for future use. | [optional] 
**Next** | Pointer to **string** | This field is for future use. | [optional] 
**Offset** | Pointer to **int32** | This field is for future use. | [optional] 
**Prev** | Pointer to **string** | This field is for future use. | [optional] 
**Total** | Pointer to **int32** | The total number of items retrieved in the result set.  &lt;br&gt;&lt;br&gt;If no items are found, this field is returned with a value of &lt;code&gt;0&lt;/code&gt;. | [optional] 

## Methods

### NewFulfillmentPolicyResponse

`func NewFulfillmentPolicyResponse() *FulfillmentPolicyResponse`

NewFulfillmentPolicyResponse instantiates a new FulfillmentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentPolicyResponseWithDefaults

`func NewFulfillmentPolicyResponseWithDefaults() *FulfillmentPolicyResponse`

NewFulfillmentPolicyResponseWithDefaults instantiates a new FulfillmentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentPolicies

`func (o *FulfillmentPolicyResponse) GetFulfillmentPolicies() []FulfillmentPolicy`

GetFulfillmentPolicies returns the FulfillmentPolicies field if non-nil, zero value otherwise.

### GetFulfillmentPoliciesOk

`func (o *FulfillmentPolicyResponse) GetFulfillmentPoliciesOk() (*[]FulfillmentPolicy, bool)`

GetFulfillmentPoliciesOk returns a tuple with the FulfillmentPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentPolicies

`func (o *FulfillmentPolicyResponse) SetFulfillmentPolicies(v []FulfillmentPolicy)`

SetFulfillmentPolicies sets FulfillmentPolicies field to given value.

### HasFulfillmentPolicies

`func (o *FulfillmentPolicyResponse) HasFulfillmentPolicies() bool`

HasFulfillmentPolicies returns a boolean if a field has been set.

### GetHref

`func (o *FulfillmentPolicyResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FulfillmentPolicyResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FulfillmentPolicyResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FulfillmentPolicyResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *FulfillmentPolicyResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FulfillmentPolicyResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FulfillmentPolicyResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FulfillmentPolicyResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *FulfillmentPolicyResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FulfillmentPolicyResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FulfillmentPolicyResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FulfillmentPolicyResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *FulfillmentPolicyResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FulfillmentPolicyResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FulfillmentPolicyResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *FulfillmentPolicyResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *FulfillmentPolicyResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *FulfillmentPolicyResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *FulfillmentPolicyResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *FulfillmentPolicyResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetTotal

`func (o *FulfillmentPolicyResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FulfillmentPolicyResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FulfillmentPolicyResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FulfillmentPolicyResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


