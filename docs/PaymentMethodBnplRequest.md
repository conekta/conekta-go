# PaymentMethodBnplRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the payment method | 
**CancelUrl** | **string** | URL to redirect the customer after a canceled payment | 
**CanNotExpire** | **bool** | Indicates if the payment method can not expire | 
**FailureUrl** | **string** | URL to redirect the customer after a failed payment | 
**ProductType** | **string** | Product type of the payment method, use for the payment method to know the product type | 
**SuccessUrl** | **string** | URL to redirect the customer after a successful payment | 

## Methods

### NewPaymentMethodBnplRequest

`func NewPaymentMethodBnplRequest(type_ string, cancelUrl string, canNotExpire bool, failureUrl string, productType string, successUrl string, ) *PaymentMethodBnplRequest`

NewPaymentMethodBnplRequest instantiates a new PaymentMethodBnplRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodBnplRequestWithDefaults

`func NewPaymentMethodBnplRequestWithDefaults() *PaymentMethodBnplRequest`

NewPaymentMethodBnplRequestWithDefaults instantiates a new PaymentMethodBnplRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodBnplRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodBnplRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodBnplRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCancelUrl

`func (o *PaymentMethodBnplRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaymentMethodBnplRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaymentMethodBnplRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetCanNotExpire

`func (o *PaymentMethodBnplRequest) GetCanNotExpire() bool`

GetCanNotExpire returns the CanNotExpire field if non-nil, zero value otherwise.

### GetCanNotExpireOk

`func (o *PaymentMethodBnplRequest) GetCanNotExpireOk() (*bool, bool)`

GetCanNotExpireOk returns a tuple with the CanNotExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanNotExpire

`func (o *PaymentMethodBnplRequest) SetCanNotExpire(v bool)`

SetCanNotExpire sets CanNotExpire field to given value.


### GetFailureUrl

`func (o *PaymentMethodBnplRequest) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *PaymentMethodBnplRequest) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *PaymentMethodBnplRequest) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.


### GetProductType

`func (o *PaymentMethodBnplRequest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PaymentMethodBnplRequest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PaymentMethodBnplRequest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetSuccessUrl

`func (o *PaymentMethodBnplRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PaymentMethodBnplRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PaymentMethodBnplRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


