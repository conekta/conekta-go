# PlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExpiryCount** | Pointer to **NullableInt32** |  | [optional] 
**Frequency** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPlanResponse

`func NewPlanResponse() *PlanResponse`

NewPlanResponse instantiates a new PlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResponseWithDefaults

`func NewPlanResponseWithDefaults() *PlanResponse`

NewPlanResponseWithDefaults instantiates a new PlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PlanResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlanResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlanResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PlanResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlanResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlanResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *PlanResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PlanResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiryCount

`func (o *PlanResponse) GetExpiryCount() int32`

GetExpiryCount returns the ExpiryCount field if non-nil, zero value otherwise.

### GetExpiryCountOk

`func (o *PlanResponse) GetExpiryCountOk() (*int32, bool)`

GetExpiryCountOk returns a tuple with the ExpiryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCount

`func (o *PlanResponse) SetExpiryCount(v int32)`

SetExpiryCount sets ExpiryCount field to given value.

### HasExpiryCount

`func (o *PlanResponse) HasExpiryCount() bool`

HasExpiryCount returns a boolean if a field has been set.

### SetExpiryCountNil

`func (o *PlanResponse) SetExpiryCountNil(b bool)`

 SetExpiryCountNil sets the value for ExpiryCount to be an explicit nil

### UnsetExpiryCount
`func (o *PlanResponse) UnsetExpiryCount()`

UnsetExpiryCount ensures that no value is present for ExpiryCount, not even an explicit nil
### GetFrequency

`func (o *PlanResponse) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PlanResponse) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PlanResponse) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PlanResponse) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetId

`func (o *PlanResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *PlanResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PlanResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLivemode

`func (o *PlanResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *PlanResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *PlanResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *PlanResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetName

`func (o *PlanResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *PlanResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PlanResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PlanResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PlanResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTrialPeriodDays

`func (o *PlanResponse) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *PlanResponse) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *PlanResponse) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *PlanResponse) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *PlanResponse) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *PlanResponse) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


