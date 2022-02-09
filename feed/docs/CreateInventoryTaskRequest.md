# CreateInventoryTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | The schemaVersion/version number of the file format (use the schema version of the API to which you are programming): LMS Version Details / Schema Version File Exchange Schema Version | [optional] 
**FeedType** | Pointer to **string** | The feed type associated with the inventory task you are about to create. Use a feedType that is available for your API. Presently, only one feed type is available: LMS_ACTIVE_INVENTORY_REPORT | [optional] 
**FilterCriteria** | Pointer to [**InventoryFilterCriteria**](InventoryFilterCriteria.md) |  | [optional] 
**InventoryFileTemplate** | Pointer to **string** | The inventory file template used to return specific types of inventory tasks. Presently not applicable for LMS_ACTIVE_INVENTORY_REPORT. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/feed/types/api:InventoryFileTemplateEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewCreateInventoryTaskRequest

`func NewCreateInventoryTaskRequest() *CreateInventoryTaskRequest`

NewCreateInventoryTaskRequest instantiates a new CreateInventoryTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryTaskRequestWithDefaults

`func NewCreateInventoryTaskRequestWithDefaults() *CreateInventoryTaskRequest`

NewCreateInventoryTaskRequestWithDefaults instantiates a new CreateInventoryTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CreateInventoryTaskRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateInventoryTaskRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateInventoryTaskRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateInventoryTaskRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetFeedType

`func (o *CreateInventoryTaskRequest) GetFeedType() string`

GetFeedType returns the FeedType field if non-nil, zero value otherwise.

### GetFeedTypeOk

`func (o *CreateInventoryTaskRequest) GetFeedTypeOk() (*string, bool)`

GetFeedTypeOk returns a tuple with the FeedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedType

`func (o *CreateInventoryTaskRequest) SetFeedType(v string)`

SetFeedType sets FeedType field to given value.

### HasFeedType

`func (o *CreateInventoryTaskRequest) HasFeedType() bool`

HasFeedType returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *CreateInventoryTaskRequest) GetFilterCriteria() InventoryFilterCriteria`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *CreateInventoryTaskRequest) GetFilterCriteriaOk() (*InventoryFilterCriteria, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *CreateInventoryTaskRequest) SetFilterCriteria(v InventoryFilterCriteria)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *CreateInventoryTaskRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetInventoryFileTemplate

`func (o *CreateInventoryTaskRequest) GetInventoryFileTemplate() string`

GetInventoryFileTemplate returns the InventoryFileTemplate field if non-nil, zero value otherwise.

### GetInventoryFileTemplateOk

`func (o *CreateInventoryTaskRequest) GetInventoryFileTemplateOk() (*string, bool)`

GetInventoryFileTemplateOk returns a tuple with the InventoryFileTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryFileTemplate

`func (o *CreateInventoryTaskRequest) SetInventoryFileTemplate(v string)`

SetInventoryFileTemplate sets InventoryFileTemplate field to given value.

### HasInventoryFileTemplate

`func (o *CreateInventoryTaskRequest) HasInventoryFileTemplate() bool`

HasInventoryFileTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


