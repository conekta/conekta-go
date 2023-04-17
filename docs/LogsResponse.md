# LogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | True, if there are more pages. | [optional] [readonly] 
**Object** | Pointer to **string** | The object type | [optional] [readonly] 
**NextPageUrl** | Pointer to **NullableString** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **NullableString** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]LogsResponseData**](LogsResponseData.md) | set to page results. | [optional] 

## Methods

### NewLogsResponse

`func NewLogsResponse() *LogsResponse`

NewLogsResponse instantiates a new LogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseWithDefaults

`func NewLogsResponseWithDefaults() *LogsResponse`

NewLogsResponseWithDefaults instantiates a new LogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *LogsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LogsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LogsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LogsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetObject

`func (o *LogsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *LogsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *LogsResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *LogsResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextPageUrl

`func (o *LogsResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *LogsResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *LogsResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *LogsResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *LogsResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *LogsResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPreviousPageUrl

`func (o *LogsResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *LogsResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *LogsResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *LogsResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### SetPreviousPageUrlNil

`func (o *LogsResponse) SetPreviousPageUrlNil(b bool)`

 SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil

### UnsetPreviousPageUrl
`func (o *LogsResponse) UnsetPreviousPageUrl()`

UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
### GetData

`func (o *LogsResponse) GetData() []LogsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogsResponse) GetDataOk() (*[]LogsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogsResponse) SetData(v []LogsResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LogsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *LogsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *LogsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


