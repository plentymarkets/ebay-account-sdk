# OrderTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDate** | Pointer to **string** | The timestamp when the task went into the COMPLETED or COMPLETED_WITH_ERROR state. This state means that eBay has compiled the report for the seller based on the seller&amp;rsquo;s filter criteria, and the seller can run a getResultFile call to download the report. | [optional] 
**CreationDate** | Pointer to **string** | The date the task was created. | [optional] 
**DetailHref** | Pointer to **string** | The path to the call URI used to retrieve the task. | [optional] 
**FeedType** | Pointer to **string** | The feed type associated with the task. | [optional] 
**FilterCriteria** | Pointer to [**OrderFilterCriteria**](OrderFilterCriteria.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version number associated with the create task. | [optional] 
**Status** | Pointer to **string** | The enumeration value that indicates the state of the task that was submitted in the request. See FeedStatusEnum for information. The values COMPLETED and COMPLETED_WITH_ERROR indicate the Order Report file is ready to download. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:FeedStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**TaskId** | Pointer to **string** | The ID of the task that was submitted in the request. | [optional] 
**UploadSummary** | Pointer to [**UploadSummary**](UploadSummary.md) |  | [optional] 

## Methods

### NewOrderTask

`func NewOrderTask() *OrderTask`

NewOrderTask instantiates a new OrderTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaskWithDefaults

`func NewOrderTaskWithDefaults() *OrderTask`

NewOrderTaskWithDefaults instantiates a new OrderTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionDate

`func (o *OrderTask) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *OrderTask) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *OrderTask) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *OrderTask) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *OrderTask) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *OrderTask) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *OrderTask) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *OrderTask) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDetailHref

`func (o *OrderTask) GetDetailHref() string`

GetDetailHref returns the DetailHref field if non-nil, zero value otherwise.

### GetDetailHrefOk

`func (o *OrderTask) GetDetailHrefOk() (*string, bool)`

GetDetailHrefOk returns a tuple with the DetailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailHref

`func (o *OrderTask) SetDetailHref(v string)`

SetDetailHref sets DetailHref field to given value.

### HasDetailHref

`func (o *OrderTask) HasDetailHref() bool`

HasDetailHref returns a boolean if a field has been set.

### GetFeedType

`func (o *OrderTask) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *OrderTask) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *OrderTask) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *OrderTask) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *OrderTask) GetFilterCriteria() OrderFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *OrderTask) GetFilterCriteriaOk() (*OrderFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *OrderTask) SetFilterCriteria(v OrderFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *OrderTask) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *OrderTask) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *OrderTask) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *OrderTask) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *OrderTask) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStatus

`func (o *OrderTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *OrderTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *OrderTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *OrderTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *OrderTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetUploadSummary

`func (o *OrderTask) GetUploadSummary() UploadSummary`

GetUploadSummary returns the UploadSummary field if non-nil, zero value otherwise.

### GetUploadSummaryOk

`func (o *OrderTask) GetUploadSummaryOk() (*UploadSummary, bool)`

GetUploadSummaryOk returns a tuple with the UploadSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSummary

`func (o *OrderTask) SetUploadSummary(v UploadSummary)`

SetUploadSummary sets UploadSummary field to given value.

### HasUploadSummary

`func (o *OrderTask) HasUploadSummary() bool`

HasUploadSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


