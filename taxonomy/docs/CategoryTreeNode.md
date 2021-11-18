# CategoryTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**CategoryTreeNodeLevel** | Pointer to **int32** | The absolute level of the current category tree node in the hierarchy of its category tree. Note: The root node of any full category tree is always at level 0. | [optional] 
**ChildCategoryTreeNodes** | Pointer to [**[]CategoryTreeNode**](CategoryTreeNode.md) | An array of one or more category tree nodes that are the immediate children of the current category tree node, as well as their children, recursively down to the leaf nodes. Returned only if the current category tree node is not a leaf node (the value of leafCategoryTreeNode is false). | [optional] 
**LeafCategoryTreeNode** | Pointer to **bool** | A value of true indicates that the current category tree node is a leaf node (it has no child nodes). A value of false indicates that the current node has one or more child nodes, which are identified by the childCategoryTreeNodes array. Returned only if the value of this field is true. | [optional] 
**ParentCategoryTreeNodeHref** | Pointer to **string** | The href portion of the getCategorySubtree call that retrieves the subtree below the parent of this category tree node. Not returned if the current category tree node is the root node of its tree. | [optional] 

## Methods

### NewCategoryTreeNode

`func NewCategoryTreeNode() *CategoryTreeNode`

NewCategoryTreeNode instantiates a new CategoryTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryTreeNodeWithDefaults

`func NewCategoryTreeNodeWithDefaults() *CategoryTreeNode`

NewCategoryTreeNodeWithDefaults instantiates a new CategoryTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategoryTreeNode) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategoryTreeNode) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategoryTreeNode) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CategoryTreeNode) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryTreeNodeLevel

`func (o *CategoryTreeNode) GetCategoryTreeNodeLevel() int32`

GetCategoryTreeNodeLevel returns the CategoryTreeNodeLevel field if non-nil, zero value otherwise.

### GetCategoryTreeNodeLevelOk

`func (o *CategoryTreeNode) GetCategoryTreeNodeLevelOk() (*int32, bool)`

GetCategoryTreeNodeLevelOk returns a tuple with the CategoryTreeNodeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeNodeLevel

`func (o *CategoryTreeNode) SetCategoryTreeNodeLevel(v int32)`

SetCategoryTreeNodeLevel sets CategoryTreeNodeLevel field to given value.

### HasCategoryTreeNodeLevel

`func (o *CategoryTreeNode) HasCategoryTreeNodeLevel() bool`

HasCategoryTreeNodeLevel returns a boolean if a field has been set.

### GetChildCategoryTreeNodes

`func (o *CategoryTreeNode) GetChildCategoryTreeNodes() []CategoryTreeNode`

GetChildCategoryTreeNodes returns the ChildCategoryTreeNodes field if non-nil, zero value otherwise.

### GetChildCategoryTreeNodesOk

`func (o *CategoryTreeNode) GetChildCategoryTreeNodesOk() (*[]CategoryTreeNode, bool)`

GetChildCategoryTreeNodesOk returns a tuple with the ChildCategoryTreeNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCategoryTreeNodes

`func (o *CategoryTreeNode) SetChildCategoryTreeNodes(v []CategoryTreeNode)`

SetChildCategoryTreeNodes sets ChildCategoryTreeNodes field to given value.

### HasChildCategoryTreeNodes

`func (o *CategoryTreeNode) HasChildCategoryTreeNodes() bool`

HasChildCategoryTreeNodes returns a boolean if a field has been set.

### GetLeafCategoryTreeNode

`func (o *CategoryTreeNode) GetLeafCategoryTreeNode() bool`

GetLeafCategoryTreeNode returns the LeafCategoryTreeNode field if non-nil, zero value otherwise.

### GetLeafCategoryTreeNodeOk

`func (o *CategoryTreeNode) GetLeafCategoryTreeNodeOk() (*bool, bool)`

GetLeafCategoryTreeNodeOk returns a tuple with the LeafCategoryTreeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCategoryTreeNode

`func (o *CategoryTreeNode) SetLeafCategoryTreeNode(v bool)`

SetLeafCategoryTreeNode sets LeafCategoryTreeNode field to given value.

### HasLeafCategoryTreeNode

`func (o *CategoryTreeNode) HasLeafCategoryTreeNode() bool`

HasLeafCategoryTreeNode returns a boolean if a field has been set.

### GetParentCategoryTreeNodeHref

`func (o *CategoryTreeNode) GetParentCategoryTreeNodeHref() string`

GetParentCategoryTreeNodeHref returns the ParentCategoryTreeNodeHref field if non-nil, zero value otherwise.

### GetParentCategoryTreeNodeHrefOk

`func (o *CategoryTreeNode) GetParentCategoryTreeNodeHrefOk() (*string, bool)`

GetParentCategoryTreeNodeHrefOk returns a tuple with the ParentCategoryTreeNodeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategoryTreeNodeHref

`func (o *CategoryTreeNode) SetParentCategoryTreeNodeHref(v string)`

SetParentCategoryTreeNodeHref sets ParentCategoryTreeNodeHref field to given value.

### HasParentCategoryTreeNodeHref

`func (o *CategoryTreeNode) HasParentCategoryTreeNodeHref() bool`

HasParentCategoryTreeNodeHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


