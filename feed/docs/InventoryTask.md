# InventoryTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** | The ID of the task. This ID is generated when the task was created by the createInventoryTask method. | [optional] 
**Status** | Pointer to **string** | The status of the task. Users must wait until status is complete before moving on to the next step (such as uploading/downloading a file). For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:FeedStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**FeedType** | Pointer to **string** | The feed type associated with the inventory task. | [optional] 
**CreationDate** | Pointer to **string** | The date the task was created. | [optional] 
**CompletionDate** | Pointer to **string** | The timestamp when the task status went into the COMPLETED, COMPLETED_WITH_ERROR, or PARTIALLY_PROCESSED state. This field is only returned if the status is one of the three completed values. | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version number associated with the task. | [optional] 
**DetailHref** | Pointer to **string** | The path to the call URI used to retrieve the task. This field points to the getInventoryTask URI. | [optional] 
**UploadSummary** | Pointer to [**UploadSummary**](UploadSummary.md) |  | [optional] 
**FilterCriteria** | Pointer to [**InventoryFilterCriteria**](InventoryFilterCriteria.md) |  | [optional] 
**InventoryFileTemplate** | Pointer to **string** | The inventory file template used to return specific types of inventory tasks, if set in the createInventoryTask method. This field does not apply to LMS_ACTIVE_INVENTORY_REPORT feed types. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:InventoryFileTemplateEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewInventoryTask

`func NewInventoryTask() *InventoryTask`

NewInventoryTask instantiates a new InventoryTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryTaskWithDefaults

`func NewInventoryTaskWithDefaults() *InventoryTask`

NewInventoryTaskWithDefaults instantiates a new InventoryTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *InventoryTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *InventoryTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *InventoryTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *InventoryTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetStatus

`func (o *InventoryTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InventoryTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InventoryTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InventoryTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFeedType

`func (o *InventoryTask) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *InventoryTask) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *InventoryTask) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *InventoryTask) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetCreationDate

`func (o *InventoryTask) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *InventoryTask) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *InventoryTask) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *InventoryTask) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *InventoryTask) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *InventoryTask) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *InventoryTask) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *InventoryTask) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *InventoryTask) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *InventoryTask) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *InventoryTask) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *InventoryTask) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetDetailHref

`func (o *InventoryTask) GetDetailHref() string`

GetDetailHref returns the DetailHref field if non-nil, zero value otherwise.

### GetDetailHrefOk

`func (o *InventoryTask) GetDetailHrefOk() (*string, bool)`

GetDetailHrefOk returns a tuple with the DetailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailHref

`func (o *InventoryTask) SetDetailHref(v string)`

SetDetailHref sets DetailHref field to given value.

### HasDetailHref

`func (o *InventoryTask) HasDetailHref() bool`

HasDetailHref returns a boolean if a field has been set.

### GetUploadSummary

`func (o *InventoryTask) GetUploadSummary() UploadSummary`

GetUploadSummary returns the UploadSummary field if non-nil, zero value otherwise.

### GetUploadSummaryOk

`func (o *InventoryTask) GetUploadSummaryOk() (*UploadSummary, bool)`

GetUploadSummaryOk returns a tuple with the UploadSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSummary

`func (o *InventoryTask) SetUploadSummary(v UploadSummary)`

SetUploadSummary sets UploadSummary field to given value.

### HasUploadSummary

`func (o *InventoryTask) HasUploadSummary() bool`

HasUploadSummary returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *InventoryTask) GetFilterCriteria() InventoryFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *InventoryTask) GetFilterCriteriaOk() (*InventoryFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *InventoryTask) SetFilterCriteria(v InventoryFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *InventoryTask) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetInventoryFileTemplate

`func (o *InventoryTask) GetInventoryFileTemplate() string`

GetInventoryFileTemplate returns the InventoryFileTemplate field if non-nil, zero value otherwise.

### GetInventoryFileTemplateOk

`func (o *InventoryTask) GetInventoryFileTemplateOk() (*string, bool)`

GetInventoryFileTemplateOk returns a tuple with the InventoryFileTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryFileTemplate

`func (o *InventoryTask) SetInventoryFileTemplate(v string)`

SetInventoryFileTemplate sets InventoryFileTemplate field to given value.

### HasInventoryFileTemplate

`func (o *InventoryTask) HasInventoryFileTemplate() bool`

HasInventoryFileTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


