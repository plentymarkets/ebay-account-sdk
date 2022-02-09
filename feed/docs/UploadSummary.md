# UploadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCount** | Pointer to **int32** | The number of records, such as the number of listings created or the number of pictures uploaded to a listing, that failed to process during the upload feed. Check the response file and correct any issues mentioned. If the feed fails before processing, no response file is provided. In this case check the REST output response. | [optional] 
**SuccessCount** | Pointer to **int32** | The number of records that were successfully processed during the upload feed. | [optional] 

## Methods

### NewUploadSummary

`func NewUploadSummary() *UploadSummary`

NewUploadSummary instantiates a new UploadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSummaryWithDefaults

`func NewUploadSummaryWithDefaults() *UploadSummary`

NewUploadSummaryWithDefaults instantiates a new UploadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCount

`func (o *UploadSummary) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *UploadSummary) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *UploadSummary) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *UploadSummary) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UploadSummary) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UploadSummary) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UploadSummary) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UploadSummary) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


