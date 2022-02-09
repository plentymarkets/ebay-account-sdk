# CreateServiceMetricsTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedType** | Pointer to **string** | The feedType specified for the task. The report lists the transaction details that contribute to the service metrics evaluation. Supported types include: CUSTOMER_SERVICE_METRICS_REPORT | [optional] 
**FilterCriteria** | Pointer to [**CustomerServiceMetricsFilterCriteria**](CustomerServiceMetricsFilterCriteria.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** | The version number of the file format. Valid value: 1.0 | [optional] 

## Methods

### NewCreateServiceMetricsTaskRequest

`func NewCreateServiceMetricsTaskRequest() *CreateServiceMetricsTaskRequest`

NewCreateServiceMetricsTaskRequest instantiates a new CreateServiceMetricsTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceMetricsTaskRequestWithDefaults

`func NewCreateServiceMetricsTaskRequestWithDefaults() *CreateServiceMetricsTaskRequest`

NewCreateServiceMetricsTaskRequestWithDefaults instantiates a new CreateServiceMetricsTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedType

`func (o *CreateServiceMetricsTaskRequest) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *CreateServiceMetricsTaskRequest) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *CreateServiceMetricsTaskRequest) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *CreateServiceMetricsTaskRequest) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *CreateServiceMetricsTaskRequest) GetFilterCriteria() CustomerServiceMetricsFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *CreateServiceMetricsTaskRequest) GetFilterCriteriaOk() (*CustomerServiceMetricsFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *CreateServiceMetricsTaskRequest) SetFilterCriteria(v CustomerServiceMetricsFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *CreateServiceMetricsTaskRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateServiceMetricsTaskRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateServiceMetricsTaskRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateServiceMetricsTaskRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateServiceMetricsTaskRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


