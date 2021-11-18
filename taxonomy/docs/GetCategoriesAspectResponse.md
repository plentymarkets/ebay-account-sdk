# GetCategoriesAspectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryTreeId** | Pointer to **string** | The unique identifier of the eBay category tree being requested. | [optional] 
**CategoryTreeVersion** | Pointer to **string** | The version of the category tree that is returned in the categoryTreeId field. | [optional] 
**CategoryAspects** | Pointer to [**[]CategoryAspect**](CategoryAspect.md) | An array of aspects that are appropriate or necessary for accurately describing items in a particular leaf category. | [optional] 

## Methods

### NewGetCategoriesAspectResponse

`func NewGetCategoriesAspectResponse() *GetCategoriesAspectResponse`

NewGetCategoriesAspectResponse instantiates a new GetCategoriesAspectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCategoriesAspectResponseWithDefaults

`func NewGetCategoriesAspectResponseWithDefaults() *GetCategoriesAspectResponse`

NewGetCategoriesAspectResponseWithDefaults instantiates a new GetCategoriesAspectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryTreeId

`func (o *GetCategoriesAspectResponse) GetCategoryTreeId() string`

GetCategoryTreeId returns the CategoryTreeId field if non-nil, zero value otherwise.

### GetCategoryTreeIdOk

`func (o *GetCategoriesAspectResponse) GetCategoryTreeIdOk() (*string, bool)`

GetCategoryTreeIdOk returns a tuple with the CategoryTreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeId

`func (o *GetCategoriesAspectResponse) SetCategoryTreeId(v string)`

SetCategoryTreeId sets CategoryTreeId field to given value.

### HasCategoryTreeId

`func (o *GetCategoriesAspectResponse) HasCategoryTreeId() bool`

HasCategoryTreeId returns a boolean if a field has been set.

### GetCategoryTreeVersion

`func (o *GetCategoriesAspectResponse) GetCategoryTreeVersion() string`

GetCategoryTreeVersion returns the CategoryTreeVersion field if non-nil, zero value otherwise.

### GetCategoryTreeVersionOk

`func (o *GetCategoriesAspectResponse) GetCategoryTreeVersionOk() (*string, bool)`

GetCategoryTreeVersionOk returns a tuple with the CategoryTreeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeVersion

`func (o *GetCategoriesAspectResponse) SetCategoryTreeVersion(v string)`

SetCategoryTreeVersion sets CategoryTreeVersion field to given value.

### HasCategoryTreeVersion

`func (o *GetCategoriesAspectResponse) HasCategoryTreeVersion() bool`

HasCategoryTreeVersion returns a boolean if a field has been set.

### GetCategoryAspects

`func (o *GetCategoriesAspectResponse) GetCategoryAspects() []CategoryAspect`

GetCategoryAspects returns the CategoryAspects field if non-nil, zero value otherwise.

### GetCategoryAspectsOk

`func (o *GetCategoriesAspectResponse) GetCategoryAspectsOk() (*[]CategoryAspect, bool)`

GetCategoryAspectsOk returns a tuple with the CategoryAspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryAspects

`func (o *GetCategoriesAspectResponse) SetCategoryAspects(v []CategoryAspect)`

SetCategoryAspects sets CategoryAspects field to given value.

### HasCategoryAspects

`func (o *GetCategoriesAspectResponse) HasCategoryAspects() bool`

HasCategoryAspects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


