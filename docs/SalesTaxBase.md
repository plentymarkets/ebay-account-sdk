# SalesTaxBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesTaxPercentage** | Pointer to **string** | The sales tax rate, as a percentage of the sale. | [optional] 
**ShippingAndHandlingTaxed** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, shipping and handling charges are taxed. | [optional] 

## Methods

### NewSalesTaxBase

`func NewSalesTaxBase() *SalesTaxBase`

NewSalesTaxBase instantiates a new SalesTaxBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesTaxBaseWithDefaults

`func NewSalesTaxBaseWithDefaults() *SalesTaxBase`

NewSalesTaxBaseWithDefaults instantiates a new SalesTaxBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesTaxPercentage

`func (o *SalesTaxBase) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *SalesTaxBase) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *SalesTaxBase) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.

### HasSalesTaxPercentage

`func (o *SalesTaxBase) HasSalesTaxPercentage() bool`

HasSalesTaxPercentage returns a boolean if a field has been set.

### GetShippingAndHandlingTaxed

`func (o *SalesTaxBase) GetShippingAndHandlingTaxed() bool`

GetShippingAndHandlingTaxed returns the ShippingAndHandlingTaxed field if non-nil, zero value otherwise.

### GetShippingAndHandlingTaxedOk

`func (o *SalesTaxBase) GetShippingAndHandlingTaxedOk() (*bool, bool)`

GetShippingAndHandlingTaxedOk returns a tuple with the ShippingAndHandlingTaxed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAndHandlingTaxed

`func (o *SalesTaxBase) SetShippingAndHandlingTaxed(v bool)`

SetShippingAndHandlingTaxed sets ShippingAndHandlingTaxed field to given value.

### HasShippingAndHandlingTaxed

`func (o *SalesTaxBase) HasShippingAndHandlingTaxed() bool`

HasShippingAndHandlingTaxed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


