# OrderDiscountLinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount to be deducted from the total sum of all payments, in cents. | 
**Code** | **string** | Discount code. | 
**Type** | **string** | It can be &#39;loyalty&#39;, &#39;campaign&#39;, &#39;coupon&#39; o &#39;sign&#39; | 

## Methods

### NewOrderDiscountLinesRequest

`func NewOrderDiscountLinesRequest(amount int64, code string, type_ string, ) *OrderDiscountLinesRequest`

NewOrderDiscountLinesRequest instantiates a new OrderDiscountLinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDiscountLinesRequestWithDefaults

`func NewOrderDiscountLinesRequestWithDefaults() *OrderDiscountLinesRequest`

NewOrderDiscountLinesRequestWithDefaults instantiates a new OrderDiscountLinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrderDiscountLinesRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderDiscountLinesRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderDiscountLinesRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCode

`func (o *OrderDiscountLinesRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderDiscountLinesRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderDiscountLinesRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *OrderDiscountLinesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderDiscountLinesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderDiscountLinesRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


