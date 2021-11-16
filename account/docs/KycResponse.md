# KycResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycChecks** | Pointer to [**[]KycCheck**](KycCheck.md) | This array contains one or more KYC checks required from a managed payments seller. The seller may need to provide more documentation and/or information about themselves, their company, or the bank account they are using for seller payouts.&lt;br/&gt;&lt;br/&gt;If no KYC checks are currently required from the seller, this array is not returned, and the seller only receives a &lt;code&gt;204 No Content&lt;/code&gt; HTTP status code. | [optional] 

## Methods

### NewKycResponse

`func NewKycResponse() *KycResponse`

NewKycResponse instantiates a new KycResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycResponseWithDefaults

`func NewKycResponseWithDefaults() *KycResponse`

NewKycResponseWithDefaults instantiates a new KycResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycChecks

`func (o *KycResponse) GetKycChecks() []KycCheck`

GetKycChecks returns the KycChecks field if non-nil, zero value otherwise.

### GetKycChecksOk

`func (o *KycResponse) GetKycChecksOk() (*[]KycCheck, bool)`

GetKycChecksOk returns a tuple with the KycChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycChecks

`func (o *KycResponse) SetKycChecks(v []KycCheck)`

SetKycChecks sets KycChecks field to given value.

### HasKycChecks

`func (o *KycResponse) HasKycChecks() bool`

HasKycChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


