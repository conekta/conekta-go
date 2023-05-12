# WebhookKeyDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the webhook key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds with the creation date of the webhook key | [optional] 
**Deleted** | Pointer to **bool** | Indicates if the webhook key is deleted | [optional] 
**Id** | Pointer to **string** | Unique identifier of the webhook key | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the webhook key is in live mode | [optional] 
**Object** | Pointer to **string** | Object name, value is webhook_key | [optional] 

## Methods

### NewWebhookKeyDeleteResponse

`func NewWebhookKeyDeleteResponse() *WebhookKeyDeleteResponse`

NewWebhookKeyDeleteResponse instantiates a new WebhookKeyDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookKeyDeleteResponseWithDefaults

`func NewWebhookKeyDeleteResponseWithDefaults() *WebhookKeyDeleteResponse`

NewWebhookKeyDeleteResponseWithDefaults instantiates a new WebhookKeyDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookKeyDeleteResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookKeyDeleteResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookKeyDeleteResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookKeyDeleteResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookKeyDeleteResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookKeyDeleteResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookKeyDeleteResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookKeyDeleteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleted

`func (o *WebhookKeyDeleteResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *WebhookKeyDeleteResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *WebhookKeyDeleteResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *WebhookKeyDeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *WebhookKeyDeleteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookKeyDeleteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookKeyDeleteResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookKeyDeleteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *WebhookKeyDeleteResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *WebhookKeyDeleteResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *WebhookKeyDeleteResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *WebhookKeyDeleteResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *WebhookKeyDeleteResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookKeyDeleteResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookKeyDeleteResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookKeyDeleteResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


