# UserScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** | The ID of the schedule. This ID is generated when the schedule was created by the createSchedule method. | [optional] 
**CreationDate** | Pointer to **string** | The creation date of the schedule in hours based on the 24-hour Coordinated Universal Time (UTC) clock. | [optional] 
**FeedType** | Pointer to **string** | The feedType associated with the schedule. | [optional] 
**LastModifiedDate** | Pointer to **string** | The date the schedule was last modified. | [optional] 
**PreferredTriggerDayOfMonth** | Pointer to **int32** | The preferred day of the month to trigger the schedule. This field can be used with preferredTriggerHour for monthly schedules. The last day of the month is used for numbers larger than the number of days in the month. | [optional] 
**PreferredTriggerDayOfWeek** | Pointer to **string** | The preferred day of the week to trigger the schedule. This field can be used with preferredTriggerHour for weekly schedules. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:DayOfWeekEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PreferredTriggerHour** | Pointer to **string** | The preferred two-digit hour of the day to trigger the schedule. Format: UTC hhZ For example, the following represents 11:00 am UTC: 11Z | [optional] 
**ScheduleEndDate** | Pointer to **string** | The timestamp on which the report generation (subscription) ends. After this date, the schedule status becomes INACTIVE. | [optional] 
**ScheduleName** | Pointer to **string** | The schedule name assigned by the user for the created schedule. Users assign this name for their reference. | [optional] 
**ScheduleStartDate** | Pointer to **string** | The timestamp that indicates the start of the report generation. | [optional] 
**ScheduleTemplateId** | Pointer to **string** | The ID of the template used to create this schedule. | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version of the feedType for the schedule. | [optional] 
**Status** | Pointer to **string** | The enumeration value that indicates the state of the schedule. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:StatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**StatusReason** | Pointer to **string** | The reason the schedule is inactive. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:StatusReasonEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewUserScheduleResponse

`func NewUserScheduleResponse() *UserScheduleResponse`

NewUserScheduleResponse instantiates a new UserScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserScheduleResponseWithDefaults

`func NewUserScheduleResponseWithDefaults() *UserScheduleResponse`

NewUserScheduleResponseWithDefaults instantiates a new UserScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleId

`func (o *UserScheduleResponse) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *UserScheduleResponse) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *UserScheduleResponse) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *UserScheduleResponse) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetCreationDate

`func (o *UserScheduleResponse) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserScheduleResponse) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserScheduleResponse) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UserScheduleResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetFeedType

