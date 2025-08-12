# CreateCompanyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the company. | [optional] 
**TypeCompany** | Pointer to **string** | The type of company, &#39;owner&#39; | [optional] 
**ComercialInfo** | Pointer to [**CreateCompanyRequestComercialInfo**](CreateCompanyRequestComercialInfo.md) |  | [optional] 
**FiscalInfo** | Pointer to [**CreateCompanyRequestFiscalInfo**](CreateCompanyRequestFiscalInfo.md) |  | [optional] 
**BankAccountInfo** | Pointer to [**CreateCompanyRequestBankAccountInfo**](CreateCompanyRequestBankAccountInfo.md) |  | [optional] 

## Methods

### NewCreateCompanyRequest

`func NewCreateCompanyRequest() *CreateCompanyRequest`

NewCreateCompanyRequest instantiates a new CreateCompanyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyRequestWithDefaults

`func NewCreateCompanyRequestWithDefaults() *CreateCompanyRequest`

NewCreateCompanyRequestWithDefaults instantiates a new CreateCompanyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCompanyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCompanyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCompanyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCompanyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeCompany

`func (o *CreateCompanyRequest) GetTypeCompany() string`

GetTypeCompany returns the TypeCompany field if non-nil, zero value otherwise.

### GetTypeCompanyOk

`func (o *CreateCompanyRequest) GetTypeCompanyOk() (*string, bool)`

GetTypeCompanyOk returns a tuple with the TypeCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCompany

`func (o *CreateCompanyRequest) SetTypeCompany(v string)`

SetTypeCompany sets TypeCompany field to given value.

### HasTypeCompany

`func (o *CreateCompanyRequest) HasTypeCompany() bool`

HasTypeCompany returns a boolean if a field has been set.

### GetComercialInfo

`func (o *CreateCompanyRequest) GetComercialInfo() CreateCompanyRequestComercialInfo`

GetComercialInfo returns the ComercialInfo field if non-nil, zero value otherwise.

### GetComercialInfoOk

`func (o *CreateCompanyRequest) GetComercialInfoOk() (*CreateCompanyRequestComercialInfo, bool)`

GetComercialInfoOk returns a tuple with the ComercialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComercialInfo

`func (o *CreateCompanyRequest) SetComercialInfo(v CreateCompanyRequestComercialInfo)`

SetComercialInfo sets ComercialInfo field to given value.

### HasComercialInfo

`func (o *CreateCompanyRequest) HasComercialInfo() bool`

HasComercialInfo returns a boolean if a field has been set.

### GetFiscalInfo

`func (o *CreateCompanyRequest) GetFiscalInfo() CreateCompanyRequestFiscalInfo`

GetFiscalInfo returns the FiscalInfo field if non-nil, zero value otherwise.

### GetFiscalInfoOk

`func (o *CreateCompanyRequest) GetFiscalInfoOk() (*CreateCompanyRequestFiscalInfo, bool)`

GetFiscalInfoOk returns a tuple with the FiscalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalInfo

`func (o *CreateCompanyRequest) SetFiscalInfo(v CreateCompanyRequestFiscalInfo)`

SetFiscalInfo sets FiscalInfo field to given value.

### HasFiscalInfo

`func (o *CreateCompanyRequest) HasFiscalInfo() bool`

HasFiscalInfo returns a boolean if a field has been set.

### GetBankAccountInfo

`func (o *CreateCompanyRequest) GetBankAccountInfo() CreateCompanyRequestBankAccountInfo`

GetBankAccountInfo returns the BankAccountInfo field if non-nil, zero value otherwise.

### GetBankAccountInfoOk

`func (o *CreateCompanyRequest) GetBankAccountInfoOk() (*CreateCompanyRequestBankAccountInfo, bool)`

GetBankAccountInfoOk returns a tuple with the BankAccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountInfo

`func (o *CreateCompanyRequest) SetBankAccountInfo(v CreateCompanyRequestBankAccountInfo)`

SetBankAccountInfo sets BankAccountInfo field to given value.

### HasBankAccountInfo

`func (o *CreateCompanyRequest) HasBankAccountInfo() bool`

HasBankAccountInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


