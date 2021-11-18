# ValueConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicableForLocalizedAspectName** | Pointer to **string** | The name of the control aspect on which the current aspect value depends. | [optional] 
**ApplicableForLocalizedAspectValues** | Pointer to **[]string** | Contains a list of the values of the control aspect on which this aspect&#39;s value depends. When the control aspect has any of the specified values, the current value of the current aspect will also be available. | [optional] 

## Methods

### NewValueConstraint

`func NewValueConstraint() *ValueConstraint`

NewValueConstraint instantiates a new ValueConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueConstraintWithDefaults

`func NewValueConstraintWithDefaults() *ValueConstraint`

NewValueConstraintWithDefaults instantiates a new ValueConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicableForLocalizedAspectName

`func (o *ValueConstraint) GetApplicableForLocalizedAspectName() string`

GetApplicableForLocalizedAspectName returns the ApplicableForLocalizedAspectName field if non-nil, zero value otherwise.

### GetApplicableForLocalizedAspectNameOk

`func (o *ValueConstraint) GetApplicableForLocalizedAspectNameOk() (*string, bool)`

GetApplicableForLocalizedAspectNameOk returns a tuple with the ApplicableForLocalizedAspectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableForLocalizedAspectName

`func (o *ValueConstraint) SetApplicableForLocalizedAspectName(v string)`

SetApplicableForLocalizedAspectName sets ApplicableForLocalizedAspectName field to given value.

### HasApplicableForLocalizedAspectName

`func (o *ValueConstraint) HasApplicableForLocalizedAspectName() bool`

HasApplicableForLocalizedAspectName returns a boolean if a field has been set.

### GetApplicableForLocalizedAspectValues

`func (o *ValueConstraint) GetApplicableForLocalizedAspectValues() []string`

GetApplicableForLocalizedAspectValues returns the ApplicableForLocalizedAspectValues field if non-nil, zero value otherwise.

### GetApplicableForLocalizedAspectValuesOk

`func (o *ValueConstraint) GetApplicableForLocalizedAspectValuesOk() (*[]string, bool)`

GetApplicableForLocalizedAspectValuesOk returns a tuple with the ApplicableForLocalizedAspectValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableForLocalizedAspectValues

`func (o *ValueConstraint) SetApplicableForLocalizedAspectValues(v []string)`

SetApplicableForLocalizedAspectValues sets ApplicableForLocalizedAspectValues field to given value.

### HasApplicableForLocalizedAspectValues

`func (o *ValueConstraint) HasApplicableForLocalizedAspectValues() bool`

HasApplicableForLocalizedAspectValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


