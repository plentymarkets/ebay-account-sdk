# InventoryFilterCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**ModifiedDateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**ListingFormat** | Pointer to **string** | The type of buying option for the order. Supports LMS_ACTIVE_INVENTORY_REPORT. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:ListingFormatEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ListingStatus** | Pointer to **string** | The status of the listing (whether the listing was unsold or is active). The UNSOLD value does not apply to LMS_ACTIVE_INVENTORY_REPORT feed types. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:ListingStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewInventoryFilterCriteria

`func NewInventoryFilterCriteria() *InventoryFilterCriteria`

NewInventoryFilterCriteria instantiates a new InventoryFilterCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryFilterCriteriaWithDefaults

`func NewInventoryFilterCriteriaWithDefaults() *InventoryFilterCriteria`

NewInventoryFilterCriteriaWithDefaults instantiates a new InventoryFilterCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDateRange

`func (o *InventoryFilterCriteria) GetCreationDateRange() DateRange`

GetCreationDateRange returns the CreationDateRange field if non-nil, zero value otherwise.

### GetCreationDateRangeOk

`func (o *InventoryFilterCriteria) GetCreationDateRangeOk() (*DateRange, bool)`

GetCreationDateRangeOk returns a tuple with the CreationDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateRange

`func (o *InventoryFilterCriteria) SetCreationDateRange(v DateRange)`

SetCreationDateRange sets CreationDateRange field to given value.

### HasCreationDateRange

`func (o *InventoryFilterCriteria) HasCreationDateRange() bool`

HasCreationDateRange returns a boolean if a field has been set.

### GetModifiedDateRange

`func (o *InventoryFilterCriteria) GetModifiedDateRange() DateRange`

GetModifiedDateRange returns the ModifiedDateRange field if non-nil, zero value otherwise.

### GetModifiedDateRangeOk

`func (o *InventoryFilterCriteria) GetModifiedDateRangeOk() (*DateRange, bool)`

GetModifiedDateRangeOk returns a tuple with the ModifiedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateRange

`func (o *InventoryFilterCriteria) SetModifiedDateRange(v DateRange)`

SetModifiedDateRange sets ModifiedDateRange field to given value.

### HasModifiedDateRange

`func (o *InventoryFilterCriteria) HasModifiedDateRange() bool`

HasModifiedDateRange returns a boolean if a field has been set.

### GetListingFormat

`func (o *InventoryFilterCriteria) GetListingFormat() string`

GetListingFormat returns the ListingFormat field if non-nil, zero value otherwise.

### GetListingFormatOk

`func (o *InventoryFilterCriteria) GetListingFormatOk() (*string, bool)`

GetListingFormatOk returns a tuple with the ListingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingFormat

`func (o *InventoryFilterCriteria) SetListingFormat(v string)`

SetListingFormat sets ListingFormat field to given value.

### HasListingFormat

`func (o *InventoryFilterCriteria) HasListingFormat() bool`

HasListingFormat returns a boolean if a field has been set.

### GetListingStatus

`func (o *InventoryFilterCriteria) GetListingStatus() string`

GetListingStatus returns the ListingStatus field if non-nil, zero value otherwise.

### GetListingStatusOk

`func (o *InventoryFilterCriteria) GetListingStatusOk() (*string, bool)`

GetListingStatusOk returns a tuple with the ListingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingStatus

`func (o *InventoryFilterCriteria) SetListingStatus(v string)`

SetListingStatus sets ListingStatus field to given value.

### HasListingStatus

`func (o *InventoryFilterCriteria) HasListingStatus() bool`

HasListingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


