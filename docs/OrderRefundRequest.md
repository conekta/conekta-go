# OrderRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Reason** | **string** |  | 

## Methods

### NewOrderRefundRequest

`func NewOrderRefundRequest(amount int32, reason string, ) *OrderRefundRequest`

NewOrderRefundRequest instantiates a new OrderRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRefundRequestWithDefaults

`func NewOrderRefundRequestWithDefaults() *OrderRefundRequest`

NewOrderRefundRequestWithDefaults instantiates a new OrderRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrderRefundRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderRefundRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderRefundRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *OrderRefundRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OrderRefundRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OrderRefundRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


