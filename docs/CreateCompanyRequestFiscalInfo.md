# CreateCompanyRequestFiscalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessPhone** | Pointer to **string** | The business phone number for fiscal purposes. | [optional] 
**FiscalType** | Pointer to **string** | The fiscal type of the company (e.g., &#39;moral&#39;, &#39;persona_fisica&#39;). | [optional] 

## Methods

### NewCreateCompanyRequestFiscalInfo

`func NewCreateCompanyRequestFiscalInfo() *CreateCompanyRequestFiscalInfo`

NewCreateCompanyRequestFiscalInfo instantiates a new CreateCompanyRequestFiscalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyRequestFiscalInfoWithDefaults

`func NewCreateCompanyRequestFiscalInfoWithDefaults() *CreateCompanyRequestFiscalInfo`

NewCreateCompanyRequestFiscalInfoWithDefaults instantiates a new CreateCompanyRequestFiscalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessPhone

`func (o *CreateCompanyRequestFiscalInfo) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *CreateCompanyRequestFiscalInfo) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *CreateCompanyRequestFiscalInfo) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *CreateCompanyRequestFiscalInfo) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### GetFiscalType

`func (o *CreateCompanyRequestFiscalInfo) GetFiscalType() string`

GetFiscalType returns the FiscalType field if non-nil, zero value otherwise.

### GetFiscalTypeOk

`func (o *CreateCompanyRequestFiscalInfo) GetFiscalTypeOk() (*string, bool)`

GetFiscalTypeOk returns a tuple with the FiscalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalType

`func (o *CreateCompanyRequestFiscalInfo) SetFiscalType(v string)`

SetFiscalType sets FiscalType field to given value.

### HasFiscalType

`func (o *CreateCompanyRequestFiscalInfo) HasFiscalType() bool`

HasFiscalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


