# GetCompatibilityPropertyValuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibilityPropertyValues** | Pointer to [**[]CompatibilityPropertyValue**](CompatibilityPropertyValue.md) | This array contains all compatible vehicle property values that match the specified eBay marketplace, specified eBay category, and filters in the request. If the compatibility_property parameter value in the request is &#39;Trim&#39;, each value returned in each value field will be a different vehicle trim, applicable to any filters that are set in the filter query parameter of the request, and also based on the eBay marketplace and category specified in the call request. | [optional] 

## Methods

### NewGetCompatibilityPropertyValuesResponse

`func NewGetCompatibilityPropertyValuesResponse() *GetCompatibilityPropertyValuesResponse`

NewGetCompatibilityPropertyValuesResponse instantiates a new GetCompatibilityPropertyValuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompatibilityPropertyValuesResponseWithDefaults

`func NewGetCompatibilityPropertyValuesResponseWithDefaults() *GetCompatibilityPropertyValuesResponse`

NewGetCompatibilityPropertyValuesResponseWithDefaults instantiates a new GetCompatibilityPropertyValuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibilityPropertyValues

`func (o *GetCompatibilityPropertyValuesResponse) GetCompatibilityPropertyValues() []CompatibilityPropertyValue`

GetCompatibilityPropertyValues returns the CompatibilityPropertyValues field if non-nil, zero value otherwise.

### GetCompatibilityPropertyValuesOk

`func (o *GetCompatibilityPropertyValuesResponse) GetCompatibilityPropertyValuesOk() (*[]CompatibilityPropertyValue, bool)`

GetCompatibilityPropertyValuesOk returns a tuple with the CompatibilityPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityPropertyValues

`func (o *GetCompatibilityPropertyValuesResponse) SetCompatibilityPropertyValues(v []CompatibilityPropertyValue)`

SetCompatibilityPropertyValues sets CompatibilityPropertyValues field to given value.

### HasCompatibilityPropertyValues

`func (o *GetCompatibilityPropertyValuesResponse) HasCompatibilityPropertyValues() bool`

HasCompatibilityPropertyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


