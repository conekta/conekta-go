# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**DevelopmentEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ProductionEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscribedEvents** | Pointer to **[]string** |  | [optional] 
**Synchronous** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

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

### GetDeleted

`func (o *WebhookResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *WebhookResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *WebhookResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *WebhookResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *WebhookResponse) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *WebhookResponse) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetDevelopmentEnabled

`func (o *WebhookResponse) GetDevelopmentEnabled() bool`

GetDevelopmentEnabled returns the DevelopmentEnabled field if non-nil, zero value otherwise.

### GetDevelopmentEnabledOk

`func (o *WebhookResponse) GetDevelopmentEnabledOk() (*bool, bool)`

GetDevelopmentEnabledOk returns a tuple with the DevelopmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentEnabled

`func (o *WebhookResponse) SetDevelopmentEnabled(v bool)`

SetDevelopmentEnabled sets DevelopmentEnabled field to given value.

### HasDevelopmentEnabled

`func (o *WebhookResponse) HasDevelopmentEnabled() bool`

HasDevelopmentEnabled returns a boolean if a field has been set.

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

### GetProductionEnabled

`func (o *WebhookResponse) GetProductionEnabled() bool`

GetProductionEnabled returns the ProductionEnabled field if non-nil, zero value otherwise.

### GetProductionEnabledOk

`func (o *WebhookResponse) GetProductionEnabledOk() (*bool, bool)`

GetProductionEnabledOk returns a tuple with the ProductionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionEnabled

`func (o *WebhookResponse) SetProductionEnabled(v bool)`

SetProductionEnabled sets ProductionEnabled field to given value.

### HasProductionEnabled

`func (o *WebhookResponse) HasProductionEnabled() bool`

HasProductionEnabled returns a boolean if a field has been set.

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

### GetSynchronous

`func (o *WebhookResponse) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *WebhookResponse) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *WebhookResponse) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *WebhookResponse) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

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


