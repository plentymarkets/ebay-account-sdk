# InventoryTaskCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The path to the call URI that produced the current page of results. | [optional] 
**Limit** | Pointer to **int32** | The value of the limit parameter submitted in the request, which is the maximum number of inventory tasks to return per page, from the result set. A result set is the complete set of tasks returned by the method. Note: Though this parameter is not required to be submitted in the request, the parameter defaults to 10 if omitted. Additionally, if this is the last or only page of the result set, the page may contain fewer tasks than the limit value submitted in the request. To determine the number of pages in a result set, divide the total value (total number of tasks matching the input criteria) by this limit value, and then round up to the next integer. For example, if the total value was 120 (120 total tasks) and the limit value was 50 (show 50 tasks per page), the total number of pages in the result set is three, so the seller would have to make three separate getInventoryTasks calls to view all tasks matching the input criteria. | [optional] 
**Next** | Pointer to **string** | The path to the call URI for the next page of results. This value is returned if there is an additional page of results to return from the result set. | [optional] 
**Offset** | Pointer to **int32** | The number of results skipped in the result set before listing the first returned result. This value can be specified in the request with the offset query parameter. Note: The items in a paginated result set use a zero-based list, where the first item in the list has an offset of 0. | [optional] 
**Prev** | Pointer to **string** | The path to the call URI for the previous page of results. This is returned if there is a previous page of results from the result set. | [optional] 
**Tasks** | Pointer to [**[]InventoryTask**](InventoryTask.md) | An array of the inventory tasks on this page. The tasks are sorted by creation date. Note: An empty array is returned if the filter criteria excludes all tasks. | [optional] 
**Total** | Pointer to **int32** | The total number of inventory tasks that match the input criteria. | [optional] 

## Methods

### NewInventoryTaskCollection

`func NewInventoryTaskCollection() *InventoryTaskCollection`

NewInventoryTaskCollection instantiates a new InventoryTaskCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryTaskCollectionWithDefaults

`func NewInventoryTaskCollectionWithDefaults() *InventoryTaskCollection`

NewInventoryTaskCollectionWithDefaults instantiates a new InventoryTaskCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *InventoryTaskCollection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InventoryTaskCollection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InventoryTaskCollection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InventoryTaskCollection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *InventoryTaskCollection) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InventoryTaskCollection) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InventoryTaskCollection) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InventoryTaskCollection) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *InventoryTaskCollection) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *InventoryTaskCollection) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *InventoryTaskCollection) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *InventoryTaskCollection) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *InventoryTaskCollection) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InventoryTaskCollection) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InventoryTaskCollection) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *InventoryTaskCollection) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *InventoryTaskCollection) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *InventoryTaskCollection) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *InventoryTaskCollection) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *InventoryTaskCollection) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetTasks

`func (o *InventoryTaskCollection) GetTasks() []InventoryTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InventoryTaskCollection) GetTasksOk() (*[]InventoryTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InventoryTaskCollection) SetTasks(v []InventoryTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InventoryTaskCollection) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTotal

`func (o *InventoryTaskCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InventoryTaskCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InventoryTaskCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InventoryTaskCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


