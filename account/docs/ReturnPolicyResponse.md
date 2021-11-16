# ReturnPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | This field is for future use. | [optional] 
**Limit** | Pointer to **int32** | This field is for future use. | [optional] 
**Next** | Pointer to **string** | This field is for future use. | [optional] 
**Offset** | Pointer to **int32** | This field is for future use. | [optional] 
**Prev** | Pointer to **string** | This field is for future use. | [optional] 
**Total** | Pointer to **int32** | The total number of items retrieved in the result set.  &lt;br&gt;&lt;br&gt;If no items are found, this field is returned with a value of &lt;code&gt;0&lt;/code&gt;. | [optional] 
**ReturnPolicies** | Pointer to [**[]ReturnPolicy**](ReturnPolicy.md) | A list of the seller&#39;s return policies. | [optional] 

## Methods

### NewReturnPolicyResponse

`func NewReturnPolicyResponse() *ReturnPolicyResponse`

NewReturnPolicyResponse instantiates a new ReturnPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnPolicyResponseWithDefaults

`func NewReturnPolicyResponseWithDefaults() *ReturnPolicyResponse`

NewReturnPolicyResponseWithDefaults instantiates a new ReturnPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ReturnPolicyResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReturnPolicyResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReturnPolicyResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ReturnPolicyResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *ReturnPolicyResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReturnPolicyResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReturnPolicyResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReturnPolicyResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *ReturnPolicyResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReturnPolicyResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReturnPolicyResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ReturnPolicyResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *ReturnPolicyResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ReturnPolicyResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ReturnPolicyResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ReturnPolicyResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *ReturnPolicyResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *ReturnPolicyResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *ReturnPolicyResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *ReturnPolicyResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetTotal

`func (o *ReturnPolicyResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReturnPolicyResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReturnPolicyResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ReturnPolicyResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetReturnPolicies

`func (o *ReturnPolicyResponse) GetReturnPolicies() []ReturnPolicy`

GetReturnPolicies returns the ReturnPolicies field if non-nil, zero value otherwise.

### GetReturnPoliciesOk

`func (o *ReturnPolicyResponse) GetReturnPoliciesOk() (*[]ReturnPolicy, bool)`

GetReturnPoliciesOk returns a tuple with the ReturnPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicies

`func (o *ReturnPolicyResponse) SetReturnPolicies(v []ReturnPolicy)`

SetReturnPolicies sets ReturnPolicies field to given value.

### HasReturnPolicies

`func (o *ReturnPolicyResponse) HasReturnPolicies() bool`

HasReturnPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


