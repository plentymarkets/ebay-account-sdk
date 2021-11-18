# GetCompatibilityMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityProperties** | Pointer to [**[]CompatibilityProperty**](CompatibilityProperty.md) | This container consists of an array of all compatible vehicle properties applicable to the specified eBay marketplace and eBay category ID. | [optional] 

## Methods

### NewGetCompatibilityMetadataResponse

`func NewGetCompatibilityMetadataResponse() *GetCompatibilityMetadataResponse`

NewGetCompatibilityMetadataResponse instantiates a new GetCompatibilityMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompatibilityMetadataResponseWithDefaults

`func NewGetCompatibilityMetadataResponseWithDefaults() *GetCompatibilityMetadataResponse`

NewGetCompatibilityMetadataResponseWithDefaults instantiates a new GetCompatibilityMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityProperties

`func (o *GetCompatibilityMetadataResponse) GetCompatibilityProperties() []CompatibilityProperty`

GetCompatibilityProperties returns the CompatibilityProperties field if non-nil, zero value otherwise.

### GetCompatibilityPropertiesOk

`func (o *GetCompatibilityMetadataResponse) GetCompatibilityPropertiesOk() (*[]CompatibilityProperty, bool)`

GetCompatibilityPropertiesOk returns a tuple with the CompatibilityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityProperties

`func (o *GetCompatibilityMetadataResponse) SetCompatibilityProperties(v []CompatibilityProperty)`

SetCompatibilityProperties sets CompatibilityProperties field to given value.

### HasCompatibilityProperties

`func (o *GetCompatibilityMetadataResponse) HasCompatibilityProperties() bool`

HasCompatibilityProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


