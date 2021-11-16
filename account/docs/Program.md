# Program

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramType** | Pointer to **string** | A seller program in to which a seller can opt-in. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:ProgramTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewProgram

`func NewProgram() *Program`

NewProgram instantiates a new Program object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramWithDefaults

`func NewProgramWithDefaults() *Program`

NewProgramWithDefaults instantiates a new Program object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramType

`func (o *Program) GetProgramType() string`

GetProgramType returns the ProgramType field if non-nil, zero value otherwise.

### GetProgramTypeOk

`func (o *Program) GetProgramTypeOk() (*string, bool)`

GetProgramTypeOk returns a tuple with the ProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramType

`func (o *Program) SetProgramType(v string)`

SetProgramType sets ProgramType field to given value.

### HasProgramType

`func (o *Program) HasProgramType() bool`

HasProgramType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


