# UpdateOrderDiscountLinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** | Discount code. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateOrderDiscountLinesRequest

`func NewUpdateOrderDiscountLinesRequest() *UpdateOrderDiscountLinesRequest`

NewUpdateOrderDiscountLinesRequest instantiates a new UpdateOrderDiscountLinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderDiscountLinesRequestWithDefaults

`func NewUpdateOrderDiscountLinesRequestWithDefaults() *UpdateOrderDiscountLinesRequest`

NewUpdateOrderDiscountLinesRequestWithDefaults instantiates a new UpdateOrderDiscountLinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdateOrderDiscountLinesRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateOrderDiscountLinesRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateOrderDiscountLinesRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateOrderDiscountLinesRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCode

`func (o *UpdateOrderDiscountLinesRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateOrderDiscountLinesRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateOrderDiscountLinesRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateOrderDiscountLinesRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *UpdateOrderDiscountLinesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOrderDiscountLinesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOrderDiscountLinesRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOrderDiscountLinesRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


