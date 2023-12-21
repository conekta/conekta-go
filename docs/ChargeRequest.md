# ChargeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**PaymentMethod** | [**ChargeRequestPaymentMethod**](ChargeRequestPaymentMethod.md) |  | 
**ReferenceId** | Pointer to **string** | Custom reference to add to the charge | [optional] 

## Methods

### NewChargeRequest

`func NewChargeRequest(paymentMethod ChargeRequestPaymentMethod, ) *ChargeRequest`

NewChargeRequest instantiates a new ChargeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRequestWithDefaults

`func NewChargeRequestWithDefaults() *ChargeRequest`

NewChargeRequestWithDefaults instantiates a new ChargeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargeRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargeRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargeRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ChargeRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *ChargeRequest) GetPaymentMethod() ChargeRequestPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ChargeRequest) GetPaymentMethodOk() (*ChargeRequestPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ChargeRequest) SetPaymentMethod(v ChargeRequestPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetReferenceId

`func (o *ChargeRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargeRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargeRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargeRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


