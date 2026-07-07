# SubscriptionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**SubscriptionDetailsCard**](SubscriptionDetailsCard.md) |  | [optional] 
**Plan** | Pointer to [**SubscriptionDetailsPlan**](SubscriptionDetailsPlan.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**NextBillingCycle** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubscriptionDetails

`func NewSubscriptionDetails() *SubscriptionDetails`

NewSubscriptionDetails instantiates a new SubscriptionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDetailsWithDefaults

`func NewSubscriptionDetailsWithDefaults() *SubscriptionDetails`

NewSubscriptionDetailsWithDefaults instantiates a new SubscriptionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *SubscriptionDetails) GetCard() SubscriptionDetailsCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *SubscriptionDetails) GetCardOk() (*SubscriptionDetailsCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *SubscriptionDetails) SetCard(v SubscriptionDetailsCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *SubscriptionDetails) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriptionDetails) GetPlan() SubscriptionDetailsPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionDetails) GetPlanOk() (*SubscriptionDetailsPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionDetails) SetPlan(v SubscriptionDetailsPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionDetails) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionDetails) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionDetails) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionDetails) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionDetails) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionDetails) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionDetails) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionDetails) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionDetails) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetNextBillingCycle

`func (o *SubscriptionDetails) GetNextBillingCycle() int64`

GetNextBillingCycle returns the NextBillingCycle field if non-nil, zero value otherwise.

### GetNextBillingCycleOk

`func (o *SubscriptionDetails) GetNextBillingCycleOk() (*int64, bool)`

GetNextBillingCycleOk returns a tuple with the NextBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBillingCycle

`func (o *SubscriptionDetails) SetNextBillingCycle(v int64)`

SetNextBillingCycle sets NextBillingCycle field to given value.

### HasNextBillingCycle

`func (o *SubscriptionDetails) HasNextBillingCycle() bool`

HasNextBillingCycle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionDetails) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionDetails) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionDetails) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionDetails) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionDetails) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionDetails) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


