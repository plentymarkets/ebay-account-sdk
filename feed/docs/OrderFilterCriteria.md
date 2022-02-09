# OrderFilterCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**ModifiedDateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**OrderStatus** | Pointer to **string** | The order status of the orders returned. If the filter is omitted from createOrderTask call, orders that are in both ACTIVE and COMPLETED states are returned. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:OrderStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewOrderFilterCriteria

`func NewOrderFilterCriteria() *OrderFilterCriteria`

NewOrderFilterCriteria instantiates a new OrderFilterCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFilterCriteriaWithDefaults

`func NewOrderFilterCriteriaWithDefaults() *OrderFilterCriteria`

NewOrderFilterCriteriaWithDefaults instantiates a new OrderFilterCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDateRange

`func (o *OrderFilterCriteria) GetCreationDateRange() DateRange`

GetCreationDateRange returns the CreationDateRange field if non-nil, zero value otherwise.

### GetCreationDateRangeOk

`func (o *OrderFilterCriteria) GetCreationDateRangeOk() (*DateRange, bool)`

GetCreationDateRangeOk returns a tuple with the CreationDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateRange

`func (o *OrderFilterCriteria) SetCreationDateRange(v DateRange)`

SetCreationDateRange sets CreationDateRange field to given value.

### HasCreationDateRange

`func (o *OrderFilterCriteria) HasCreationDateRange() bool`

HasCreationDateRange returns a boolean if a field has been set.

### GetModifiedDateRange

`func (o *OrderFilterCriteria) GetModifiedDateRange() DateRange`

GetModifiedDateRange returns the ModifiedDateRange field if non-nil, zero value otherwise.

### GetModifiedDateRangeOk

`func (o *OrderFilterCriteria) GetModifiedDateRangeOk() (*DateRange, bool)`

GetModifiedDateRangeOk returns a tuple with the ModifiedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateRange

`func (o *OrderFilterCriteria) SetModifiedDateRange(v DateRange)`

SetModifiedDateRange sets ModifiedDateRange field to given value.

### HasModifiedDateRange

`func (o *OrderFilterCriteria) HasModifiedDateRange() bool`

HasModifiedDateRange returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderFilterCriteria) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderFilterCriteria) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderFilterCriteria) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderFilterCriteria) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


