# CustomerShippingContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone contact | [optional] 
**Receiver** | Pointer to **string** | Name of the person who will receive the order | [optional] 
**BetweenStreets** | Pointer to **string** | The street names between which the order will be delivered. | [optional] 
**Address** | [**CustomerShippingContactsAddress**](CustomerShippingContactsAddress.md) |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the shipping contact | [optional] 

## Methods

### NewCustomerShippingContacts

`func NewCustomerShippingContacts(address CustomerShippingContactsAddress, ) *CustomerShippingContacts`

NewCustomerShippingContacts instantiates a new CustomerShippingContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerShippingContactsWithDefaults

`func NewCustomerShippingContactsWithDefaults() *CustomerShippingContacts`

NewCustomerShippingContactsWithDefaults instantiates a new CustomerShippingContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerShippingContacts) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerShippingContacts) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerShippingContacts) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerShippingContacts) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerShippingContacts) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerShippingContacts) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerShippingContacts) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerShippingContacts) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerShippingContacts) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerShippingContacts) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerShippingContacts) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerShippingContacts) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerShippingContacts) GetAddress() CustomerShippingContactsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerShippingContacts) GetAddressOk() (*CustomerShippingContactsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerShippingContacts) SetAddress(v CustomerShippingContactsAddress)`

SetAddress sets Address field to given value.


### GetParentId

`func (o *CustomerShippingContacts) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerShippingContacts) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerShippingContacts) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerShippingContacts) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerShippingContacts) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerShippingContacts) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerShippingContacts) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerShippingContacts) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CustomerShippingContacts) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CustomerShippingContacts) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDeleted

`func (o *CustomerShippingContacts) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerShippingContacts) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerShippingContacts) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerShippingContacts) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *CustomerShippingContacts) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *CustomerShippingContacts) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetMetadata

`func (o *CustomerShippingContacts) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerShippingContacts) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerShippingContacts) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerShippingContacts) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


