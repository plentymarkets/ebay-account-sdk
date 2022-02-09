# ServiceMetricsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDate** | Pointer to **string** | The timestamp when the customer service metrics task went into the COMPLETED or COMPLETED_WITH_ERROR state. This field is only returned if the status is one of the two completed values. This state means that eBay has compiled the report for the seller based on the seller&amp;rsquo;s filter criteria, and the seller can run a getResultFile call to download the report. | [optional] 
**CreationDate** | Pointer to **string** | The date the customer service metrics task was created. | [optional] 
**DetailHref** | Pointer to **string** | The relative getCustomerServiceMetricTask call URI path to retrieve the corresponding task. | [optional] 
**FeedType** | Pointer to **string** | The feed type associated with the task. | [optional] 
**FilterCriteria** | Pointer to [**CustomerServiceMetricsFilterCriteria**](CustomerServiceMetricsFilterCriteria.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version number of the file format. If omitted, the default value is used. Default value: 1.0 | [optional] 
**Status** | Pointer to **string** | An enumeration value that indicates the state of the task. See FeedStatusEnum for values. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:FeedStatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**TaskId** | Pointer to **string** | The unique eBay-assigned ID of the task. | [optional] 

## Methods

### NewServiceMetricsTask

`func NewServiceMetricsTask() *ServiceMetricsTask`

NewServiceMetricsTask instantiates a new ServiceMetricsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetricsTaskWithDefaults

`func NewServiceMetricsTaskWithDefaults() *ServiceMetricsTask`

NewServiceMetricsTaskWithDefaults instantiates a new ServiceMetricsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionDate

`func (o *ServiceMetricsTask) GetCompletionDate() string`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *ServiceMetricsTask) GetCompletionDateOk() (*string, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *ServiceMetricsTask) SetCompletionDate(v string)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *ServiceMetricsTask) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *ServiceMetricsTask) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ServiceMetricsTask) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ServiceMetricsTask) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ServiceMetricsTask) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDetailHref

`func (o *ServiceMetricsTask) GetDetailHref() string`

GetDetailHref returns the DetailHref field if non-nil, zero value otherwise.

### GetDetailHrefOk

`func (o *ServiceMetricsTask) GetDetailHrefOk() (*string, bool)`

GetDetailHrefOk returns a tuple with the DetailHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailHref

`func (o *ServiceMetricsTask) SetDetailHref(v string)`

SetDetailHref sets DetailHref field to given value.

### HasDetailHref

`func (o *ServiceMetricsTask) HasDetailHref() bool`

HasDetailHref returns a boolean if a field has been set.

### GetFeedType

`func (o *ServiceMetricsTask) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *ServiceMetricsTask) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *ServiceMetricsTask) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *ServiceMetricsTask) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *ServiceMetricsTask) GetFilterCriteria() CustomerServiceMetricsFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *ServiceMetricsTask) GetFilterCriteriaOk() (*CustomerServiceMetricsFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *ServiceMetricsTask) SetFilterCriteria(v CustomerServiceMetricsFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *ServiceMetricsTask) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ServiceMetricsTask) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ServiceMetricsTask) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ServiceMetricsTask) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ServiceMetricsTask) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceMetricsTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceMetricsTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceMetricsTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceMetricsTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *ServiceMetricsTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ServiceMetricsTask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ServiceMetricsTask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ServiceMetricsTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


