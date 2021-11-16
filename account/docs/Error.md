# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category type for this error or warning. It is a string that can have one of three values:&lt;ul&gt;&lt;li&gt;&lt;code&gt;Application&lt;/code&gt;: Indicates an exception or error occurred in the application code or at runtime. Examples include catching an exception in a service&#39;s business logic, system failures, or request errors from a dependency.&lt;/li&gt;&lt;li&gt;&lt;code&gt;Business&lt;/code&gt;: Used when your service or a dependent service refused to continue processing on the resource because of a business rule violation such as \&quot;Seller does not ship item to Antarctica\&quot; or \&quot;Buyer ineligible to purchase an alcoholic item\&quot;. Business errors are not syntactical input errors.&lt;/li&gt;&lt;li&gt;&lt;code&gt;Request&lt;/code&gt;: Used when there is anything wrong with the request, such as authentication, syntactical errors, rate limiting or missing headers, bad HTTP header values, and so on.&lt;/li&gt;&lt;/ul&gt; | [optional] 
**Domain** | Pointer to **string** | Name of the domain ,or primary system, of the service or application where the error occurred. | [optional] 
**ErrorId** | Pointer to **int32** | A positive integer that uniquely identifies the specific error condition that occurred. Your application can use error codes as identifiers in your customized error-handling algorithms. | [optional] 
**InputRefIds** | Pointer to **[]string** | Identifies specific request elements associated with the error, if any. inputRefId&#39;s response is format specific. For JSON, use &lt;i&gt;JSONPath&lt;/i&gt; notation. | [optional] 
**LongMessage** | Pointer to **string** | A more detailed explanation of the error than given in the &lt;code&gt;message&lt;/code&gt; error field. | [optional] 
**Message** | Pointer to **string** | Information on how to correct the problem, in the end user&#39;s terms and language where applicable. Its value is at most 50 characters long. If applicable, the value is localized in the end user&#39;s requested locale. | [optional] 
**OutputRefIds** | Pointer to **[]string** | Identifies specific response elements associated with the error, if any. Path format is the same as &lt;code&gt;inputRefId&lt;/code&gt;. | [optional] 
**Parameters** | Pointer to [**[]ErrorParameter**](ErrorParameter.md) | This optional list of name/value pairs that contain context-specific &lt;code&gt;ErrorParameter&lt;/code&gt; objects, with each item in the list being a parameter (or input field name) that caused an error condition. Each &lt;code&gt;ErrorParameter&lt;/code&gt; object consists of two fields, a &lt;code&gt;name&lt;/code&gt; and a &lt;code&gt;value&lt;/code&gt;. | [optional] 
**Subdomain** | Pointer to **string** | If present, indicates the subsystem in which the error occurred. | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Error) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Error) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Error) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Error) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDomain

`func (o *Error) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Error) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Error) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Error) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetErrorId

`func (o *Error) GetErrorId() int32`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *Error) GetErrorIdOk() (*int32, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *Error) SetErrorId(v int32)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *Error) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetInputRefIds

`func (o *Error) GetInputRefIds() []string`

GetInputRefIds returns the InputRefIds field if non-nil, zero value otherwise.

### GetInputRefIdsOk

`func (o *Error) GetInputRefIdsOk() (*[]string, bool)`

GetInputRefIdsOk returns a tuple with the InputRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRefIds

`func (o *Error) SetInputRefIds(v []string)`

SetInputRefIds sets InputRefIds field to given value.

### HasInputRefIds

`func (o *Error) HasInputRefIds() bool`

HasInputRefIds returns a boolean if a field has been set.

### GetLongMessage

`func (o *Error) GetLongMessage() string`

GetLongMessage returns the LongMessage field if non-nil, zero value otherwise.

### GetLongMessageOk

`func (o *Error) GetLongMessageOk() (*string, bool)`

GetLongMessageOk returns a tuple with the LongMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessage

`func (o *Error) SetLongMessage(v string)`

SetLongMessage sets LongMessage field to given value.

### HasLongMessage

`func (o *Error) HasLongMessage() bool`

HasLongMessage returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutputRefIds

`func (o *Error) GetOutputRefIds() []string`

GetOutputRefIds returns the OutputRefIds field if non-nil, zero value otherwise.

### GetOutputRefIdsOk

`func (o *Error) GetOutputRefIdsOk() (*[]string, bool)`

GetOutputRefIdsOk returns a tuple with the OutputRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRefIds

`func (o *Error) SetOutputRefIds(v []string)`

SetOutputRefIds sets OutputRefIds field to given value.

### HasOutputRefIds

`func (o *Error) HasOutputRefIds() bool`

HasOutputRefIds returns a boolean if a field has been set.

### GetParameters

`func (o *Error) GetParameters() []ErrorParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Error) GetParametersOk() (*[]ErrorParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Error) SetParameters(v []ErrorParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Error) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSubdomain

`func (o *Error) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Error) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Error) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Error) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


