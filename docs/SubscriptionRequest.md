# SubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**CardId** | Pointer to **string** |  | [optional] 
**TrialEnd** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionRequest

`func NewSubscriptionRequest(planId string, ) *SubscriptionRequest`

NewSubscriptionRequest instantiates a new SubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRequestWithDefaults

`func NewSubscriptionRequestWithDefaults() *SubscriptionRequest`

NewSubscriptionRequestWithDefaults instantiates a new SubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *SubscriptionRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetCardId

`func (o *SubscriptionRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *SubscriptionRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *SubscriptionRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *SubscriptionRequest) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetTrialEnd

`func (o *SubscriptionRequest) GetTrialEnd() int32`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionRequest) GetTrialEndOk() (*int32, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionRequest) SetTrialEnd(v int32)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


