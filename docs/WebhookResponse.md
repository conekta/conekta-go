# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the webhook | [optional] 
**Description** | Pointer to **string** | A name or brief explanation of what this webhook is used for | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the webhook is in production | [optional] 
**Active** | Pointer to **bool** | Indicates if the webhook is actived or not | [optional] 
**Object** | Pointer to **string** | Object name, value is &#39;webhook&#39; | [optional] 
**Status** | Pointer to **string** | Indicates if the webhook is ready to receive events or failing | [optional] 
**SubscribedEvents** | Pointer to **[]string** | lists the events that will be sent to the webhook | [optional] 
**Url** | Pointer to **string** | url or endpoint of the webhook | [optional] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse() *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLivemode

`func (o *WebhookResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *WebhookResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *WebhookResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *WebhookResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetActive

`func (o *WebhookResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetObject

`func (o *WebhookResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscribedEvents

`func (o *WebhookResponse) GetSubscribedEvents() []string`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *WebhookResponse) GetSubscribedEventsOk() (*[]string, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *WebhookResponse) SetSubscribedEvents(v []string)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *WebhookResponse) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


