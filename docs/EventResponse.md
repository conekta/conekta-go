# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WebhookLogs** | Pointer to [**[]WebhookLog**](WebhookLog.md) |  | [optional] 
**WebhookStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewEventResponse

`func NewEventResponse() *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EventResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EventResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *EventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *EventResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *EventResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *EventResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *EventResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *EventResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EventResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EventResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *EventResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *EventResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebhookLogs

`func (o *EventResponse) GetWebhookLogs() []WebhookLog`

GetWebhookLogs returns the WebhookLogs field if non-nil, zero value otherwise.

### GetWebhookLogsOk

`func (o *EventResponse) GetWebhookLogsOk() (*[]WebhookLog, bool)`

GetWebhookLogsOk returns a tuple with the WebhookLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookLogs

`func (o *EventResponse) SetWebhookLogs(v []WebhookLog)`

SetWebhookLogs sets WebhookLogs field to given value.

### HasWebhookLogs

`func (o *EventResponse) HasWebhookLogs() bool`

HasWebhookLogs returns a boolean if a field has been set.

### GetWebhookStatus

`func (o *EventResponse) GetWebhookStatus() string`

GetWebhookStatus returns the WebhookStatus field if non-nil, zero value otherwise.

### GetWebhookStatusOk

`func (o *EventResponse) GetWebhookStatusOk() (*string, bool)`

GetWebhookStatusOk returns a tuple with the WebhookStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookStatus

`func (o *EventResponse) SetWebhookStatus(v string)`

SetWebhookStatus sets WebhookStatus field to given value.

### HasWebhookStatus

`func (o *EventResponse) HasWebhookStatus() bool`

HasWebhookStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


