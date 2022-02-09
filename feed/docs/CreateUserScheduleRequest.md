# CreateUserScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedType** | Pointer to **string** | The name of the feed type for the created schedule. Match the feed_type from the schedule template associated with this schedule. | [optional] 
**PreferredTriggerDayOfMonth** | Pointer to **int32** | The preferred day of the month to trigger the schedule. This field can be used with preferredTriggerHour for monthly schedules. The last day of the month is used for numbers larger than the actual number of days in the month. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Minimum: 1 Maximum: 31 | [optional] 
**PreferredTriggerDayOfWeek** | Pointer to **string** | The preferred day of the week to trigger the schedule. This field can be used with preferredTriggerHour for weekly schedules. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:DayOfWeekEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PreferredTriggerHour** | Pointer to **string** | The preferred two-digit hour of the day to trigger the schedule. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC hhZ For example, the following represents 11:00 am UTC: 11Z | [optional] 
**ScheduleEndDate** | Pointer to **string** | The timestamp on which the report generation (subscription) ends. After this date, the schedule status becomes INACTIVE. Use this field, if available, to end the schedule in the future. This value must be later than scheduleStartDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents UTC October 10, 2021 at 10:00 AM: 2021-10-10T10Z | [optional] 
**ScheduleName** | Pointer to **string** | The schedule name assigned by the user for the created schedule. | [optional] 
**ScheduleStartDate** | Pointer to **string** | The timestamp to start generating the report. After this timestamp, the schedule status becomes active until either the scheduleEndDate occurs or the scheduleTemplateId becomes inactive. Use this field, if available, to start the schedule in the future but before the scheduleEndDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents a schedule start date of UTC October 01, 2020 at 12:00 PM: 2020-01-01T12Z | [optional] 
**ScheduleTemplateId** | Pointer to **string** | The ID of the template associated with the schedule ID. You can get this ID from the documentation or by calling the getScheduleTemplates method. This method requires a schedule template ID that is ACTIVE. | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version of the schedule feedType. This field is required if the feedType has a schema version. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. | [optional] 

## Methods

### NewCreateUserScheduleRequest

`func NewCreateUserScheduleRequest() *CreateUserScheduleRequest`

NewCreateUserScheduleRequest instantiates a new CreateUserScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserScheduleRequestWithDefaults

`func NewCreateUserScheduleRequestWithDefaults() *CreateUserScheduleRequest`

NewCreateUserScheduleRequestWithDefaults instantiates a new CreateUserScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedType

`func (o *CreateUserScheduleRequest) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *CreateUserScheduleRequest) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *CreateUserScheduleRequest) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *CreateUserScheduleRequest) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetPreferredTriggerDayOfMonth

`func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfMonth() int32`

GetPreferredTriggerDayOfMonth returns the PreferredTriggerDayOfMonth field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfMonthOk

`func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfMonthOk() (*int32, bool)`

GetPreferredTriggerDayOfMonthOk returns a tuple with the PreferredTriggerDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfMonth

`func (o *CreateUserScheduleRequest) SetPreferredTriggerDayOfMonth(v int32)`

SetPreferredTriggerDayOfMonth sets PreferredTriggerDayOfMonth field to given value.

### HasPreferredTriggerDayOfMonth

`func (o *CreateUserScheduleRequest) HasPreferredTriggerDayOfMonth() bool`

HasPreferredTriggerDayOfMonth returns a boolean if a field has been set.

### GetPreferredTriggerDayOfWeek

`func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfWeek() string`

GetPreferredTriggerDayOfWeek returns the PreferredTriggerDayOfWeek field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfWeekOk

`func (o *CreateUserScheduleRequest) GetPreferredTriggerDayOfWeekOk() (*string, bool)`

GetPreferredTriggerDayOfWeekOk returns a tuple with the PreferredTriggerDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfWeek

`func (o *CreateUserScheduleRequest) SetPreferredTriggerDayOfWeek(v string)`

SetPreferredTriggerDayOfWeek sets PreferredTriggerDayOfWeek field to given value.

### HasPreferredTriggerDayOfWeek

`func (o *CreateUserScheduleRequest) HasPreferredTriggerDayOfWeek() bool`

HasPreferredTriggerDayOfWeek returns a boolean if a field has been set.

### GetPreferredTriggerHour

`func (o *CreateUserScheduleRequest) GetPreferredTriggerHour() string`

GetPreferredTriggerHour returns the PreferredTriggerHour field if non-nil, zero value otherwise.

### GetPreferredTriggerHourOk

`func (o *CreateUserScheduleRequest) GetPreferredTriggerHourOk() (*string, bool)`

GetPreferredTriggerHourOk returns a tuple with the PreferredTriggerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerHour

`func (o *CreateUserScheduleRequest) SetPreferredTriggerHour(v string)`

SetPreferredTriggerHour sets PreferredTriggerHour field to given value.

### HasPreferredTriggerHour

`func (o *CreateUserScheduleRequest) HasPreferredTriggerHour() bool`

HasPreferredTriggerHour returns a boolean if a field has been set.

### GetScheduleEndDate

`func (o *CreateUserScheduleRequest) GetScheduleEndDate() string`

GetScheduleEndDate returns the ScheduleEndDate field if non-nil, zero value otherwise.

### GetScheduleEndDateOk

`func (o *CreateUserScheduleRequest) GetScheduleEndDateOk() (*string, bool)`

GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndDate

`func (o *CreateUserScheduleRequest) SetScheduleEndDate(v string)`

SetScheduleEndDate sets ScheduleEndDate field to given value.

### HasScheduleEndDate

`func (o *CreateUserScheduleRequest) HasScheduleEndDate() bool`

HasScheduleEndDate returns a boolean if a field has been set.

### GetScheduleName

`func (o *CreateUserScheduleRequest) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *CreateUserScheduleRequest) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *CreateUserScheduleRequest) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *CreateUserScheduleRequest) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduleStartDate

`func (o *CreateUserScheduleRequest) GetScheduleStartDate() string`

GetScheduleStartDate returns the ScheduleStartDate field if non-nil, zero value otherwise.

### GetScheduleStartDateOk

`func (o *CreateUserScheduleRequest) GetScheduleStartDateOk() (*string, bool)`

GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartDate

`func (o *CreateUserScheduleRequest) SetScheduleStartDate(v string)`

SetScheduleStartDate sets ScheduleStartDate field to given value.

### HasScheduleStartDate

`func (o *CreateUserScheduleRequest) HasScheduleStartDate() bool`

HasScheduleStartDate returns a boolean if a field has been set.

### GetScheduleTemplateId

`func (o *CreateUserScheduleRequest) GetScheduleTemplateId() string`

GetScheduleTemplateId returns the ScheduleTemplateId field if non-nil, zero value otherwise.

### GetScheduleTemplateIdOk

`func (o *CreateUserScheduleRequest) GetScheduleTemplateIdOk() (*string, bool)`

GetScheduleTemplateIdOk returns a tuple with the ScheduleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTemplateId

`func (o *CreateUserScheduleRequest) SetScheduleTemplateId(v string)`

SetScheduleTemplateId sets ScheduleTemplateId field to given value.

### HasScheduleTemplateId

`func (o *CreateUserScheduleRequest) HasScheduleTemplateId() bool`

HasScheduleTemplateId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateUserScheduleRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateUserScheduleRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateUserScheduleRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateUserScheduleRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


