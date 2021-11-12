# SetReturnPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTypes** | Pointer to [**[]CategoryType**](CategoryType.md) | For return policies, this field always returns &lt;code&gt;ALL_EXCLUDING_MOTORS_VEHICLES&lt;/code&gt; (returns on motor vehicles are not processed through eBay flows.) &lt;br&gt;&lt;br&gt;&lt;b&gt;Default&lt;/b&gt;: &lt;code&gt;ALL_EXCLUDING_MOTORS_VEHICLES&lt;/code&gt; (for return policies only) | [optional] 
**Description** | Pointer to **string** | An optional seller-defined description of the return policy for internal use (this value is not displayed to end users). | [optional] 
**ExtendedHolidayReturnsOffered** | Pointer to **bool** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is ignored, it is neither read nor returned.&lt;/p&gt;  &lt;p&gt;If set to &lt;code&gt;true&lt;/code&gt;, the seller offers an &lt;em&gt;Extended Holiday Returns&lt;/em&gt; policy for their listings. &lt;p&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;IMPORTANT:&lt;/strong&gt; Extended Holiday Returns is a seasonally available feature that is offered on some eBay marketplaces. To see if the feature is enabled in any given year, check the eBay Seller Center &lt;a href&#x3D;\&quot;http://pages.ebay.com/sellerinformation/returns-on-eBay/\&quot;&gt;Returns on eBay&lt;/a&gt; page of before the holiday season begins.&lt;/p&gt; | [optional] 
**InternationalOverride** | Pointer to [**InternationalReturnOverrideType**](InternationalReturnOverrideType.md) |  | [optional] 
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which this return policy applies. If this value is not specified, value defaults to the seller&#39;s eBay registration site. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | A user-defined name for this return policy. Names must be unique for policies assigned to the same marketplace. &lt;br&gt;&lt;br&gt;&lt;b&gt;Max length&lt;/b&gt;: 64 | [optional] 
**RefundMethod** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value other than &lt;code&gt;MONEY_BACK&lt;/code&gt; will be treated as &lt;code&gt;MONEY_BACK&lt;/code&gt; (although for a period of time, eBay will store and return the legacy values to preserve backwards compatibility).&lt;/p&gt;  Indicates the method the seller uses to compensate the buyer for returned items. The return method specified applies only to &lt;a href&#x3D;\&quot;http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\&quot; target&#x3D;\&quot;_blank\&quot;&gt;remorse returns&lt;/a&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:RefundMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**RestockingFeePercentage** | Pointer to **string** | &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;span  style&#x3D;\&quot;color: #dd1e31;\&quot;&gt;&lt;b&gt;Important!&lt;/b&gt;&lt;/span&gt; This field has been deprecated as of version 1.2.0, released on May 31, 2018. Any value supplied in this field is ignored, it is neither read nor returned.&lt;/p&gt;  &lt;p&gt;Optionally set by the seller, the percentage charged if the seller charges buyers a a restocking fee when items are returned due to buyer remorse and/or a purchasing mistake. The total amount charged to the buyer is the cost of the item multiplied by the percentage indicated in this field. | [optional] 
**ReturnInstructions** | Pointer to **string** | This field contains the seller&#39;s detailed explanation for their return policy and is displayed in the Return Policy section of the View Item page. This field is valid in only the following marketplaces (the field is otherwise ignored): &lt;ul&gt; &lt;li&gt;Germany (DE)&lt;/li&gt; &lt;li&gt;Spain (ES)&lt;/li&gt; &lt;li&gt;France (FR)&lt;/li&gt; &lt;li&gt;Italy (IT)&lt;/li&gt; &lt;/ul&gt; | [optional] 
**ReturnMethod** | Pointer to **string** | This field indicates the method in which the seller handles non-money back return requests for &lt;a href&#x3D;\&quot;http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\&quot; target&#x3D;\&quot;_blank\&quot;&gt;remorse returns&lt;/a&gt;. This field is valid in only the US marketplace and the only valid value is &lt;code&gt;REPLACEMENT&lt;/code&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ReturnPeriod** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ReturnPolicyId** | Pointer to **string** | A unique eBay-assigned ID for a return policy. This ID is generated when the policy is created. | [optional] 
**ReturnsAccepted** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller accepts returns. If set to &lt;code&gt;false&lt;/code&gt;, this field indicates that the seller does not accept returns. | [optional] 
**ReturnShippingCostPayer** | Pointer to **string** | This field indicates who is responsible for paying for the shipping charges for returned items. The field can be set to either &lt;code&gt;BUYER&lt;/code&gt; or &lt;code&gt;SELLER&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;Depending on the return policy and specifics of the return, either the buyer or the seller can be responsible for the return shipping costs. Note that the seller is always responsible for return shipping costs for SNAD-related issues. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnShippingCostPayerEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | A list of warnings related to request. This field normally returns empty, which indicates the request did not generate any warnings. | [optional] 

