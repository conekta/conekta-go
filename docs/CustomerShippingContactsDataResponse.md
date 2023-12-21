# CustomerShippingContactsDataResponse

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
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 

## Methods

### NewCustomerShippingContactsDataResponse

`func NewCustomerShippingContactsDataResponse(address CustomerShippingContactsAddress, id string, object string, createdAt int64, ) *CustomerShippingContactsDataResponse`

NewCustomerShippingContactsDataResponse instantiates a new CustomerShippingContactsDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerShippingContactsDataResponseWithDefaults

`func NewCustomerShippingContactsDataResponseWithDefaults() *CustomerShippingContactsDataResponse`

NewCustomerShippingContactsDataResponseWithDefaults instantiates a new CustomerShippingContactsDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerShippingContactsDataResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerShippingContactsDataResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerShippingContactsDataResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerShippingContactsDataResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerShippingContactsDataResponse) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerShippingContactsDataResponse) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerShippingContactsDataResponse) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerShippingContactsDataResponse) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerShippingContactsDataResponse) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerShippingContactsDataResponse) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerShippingContactsDataResponse) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerShippingContactsDataResponse) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerShippingContactsDataResponse) GetAddress() CustomerShippingContactsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerShippingContactsDataResponse) GetAddressOk() (*CustomerShippingContactsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerShippingContactsDataResponse) SetAddress(v CustomerShippingContactsAddress)`

SetAddress sets Address field to given value.


### GetParentId

`func (o *CustomerShippingContactsDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerShippingContactsDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerShippingContactsDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerShippingContactsDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerShippingContactsDataResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerShippingContactsDataResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerShippingContactsDataResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerShippingContactsDataResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CustomerShippingContactsDataResponse) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CustomerShippingContactsDataResponse) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDeleted

`func (o *CustomerShippingContactsDataResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerShippingContactsDataResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerShippingContactsDataResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerShippingContactsDataResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *CustomerShippingContactsDataResponse) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *CustomerShippingContactsDataResponse) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetMetadata

`func (o *CustomerShippingContactsDataResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerShippingContactsDataResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerShippingContactsDataResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerShippingContactsDataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *CustomerShippingContactsDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerShippingContactsDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerShippingContactsDataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CustomerShippingContactsDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerShippingContactsDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerShippingContactsDataResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *CustomerShippingContactsDataResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerShippingContactsDataResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerShippingContactsDataResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


