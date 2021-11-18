# CategoryAspect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Aspects** | Pointer to [**[]Aspect**](Aspect.md) | A list of aspect metadata that is used to describe the items in a particular leaf category. | [optional] 

## Methods

### NewCategoryAspect

`func NewCategoryAspect() *CategoryAspect`

NewCategoryAspect instantiates a new CategoryAspect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryAspectWithDefaults

`func NewCategoryAspectWithDefaults() *CategoryAspect`

NewCategoryAspectWithDefaults instantiates a new CategoryAspect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategoryAspect) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategoryAspect) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategoryAspect) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CategoryAspect) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAspects

`func (o *CategoryAspect) GetAspects() []Aspect`

GetAspects returns the Aspects field if non-nil, zero value otherwise.

### GetAspectsOk

`func (o *CategoryAspect) GetAspectsOk() (*[]Aspect, bool)`

GetAspectsOk returns a tuple with the Aspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspects

`func (o *CategoryAspect) SetAspects(v []Aspect)`

SetAspects sets Aspects field to given value.

### HasAspects

`func (o *CategoryAspect) HasAspects() bool`

HasAspects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


