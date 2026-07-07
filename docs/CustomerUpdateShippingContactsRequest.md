# CustomerUpdateShippingContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Phone contact | [optional] 
**Receiver** | Pointer to **string** | Name of the person who will receive the order | [optional] 
**BetweenStreets** | Pointer to **string** | The street names between which the order will be delivered. | [optional] 
**Address** | Pointer to [**CustomerShippingContactsRequestAddress**](CustomerShippingContactsRequestAddress.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerUpdateShippingContactsRequest

`func NewCustomerUpdateShippingContactsRequest() *CustomerUpdateShippingContactsRequest`

NewCustomerUpdateShippingContactsRequest instantiates a new CustomerUpdateShippingContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateShippingContactsRequestWithDefaults

`func NewCustomerUpdateShippingContactsRequestWithDefaults() *CustomerUpdateShippingContactsRequest`

NewCustomerUpdateShippingContactsRequestWithDefaults instantiates a new CustomerUpdateShippingContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CustomerUpdateShippingContactsRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerUpdateShippingContactsRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerUpdateShippingContactsRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerUpdateShippingContactsRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetReceiver

`func (o *CustomerUpdateShippingContactsRequest) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *CustomerUpdateShippingContactsRequest) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *CustomerUpdateShippingContactsRequest) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *CustomerUpdateShippingContactsRequest) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetBetweenStreets

`func (o *CustomerUpdateShippingContactsRequest) GetBetweenStreets() string`

GetBetweenStreets returns the BetweenStreets field if non-nil, zero value otherwise.

### GetBetweenStreetsOk

`func (o *CustomerUpdateShippingContactsRequest) GetBetweenStreetsOk() (*string, bool)`

GetBetweenStreetsOk returns a tuple with the BetweenStreets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenStreets

`func (o *CustomerUpdateShippingContactsRequest) SetBetweenStreets(v string)`

SetBetweenStreets sets BetweenStreets field to given value.

### HasBetweenStreets

`func (o *CustomerUpdateShippingContactsRequest) HasBetweenStreets() bool`

HasBetweenStreets returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerUpdateShippingContactsRequest) GetAddress() CustomerShippingContactsRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerUpdateShippingContactsRequest) GetAddressOk() (*CustomerShippingContactsRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerUpdateShippingContactsRequest) SetAddress(v CustomerShippingContactsRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerUpdateShippingContactsRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetParentId

`func (o *CustomerUpdateShippingContactsRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CustomerUpdateShippingContactsRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CustomerUpdateShippingContactsRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CustomerUpdateShippingContactsRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerUpdateShippingContactsRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerUpdateShippingContactsRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerUpdateShippingContactsRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerUpdateShippingContactsRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDeleted

`func (o *CustomerUpdateShippingContactsRequest) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CustomerUpdateShippingContactsRequest) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CustomerUpdateShippingContactsRequest) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CustomerUpdateShippingContactsRequest) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


