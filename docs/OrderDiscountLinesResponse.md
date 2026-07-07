# OrderDiscountLinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**NextPageUrl** | Pointer to **string** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **string** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]DiscountLinesDataResponse**](DiscountLinesDataResponse.md) |  | [optional] 

## Methods

### NewOrderDiscountLinesResponse

`func NewOrderDiscountLinesResponse(hasMore bool, object string, ) *OrderDiscountLinesResponse`

NewOrderDiscountLinesResponse instantiates a new OrderDiscountLinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDiscountLinesResponseWithDefaults

`func NewOrderDiscountLinesResponseWithDefaults() *OrderDiscountLinesResponse`

NewOrderDiscountLinesResponseWithDefaults instantiates a new OrderDiscountLinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OrderDiscountLinesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderDiscountLinesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderDiscountLinesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *OrderDiscountLinesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderDiscountLinesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderDiscountLinesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetNextPageUrl

`func (o *OrderDiscountLinesResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *OrderDiscountLinesResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *OrderDiscountLinesResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *OrderDiscountLinesResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPreviousPageUrl

`func (o *OrderDiscountLinesResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *OrderDiscountLinesResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *OrderDiscountLinesResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *OrderDiscountLinesResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### GetData

`func (o *OrderDiscountLinesResponse) GetData() []DiscountLinesDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderDiscountLinesResponse) GetDataOk() (*[]DiscountLinesDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderDiscountLinesResponse) SetData(v []DiscountLinesDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OrderDiscountLinesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