`func (o *UserScheduleResponse) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *UserScheduleResponse) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *UserScheduleResponse) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *UserScheduleResponse) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *UserScheduleResponse) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *UserScheduleResponse) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *UserScheduleResponse) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *UserScheduleResponse) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetPreferredTriggerDayOfMonth

`func (o *UserScheduleResponse) GetPreferredTriggerDayOfMonth() int32`

GetPreferredTriggerDayOfMonth returns the PreferredTriggerDayOfMonth field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfMonthOk

`func (o *UserScheduleResponse) GetPreferredTriggerDayOfMonthOk() (*int32, bool)`

GetPreferredTriggerDayOfMonthOk returns a tuple with the PreferredTriggerDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfMonth

`func (o *UserScheduleResponse) SetPreferredTriggerDayOfMonth(v int32)`

SetPreferredTriggerDayOfMonth sets PreferredTriggerDayOfMonth field to given value.

### HasPreferredTriggerDayOfMonth

`func (o *UserScheduleResponse) HasPreferredTriggerDayOfMonth() bool`

HasPreferredTriggerDayOfMonth returns a boolean if a field has been set.

### GetPreferredTriggerDayOfWeek

`func (o *UserScheduleResponse) GetPreferredTriggerDayOfWeek() string`

GetPreferredTriggerDayOfWeek returns the PreferredTriggerDayOfWeek field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfWeekOk

`func (o *UserScheduleResponse) GetPreferredTriggerDayOfWeekOk() (*string, bool)`

GetPreferredTriggerDayOfWeekOk returns a tuple with the PreferredTriggerDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfWeek

`func (o *UserScheduleResponse) SetPreferredTriggerDayOfWeek(v string)`

SetPreferredTriggerDayOfWeek sets PreferredTriggerDayOfWeek field to given value.

### HasPreferredTriggerDayOfWeek

`func (o *UserScheduleResponse) HasPreferredTriggerDayOfWeek() bool`

HasPreferredTriggerDayOfWeek returns a boolean if a field has been set.

### GetPreferredTriggerHour

`func (o *UserScheduleResponse) GetPreferredTriggerHour() string`

GetPreferredTriggerHour returns the PreferredTriggerHour field if non-nil, zero value otherwise.

### GetPreferredTriggerHourOk

`func (o *UserScheduleResponse) GetPreferredTriggerHourOk() (*string, bool)`

GetPreferredTriggerHourOk returns a tuple with the PreferredTriggerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerHour

`func (o *UserScheduleResponse) SetPreferredTriggerHour(v string)`

SetPreferredTriggerHour sets PreferredTriggerHour field to given value.

### HasPreferredTriggerHour

`func (o *UserScheduleResponse) HasPreferredTriggerHour() bool`

HasPreferredTriggerHour returns a boolean if a field has been set.

### GetScheduleEndDate

`func (o *UserScheduleResponse) GetScheduleEndDate() string`

GetScheduleEndDate returns the ScheduleEndDate field if non-nil, zero value otherwise.

### GetScheduleEndDateOk

`func (o *UserScheduleResponse) GetScheduleEndDateOk() (*string, bool)`

GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndDate

`func (o *UserScheduleResponse) SetScheduleEndDate(v string)`

SetScheduleEndDate sets ScheduleEndDate field to given value.

### HasScheduleEndDate

`func (o *UserScheduleResponse) HasScheduleEndDate() bool`

HasScheduleEndDate returns a boolean if a field has been set.

### GetScheduleName

`func (o *UserScheduleResponse) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *UserScheduleResponse) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *UserScheduleResponse) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *UserScheduleResponse) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduleStartDate

`func (o *UserScheduleResponse) GetScheduleStartDate() string`

GetScheduleStartDate returns the ScheduleStartDate field if non-nil, zero value otherwise.

### GetScheduleStartDateOk

`func (o *UserScheduleResponse) GetScheduleStartDateOk() (*string, bool)`

GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartDate

`func (o *UserScheduleResponse) SetScheduleStartDate(v string)`

SetScheduleStartDate sets ScheduleStartDate field to given value.

### HasScheduleStartDate

`func (o *UserScheduleResponse) HasScheduleStartDate() bool`

HasScheduleStartDate returns a boolean if a field has been set.

### GetScheduleTemplateId

`func (o *UserScheduleResponse) GetScheduleTemplateId() string`

GetScheduleTemplateId returns the ScheduleTemplateId field if non-nil, zero value otherwise.

### GetScheduleTemplateIdOk

`func (o *UserScheduleResponse) GetScheduleTemplateIdOk() (*string, bool)`

GetScheduleTemplateIdOk returns a tuple with the ScheduleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTemplateId

`func (o *UserScheduleResponse) SetScheduleTemplateId(v string)`

SetScheduleTemplateId sets ScheduleTemplateId field to given value.

### HasScheduleTemplateId

`func (o *UserScheduleResponse) HasScheduleTemplateId() bool`

HasScheduleTemplateId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *UserScheduleResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *UserScheduleResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *UserScheduleResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *UserScheduleResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStatus

`func (o *UserScheduleResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserScheduleResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserScheduleResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserScheduleResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *UserScheduleResponse) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *UserScheduleResponse) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *UserScheduleResponse) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *UserScheduleResponse) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