## Methods

### NewSetReturnPolicyResponse

`func NewSetReturnPolicyResponse() *SetReturnPolicyResponse`

NewSetReturnPolicyResponse instantiates a new SetReturnPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetReturnPolicyResponseWithDefaults

`func NewSetReturnPolicyResponseWithDefaults() *SetReturnPolicyResponse`

NewSetReturnPolicyResponseWithDefaults instantiates a new SetReturnPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTypes

`func (o *SetReturnPolicyResponse) GetCategoryTypes() []CategoryType`

GetCategoryTypes returns the CategoryTypes field if non-nil, zero value otherwise.

### GetCategoryTypesOk

`func (o *SetReturnPolicyResponse) GetCategoryTypesOk() (*[]CategoryType, bool)`

GetCategoryTypesOk returns a tuple with the CategoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTypes

`func (o *SetReturnPolicyResponse) SetCategoryTypes(v []CategoryType)`

SetCategoryTypes sets CategoryTypes field to given value.

### HasCategoryTypes

`func (o *SetReturnPolicyResponse) HasCategoryTypes() bool`

HasCategoryTypes returns a boolean if a field has been set.

### GetDescription

`func (o *SetReturnPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetReturnPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetReturnPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetReturnPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtendedHolidayReturnsOffered

`func (o *SetReturnPolicyResponse) GetExtendedHolidayReturnsOffered() bool`

GetExtendedHolidayReturnsOffered returns the ExtendedHolidayReturnsOffered field if non-nil, zero value otherwise.

### GetExtendedHolidayReturnsOfferedOk

`func (o *SetReturnPolicyResponse) GetExtendedHolidayReturnsOfferedOk() (*bool, bool)`

GetExtendedHolidayReturnsOfferedOk returns a tuple with the ExtendedHolidayReturnsOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHolidayReturnsOffered

`func (o *SetReturnPolicyResponse) SetExtendedHolidayReturnsOffered(v bool)`

SetExtendedHolidayReturnsOffered sets ExtendedHolidayReturnsOffered field to given value.

### HasExtendedHolidayReturnsOffered

`func (o *SetReturnPolicyResponse) HasExtendedHolidayReturnsOffered() bool`

HasExtendedHolidayReturnsOffered returns a boolean if a field has been set.

### GetInternationalOverride

`func (o *SetReturnPolicyResponse) GetInternationalOverride() InternationalReturnOverrideType`

GetInternationalOverride returns the InternationalOverride field if non-nil, zero value otherwise.

### GetInternationalOverrideOk

`func (o *SetReturnPolicyResponse) GetInternationalOverrideOk() (*InternationalReturnOverrideType, bool)`

GetInternationalOverrideOk returns a tuple with the InternationalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOverride

`func (o *SetReturnPolicyResponse) SetInternationalOverride(v InternationalReturnOverrideType)`

SetInternationalOverride sets InternationalOverride field to given value.

### HasInternationalOverride

`func (o *SetReturnPolicyResponse) HasInternationalOverride() bool`

HasInternationalOverride returns a boolean if a field has been set.

### GetMarketplaceId

`func (o *SetReturnPolicyResponse) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *SetReturnPolicyResponse) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *SetReturnPolicyResponse) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *SetReturnPolicyResponse) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetName

