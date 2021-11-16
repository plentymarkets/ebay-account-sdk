# SellingPrivileges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerRegistrationCompleted** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, the seller&#39;s registration is completed. | [optional] 
**SellingLimit** | Pointer to [**SellingLimit**](SellingLimit.md) |  | [optional] 

## Methods

### NewSellingPrivileges

`func NewSellingPrivileges() *SellingPrivileges`

NewSellingPrivileges instantiates a new SellingPrivileges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellingPrivilegesWithDefaults

`func NewSellingPrivilegesWithDefaults() *SellingPrivileges`

NewSellingPrivilegesWithDefaults instantiates a new SellingPrivileges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellerRegistrationCompleted

`func (o *SellingPrivileges) GetSellerRegistrationCompleted() bool`

GetSellerRegistrationCompleted returns the SellerRegistrationCompleted field if non-nil, zero value otherwise.

### GetSellerRegistrationCompletedOk

`func (o *SellingPrivileges) GetSellerRegistrationCompletedOk() (*bool, bool)`

GetSellerRegistrationCompletedOk returns a tuple with the SellerRegistrationCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRegistrationCompleted

`func (o *SellingPrivileges) SetSellerRegistrationCompleted(v bool)`

SetSellerRegistrationCompleted sets SellerRegistrationCompleted field to given value.

### HasSellerRegistrationCompleted

`func (o *SellingPrivileges) HasSellerRegistrationCompleted() bool`

HasSellerRegistrationCompleted returns a boolean if a field has been set.

### GetSellingLimit

`func (o *SellingPrivileges) GetSellingLimit() SellingLimit`

GetSellingLimit returns the SellingLimit field if non-nil, zero value otherwise.

### GetSellingLimitOk

`func (o *SellingPrivileges) GetSellingLimitOk() (*SellingLimit, bool)`

GetSellingLimitOk returns a tuple with the SellingLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingLimit

`func (o *SellingPrivileges) SetSellingLimit(v SellingLimit)`

SetSellingLimit sets SellingLimit field to given value.

### HasSellingLimit

`func (o *SellingPrivileges) HasSellingLimit() bool`

HasSellingLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


