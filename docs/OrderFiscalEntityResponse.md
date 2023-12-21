# OrderFiscalEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**OrderFiscalEntityAddressResponse**](OrderFiscalEntityAddressResponse.md) |  | 
**Email** | Pointer to **NullableString** | Email of the fiscal entity | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the fiscal entity | [optional] 
**Name** | Pointer to **NullableString** | Name of the fiscal entity | [optional] 
**TaxId** | Pointer to **NullableString** | Tax ID of the fiscal entity | [optional] 
**Id** | **string** | ID of the fiscal entity | 
**CreatedAt** | **int64** | The time at which the object was created in seconds since the Unix epoch | 
**Object** | **string** |  | 
**Phone** | Pointer to **NullableString** | Phone of the fiscal entity | [optional] 

## Methods

### NewOrderFiscalEntityResponse

`func NewOrderFiscalEntityResponse(address OrderFiscalEntityAddressResponse, id string, createdAt int64, object string, ) *OrderFiscalEntityResponse`

NewOrderFiscalEntityResponse instantiates a new OrderFiscalEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFiscalEntityResponseWithDefaults

`func NewOrderFiscalEntityResponseWithDefaults() *OrderFiscalEntityResponse`

NewOrderFiscalEntityResponseWithDefaults instantiates a new OrderFiscalEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrderFiscalEntityResponse) GetAddress() OrderFiscalEntityAddressResponse`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderFiscalEntityResponse) GetAddressOk() (*OrderFiscalEntityAddressResponse, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderFiscalEntityResponse) SetAddress(v OrderFiscalEntityAddressResponse)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *OrderFiscalEntityResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderFiscalEntityResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderFiscalEntityResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderFiscalEntityResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrderFiscalEntityResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrderFiscalEntityResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMetadata

`func (o *OrderFiscalEntityResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderFiscalEntityResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderFiscalEntityResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderFiscalEntityResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrderFiscalEntityResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderFiscalEntityResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderFiscalEntityResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderFiscalEntityResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrderFiscalEntityResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrderFiscalEntityResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxId

`func (o *OrderFiscalEntityResponse) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *OrderFiscalEntityResponse) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *OrderFiscalEntityResponse) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *OrderFiscalEntityResponse) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *OrderFiscalEntityResponse) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *OrderFiscalEntityResponse) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetId

`func (o *OrderFiscalEntityResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderFiscalEntityResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderFiscalEntityResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrderFiscalEntityResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderFiscalEntityResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderFiscalEntityResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetObject

`func (o *OrderFiscalEntityResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderFiscalEntityResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderFiscalEntityResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetPhone

`func (o *OrderFiscalEntityResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderFiscalEntityResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderFiscalEntityResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderFiscalEntityResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *OrderFiscalEntityResponse) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *OrderFiscalEntityResponse) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


