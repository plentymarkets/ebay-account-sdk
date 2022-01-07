# CustomPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomPolicyId** | Pointer to **string** | The unique custom policy identifier for a policy.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This value is automatically assigned by the system when the policy is created.&lt;/span&gt; | [optional] 
**Description** | Pointer to **string** | Details of the seller&#39;s specific policy and terms associated with the policy. Buyers access this information from the View Item page for items to which the policy has been applied.&lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 15,000 | [optional] 
**Label** | Pointer to **string** | Customer-facing label shown on View Item pages for items to which the policy applies. This seller-defined string is displayed as a system-generated hyperlink pointing to detailed policy information.&lt;br/&gt;&lt;br/&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 
**Name** | Pointer to **string** | The seller-defined name for the custom policy. Names must be unique for policies assigned to the same seller, policy type, and eBay marketplace.&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This field is visible only to the seller. &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 
**PolicyType** | Pointer to **string** | Specifies the type of Custom Policy. &lt;br/&gt;&lt;br/&gt;Two Custom Policy types are supported: &lt;ul&gt;&lt;li&gt;Product Compliance (PRODUCT_COMPLIANCE)&lt;/li&gt; &lt;li&gt;Takeback (TAKE_BACK)&lt;/li&gt;&lt;/ul&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:CustomPolicyTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewCustomPolicy

`func NewCustomPolicy() *CustomPolicy`

NewCustomPolicy instantiates a new CustomPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPolicyWithDefaults

`func NewCustomPolicyWithDefaults() *CustomPolicy`

NewCustomPolicyWithDefaults instantiates a new CustomPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomPolicyId

`func (o *CustomPolicy) GetCustomPolicyId() string`

GetCustomPolicyId returns the CustomPolicyId field if non-nil, zero value otherwise.

### GetCustomPolicyIdOk

`func (o *CustomPolicy) GetCustomPolicyIdOk() (*string, bool)`

GetCustomPolicyIdOk returns a tuple with the CustomPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPolicyId

`func (o *CustomPolicy) SetCustomPolicyId(v string)`

SetCustomPolicyId sets CustomPolicyId field to given value.

### HasCustomPolicyId

`func (o *CustomPolicy) HasCustomPolicyId() bool`

HasCustomPolicyId returns a boolean if a field has been set.

### GetDescription

`func (o *CustomPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *CustomPolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomPolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomPolicy) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomPolicy) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CustomPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyType

`func (o *CustomPolicy) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CustomPolicy) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CustomPolicy) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *CustomPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


