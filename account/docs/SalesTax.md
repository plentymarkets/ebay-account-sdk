# SalesTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The country code identifying the country to which this tax rate applies. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**SalesTaxJurisdictionId** | Pointer to **string** | A unique ID that identifies the sales tax jurisdiction to which the tax rate applies (for example a state within the United States). | [optional] 
**SalesTaxPercentage** | Pointer to **string** | The sales tax rate (as a percentage of the sale) applied to sales transactions made in this country and sales tax jurisdiction. | [optional] 
**ShippingAndHandlingTaxed** | Pointer to **bool** | If set to &lt;code&gt;true&lt;/code&gt;, shipping and handling charges are taxed. | [optional] 

## Methods

### NewSalesTax

`func NewSalesTax() *SalesTax`

NewSalesTax instantiates a new SalesTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesTaxWithDefaults

`func NewSalesTaxWithDefaults() *SalesTax`

NewSalesTaxWithDefaults instantiates a new SalesTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *SalesTax) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SalesTax) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SalesTax) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SalesTax) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetSalesTaxJurisdictionId

`func (o *SalesTax) GetSalesTaxJurisdictionId() string`

GetSalesTaxJurisdictionId returns the SalesTaxJurisdictionId field if non-nil, zero value otherwise.

### GetSalesTaxJurisdictionIdOk

`func (o *SalesTax) GetSalesTaxJurisdictionIdOk() (*string, bool)`

GetSalesTaxJurisdictionIdOk returns a tuple with the SalesTaxJurisdictionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxJurisdictionId

`func (o *SalesTax) SetSalesTaxJurisdictionId(v string)`

SetSalesTaxJurisdictionId sets SalesTaxJurisdictionId field to given value.

### HasSalesTaxJurisdictionId

`func (o *SalesTax) HasSalesTaxJurisdictionId() bool`

HasSalesTaxJurisdictionId returns a boolean if a field has been set.

### GetSalesTaxPercentage

`func (o *SalesTax) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *SalesTax) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *SalesTax) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.

### HasSalesTaxPercentage

`func (o *SalesTax) HasSalesTaxPercentage() bool`

HasSalesTaxPercentage returns a boolean if a field has been set.

### GetShippingAndHandlingTaxed

`func (o *SalesTax) GetShippingAndHandlingTaxed() bool`

GetShippingAndHandlingTaxed returns the ShippingAndHandlingTaxed field if non-nil, zero value otherwise.

### GetShippingAndHandlingTaxedOk

`func (o *SalesTax) GetShippingAndHandlingTaxedOk() (*bool, bool)`

GetShippingAndHandlingTaxedOk returns a tuple with the ShippingAndHandlingTaxed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAndHandlingTaxed

`func (o *SalesTax) SetShippingAndHandlingTaxed(v bool)`

SetShippingAndHandlingTaxed sets ShippingAndHandlingTaxed field to given value.

### HasShippingAndHandlingTaxed

`func (o *SalesTax) HasShippingAndHandlingTaxed() bool`

HasShippingAndHandlingTaxed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


