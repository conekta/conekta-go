# CustomerFiscalEntitiesDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**CustomerAddress**](CustomerAddress.md) |  | 
**TaxId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerFiscalEntitiesDataResponse

`func NewCustomerFiscalEntitiesDataResponse(address CustomerAddress, id string, object string, createdAt int64, ) *CustomerFiscalEntitiesDataResponse`

NewCustomerFiscalEntitiesDataResponse instantiates a new CustomerFiscalEntitiesDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerFiscalEntitiesDataResponseWithDefaults

`func NewCustomerFiscalEntitiesDataResponseWithDefaults() *CustomerFiscalEntitiesDataResponse`

NewCustomerFiscalEntitiesDataResponseWithDefaults instantiates a new CustomerFiscalEntitiesDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CustomerFiscalEntitiesDataResponse) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerFiscalEntitiesDataResponse) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerFiscalEntitiesDataResponse) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.


### GetTaxId

`func (o *CustomerFiscalEntitiesDataResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerFiscalEntitiesDataResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerFiscalEntitiesDataResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerFiscalEntitiesDataResponse) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerFiscalEntitiesDataResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerFiscalEntitiesDataResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerFiscalEntitiesDataResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerFiscalEntitiesDataResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerFiscalEntitiesDataResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerFiscalEntitiesDataResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerFiscalEntitiesDataResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerFiscalEntitiesDataResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerFiscalEntitiesDataResponse) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerFiscalEntitiesDataResponse) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerFiscalEntitiesDataResponse) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerFiscalEntitiesDataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerFiscalEntitiesDataResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerFiscalEntitiesDataResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerFiscalEntitiesDataResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerFiscalEntitiesDataResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetId

`func (o *CustomerFiscalEntitiesDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerFiscalEntitiesDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerFiscalEntitiesDataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CustomerFiscalEntitiesDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerFiscalEntitiesDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerFiscalEntitiesDataResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *CustomerFiscalEntitiesDataResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerFiscalEntitiesDataResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerFiscalEntitiesDataResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *CustomerFiscalEntitiesDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerFiscalEntitiesDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerFiscalEntitiesDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerFiscalEntitiesDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerFiscalEntitiesDataResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerFiscalEntitiesDataResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerFiscalEntitiesDataResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerFiscalEntitiesDataResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


