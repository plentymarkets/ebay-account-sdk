# AspectConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectApplicableTo** | Pointer to **[]string** | This value indicate if the aspect identified by the aspects.localizedAspectName field is a product aspect (relevant to catalog products in the category) or an item/instance aspect, which is an aspect whose value will vary based on a particular instance of the product. | [optional] 
**AspectDataType** | Pointer to **string** | The data type of this aspect. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:AspectDataTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**AspectEnabledForVariations** | Pointer to **bool** | A value of true indicates that this aspect can be used to help identify item variations. | [optional] 
**AspectFormat** | Pointer to **string** | Returned only if the value of aspectDataType identifies a data type that requires specific formatting. Currently, this field provides formatting hints as follows: DATE: YYYY, YYYYMM, YYYYMMDD NUMBER: int32, double | [optional] 
**AspectMaxLength** | Pointer to **int32** | The maximum length of the item/instance aspect&#39;s value. The seller must make sure not to exceed this length when specifying the instance aspect&#39;s value for a product. This field is only returned for instance aspects. | [optional] 
**AspectMode** | Pointer to **string** | The manner in which values of this aspect must be specified by the seller (as free text or by selecting from available options). For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:AspectModeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**AspectRequired** | Pointer to **bool** | A value of true indicates that this aspect is required when offering items in the specified category. | [optional] 
**AspectUsage** | Pointer to **string** | The enumeration value returned in this field will indicate if the corresponding aspect is recommended or optional. Note: This field is always returned, even for hard-mandated/required aspects (where aspectRequired: true). The value returned for required aspects will be RECOMMENDED, but they are actually required and a seller will be blocked from listing or revising an item without these aspects. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:AspectUsageEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ExpectedRequiredByDate** | Pointer to **string** | The expected date after which the aspect will be required. Note: The value returned in this field specifies only an approximate date, which may not reflect the actual date after which the aspect is required. | [optional] 
**ItemToAspectCardinality** | Pointer to **string** | Indicates whether this aspect can accept single or multiple values for items in the specified category. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/commerce/taxonomy/types/txn:ItemToAspectCardinalityEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewAspectConstraint

`func NewAspectConstraint() *AspectConstraint`

NewAspectConstraint instantiates a new AspectConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectConstraintWithDefaults

`func NewAspectConstraintWithDefaults() *AspectConstraint`

NewAspectConstraintWithDefaults instantiates a new AspectConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectApplicableTo

`func (o *AspectConstraint) GetAspectApplicableTo() []string`

GetAspectApplicableTo returns the AspectApplicableTo field if non-nil, zero value otherwise.

### GetAspectApplicableToOk

`func (o *AspectConstraint) GetAspectApplicableToOk() (*[]string, bool)`

GetAspectApplicableToOk returns a tuple with the AspectApplicableTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectApplicableTo

`func (o *AspectConstraint) SetAspectApplicableTo(v []string)`

SetAspectApplicableTo sets AspectApplicableTo field to given value.

### HasAspectApplicableTo

`func (o *AspectConstraint) HasAspectApplicableTo() bool`

HasAspectApplicableTo returns a boolean if a field has been set.

### GetAspectDataType

`func (o *AspectConstraint) GetAspectDataType() string`

GetAspectDataType returns the AspectDataType field if non-nil, zero value otherwise.

### GetAspectDataTypeOk

`func (o *AspectConstraint) GetAspectDataTypeOk() (*string, bool)`

GetAspectDataTypeOk returns a tuple with the AspectDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectDataType

`func (o *AspectConstraint) SetAspectDataType(v string)`

SetAspectDataType sets AspectDataType field to given value.

### HasAspectDataType

`func (o *AspectConstraint) HasAspectDataType() bool`

HasAspectDataType returns a boolean if a field has been set.

### GetAspectEnabledForVariations

`func (o *AspectConstraint) GetAspectEnabledForVariations() bool`

GetAspectEnabledForVariations returns the AspectEnabledForVariations field if non-nil, zero value otherwise.

### GetAspectEnabledForVariationsOk

`func (o *AspectConstraint) GetAspectEnabledForVariationsOk() (*bool, bool)`

GetAspectEnabledForVariationsOk returns a tuple with the AspectEnabledForVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectEnabledForVariations

`func (o *AspectConstraint) SetAspectEnabledForVariations(v bool)`

SetAspectEnabledForVariations sets AspectEnabledForVariations field to given value.

### HasAspectEnabledForVariations

`func (o *AspectConstraint) HasAspectEnabledForVariations() bool`

HasAspectEnabledForVariations returns a boolean if a field has been set.

### GetAspectFormat

`func (o *AspectConstraint) GetAspectFormat() string`

