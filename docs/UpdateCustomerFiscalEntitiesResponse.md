# UpdateCustomerFiscalEntitiesResponse

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

### NewUpdateCustomerFiscalEntitiesResponse

`func NewUpdateCustomerFiscalEntitiesResponse(address CustomerAddress, id string, object string, createdAt int64, ) *UpdateCustomerFiscalEntitiesResponse`

NewUpdateCustomerFiscalEntitiesResponse instantiates a new UpdateCustomerFiscalEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerFiscalEntitiesResponseWithDefaults

`func NewUpdateCustomerFiscalEntitiesResponseWithDefaults() *UpdateCustomerFiscalEntitiesResponse`

NewUpdateCustomerFiscalEntitiesResponseWithDefaults instantiates a new UpdateCustomerFiscalEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateCustomerFiscalEntitiesResponse) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateCustomerFiscalEntitiesResponse) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.


### GetTaxId

`func (o *UpdateCustomerFiscalEntitiesResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *UpdateCustomerFiscalEntitiesResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *UpdateCustomerFiscalEntitiesResponse) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCustomerFiscalEntitiesResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCustomerFiscalEntitiesResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCustomerFiscalEntitiesResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateCustomerFiscalEntitiesResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateCustomerFiscalEntitiesResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateCustomerFiscalEntitiesResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateCustomerFiscalEntitiesResponse) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateCustomerFiscalEntitiesResponse) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateCustomerFiscalEntitiesResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCompanyName

`func (o *UpdateCustomerFiscalEntitiesResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateCustomerFiscalEntitiesResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateCustomerFiscalEntitiesResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetId

`func (o *UpdateCustomerFiscalEntitiesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomerFiscalEntitiesResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *UpdateCustomerFiscalEntitiesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UpdateCustomerFiscalEntitiesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *UpdateCustomerFiscalEntitiesResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateCustomerFiscalEntitiesResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *UpdateCustomerFiscalEntitiesResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateCustomerFiscalEntitiesResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateCustomerFiscalEntitiesResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *UpdateCustomerFiscalEntitiesResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UpdateCustomerFiscalEntitiesResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UpdateCustomerFiscalEntitiesResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UpdateCustomerFiscalEntitiesResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


