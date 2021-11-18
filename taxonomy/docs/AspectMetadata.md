# AspectMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aspects** | Pointer to [**[]Aspect**](Aspect.md) | A list of item aspects (for example, color) that are appropriate or necessary for accurately describing items in a particular leaf category. Each category has a different set of aspects and different requirements for aspect values. Sellers are required or encouraged to provide one or more acceptable values for each aspect when offering an item in that category on eBay. | [optional] 

## Methods

### NewAspectMetadata

`func NewAspectMetadata() *AspectMetadata`

NewAspectMetadata instantiates a new AspectMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectMetadataWithDefaults

`func NewAspectMetadataWithDefaults() *AspectMetadata`

NewAspectMetadataWithDefaults instantiates a new AspectMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspects

`func (o *AspectMetadata) GetAspects() []Aspect`

GetAspects returns the Aspects field if non-nil, zero value otherwise.

### GetAspectsOk

`func (o *AspectMetadata) GetAspectsOk() (*[]Aspect, bool)`

GetAspectsOk returns a tuple with the Aspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspects

`func (o *AspectMetadata) SetAspects(v []Aspect)`

SetAspects sets Aspects field to given value.

### HasAspects

`func (o *AspectMetadata) HasAspects() bool`

HasAspects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


