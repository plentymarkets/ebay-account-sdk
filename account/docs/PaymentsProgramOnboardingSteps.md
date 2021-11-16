# PaymentsProgramOnboardingSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the step in the steps array. Over time, these names are subject to change as processes change. The output sample contains example step names. Review an actual call response for updated step names.  | [optional] 
**Status** | Pointer to **string** | This enumeration value indicates the status of the associated step. &lt;p&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;strong&gt;Note:&lt;/strong&gt; Only one step can be &lt;code&gt;IN_PROGRESS&lt;/code&gt; at a time.&lt;/span&gt;&lt;/p&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:PaymentsProgramOnboardingStepStatus&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**WebUrl** | Pointer to **string** | This URL provides access to the &lt;code&gt;IN_PROGRESS&lt;/code&gt; step. | [optional] 

## Methods

### NewPaymentsProgramOnboardingSteps

`func NewPaymentsProgramOnboardingSteps() *PaymentsProgramOnboardingSteps`

NewPaymentsProgramOnboardingSteps instantiates a new PaymentsProgramOnboardingSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsProgramOnboardingStepsWithDefaults

`func NewPaymentsProgramOnboardingStepsWithDefaults() *PaymentsProgramOnboardingSteps`

NewPaymentsProgramOnboardingStepsWithDefaults instantiates a new PaymentsProgramOnboardingSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaymentsProgramOnboardingSteps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentsProgramOnboardingSteps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentsProgramOnboardingSteps) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentsProgramOnboardingSteps) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentsProgramOnboardingSteps) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentsProgramOnboardingSteps) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentsProgramOnboardingSteps) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentsProgramOnboardingSteps) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebUrl

`func (o *PaymentsProgramOnboardingSteps) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *PaymentsProgramOnboardingSteps) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *PaymentsProgramOnboardingSteps) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *PaymentsProgramOnboardingSteps) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


