# BalanceCommonField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The balance&#39;s amount | [optional] 
**Currency** | Pointer to **string** | The balance&#39;s currency | [optional] 

## Methods

### NewBalanceCommonField

`func NewBalanceCommonField() *BalanceCommonField`

NewBalanceCommonField instantiates a new BalanceCommonField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceCommonFieldWithDefaults

`func NewBalanceCommonFieldWithDefaults() *BalanceCommonField`

NewBalanceCommonFieldWithDefaults instantiates a new BalanceCommonField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BalanceCommonField) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BalanceCommonField) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BalanceCommonField) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BalanceCommonField) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *BalanceCommonField) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BalanceCommonField) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BalanceCommonField) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BalanceCommonField) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


