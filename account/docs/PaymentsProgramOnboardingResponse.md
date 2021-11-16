# PaymentsProgramOnboardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnboardingStatus** | Pointer to **string** | This enumeration value indicates the eligibility of payment onboarding for the registered site. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/api:PaymentsProgramOnboardingStatus&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Steps** | Pointer to [**[]PaymentsProgramOnboardingSteps**](PaymentsProgramOnboardingSteps.md) | An array of the active process steps for payment onboarding and the status of each step. This array includes the step &lt;strong&gt;name&lt;/strong&gt;, step &lt;strong&gt;status&lt;/strong&gt;, and a &lt;strong&gt;webUrl&lt;/strong&gt; to the &lt;code&gt;IN_PROGRESS&lt;/code&gt; step. The step names are returned in sequential order.  | [optional] 

## Methods

### NewPaymentsProgramOnboardingResponse

`func NewPaymentsProgramOnboardingResponse() *PaymentsProgramOnboardingResponse`

NewPaymentsProgramOnboardingResponse instantiates a new PaymentsProgramOnboardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsProgramOnboardingResponseWithDefaults

`func NewPaymentsProgramOnboardingResponseWithDefaults() *PaymentsProgramOnboardingResponse`

NewPaymentsProgramOnboardingResponseWithDefaults instantiates a new PaymentsProgramOnboardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnboardingStatus

`func (o *PaymentsProgramOnboardingResponse) GetOnboardingStatus() string`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *PaymentsProgramOnboardingResponse) GetOnboardingStatusOk() (*string, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *PaymentsProgramOnboardingResponse) SetOnboardingStatus(v string)`

SetOnboardingStatus sets OnboardingStatus field to given value.

### HasOnboardingStatus

`func (o *PaymentsProgramOnboardingResponse) HasOnboardingStatus() bool`

HasOnboardingStatus returns a boolean if a field has been set.

### GetSteps

`func (o *PaymentsProgramOnboardingResponse) GetSteps() []PaymentsProgramOnboardingSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *PaymentsProgramOnboardingResponse) GetStepsOk() (*[]PaymentsProgramOnboardingSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *PaymentsProgramOnboardingResponse) SetSteps(v []PaymentsProgramOnboardingSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *PaymentsProgramOnboardingResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


