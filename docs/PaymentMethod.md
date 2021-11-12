# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brands** | Pointer to **[]string** | It&#39;s important to note that the credit card brands Visa and MasterCard must both be listed if either one is listed, as is shown in the following code fragment:  &lt;br&gt;&lt;br&gt;&lt;code&gt;\&quot;paymentMethods\&quot;: [{ \&quot;brands\&quot;: [VISA, MASTERCARD] }] ...&lt;/code&gt; &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; Different eBay marketplaces may or may not support this field. &lt;br&gt;&lt;br&gt;Use the Trading API &lt;b&gt;GetCategoryFeatures&lt;/b&gt; call with &lt;b&gt;FeatureID&lt;/b&gt; set to &lt;code&gt;PaymentMethods&lt;/code&gt; and &lt;b&gt;DetailLevel&lt;/b&gt; set to &lt;code&gt;ReturnAll&lt;/code&gt; to see what credit card brands different marketplaces support. If the &lt;b&gt;GetCategoryFeatures&lt;/b&gt; call returns details on credit card brands for the categories in which you sell, then you can use this field to list the credit card brands the seller accepts. If, on the other hand, &lt;b&gt;GetCategoryFeatures&lt;/b&gt; does not enumerate credit card brands for your target site (for example, if it returns &lt;b&gt;PaymentMethod&lt;/b&gt; set to &lt;code&gt;CCAccepted&lt;/code&gt;), then you cannot enumerate specific credit card brands with this field for that marketplace.&lt;/p&gt; &lt;p&gt;&lt;i&gt;Required if &lt;/i&gt; &lt;b&gt;paymentMethodType&lt;/b&gt; is set to &lt;code&gt;CREDIT_CARD&lt;/code&gt;.  &lt;br&gt;&lt;br&gt;A list of credit card brands accepted by the seller.&lt;/p&gt; | [optional] 
**PaymentMethodType** | Pointer to **string** | The payment method, selected from the supported payment method types.  &lt;p&gt;Use &lt;b&gt;GetCategoryFeatures&lt;/b&gt; in the Trading API to retrieve the payment methods allowed for a category on a specific marketplace, as well as the default payment method for that marketplace (review the &lt;b&gt;SiteDefaults.PaymentMethod&lt;/b&gt; field). For example, the response from &lt;b&gt;GetCategoryFeatures&lt;/b&gt; shows that on the US marketplace, most categories allow only electronic payments via credit cards, PayPal, and the like. Also, note that &lt;b&gt;GeteBayDetails&lt;/b&gt; does not return payment method information.  &lt;p class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; If you create item listings using the Inventory API, you must set this field to &lt;code&gt;PAYPAL&lt;/code&gt; (currently, the Inventory API supports only fixed-prince GTC items where the only supported &lt;b&gt;paymentMethod&lt;/b&gt; is PayPal).&lt;/p&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:PaymentMethodTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**RecipientAccountReference** | Pointer to [**RecipientAccountReference**](RecipientAccountReference.md) |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrands

`func (o *PaymentMethod) GetBrands() []string`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *PaymentMethod) GetBrandsOk() (*[]string, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *PaymentMethod) SetBrands(v []string)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *PaymentMethod) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *PaymentMethod) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *PaymentMethod) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *PaymentMethod) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *PaymentMethod) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetRecipientAccountReference

`func (o *PaymentMethod) GetRecipientAccountReference() RecipientAccountReference`

GetRecipientAccountReference returns the RecipientAccountReference field if non-nil, zero value otherwise.

### GetRecipientAccountReferenceOk

`func (o *PaymentMethod) GetRecipientAccountReferenceOk() (*RecipientAccountReference, bool)`

GetRecipientAccountReferenceOk returns a tuple with the RecipientAccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAccountReference

`func (o *PaymentMethod) SetRecipientAccountReference(v RecipientAccountReference)`

SetRecipientAccountReference sets RecipientAccountReference field to given value.

### HasRecipientAccountReference

`func (o *PaymentMethod) HasRecipientAccountReference() bool`

HasRecipientAccountReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


