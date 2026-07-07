# LogsResponseForRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | True, if there are more pages. | [optional] [readonly] 
**Object** | Pointer to **string** | The object type | [optional] [readonly] 
**NextPageUrl** | Pointer to **string** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **string** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]LogsResponseData**](LogsResponseData.md) | set to page results. | [optional] 

## Methods

### NewLogsResponseForRequest

`func NewLogsResponseForRequest() *LogsResponseForRequest`

NewLogsResponseForRequest instantiates a new LogsResponseForRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseForRequestWithDefaults

`func NewLogsResponseForRequestWithDefaults() *LogsResponseForRequest`

NewLogsResponseForRequestWithDefaults instantiates a new LogsResponseForRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *LogsResponseForRequest) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LogsResponseForRequest) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LogsResponseForRequest) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LogsResponseForRequest) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetObject

`func (o *LogsResponseForRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *LogsResponseForRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *LogsResponseForRequest) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *LogsResponseForRequest) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextPageUrl

`func (o *LogsResponseForRequest) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *LogsResponseForRequest) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *LogsResponseForRequest) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *LogsResponseForRequest) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### GetPreviousPageUrl

`func (o *LogsResponseForRequest) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *LogsResponseForRequest) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *LogsResponseForRequest) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *LogsResponseForRequest) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### GetData

`func (o *LogsResponseForRequest) GetData() []LogsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogsResponseForRequest) GetDataOk() (*[]LogsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogsResponseForRequest) SetData(v []LogsResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LogsResponseForRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


