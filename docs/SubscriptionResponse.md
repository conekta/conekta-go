# SubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleStart** | Pointer to **NullableInt64** |  | [optional] 
**BillingCycleEnd** | Pointer to **NullableInt64** |  | [optional] 
**CanceledAt** | Pointer to **NullableInt64** |  | [optional] 
**CanceledReason** | Pointer to **string** | Reason for cancellation. This field appears when the subscription status is &#39;canceled&#39;. | [optional] 
**CardId** | Pointer to **string** |  | [optional] 
**ChargeId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CustomerCustomReference** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastBillingCycleOrderId** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**PausedAt** | Pointer to **NullableInt64** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionStart** | Pointer to **int32** |  | [optional] 
**TrialStart** | Pointer to **NullableInt64** |  | [optional] 
**TrialEnd** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewSubscriptionResponse

`func NewSubscriptionResponse() *SubscriptionResponse`

NewSubscriptionResponse instantiates a new SubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResponseWithDefaults

`func NewSubscriptionResponseWithDefaults() *SubscriptionResponse`

NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleStart

`func (o *SubscriptionResponse) GetBillingCycleStart() int64`

GetBillingCycleStart returns the BillingCycleStart field if non-nil, zero value otherwise.

### GetBillingCycleStartOk

`func (o *SubscriptionResponse) GetBillingCycleStartOk() (*int64, bool)`

GetBillingCycleStartOk returns a tuple with the BillingCycleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleStart

`func (o *SubscriptionResponse) SetBillingCycleStart(v int64)`

SetBillingCycleStart sets BillingCycleStart field to given value.

### HasBillingCycleStart

`func (o *SubscriptionResponse) HasBillingCycleStart() bool`

HasBillingCycleStart returns a boolean if a field has been set.

### SetBillingCycleStartNil

`func (o *SubscriptionResponse) SetBillingCycleStartNil(b bool)`

 SetBillingCycleStartNil sets the value for BillingCycleStart to be an explicit nil

### UnsetBillingCycleStart
`func (o *SubscriptionResponse) UnsetBillingCycleStart()`

UnsetBillingCycleStart ensures that no value is present for BillingCycleStart, not even an explicit nil
### GetBillingCycleEnd

`func (o *SubscriptionResponse) GetBillingCycleEnd() int64`

GetBillingCycleEnd returns the BillingCycleEnd field if non-nil, zero value otherwise.

### GetBillingCycleEndOk

`func (o *SubscriptionResponse) GetBillingCycleEndOk() (*int64, bool)`

GetBillingCycleEndOk returns a tuple with the BillingCycleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleEnd

`func (o *SubscriptionResponse) SetBillingCycleEnd(v int64)`

SetBillingCycleEnd sets BillingCycleEnd field to given value.

### HasBillingCycleEnd

`func (o *SubscriptionResponse) HasBillingCycleEnd() bool`

HasBillingCycleEnd returns a boolean if a field has been set.

### SetBillingCycleEndNil

`func (o *SubscriptionResponse) SetBillingCycleEndNil(b bool)`

 SetBillingCycleEndNil sets the value for BillingCycleEnd to be an explicit nil

### UnsetBillingCycleEnd
`func (o *SubscriptionResponse) UnsetBillingCycleEnd()`

UnsetBillingCycleEnd ensures that no value is present for BillingCycleEnd, not even an explicit nil
### GetCanceledAt

`func (o *SubscriptionResponse) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SubscriptionResponse) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SubscriptionResponse) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *SubscriptionResponse) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *SubscriptionResponse) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *SubscriptionResponse) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetCanceledReason

`func (o *SubscriptionResponse) GetCanceledReason() string`

GetCanceledReason returns the CanceledReason field if non-nil, zero value otherwise.

### GetCanceledReasonOk

`func (o *SubscriptionResponse) GetCanceledReasonOk() (*string, bool)`

GetCanceledReasonOk returns a tuple with the CanceledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledReason

`func (o *SubscriptionResponse) SetCanceledReason(v string)`

SetCanceledReason sets CanceledReason field to given value.

### HasCanceledReason

`func (o *SubscriptionResponse) HasCanceledReason() bool`

HasCanceledReason returns a boolean if a field has been set.

### GetCardId

`func (o *SubscriptionResponse) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *SubscriptionResponse) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *SubscriptionResponse) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *SubscriptionResponse) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetChargeId

