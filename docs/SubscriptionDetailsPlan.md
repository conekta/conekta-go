# SubscriptionDetailsPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **int32** |  | [optional] 
**TrialPeriodDays** | Pointer to **int32** |  | [optional] 
**ExpiryCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubscriptionDetailsPlan

`func NewSubscriptionDetailsPlan() *SubscriptionDetailsPlan`

NewSubscriptionDetailsPlan instantiates a new SubscriptionDetailsPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDetailsPlanWithDefaults

`func NewSubscriptionDetailsPlanWithDefaults() *SubscriptionDetailsPlan`

NewSubscriptionDetailsPlanWithDefaults instantiates a new SubscriptionDetailsPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionDetailsPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDetailsPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDetailsPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionDetailsPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionDetailsPlan) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionDetailsPlan) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionDetailsPlan) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionDetailsPlan) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionDetailsPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionDetailsPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionDetailsPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionDetailsPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAmount

`func (o *SubscriptionDetailsPlan) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionDetailsPlan) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionDetailsPlan) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SubscriptionDetailsPlan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionDetailsPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionDetailsPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionDetailsPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionDetailsPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *SubscriptionDetailsPlan) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SubscriptionDetailsPlan) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SubscriptionDetailsPlan) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *SubscriptionDetailsPlan) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetFrequency

`func (o *SubscriptionDetailsPlan) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SubscriptionDetailsPlan) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SubscriptionDetailsPlan) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SubscriptionDetailsPlan) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTrialPeriodDays

`func (o *SubscriptionDetailsPlan) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *SubscriptionDetailsPlan) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *SubscriptionDetailsPlan) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *SubscriptionDetailsPlan) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetExpiryCount

`func (o *SubscriptionDetailsPlan) GetExpiryCount() int32`

GetExpiryCount returns the ExpiryCount field if non-nil, zero value otherwise.

### GetExpiryCountOk

`func (o *SubscriptionDetailsPlan) GetExpiryCountOk() (*int32, bool)`

GetExpiryCountOk returns a tuple with the ExpiryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCount

`func (o *SubscriptionDetailsPlan) SetExpiryCount(v int32)`

SetExpiryCount sets ExpiryCount field to given value.

### HasExpiryCount

`func (o *SubscriptionDetailsPlan) HasExpiryCount() bool`

HasExpiryCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionDetailsPlan) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionDetailsPlan) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionDetailsPlan) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionDetailsPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


