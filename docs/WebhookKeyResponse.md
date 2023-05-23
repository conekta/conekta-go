# WebhookKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the webhook key | [optional] 
**Active** | Pointer to **bool** | Indicates if the webhook key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds with the creation date of the webhook key | [optional] 
**DeactivatedAt** | Pointer to **NullableInt64** | Unix timestamp in seconds with the deactivation date of the webhook key | [optional] 
**PublicKey** | Pointer to **string** | Public key to be used in the webhook | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the webhook key is in live mode | [optional] 
**Object** | Pointer to **string** | Object name, value is webhook_key | [optional] 

## Methods

### NewWebhookKeyResponse

`func NewWebhookKeyResponse() *WebhookKeyResponse`

NewWebhookKeyResponse instantiates a new WebhookKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookKeyResponseWithDefaults

`func NewWebhookKeyResponseWithDefaults() *WebhookKeyResponse`

NewWebhookKeyResponseWithDefaults instantiates a new WebhookKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookKeyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *WebhookKeyResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookKeyResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookKeyResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookKeyResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookKeyResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookKeyResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookKeyResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookKeyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *WebhookKeyResponse) GetDeactivatedAt() int64`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *WebhookKeyResponse) GetDeactivatedAtOk() (*int64, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *WebhookKeyResponse) SetDeactivatedAt(v int64)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *WebhookKeyResponse) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### SetDeactivatedAtNil

`func (o *WebhookKeyResponse) SetDeactivatedAtNil(b bool)`

 SetDeactivatedAtNil sets the value for DeactivatedAt to be an explicit nil

### UnsetDeactivatedAt
`func (o *WebhookKeyResponse) UnsetDeactivatedAt()`

UnsetDeactivatedAt ensures that no value is present for DeactivatedAt, not even an explicit nil
### GetPublicKey

`func (o *WebhookKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WebhookKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WebhookKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WebhookKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLivemode

`func (o *WebhookKeyResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *WebhookKeyResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *WebhookKeyResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *WebhookKeyResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *WebhookKeyResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookKeyResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookKeyResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookKeyResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


