# ChargesDataResponse

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
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**PaidAt** | Pointer to **NullableInt32** |  | [optional] 
**PaymentMethod** | Pointer to [**ChargeResponsePaymentMethod**](ChargeResponsePaymentMethod.md) |  | [optional] 
**ReferenceId** | Pointer to **NullableString** | Reference ID of the charge | [optional] 
**Refunds** | Pointer to [**NullableChargeResponseRefunds**](ChargeResponseRefunds.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewChargesDataResponse

`func NewChargesDataResponse() *ChargesDataResponse`

NewChargesDataResponse instantiates a new ChargesDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargesDataResponseWithDefaults

`func NewChargesDataResponseWithDefaults() *ChargesDataResponse`

NewChargesDataResponseWithDefaults instantiates a new ChargesDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargesDataResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargesDataResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargesDataResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ChargesDataResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *ChargesDataResponse) GetChannel() ChargeResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChargesDataResponse) GetChannelOk() (*ChargeResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChargesDataResponse) SetChannel(v ChargeResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChargesDataResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargesDataResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargesDataResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargesDataResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChargesDataResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *ChargesDataResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargesDataResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargesDataResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChargesDataResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *ChargesDataResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargesDataResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargesDataResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ChargesDataResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *ChargesDataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargesDataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargesDataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargesDataResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *ChargesDataResponse) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *ChargesDataResponse) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *ChargesDataResponse) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *ChargesDataResponse) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetFailureCode

`func (o *ChargesDataResponse) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargesDataResponse) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargesDataResponse) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *ChargesDataResponse) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ChargesDataResponse) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargesDataResponse) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargesDataResponse) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ChargesDataResponse) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetId

`func (o *ChargesDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargesDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargesDataResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargesDataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ChargesDataResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ChargesDataResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ChargesDataResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ChargesDataResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *ChargesDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargesDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargesDataResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChargesDataResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrderId

`func (o *ChargesDataResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ChargesDataResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ChargesDataResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ChargesDataResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPaidAt

`func (o *ChargesDataResponse) GetPaidAt() int32`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *ChargesDataResponse) GetPaidAtOk() (*int32, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *ChargesDataResponse) SetPaidAt(v int32)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *ChargesDataResponse) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *ChargesDataResponse) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *ChargesDataResponse) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetPaymentMethod

`func (o *ChargesDataResponse) GetPaymentMethod() ChargeResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChargesDataResponse) GetPaymentMethodOk() (*ChargeResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChargesDataResponse) SetPaymentMethod(v ChargeResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ChargesDataResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetReferenceId

`func (o *ChargesDataResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargesDataResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargesDataResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargesDataResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *ChargesDataResponse) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *ChargesDataResponse) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRefunds

`func (o *ChargesDataResponse) GetRefunds() ChargeResponseRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *ChargesDataResponse) GetRefundsOk() (*ChargeResponseRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *ChargesDataResponse) SetRefunds(v ChargeResponseRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *ChargesDataResponse) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### SetRefundsNil

`func (o *ChargesDataResponse) SetRefundsNil(b bool)`

 SetRefundsNil sets the value for Refunds to be an explicit nil

### UnsetRefunds
`func (o *ChargesDataResponse) UnsetRefunds()`

UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
### GetStatus

`func (o *ChargesDataResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargesDataResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargesDataResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargesDataResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


