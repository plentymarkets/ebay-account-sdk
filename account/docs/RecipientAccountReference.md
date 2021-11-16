# RecipientAccountReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** | Contains the PayPal email address of the recipient (seller) if &lt;b&gt;referenceType&lt;/b&gt; is set to &lt;code&gt;PAYPAL_EMAIL&lt;/code&gt;. | [optional] 
**ReferenceType** | Pointer to **string** | This field&#39;s value indicates the type of account where payment will be received. Currently, the only supported value is PAYPAL_EMAIL. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:RecipientAccountReferenceTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewRecipientAccountReference

`func NewRecipientAccountReference() *RecipientAccountReference`

NewRecipientAccountReference instantiates a new RecipientAccountReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientAccountReferenceWithDefaults

`func NewRecipientAccountReferenceWithDefaults() *RecipientAccountReference`

NewRecipientAccountReferenceWithDefaults instantiates a new RecipientAccountReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *RecipientAccountReference) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *RecipientAccountReference) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *RecipientAccountReference) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *RecipientAccountReference) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceType

`func (o *RecipientAccountReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *RecipientAccountReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *RecipientAccountReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *RecipientAccountReference) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


