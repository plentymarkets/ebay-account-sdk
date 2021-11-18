# CategorySubtree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorySubtreeNode** | Pointer to [**CategoryTreeNode**](CategoryTreeNode.md) |  | [optional] 
**CategoryTreeId** | Pointer to **string** | The unique identifier of the eBay category tree to which this subtree belongs. | [optional] 
**CategoryTreeVersion** | Pointer to **string** | The version of the category tree identified by categoryTreeId. It&#39;s a good idea to cache this value for comparison so you can determine if this category tree has been modified in subsequent calls. | [optional] 

## Methods

### NewCategorySubtree

`func NewCategorySubtree() *CategorySubtree`

NewCategorySubtree instantiates a new CategorySubtree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySubtreeWithDefaults

`func NewCategorySubtreeWithDefaults() *CategorySubtree`

NewCategorySubtreeWithDefaults instantiates a new CategorySubtree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorySubtreeNode

`func (o *CategorySubtree) GetCategorySubtreeNode() CategoryTreeNode`

GetCategorySubtreeNode returns the CategorySubtreeNode field if non-nil, zero value otherwise.

### GetCategorySubtreeNodeOk

`func (o *CategorySubtree) GetCategorySubtreeNodeOk() (*CategoryTreeNode, bool)`

GetCategorySubtreeNodeOk returns a tuple with the CategorySubtreeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySubtreeNode

`func (o *CategorySubtree) SetCategorySubtreeNode(v CategoryTreeNode)`

SetCategorySubtreeNode sets CategorySubtreeNode field to given value.

### HasCategorySubtreeNode

`func (o *CategorySubtree) HasCategorySubtreeNode() bool`

HasCategorySubtreeNode returns a boolean if a field has been set.

### GetCategoryTreeId

`func (o *CategorySubtree) GetCategoryTreeId() string`

GetCategoryTreeId returns the CategoryTreeId field if non-nil, zero value otherwise.

### GetCategoryTreeIdOk

`func (o *CategorySubtree) GetCategoryTreeIdOk() (*string, bool)`

GetCategoryTreeIdOk returns a tuple with the CategoryTreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeId

`func (o *CategorySubtree) SetCategoryTreeId(v string)`

SetCategoryTreeId sets CategoryTreeId field to given value.

### HasCategoryTreeId

`func (o *CategorySubtree) HasCategoryTreeId() bool`

HasCategoryTreeId returns a boolean if a field has been set.

### GetCategoryTreeVersion

`func (o *CategorySubtree) GetCategoryTreeVersion() string`

GetCategoryTreeVersion returns the CategoryTreeVersion field if non-nil, zero value otherwise.

### GetCategoryTreeVersionOk

`func (o *CategorySubtree) GetCategoryTreeVersionOk() (*string, bool)`

GetCategoryTreeVersionOk returns a tuple with the CategoryTreeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeVersion

`func (o *CategorySubtree) SetCategoryTreeVersion(v string)`

SetCategoryTreeVersion sets CategoryTreeVersion field to given value.

### HasCategoryTreeVersion

`func (o *CategorySubtree) HasCategoryTreeVersion() bool`

HasCategoryTreeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


