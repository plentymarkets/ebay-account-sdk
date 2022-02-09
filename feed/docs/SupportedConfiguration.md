# SupportedConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | The default value for the property. If a value is omitted from the schedule and a default value is supplied, the default value is used. | [optional] 
**Property** | Pointer to **string** | Properties supported by the template. Properties can include the following: scheduleStartDate: The timestamp that the report generation (subscription) begins. After this timestamp, the schedule status becomes active until either the scheduleEndDate occurs or the scheduleTemplate becomes inactive. Format: UTC yyyy-MM-ddTHHZ scheduleEndDate: The timestamp that the report generation (subscription) ends. After this date, the schedule status becomes INACTIVE. Format: UTC yyyy-MM-ddTHHZ schemaVersion: The schema version of the schedule templates feedType. This field is required if the feedType has a schema version. preferredTriggerDayOfMonth: The preferred day of the month to trigger the schedule. preferredTriggerDayOfWeek: The preferred day of the week to trigger the schedule. preferredTriggerHour: The preferred two-digit hour of the day to trigger the schedule. Format: UTC hhZ | [optional] 
**Usage** | Pointer to **string** | Whether the specified property is REQUIRED or OPTIONAL. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:ConfigurationsUsageEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewSupportedConfiguration

`func NewSupportedConfiguration() *SupportedConfiguration`

NewSupportedConfiguration instantiates a new SupportedConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedConfigurationWithDefaults

`func NewSupportedConfigurationWithDefaults() *SupportedConfiguration`

NewSupportedConfigurationWithDefaults instantiates a new SupportedConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *SupportedConfiguration) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SupportedConfiguration) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SupportedConfiguration) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SupportedConfiguration) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetProperty

`func (o *SupportedConfiguration) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SupportedConfiguration) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SupportedConfiguration) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SupportedConfiguration) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetUsage

`func (o *SupportedConfiguration) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SupportedConfiguration) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SupportedConfiguration) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SupportedConfiguration) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


