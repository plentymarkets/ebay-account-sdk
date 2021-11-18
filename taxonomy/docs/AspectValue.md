# AspectValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedValue** | Pointer to **string** | The localized value of this aspect. Note: This value is always localized for the specified marketplace. | [optional] 
**ValueConstraints** | Pointer to [**[]ValueConstraint**](ValueConstraint.md) | Not returned if the value of the localizedValue field can always be selected for this aspect of the specified category. Contains a list of the dependencies that identify when the value of the localizedValue field is available for the current aspect. Each dependency specifies the values of another aspect of the same category (a control aspect), for which the current value of the current aspect can also be selected by the seller. Example: A shirt is available in three sizes and three colors, but only the Small and Medium sizes come in Green. Thus for the Color aspect, the value Green is constrained by its dependency on Size (the control aspect). Only when the Size aspect value is Small or Medium, can the Color aspect value of Green be selected by the seller. | [optional] 

## Methods

### NewAspectValue

`func NewAspectValue() *AspectValue`

NewAspectValue instantiates a new AspectValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectValueWithDefaults

`func NewAspectValueWithDefaults() *AspectValue`

NewAspectValueWithDefaults instantiates a new AspectValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedValue

`func (o *AspectValue) GetLocalizedValue() string`

GetLocalizedValue returns the LocalizedValue field if non-nil, zero value otherwise.

### GetLocalizedValueOk

`func (o *AspectValue) GetLocalizedValueOk() (*string, bool)`

GetLocalizedValueOk returns a tuple with the LocalizedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedValue

`func (o *AspectValue) SetLocalizedValue(v string)`

SetLocalizedValue sets LocalizedValue field to given value.

### HasLocalizedValue

`func (o *AspectValue) HasLocalizedValue() bool`

HasLocalizedValue returns a boolean if a field has been set.

### GetValueConstraints

`func (o *AspectValue) GetValueConstraints() []ValueConstraint`

GetValueConstraints returns the ValueConstraints field if non-nil, zero value otherwise.

### GetValueConstraintsOk

`func (o *AspectValue) GetValueConstraintsOk() (*[]ValueConstraint, bool)`

GetValueConstraintsOk returns a tuple with the ValueConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueConstraints

`func (o *AspectValue) SetValueConstraints(v []ValueConstraint)`

SetValueConstraints sets ValueConstraints field to given value.

### HasValueConstraints

`func (o *AspectValue) HasValueConstraints() bool`

HasValueConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


