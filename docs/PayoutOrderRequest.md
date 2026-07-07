# PayoutOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPayoutMethods** | **[]string** | The payout methods that are allowed for the payout order. | 
**Amount** | **int64** | The amount of the payout order. | 
**Currency** | **string** | The currency in which the payout order is made. | [default to "MXN"]
**CustomerInfo** | [**PayoutOrderRequestCustomerInfo**](PayoutOrderRequestCustomerInfo.md) |  | 
**ExpiresAt** | **int64** | The expiration time of the payout order in Unix timestamp. | 
**Metadata** | Pointer to **map[string]interface{}** | The metadata of the payout order. | [optional] 
**Payout** | [**Payout**](Payout.md) |  | 
**Reason** | **string** | The reason for the payout order. | 

## Methods

### NewPayoutOrderRequest

`func NewPayoutOrderRequest(allowedPayoutMethods []string, amount int64, currency string, customerInfo PayoutOrderRequestCustomerInfo, expiresAt int64, payout Payout, reason string, ) *PayoutOrderRequest`

NewPayoutOrderRequest instantiates a new PayoutOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOrderRequestWithDefaults

`func NewPayoutOrderRequestWithDefaults() *PayoutOrderRequest`

NewPayoutOrderRequestWithDefaults instantiates a new PayoutOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPayoutMethods

`func (o *PayoutOrderRequest) GetAllowedPayoutMethods() []string`

GetAllowedPayoutMethods returns the AllowedPayoutMethods field if non-nil, zero value otherwise.

### GetAllowedPayoutMethodsOk

`func (o *PayoutOrderRequest) GetAllowedPayoutMethodsOk() (*[]string, bool)`

GetAllowedPayoutMethodsOk returns a tuple with the AllowedPayoutMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPayoutMethods

`func (o *PayoutOrderRequest) SetAllowedPayoutMethods(v []string)`

SetAllowedPayoutMethods sets AllowedPayoutMethods field to given value.


### GetAmount

`func (o *PayoutOrderRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutOrderRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutOrderRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PayoutOrderRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayoutOrderRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayoutOrderRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerInfo

`func (o *PayoutOrderRequest) GetCustomerInfo() PayoutOrderRequestCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *PayoutOrderRequest) GetCustomerInfoOk() (*PayoutOrderRequestCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *PayoutOrderRequest) SetCustomerInfo(v PayoutOrderRequestCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.


### GetExpiresAt

`func (o *PayoutOrderRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PayoutOrderRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PayoutOrderRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetMetadata

`func (o *PayoutOrderRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PayoutOrderRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PayoutOrderRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PayoutOrderRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPayout

`func (o *PayoutOrderRequest) GetPayout() Payout`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *PayoutOrderRequest) GetPayoutOk() (*Payout, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *PayoutOrderRequest) SetPayout(v Payout)`

SetPayout sets Payout field to given value.


### GetReason

`func (o *PayoutOrderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PayoutOrderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PayoutOrderRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


