# DiscountLinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount to be deducted from the total sum of all payments, in cents. | 
**Code** | **string** | Discount code. | 
**Type** | **string** | It can be &#39;loyalty&#39;, &#39;campaign&#39;, &#39;coupon&#39; o &#39;sign&#39; | 
**Id** | **string** | The discount line id | 
**Object** | **string** | The object name | 
**ParentId** | **string** | The order id | 

## Methods

### NewDiscountLinesResponse

`func NewDiscountLinesResponse(amount int64, code string, type_ string, id string, object string, parentId string, ) *DiscountLinesResponse`

NewDiscountLinesResponse instantiates a new DiscountLinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountLinesResponseWithDefaults

`func NewDiscountLinesResponseWithDefaults() *DiscountLinesResponse`

NewDiscountLinesResponseWithDefaults instantiates a new DiscountLinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DiscountLinesResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DiscountLinesResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DiscountLinesResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCode

`func (o *DiscountLinesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiscountLinesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiscountLinesResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *DiscountLinesResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscountLinesResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscountLinesResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DiscountLinesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountLinesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountLinesResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *DiscountLinesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DiscountLinesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DiscountLinesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetParentId

`func (o *DiscountLinesResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DiscountLinesResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DiscountLinesResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


