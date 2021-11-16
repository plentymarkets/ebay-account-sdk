# PaymentPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | The &lt;b&gt;CategoryTypeEnum&lt;/b&gt; value to which this policy applies. This container is used to discern accounts that sell motor vehicles from those that do not.&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction:&lt;/b&gt; Currently, each policy can be set to only one &lt;b&gt;categoryTypes&lt;/b&gt; value at a time. | [optional] 
**Deposit** | Pointer to [**Deposit**](Deposit.md) |  | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the payment policy for internal use (this value is not displayed to end users).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 250 | [optional] 
**FullPaymentDueIn** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ImmediatePay** | Pointer to **bool** | This field should be included and set to &#39;true&#39; if the seller wants to require immediate payment from the buyer for a fixed-price item or for an auction item where the buyer is using the &#39;Buy it Now&#39; option.&lt;br /&gt;&lt;br /&gt;Sellers who are not onboarded with eBay managed payments have the following additional requirements:&lt;ul&gt;&lt;li&gt;Offer PayPal as the only payment method&lt;/li&gt;&lt;li&gt;Provide a valid PayPal payment email address in the &lt;b&gt;paymentMethods.recipientAccountReference.referenceId&lt;/b&gt; field&lt;/li&gt;&lt;/ul&gt;&lt;br /&gt;&lt;br /&gt;&lt;b&gt;Default:&lt;/b&gt; False | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this payment policy applies. If this value is not specified, the value defaults to the seller&#39;s eBay registration site. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; A limited number of sellers, on a limited number of eBay marketplaces, are currently opted-in to the eBay managed payments program. To view the eBay marketplaces where managed payments are currently supported, see the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; landing page.&lt;/p&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this payment policy. Names must be unique for policies assigned to the same marketplace. &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; eBay will create a new payment policy for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program.&lt;/p&gt;&lt;b&gt;Max length:&lt;/b&gt; 64 | [optional] 
**PaymentInstructions** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; DO NOT USE THIS FIELD. Payment instructions are no longer supported by payment business policies.&lt;/p&gt;A free-form string field that allows sellers to add detailed payment instructions to their listings. The payment instructions appear on eBay&#39;s View Item and Checkout pages. &lt;br&gt;&lt;br&gt;eBay recommends sellers use this field to clarify payment policies for motor vehicle listings on eBay Motors. For example, sellers can include the specifics on the deposit (if required), pickup/delivery arrangements, and full payment details on the vehicle. &lt;br&gt;&lt;br&gt;The field allows only 500 characters as input, but due to the way the eBay web site UI treats characters, this field can return more than 500 characters in the response. For example, characters like &amp; and &#39; (ampersand and single quote) count as 5 characters each. &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction:&lt;/b&gt; This container is not supported for sellers who opt-in to the &lt;a href&#x3D;\&quot;/managed-payments\&quot; title&#x3D;\&quot;eBay Developers Program page\&quot; target&#x3D;\&quot;_blank\&quot;&gt;managed payments&lt;/a&gt; program.  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length:&lt;/b&gt; 1000 | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | This array is used to specify one or more payment methods that will be accepted for order payment.&lt;br /&gt;&lt;br /&gt;Because available electronic payments are strictly managed by eBay for sellers who are onboarded with eBay managed payments, this array is only needed by those sellers to specify offline payments (if applicable).&lt;br /&gt;&lt;br /&gt;For sellers who are not onboarded with managed payments, at least one electronic payment method is required. For most marketplaces and categories, this payment method is &#39;PAYPAL&#39;.&lt;br /&gt;&lt;br /&gt;For a motor vehicle payment policy, the payment policy must specify at least one of the following as an offline payment method:&lt;ul&gt;&lt;li&gt;CASH_ON_PICKUP (not available in the US)&lt;/li&gt;&lt;li&gt;CASHIER_CHECK&lt;/li&gt;&lt;li&gt;MONEY_ORDER&lt;/li&gt;&lt;li&gt;PERSONAL_CHECK&lt;/li&gt;&lt;/ul&gt; | [optional] 

## Methods

### NewPaymentPolicyRequest

`func NewPaymentPolicyRequest() *PaymentPolicyRequest`

NewPaymentPolicyRequest instantiates a new PaymentPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentPolicyRequestWithDefaults

`func NewPaymentPolicyRequestWithDefaults() *PaymentPolicyRequest`

NewPaymentPolicyRequestWithDefaults instantiates a new PaymentPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *PaymentPolicyRequest) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *PaymentPolicyRequest) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *PaymentPolicyRequest) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *PaymentPolicyRequest) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDeposit

`func (o *PaymentPolicyRequest) GetDeposit() Deposit`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *PaymentPolicyRequest) GetDepositOk() (*Deposit, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *PaymentPolicyRequest) SetDeposit(v Deposit)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *PaymentPolicyRequest) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFullPaymentDueIn

`func (o *PaymentPolicyRequest) GetFullPaymentDueIn() TimeDuration`

GetFullPaymentDueIn returns the FullPaymentDueIn field if non-nil, zero value otherwise.

### GetFullPaymentDueInOk

`func (o *PaymentPolicyRequest) GetFullPaymentDueInOk() (*TimeDuration, bool)`

GetFullPaymentDueInOk returns a tuple with the FullPaymentDueIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPaymentDueIn

`func (o *PaymentPolicyRequest) SetFullPaymentDueIn(v TimeDuration)`

SetFullPaymentDueIn sets FullPaymentDueIn field to given value.

### HasFullPaymentDueIn

`func (o *PaymentPolicyRequest) HasFullPaymentDueIn() bool`

HasFullPaymentDueIn returns a boolean if a field has been set.

### GetImmediatePay

`func (o *PaymentPolicyRequest) GetImmediatePay() bool`

GetImmediatePay returns the ImmediatePay field if non-nil, zero value otherwise.

### GetImmediatePayOk

`func (o *PaymentPolicyRequest) GetImmediatePayOk() (*bool, bool)`

GetImmediatePayOk returns a tuple with the ImmediatePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediatePay

`func (o *PaymentPolicyRequest) SetImmediatePay(v bool)`

SetImmediatePay sets ImmediatePay field to given value.

### HasImmediatePay

`func (o *PaymentPolicyRequest) HasImmediatePay() bool`

HasImmediatePay returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *PaymentPolicyRequest) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *PaymentPolicyRequest) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *PaymentPolicyRequest) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *PaymentPolicyRequest) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *PaymentPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *PaymentPolicyRequest) GetPaymentInstructions() string`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *PaymentPolicyRequest) GetPaymentInstructionsOk() (*string, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *PaymentPolicyRequest) SetPaymentInstructions(v string)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *PaymentPolicyRequest) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *PaymentPolicyRequest) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentPolicyRequest) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentPolicyRequest) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentPolicyRequest) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


