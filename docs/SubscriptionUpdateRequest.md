# SubscriptionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** |  | [optional] 
**CardId** | Pointer to **string** |  | [optional] 
**TrialEnd** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionUpdateRequest

`func NewSubscriptionUpdateRequest() *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequest instantiates a new SubscriptionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateRequestWithDefaults

`func NewSubscriptionUpdateRequestWithDefaults() *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequestWithDefaults instantiates a new SubscriptionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *SubscriptionUpdateRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionUpdateRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionUpdateRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionUpdateRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetCardId

`func (o *SubscriptionUpdateRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *SubscriptionUpdateRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *SubscriptionUpdateRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *SubscriptionUpdateRequest) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetTrialEnd

`func (o *SubscriptionUpdateRequest) GetTrialEnd() int32`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionUpdateRequest) GetTrialEndOk() (*int32, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionUpdateRequest) SetTrialEnd(v int32)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionUpdateRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


