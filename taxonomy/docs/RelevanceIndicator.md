# RelevanceIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchCount** | Pointer to **int32** | The number of recent searches (based on 30 days of data) for the aspect. | [optional] 

## Methods

### NewRelevanceIndicator

`func NewRelevanceIndicator() *RelevanceIndicator`

NewRelevanceIndicator instantiates a new RelevanceIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelevanceIndicatorWithDefaults

`func NewRelevanceIndicatorWithDefaults() *RelevanceIndicator`

NewRelevanceIndicatorWithDefaults instantiates a new RelevanceIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchCount

`func (o *RelevanceIndicator) GetSearchCount() int32`

GetSearchCount returns the SearchCount field if non-nil, zero value otherwise.

### GetSearchCountOk

`func (o *RelevanceIndicator) GetSearchCountOk() (*int32, bool)`

GetSearchCountOk returns a tuple with the SearchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCount

`func (o *RelevanceIndicator) SetSearchCount(v int32)`

SetSearchCount sets SearchCount field to given value.

### HasSearchCount

`func (o *RelevanceIndicator) HasSearchCount() bool`

HasSearchCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


