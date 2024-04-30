# PayoutOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPayoutMethods** | **[]string** | The payout methods that are allowed for the payout order. | 
**Amount** | **int32** | The amount of the payout order. | 
**CreatedAt** | **int64** | The creation date of the payout order. | 
**Currency** | **string** | The currency in which the payout order is made. | [default to "MXN"]
**CustomerInfo** | [**PayoutOrderResponseCustomerInfo**](PayoutOrderResponseCustomerInfo.md) |  | 
**ExpiresAt** | Pointer to **int64** | The expiration date of the payout order. | [optional] 
**Id** | **string** | The id of the payout order. | 
**Livemode** | **bool** | The live mode of the payout order. | 
**Object** | **string** | The object of the payout order. | 
**Metadata** | Pointer to **map[string]interface{}** | The metadata of the payout order. | [optional] 
**Payouts** | [**[]PayoutOrderPayoutsItem**](PayoutOrderPayoutsItem.md) | The payout information of the payout order. | 
**Reason** | **string** | The reason for the payout order. | 
**Status** | Pointer to **string** | The status of the payout order. | [optional] 
**UpdatedAt** | **int64** | The update date of the payout order. | 

## Methods

### NewPayoutOrderResponse

`func NewPayoutOrderResponse(allowedPayoutMethods []string, amount int32, createdAt int64, currency string, customerInfo PayoutOrderResponseCustomerInfo, id string, livemode bool, object string, payouts []PayoutOrderPayoutsItem, reason string, updatedAt int64, ) *PayoutOrderResponse`

NewPayoutOrderResponse instantiates a new PayoutOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOrderResponseWithDefaults

`func NewPayoutOrderResponseWithDefaults() *PayoutOrderResponse`

NewPayoutOrderResponseWithDefaults instantiates a new PayoutOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPayoutMethods

`func (o *PayoutOrderResponse) GetAllowedPayoutMethods() []string`

GetAllowedPayoutMethods returns the AllowedPayoutMethods field if non-nil, zero value otherwise.

### GetAllowedPayoutMethodsOk

`func (o *PayoutOrderResponse) GetAllowedPayoutMethodsOk() (*[]string, bool)`

GetAllowedPayoutMethodsOk returns a tuple with the AllowedPayoutMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPayoutMethods

`func (o *PayoutOrderResponse) SetAllowedPayoutMethods(v []string)`

SetAllowedPayoutMethods sets AllowedPayoutMethods field to given value.


### GetAmount

`func (o *PayoutOrderResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutOrderResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutOrderResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreatedAt

`func (o *PayoutOrderResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PayoutOrderResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PayoutOrderResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *PayoutOrderResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayoutOrderResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayoutOrderResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerInfo

`func (o *PayoutOrderResponse) GetCustomerInfo() PayoutOrderResponseCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *PayoutOrderResponse) GetCustomerInfoOk() (*PayoutOrderResponseCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *PayoutOrderResponse) SetCustomerInfo(v PayoutOrderResponseCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.


### GetExpiresAt

`func (o *PayoutOrderResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PayoutOrderResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PayoutOrderResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PayoutOrderResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *PayoutOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutOrderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLivemode

`func (o *PayoutOrderResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *PayoutOrderResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *PayoutOrderResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetObject

`func (o *PayoutOrderResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PayoutOrderResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PayoutOrderResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetMetadata

`func (o *PayoutOrderResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PayoutOrderResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PayoutOrderResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PayoutOrderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPayouts

`func (o *PayoutOrderResponse) GetPayouts() []PayoutOrderPayoutsItem`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *PayoutOrderResponse) GetPayoutsOk() (*[]PayoutOrderPayoutsItem, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *PayoutOrderResponse) SetPayouts(v []PayoutOrderPayoutsItem)`

SetPayouts sets Payouts field to given value.


### GetReason

`func (o *PayoutOrderResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PayoutOrderResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PayoutOrderResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *PayoutOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PayoutOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PayoutOrderResponse) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PayoutOrderResponse) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PayoutOrderResponse) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


