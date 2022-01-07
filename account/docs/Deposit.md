# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**DueIn** | Pointer to [**TimeDuration**](TimeDuration.md) |  | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | For deposits (which are applicable to only motor listings), the &lt;b&gt;paymentMethodType&lt;/b&gt; must be set to &#39;&lt;code&gt;PAYPAL&lt;/code&gt;&#39; and you must also populate the fields in the &lt;b&gt;recipientAccountReference&lt;/b&gt; object.  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if&lt;/i&gt; your motor vehicles listing requires a deposit. | [optional] 

## Methods

### NewDeposit

`func NewDeposit() *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Deposit) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Deposit) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Deposit) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Deposit) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDueIn

`func (o *Deposit) GetDueIn() TimeDuration`

GetDueIn returns the DueIn field if non-nil, zero value otherwise.

### GetDueInOk

`func (o *Deposit) GetDueInOk() (*TimeDuration, bool)`

GetDueInOk returns a tuple with the DueIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueIn

`func (o *Deposit) SetDueIn(v TimeDuration)`

SetDueIn sets DueIn field to given value.

### HasDueIn

`func (o *Deposit) HasDueIn() bool`

HasDueIn returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *Deposit) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *Deposit) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *Deposit) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *Deposit) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


