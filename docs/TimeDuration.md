# TimeDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** | A time-measurement unit that specifies a singular period of time.  &lt;br&gt;&lt;br&gt;A span of time is defined when you apply the value specified in the &lt;b&gt;value&lt;/b&gt; field to the value specified for &lt;b&gt;unit&lt;/b&gt;. &lt;br&gt;&lt;br&gt;Time-measurement units can be YEAR, MONTH, DAY, and so on. See &lt;b&gt;TimeDurationUnitEnum&lt;/b&gt; for a complete list of possible time-measurement units. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/ba:TimeDurationUnitEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Value** | Pointer to **int32** | An integer that represents an amount of time, as measured by the time-measurement unit specified in the &lt;b&gt;unit&lt;/b&gt; field. | [optional] 

## Methods

### NewTimeDuration

`func NewTimeDuration() *TimeDuration`

NewTimeDuration instantiates a new TimeDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeDurationWithDefaults

`func NewTimeDurationWithDefaults() *TimeDuration`

NewTimeDurationWithDefaults instantiates a new TimeDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *TimeDuration) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeDuration) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeDuration) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeDuration) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValue

`func (o *TimeDuration) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimeDuration) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimeDuration) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *TimeDuration) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


