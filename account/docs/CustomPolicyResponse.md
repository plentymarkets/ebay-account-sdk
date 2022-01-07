# CustomPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomPolicies** | Pointer to [**[]CompactCustomPolicyResponse**](CompactCustomPolicyResponse.md) | This array contains the custom policies that match the input criteria. | [optional] 
**Href** | Pointer to **string** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 
**Limit** | Pointer to **int32** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 
**Next** | Pointer to **string** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 
**Offset** | Pointer to **int32** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 
**Prev** | Pointer to **string** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 
**Total** | Pointer to **int32** | &lt;i&gt;This field is for future use.&lt;/i&gt; | [optional] 

## Methods

### NewCustomPolicyResponse

`func NewCustomPolicyResponse() *CustomPolicyResponse`

NewCustomPolicyResponse instantiates a new CustomPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPolicyResponseWithDefaults

`func NewCustomPolicyResponseWithDefaults() *CustomPolicyResponse`

NewCustomPolicyResponseWithDefaults instantiates a new CustomPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomPolicies

`func (o *CustomPolicyResponse) GetCustomPolicies() []CompactCustomPolicyResponse`

GetCustomPolicies returns the CustomPolicies field if non-nil, zero value otherwise.

### GetCustomPoliciesOk

`func (o *CustomPolicyResponse) GetCustomPoliciesOk() (*[]CompactCustomPolicyResponse, bool)`

GetCustomPoliciesOk returns a tuple with the CustomPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPolicies

`func (o *CustomPolicyResponse) SetCustomPolicies(v []CompactCustomPolicyResponse)`

SetCustomPolicies sets CustomPolicies field to given value.

### HasCustomPolicies

`func (o *CustomPolicyResponse) HasCustomPolicies() bool`

HasCustomPolicies returns a boolean if a field has been set.

### GetHref

`func (o *CustomPolicyResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CustomPolicyResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CustomPolicyResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CustomPolicyResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *CustomPolicyResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CustomPolicyResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CustomPolicyResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CustomPolicyResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *CustomPolicyResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CustomPolicyResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CustomPolicyResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CustomPolicyResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *CustomPolicyResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CustomPolicyResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CustomPolicyResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CustomPolicyResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *CustomPolicyResponse) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *CustomPolicyResponse) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *CustomPolicyResponse) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *CustomPolicyResponse) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetTotal

`func (o *CustomPolicyResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomPolicyResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CustomPolicyResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CustomPolicyResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


