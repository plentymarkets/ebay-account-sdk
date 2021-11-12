# PaymentsProgramResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceId** | Pointer to **string** | The ID of the eBay marketplace to which the payment program applies. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:MarketplaceIdEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PaymentsProgramType** | Pointer to **string** | This parameter specifies the payment program whose status is returned by the call.  &lt;br&gt;&lt;br&gt;Currently the only supported payments program is &lt;code&gt;EBAY_PAYMENTS&lt;/code&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:PaymentsProgramType&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Status** | Pointer to **string** | The enumeration value returned in this field indicates whether or not the seller is enabled for the payments program. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:PaymentsProgramStatus&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**WasPreviouslyOptedIn** | Pointer to **bool** | If returned as &lt;code&gt;true&lt;/code&gt;, the seller was at one point opted-in to the associated payment program, but they later opted out of the program. A value of &lt;code&gt;false&lt;/code&gt; indicates the seller never opted-in to the program or if they did opt-in to the program, they never opted-out of it.  &lt;br&gt;&lt;br&gt;It&#39;s important to note that the setting of this field does not indicate the seller&#39;s current status regarding the payment program. It is possible for this field to return &lt;code&gt;true&lt;/code&gt; while the &lt;b&gt;status&lt;/b&gt; field returns &lt;code&gt;OPTED_IN&lt;/code&gt;. | [optional] 

## Methods

### NewPaymentsProgramResponse

`func NewPaymentsProgramResponse() *PaymentsProgramResponse`

NewPaymentsProgramResponse instantiates a new PaymentsProgramResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsProgramResponseWithDefaults

`func NewPaymentsProgramResponseWithDefaults() *PaymentsProgramResponse`

NewPaymentsProgramResponseWithDefaults instantiates a new PaymentsProgramResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceId

`func (o *PaymentsProgramResponse) GetMarketplaceId() string`

GetMarketplaceId returns the MarketplaceId field if non-nil, zero value otherwise.

### GetMarketplaceIdOk

`func (o *PaymentsProgramResponse) GetMarketplaceIdOk() (*string, bool)`

GetMarketplaceIdOk returns a tuple with the MarketplaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceId

`func (o *PaymentsProgramResponse) SetMarketplaceId(v string)`

SetMarketplaceId sets MarketplaceId field to given value.

### HasMarketplaceId

`func (o *PaymentsProgramResponse) HasMarketplaceId() bool`

HasMarketplaceId returns a boolean if a field has been set.

### GetPaymentsProgramType

`func (o *PaymentsProgramResponse) GetPaymentsProgramType() string`

GetPaymentsProgramType returns the PaymentsProgramType field if non-nil, zero value otherwise.

### GetPaymentsProgramTypeOk

`func (o *PaymentsProgramResponse) GetPaymentsProgramTypeOk() (*string, bool)`

GetPaymentsProgramTypeOk returns a tuple with the PaymentsProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsProgramType

`func (o *PaymentsProgramResponse) SetPaymentsProgramType(v string)`

SetPaymentsProgramType sets PaymentsProgramType field to given value.

### HasPaymentsProgramType

`func (o *PaymentsProgramResponse) HasPaymentsProgramType() bool`

HasPaymentsProgramType returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentsProgramResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentsProgramResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentsProgramResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentsProgramResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWasPreviouslyOptedIn

`func (o *PaymentsProgramResponse) GetWasPreviouslyOptedIn() bool`

GetWasPreviouslyOptedIn returns the WasPreviouslyOptedIn field if non-nil, zero value otherwise.

### GetWasPreviouslyOptedInOk

`func (o *PaymentsProgramResponse) GetWasPreviouslyOptedInOk() (*bool, bool)`

GetWasPreviouslyOptedInOk returns a tuple with the WasPreviouslyOptedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasPreviouslyOptedIn

`func (o *PaymentsProgramResponse) SetWasPreviouslyOptedIn(v bool)`

SetWasPreviouslyOptedIn sets WasPreviouslyOptedIn field to given value.

### HasWasPreviouslyOptedIn

`func (o *PaymentsProgramResponse) HasWasPreviouslyOptedIn() bool`

HasWasPreviouslyOptedIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


