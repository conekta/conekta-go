# CustomerShippingContactsResponse

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

### NewCustomerShippingContactsResponse

`func NewCustomerShippingContactsResponse() *CustomerShippingContactsResponse`

NewCustomerShippingContactsResponse instantiates a new CustomerShippingContactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerShippingContactsResponseWithDefaults

`func NewCustomerShippingContactsResponseWithDefaults() *CustomerShippingContactsResponse`

NewCustomerShippingContactsResponseWithDefaults instantiates a new CustomerShippingContactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerShippingContactsResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerShippingContactsResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerShippingContactsResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerShippingContactsResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerShippingContactsResponse) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerShippingContactsResponse) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerShippingContactsResponse) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerShippingContactsResponse) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerShippingContactsResponse) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerShippingContactsResponse) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerShippingContactsResponse) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerShippingContactsResponse) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### SetBetweenStreetsNil

`func (o *CustomerShippingContactsResponse) SetBetweenStreetsNil(b bool)`

 SetBetweenStreetsNil sets the value for BetweenStreets to be an explicit nil

### UnsetBetweenStreets
`func (o *CustomerShippingContactsResponse) UnsetBetweenStreets()`

UnsetBetweenStreets ensures that no value is present for BetweenStreets, not even an explicit nil
### GetAddress

`func (o *CustomerShippingContactsResponse) GetAddress() CustomerShippingContactsResponseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerShippingContactsResponse) GetAddressOk() (*CustomerShippingContactsResponseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerShippingContactsResponse) SetAddress(v CustomerShippingContactsResponseAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerShippingContactsResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetParentId

`func (o *CustomerShippingContactsResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerShippingContactsResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerShippingContactsResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerShippingContactsResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerShippingContactsResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerShippingContactsResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerShippingContactsResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerShippingContactsResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *CustomerShippingContactsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerShippingContactsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerShippingContactsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerShippingContactsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerShippingContactsResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerShippingContactsResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerShippingContactsResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerShippingContactsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerShippingContactsResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerShippingContactsResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerShippingContactsResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerShippingContactsResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *CustomerShippingContactsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerShippingContactsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerShippingContactsResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerShippingContactsResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDeleted

`func (o *CustomerShippingContactsResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerShippingContactsResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerShippingContactsResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerShippingContactsResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


