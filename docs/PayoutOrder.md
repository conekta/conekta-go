# PayoutOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPayoutMethods** | **[]string** | The payout methods that are allowed for the payout order. | 
**Amount** | **int32** | The amount of the payout order. | 
**Currency** | **string** | The currency in which the payout order is made. | [default to "MXN"]
**CustomerInfo** | [**CustomerInfoJustCustomerId**](CustomerInfoJustCustomerId.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** | The metadata of the payout order. | [optional] 
**Payout** | [**Payout**](Payout.md) |  | 
**Reason** | **string** | The reason for the payout order. | 

## Methods

### NewPayoutOrder

`func NewPayoutOrder(allowedPayoutMethods []string, amount int32, currency string, customerInfo CustomerInfoJustCustomerId, payout Payout, reason string, ) *PayoutOrder`

NewPayoutOrder instantiates a new PayoutOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOrderWithDefaults

`func NewPayoutOrderWithDefaults() *PayoutOrder`

NewPayoutOrderWithDefaults instantiates a new PayoutOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPayoutMethods

`func (o *PayoutOrder) GetAllowedPayoutMethods() []string`

GetAllowedPayoutMethods returns the AllowedPayoutMethods field if non-nil, zero value otherwise.

### GetAllowedPayoutMethodsOk

`func (o *PayoutOrder) GetAllowedPayoutMethodsOk() (*[]string, bool)`

GetAllowedPayoutMethodsOk returns a tuple with the AllowedPayoutMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPayoutMethods

`func (o *PayoutOrder) SetAllowedPayoutMethods(v []string)`

SetAllowedPayoutMethods sets AllowedPayoutMethods field to given value.


### GetAmount

`func (o *PayoutOrder) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayoutOrder) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayoutOrder) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PayoutOrder) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayoutOrder) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayoutOrder) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerInfo

`func (o *PayoutOrder) GetCustomerInfo() CustomerInfoJustCustomerId`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *PayoutOrder) GetCustomerInfoOk() (*CustomerInfoJustCustomerId, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *PayoutOrder) SetCustomerInfo(v CustomerInfoJustCustomerId)`

SetCustomerInfo sets CustomerInfo field to given value.


### GetMetadata

`func (o *PayoutOrder) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PayoutOrder) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PayoutOrder) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PayoutOrder) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPayout

`func (o *PayoutOrder) GetPayout() Payout`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *PayoutOrder) GetPayoutOk() (*Payout, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *PayoutOrder) SetPayout(v Payout)`

SetPayout sets Payout field to given value.


### GetReason

`func (o *PayoutOrder) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PayoutOrder) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PayoutOrder) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


