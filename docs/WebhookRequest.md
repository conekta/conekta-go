# WebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Here you must place the URL of your Webhook remember that you must program what you will do with the events received. Also do not forget to handle the HTTPS protocol for greater security. | 
**Synchronous** | **bool** | It is a value that allows to decide if the events will be synchronous or asynchronous. We recommend asynchronous &#x3D; false | [default to false]

## Methods

### NewWebhookRequest

`func NewWebhookRequest(url string, synchronous bool, ) *WebhookRequest`

NewWebhookRequest instantiates a new WebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestWithDefaults

`func NewWebhookRequestWithDefaults() *WebhookRequest`

NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSynchronous

`func (o *WebhookRequest) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *WebhookRequest) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *WebhookRequest) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


