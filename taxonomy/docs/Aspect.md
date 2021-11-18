# Aspect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectConstraint** | Pointer to [**AspectConstraint**](AspectConstraint.md) |  | [optional] 
**AspectValues** | Pointer to [**[]AspectValue**](AspectValue.md) | A list of valid values for this aspect (for example: Red, Green, and Blue), along with any constraints on those values. | [optional] 
**LocalizedAspectName** | Pointer to **string** | The localized name of this aspect (for example: Colour on the eBay UK site). Note: This name is always localized for the specified marketplace. | [optional] 
**RelevanceIndicator** | Pointer to [**RelevanceIndicator**](RelevanceIndicator.md) |  | [optional] 

## Methods

### NewAspect

`func NewAspect() *Aspect`

NewAspect instantiates a new Aspect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectWithDefaults

`func NewAspectWithDefaults() *Aspect`

NewAspectWithDefaults instantiates a new Aspect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectConstraint

`func (o *Aspect) GetAspectConstraint() AspectConstraint`

GetAspectConstraint returns the AspectConstraint field if non-nil, zero value otherwise.

### GetAspectConstraintOk

`func (o *Aspect) GetAspectConstraintOk() (*AspectConstraint, bool)`

GetAspectConstraintOk returns a tuple with the AspectConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectConstraint

`func (o *Aspect) SetAspectConstraint(v AspectConstraint)`

SetAspectConstraint sets AspectConstraint field to given value.

### HasAspectConstraint

`func (o *Aspect) HasAspectConstraint() bool`

HasAspectConstraint returns a boolean if a field has been set.

### GetAspectValues

`func (o *Aspect) GetAspectValues() []AspectValue`

GetAspectValues returns the AspectValues field if non-nil, zero value otherwise.

### GetAspectValuesOk

`func (o *Aspect) GetAspectValuesOk() (*[]AspectValue, bool)`

GetAspectValuesOk returns a tuple with the AspectValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectValues

`func (o *Aspect) SetAspectValues(v []AspectValue)`

SetAspectValues sets AspectValues field to given value.

### HasAspectValues

`func (o *Aspect) HasAspectValues() bool`

HasAspectValues returns a boolean if a field has been set.

### GetLocalizedAspectName

`func (o *Aspect) GetLocalizedAspectName() string`

GetLocalizedAspectName returns the LocalizedAspectName field if non-nil, zero value otherwise.

### GetLocalizedAspectNameOk

`func (o *Aspect) GetLocalizedAspectNameOk() (*string, bool)`

GetLocalizedAspectNameOk returns a tuple with the LocalizedAspectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedAspectName

`func (o *Aspect) SetLocalizedAspectName(v string)`

SetLocalizedAspectName sets LocalizedAspectName field to given value.

### HasLocalizedAspectName

`func (o *Aspect) HasLocalizedAspectName() bool`

HasLocalizedAspectName returns a boolean if a field has been set.

### GetRelevanceIndicator

`func (o *Aspect) GetRelevanceIndicator() RelevanceIndicator`

GetRelevanceIndicator returns the RelevanceIndicator field if non-nil, zero value otherwise.

### GetRelevanceIndicatorOk

`func (o *Aspect) GetRelevanceIndicatorOk() (*RelevanceIndicator, bool)`

GetRelevanceIndicatorOk returns a tuple with the RelevanceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceIndicator

`func (o *Aspect) SetRelevanceIndicator(v RelevanceIndicator)`

SetRelevanceIndicator sets RelevanceIndicator field to given value.

### HasRelevanceIndicator

`func (o *Aspect) HasRelevanceIndicator() bool`

HasRelevanceIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


