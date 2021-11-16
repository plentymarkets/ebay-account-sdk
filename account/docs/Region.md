# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionName** | Pointer to **string** | A string that indicates the name of a region, as defined by eBay. A \&quot;region\&quot; can be either a &#39;world region&#39; (e.g., the \&quot;Middle East\&quot; or \&quot;Southeast Asia\&quot;) or a country, as represented with a two-letter country code. Use &lt;b&gt;GeteBayDetails&lt;/b&gt; to get the values accepted by this field. &lt;p&gt;The values that you&#39;re allowed to use for a specific &lt;b&gt;regionName&lt;/b&gt; field depend on the context in which you are setting the value. For details on how to set the values for this field, see &lt;a href&#x3D;\&quot;/api-docs/sell/static/seller-accounts/ht_shipping-worldwide.html#shipToLocations\&quot;&gt;The shipToLocations container&lt;/a&gt;. | [optional] 
**RegionType** | Pointer to **string** | Reserved for future use. &lt;!--The region&#39;s type, which can be one of the following: &#39;COUNTRY&#39;, &#39;COUNTRY_REGION&#39;, &#39;STATE_OR_PROVINCE&#39;, &#39;WORLD_REGION&#39;, or &#39;WORLDWIDE&#39;.--&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:RegionTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionName

`func (o *Region) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *Region) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *Region) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *Region) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRegionType

`func (o *Region) GetRegionType() string`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *Region) GetRegionTypeOk() (*string, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *Region) SetRegionType(v string)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *Region) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


