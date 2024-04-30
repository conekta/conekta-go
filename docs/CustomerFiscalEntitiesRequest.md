# CustomerFiscalEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**CustomerAddress**](CustomerAddress.md) |  | 
**TaxId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerFiscalEntitiesRequest

`func NewCustomerFiscalEntitiesRequest(address CustomerAddress, ) *CustomerFiscalEntitiesRequest`

NewCustomerFiscalEntitiesRequest instantiates a new CustomerFiscalEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerFiscalEntitiesRequestWithDefaults

`func NewCustomerFiscalEntitiesRequestWithDefaults() *CustomerFiscalEntitiesRequest`

NewCustomerFiscalEntitiesRequestWithDefaults instantiates a new CustomerFiscalEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CustomerFiscalEntitiesRequest) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerFiscalEntitiesRequest) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerFiscalEntitiesRequest) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.


### GetTaxId

`func (o *CustomerFiscalEntitiesRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerFiscalEntitiesRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerFiscalEntitiesRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerFiscalEntitiesRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerFiscalEntitiesRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerFiscalEntitiesRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerFiscalEntitiesRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerFiscalEntitiesRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerFiscalEntitiesRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerFiscalEntitiesRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerFiscalEntitiesRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerFiscalEntitiesRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerFiscalEntitiesRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerFiscalEntitiesRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerFiscalEntitiesRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerFiscalEntitiesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerFiscalEntitiesRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerFiscalEntitiesRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerFiscalEntitiesRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerFiscalEntitiesRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


