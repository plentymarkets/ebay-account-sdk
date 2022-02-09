# ScheduleTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedType** | Pointer to **string** | The feed type of the schedule template. Note: When calling createSchedule and updateSchedule methods you must match the feed type specified by the schedule template (this feedType). | [optional] 
**Frequency** | Pointer to **string** | This field specifies how often the schedule is generated. If set to HALF_HOUR or ONE_HOUR, you cannot set a preferredTriggerHour using createSchedule or updateSchedule. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:FrequencyEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Name** | Pointer to **string** | The template name provided by the template. | [optional] 
**ScheduleTemplateId** | Pointer to **string** | The ID of the template. Use this ID to create a schedule based on the properties of this schedule template. | [optional] 
**Status** | Pointer to **string** | The present status of the template. You cannot create or modify a schedule using a template with an INACTIVE status. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:StatusEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**SupportedConfigurations** | Pointer to [**[]SupportedConfiguration**](SupportedConfiguration.md) | An array of the configuration supported by this template. | [optional] 

## Methods

### NewScheduleTemplateResponse

`func NewScheduleTemplateResponse() *ScheduleTemplateResponse`

NewScheduleTemplateResponse instantiates a new ScheduleTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTemplateResponseWithDefaults

`func NewScheduleTemplateResponseWithDefaults() *ScheduleTemplateResponse`

NewScheduleTemplateResponseWithDefaults instantiates a new ScheduleTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedType

`func (o *ScheduleTemplateResponse) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *ScheduleTemplateResponse) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *ScheduleTemplateResponse) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *ScheduleTemplateResponse) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFrequency

`func (o *ScheduleTemplateResponse) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScheduleTemplateResponse) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScheduleTemplateResponse) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ScheduleTemplateResponse) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *ScheduleTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleTemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleTemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleTemplateId

`func (o *ScheduleTemplateResponse) GetScheduleTemplateId() string`

GetScheduleTemplateId returns the ScheduleTemplateId field if non-nil, zero value otherwise.

### GetScheduleTemplateIdOk

`func (o *ScheduleTemplateResponse) GetScheduleTemplateIdOk() (*string, bool)`

GetScheduleTemplateIdOk returns a tuple with the ScheduleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTemplateId

`func (o *ScheduleTemplateResponse) SetScheduleTemplateId(v string)`

SetScheduleTemplateId sets ScheduleTemplateId field to given value.

### HasScheduleTemplateId

`func (o *ScheduleTemplateResponse) HasScheduleTemplateId() bool`

HasScheduleTemplateId returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduleTemplateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleTemplateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleTemplateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduleTemplateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportedConfigurations

`func (o *ScheduleTemplateResponse) GetSupportedConfigurations() []SupportedConfiguration`

GetSupportedConfigurations returns the SupportedConfigurations field if non-nil, zero value otherwise.

### GetSupportedConfigurationsOk

`func (o *ScheduleTemplateResponse) GetSupportedConfigurationsOk() (*[]SupportedConfiguration, bool)`

GetSupportedConfigurationsOk returns a tuple with the SupportedConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedConfigurations

`func (o *ScheduleTemplateResponse) SetSupportedConfigurations(v []SupportedConfiguration)`

SetSupportedConfigurations sets SupportedConfigurations field to given value.

### HasSupportedConfigurations

`func (o *ScheduleTemplateResponse) HasSupportedConfigurations() bool`

HasSupportedConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


