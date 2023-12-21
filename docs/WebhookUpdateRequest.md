# WebhookUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Here you must place the URL of your Webhook remember that you must program what you will do with the events received. Also do not forget to handle the HTTPS protocol for greater security. | 
**Synchronous** | Pointer to **bool** | It is a value that allows to decide if the events will be synchronous or asynchronous. We recommend asynchronous &#x3D; false | [optional] [default to false]
**Events** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebhookUpdateRequest

`func NewWebhookUpdateRequest(url string, ) *WebhookUpdateRequest`

NewWebhookUpdateRequest instantiates a new WebhookUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateRequestWithDefaults

`func NewWebhookUpdateRequestWithDefaults() *WebhookUpdateRequest`

NewWebhookUpdateRequestWithDefaults instantiates a new WebhookUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSynchronous

`func (o *WebhookUpdateRequest) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *WebhookUpdateRequest) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *WebhookUpdateRequest) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *WebhookUpdateRequest) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookUpdateRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookUpdateRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookUpdateRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookUpdateRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


