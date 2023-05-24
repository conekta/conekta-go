# OrderResponseDiscountLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**Data** | Pointer to [**[]DiscountLinesDataResponse**](DiscountLinesDataResponse.md) |  | [optional] 

## Methods

### NewOrderResponseDiscountLines

`func NewOrderResponseDiscountLines(hasMore bool, object string, ) *OrderResponseDiscountLines`

NewOrderResponseDiscountLines instantiates a new OrderResponseDiscountLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseDiscountLinesWithDefaults

`func NewOrderResponseDiscountLinesWithDefaults() *OrderResponseDiscountLines`

NewOrderResponseDiscountLinesWithDefaults instantiates a new OrderResponseDiscountLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OrderResponseDiscountLines) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderResponseDiscountLines) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderResponseDiscountLines) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *OrderResponseDiscountLines) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderResponseDiscountLines) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderResponseDiscountLines) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *OrderResponseDiscountLines) GetData() []DiscountLinesDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderResponseDiscountLines) GetDataOk() (*[]DiscountLinesDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderResponseDiscountLines) SetData(v []DiscountLinesDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OrderResponseDiscountLines) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


