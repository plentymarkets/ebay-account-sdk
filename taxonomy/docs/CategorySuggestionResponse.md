# CategorySuggestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorySuggestions** | Pointer to [**[]CategorySuggestion**](CategorySuggestion.md) | Contains details about one or more suggested categories that correspond to the provided keywords. The array of suggested categories is sorted in order of eBay&#39;s confidence of the relevance of each category (the first category is the most relevant). Important: This call is not supported in the Sandbox environment. It will return a response payload in which the categoryName fields contain random or boilerplate text regardless of the query submitted. | [optional] 
**CategoryTreeId** | Pointer to **string** | The unique identifier of the eBay category tree from which suggestions are returned. | [optional] 
**CategoryTreeVersion** | Pointer to **string** | The version of the category tree identified by categoryTreeId. It&#39;s a good idea to cache this value for comparison so you can determine if this category tree has been modified in subsequent calls. | [optional] 

## Methods

### NewCategorySuggestionResponse

`func NewCategorySuggestionResponse() *CategorySuggestionResponse`

NewCategorySuggestionResponse instantiates a new CategorySuggestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorySuggestionResponseWithDefaults

`func NewCategorySuggestionResponseWithDefaults() *CategorySuggestionResponse`

NewCategorySuggestionResponseWithDefaults instantiates a new CategorySuggestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorySuggestions

`func (o *CategorySuggestionResponse) GetCategorySuggestions() []CategorySuggestion`

GetCategorySuggestions returns the CategorySuggestions field if non-nil, zero value otherwise.

### GetCategorySuggestionsOk

`func (o *CategorySuggestionResponse) GetCategorySuggestionsOk() (*[]CategorySuggestion, bool)`

GetCategorySuggestionsOk returns a tuple with the CategorySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySuggestions

`func (o *CategorySuggestionResponse) SetCategorySuggestions(v []CategorySuggestion)`

SetCategorySuggestions sets CategorySuggestions field to given value.

### HasCategorySuggestions

`func (o *CategorySuggestionResponse) HasCategorySuggestions() bool`

HasCategorySuggestions returns a boolean if a field has been set.

### GetCategoryTreeId

`func (o *CategorySuggestionResponse) GetCategoryTreeId() string`

GetCategoryTreeId returns the CategoryTreeId field if non-nil, zero value otherwise.

### GetCategoryTreeIdOk

`func (o *CategorySuggestionResponse) GetCategoryTreeIdOk() (*string, bool)`

GetCategoryTreeIdOk returns a tuple with the CategoryTreeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeId

`func (o *CategorySuggestionResponse) SetCategoryTreeId(v string)`

SetCategoryTreeId sets CategoryTreeId field to given value.

### HasCategoryTreeId

`func (o *CategorySuggestionResponse) HasCategoryTreeId() bool`

HasCategoryTreeId returns a boolean if a field has been set.

### GetCategoryTreeVersion

`func (o *CategorySuggestionResponse) GetCategoryTreeVersion() string`

GetCategoryTreeVersion returns the CategoryTreeVersion field if non-nil, zero value otherwise.

### GetCategoryTreeVersionOk

`func (o *CategorySuggestionResponse) GetCategoryTreeVersionOk() (*string, bool)`

GetCategoryTreeVersionOk returns a tuple with the CategoryTreeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryTreeVersion

`func (o *CategorySuggestionResponse) SetCategoryTreeVersion(v string)`

SetCategoryTreeVersion sets CategoryTreeVersion field to given value.

### HasCategoryTreeVersion

`func (o *CategorySuggestionResponse) HasCategoryTreeVersion() bool`

HasCategoryTreeVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


