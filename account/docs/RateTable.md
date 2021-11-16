# RateTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | A two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; title&#x3D;\&quot;https://www.iso.org\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; country code representing the eBay marketplace where an item is listed. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Locality** | Pointer to **string** | The region covered by the shipping rate table, either DOMESTIC or INTERNATIONAL. &lt;br&gt;&lt;br&gt;DOMESTIC indicates that the shipping rate table applies to regions within the country where an item is listed (the &lt;i&gt;source&lt;/i&gt; country) while INTERNATIONAL indicates that the shipping rate table applies to regions outside of the country where an item is listed. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ShippingOptionTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | The user-defined name for the shipping rate table. Sellers can access Seller Hub (or &lt;b&gt;My eBay &gt; Account &gt; Site Preferences &gt; Shipping preferences&lt;/b&gt;) to create and assign names to their shipping rate tables. | [optional] 
**RateTableId** | Pointer to **string** | A unique eBay-assigned ID for a seller&#39;s shipping rate table. Call &lt;b&gt;getRateTables&lt;/b&gt; to retrieve the seller&#39;s current rate table IDs. | [optional] 

## Methods

### NewRateTable

`func NewRateTable() *RateTable`

NewRateTable instantiates a new RateTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateTableWithDefaults

`func NewRateTableWithDefaults() *RateTable`

NewRateTableWithDefaults instantiates a new RateTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *RateTable) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RateTable) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RateTable) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *RateTable) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocality

`func (o *RateTable) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *RateTable) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *RateTable) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *RateTable) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetName

`func (o *RateTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RateTable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRateTableId

`func (o *RateTable) GetRateTableId() string`

GetRateTableId returns the RateTableId field if non-nil, zero value otherwise.

### GetRateTableIdOk

`func (o *RateTable) GetRateTableIdOk() (*string, bool)`

GetRateTableIdOk returns a tuple with the RateTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTableId

`func (o *RateTable) SetRateTableId(v string)`

SetRateTableId sets RateTableId field to given value.

### HasRateTableId

`func (o *RateTable) HasRateTableId() bool`

HasRateTableId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


