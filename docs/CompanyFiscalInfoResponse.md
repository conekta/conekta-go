# CompanyFiscalInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | The resource&#39;s type | [optional] 
**TaxId** | Pointer to **string** | Tax ID of the company | [optional] 
**LegalEntityName** | Pointer to **string** | Legal name of the company | [optional] 
**BusinessType** | Pointer to **string** | Business type of the company | [optional] 
**Phone** | Pointer to **string** | Phone number of the company | [optional] 
**PhysicalPersonBusinessType** | Pointer to **string** | Business type if &#39;persona_fisica&#39; | [optional] 
**Address** | Pointer to [**CompanyFiscalInfoAddressResponse**](CompanyFiscalInfoAddressResponse.md) |  | [optional] 

## Methods

### NewCompanyFiscalInfoResponse

`func NewCompanyFiscalInfoResponse() *CompanyFiscalInfoResponse`

NewCompanyFiscalInfoResponse instantiates a new CompanyFiscalInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFiscalInfoResponseWithDefaults

`func NewCompanyFiscalInfoResponseWithDefaults() *CompanyFiscalInfoResponse`

NewCompanyFiscalInfoResponseWithDefaults instantiates a new CompanyFiscalInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CompanyFiscalInfoResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompanyFiscalInfoResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompanyFiscalInfoResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompanyFiscalInfoResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTaxId

`func (o *CompanyFiscalInfoResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CompanyFiscalInfoResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CompanyFiscalInfoResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CompanyFiscalInfoResponse) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetLegalEntityName

`func (o *CompanyFiscalInfoResponse) GetLegalEntityName() string`

GetLegalEntityName returns the LegalEntityName field if non-nil, zero value otherwise.

### GetLegalEntityNameOk

`func (o *CompanyFiscalInfoResponse) GetLegalEntityNameOk() (*string, bool)`

GetLegalEntityNameOk returns a tuple with the LegalEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityName

`func (o *CompanyFiscalInfoResponse) SetLegalEntityName(v string)`

SetLegalEntityName sets LegalEntityName field to given value.

### HasLegalEntityName

`func (o *CompanyFiscalInfoResponse) HasLegalEntityName() bool`

HasLegalEntityName returns a boolean if a field has been set.

### GetBusinessType

`func (o *CompanyFiscalInfoResponse) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *CompanyFiscalInfoResponse) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *CompanyFiscalInfoResponse) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *CompanyFiscalInfoResponse) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetPhone

`func (o *CompanyFiscalInfoResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyFiscalInfoResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyFiscalInfoResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyFiscalInfoResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhysicalPersonBusinessType

`func (o *CompanyFiscalInfoResponse) GetPhysicalPersonBusinessType() string`

GetPhysicalPersonBusinessType returns the PhysicalPersonBusinessType field if non-nil, zero value otherwise.

### GetPhysicalPersonBusinessTypeOk

`func (o *CompanyFiscalInfoResponse) GetPhysicalPersonBusinessTypeOk() (*string, bool)`

GetPhysicalPersonBusinessTypeOk returns a tuple with the PhysicalPersonBusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPersonBusinessType

`func (o *CompanyFiscalInfoResponse) SetPhysicalPersonBusinessType(v string)`

SetPhysicalPersonBusinessType sets PhysicalPersonBusinessType field to given value.

### HasPhysicalPersonBusinessType

`func (o *CompanyFiscalInfoResponse) HasPhysicalPersonBusinessType() bool`

HasPhysicalPersonBusinessType returns a boolean if a field has been set.

### GetAddress

`func (o *CompanyFiscalInfoResponse) GetAddress() CompanyFiscalInfoAddressResponse`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyFiscalInfoResponse) GetAddressOk() (*CompanyFiscalInfoAddressResponse, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyFiscalInfoResponse) SetAddress(v CompanyFiscalInfoAddressResponse)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CompanyFiscalInfoResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