`func (o *SetReturnPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetReturnPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetReturnPolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SetReturnPolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefundMethod

`func (o *SetReturnPolicyResponse) GetRefundMethod() string`

GetRefundMethod returns the RefundMethod field if non-nil, zero value otherwise.

### GetRefundMethodOk

`func (o *SetReturnPolicyResponse) GetRefundMethodOk() (*string, bool)`

GetRefundMethodOk returns a tuple with the RefundMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundMethod

`func (o *SetReturnPolicyResponse) SetRefundMethod(v string)`

SetRefundMethod sets RefundMethod field to given value.

### HasRefundMethod

`func (o *SetReturnPolicyResponse) HasRefundMethod() bool`

HasRefundMethod returns a boolean if a field has been set.

### GetRestockingFeePercentage

`func (o *SetReturnPolicyResponse) GetRestockingFeePercentage() string`

GetRestockingFeePercentage returns the RestockingFeePercentage field if non-nil, zero value otherwise.

### GetRestockingFeePercentageOk

`func (o *SetReturnPolicyResponse) GetRestockingFeePercentageOk() (*string, bool)`

GetRestockingFeePercentageOk returns a tuple with the RestockingFeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockingFeePercentage

`func (o *SetReturnPolicyResponse) SetRestockingFeePercentage(v string)`

SetRestockingFeePercentage sets RestockingFeePercentage field to given value.

### HasRestockingFeePercentage

`func (o *SetReturnPolicyResponse) HasRestockingFeePercentage() bool`

HasRestockingFeePercentage returns a boolean if a field has been set.

### GetReturnInstructions

`func (o *SetReturnPolicyResponse) GetReturnInstructions() string`

GetReturnInstructions returns the ReturnInstructions field if non-nil, zero value otherwise.

### GetReturnInstructionsOk

`func (o *SetReturnPolicyResponse) GetReturnInstructionsOk() (*string, bool)`

GetReturnInstructionsOk returns a tuple with the ReturnInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnInstructions

`func (o *SetReturnPolicyResponse) SetReturnInstructions(v string)`

SetReturnInstructions sets ReturnInstructions field to given value.

### HasReturnInstructions

`func (o *SetReturnPolicyResponse) HasReturnInstructions() bool`

HasReturnInstructions returns a boolean if a field has been set.

### GetReturnMethod

`func (o *SetReturnPolicyResponse) GetReturnMethod() string`

GetReturnMethod returns the ReturnMethod field if non-nil, zero value otherwise.

### GetReturnMethodOk

`func (o *SetReturnPolicyResponse) GetReturnMethodOk() (*string, bool)`

GetReturnMethodOk returns a tuple with the ReturnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMethod

`func (o *SetReturnPolicyResponse) SetReturnMethod(v string)`

SetReturnMethod sets ReturnMethod field to given value.

### HasReturnMethod

`func (o *SetReturnPolicyResponse) HasReturnMethod() bool`

HasReturnMethod returns a boolean if a field has been set.

### GetReturnPeriod

`func (o *SetReturnPolicyResponse) GetReturnPeriod() TimeDuration`

GetReturnPeriod returns the ReturnPeriod field if non-nil, zero value otherwise.

### GetReturnPeriodOk

`func (o *SetReturnPolicyResponse) GetReturnPeriodOk() (*TimeDuration, bool)`

GetReturnPeriodOk returns a tuple with the ReturnPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPeriod

`func (o *SetReturnPolicyResponse) SetReturnPeriod(v TimeDuration)`

SetReturnPeriod sets ReturnPeriod field to given value.

### HasReturnPeriod

`func (o *SetReturnPolicyResponse) HasReturnPeriod() bool`

HasReturnPeriod returns a boolean if a field has been set.

### GetReturnPolicyId

`func (o *SetReturnPolicyResponse) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *SetReturnPolicyResponse) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *SetReturnPolicyResponse) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *SetReturnPolicyResponse) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### GetReturnsAccepted

`func (o *SetReturnPolicyResponse) GetReturnsAccepted() bool`

GetReturnsAccepted returns the ReturnsAccepted field if non-nil, zero value otherwise.

### GetReturnsAcceptedOk

`func (o *SetReturnPolicyResponse) GetReturnsAcceptedOk() (*bool, bool)`

GetReturnsAcceptedOk returns a tuple with the ReturnsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsAccepted

`func (o *SetReturnPolicyResponse) SetReturnsAccepted(v bool)`

SetReturnsAccepted sets ReturnsAccepted field to given value.

### HasReturnsAccepted

`func (o *SetReturnPolicyResponse) HasReturnsAccepted() bool`

HasReturnsAccepted returns a boolean if a field has been set.

### GetReturnShippingCostPayer

`func (o *SetReturnPolicyResponse) GetReturnShippingCostPayer() string`

GetReturnShippingCostPayer returns the ReturnShippingCostPayer field if non-nil, zero value otherwise.

### GetReturnShippingCostPayerOk

`func (o *SetReturnPolicyResponse) GetReturnShippingCostPayerOk() (*string, bool)`

GetReturnShippingCostPayerOk returns a tuple with the ReturnShippingCostPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnShippingCostPayer

`func (o *SetReturnPolicyResponse) SetReturnShippingCostPayer(v string)`

SetReturnShippingCostPayer sets ReturnShippingCostPayer field to given value.

### HasReturnShippingCostPayer

`func (o *SetReturnPolicyResponse) HasReturnShippingCostPayer() bool`

HasReturnShippingCostPayer returns a boolean if a field has been set.

### GetWarnings

`func (o *SetReturnPolicyResponse) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SetReturnPolicyResponse) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SetReturnPolicyResponse) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SetReturnPolicyResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


