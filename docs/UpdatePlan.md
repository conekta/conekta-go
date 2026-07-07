# UpdatePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount in cents that will be charged on the interval specified. | [optional] 
**Currency** | Pointer to **string** | ISO 4217 for currencies, for the Mexican peso it is MXN/USD | [optional] 
**ExpiryCount** | Pointer to **int32** | Number of repetitions of the frequency NUMBER OF CHARGES TO BE MADE, considering the interval and frequency, this evolves over time, but is subject to the expiration count. | [optional] 
**Name** | Pointer to **string** | The name of the plan. | [optional] 

## Methods

### NewUpdatePlan

`func NewUpdatePlan() *UpdatePlan`

NewUpdatePlan instantiates a new UpdatePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlanWithDefaults

`func NewUpdatePlanWithDefaults() *UpdatePlan`

NewUpdatePlanWithDefaults instantiates a new UpdatePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdatePlan) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdatePlan) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdatePlan) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdatePlan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdatePlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdatePlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdatePlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdatePlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiryCount

`func (o *UpdatePlan) GetExpiryCount() int32`

GetExpiryCount returns the ExpiryCount field if non-nil, zero value otherwise.

### GetExpiryCountOk

`func (o *UpdatePlan) GetExpiryCountOk() (*int32, bool)`

GetExpiryCountOk returns a tuple with the ExpiryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryCount

`func (o *UpdatePlan) SetExpiryCount(v int32)`

SetExpiryCount sets ExpiryCount field to given value.

### HasExpiryCount

`func (o *UpdatePlan) HasExpiryCount() bool`

HasExpiryCount returns a boolean if a field has been set.

### GetName

`func (o *UpdatePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlan) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


