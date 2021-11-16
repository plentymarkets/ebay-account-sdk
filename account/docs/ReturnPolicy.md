# ReturnPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | For return policies, this field can be set to only &lt;code&gt;ALL_EXCLUDING_MOTORS_VEHICLES&lt;/code&gt; (returns on motor vehicles are not processed through eBay flows) &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: &lt;code&gt;ALL_EXCLUDING_MOTORS_VEHICLES&lt;/code&gt; (for return policies only) | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the return policy for internal use (this value is not displayed to end users).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 250 | [optional] 
**ExtendedHolidayReturnsOffered** | Pointer to **bool** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is neither read nor returned.&lt;/p&gt;  &lt;p&gt;If set to &lt;code&gt;true&lt;/code&gt;, the seller offers an &lt;em&gt;Extended Holiday Returns&lt;/em&gt; policy for their listings.&lt;/p&gt;  &lt;p&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;IMPORTANT:&lt;/strong&gt; Extended Holiday Returns is a seasonally available feature that is offered on some eBay marketplaces. To see if the feature is enabled in any given year, check the &lt;a href&#x3D;\&quot;http://pages.ebay.com/seller-center/shipping/returns.html\&quot;&gt;Returns on eBay&lt;/a&gt; page before the holiday season begins. If the feature is not enabled for the season, this field is ignored.&lt;/span&gt;&lt;/p&gt;  &lt;p&gt;The extended holiday returns period is defined by three dates:&lt;/p&gt; &lt;ul&gt;&lt;li&gt;The start date &#x3D; start of November.&lt;/li&gt;&lt;li&gt;The purchase cutoff date &#x3D; end of the year.&lt;/li&gt;&lt;li&gt;The end date &#x3D; end of January.&lt;/li&gt;&lt;/ul&gt;  &lt;p&gt;The above dates may vary by a few days each year. Sellers are notified of the current dates on their eBay marketplace before the holiday period starts.&lt;/p&gt;  &lt;p&gt;Sellers can specify Extended Holiday Returns (as well as their regular non-holiday returns period) for chosen listings at any time during the year. The Extended Holiday Returns offer is not visible in listings until the start date of current year&#39;s holiday returns period, at which point it overrides the non-holiday returns policy. Buyers can see the Extended Holiday Returns offer in listings displayed through the purchase cutoff date and are able to return those purchases until the end date of the period.&lt;/p&gt;  &lt;p&gt;After the purchase cutoff date, the Extended Holiday Returns offer automatically disappears from the listings and the seller&#39;s non-holiday returns period reappears. Purchases made from that point on are subject to the non-holiday returns period, while purchases made before the cutoff date still have until the end of the period to return under the program.&lt;/p&gt;  &lt;p&gt;If the value of &lt;strong&gt;holidayReturns&lt;/strong&gt; is &lt;code&gt;false&lt;/code&gt; for an item, the return period specified by the &lt;strong&gt;returnsWithinOption&lt;/strong&gt; field applies, regardless of the purchase date. If the item is listed with a policy of no returns, &lt;strong&gt;holidayReturns&lt;/strong&gt; is automatically reset to &lt;code&gt;false&lt;/code&gt;.&lt;/p&gt; | [optional] 
**InternationalOverride** | Pointer to [**InternationalReturnOverrideType**](InternationalReturnOverrideType.md) |  | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this return policy applies. If this value is not specified, value defaults to the seller&#39;s eBay registration site. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this return policy. Names must be unique for policies assigned to the same marketplace. &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 64 | [optional] 
**RefundMethod** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; this field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value other than &lt;code&gt;MONEY_BACK&lt;/code&gt; will be treated as &lt;code&gt;MONEY_BACK&lt;/code&gt; (although for a period of time, eBay will store and return the legacy values to preserve backwards compatibility).&lt;/p&gt;  Indicates the method the seller uses to compensate the buyer for returned items. The return method specified applies only to &lt;a href&#x3D;\&quot;http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\&quot; target&#x3D;\&quot;_blank\&quot;&gt;remorse returns&lt;/a&gt;. &lt;br&gt;&lt;br&gt;Each eBay marketplace may support different sets of refund methods and marketplaces can also have differing default values for this field. Sellers are obligated to honor the refund method displayed in their listings. Call &lt;b&gt;GeteBayDetails&lt;/b&gt; in the Trading API to see the refund methods supported by the marketplaces you sell into. &lt;br&gt;&lt;br&gt;We recommend you set this field to the value of your preferred refund method and that you use the &lt;b&gt;description&lt;/b&gt; field to detail the seller&#39;s return policy (such as indicating how quickly the seller will process a refund, whether the seller must receive the item before processing a refund, and other similar useful details). &lt;br&gt;&lt;br&gt;You cannot modify this value in a Revise item call if (1) the listing has bids or (2) the listing ends within 12 hours. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:RefundMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**RestockingFeePercentage** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is ignored, it is neither read nor returned.&lt;/p&gt;  &lt;p&gt;Sellers who accept returns should include this field if they charge buyers a restocking fee when items are returned. A restocking fee comes into play only when an item is returned due to buyer remorse and/or a purchasing mistake, but sellers cannot charge a restocking fee for SNAD-related returns. The total amount returned to the buyer is reduced by the cost of the item multiplied by the percentage indicated by this field. &lt;p&gt;Allowable restocking fee values are:&lt;/p&gt; &lt;ul&gt;&lt;li&gt;&lt;code&gt;0.0&lt;/code&gt;: No restocking fee is charged to the buyer&lt;/li&gt;&lt;li&gt;&lt;code&gt;10.0&lt;/code&gt;: 10 percent of the item price is charged to the buyer&lt;/li&gt;&lt;li&gt;&lt;code&gt;15.0&lt;/code&gt;: 15 percent of the item price is charged to the buyer&lt;/li&gt; &lt;li&gt;&lt;code&gt;20.0&lt;/code&gt;: Up to 20 percent of the item price is charged to the buyer&lt;/li&gt;&lt;/ul&gt; | [optional] 
**ReturnInstructions** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field is being deprecated on many marketplaces. Once deprecated, this field will be ignored on marketplaces where it is not supported and it will neither be read nor returned.&lt;/p&gt;  &lt;p&gt;This optional field contains the seller&#39;s detailed explanation for their return policy and is displayed in the Return Policy section of the View Item page. This field is valid in only the following marketplaces (the field is otherwise ignored):&lt;/p&gt; &lt;ul&gt; &lt;li&gt;Germany (DE)&lt;/li&gt; &lt;li&gt;Spain (ES)&lt;/li&gt; &lt;li&gt;France (FR)&lt;/li&gt; &lt;li&gt;Italy (IT)&lt;/li&gt; &lt;/ul&gt; Where valid, sellers can use this field to add details about their return policies. eBay uses this text string as-is in the Return Policy section of the View Item page. Avoid HTML and avoid character entity references (such as &amp;amp;amp;pound; or &amp;amp;amp;#163;). To include special characters in the return policy description, use the literal UTF-8 or ISO-8559-1 character (e.g. &amp;amp;#163;).  &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 5000 (8000 for DE) | [optional] 
**ReturnMethod** | Pointer to **string** | Valid in the US marketplace only, this optional field indicates additional services (other than money-back) that sellers can offer buyers for &lt;a href&#x3D;\&quot;http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\&quot; target&#x3D;\&quot;_blank\&quot;&gt;remorse returns&lt;/a&gt;.  &lt;br&gt;&lt;br&gt;As of version 1.2.0, the only accepted value for this field is &lt;code&gt;REPLACEMENT&lt;/code&gt;. This field is valid in only the US marketplace, any supplied value is ignored in other marketplaces. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ReturnPeriod** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ReturnPolicyId** | Pointer to **string** | A unique eBay-assigned ID for this policy. This ID value is appended to the end of the &lt;code&gt;Location&lt;/code&gt; URI that is returned as a response header when you call &lt;b&gt;createReturnPolicy&lt;/b&gt;). | [optional] 
**ReturnsAccepted** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller accepts returns.  &lt;br&gt;&lt;br&gt;Call the &lt;b&gt;getReturnPolicies&lt;/b&gt; in the Metadata API to see what categories require returns to be offered for listings in each category. Also, note that some European marketplaces (for example, UK, IE, and DE) require sellers to accept returns for fixed-price items and auctions listed with Buy It Now. For details, see &lt;a href&#x3D;\&quot;http://pages.ebay.co.uk/help/policies/user-agreement.html#returns\&quot;&gt;Returns and the Law (UK)&lt;/a&gt;.  &lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt;Top-Rated sellers must accept item returns and the &lt;b&gt;handlingTime&lt;/b&gt; should be set to zero days or one day for a listing to receive a Top-Rated Plus badge on the View Item or search result pages. For more information on eBay&#39;s Top-Rated seller program, see &lt;a href&#x3D;\&quot;http://pages.ebay.com/help/sell/top-rated.html\&quot;&gt;Becoming a Top Rated Seller and qualifying for Top Rated Plus benefits&lt;/a&gt;.&lt;/span&gt; | [optional] 
**ReturnShippingCostPayer** | Pointer to **string** | This field indicates who is responsible for paying for the shipping charges for returned items. The field can be set to either &lt;code&gt;BUYER&lt;/code&gt; or &lt;code&gt;SELLER&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;Depending on the return policy and specifics of the return, either the buyer or the seller can be responsible for the return shipping costs. Note that the seller is always responsible for return shipping costs for SNAD-related issues. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnShippingCostPayerEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewReturnPolicy

