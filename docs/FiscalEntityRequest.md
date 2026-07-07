# FiscalEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**FiscalEntityRequestAddress**](FiscalEntityRequestAddress.md) |  | 
**TaxId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to  |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 

## Methods

### NewFiscalEntityRequest

`func NewFiscalEntityRequest(address FiscalEntityRequestAddress, ) *FiscalEntityRequest`

NewFiscalEntityRequest instantiates a new FiscalEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalEntityRequestWithDefaults

`func NewFiscalEntityRequestWithDefaults() *FiscalEntityRequest`

NewFiscalEntityRequestWithDefaults instantiates a new FiscalEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FiscalEntityRequest) GetAddress() FiscalEntityRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FiscalEntityRequest) GetAddressOk() (*FiscalEntityRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FiscalEntityRequest) SetAddress(v FiscalEntityRequestAddress)`

SetAddress sets Address field to given value.


### GetTaxId

`func (o *FiscalEntityRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *FiscalEntityRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *FiscalEntityRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *FiscalEntityRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *FiscalEntityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FiscalEntityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FiscalEntityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FiscalEntityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *FiscalEntityRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *FiscalEntityRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *FiscalEntityRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *FiscalEntityRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *FiscalEntityRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FiscalEntityRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FiscalEntityRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FiscalEntityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *FiscalEntityRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FiscalEntityRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCompanyName

`func (o *FiscalEntityRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *FiscalEntityRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *FiscalEntityRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *FiscalEntityRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


