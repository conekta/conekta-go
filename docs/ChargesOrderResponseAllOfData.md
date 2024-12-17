# ChargesOrderResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Channel** | Pointer to [**ChargeResponseChannel**](ChargeResponseChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceFingerprint** | Pointer to **string** |  | [optional] 
**FailureCode** | Pointer to **string** |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Charge ID | [optional] 
**Livemode** | Pointer to **bool** | Whether the charge was made in live mode or not | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** | Order ID | [optional] 
**PaidAt** | Pointer to **NullableInt64** | Payment date | [optional] 
**PaymentMethod** | Pointer to [**ChargeResponsePaymentMethod**](ChargeResponsePaymentMethod.md) |  | [optional] 
**ReferenceId** | Pointer to **NullableString** | Reference ID of the charge | [optional] 
**Refunds** | Pointer to [**NullableChargeResponseRefunds**](ChargeResponseRefunds.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewChargesOrderResponseAllOfData

`func NewChargesOrderResponseAllOfData() *ChargesOrderResponseAllOfData`

NewChargesOrderResponseAllOfData instantiates a new ChargesOrderResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargesOrderResponseAllOfDataWithDefaults

`func NewChargesOrderResponseAllOfDataWithDefaults() *ChargesOrderResponseAllOfData`

NewChargesOrderResponseAllOfDataWithDefaults instantiates a new ChargesOrderResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargesOrderResponseAllOfData) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargesOrderResponseAllOfData) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargesOrderResponseAllOfData) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ChargesOrderResponseAllOfData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *ChargesOrderResponseAllOfData) GetChannel() ChargeResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChargesOrderResponseAllOfData) GetChannelOk() (*ChargeResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChargesOrderResponseAllOfData) SetChannel(v ChargeResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChargesOrderResponseAllOfData) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargesOrderResponseAllOfData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargesOrderResponseAllOfData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargesOrderResponseAllOfData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChargesOrderResponseAllOfData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *ChargesOrderResponseAllOfData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargesOrderResponseAllOfData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargesOrderResponseAllOfData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChargesOrderResponseAllOfData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *ChargesOrderResponseAllOfData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargesOrderResponseAllOfData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargesOrderResponseAllOfData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ChargesOrderResponseAllOfData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *ChargesOrderResponseAllOfData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargesOrderResponseAllOfData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargesOrderResponseAllOfData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargesOrderResponseAllOfData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *ChargesOrderResponseAllOfData) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *ChargesOrderResponseAllOfData) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *ChargesOrderResponseAllOfData) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *ChargesOrderResponseAllOfData) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetFailureCode

`func (o *ChargesOrderResponseAllOfData) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargesOrderResponseAllOfData) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargesOrderResponseAllOfData) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *ChargesOrderResponseAllOfData) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ChargesOrderResponseAllOfData) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargesOrderResponseAllOfData) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargesOrderResponseAllOfData) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ChargesOrderResponseAllOfData) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *ChargesOrderResponseAllOfData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargesOrderResponseAllOfData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargesOrderResponseAllOfData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargesOrderResponseAllOfData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ChargesOrderResponseAllOfData) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ChargesOrderResponseAllOfData) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ChargesOrderResponseAllOfData) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ChargesOrderResponseAllOfData) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *ChargesOrderResponseAllOfData) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargesOrderResponseAllOfData) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargesOrderResponseAllOfData) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChargesOrderResponseAllOfData) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrderId

`func (o *ChargesOrderResponseAllOfData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChargesOrderResponseAllOfData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChargesOrderResponseAllOfData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ChargesOrderResponseAllOfData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPaidAt

`func (o *ChargesOrderResponseAllOfData) GetPaidAt() int64`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *ChargesOrderResponseAllOfData) GetPaidAtOk() (*int64, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *ChargesOrderResponseAllOfData) SetPaidAt(v int64)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *ChargesOrderResponseAllOfData) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *ChargesOrderResponseAllOfData) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *ChargesOrderResponseAllOfData) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetPaymentMethod

`func (o *ChargesOrderResponseAllOfData) GetPaymentMethod() ChargeResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChargesOrderResponseAllOfData) GetPaymentMethodOk() (*ChargeResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChargesOrderResponseAllOfData) SetPaymentMethod(v ChargeResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ChargesOrderResponseAllOfData) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetReferenceId

`func (o *ChargesOrderResponseAllOfData) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargesOrderResponseAllOfData) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargesOrderResponseAllOfData) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargesOrderResponseAllOfData) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ChargesOrderResponseAllOfData) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ChargesOrderResponseAllOfData) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRefunds

`func (o *ChargesOrderResponseAllOfData) GetRefunds() ChargeResponseRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *ChargesOrderResponseAllOfData) GetRefundsOk() (*ChargeResponseRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *ChargesOrderResponseAllOfData) SetRefunds(v ChargeResponseRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *ChargesOrderResponseAllOfData) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### SetRefundsNil

`func (o *ChargesOrderResponseAllOfData) SetRefundsNil(b bool)`

 SetRefundsNil sets the value for Refunds to be an explicit nil

### UnsetRefunds
`func (o *ChargesOrderResponseAllOfData) UnsetRefunds()`

UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
### GetStatus

`func (o *ChargesOrderResponseAllOfData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargesOrderResponseAllOfData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargesOrderResponseAllOfData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargesOrderResponseAllOfData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


