# CategoryType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Specifies the default policy for a &lt;b&gt;marketplaceId&lt;/b&gt; and &lt;b&gt;categoryTypes.name&lt;/b&gt; pair. Sellers can create multiple policies for any &lt;b&gt;marketplaceId&lt;/b&gt; and &lt;b&gt;categoryTypes.name&lt;/b&gt; combination. For example, you can create multiple fulfillment policies for one marketplace, where they all target the same category type &lt;b&gt;name&lt;/b&gt;. However, only one policy can be the default for any &lt;b&gt;marketplaceId&lt;/b&gt; and &lt;b&gt;name&lt;/b&gt; combination, and eBay designates the first policy created for a combination as the default.  &lt;br&gt;&lt;br&gt;If set to &lt;code&gt;true&lt;/code&gt;, this policy is the default policy for the associated &lt;b&gt;categoryTypes.name&lt;/b&gt; and &lt;b&gt;marketplaceId&lt;/b&gt; pair.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note&lt;/b&gt;: eBay considers the status of this field only when you create listings through the Web flow. If you create listings using the APIs, you must specifically set the policies you want applied to a listing in the payload of the call you use to create the listing. If you use the Web flow to create item listings, eBay uses the default policy for the marketplace and category type specified, unless you override the default.&lt;/span&gt; &lt;br&gt;&lt;br&gt;For more on default policies, see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/business-policies.html#default_policy\&quot;&gt;Changing the default policy for a category type&lt;/a&gt;. | [optional] 
**Name** | Pointer to **string** | The category type to which the policy applies (motor vehicles or non-motor vehicles). &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restrictions&lt;/b&gt;: &lt;ul&gt;&lt;li&gt;The &lt;code&gt;MOTORS_VEHICLES&lt;/code&gt; category type is not valid for return policies. eBay flows do not support the return of motor vehicles.&lt;/li&gt;&lt;li&gt;Only the &lt;code&gt;ALL_EXCLUDING_MOTORS_VEHICLES&lt;/code&gt; category type is supported for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program. Managed payments does not currently support the sale of motor vehicles.&lt;/li&gt;&lt;/ul&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:CategoryTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewCategoryType

`func NewCategoryType() *CategoryType`

NewCategoryType instantiates a new CategoryType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryTypeWithDefaults

`func NewCategoryTypeWithDefaults() *CategoryType`

NewCategoryTypeWithDefaults instantiates a new CategoryType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *CategoryType) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CategoryType) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CategoryType) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CategoryType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *CategoryType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryType) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