`func NewReturnPolicy() *ReturnPolicy`

NewReturnPolicy instantiates a new ReturnPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnPolicyWithDefaults

`func NewReturnPolicyWithDefaults() *ReturnPolicy`

NewReturnPolicyWithDefaults instantiates a new ReturnPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *ReturnPolicy) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *ReturnPolicy) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *ReturnPolicy) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *ReturnPolicy) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDescription

`func (o *ReturnPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReturnPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReturnPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReturnPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtendedHolidayReturnsOffered

`func (o *ReturnPolicy) GetExtendedHolidayReturnsOffered() bool`

GetExtendedHolidayReturnsOffered returns the ExtendedHolidayReturnsOffered field if non-nil, zero value otherwise.

### GetExtendedHolidayReturnsOfferedOk

`func (o *ReturnPolicy) GetExtendedHolidayReturnsOfferedOk() (*bool, bool)`

GetExtendedHolidayReturnsOfferedOk returns a tuple with the ExtendedHolidayReturnsOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHolidayReturnsOffered

`func (o *ReturnPolicy) SetExtendedHolidayReturnsOffered(v bool)`

SetExtendedHolidayReturnsOffered sets ExtendedHolidayReturnsOffered field to given value.

### HasExtendedHolidayReturnsOffered

`func (o *ReturnPolicy) HasExtendedHolidayReturnsOffered() bool`

HasExtendedHolidayReturnsOffered returns a boolean if a field has been set.

### GetInternationalOverride

`func (o *ReturnPolicy) GetInternationalOverride() InternationalReturnOverrideType`

GetInternationalOverride returns the InternationalOverride field if non-nil, zero value otherwise.

### GetInternationalOverrideOk

`func (o *ReturnPolicy) GetInternationalOverrideOk() (*InternationalReturnOverrideType, bool)`

GetInternationalOverrideOk returns a tuple with the InternationalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOverride

`func (o *ReturnPolicy) SetInternationalOverride(v InternationalReturnOverrideType)`

SetInternationalOverride sets InternationalOverride field to given value.

### HasInternationalOverride

`func (o *ReturnPolicy) HasInternationalOverride() bool`

HasInternationalOverride returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *ReturnPolicy) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *ReturnPolicy) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *ReturnPolicy) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *ReturnPolicy) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *ReturnPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReturnPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReturnPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReturnPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefundMethod

`func (o *ReturnPolicy) GetRefundMethod() string`

GetRefundMethod returns the RefundMethod field if non-nil, zero value otherwise.

### GetRefundMethodOk

`func (o *ReturnPolicy) GetRefundMethodOk() (*string, bool)`

GetRefundMethodOk returns a tuple with the RefundMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundMethod

`func (o *ReturnPolicy) SetRefundMethod(v string)`

SetRefundMethod sets RefundMethod field to given value.

### HasRefundMethod

`func (o *ReturnPolicy) HasRefundMethod() bool`

HasRefundMethod returns a boolean if a field has been set.

### GetRestockingFeePercentage

`func (o *ReturnPolicy) GetRestockingFeePercentage() string`

GetRestockingFeePercentage returns the RestockingFeePercentage field if non-nil, zero value otherwise.

### GetRestockingFeePercentageOk

`func (o *ReturnPolicy) GetRestockingFeePercentageOk() (*string, bool)`

GetRestockingFeePercentageOk returns a tuple with the RestockingFeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockingFeePercentage

`func (o *ReturnPolicy) SetRestockingFeePercentage(v string)`

SetRestockingFeePercentage sets RestockingFeePercentage field to given value.

### HasRestockingFeePercentage

`func (o *ReturnPolicy) HasRestockingFeePercentage() bool`

HasRestockingFeePercentage returns a boolean if a field has been set.

### GetReturnInstructions

`func (o *ReturnPolicy) GetReturnInstructions() string`

GetReturnInstructions returns the ReturnInstructions field if non-nil, zero value otherwise.

### GetReturnInstructionsOk

`func (o *ReturnPolicy) GetReturnInstructionsOk() (*string, bool)`

GetReturnInstructionsOk returns a tuple with the ReturnInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnInstructions

`func (o *ReturnPolicy) SetReturnInstructions(v string)`

SetReturnInstructions sets ReturnInstructions field to given value.

### HasReturnInstructions

`func (o *ReturnPolicy) HasReturnInstructions() bool`

HasReturnInstructions returns a boolean if a field has been set.

### GetReturnMethod

`func (o *ReturnPolicy) GetReturnMethod() string`

GetReturnMethod returns the ReturnMethod field if non-nil, zero value otherwise.

### GetReturnMethodOk

`func (o *ReturnPolicy) GetReturnMethodOk() (*string, bool)`

GetReturnMethodOk returns a tuple with the ReturnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMethod

`func (o *ReturnPolicy) SetReturnMethod(v string)`

SetReturnMethod sets ReturnMethod field to given value.

### HasReturnMethod

`func (o *ReturnPolicy) HasReturnMethod() bool`

HasReturnMethod returns a boolean if a field has been set.

### GetReturnPeriod

`func (o *ReturnPolicy) GetReturnPeriod() TimeDuration`

GetReturnPeriod returns the ReturnPeriod field if non-nil, zero value otherwise.

### GetReturnPeriodOk

`func (o *ReturnPolicy) GetReturnPeriodOk() (*TimeDuration, bool)`

GetReturnPeriodOk returns a tuple with the ReturnPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPeriod

`func (o *ReturnPolicy) SetReturnPeriod(v TimeDuration)`

SetReturnPeriod sets ReturnPeriod field to given value.

### HasReturnPeriod

`func (o *ReturnPolicy) HasReturnPeriod() bool`

HasReturnPeriod returns a boolean if a field has been set.

### GetReturnPolicyId

`func (o *ReturnPolicy) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *ReturnPolicy) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *ReturnPolicy) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *ReturnPolicy) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### GetReturnsAccepted

