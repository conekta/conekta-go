# WebhookKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the webhook key is active | [optional] [default to false]

## Methods

### NewWebhookKeyUpdateRequest

`func NewWebhookKeyUpdateRequest() *WebhookKeyUpdateRequest`

NewWebhookKeyUpdateRequest instantiates a new WebhookKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookKeyUpdateRequestWithDefaults

`func NewWebhookKeyUpdateRequestWithDefaults() *WebhookKeyUpdateRequest`

NewWebhookKeyUpdateRequestWithDefaults instantiates a new WebhookKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookKeyUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookKeyUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookKeyUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookKeyUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


