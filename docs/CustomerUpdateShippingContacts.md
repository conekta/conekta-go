# CustomerUpdateShippingContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone contact | [optional] 
**Receiver** | Pointer to **string** | Name of the person who will receive the order | [optional] 
**BetweenStreets** | Pointer to **string** | The street names between which the order will be delivered. | [optional] 
**Address** | Pointer to [**CustomerShippingContactsAddress**](CustomerShippingContactsAddress.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCustomerUpdateShippingContacts

`func NewCustomerUpdateShippingContacts() *CustomerUpdateShippingContacts`

NewCustomerUpdateShippingContacts instantiates a new CustomerUpdateShippingContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateShippingContactsWithDefaults

`func NewCustomerUpdateShippingContactsWithDefaults() *CustomerUpdateShippingContacts`

NewCustomerUpdateShippingContactsWithDefaults instantiates a new CustomerUpdateShippingContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerUpdateShippingContacts) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerUpdateShippingContacts) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerUpdateShippingContacts) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerUpdateShippingContacts) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerUpdateShippingContacts) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerUpdateShippingContacts) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerUpdateShippingContacts) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerUpdateShippingContacts) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerUpdateShippingContacts) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerUpdateShippingContacts) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerUpdateShippingContacts) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerUpdateShippingContacts) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerUpdateShippingContacts) GetAddress() CustomerShippingContactsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerUpdateShippingContacts) GetAddressOk() (*CustomerShippingContactsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerUpdateShippingContacts) SetAddress(v CustomerShippingContactsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerUpdateShippingContacts) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetParentId

`func (o *CustomerUpdateShippingContacts) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerUpdateShippingContacts) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerUpdateShippingContacts) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerUpdateShippingContacts) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerUpdateShippingContacts) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerUpdateShippingContacts) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerUpdateShippingContacts) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerUpdateShippingContacts) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CustomerUpdateShippingContacts) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CustomerUpdateShippingContacts) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDeleted

`func (o *CustomerUpdateShippingContacts) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerUpdateShippingContacts) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerUpdateShippingContacts) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerUpdateShippingContacts) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *CustomerUpdateShippingContacts) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *CustomerUpdateShippingContacts) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


