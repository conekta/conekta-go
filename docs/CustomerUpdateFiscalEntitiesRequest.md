# CustomerUpdateFiscalEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**CustomerAddress**](CustomerAddress.md) |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerUpdateFiscalEntitiesRequest

`func NewCustomerUpdateFiscalEntitiesRequest() *CustomerUpdateFiscalEntitiesRequest`

NewCustomerUpdateFiscalEntitiesRequest instantiates a new CustomerUpdateFiscalEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateFiscalEntitiesRequestWithDefaults

`func NewCustomerUpdateFiscalEntitiesRequestWithDefaults() *CustomerUpdateFiscalEntitiesRequest`

NewCustomerUpdateFiscalEntitiesRequestWithDefaults instantiates a new CustomerUpdateFiscalEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CustomerUpdateFiscalEntitiesRequest) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerUpdateFiscalEntitiesRequest) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerUpdateFiscalEntitiesRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTaxId

`func (o *CustomerUpdateFiscalEntitiesRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerUpdateFiscalEntitiesRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerUpdateFiscalEntitiesRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerUpdateFiscalEntitiesRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerUpdateFiscalEntitiesRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerUpdateFiscalEntitiesRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerUpdateFiscalEntitiesRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerUpdateFiscalEntitiesRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerUpdateFiscalEntitiesRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerUpdateFiscalEntitiesRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerUpdateFiscalEntitiesRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerUpdateFiscalEntitiesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerUpdateFiscalEntitiesRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerUpdateFiscalEntitiesRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerUpdateFiscalEntitiesRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerUpdateFiscalEntitiesRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


