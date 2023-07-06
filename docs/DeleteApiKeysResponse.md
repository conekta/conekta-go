# DeleteApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the api key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds of when the api key was created | [optional] 
**Description** | Pointer to **string** | A name or brief explanation of what this api key is used for | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the api key is in production | [optional] 
**Prefix** | Pointer to **string** | The first few characters of the authentication_token | [optional] 
**Id** | Pointer to **string** | Unique identifier of the api key | [optional] 
**Object** | Pointer to **string** | Object name, value is &#39;api_key&#39; | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** | Indicates if the api key is private or public | [optional] 

## Methods

### NewDeleteApiKeysResponse

`func NewDeleteApiKeysResponse() *DeleteApiKeysResponse`

NewDeleteApiKeysResponse instantiates a new DeleteApiKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApiKeysResponseWithDefaults

`func NewDeleteApiKeysResponseWithDefaults() *DeleteApiKeysResponse`

NewDeleteApiKeysResponseWithDefaults instantiates a new DeleteApiKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DeleteApiKeysResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DeleteApiKeysResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DeleteApiKeysResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DeleteApiKeysResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeleteApiKeysResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeleteApiKeysResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeleteApiKeysResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeleteApiKeysResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DeleteApiKeysResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeleteApiKeysResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeleteApiKeysResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeleteApiKeysResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLivemode

`func (o *DeleteApiKeysResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *DeleteApiKeysResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *DeleteApiKeysResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *DeleteApiKeysResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetPrefix

`func (o *DeleteApiKeysResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DeleteApiKeysResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DeleteApiKeysResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DeleteApiKeysResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetId

`func (o *DeleteApiKeysResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteApiKeysResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteApiKeysResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteApiKeysResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *DeleteApiKeysResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DeleteApiKeysResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DeleteApiKeysResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DeleteApiKeysResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDeleted

`func (o *DeleteApiKeysResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteApiKeysResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteApiKeysResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteApiKeysResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRole

`func (o *DeleteApiKeysResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DeleteApiKeysResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DeleteApiKeysResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DeleteApiKeysResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


