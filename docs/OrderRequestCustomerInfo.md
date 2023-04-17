# OrderRequestCustomerInfo

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

### NewOrderRequestCustomerInfo

`func NewOrderRequestCustomerInfo(name string, email string, phone string, customerId string, ) *OrderRequestCustomerInfo`

NewOrderRequestCustomerInfo instantiates a new OrderRequestCustomerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestCustomerInfoWithDefaults

`func NewOrderRequestCustomerInfoWithDefaults() *OrderRequestCustomerInfo`

NewOrderRequestCustomerInfoWithDefaults instantiates a new OrderRequestCustomerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrderRequestCustomerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderRequestCustomerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderRequestCustomerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *OrderRequestCustomerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderRequestCustomerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderRequestCustomerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *OrderRequestCustomerInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderRequestCustomerInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderRequestCustomerInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCorporate

`func (o *OrderRequestCustomerInfo) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *OrderRequestCustomerInfo) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *OrderRequestCustomerInfo) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *OrderRequestCustomerInfo) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *OrderRequestCustomerInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderRequestCustomerInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderRequestCustomerInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderRequestCustomerInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCustomerId

`func (o *OrderRequestCustomerInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OrderRequestCustomerInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OrderRequestCustomerInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


