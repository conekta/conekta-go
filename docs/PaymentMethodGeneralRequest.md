# PaymentMethodGeneralRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int64** | Method expiration date as unix timestamp | [optional] 
**MonthlyInstallments** | Pointer to **int32** | How many months without interest to apply, it can be 3, 6, 9, 12 or 18 | [optional] 
**Type** | **string** | Type of payment method | 
**TokenId** | Pointer to **string** |  | [optional] 
**PaymentSourceId** | Pointer to **string** |  | [optional] 
**Cvc** | Pointer to **string** | Optional, It is a value that allows identifying the security code of the card. Only for PCI merchants | [optional] 
**ContractId** | Pointer to **string** | Optional id sent to indicate the bank contract for recurrent card charges. | [optional] 
**CustomerIpAddress** | Pointer to **string** | Optional field used to capture the customer&#39;s IP address for fraud prevention and security monitoring purposes | [optional] 

## Methods

### NewPaymentMethodGeneralRequest

`func NewPaymentMethodGeneralRequest(type_ string, ) *PaymentMethodGeneralRequest`

NewPaymentMethodGeneralRequest instantiates a new PaymentMethodGeneralRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodGeneralRequestWithDefaults

`func NewPaymentMethodGeneralRequestWithDefaults() *PaymentMethodGeneralRequest`

NewPaymentMethodGeneralRequestWithDefaults instantiates a new PaymentMethodGeneralRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *PaymentMethodGeneralRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodGeneralRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodGeneralRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodGeneralRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMonthlyInstallments

`func (o *PaymentMethodGeneralRequest) GetMonthlyInstallments() int32`

GetMonthlyInstallments returns the MonthlyInstallments field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOk

`func (o *PaymentMethodGeneralRequest) GetMonthlyInstallmentsOk() (*int32, bool)`

GetMonthlyInstallmentsOk returns a tuple with the MonthlyInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallments

`func (o *PaymentMethodGeneralRequest) SetMonthlyInstallments(v int32)`

SetMonthlyInstallments sets MonthlyInstallments field to given value.

### HasMonthlyInstallments

`func (o *PaymentMethodGeneralRequest) HasMonthlyInstallments() bool`

HasMonthlyInstallments returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodGeneralRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodGeneralRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodGeneralRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTokenId

`func (o *PaymentMethodGeneralRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentMethodGeneralRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentMethodGeneralRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentMethodGeneralRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetPaymentSourceId

`func (o *PaymentMethodGeneralRequest) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *PaymentMethodGeneralRequest) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *PaymentMethodGeneralRequest) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.

### HasPaymentSourceId

`func (o *PaymentMethodGeneralRequest) HasPaymentSourceId() bool`

HasPaymentSourceId returns a boolean if a field has been set.

### GetCvc

`func (o *PaymentMethodGeneralRequest) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *PaymentMethodGeneralRequest) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *PaymentMethodGeneralRequest) SetCvc(v string)`

SetCvc sets Cvc field to given value.

### HasCvc

`func (o *PaymentMethodGeneralRequest) HasCvc() bool`

HasCvc returns a boolean if a field has been set.

### GetContractId

`func (o *PaymentMethodGeneralRequest) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *PaymentMethodGeneralRequest) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *PaymentMethodGeneralRequest) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *PaymentMethodGeneralRequest) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCustomerIpAddress

`func (o *PaymentMethodGeneralRequest) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *PaymentMethodGeneralRequest) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *PaymentMethodGeneralRequest) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *PaymentMethodGeneralRequest) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


