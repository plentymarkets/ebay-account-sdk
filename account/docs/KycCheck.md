# KycCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRequired** | Pointer to **string** | The enumeration value returned in this field categorizes the type of details needed for the KYC check. More information about the check is shown in the &lt;b&gt;detailMessage&lt;/b&gt; and other applicable, corresponding fields. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/sell/account/types/kyc:DetailsType&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**DueDate** | Pointer to **string** | The timestamp in this field indicates the date by which the seller should resolve the KYC requirement.&lt;br&gt;&lt;br&gt;The timestamp in this field uses the UTC date and time format described in the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-8601-date-and-time-format.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 8601 Standard&lt;/a&gt;. See below for this format and an example: &lt;br&gt;&lt;br&gt;&lt;i&gt;MM-DD-YYYY HH:MM:SS&lt;/i&gt;&lt;br/&gt;&lt;code&gt;06-05-2020 10:34:18&lt;/code&gt; | [optional] 
**RemedyUrl** | Pointer to **string** | If applicable and available, a URL will be returned in this field, and the link will take the seller to an eBay page where they can provide the requested information. | [optional] 
**Alert** | Pointer to **string** | This field gives a short summary of what is required from the seller. An example might be, &#39;&lt;code&gt;Upload bank document now.&lt;/code&gt;&#39;. The &lt;b&gt;detailMessage&lt;/b&gt; field will often provide more details on what is required of the seller. | [optional] 
**DetailMessage** | Pointer to **string** | This field gives a detailed message about what is required from the seller. An example might be, &#39;&lt;code&gt;Please upload a bank document by 2020-08-01 to get your account back in good standing.&lt;/code&gt;&#39;. | [optional] 

## Methods

### NewKycCheck

`func NewKycCheck() *KycCheck`

NewKycCheck instantiates a new KycCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycCheckWithDefaults

`func NewKycCheckWithDefaults() *KycCheck`

NewKycCheckWithDefaults instantiates a new KycCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRequired

`func (o *KycCheck) GetDataRequired() string`

GetDataRequired returns the DataRequired field if non-nil, zero value otherwise.

### GetDataRequiredOk

`func (o *KycCheck) GetDataRequiredOk() (*string, bool)`

GetDataRequiredOk returns a tuple with the DataRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRequired

`func (o *KycCheck) SetDataRequired(v string)`

SetDataRequired sets DataRequired field to given value.

### HasDataRequired

`func (o *KycCheck) HasDataRequired() bool`

HasDataRequired returns a boolean if a field has been set.

### GetDueDate

`func (o *KycCheck) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *KycCheck) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *KycCheck) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *KycCheck) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetRemedyUrl

`func (o *KycCheck) GetRemedyUrl() string`

GetRemedyUrl returns the RemedyUrl field if non-nil, zero value otherwise.

### GetRemedyUrlOk

`func (o *KycCheck) GetRemedyUrlOk() (*string, bool)`

GetRemedyUrlOk returns a tuple with the RemedyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyUrl

`func (o *KycCheck) SetRemedyUrl(v string)`

SetRemedyUrl sets RemedyUrl field to given value.

### HasRemedyUrl

`func (o *KycCheck) HasRemedyUrl() bool`

HasRemedyUrl returns a boolean if a field has been set.

### GetAlert

`func (o *KycCheck) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *KycCheck) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *KycCheck) SetAlert(v string)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *KycCheck) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetDetailMessage

`func (o *KycCheck) GetDetailMessage() string`

GetDetailMessage returns the DetailMessage field if non-nil, zero value otherwise.

### GetDetailMessageOk

`func (o *KycCheck) GetDetailMessageOk() (*string, bool)`

GetDetailMessageOk returns a tuple with the DetailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailMessage

`func (o *KycCheck) SetDetailMessage(v string)`

SetDetailMessage sets DetailMessage field to given value.

### HasDetailMessage

`func (o *KycCheck) HasDetailMessage() bool`

HasDetailMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


