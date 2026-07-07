# CustomerShippingContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone contact | [optional] 
**Receiver** | Pointer to **string** | Name of the person who will receive the order | [optional] 
**BetweenStreets** | Pointer to **string** | The street names between which the order will be delivered. | [optional] 
**Address** | [**CustomerShippingContactsRequestAddress**](CustomerShippingContactsRequestAddress.md) |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the shipping contact | [optional] 

## Methods

### NewCustomerShippingContactsRequest

`func NewCustomerShippingContactsRequest(address CustomerShippingContactsRequestAddress, ) *CustomerShippingContactsRequest`

NewCustomerShippingContactsRequest instantiates a new CustomerShippingContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerShippingContactsRequestWithDefaults

`func NewCustomerShippingContactsRequestWithDefaults() *CustomerShippingContactsRequest`

NewCustomerShippingContactsRequestWithDefaults instantiates a new CustomerShippingContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerShippingContactsRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerShippingContactsRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerShippingContactsRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerShippingContactsRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerShippingContactsRequest) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerShippingContactsRequest) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerShippingContactsRequest) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerShippingContactsRequest) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerShippingContactsRequest) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerShippingContactsRequest) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerShippingContactsRequest) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerShippingContactsRequest) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerShippingContactsRequest) GetAddress() CustomerShippingContactsRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerShippingContactsRequest) GetAddressOk() (*CustomerShippingContactsRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerShippingContactsRequest) SetAddress(v CustomerShippingContactsRequestAddress)`

SetAddress sets Address field to given value.


### GetParentId

`func (o *CustomerShippingContactsRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerShippingContactsRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerShippingContactsRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerShippingContactsRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerShippingContactsRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerShippingContactsRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerShippingContactsRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerShippingContactsRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDeleted

`func (o *CustomerShippingContactsRequest) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerShippingContactsRequest) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerShippingContactsRequest) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerShippingContactsRequest) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerShippingContactsRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerShippingContactsRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerShippingContactsRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerShippingContactsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