GetAspectFormat returns the AspectFormat field if non-nil, zero value otherwise.

### GetAspectFormatOk

`func (o *AspectConstraint) GetAspectFormatOk() (*string, bool)`

GetAspectFormatOk returns a tuple with the AspectFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectFormat

`func (o *AspectConstraint) SetAspectFormat(v string)`

SetAspectFormat sets AspectFormat field to given value.

### HasAspectFormat

`func (o *AspectConstraint) HasAspectFormat() bool`

HasAspectFormat returns a boolean if a field has been set.

### GetAspectMaxLength

`func (o *AspectConstraint) GetAspectMaxLength() int32`

GetAspectMaxLength returns the AspectMaxLength field if non-nil, zero value otherwise.

### GetAspectMaxLengthOk

`func (o *AspectConstraint) GetAspectMaxLengthOk() (*int32, bool)`

GetAspectMaxLengthOk returns a tuple with the AspectMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectMaxLength

`func (o *AspectConstraint) SetAspectMaxLength(v int32)`

SetAspectMaxLength sets AspectMaxLength field to given value.

### HasAspectMaxLength

`func (o *AspectConstraint) HasAspectMaxLength() bool`

HasAspectMaxLength returns a boolean if a field has been set.

### GetAspectMode

`func (o *AspectConstraint) GetAspectMode() string`

GetAspectMode returns the AspectMode field if non-nil, zero value otherwise.

### GetAspectModeOk

`func (o *AspectConstraint) GetAspectModeOk() (*string, bool)`

GetAspectModeOk returns a tuple with the AspectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectMode

`func (o *AspectConstraint) SetAspectMode(v string)`

SetAspectMode sets AspectMode field to given value.

### HasAspectMode

`func (o *AspectConstraint) HasAspectMode() bool`

HasAspectMode returns a boolean if a field has been set.

### GetAspectRequired

`func (o *AspectConstraint) GetAspectRequired() bool`

GetAspectRequired returns the AspectRequired field if non-nil, zero value otherwise.

### GetAspectRequiredOk

`func (o *AspectConstraint) GetAspectRequiredOk() (*bool, bool)`

GetAspectRequiredOk returns a tuple with the AspectRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRequired

`func (o *AspectConstraint) SetAspectRequired(v bool)`

SetAspectRequired sets AspectRequired field to given value.

### HasAspectRequired

`func (o *AspectConstraint) HasAspectRequired() bool`

HasAspectRequired returns a boolean if a field has been set.

### GetAspectUsage

`func (o *AspectConstraint) GetAspectUsage() string`

GetAspectUsage returns the AspectUsage field if non-nil, zero value otherwise.

### GetAspectUsageOk

`func (o *AspectConstraint) GetAspectUsageOk() (*string, bool)`

GetAspectUsageOk returns a tuple with the AspectUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectUsage

`func (o *AspectConstraint) SetAspectUsage(v string)`

SetAspectUsage sets AspectUsage field to given value.

### HasAspectUsage

`func (o *AspectConstraint) HasAspectUsage() bool`

HasAspectUsage returns a boolean if a field has been set.

### GetExpectedRequiredByDate

`func (o *AspectConstraint) GetExpectedRequiredByDate() string`

GetExpectedRequiredByDate returns the ExpectedRequiredByDate field if non-nil, zero value otherwise.

### GetExpectedRequiredByDateOk

`func (o *AspectConstraint) GetExpectedRequiredByDateOk() (*string, bool)`

GetExpectedRequiredByDateOk returns a tuple with the ExpectedRequiredByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRequiredByDate

`func (o *AspectConstraint) SetExpectedRequiredByDate(v string)`

SetExpectedRequiredByDate sets ExpectedRequiredByDate field to given value.

### HasExpectedRequiredByDate

`func (o *AspectConstraint) HasExpectedRequiredByDate() bool`

HasExpectedRequiredByDate returns a boolean if a field has been set.

### GetItemToAspectCardinality

`func (o *AspectConstraint) GetItemToAspectCardinality() string`

GetItemToAspectCardinality returns the ItemToAspectCardinality field if non-nil, zero value otherwise.

### GetItemToAspectCardinalityOk

`func (o *AspectConstraint) GetItemToAspectCardinalityOk() (*string, bool)`

GetItemToAspectCardinalityOk returns a tuple with the ItemToAspectCardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemToAspectCardinality

`func (o *AspectConstraint) SetItemToAspectCardinality(v string)`

SetItemToAspectCardinality sets ItemToAspectCardinality field to given value.

### HasItemToAspectCardinality

`func (o *AspectConstraint) HasItemToAspectCardinality() bool`

HasItemToAspectCardinality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


