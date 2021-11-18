# BaseCategoryTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTreeId** | Pointer to **string** | The unique identifier of the eBay category tree for the specified marketplace. | [optional] 
**CategoryTreeVersion** | Pointer to **string** | The version of the category tree identified by categoryTreeId. It&#39;s a good idea to cache this value for comparison so you can determine if this category tree has been modified in subsequent calls. | [optional] 

## Methods

### NewBaseCategoryTree

`func NewBaseCategoryTree() *BaseCategoryTree`

NewBaseCategoryTree instantiates a new BaseCategoryTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCategoryTreeWithDefaults

`func NewBaseCategoryTreeWithDefaults() *BaseCategoryTree`

NewBaseCategoryTreeWithDefaults instantiates a new BaseCategoryTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTreeId

`func (o *BaseCategoryTree) GetCategoryTreeId() string`

GetCategoryTreeId returns the CategoryTreeId field if non-nil, zero value otherwise.

### GetCategoryTreeIdOk

`func (o *BaseCategoryTree) GetCategoryTreeIdOk() (*string, bool)`

GetCategoryTreeIdOk returns a tuple with the CategoryTreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeId

`func (o *BaseCategoryTree) SetCategoryTreeId(v string)`

SetCategoryTreeId sets CategoryTreeId field to given value.

### HasCategoryTreeId

`func (o *BaseCategoryTree) HasCategoryTreeId() bool`

HasCategoryTreeId returns a boolean if a field has been set.

### GetCategoryTreeVersion

`func (o *BaseCategoryTree) GetCategoryTreeVersion() string`

GetCategoryTreeVersion returns the CategoryTreeVersion field if non-nil, zero value otherwise.

### GetCategoryTreeVersionOk

`func (o *BaseCategoryTree) GetCategoryTreeVersionOk() (*string, bool)`

GetCategoryTreeVersionOk returns a tuple with the CategoryTreeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeVersion

`func (o *BaseCategoryTree) SetCategoryTreeVersion(v string)`

SetCategoryTreeVersion sets CategoryTreeVersion field to given value.

### HasCategoryTreeVersion

`func (o *BaseCategoryTree) HasCategoryTreeVersion() bool`

HasCategoryTreeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