`func (o *SubscriptionResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *SubscriptionResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *SubscriptionResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *SubscriptionResponse) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### SetChargeIdNil

`func (o *SubscriptionResponse) SetChargeIdNil(b bool)`

 SetChargeIdNil sets the value for ChargeId to be an explicit nil

### UnsetChargeId
`func (o *SubscriptionResponse) UnsetChargeId()`

UnsetChargeId ensures that no value is present for ChargeId, not even an explicit nil
### GetCreatedAt

`func (o *SubscriptionResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerCustomReference

`func (o *SubscriptionResponse) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *SubscriptionResponse) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *SubscriptionResponse) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *SubscriptionResponse) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastBillingCycleOrderId

`func (o *SubscriptionResponse) GetLastBillingCycleOrderId() string`

GetLastBillingCycleOrderId returns the LastBillingCycleOrderId field if non-nil, zero value otherwise.

### GetLastBillingCycleOrderIdOk

`func (o *SubscriptionResponse) GetLastBillingCycleOrderIdOk() (*string, bool)`

GetLastBillingCycleOrderIdOk returns a tuple with the LastBillingCycleOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBillingCycleOrderId

`func (o *SubscriptionResponse) SetLastBillingCycleOrderId(v string)`

SetLastBillingCycleOrderId sets LastBillingCycleOrderId field to given value.

### HasLastBillingCycleOrderId

`func (o *SubscriptionResponse) HasLastBillingCycleOrderId() bool`

HasLastBillingCycleOrderId returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPausedAt

`func (o *SubscriptionResponse) GetPausedAt() int64`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *SubscriptionResponse) GetPausedAtOk() (*int64, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *SubscriptionResponse) SetPausedAt(v int64)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *SubscriptionResponse) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### SetPausedAtNil

`func (o *SubscriptionResponse) SetPausedAtNil(b bool)`

 SetPausedAtNil sets the value for PausedAt to be an explicit nil

### UnsetPausedAt
`func (o *SubscriptionResponse) UnsetPausedAt()`

UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
### GetPlanId

`func (o *SubscriptionResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionStart

`func (o *SubscriptionResponse) GetSubscriptionStart() int32`

GetSubscriptionStart returns the SubscriptionStart field if non-nil, zero value otherwise.

### GetSubscriptionStartOk

`func (o *SubscriptionResponse) GetSubscriptionStartOk() (*int32, bool)`

GetSubscriptionStartOk returns a tuple with the SubscriptionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStart

`func (o *SubscriptionResponse) SetSubscriptionStart(v int32)`

SetSubscriptionStart sets SubscriptionStart field to given value.

### HasSubscriptionStart

`func (o *SubscriptionResponse) HasSubscriptionStart() bool`

HasSubscriptionStart returns a boolean if a field has been set.

### GetTrialStart

`func (o *SubscriptionResponse) GetTrialStart() int64`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *SubscriptionResponse) GetTrialStartOk() (*int64, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *SubscriptionResponse) SetTrialStart(v int64)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *SubscriptionResponse) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### SetTrialStartNil

`func (o *SubscriptionResponse) SetTrialStartNil(b bool)`

 SetTrialStartNil sets the value for TrialStart to be an explicit nil

### UnsetTrialStart
`func (o *SubscriptionResponse) UnsetTrialStart()`

UnsetTrialStart ensures that no value is present for TrialStart, not even an explicit nil
### GetTrialEnd

`func (o *SubscriptionResponse) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionResponse) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionResponse) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionResponse) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *SubscriptionResponse) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *SubscriptionResponse) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


