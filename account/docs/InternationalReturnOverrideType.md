# InternationalReturnOverrideType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnMethod** | Pointer to **string** | Valid in the US marketplace only, this optional field indicates additional services (other than money-back) that sellers can offer buyers for &lt;a href&#x3D;\&quot;http://developer.ebay.com/DevZone/guides/features-guide/default.html#Development/Post-Order-Returns.html#return-reasons\&quot; target&#x3D;\&quot;_blank\&quot;&gt;remorse returns&lt;/a&gt;.  &lt;br&gt;&lt;br&gt;As of version 1.2.0, the only accepted value for this field is &lt;code&gt;REPLACEMENT&lt;/code&gt;. This field is valid in only the US marketplace, any supplied value is ignored in other marketplaces. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnMethodEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ReturnPeriod** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**ReturnsAccepted** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller allows international returns. If set to &lt;code&gt;false&lt;/code&gt;, the seller does not accept international returns.  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if &lt;/i&gt; the seller wants to set an international return policy that differs from their domestic return policy. | [optional] 
**ReturnShippingCostPayer** | Pointer to **string** | This field indicates who is responsible for paying for the shipping charges for returned items. The field can be set to either &lt;code&gt;BUYER&lt;/code&gt; or &lt;code&gt;SELLER&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;Depending on the return policy and specifics of the return, either the buyer or the seller can be responsible for the return shipping costs. Note that the seller is always responsible for return shipping costs for SNAD-related issues.  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if &lt;/i&gt; the &lt;b&gt;internationalOverride.returnsAccepted&lt;/b&gt; field is set to &lt;code&gt;true&lt;/code&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ReturnShippingCostPayerEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewInternationalReturnOverrideType

`func NewInternationalReturnOverrideType() *InternationalReturnOverrideType`

NewInternationalReturnOverrideType instantiates a new InternationalReturnOverrideType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalReturnOverrideTypeWithDefaults

`func NewInternationalReturnOverrideTypeWithDefaults() *InternationalReturnOverrideType`

NewInternationalReturnOverrideTypeWithDefaults instantiates a new InternationalReturnOverrideType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnMethod

`func (o *InternationalReturnOverrideType) GetReturnMethod() string`

GetReturnMethod returns the ReturnMethod field if non-nil, zero value otherwise.

### GetReturnMethodOk

`func (o *InternationalReturnOverrideType) GetReturnMethodOk() (*string, bool)`

GetReturnMethodOk returns a tuple with the ReturnMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMethod

`func (o *InternationalReturnOverrideType) SetReturnMethod(v string)`

SetReturnMethod sets ReturnMethod field to given value.

### HasReturnMethod

`func (o *InternationalReturnOverrideType) HasReturnMethod() bool`

HasReturnMethod returns a boolean if a field has been set.

### GetReturnPeriod

`func (o *InternationalReturnOverrideType) GetReturnPeriod() TimeDuration`

GetReturnPeriod returns the ReturnPeriod field if non-nil, zero value otherwise.

### GetReturnPeriodOk

`func (o *InternationalReturnOverrideType) GetReturnPeriodOk() (*TimeDuration, bool)`

GetReturnPeriodOk returns a tuple with the ReturnPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPeriod

`func (o *InternationalReturnOverrideType) SetReturnPeriod(v TimeDuration)`

SetReturnPeriod sets ReturnPeriod field to given value.

### HasReturnPeriod

`func (o *InternationalReturnOverrideType) HasReturnPeriod() bool`

HasReturnPeriod returns a boolean if a field has been set.

### GetReturnsAccepted

`func (o *InternationalReturnOverrideType) GetReturnsAccepted() bool`

GetReturnsAccepted returns the ReturnsAccepted field if non-nil, zero value otherwise.

### GetReturnsAcceptedOk

`func (o *InternationalReturnOverrideType) GetReturnsAcceptedOk() (*bool, bool)`

GetReturnsAcceptedOk returns a tuple with the ReturnsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsAccepted

`func (o *InternationalReturnOverrideType) SetReturnsAccepted(v bool)`

SetReturnsAccepted sets ReturnsAccepted field to given value.

### HasReturnsAccepted

`func (o *InternationalReturnOverrideType) HasReturnsAccepted() bool`

HasReturnsAccepted returns a boolean if a field has been set.

### GetReturnShippingCostPayer

`func (o *InternationalReturnOverrideType) GetReturnShippingCostPayer() string`

GetReturnShippingCostPayer returns the ReturnShippingCostPayer field if non-nil, zero value otherwise.

### GetReturnShippingCostPayerOk

`func (o *InternationalReturnOverrideType) GetReturnShippingCostPayerOk() (*string, bool)`

GetReturnShippingCostPayerOk returns a tuple with the ReturnShippingCostPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnShippingCostPayer

`func (o *InternationalReturnOverrideType) SetReturnShippingCostPayer(v string)`

SetReturnShippingCostPayer sets ReturnShippingCostPayer field to given value.

### HasReturnShippingCostPayer

`func (o *InternationalReturnOverrideType) HasReturnShippingCostPayer() bool`

HasReturnShippingCostPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


