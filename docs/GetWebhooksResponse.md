# GetWebhooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates if there are more pages to be requested | 
**Object** | **string** | Object type, in this case is list | 
**NextPageUrl** | Pointer to **NullableString** | URL of the next page. | [optional] 
**PreviousPageUrl** | Pointer to **NullableString** | Url of the previous page. | [optional] 
**Data** | Pointer to [**[]WebhookResponse**](WebhookResponse.md) |  | [optional] 

## Methods

### NewGetWebhooksResponse

`func NewGetWebhooksResponse(hasMore bool, object string, ) *GetWebhooksResponse`

NewGetWebhooksResponse instantiates a new GetWebhooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhooksResponseWithDefaults

`func NewGetWebhooksResponseWithDefaults() *GetWebhooksResponse`

NewGetWebhooksResponseWithDefaults instantiates a new GetWebhooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetWebhooksResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetWebhooksResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetWebhooksResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetObject

`func (o *GetWebhooksResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetWebhooksResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetWebhooksResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetNextPageUrl

`func (o *GetWebhooksResponse) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *GetWebhooksResponse) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *GetWebhooksResponse) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *GetWebhooksResponse) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *GetWebhooksResponse) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *GetWebhooksResponse) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPreviousPageUrl

`func (o *GetWebhooksResponse) GetPreviousPageUrl() string`

GetPreviousPageUrl returns the PreviousPageUrl field if non-nil, zero value otherwise.

### GetPreviousPageUrlOk

`func (o *GetWebhooksResponse) GetPreviousPageUrlOk() (*string, bool)`

GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUrl

`func (o *GetWebhooksResponse) SetPreviousPageUrl(v string)`

SetPreviousPageUrl sets PreviousPageUrl field to given value.

### HasPreviousPageUrl

`func (o *GetWebhooksResponse) HasPreviousPageUrl() bool`

HasPreviousPageUrl returns a boolean if a field has been set.

### SetPreviousPageUrlNil

`func (o *GetWebhooksResponse) SetPreviousPageUrlNil(b bool)`

 SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil

### UnsetPreviousPageUrl
`func (o *GetWebhooksResponse) UnsetPreviousPageUrl()`

UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
### GetData

`func (o *GetWebhooksResponse) GetData() []WebhookResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetWebhooksResponse) GetDataOk() (*[]WebhookResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetWebhooksResponse) SetData(v []WebhookResponse)`

SetData sets Data field to given value.

### HasData

`func (o *GetWebhooksResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


