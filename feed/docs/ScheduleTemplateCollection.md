# ScheduleTemplateCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The path to the call URI that produced the current page of results. | [optional] 
**Limit** | Pointer to **int32** | The value of the limit parameter submitted in the request, which is the maximum number of schedule templates to return per page, from the result set. A result set is the complete set of schedule templates returned by the method. Note: Though this parameter is not required to be submitted in the request, the parameter defaults to 10 if omitted. Note: If this is the last or only page of the result set, the page may contain fewer tasks than the limit value. To determine the number of pages in a result set, divide the total value (total number of tasks matching input criteria) by this limit value, and then round up to the next integer. For example, if the total value was 120 (120 total tasks) and the limit value was 50 (show 50 tasks per page), the total number of pages in the result set is three, so the seller would have to make three separate getScheduleTemplates calls to view all tasks matching the input criteria. | [optional] 
**Next** | Pointer to **string** | The path to the call URI for the next page of results. This value is returned if there is an additional page of results to return from the result set. | [optional] 
**Offset** | Pointer to **int32** | The number of results skipped in the result set before listing the first returned result. This value can be set in the request with the offset query parameter. Note: The items in a paginated result set use a zero-based list where the first item in the list has an offset of 0. | [optional] 
**Prev** | Pointer to **string** | The path to the call URI for the previous page of results. This is returned if there is a previous page of results from the result set. | [optional] 
**ScheduleTemplates** | Pointer to [**[]ScheduleTemplateResponse**](ScheduleTemplateResponse.md) | An array of the schedule templates on this page. An empty array is returned if the filter criteria excludes all tasks. | [optional] 
**Total** | Pointer to **int32** | The total number of schedule templates that match the input criteria. | [optional] 

## Methods

### NewScheduleTemplateCollection

`func NewScheduleTemplateCollection() *ScheduleTemplateCollection`

NewScheduleTemplateCollection instantiates a new ScheduleTemplateCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTemplateCollectionWithDefaults

`func NewScheduleTemplateCollectionWithDefaults() *ScheduleTemplateCollection`

NewScheduleTemplateCollectionWithDefaults instantiates a new ScheduleTemplateCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ScheduleTemplateCollection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ScheduleTemplateCollection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ScheduleTemplateCollection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ScheduleTemplateCollection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *ScheduleTemplateCollection) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ScheduleTemplateCollection) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ScheduleTemplateCollection) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ScheduleTemplateCollection) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *ScheduleTemplateCollection) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ScheduleTemplateCollection) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ScheduleTemplateCollection) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ScheduleTemplateCollection) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *ScheduleTemplateCollection) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ScheduleTemplateCollection) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ScheduleTemplateCollection) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ScheduleTemplateCollection) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *ScheduleTemplateCollection) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *ScheduleTemplateCollection) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *ScheduleTemplateCollection) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *ScheduleTemplateCollection) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetScheduleTemplates

`func (o *ScheduleTemplateCollection) GetScheduleTemplates() []ScheduleTemplateResponse`

GetScheduleTemplates returns the ScheduleTemplates field if non-nil, zero value otherwise.

### GetScheduleTemplatesOk

`func (o *ScheduleTemplateCollection) GetScheduleTemplatesOk() (*[]ScheduleTemplateResponse, bool)`

GetScheduleTemplatesOk returns a tuple with the ScheduleTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTemplates

`func (o *ScheduleTemplateCollection) SetScheduleTemplates(v []ScheduleTemplateResponse)`

SetScheduleTemplates sets ScheduleTemplates field to given value.

### HasScheduleTemplates

`func (o *ScheduleTemplateCollection) HasScheduleTemplates() bool`

HasScheduleTemplates returns a boolean if a field has been set.

### GetTotal

`func (o *ScheduleTemplateCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ScheduleTemplateCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ScheduleTemplateCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ScheduleTemplateCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


