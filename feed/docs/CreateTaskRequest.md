# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedType** | Pointer to **string** | The feed type associated with the task. Only use a feedType that is available for your API. Available feed types: LMS FeedTypes: All LMS feed types except LMS_ORDER_REPORT and LMS_ACTIVE_INVENTORY_REPORT File Exchange FeedTypes | [optional] 
**SchemaVersion** | Pointer to **string** | The schemaVersion/version number of the file format (use the schema version of the API to which you are programming): LMS Version Details / Schema Version File Exchange Schema Version | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest() *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedType

`func (o *CreateTaskRequest) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *CreateTaskRequest) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *CreateTaskRequest) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *CreateTaskRequest) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateTaskRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateTaskRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateTaskRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateTaskRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


