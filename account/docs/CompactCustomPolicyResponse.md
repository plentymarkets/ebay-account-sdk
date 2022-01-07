# CompactCustomPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomPolicyId** | Pointer to **string** | The unique custom policy identifier for the policy being returned.&lt;br/&gt;&lt;br/&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This value is automatically assigned by the system when the policy is created.&lt;/span&gt; | [optional] 
**Label** | Pointer to **string** | Customer-facing label shown on View Item pages for items to which the policy applies. This seller-defined string is displayed as a system-generated hyperlink pointing to detailed policy information.&lt;br/&gt;&lt;br/&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 
**Name** | Pointer to **string** | The seller-defined name for the custom policy. Names must be unique for policies assigned to the same seller, policy type, and eBay marketplace.&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; This field is visible only to the seller. &lt;/span&gt;&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Max length:&lt;/b&gt; 65 | [optional] 
**PolicyType** | Pointer to **string** | Specifies the type of Custom Policy being returned. &lt;br/&gt;&lt;br/&gt;Two Custom Policy types are supported: &lt;ul&gt;&lt;li&gt;Product Compliance (PRODUCT_COMPLIANCE)&lt;/li&gt; &lt;li&gt;Takeback (TAKE_BACK)&lt;/li&gt;&lt;/ul&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:CustomPolicyTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewCompactCustomPolicyResponse

`func NewCompactCustomPolicyResponse() *CompactCustomPolicyResponse`

NewCompactCustomPolicyResponse instantiates a new CompactCustomPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompactCustomPolicyResponseWithDefaults

`func NewCompactCustomPolicyResponseWithDefaults() *CompactCustomPolicyResponse`

NewCompactCustomPolicyResponseWithDefaults instantiates a new CompactCustomPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomPolicyId

`func (o *CompactCustomPolicyResponse) GetCustomPolicyId() string`

GetCustomPolicyId returns the CustomPolicyId field if non-nil, zero value otherwise.

### GetCustomPolicyIdOk

`func (o *CompactCustomPolicyResponse) GetCustomPolicyIdOk() (*string, bool)`

GetCustomPolicyIdOk returns a tuple with the CustomPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPolicyId

`func (o *CompactCustomPolicyResponse) SetCustomPolicyId(v string)`

SetCustomPolicyId sets CustomPolicyId field to given value.

### HasCustomPolicyId

`func (o *CompactCustomPolicyResponse) HasCustomPolicyId() bool`

HasCustomPolicyId returns a boolean if a field has been set.

### GetLabel

`func (o *CompactCustomPolicyResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CompactCustomPolicyResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CompactCustomPolicyResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CompactCustomPolicyResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CompactCustomPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompactCustomPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompactCustomPolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompactCustomPolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyType

`func (o *CompactCustomPolicyResponse) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CompactCustomPolicyResponse) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CompactCustomPolicyResponse) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *CompactCustomPolicyResponse) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


