# OrderUpdateFiscalEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**FiscalEntityAddress**](FiscalEntityAddress.md) |  | 
**Email** | Pointer to **string** | Email of the fiscal entity | [optional] 
**Name** | Pointer to **NullableString** | Name of the fiscal entity | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the fiscal entity | [optional] 
**Phone** | Pointer to **string** | Phone of the fiscal entity | [optional] 
**TaxId** | Pointer to **NullableString** | Tax ID of the fiscal entity | [optional] 

## Methods

### NewOrderUpdateFiscalEntityRequest

`func NewOrderUpdateFiscalEntityRequest(address FiscalEntityAddress, ) *OrderUpdateFiscalEntityRequest`

NewOrderUpdateFiscalEntityRequest instantiates a new OrderUpdateFiscalEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateFiscalEntityRequestWithDefaults

`func NewOrderUpdateFiscalEntityRequestWithDefaults() *OrderUpdateFiscalEntityRequest`

NewOrderUpdateFiscalEntityRequestWithDefaults instantiates a new OrderUpdateFiscalEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrderUpdateFiscalEntityRequest) GetAddress() FiscalEntityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderUpdateFiscalEntityRequest) GetAddressOk() (*FiscalEntityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderUpdateFiscalEntityRequest) SetAddress(v FiscalEntityAddress)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *OrderUpdateFiscalEntityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderUpdateFiscalEntityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderUpdateFiscalEntityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderUpdateFiscalEntityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *OrderUpdateFiscalEntityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderUpdateFiscalEntityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderUpdateFiscalEntityRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderUpdateFiscalEntityRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrderUpdateFiscalEntityRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrderUpdateFiscalEntityRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMetadata

`func (o *OrderUpdateFiscalEntityRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderUpdateFiscalEntityRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderUpdateFiscalEntityRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderUpdateFiscalEntityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPhone

`func (o *OrderUpdateFiscalEntityRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderUpdateFiscalEntityRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderUpdateFiscalEntityRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderUpdateFiscalEntityRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTaxId

`func (o *OrderUpdateFiscalEntityRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *OrderUpdateFiscalEntityRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *OrderUpdateFiscalEntityRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *OrderUpdateFiscalEntityRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *OrderUpdateFiscalEntityRequest) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *OrderUpdateFiscalEntityRequest) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


