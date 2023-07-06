# OrderUpdateRequestCustomerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Phone** | **string** |  | 
**Corporate** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CustomerId** | **string** |  | 

## Methods

### NewOrderUpdateRequestCustomerInfo

`func NewOrderUpdateRequestCustomerInfo(name string, email string, phone string, customerId string, ) *OrderUpdateRequestCustomerInfo`

NewOrderUpdateRequestCustomerInfo instantiates a new OrderUpdateRequestCustomerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateRequestCustomerInfoWithDefaults

`func NewOrderUpdateRequestCustomerInfoWithDefaults() *OrderUpdateRequestCustomerInfo`

NewOrderUpdateRequestCustomerInfoWithDefaults instantiates a new OrderUpdateRequestCustomerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrderUpdateRequestCustomerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderUpdateRequestCustomerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderUpdateRequestCustomerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *OrderUpdateRequestCustomerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderUpdateRequestCustomerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderUpdateRequestCustomerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *OrderUpdateRequestCustomerInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderUpdateRequestCustomerInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderUpdateRequestCustomerInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCorporate

`func (o *OrderUpdateRequestCustomerInfo) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *OrderUpdateRequestCustomerInfo) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *OrderUpdateRequestCustomerInfo) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *OrderUpdateRequestCustomerInfo) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *OrderUpdateRequestCustomerInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderUpdateRequestCustomerInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderUpdateRequestCustomerInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderUpdateRequestCustomerInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCustomerId

`func (o *OrderUpdateRequestCustomerInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OrderUpdateRequestCustomerInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OrderUpdateRequestCustomerInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


