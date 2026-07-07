# OrderTaxLinesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**NextPageUrl** | Pointer to **string** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **string** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]TaxLinesDataResponse**](TaxLinesDataResponse.md) |  | [optional] 

## Methods

### NewOrderTaxLinesResponse

`func NewOrderTaxLinesResponse(hasMore bool, object string, ) *OrderTaxLinesResponse`

NewOrderTaxLinesResponse instantiates a new OrderTaxLinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaxLinesResponseWithDefaults

`func NewOrderTaxLinesResponseWithDefaults() *OrderTaxLinesResponse`

NewOrderTaxLinesResponseWithDefaults instantiates a new OrderTaxLinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *OrderTaxLinesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OrderTaxLinesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OrderTaxLinesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *OrderTaxLinesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderTaxLinesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderTaxLinesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetNextPageUrl

`func (o *OrderTaxLinesResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *OrderTaxLinesResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *OrderTaxLinesResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *OrderTaxLinesResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPreviousPageUrl

`func (o *OrderTaxLinesResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *OrderTaxLinesResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *OrderTaxLinesResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *OrderTaxLinesResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### GetData

`func (o *OrderTaxLinesResponse) GetData() []TaxLinesDataResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderTaxLinesResponse) GetDataOk() (*[]TaxLinesDataResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderTaxLinesResponse) SetData(v []TaxLinesDataResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OrderTaxLinesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


