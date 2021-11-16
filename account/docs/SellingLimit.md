# SellingLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Quantity** | Pointer to **int32** | The maximum quantity of items that can be listed by the seller per calendar month. Note that for a listing with variations, all of the items listed in the variation count as individual items. | [optional] 

## Methods

### NewSellingLimit

`func NewSellingLimit() *SellingLimit`

NewSellingLimit instantiates a new SellingLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellingLimitWithDefaults

`func NewSellingLimitWithDefaults() *SellingLimit`

NewSellingLimitWithDefaults instantiates a new SellingLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SellingLimit) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SellingLimit) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SellingLimit) SetAmount(v Amount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SellingLimit) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *SellingLimit) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SellingLimit) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SellingLimit) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SellingLimit) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


