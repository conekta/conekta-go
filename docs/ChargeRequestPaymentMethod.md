# ChargeRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of payment method | 
**ExpiresAt** | Pointer to **int64** | Method expiration date as unix timestamp | [optional] 
**ProductType** | **string** | Product type of the payment method, use for the payment method to know the product type | 
**CancelUrl** | **string** | URL to redirect the customer after a canceled payment | 
**CanNotExpire** | **bool** | Indicates if the payment method can not expire | 
**FailureUrl** | **string** | URL to redirect the customer after a failed payment | 
**SuccessUrl** | **string** | URL to redirect the customer after a successful payment | 
**Cvc** | **string** | Optional, It is a value that allows identifying the security code of the card. Only for PCI merchants | 
**ExpMonth** | **string** | Card expiration month | 
**ExpYear** | **string** | Card expiration year | 
**Name** | **string** | Cardholder name | 
**Number** | **string** | Card number | 
**CustomerIpAddress** | Pointer to **string** | Optional field used to capture the customer&#39;s IP address for fraud prevention and security monitoring purposes | [optional] 
**MonthlyInstallments** | Pointer to **int32** | How many months without interest to apply, it can be 3, 6, 9, 12 or 18 | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**PaymentSourceId** | Pointer to **string** |  | [optional] 
**ContractId** | Pointer to **string** | Optional id sent to indicate the bank contract for recurrent card charges. | [optional] 

## Methods

### NewChargeRequestPaymentMethod

`func NewChargeRequestPaymentMethod(type_ string, productType string, cancelUrl string, canNotExpire bool, failureUrl string, successUrl string, cvc string, expMonth string, expYear string, name string, number string, ) *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethod instantiates a new ChargeRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRequestPaymentMethodWithDefaults

`func NewChargeRequestPaymentMethodWithDefaults() *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethodWithDefaults instantiates a new ChargeRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargeRequestPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeRequestPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeRequestPaymentMethod) SetType(v string)`

SetType sets Type field to given value.


### GetExpiresAt

`func (o *ChargeRequestPaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeRequestPaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeRequestPaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeRequestPaymentMethod) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProductType

`func (o *ChargeRequestPaymentMethod) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ChargeRequestPaymentMethod) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ChargeRequestPaymentMethod) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetCancelUrl

`func (o *ChargeRequestPaymentMethod) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *ChargeRequestPaymentMethod) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *ChargeRequestPaymentMethod) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetCanNotExpire

`func (o *ChargeRequestPaymentMethod) GetCanNotExpire() bool`

GetCanNotExpire returns the CanNotExpire field if non-nil, zero value otherwise.

### GetCanNotExpireOk

`func (o *ChargeRequestPaymentMethod) GetCanNotExpireOk() (*bool, bool)`

GetCanNotExpireOk returns a tuple with the CanNotExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanNotExpire

`func (o *ChargeRequestPaymentMethod) SetCanNotExpire(v bool)`

SetCanNotExpire sets CanNotExpire field to given value.


### GetFailureUrl

`func (o *ChargeRequestPaymentMethod) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *ChargeRequestPaymentMethod) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *ChargeRequestPaymentMethod) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.


### GetSuccessUrl

`func (o *ChargeRequestPaymentMethod) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *ChargeRequestPaymentMethod) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *ChargeRequestPaymentMethod) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCvc

`func (o *ChargeRequestPaymentMethod) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *ChargeRequestPaymentMethod) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *ChargeRequestPaymentMethod) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetExpMonth

`func (o *ChargeRequestPaymentMethod) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *ChargeRequestPaymentMethod) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *ChargeRequestPaymentMethod) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *ChargeRequestPaymentMethod) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *ChargeRequestPaymentMethod) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *ChargeRequestPaymentMethod) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.


### GetName

`func (o *ChargeRequestPaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeRequestPaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeRequestPaymentMethod) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *ChargeRequestPaymentMethod) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ChargeRequestPaymentMethod) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ChargeRequestPaymentMethod) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCustomerIpAddress

`func (o *ChargeRequestPaymentMethod) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *ChargeRequestPaymentMethod) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *ChargeRequestPaymentMethod) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *ChargeRequestPaymentMethod) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.

### GetMonthlyInstallments

`func (o *ChargeRequestPaymentMethod) GetMonthlyInstallments() int32`

GetMonthlyInstallments returns the MonthlyInstallments field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOk

`func (o *ChargeRequestPaymentMethod) GetMonthlyInstallmentsOk() (*int32, bool)`

GetMonthlyInstallmentsOk returns a tuple with the MonthlyInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallments

`func (o *ChargeRequestPaymentMethod) SetMonthlyInstallments(v int32)`

SetMonthlyInstallments sets MonthlyInstallments field to given value.

### HasMonthlyInstallments

`func (o *ChargeRequestPaymentMethod) HasMonthlyInstallments() bool`

HasMonthlyInstallments returns a boolean if a field has been set.

### GetTokenId

`func (o *ChargeRequestPaymentMethod) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ChargeRequestPaymentMethod) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ChargeRequestPaymentMethod) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ChargeRequestPaymentMethod) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetPaymentSourceId

`func (o *ChargeRequestPaymentMethod) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *ChargeRequestPaymentMethod) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *ChargeRequestPaymentMethod) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.

### HasPaymentSourceId

`func (o *ChargeRequestPaymentMethod) HasPaymentSourceId() bool`

HasPaymentSourceId returns a boolean if a field has been set.

### GetContractId

`func (o *ChargeRequestPaymentMethod) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ChargeRequestPaymentMethod) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ChargeRequestPaymentMethod) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ChargeRequestPaymentMethod) HasContractId() bool`

HasContractId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


