# UpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Here you must place the URL of your Webhook remember that you must program what you will do with the events received. Also do not forget to handle the HTTPS protocol for greater security. | [optional] 
**SubscribedEvents** | Pointer to **[]string** | events that will be sent to the webhook | [optional] 
**Active** | Pointer to **bool** | whether the webhook is active or not | [optional] 

## Methods

### NewUpdateWebhook

`func NewUpdateWebhook() *UpdateWebhook`

NewUpdateWebhook instantiates a new UpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookWithDefaults

`func NewUpdateWebhookWithDefaults() *UpdateWebhook`

NewUpdateWebhookWithDefaults instantiates a new UpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSubscribedEvents

`func (o *UpdateWebhook) GetSubscribedEvents() []string`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *UpdateWebhook) GetSubscribedEventsOk() (*[]string, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *UpdateWebhook) SetSubscribedEvents(v []string)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *UpdateWebhook) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetActive

`func (o *UpdateWebhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateWebhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateWebhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateWebhook) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


