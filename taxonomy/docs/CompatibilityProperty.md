# CompatibilityProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This is the actual name of the compatible vehicle property as it is known on the specified eBay marketplace and in the eBay category. This is the string value that should be used in the compatibility_property and filter query parameters of a getCompatibilityPropertyValues request URI. Typical vehicle properties are &#39;Make&#39;, &#39;Model&#39;, &#39;Year&#39;, &#39;Engine&#39;, and &#39;Trim&#39;, but will vary based on the eBay marketplace and the eBay category. | [optional] 
**LocalizedName** | Pointer to **string** | This is the localized name of the compatible vehicle property. The language that is used will depend on the user making the call, or based on the language specified if the Content-Language HTTP header is used. In some instances, the string value in this field may be the same as the string in the corresponding name field. | [optional] 

## Methods

### NewCompatibilityProperty

`func NewCompatibilityProperty() *CompatibilityProperty`

NewCompatibilityProperty instantiates a new CompatibilityProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibilityPropertyWithDefaults

`func NewCompatibilityPropertyWithDefaults() *CompatibilityProperty`

NewCompatibilityPropertyWithDefaults instantiates a new CompatibilityProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompatibilityProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompatibilityProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompatibilityProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompatibilityProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocalizedName

`func (o *CompatibilityProperty) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *CompatibilityProperty) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *CompatibilityProperty) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *CompatibilityProperty) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


