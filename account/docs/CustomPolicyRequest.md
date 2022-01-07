# CustomPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Details of the seller&#39;s specific policy and terms for this policy.&lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 15,000 | [optional] 
**Label** | Pointer to **string** | Customer-facing label shown on View Item pages for items to which the policy applies. This seller-defined string is displayed as a system-generated hyperlink pointing to detailed policy information.&lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 
**Name** | Pointer to **string** | The seller-defined name for the custom policy. Names must be unique for policies assigned to the same seller, policy type, and eBay marketplace.&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This field is visible only to the seller. &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 

## Methods

### NewCustomPolicyRequest

`func NewCustomPolicyRequest() *CustomPolicyRequest`

NewCustomPolicyRequest instantiates a new CustomPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPolicyRequestWithDefaults

`func NewCustomPolicyRequestWithDefaults() *CustomPolicyRequest`

NewCustomPolicyRequestWithDefaults instantiates a new CustomPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *CustomPolicyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomPolicyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomPolicyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomPolicyRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CustomPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


