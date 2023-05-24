# EventsResendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAttempts** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastAttemptedAt** | Pointer to **int32** |  | [optional] 
**LastHttpResponseStatus** | Pointer to **int32** |  | [optional] 
**ResponseData** | Pointer to **map[string]interface{}** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsResendResponse

`func NewEventsResendResponse() *EventsResendResponse`

NewEventsResendResponse instantiates a new EventsResendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsResendResponseWithDefaults

`func NewEventsResendResponseWithDefaults() *EventsResendResponse`

NewEventsResendResponseWithDefaults instantiates a new EventsResendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAttempts

`func (o *EventsResendResponse) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *EventsResendResponse) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *EventsResendResponse) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *EventsResendResponse) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### GetId

`func (o *EventsResendResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsResendResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsResendResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventsResendResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAttemptedAt

`func (o *EventsResendResponse) GetLastAttemptedAt() int32`

GetLastAttemptedAt returns the LastAttemptedAt field if non-nil, zero value otherwise.

### GetLastAttemptedAtOk

`func (o *EventsResendResponse) GetLastAttemptedAtOk() (*int32, bool)`

GetLastAttemptedAtOk returns a tuple with the LastAttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptedAt

`func (o *EventsResendResponse) SetLastAttemptedAt(v int32)`

SetLastAttemptedAt sets LastAttemptedAt field to given value.

### HasLastAttemptedAt

`func (o *EventsResendResponse) HasLastAttemptedAt() bool`

HasLastAttemptedAt returns a boolean if a field has been set.

### GetLastHttpResponseStatus

`func (o *EventsResendResponse) GetLastHttpResponseStatus() int32`

GetLastHttpResponseStatus returns the LastHttpResponseStatus field if non-nil, zero value otherwise.

### GetLastHttpResponseStatusOk

`func (o *EventsResendResponse) GetLastHttpResponseStatusOk() (*int32, bool)`

GetLastHttpResponseStatusOk returns a tuple with the LastHttpResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHttpResponseStatus

`func (o *EventsResendResponse) SetLastHttpResponseStatus(v int32)`

SetLastHttpResponseStatus sets LastHttpResponseStatus field to given value.

### HasLastHttpResponseStatus

`func (o *EventsResendResponse) HasLastHttpResponseStatus() bool`

HasLastHttpResponseStatus returns a boolean if a field has been set.

### GetResponseData

`func (o *EventsResendResponse) GetResponseData() map[string]interface{}`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *EventsResendResponse) GetResponseDataOk() (*map[string]interface{}, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *EventsResendResponse) SetResponseData(v map[string]interface{})`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *EventsResendResponse) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### GetUrl

`func (o *EventsResendResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsResendResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsResendResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsResendResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


