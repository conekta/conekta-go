# CustomerSubscriptionResponse

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

### NewCustomerSubscriptionResponse

`func NewCustomerSubscriptionResponse() *CustomerSubscriptionResponse`

NewCustomerSubscriptionResponse instantiates a new CustomerSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSubscriptionResponseWithDefaults

`func NewCustomerSubscriptionResponseWithDefaults() *CustomerSubscriptionResponse`

NewCustomerSubscriptionResponseWithDefaults instantiates a new CustomerSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleStart

`func (o *CustomerSubscriptionResponse) GetBillingCycleStart() int64`

GetBillingCycleStart returns the BillingCycleStart field if non-nil, zero value otherwise.

### GetBillingCycleStartOk

`func (o *CustomerSubscriptionResponse) GetBillingCycleStartOk() (*int64, bool)`

GetBillingCycleStartOk returns a tuple with the BillingCycleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleStart

`func (o *CustomerSubscriptionResponse) SetBillingCycleStart(v int64)`

SetBillingCycleStart sets BillingCycleStart field to given value.

### HasBillingCycleStart

`func (o *CustomerSubscriptionResponse) HasBillingCycleStart() bool`

HasBillingCycleStart returns a boolean if a field has been set.

### GetBillingCycleEnd

`func (o *CustomerSubscriptionResponse) GetBillingCycleEnd() int64`

GetBillingCycleEnd returns the BillingCycleEnd field if non-nil, zero value otherwise.

### GetBillingCycleEndOk

`func (o *CustomerSubscriptionResponse) GetBillingCycleEndOk() (*int64, bool)`

GetBillingCycleEndOk returns a tuple with the BillingCycleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleEnd

`func (o *CustomerSubscriptionResponse) SetBillingCycleEnd(v int64)`

SetBillingCycleEnd sets BillingCycleEnd field to given value.

### HasBillingCycleEnd

`func (o *CustomerSubscriptionResponse) HasBillingCycleEnd() bool`

HasBillingCycleEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *CustomerSubscriptionResponse) GetCanceledAt() int64`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *CustomerSubscriptionResponse) GetCanceledAtOk() (*int64, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *CustomerSubscriptionResponse) SetCanceledAt(v int64)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *CustomerSubscriptionResponse) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCanceledReason

`func (o *CustomerSubscriptionResponse) GetCanceledReason() string`

GetCanceledReason returns the CanceledReason field if non-nil, zero value otherwise.

### GetCanceledReasonOk

`func (o *CustomerSubscriptionResponse) GetCanceledReasonOk() (*string, bool)`

GetCanceledReasonOk returns a tuple with the CanceledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledReason

`func (o *CustomerSubscriptionResponse) SetCanceledReason(v string)`

SetCanceledReason sets CanceledReason field to given value.

### HasCanceledReason

`func (o *CustomerSubscriptionResponse) HasCanceledReason() bool`

HasCanceledReason returns a boolean if a field has been set.

### GetCardId

`func (o *CustomerSubscriptionResponse) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CustomerSubscriptionResponse) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CustomerSubscriptionResponse) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *CustomerSubscriptionResponse) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetChargeId

`func (o *CustomerSubscriptionResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *CustomerSubscriptionResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *CustomerSubscriptionResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *CustomerSubscriptionResponse) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerSubscriptionResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerSubscriptionResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerSubscriptionResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerSubscriptionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomerCustomReference

`func (o *CustomerSubscriptionResponse) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *CustomerSubscriptionResponse) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *CustomerSubscriptionResponse) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *CustomerSubscriptionResponse) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomerSubscriptionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerSubscriptionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerSubscriptionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerSubscriptionResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *CustomerSubscriptionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerSubscriptionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerSubscriptionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerSubscriptionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastBillingCycleOrderId

`func (o *CustomerSubscriptionResponse) GetLastBillingCycleOrderId() string`

GetLastBillingCycleOrderId returns the LastBillingCycleOrderId field if non-nil, zero value otherwise.

### GetLastBillingCycleOrderIdOk

`func (o *CustomerSubscriptionResponse) GetLastBillingCycleOrderIdOk() (*string, bool)`

GetLastBillingCycleOrderIdOk returns a tuple with the LastBillingCycleOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBillingCycleOrderId

`func (o *CustomerSubscriptionResponse) SetLastBillingCycleOrderId(v string)`

SetLastBillingCycleOrderId sets LastBillingCycleOrderId field to given value.

### HasLastBillingCycleOrderId

`func (o *CustomerSubscriptionResponse) HasLastBillingCycleOrderId() bool`

HasLastBillingCycleOrderId returns a boolean if a field has been set.

### GetObject

`func (o *CustomerSubscriptionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerSubscriptionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerSubscriptionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerSubscriptionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPausedAt

`func (o *CustomerSubscriptionResponse) GetPausedAt() int64`

GetPausedAt returns the PausedAt field if non-nil, zero value otherwise.

### GetPausedAtOk

`func (o *CustomerSubscriptionResponse) GetPausedAtOk() (*int64, bool)`

GetPausedAtOk returns a tuple with the PausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedAt

`func (o *CustomerSubscriptionResponse) SetPausedAt(v int64)`

SetPausedAt sets PausedAt field to given value.

### HasPausedAt

`func (o *CustomerSubscriptionResponse) HasPausedAt() bool`

HasPausedAt returns a boolean if a field has been set.

### GetPlanId

`func (o *CustomerSubscriptionResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CustomerSubscriptionResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CustomerSubscriptionResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CustomerSubscriptionResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerSubscriptionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerSubscriptionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerSubscriptionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerSubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionStart

`func (o *CustomerSubscriptionResponse) GetSubscriptionStart() int32`

GetSubscriptionStart returns the SubscriptionStart field if non-nil, zero value otherwise.

### GetSubscriptionStartOk

`func (o *CustomerSubscriptionResponse) GetSubscriptionStartOk() (*int32, bool)`

GetSubscriptionStartOk returns a tuple with the SubscriptionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStart

`func (o *CustomerSubscriptionResponse) SetSubscriptionStart(v int32)`

SetSubscriptionStart sets SubscriptionStart field to given value.

### HasSubscriptionStart

`func (o *CustomerSubscriptionResponse) HasSubscriptionStart() bool`

HasSubscriptionStart returns a boolean if a field has been set.

### GetTrialStart

`func (o *CustomerSubscriptionResponse) GetTrialStart() int64`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *CustomerSubscriptionResponse) GetTrialStartOk() (*int64, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *CustomerSubscriptionResponse) SetTrialStart(v int64)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *CustomerSubscriptionResponse) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetTrialEnd

`func (o *CustomerSubscriptionResponse) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *CustomerSubscriptionResponse) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *CustomerSubscriptionResponse) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *CustomerSubscriptionResponse) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


