# OrderFiscalEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**FiscalEntityAddress**](FiscalEntityAddress.md) |  | 
**Email** | Pointer to **string** | Email of the fiscal entity | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the fiscal entity | [optional] 
**Name** | Pointer to **NullableString** | Name of the fiscal entity | [optional] 
**Phone** | Pointer to **string** | Phone of the fiscal entity | [optional] 
**TaxId** | Pointer to **NullableString** | Tax ID of the fiscal entity | [optional] 

## Methods

### NewOrderFiscalEntityRequest

`func NewOrderFiscalEntityRequest(address FiscalEntityAddress, ) *OrderFiscalEntityRequest`

NewOrderFiscalEntityRequest instantiates a new OrderFiscalEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFiscalEntityRequestWithDefaults

`func NewOrderFiscalEntityRequestWithDefaults() *OrderFiscalEntityRequest`

NewOrderFiscalEntityRequestWithDefaults instantiates a new OrderFiscalEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrderFiscalEntityRequest) GetAddress() FiscalEntityAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderFiscalEntityRequest) GetAddressOk() (*FiscalEntityAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderFiscalEntityRequest) SetAddress(v FiscalEntityAddress)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *OrderFiscalEntityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderFiscalEntityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderFiscalEntityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderFiscalEntityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderFiscalEntityRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderFiscalEntityRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderFiscalEntityRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderFiscalEntityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrderFiscalEntityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderFiscalEntityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderFiscalEntityRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderFiscalEntityRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrderFiscalEntityRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrderFiscalEntityRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *OrderFiscalEntityRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderFiscalEntityRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderFiscalEntityRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderFiscalEntityRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTaxId

`func (o *OrderFiscalEntityRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *OrderFiscalEntityRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *OrderFiscalEntityRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *OrderFiscalEntityRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *OrderFiscalEntityRequest) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *OrderFiscalEntityRequest) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


