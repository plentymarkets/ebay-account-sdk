# CategorySuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**CategoryTreeNodeAncestors** | Pointer to [**[]AncestorReference**](AncestorReference.md) | An ordered list of category references that describes the location of the suggested category in the specified category tree. The list identifies the category&#39;s ancestry as a sequence of parent nodes, from the current node&#39;s immediate parent to the root node of the category tree. Note: The root node of a full default category tree includes categoryId and categoryName fields, but their values should not be relied upon. They provide no useful information for application development. | [optional] 
**CategoryTreeNodeLevel** | Pointer to **int32** | The absolute level of the category tree node in the hierarchy of its category tree. Note: The root node of any full category tree is always at level 0. | [optional] 
**Relevancy** | Pointer to **string** | This field is reserved for internal or future use. | [optional] 

## Methods

### NewCategorySuggestion

`func NewCategorySuggestion() *CategorySuggestion`

NewCategorySuggestion instantiates a new CategorySuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySuggestionWithDefaults

`func NewCategorySuggestionWithDefaults() *CategorySuggestion`

NewCategorySuggestionWithDefaults instantiates a new CategorySuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategorySuggestion) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategorySuggestion) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategorySuggestion) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CategorySuggestion) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryTreeNodeAncestors

`func (o *CategorySuggestion) GetCategoryTreeNodeAncestors() []AncestorReference`

GetCategoryTreeNodeAncestors returns the CategoryTreeNodeAncestors field if non-nil, zero value otherwise.

### GetCategoryTreeNodeAncestorsOk

`func (o *CategorySuggestion) GetCategoryTreeNodeAncestorsOk() (*[]AncestorReference, bool)`

GetCategoryTreeNodeAncestorsOk returns a tuple with the CategoryTreeNodeAncestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeNodeAncestors

`func (o *CategorySuggestion) SetCategoryTreeNodeAncestors(v []AncestorReference)`

SetCategoryTreeNodeAncestors sets CategoryTreeNodeAncestors field to given value.

### HasCategoryTreeNodeAncestors

`func (o *CategorySuggestion) HasCategoryTreeNodeAncestors() bool`

HasCategoryTreeNodeAncestors returns a boolean if a field has been set.

### GetCategoryTreeNodeLevel

`func (o *CategorySuggestion) GetCategoryTreeNodeLevel() int32`

GetCategoryTreeNodeLevel returns the CategoryTreeNodeLevel field if non-nil, zero value otherwise.

### GetCategoryTreeNodeLevelOk

`func (o *CategorySuggestion) GetCategoryTreeNodeLevelOk() (*int32, bool)`

GetCategoryTreeNodeLevelOk returns a tuple with the CategoryTreeNodeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeNodeLevel

`func (o *CategorySuggestion) SetCategoryTreeNodeLevel(v int32)`

SetCategoryTreeNodeLevel sets CategoryTreeNodeLevel field to given value.

### HasCategoryTreeNodeLevel

`func (o *CategorySuggestion) HasCategoryTreeNodeLevel() bool`

HasCategoryTreeNodeLevel returns a boolean if a field has been set.

### GetRelevancy

`func (o *CategorySuggestion) GetRelevancy() string`

GetRelevancy returns the Relevancy field if non-nil, zero value otherwise.

### GetRelevancyOk

`func (o *CategorySuggestion) GetRelevancyOk() (*string, bool)`

GetRelevancyOk returns a tuple with the Relevancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevancy

`func (o *CategorySuggestion) SetRelevancy(v string)`

SetRelevancy sets Relevancy field to given value.

### HasRelevancy

`func (o *CategorySuggestion) HasRelevancy() bool`

HasRelevancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


