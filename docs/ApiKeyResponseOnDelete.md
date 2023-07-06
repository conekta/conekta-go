# ApiKeyResponseOnDelete

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
**Deleted** | Pointer to **bool** | Indicates if the api key was deleted | [optional] 
**Role** | Pointer to **string** | Indicates if the api key is private or public | [optional] 

## Methods

### NewApiKeyResponseOnDelete

`func NewApiKeyResponseOnDelete() *ApiKeyResponseOnDelete`

NewApiKeyResponseOnDelete instantiates a new ApiKeyResponseOnDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyResponseOnDeleteWithDefaults

`func NewApiKeyResponseOnDeleteWithDefaults() *ApiKeyResponseOnDelete`

NewApiKeyResponseOnDeleteWithDefaults instantiates a new ApiKeyResponseOnDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyResponseOnDelete) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyResponseOnDelete) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyResponseOnDelete) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyResponseOnDelete) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKeyResponseOnDelete) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyResponseOnDelete) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyResponseOnDelete) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKeyResponseOnDelete) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyResponseOnDelete) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyResponseOnDelete) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyResponseOnDelete) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyResponseOnDelete) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLivemode

`func (o *ApiKeyResponseOnDelete) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ApiKeyResponseOnDelete) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ApiKeyResponseOnDelete) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ApiKeyResponseOnDelete) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetPrefix

`func (o *ApiKeyResponseOnDelete) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKeyResponseOnDelete) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKeyResponseOnDelete) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiKeyResponseOnDelete) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetId

`func (o *ApiKeyResponseOnDelete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyResponseOnDelete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyResponseOnDelete) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyResponseOnDelete) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ApiKeyResponseOnDelete) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ApiKeyResponseOnDelete) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ApiKeyResponseOnDelete) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ApiKeyResponseOnDelete) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDeleted

`func (o *ApiKeyResponseOnDelete) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ApiKeyResponseOnDelete) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ApiKeyResponseOnDelete) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ApiKeyResponseOnDelete) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRole

`func (o *ApiKeyResponseOnDelete) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiKeyResponseOnDelete) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiKeyResponseOnDelete) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiKeyResponseOnDelete) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


