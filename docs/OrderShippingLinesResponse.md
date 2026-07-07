# OrderShippingLinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**NextPageUrl** | Pointer to **string** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **string** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]ShippingLinesDataResponse**](ShippingLinesDataResponse.md) |  | [optional] 

## Methods

### NewOrderShippingLinesResponse

`func NewOrderShippingLinesResponse(hasMore bool, object string, ) *OrderShippingLinesResponse`

NewOrderShippingLinesResponse instantiates a new OrderShippingLinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShippingLinesResponseWithDefaults

`func NewOrderShippingLinesResponseWithDefaults() *OrderShippingLinesResponse`

NewOrderShippingLinesResponseWithDefaults instantiates a new OrderShippingLinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OrderShippingLinesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderShippingLinesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderShippingLinesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *OrderShippingLinesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderShippingLinesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderShippingLinesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetNextPageUrl

`func (o *OrderShippingLinesResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *OrderShippingLinesResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *OrderShippingLinesResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *OrderShippingLinesResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPreviousPageUrl

`func (o *OrderShippingLinesResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *OrderShippingLinesResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *OrderShippingLinesResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *OrderShippingLinesResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### GetData

`func (o *OrderShippingLinesResponse) GetData() []ShippingLinesDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderShippingLinesResponse) GetDataOk() (*[]ShippingLinesDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderShippingLinesResponse) SetData(v []ShippingLinesDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OrderShippingLinesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


