# PaymentMethodCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of payment method | 
**Cvc** | **string** | Card security code | 
**ExpMonth** | **string** | Card expiration month | 
**ExpYear** | **string** | Card expiration year | 
**Name** | **string** | Cardholder name | 
**Number** | **string** | Card number | 
**CustomerIpAddress** | Pointer to **string** | Optional field used to capture the customer&#39;s IP address for fraud prevention and security monitoring purposes | [optional] 

## Methods

### NewPaymentMethodCardRequest

`func NewPaymentMethodCardRequest(type_ string, cvc string, expMonth string, expYear string, name string, number string, ) *PaymentMethodCardRequest`

NewPaymentMethodCardRequest instantiates a new PaymentMethodCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardRequestWithDefaults

`func NewPaymentMethodCardRequestWithDefaults() *PaymentMethodCardRequest`

NewPaymentMethodCardRequestWithDefaults instantiates a new PaymentMethodCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCardRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCardRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCardRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCvc

`func (o *PaymentMethodCardRequest) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *PaymentMethodCardRequest) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *PaymentMethodCardRequest) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetExpMonth

`func (o *PaymentMethodCardRequest) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethodCardRequest) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethodCardRequest) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *PaymentMethodCardRequest) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethodCardRequest) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethodCardRequest) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.


### GetName

`func (o *PaymentMethodCardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodCardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodCardRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *PaymentMethodCardRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaymentMethodCardRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaymentMethodCardRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCustomerIpAddress

`func (o *PaymentMethodCardRequest) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *PaymentMethodCardRequest) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *PaymentMethodCardRequest) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *PaymentMethodCardRequest) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


