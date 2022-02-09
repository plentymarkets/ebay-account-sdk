# DateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | The beginning date in the range. If the parent type is included, both the from and/or the to fields become conditionally required. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks within a range yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ Tasks created on March 31, 2021 2021-03-31T00:00:00.000Z..2021-03-31T00:00:00.000Z | [optional] 
**To** | Pointer to **string** | The end date for the date range, which is inclusive. If the parent type is included, both the from and/or the to fields become conditionally required. For example: Tasks within a range yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ Tasks created on March 31, 2021 2021-03-31T00:00:00.000Z..2021-03-31T00:00:00.000Z | [optional] 

## Methods

### NewDateRange

`func NewDateRange() *DateRange`

NewDateRange instantiates a new DateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeWithDefaults

`func NewDateRangeWithDefaults() *DateRange`

NewDateRangeWithDefaults instantiates a new DateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *DateRange) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DateRange) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DateRange) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DateRange) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *DateRange) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DateRange) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DateRange) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *DateRange) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


