# GetChargesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**NextPageUrl** | Pointer to **NullableString** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **NullableString** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]ChargeResponse**](ChargeResponse.md) |  | [optional] 

## Methods

### NewGetChargesResponse

`func NewGetChargesResponse(hasMore bool, object string, ) *GetChargesResponse`

NewGetChargesResponse instantiates a new GetChargesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChargesResponseWithDefaults

`func NewGetChargesResponseWithDefaults() *GetChargesResponse`

NewGetChargesResponseWithDefaults instantiates a new GetChargesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetChargesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetChargesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetChargesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *GetChargesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetChargesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetChargesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetNextPageUrl

`func (o *GetChargesResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *GetChargesResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *GetChargesResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *GetChargesResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *GetChargesResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *GetChargesResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPreviousPageUrl

`func (o *GetChargesResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *GetChargesResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *GetChargesResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *GetChargesResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### SetPreviousPageUrlNil

`func (o *GetChargesResponse) SetPreviousPageUrlNil(b bool)`

 SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil

### UnsetPreviousPageUrl
`func (o *GetChargesResponse) UnsetPreviousPageUrl()`

UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
### GetData

`func (o *GetChargesResponse) GetData() []ChargeResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChargesResponse) GetDataOk() (*[]ChargeResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChargesResponse) SetData(v []ChargeResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetChargesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


