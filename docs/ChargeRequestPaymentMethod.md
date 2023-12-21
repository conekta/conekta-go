# ChargeRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int64** | Method expiration date as unix timestamp | [optional] 
**MonthlyInstallments** | Pointer to **int32** | How many months without interest to apply, it can be 3, 6, 9, 12 or 18 | [optional] 
**Type** | **string** |  | 
**TokenId** | Pointer to **string** |  | [optional] 
**PaymentSourceId** | Pointer to **string** |  | [optional] 
**ContractId** | Pointer to **string** | Optional id sent to indicate the bank contract for recurrent card charges. | [optional] 

## Methods

### NewChargeRequestPaymentMethod

`func NewChargeRequestPaymentMethod(type_ string, ) *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethod instantiates a new ChargeRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRequestPaymentMethodWithDefaults

`func NewChargeRequestPaymentMethodWithDefaults() *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethodWithDefaults instantiates a new ChargeRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


