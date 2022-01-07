# SetPaymentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | The &lt;b&gt;CategoryTypeEnum&lt;/b&gt; value to which this policy applies. This container is used to discern accounts that sell motor vehicles from those that do not.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction:&lt;/b&gt; Currently, each policy can be set to only one &lt;b&gt;categoryTypes&lt;/b&gt; value at a time. | [optional] 
**Deposit** | Pointer to [**Deposit**](Deposit.md) |  | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the payment policy. | [optional] 
**FullPaymentDueIn** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ImmediatePay** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;,  payment is due upon receipt (eBay generates a receipt when the buyer agrees to purchase an item). The items will be available for other buyers until the payment is complete.&lt;br /&gt;&lt;br /&gt;This boolean must be set in the payment policy if the seller wants to create a listing that has an &lt;i&gt;immediate payment&lt;/i&gt; requirement. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; This container can be used for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program, but some requirements do not apply.&lt;/p&gt;&lt;b&gt;Default:&lt;/b&gt; False | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this payment policy applies. If this value is not specified, the value defaults to the seller&#39;s eBay registration site. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; A limited number of sellers, on a limited number of eBay marketplaces, are currently opted-in to the eBay managed payments program. To view the eBay marketplaces where managed payments are currently supported, see the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; landing page.&lt;/p&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this payment policy. Names must be unique for policies assigned to the same marketplace. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note&lt;/b&gt;: eBay will create a new payment policy for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program.&lt;/p&gt;&lt;b&gt;Max length:&lt;/b&gt; 64 | [optional] 
**PaymentInstructions** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; NO LONGER SUPPORTED. Although this field may be returned for some older payment business policies, payment instructions are no longer supported by payment business policies. If this field is returned, it can be ignored and these payment instructions will not appear in any listings that use the corresponding business policy.&lt;/p&gt;A free-form string field that allows sellers to add detailed payment instructions to their listings. The payment instructions appear on eBay&#39;s View Item and Checkout pages. &lt;br&gt;&lt;br&gt;eBay recommends sellers use this field to clarify payment policies for motor vehicle listings on eBay Motors. For example, sellers can include the specifics on the deposit (if required), pickup/delivery arrangements, and full payment details on the vehicle. &lt;br&gt;&lt;br&gt;The field allows only 500 characters as input, but due to the way the eBay web site UI treats characters, this field can return more than 500 characters in the response. For example, characters like &amp; and &#39; (ampersand and single quote) count as 5 characters each. &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction:&lt;/b&gt; This container is not supported for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program.  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 1000 | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | If the seller is not opted-in to managed payments, this container returns a list of the payment methods accepted by the seller.  &lt;br&gt;&lt;br&gt;When not opted-in to managed payments, each payment policy must specify at least one payment method. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; The &lt;b&gt;paymentMethods&lt;/b&gt; container is not returned if the seller is opted-in to the &lt;a href&#x3D;\&quot;https://pages.ebay.com/seller-center/service-and-payments/managed-payments-on-ebay.html\&quot; title&#x3D;\&quot;eBay Seller Center page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program.&lt;/p&gt; | [optional] 
**PaymentPolicyId** | Pointer to **string** | A unique eBay-assigned ID for a payment policy. This ID is generated when the policy is created. | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | A list of warnings related to request. This field normally returns empty, which indicates the request did not generate any warnings. | [optional] 

## Methods

### NewSetPaymentPolicyResponse

`func NewSetPaymentPolicyResponse() *SetPaymentPolicyResponse`

NewSetPaymentPolicyResponse instantiates a new SetPaymentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPaymentPolicyResponseWithDefaults

`func NewSetPaymentPolicyResponseWithDefaults() *SetPaymentPolicyResponse`

NewSetPaymentPolicyResponseWithDefaults instantiates a new SetPaymentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *SetPaymentPolicyResponse) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *SetPaymentPolicyResponse) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *SetPaymentPolicyResponse) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *SetPaymentPolicyResponse) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDeposit

`func (o *SetPaymentPolicyResponse) GetDeposit() Deposit`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *SetPaymentPolicyResponse) GetDepositOk() (*Deposit, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *SetPaymentPolicyResponse) SetDeposit(v Deposit)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *SetPaymentPolicyResponse) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDescription

`func (o *SetPaymentPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetPaymentPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetPaymentPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetPaymentPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullPaymentDueIn

`func (o *SetPaymentPolicyResponse) GetFullPaymentDueIn() TimeDuration`

GetFullPaymentDueIn returns the FullPaymentDueIn field if non-nil, zero value otherwise.

### GetFullPaymentDueInOk

`func (o *SetPaymentPolicyResponse) GetFullPaymentDueInOk() (*TimeDuration, bool)`

GetFullPaymentDueInOk returns a tuple with the FullPaymentDueIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPaymentDueIn

`func (o *SetPaymentPolicyResponse) SetFullPaymentDueIn(v TimeDuration)`

SetFullPaymentDueIn sets FullPaymentDueIn field to given value.

### HasFullPaymentDueIn

`func (o *SetPaymentPolicyResponse) HasFullPaymentDueIn() bool`

HasFullPaymentDueIn returns a boolean if a field has been set.

### GetImmediatePay

`func (o *SetPaymentPolicyResponse) GetImmediatePay() bool`

GetImmediatePay returns the ImmediatePay field if non-nil, zero value otherwise.

### GetImmediatePayOk

`func (o *SetPaymentPolicyResponse) GetImmediatePayOk() (*bool, bool)`

GetImmediatePayOk returns a tuple with the ImmediatePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediatePay

`func (o *SetPaymentPolicyResponse) SetImmediatePay(v bool)`

SetImmediatePay sets ImmediatePay field to given value.

### HasImmediatePay

`func (o *SetPaymentPolicyResponse) HasImmediatePay() bool`

HasImmediatePay returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *SetPaymentPolicyResponse) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *SetPaymentPolicyResponse) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *SetPaymentPolicyResponse) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *SetPaymentPolicyResponse) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *SetPaymentPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetPaymentPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetPaymentPolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SetPaymentPolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *SetPaymentPolicyResponse) GetPaymentInstructions() string`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *SetPaymentPolicyResponse) GetPaymentInstructionsOk() (*string, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *SetPaymentPolicyResponse) SetPaymentInstructions(v string)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *SetPaymentPolicyResponse) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *SetPaymentPolicyResponse) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *SetPaymentPolicyResponse) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *SetPaymentPolicyResponse) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *SetPaymentPolicyResponse) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPaymentPolicyId

`func (o *SetPaymentPolicyResponse) GetPaymentPolicyId() string`

GetPaymentPolicyId returns the PaymentPolicyId field if non-nil, zero value otherwise.

### GetPaymentPolicyIdOk

`func (o *SetPaymentPolicyResponse) GetPaymentPolicyIdOk() (*string, bool)`

GetPaymentPolicyIdOk returns a tuple with the PaymentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPolicyId

`func (o *SetPaymentPolicyResponse) SetPaymentPolicyId(v string)`

SetPaymentPolicyId sets PaymentPolicyId field to given value.

### HasPaymentPolicyId

`func (o *SetPaymentPolicyResponse) HasPaymentPolicyId() bool`

HasPaymentPolicyId returns a boolean if a field has been set.

### GetWarnings

`func (o *SetPaymentPolicyResponse) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SetPaymentPolicyResponse) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SetPaymentPolicyResponse) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SetPaymentPolicyResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


