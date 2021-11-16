# RegionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionExcluded** | Pointer to [**[]Region**](Region.md) | A list of one or more &lt;b&gt;regionsName&lt;/b&gt; fields that specify the areas to where a seller does not ship. Populate &lt;b&gt;regionExcluded&lt;/b&gt; in only the top-level &lt;b&gt;shipToLocations&lt;/b&gt; container (do not populate this field within the &lt;b&gt;shippingOptions&lt;/b&gt; container). &lt;p&gt;Normally a seller ships to as many areas as possible using both DOMESTIC and INTERNATIONAL shipping options and they don&#39;t have a need to exclude any regions from their ship-to locations. With this, there&#39;s no reason to set &lt;b&gt;regionExclude&lt;/b&gt; fields. However, it makes sense to set the &lt;b&gt;regionExcluded&lt;/b&gt; field when a seller wants to exclude a small area that&#39;s located within a larger area they service. For example, suppose a seller indicates they ship &#39;Worldwide&#39;, but for some reason must exclude a specific country, or world region, from the larger world area they ship to. &lt;br&gt;&lt;br&gt;To retrieve the regions you can specify in the associated &lt;b&gt;regionName&lt;/b&gt; field, call &lt;b&gt;GeteBayDetails&lt;/b&gt; with &lt;b&gt;DetailName&lt;/b&gt; set to &lt;code&gt;ExcludeShippingLocationDetails&lt;/code&gt;, then review the &lt;b&gt;Location&lt;/b&gt; fields in the response for the strings that you can specify &lt;b&gt;regionExcluded.regionName&lt;/b&gt;.  &lt;br&gt;&lt;br&gt;Note that if a buyer&#39;s primary ship-to location is a region that a seller has excluded in their fulfillment policy (or if the buyer does not have a primary ship-to location), they will receive an error message if they attempt to buy or place a bid on an item that uses that fulfillment policy. &lt;br&gt;&lt;br&gt;For details on setting this field, see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/ht_shipping-exclude-regions.html\&quot;&gt;Excluding specific regions from included shipping areas&lt;/a&gt;. | [optional] 
**RegionIncluded** | Pointer to [**[]Region**](Region.md) | A list of one or more &lt;b&gt;regionsName&lt;/b&gt; fields that specify the areas to where a seller ships. &lt;br&gt;&lt;br&gt;&lt;b&gt;Important:&lt;/b&gt; Populate this field only when the parent &lt;b&gt;shipToLocations&lt;/b&gt; object is located within a &lt;b&gt;shippingOptions&lt;/b&gt; container (that is, the parent &lt;b&gt;shipTolocations&lt;/b&gt; object must not be the one at the top-level of the policy). Also, this field needs to be populated only when the associated &lt;b&gt;shippingOptions&lt;/b&gt; container has &lt;b&gt;optionType&lt;/b&gt; set to &lt;code&gt;INTERNATIONAL&lt;/code&gt;. &lt;br&gt;&lt;br&gt;Withing an international shipping option, set this value to &lt;code&gt;Worldwide&lt;/code&gt; to indicate the seller ships to all world regions. If needed, use the &lt;b&gt;regionExcluded&lt;/b&gt; field to exclude any regions in the world to where the seller does not ship. &lt;br&gt;&lt;br&gt;Each eBay marketplace supports its own set of allowable shipping locations. Obtain the valid &#39;Ship-To Locations&#39; for a marketplace by calling &lt;b&gt;GeteBayDetails&lt;/b&gt; with &lt;b&gt;DetailName&lt;/b&gt; set to &lt;code&gt;ShippingLocationDetails&lt;/code&gt;, then review the &lt;b&gt;ShippingLocation&lt;/b&gt; fields in the response for the strings that you can specify in the &lt;b&gt;regionIncluded.regionName&lt;/b&gt; field. &lt;br&gt;&lt;br&gt;For DOMESTIC shipping options, eBay automatically uses the seller&#39;s listing country as the default &lt;b&gt;regionIncluded&lt;/b&gt; country. For details on setting this field, see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/ht_shipping-worldwide.html\&quot;&gt;How to set up worldwide shipping&lt;/a&gt;.  &lt;br&gt;&lt;br&gt;This field is always returned in the shipping policy response.  &lt;br&gt;&lt;br&gt;&lt;i&gt;Required if&lt;/i&gt; &lt;b&gt;optionType&lt;/b&gt; set to &lt;code&gt;INTERNATIONAL&lt;/code&gt;. | [optional] 

## Methods

### NewRegionSet

`func NewRegionSet() *RegionSet`

NewRegionSet instantiates a new RegionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionSetWithDefaults

`func NewRegionSetWithDefaults() *RegionSet`

NewRegionSetWithDefaults instantiates a new RegionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionExcluded

`func (o *RegionSet) GetRegionExcluded() []Region`

GetRegionExcluded returns the RegionExcluded field if non-nil, zero value otherwise.

### GetRegionExcludedOk

`func (o *RegionSet) GetRegionExcludedOk() (*[]Region, bool)`

GetRegionExcludedOk returns a tuple with the RegionExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionExcluded

`func (o *RegionSet) SetRegionExcluded(v []Region)`

SetRegionExcluded sets RegionExcluded field to given value.

### HasRegionExcluded

`func (o *RegionSet) HasRegionExcluded() bool`

HasRegionExcluded returns a boolean if a field has been set.

### GetRegionIncluded

`func (o *RegionSet) GetRegionIncluded() []Region`

GetRegionIncluded returns the RegionIncluded field if non-nil, zero value otherwise.

### GetRegionIncludedOk

`func (o *RegionSet) GetRegionIncludedOk() (*[]Region, bool)`

GetRegionIncludedOk returns a tuple with the RegionIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIncluded

`func (o *RegionSet) SetRegionIncluded(v []Region)`

SetRegionIncluded sets RegionIncluded field to given value.

### HasRegionIncluded

`func (o *RegionSet) HasRegionIncluded() bool`

HasRegionIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


