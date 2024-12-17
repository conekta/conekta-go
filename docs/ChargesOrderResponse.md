# ChargesOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**Data** | Pointer to [**[]ChargesOrderResponseAllOfData**](ChargesOrderResponseAllOfData.md) |  | [optional] 

## Methods

### NewChargesOrderResponse

`func NewChargesOrderResponse(hasMore bool, object string, ) *ChargesOrderResponse`

NewChargesOrderResponse instantiates a new ChargesOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargesOrderResponseWithDefaults

`func NewChargesOrderResponseWithDefaults() *ChargesOrderResponse`

NewChargesOrderResponseWithDefaults instantiates a new ChargesOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ChargesOrderResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChargesOrderResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChargesOrderResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *ChargesOrderResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargesOrderResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargesOrderResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *ChargesOrderResponse) GetData() []ChargesOrderResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChargesOrderResponse) GetDataOk() (*[]ChargesOrderResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChargesOrderResponse) SetData(v []ChargesOrderResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *ChargesOrderResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


