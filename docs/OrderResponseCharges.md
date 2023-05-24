# OrderResponseCharges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**Data** | Pointer to [**[]ChargesDataResponse**](ChargesDataResponse.md) |  | [optional] 

## Methods

### NewOrderResponseCharges

`func NewOrderResponseCharges(hasMore bool, object string, ) *OrderResponseCharges`

NewOrderResponseCharges instantiates a new OrderResponseCharges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseChargesWithDefaults

`func NewOrderResponseChargesWithDefaults() *OrderResponseCharges`

NewOrderResponseChargesWithDefaults instantiates a new OrderResponseCharges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OrderResponseCharges) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderResponseCharges) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderResponseCharges) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *OrderResponseCharges) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderResponseCharges) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderResponseCharges) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *OrderResponseCharges) GetData() []ChargesDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderResponseCharges) GetDataOk() (*[]ChargesDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderResponseCharges) SetData(v []ChargesDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OrderResponseCharges) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


