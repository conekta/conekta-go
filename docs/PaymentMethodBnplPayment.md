# PaymentMethodBnplPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**CancelUrl** | Pointer to **string** | URL to redirect the customer after a canceled payment | [optional] 
**ExpiresAt** | **int64** | Expiration date of the charge | 
**FailureUrl** | Pointer to **string** | URL to redirect the customer after a failed payment | [optional] 
**ProductType** | **string** | Product type of the charge | 
**RedirectUrl** | Pointer to **string** | URL to redirect the customer to complete the payment | [optional] 
**SuccessUrl** | Pointer to **string** | URL to redirect the customer after a successful payment | [optional] 

## Methods

### NewPaymentMethodBnplPayment

`func NewPaymentMethodBnplPayment(object string, expiresAt int64, productType string, ) *PaymentMethodBnplPayment`

NewPaymentMethodBnplPayment instantiates a new PaymentMethodBnplPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodBnplPaymentWithDefaults

`func NewPaymentMethodBnplPaymentWithDefaults() *PaymentMethodBnplPayment`

NewPaymentMethodBnplPaymentWithDefaults instantiates a new PaymentMethodBnplPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodBnplPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodBnplPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodBnplPayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodBnplPayment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodBnplPayment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodBnplPayment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodBnplPayment) SetObject(v string)`

SetObject sets Object field to given value.


### GetCancelUrl

`func (o *PaymentMethodBnplPayment) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaymentMethodBnplPayment) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaymentMethodBnplPayment) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *PaymentMethodBnplPayment) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodBnplPayment) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodBnplPayment) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodBnplPayment) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFailureUrl

`func (o *PaymentMethodBnplPayment) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *PaymentMethodBnplPayment) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *PaymentMethodBnplPayment) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *PaymentMethodBnplPayment) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetProductType

`func (o *PaymentMethodBnplPayment) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PaymentMethodBnplPayment) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PaymentMethodBnplPayment) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRedirectUrl

`func (o *PaymentMethodBnplPayment) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentMethodBnplPayment) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentMethodBnplPayment) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PaymentMethodBnplPayment) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *PaymentMethodBnplPayment) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PaymentMethodBnplPayment) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PaymentMethodBnplPayment) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *PaymentMethodBnplPayment) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


