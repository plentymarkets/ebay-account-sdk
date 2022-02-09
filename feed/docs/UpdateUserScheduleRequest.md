# UpdateUserScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredTriggerDayOfMonth** | Pointer to **int32** | The preferred day of the month to trigger the schedule. This field can be used with preferredTriggerHour for monthly schedules. The last day of the month is used for numbers larger than the actual number of days in the month. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Minimum: 1 Maximum: 31 | [optional] 
**PreferredTriggerDayOfWeek** | Pointer to **string** | The preferred day of the week to trigger the schedule. This field can be used with preferredTriggerHour for weekly schedules. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:DayOfWeekEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PreferredTriggerHour** | Pointer to **string** | The preferred two-digit hour of the day to trigger the schedule. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC hhZ For example, the following represents 11:00 am UTC: 11Z Minimum: 00Z Maximum: 23Z | [optional] 
**ScheduleEndDate** | Pointer to **string** | The timestamp on which the schedule (report generation) ends. After this date, the schedule status becomes INACTIVE. Use this field, if available, to end the schedule in the future. This value must be later than scheduleStartDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents UTC October 10, 2021 at 10:00 AM: 2021-10-10T10Z | [optional] 
**ScheduleName** | Pointer to **string** | The schedule name assigned by the user for the created schedule. | [optional] 
**ScheduleStartDate** | Pointer to **string** | The timestamp to start generating the report. After this timestamp, the schedule status becomes active until either the scheduleEndDate occurs or the scheduleTemplateId becomes inactive. Use this field, if available, to start the schedule in the future but before the scheduleEndDate (if supplied). This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. Format: UTC yyyy-MM-ddTHHZ For example, the following represents a schedule start date of UTC October 01, 2020 at 12:00 PM: 2020-01-01T12Z | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version of the feedType for the schedule. This field is required if the feedType has a schema version. This field is available as specified by the template (scheduleTemplateId). The template can specify this field as optional or required, and optionally provides a default value. | [optional] 

## Methods

### NewUpdateUserScheduleRequest

`func NewUpdateUserScheduleRequest() *UpdateUserScheduleRequest`

NewUpdateUserScheduleRequest instantiates a new UpdateUserScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserScheduleRequestWithDefaults

`func NewUpdateUserScheduleRequestWithDefaults() *UpdateUserScheduleRequest`

NewUpdateUserScheduleRequestWithDefaults instantiates a new UpdateUserScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredTriggerDayOfMonth

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerDayOfMonth() int32`

GetPreferredTriggerDayOfMonth returns the PreferredTriggerDayOfMonth field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfMonthOk

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerDayOfMonthOk() (*int32, bool)`

GetPreferredTriggerDayOfMonthOk returns a tuple with the PreferredTriggerDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfMonth

`func (o *UpdateUserScheduleRequest) SetPreferredTriggerDayOfMonth(v int32)`

SetPreferredTriggerDayOfMonth sets PreferredTriggerDayOfMonth field to given value.

### HasPreferredTriggerDayOfMonth

`func (o *UpdateUserScheduleRequest) HasPreferredTriggerDayOfMonth() bool`

HasPreferredTriggerDayOfMonth returns a boolean if a field has been set.

### GetPreferredTriggerDayOfWeek

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerDayOfWeek() string`

GetPreferredTriggerDayOfWeek returns the PreferredTriggerDayOfWeek field if non-nil, zero value otherwise.

### GetPreferredTriggerDayOfWeekOk

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerDayOfWeekOk() (*string, bool)`

GetPreferredTriggerDayOfWeekOk returns a tuple with the PreferredTriggerDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerDayOfWeek

`func (o *UpdateUserScheduleRequest) SetPreferredTriggerDayOfWeek(v string)`

SetPreferredTriggerDayOfWeek sets PreferredTriggerDayOfWeek field to given value.

### HasPreferredTriggerDayOfWeek

`func (o *UpdateUserScheduleRequest) HasPreferredTriggerDayOfWeek() bool`

HasPreferredTriggerDayOfWeek returns a boolean if a field has been set.

### GetPreferredTriggerHour

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerHour() string`

GetPreferredTriggerHour returns the PreferredTriggerHour field if non-nil, zero value otherwise.

### GetPreferredTriggerHourOk

`func (o *UpdateUserScheduleRequest) GetPreferredTriggerHourOk() (*string, bool)`

GetPreferredTriggerHourOk returns a tuple with the PreferredTriggerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredTriggerHour

`func (o *UpdateUserScheduleRequest) SetPreferredTriggerHour(v string)`

SetPreferredTriggerHour sets PreferredTriggerHour field to given value.

### HasPreferredTriggerHour

`func (o *UpdateUserScheduleRequest) HasPreferredTriggerHour() bool`

HasPreferredTriggerHour returns a boolean if a field has been set.

### GetScheduleEndDate

`func (o *UpdateUserScheduleRequest) GetScheduleEndDate() string`

GetScheduleEndDate returns the ScheduleEndDate field if non-nil, zero value otherwise.

### GetScheduleEndDateOk

`func (o *UpdateUserScheduleRequest) GetScheduleEndDateOk() (*string, bool)`

GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndDate

`func (o *UpdateUserScheduleRequest) SetScheduleEndDate(v string)`

SetScheduleEndDate sets ScheduleEndDate field to given value.

### HasScheduleEndDate

`func (o *UpdateUserScheduleRequest) HasScheduleEndDate() bool`

HasScheduleEndDate returns a boolean if a field has been set.

### GetScheduleName

`func (o *UpdateUserScheduleRequest) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *UpdateUserScheduleRequest) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *UpdateUserScheduleRequest) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *UpdateUserScheduleRequest) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetScheduleStartDate

`func (o *UpdateUserScheduleRequest) GetScheduleStartDate() string`

GetScheduleStartDate returns the ScheduleStartDate field if non-nil, zero value otherwise.

### GetScheduleStartDateOk

`func (o *UpdateUserScheduleRequest) GetScheduleStartDateOk() (*string, bool)`

GetScheduleStartDateOk returns a tuple with the ScheduleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStartDate

`func (o *UpdateUserScheduleRequest) SetScheduleStartDate(v string)`

SetScheduleStartDate sets ScheduleStartDate field to given value.

### HasScheduleStartDate

`func (o *UpdateUserScheduleRequest) HasScheduleStartDate() bool`

HasScheduleStartDate returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *UpdateUserScheduleRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *UpdateUserScheduleRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *UpdateUserScheduleRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *UpdateUserScheduleRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


