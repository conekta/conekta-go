# ChargeOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Channel** | Pointer to [**ChargeResponseChannel**](ChargeResponseChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceFingerprint** | Pointer to **NullableString** |  | [optional] 
**FailureCode** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**MonthlyInstallments** | Pointer to **NullableInt32** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**PaidAt** | Pointer to **NullableInt32** |  | [optional] 
**PaymentMethod** | Pointer to [**ChargeOrderResponsePaymentMethod**](ChargeOrderResponsePaymentMethod.md) |  | [optional] 
**ReferenceId** | Pointer to **NullableString** | Reference ID of the charge | [optional] 
**Refunds** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeOrderResponse

`func NewChargeOrderResponse() *ChargeOrderResponse`

NewChargeOrderResponse instantiates a new ChargeOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeOrderResponseWithDefaults

`func NewChargeOrderResponseWithDefaults() *ChargeOrderResponse`

NewChargeOrderResponseWithDefaults instantiates a new ChargeOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargeOrderResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargeOrderResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargeOrderResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ChargeOrderResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *ChargeOrderResponse) GetChannel() ChargeResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChargeOrderResponse) GetChannelOk() (*ChargeResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChargeOrderResponse) SetChannel(v ChargeResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChargeOrderResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargeOrderResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeOrderResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeOrderResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChargeOrderResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *ChargeOrderResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargeOrderResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargeOrderResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChargeOrderResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *ChargeOrderResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargeOrderResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargeOrderResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ChargeOrderResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *ChargeOrderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeOrderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeOrderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargeOrderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *ChargeOrderResponse) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *ChargeOrderResponse) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *ChargeOrderResponse) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *ChargeOrderResponse) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### SetDeviceFingerprintNil

`func (o *ChargeOrderResponse) SetDeviceFingerprintNil(b bool)`

 SetDeviceFingerprintNil sets the value for DeviceFingerprint to be an explicit nil

### UnsetDeviceFingerprint
`func (o *ChargeOrderResponse) UnsetDeviceFingerprint()`

UnsetDeviceFingerprint ensures that no value is present for DeviceFingerprint, not even an explicit nil
### GetFailureCode

`func (o *ChargeOrderResponse) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargeOrderResponse) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargeOrderResponse) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *ChargeOrderResponse) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ChargeOrderResponse) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargeOrderResponse) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargeOrderResponse) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ChargeOrderResponse) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *ChargeOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ChargeOrderResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ChargeOrderResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ChargeOrderResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ChargeOrderResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMonthlyInstallments

`func (o *ChargeOrderResponse) GetMonthlyInstallments() int32`

GetMonthlyInstallments returns the MonthlyInstallments field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOk

`func (o *ChargeOrderResponse) GetMonthlyInstallmentsOk() (*int32, bool)`

GetMonthlyInstallmentsOk returns a tuple with the MonthlyInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallments

`func (o *ChargeOrderResponse) SetMonthlyInstallments(v int32)`

SetMonthlyInstallments sets MonthlyInstallments field to given value.

### HasMonthlyInstallments

`func (o *ChargeOrderResponse) HasMonthlyInstallments() bool`

HasMonthlyInstallments returns a boolean if a field has been set.

### SetMonthlyInstallmentsNil

`func (o *ChargeOrderResponse) SetMonthlyInstallmentsNil(b bool)`

 SetMonthlyInstallmentsNil sets the value for MonthlyInstallments to be an explicit nil

### UnsetMonthlyInstallments
`func (o *ChargeOrderResponse) UnsetMonthlyInstallments()`

UnsetMonthlyInstallments ensures that no value is present for MonthlyInstallments, not even an explicit nil
### GetObject

`func (o *ChargeOrderResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeOrderResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeOrderResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChargeOrderResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrderId

`func (o *ChargeOrderResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChargeOrderResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChargeOrderResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ChargeOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPaidAt

`func (o *ChargeOrderResponse) GetPaidAt() int32`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *ChargeOrderResponse) GetPaidAtOk() (*int32, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *ChargeOrderResponse) SetPaidAt(v int32)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *ChargeOrderResponse) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *ChargeOrderResponse) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *ChargeOrderResponse) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetPaymentMethod

`func (o *ChargeOrderResponse) GetPaymentMethod() ChargeOrderResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChargeOrderResponse) GetPaymentMethodOk() (*ChargeOrderResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChargeOrderResponse) SetPaymentMethod(v ChargeOrderResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ChargeOrderResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetReferenceId

`func (o *ChargeOrderResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargeOrderResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargeOrderResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargeOrderResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ChargeOrderResponse) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ChargeOrderResponse) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRefunds

`func (o *ChargeOrderResponse) GetRefunds() []map[string]interface{}`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *ChargeOrderResponse) GetRefundsOk() (*[]map[string]interface{}, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *ChargeOrderResponse) SetRefunds(v []map[string]interface{})`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *ChargeOrderResponse) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetStatus

`func (o *ChargeOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargeOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


