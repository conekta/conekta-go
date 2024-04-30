# OrderResponseShippingContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**BetweenStreets** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**CustomerShippingContactsResponseAddress**](CustomerShippingContactsResponseAddress.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the shipping contact | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrderResponseShippingContact

`func NewOrderResponseShippingContact() *OrderResponseShippingContact`

NewOrderResponseShippingContact instantiates a new OrderResponseShippingContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseShippingContactWithDefaults

`func NewOrderResponseShippingContactWithDefaults() *OrderResponseShippingContact`

NewOrderResponseShippingContactWithDefaults instantiates a new OrderResponseShippingContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *OrderResponseShippingContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderResponseShippingContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderResponseShippingContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderResponseShippingContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *OrderResponseShippingContact) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *OrderResponseShippingContact) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *OrderResponseShippingContact) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *OrderResponseShippingContact) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *OrderResponseShippingContact) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *OrderResponseShippingContact) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *OrderResponseShippingContact) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *OrderResponseShippingContact) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### SetBetweenStreetsNil

`func (o *OrderResponseShippingContact) SetBetweenStreetsNil(b bool)`

 SetBetweenStreetsNil sets the value for BetweenStreets to be an explicit nil

### UnsetBetweenStreets
`func (o *OrderResponseShippingContact) UnsetBetweenStreets()`

UnsetBetweenStreets ensures that no value is present for BetweenStreets, not even an explicit nil
### GetAddress

`func (o *OrderResponseShippingContact) GetAddress() CustomerShippingContactsResponseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderResponseShippingContact) GetAddressOk() (*CustomerShippingContactsResponseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderResponseShippingContact) SetAddress(v CustomerShippingContactsResponseAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrderResponseShippingContact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetParentId

`func (o *OrderResponseShippingContact) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *OrderResponseShippingContact) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *OrderResponseShippingContact) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *OrderResponseShippingContact) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *OrderResponseShippingContact) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OrderResponseShippingContact) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OrderResponseShippingContact) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OrderResponseShippingContact) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *OrderResponseShippingContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderResponseShippingContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderResponseShippingContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderResponseShippingContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderResponseShippingContact) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderResponseShippingContact) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderResponseShippingContact) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderResponseShippingContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderResponseShippingContact) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderResponseShippingContact) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderResponseShippingContact) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderResponseShippingContact) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *OrderResponseShippingContact) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderResponseShippingContact) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderResponseShippingContact) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderResponseShippingContact) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDeleted

`func (o *OrderResponseShippingContact) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrderResponseShippingContact) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrderResponseShippingContact) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OrderResponseShippingContact) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


