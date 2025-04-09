# PlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount in cents that will be charged on the interval specified. | 
**Currency** | Pointer to **string** | ISO 4217 for currencies, for the Mexican peso it is MXN/USD | [optional] 
**ExpiryCount** | Pointer to **int32** | Number of repetitions of the frequency NUMBER OF CHARGES TO BE MADE, considering the interval and frequency, this evolves over time, but is subject to the expiration count. | [optional] 
**Frequency** | **int32** | Frequency of the charge, which together with the interval, can be every 3 weeks, every 4 months, every 2 years, every 5 fortnights | 
**Id** | Pointer to **string** | internal reference id | [optional] 
**Interval** | **string** | The interval of time between each charge. | 
**Name** | **string** | The name of the plan. | 
**TrialPeriodDays** | Pointer to **int32** | The number of days the customer will have a free trial. | [optional] 
**MaxRetries** | Pointer to **int32** | (optional) Specifies the maximum number of retry attempts for a subscription payment before it is canceled. | [optional] 
**RetryDelayHours** | Pointer to **int32** | (optional)  Defines the number of hours between subscription payment retry attempts. | [optional] 

## Methods

### NewPlanRequest

`func NewPlanRequest(amount int32, frequency int32, interval string, name string, ) *PlanRequest`

NewPlanRequest instantiates a new PlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanRequestWithDefaults

`func NewPlanRequestWithDefaults() *PlanRequest`

NewPlanRequestWithDefaults instantiates a new PlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PlanRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlanRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlanRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PlanRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PlanRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiryCount

`func (o *PlanRequest) GetExpiryCount() int32`

GetExpiryCount returns the ExpiryCount field if non-nil, zero value otherwise.

### GetExpiryCountOk

`func (o *PlanRequest) GetExpiryCountOk() (*int32, bool)`

GetExpiryCountOk returns a tuple with the ExpiryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCount

`func (o *PlanRequest) SetExpiryCount(v int32)`

SetExpiryCount sets ExpiryCount field to given value.

### HasExpiryCount

`func (o *PlanRequest) HasExpiryCount() bool`

HasExpiryCount returns a boolean if a field has been set.

### GetFrequency

`func (o *PlanRequest) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PlanRequest) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PlanRequest) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *PlanRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *PlanRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetName

`func (o *PlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTrialPeriodDays

`func (o *PlanRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *PlanRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *PlanRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *PlanRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetMaxRetries

`func (o *PlanRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PlanRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PlanRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PlanRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetRetryDelayHours

`func (o *PlanRequest) GetRetryDelayHours() int32`

GetRetryDelayHours returns the RetryDelayHours field if non-nil, zero value otherwise.

### GetRetryDelayHoursOk

`func (o *PlanRequest) GetRetryDelayHoursOk() (*int32, bool)`

GetRetryDelayHoursOk returns a tuple with the RetryDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelayHours

`func (o *PlanRequest) SetRetryDelayHours(v int32)`

SetRetryDelayHours sets RetryDelayHours field to given value.

### HasRetryDelayHours

`func (o *PlanRequest) HasRetryDelayHours() bool`

HasRetryDelayHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