`func (o *ReturnPolicy) GetReturnsAccepted() bool`

GetReturnsAccepted returns the ReturnsAccepted field if non-nil, zero value otherwise.

### GetReturnsAcceptedOk

`func (o *ReturnPolicy) GetReturnsAcceptedOk() (*bool, bool)`

GetReturnsAcceptedOk returns a tuple with the ReturnsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsAccepted

`func (o *ReturnPolicy) SetReturnsAccepted(v bool)`

SetReturnsAccepted sets ReturnsAccepted field to given value.

### HasReturnsAccepted

`func (o *ReturnPolicy) HasReturnsAccepted() bool`

HasReturnsAccepted returns a boolean if a field has been set.

### GetReturnShippingCostPayer

`func (o *ReturnPolicy) GetReturnShippingCostPayer() string`

GetReturnShippingCostPayer returns the ReturnShippingCostPayer field if non-nil, zero value otherwise.

### GetReturnShippingCostPayerOk

`func (o *ReturnPolicy) GetReturnShippingCostPayerOk() (*string, bool)`

GetReturnShippingCostPayerOk returns a tuple with the ReturnShippingCostPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnShippingCostPayer

`func (o *ReturnPolicy) SetReturnShippingCostPayer(v string)`

SetReturnShippingCostPayer sets ReturnShippingCostPayer field to given value.

### HasReturnShippingCostPayer

`func (o *ReturnPolicy) HasReturnShippingCostPayer() bool`

HasReturnShippingCostPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


