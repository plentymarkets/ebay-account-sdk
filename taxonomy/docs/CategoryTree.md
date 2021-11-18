# CategoryTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicableMarketplaceIds** | Pointer to **[]string** | A list of one or more identifiers of the eBay marketplaces that use this category tree. | [optional] 
**CategoryTreeId** | Pointer to **string** | The unique identifier of this eBay category tree. | [optional] 
**CategoryTreeVersion** | Pointer to **string** | The version of this category tree. It&#39;s a good idea to cache this value for comparison so you can determine if this category tree has been modified in subsequent calls. | [optional] 
**RootCategoryNode** | Pointer to [**CategoryTreeNode**](CategoryTreeNode.md) |  | [optional] 

## Methods

### NewCategoryTree

`func NewCategoryTree() *CategoryTree`

NewCategoryTree instantiates a new CategoryTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryTreeWithDefaults

`func NewCategoryTreeWithDefaults() *CategoryTree`

NewCategoryTreeWithDefaults instantiates a new CategoryTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicableMarketplaceIds

`func (o *CategoryTree) GetApplicableMarketplaceIds() []string`

GetApplicableMarketplaceIds returns the ApplicableMarketplaceIds field if non-nil, zero value otherwise.

### GetApplicableMarketplaceIdsOk

`func (o *CategoryTree) GetApplicableMarketplaceIdsOk() (*[]string, bool)`

GetApplicableMarketplaceIdsOk returns a tuple with the ApplicableMarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableMarketplaceIds

`func (o *CategoryTree) SetApplicableMarketplaceIds(v []string)`

SetApplicableMarketplaceIds sets ApplicableMarketplaceIds field to given value.

### HasApplicableMarketplaceIds

`func (o *CategoryTree) HasApplicableMarketplaceIds() bool`

HasApplicableMarketplaceIds returns a boolean if a field has been set.

### GetCategoryTreeId

`func (o *CategoryTree) GetCategoryTreeId() string`

GetCategoryTreeId returns the CategoryTreeId field if non-nil, zero value otherwise.

### GetCategoryTreeIdOk

`func (o *CategoryTree) GetCategoryTreeIdOk() (*string, bool)`

GetCategoryTreeIdOk returns a tuple with the CategoryTreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeId

`func (o *CategoryTree) SetCategoryTreeId(v string)`

SetCategoryTreeId sets CategoryTreeId field to given value.

### HasCategoryTreeId

`func (o *CategoryTree) HasCategoryTreeId() bool`

HasCategoryTreeId returns a boolean if a field has been set.

### GetCategoryTreeVersion

`func (o *CategoryTree) GetCategoryTreeVersion() string`

GetCategoryTreeVersion returns the CategoryTreeVersion field if non-nil, zero value otherwise.

### GetCategoryTreeVersionOk

`func (o *CategoryTree) GetCategoryTreeVersionOk() (*string, bool)`

GetCategoryTreeVersionOk returns a tuple with the CategoryTreeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeVersion

`func (o *CategoryTree) SetCategoryTreeVersion(v string)`

SetCategoryTreeVersion sets CategoryTreeVersion field to given value.

### HasCategoryTreeVersion

`func (o *CategoryTree) HasCategoryTreeVersion() bool`

HasCategoryTreeVersion returns a boolean if a field has been set.

### GetRootCategoryNode

`func (o *CategoryTree) GetRootCategoryNode() CategoryTreeNode`

GetRootCategoryNode returns the RootCategoryNode field if non-nil, zero value otherwise.

### GetRootCategoryNodeOk

`func (o *CategoryTree) GetRootCategoryNodeOk() (*CategoryTreeNode, bool)`

GetRootCategoryNodeOk returns a tuple with the RootCategoryNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCategoryNode

`func (o *CategoryTree) SetRootCategoryNode(v CategoryTreeNode)`

SetRootCategoryNode sets RootCategoryNode field to given value.

### HasRootCategoryNode

`func (o *CategoryTree) HasRootCategoryNode() bool`

HasRootCategoryNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


