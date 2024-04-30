# PayoutOrderPayoutsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the payout. | 
**Currency** | **string** | The currency in which the payout is made. | 
**ExpiresAt** | Pointer to **int64** | The expiration date of the payout. | [optional] 
**Id** | **string** | The id of the payout. | 
**Livemode** | **bool** | The live mode of the payout. | 
**Object** | **string** | The object of the payout. | 
**PayoutOrderId** | Pointer to **string** | The id of the payout order. | [optional] 
**Status** | Pointer to **string** | The status of the payout. | [optional] 

## Methods

### NewPayoutOrderPayoutsItem

`func NewPayoutOrderPayoutsItem(amount int32, currency string, id string, livemode bool, object string, ) *PayoutOrderPayoutsItem`

NewPayoutOrderPayoutsItem instantiates a new PayoutOrderPayoutsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOrderPayoutsItemWithDefaults

`func NewPayoutOrderPayoutsItemWithDefaults() *PayoutOrderPayoutsItem`

NewPayoutOrderPayoutsItemWithDefaults instantiates a new PayoutOrderPayoutsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PayoutOrderPayoutsItem) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutOrderPayoutsItem) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutOrderPayoutsItem) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PayoutOrderPayoutsItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayoutOrderPayoutsItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayoutOrderPayoutsItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExpiresAt

`func (o *PayoutOrderPayoutsItem) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PayoutOrderPayoutsItem) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PayoutOrderPayoutsItem) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PayoutOrderPayoutsItem) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *PayoutOrderPayoutsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutOrderPayoutsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutOrderPayoutsItem) SetId(v string)`

SetId sets Id field to given value.


### GetLivemode

`func (o *PayoutOrderPayoutsItem) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *PayoutOrderPayoutsItem) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *PayoutOrderPayoutsItem) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetObject

`func (o *PayoutOrderPayoutsItem) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PayoutOrderPayoutsItem) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PayoutOrderPayoutsItem) SetObject(v string)`

SetObject sets Object field to given value.


### GetPayoutOrderId

`func (o *PayoutOrderPayoutsItem) GetPayoutOrderId() string`

GetPayoutOrderId returns the PayoutOrderId field if non-nil, zero value otherwise.

### GetPayoutOrderIdOk

`func (o *PayoutOrderPayoutsItem) GetPayoutOrderIdOk() (*string, bool)`

GetPayoutOrderIdOk returns a tuple with the PayoutOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutOrderId

`func (o *PayoutOrderPayoutsItem) SetPayoutOrderId(v string)`

SetPayoutOrderId sets PayoutOrderId field to given value.

### HasPayoutOrderId

`func (o *PayoutOrderPayoutsItem) HasPayoutOrderId() bool`

HasPayoutOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutOrderPayoutsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutOrderPayoutsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutOrderPayoutsItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PayoutOrderPayoutsItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


