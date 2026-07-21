# SubscriptionResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleStart** | Pointer to **int64** |  | [optional] 
**BillingCycleEnd** | Pointer to **int64** |  | [optional] 
**CanceledAt** | Pointer to **int64** |  | [optional] 
**CanceledReason** | Pointer to **string** | Reason for cancellation. This field appears when the subscription status is &#39;canceled&#39;. | [optional] 
**CardId** | Pointer to **string** |  | [optional] 
**ChargeId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CustomerCustomReference** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastBillingCycleOrderId** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**PausedAt** | Pointer to **int64** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionStart** | Pointer to **int32** |  | [optional] 
**TrialStart** | Pointer to **int64** |  | [optional] 
**TrialEnd** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubscriptionResponse1

`func NewSubscriptionResponse1() *SubscriptionResponse1`

NewSubscriptionResponse1 instantiates a new SubscriptionResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResponse1WithDefaults

`func NewSubscriptionResponse1WithDefaults() *SubscriptionResponse1`

NewSubscriptionResponse1WithDefaults instantiates a new SubscriptionResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleStart

`func (o *SubscriptionResponse1) GetBillingCycleStart() int64`

GetBillingCycleStart returns the BillingCycleStart field if non-nil, zero value otherwise.

### GetBillingCycleStartOk

`func (o *SubscriptionResponse1) GetBillingCycleStartOk() (*int64, bool)`

GetBillingCycleStartOk returns a tuple with the BillingCycleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleStart

`func (o *SubscriptionResponse1) SetBillingCycleStart(v int64)`

SetBillingCycleStart sets BillingCycleStart field to given value.

### HasBillingCycleStart

`func (o *SubscriptionResponse1) HasBillingCycleStart() bool`

HasBillingCycleStart returns a boolean if a field has been set.

### GetBillingCycleEnd

`func (o *SubscriptionResponse1) GetBillingCycleEnd() int64`

GetBillingCycleEnd returns the BillingCycleEnd field if non-nil, zero value otherwise.

### GetBillingCycleEndOk

`func (o *SubscriptionResponse1) GetBillingCycleEndOk() (*int64, bool)`

GetBillingCycleEndOk returns a tuple with the BillingCycleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleEnd

`func (o *SubscriptionResponse1) SetBillingCycleEnd(v int64)`

SetBillingCycleEnd sets BillingCycleEnd field to given value.

### HasBillingCycleEnd

`func (o *SubscriptionResponse1) HasBillingCycleEnd() bool`

HasBillingCycleEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *SubscriptionResponse1) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SubscriptionResponse1) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SubscriptionResponse1) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *SubscriptionResponse1) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCanceledReason

`func (o *SubscriptionResponse1) GetCanceledReason() string`

GetCanceledReason returns the CanceledReason field if non-nil, zero value otherwise.

### GetCanceledReasonOk

`func (o *SubscriptionResponse1) GetCanceledReasonOk() (*string, bool)`

GetCanceledReasonOk returns a tuple with the CanceledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledReason

`func (o *SubscriptionResponse1) SetCanceledReason(v string)`

SetCanceledReason sets CanceledReason field to given value.

### HasCanceledReason

`func (o *SubscriptionResponse1) HasCanceledReason() bool`

HasCanceledReason returns a boolean if a field has been set.

### GetCardId

`func (o *SubscriptionResponse1) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *SubscriptionResponse1) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *SubscriptionResponse1) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *SubscriptionResponse1) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetChargeId

`func (o *SubscriptionResponse1) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *SubscriptionResponse1) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *SubscriptionResponse1) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *SubscriptionResponse1) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionResponse1) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionResponse1) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionResponse1) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionResponse1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerCustomReference

`func (o *SubscriptionResponse1) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *SubscriptionResponse1) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *SubscriptionResponse1) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *SubscriptionResponse1) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionResponse1) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionResponse1) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionResponse1) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionResponse1) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionResponse1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionResponse1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastBillingCycleOrderId

`func (o *SubscriptionResponse1) GetLastBillingCycleOrderId() string`

GetLastBillingCycleOrderId returns the LastBillingCycleOrderId field if non-nil, zero value otherwise.

### GetLastBillingCycleOrderIdOk

`func (o *SubscriptionResponse1) GetLastBillingCycleOrderIdOk() (*string, bool)`

GetLastBillingCycleOrderIdOk returns a tuple with the LastBillingCycleOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBillingCycleOrderId

`func (o *SubscriptionResponse1) SetLastBillingCycleOrderId(v string)`

SetLastBillingCycleOrderId sets LastBillingCycleOrderId field to given value.

### HasLastBillingCycleOrderId

`func (o *SubscriptionResponse1) HasLastBillingCycleOrderId() bool`

HasLastBillingCycleOrderId returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionResponse1) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionResponse1) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionResponse1) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionResponse1) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPausedAt

`func (o *SubscriptionResponse1) GetPausedAt() int64`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *SubscriptionResponse1) GetPausedAtOk() (*int64, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *SubscriptionResponse1) SetPausedAt(v int64)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *SubscriptionResponse1) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionResponse1) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionResponse1) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionResponse1) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionResponse1) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionResponse1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionResponse1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionResponse1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionResponse1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionStart

`func (o *SubscriptionResponse1) GetSubscriptionStart() int32`

GetSubscriptionStart returns the SubscriptionStart field if non-nil, zero value otherwise.

### GetSubscriptionStartOk

`func (o *SubscriptionResponse1) GetSubscriptionStartOk() (*int32, bool)`

GetSubscriptionStartOk returns a tuple with the SubscriptionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStart

`func (o *SubscriptionResponse1) SetSubscriptionStart(v int32)`

SetSubscriptionStart sets SubscriptionStart field to given value.

### HasSubscriptionStart

`func (o *SubscriptionResponse1) HasSubscriptionStart() bool`

HasSubscriptionStart returns a boolean if a field has been set.

### GetTrialStart

`func (o *SubscriptionResponse1) GetTrialStart() int64`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *SubscriptionResponse1) GetTrialStartOk() (*int64, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *SubscriptionResponse1) SetTrialStart(v int64)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *SubscriptionResponse1) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetTrialEnd

`func (o *SubscriptionResponse1) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionResponse1) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionResponse1) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionResponse1) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


