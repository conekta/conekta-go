# OrderCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Amount to capture | 

## Methods

### NewOrderCaptureRequest

`func NewOrderCaptureRequest(amount int64, ) *OrderCaptureRequest`

NewOrderCaptureRequest instantiates a new OrderCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCaptureRequestWithDefaults

`func NewOrderCaptureRequestWithDefaults() *OrderCaptureRequest`

NewOrderCaptureRequestWithDefaults instantiates a new OrderCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrderCaptureRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderCaptureRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderCaptureRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


