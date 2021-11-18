# AncestorReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | The unique identifier of the eBay ancestor category. Note: The root node of a full default category tree includes the categoryId field, but its value should not be relied upon. It provides no useful information for application development. | [optional] 
**CategoryName** | Pointer to **string** | The name of the ancestor category identified by categoryId. | [optional] 
**CategorySubtreeNodeHref** | Pointer to **string** | The href portion of the getCategorySubtree call that retrieves the subtree below the ancestor category node. | [optional] 
**CategoryTreeNodeLevel** | Pointer to **int32** | The absolute level of the ancestor category node in the hierarchy of its category tree. Note: The root node of any full category tree is always at level 0. | [optional] 

## Methods

### NewAncestorReference

`func NewAncestorReference() *AncestorReference`

NewAncestorReference instantiates a new AncestorReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAncestorReferenceWithDefaults

`func NewAncestorReferenceWithDefaults() *AncestorReference`

NewAncestorReferenceWithDefaults instantiates a new AncestorReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *AncestorReference) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *AncestorReference) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *AncestorReference) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *AncestorReference) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryName

`func (o *AncestorReference) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *AncestorReference) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *AncestorReference) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *AncestorReference) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetCategorySubtreeNodeHref

`func (o *AncestorReference) GetCategorySubtreeNodeHref() string`

GetCategorySubtreeNodeHref returns the CategorySubtreeNodeHref field if non-nil, zero value otherwise.

### GetCategorySubtreeNodeHrefOk

`func (o *AncestorReference) GetCategorySubtreeNodeHrefOk() (*string, bool)`

GetCategorySubtreeNodeHrefOk returns a tuple with the CategorySubtreeNodeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySubtreeNodeHref

`func (o *AncestorReference) SetCategorySubtreeNodeHref(v string)`

SetCategorySubtreeNodeHref sets CategorySubtreeNodeHref field to given value.

### HasCategorySubtreeNodeHref

`func (o *AncestorReference) HasCategorySubtreeNodeHref() bool`

HasCategorySubtreeNodeHref returns a boolean if a field has been set.

### GetCategoryTreeNodeLevel

`func (o *AncestorReference) GetCategoryTreeNodeLevel() int32`

GetCategoryTreeNodeLevel returns the CategoryTreeNodeLevel field if non-nil, zero value otherwise.

### GetCategoryTreeNodeLevelOk

`func (o *AncestorReference) GetCategoryTreeNodeLevelOk() (*int32, bool)`

GetCategoryTreeNodeLevelOk returns a tuple with the CategoryTreeNodeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeNodeLevel

`func (o *AncestorReference) SetCategoryTreeNodeLevel(v int32)`

SetCategoryTreeNodeLevel sets CategoryTreeNodeLevel field to given value.

### HasCategoryTreeNodeLevel

`func (o *AncestorReference) HasCategoryTreeNodeLevel() bool`

HasCategoryTreeNodeLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


