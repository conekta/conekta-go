# WebhookKeyCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the webhook key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds with the creation date of the webhook key | [optional] 
**Id** | Pointer to **string** | Unique identifier of the webhook key | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the webhook key is in live mode | [optional] 
**Object** | Pointer to **string** | Object name, value is webhook_key | [optional] 
**PublicKey** | Pointer to **string** | Public key to be used in the webhook | [optional] 

## Methods

### NewWebhookKeyCreateResponse

`func NewWebhookKeyCreateResponse() *WebhookKeyCreateResponse`

NewWebhookKeyCreateResponse instantiates a new WebhookKeyCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookKeyCreateResponseWithDefaults

`func NewWebhookKeyCreateResponseWithDefaults() *WebhookKeyCreateResponse`

NewWebhookKeyCreateResponseWithDefaults instantiates a new WebhookKeyCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookKeyCreateResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookKeyCreateResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookKeyCreateResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookKeyCreateResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookKeyCreateResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookKeyCreateResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookKeyCreateResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookKeyCreateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WebhookKeyCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookKeyCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookKeyCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookKeyCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *WebhookKeyCreateResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *WebhookKeyCreateResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *WebhookKeyCreateResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *WebhookKeyCreateResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *WebhookKeyCreateResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookKeyCreateResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookKeyCreateResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookKeyCreateResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPublicKey

`func (o *WebhookKeyCreateResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WebhookKeyCreateResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WebhookKeyCreateResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WebhookKeyCreateResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


