# CreateOrderTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedType** | Pointer to **string** | The feed type associated with the task. The only presently supported value is LMS_ORDER_REPORT. | [optional] 
**FilterCriteria** | Pointer to [**OrderFilterCriteria**](OrderFilterCriteria.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** | The schema version of the LMS OrderReport. For the LMS_ORDER_REPORT feed type, see the OrderReport reference page to see the present schema version. The schemaVersion value is the version number shown at the top of the OrderReport page. Restriction: This value must be 1113 or higher. The OrderReport schema version is updated about every two weeks. All version numbers are odd numbers (even numbers are skipped). For example, the next release version after &#39;1113&#39; is &#39;1115&#39;. | [optional] 

## Methods

### NewCreateOrderTaskRequest

`func NewCreateOrderTaskRequest() *CreateOrderTaskRequest`

NewCreateOrderTaskRequest instantiates a new CreateOrderTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderTaskRequestWithDefaults

`func NewCreateOrderTaskRequestWithDefaults() *CreateOrderTaskRequest`

NewCreateOrderTaskRequestWithDefaults instantiates a new CreateOrderTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedType

`func (o *CreateOrderTaskRequest) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *CreateOrderTaskRequest) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *CreateOrderTaskRequest) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *CreateOrderTaskRequest) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *CreateOrderTaskRequest) GetFilterCriteria() OrderFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *CreateOrderTaskRequest) GetFilterCriteriaOk() (*OrderFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *CreateOrderTaskRequest) SetFilterCriteria(v OrderFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *CreateOrderTaskRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateOrderTaskRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateOrderTaskRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateOrderTaskRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateOrderTaskRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


