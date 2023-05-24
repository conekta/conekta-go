# ApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the api key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds with the creation date of the api key | [optional] 
**Description** | Pointer to **string** | Detail of the use that will be given to the api key | [optional] 
**Id** | Pointer to **string** | Unique identifier of the api key | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the api key is in live mode | [optional] 
**Object** | Pointer to **string** | Object name, value is api_key | [optional] 
**Prefix** | Pointer to **string** | The first few characters of the authentication_token | [optional] 
**Role** | Pointer to **string** | Indicates the user account private&#x3D;owner or public&#x3D;public | [optional] 

## Methods

### NewApiKeyResponse

`func NewApiKeyResponse() *ApiKeyResponse`

NewApiKeyResponse instantiates a new ApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyResponseWithDefaults

`func NewApiKeyResponseWithDefaults() *ApiKeyResponse`

NewApiKeyResponseWithDefaults instantiates a new ApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKeyResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKeyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ApiKeyResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ApiKeyResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ApiKeyResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ApiKeyResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *ApiKeyResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ApiKeyResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ApiKeyResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ApiKeyResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPrefix

`func (o *ApiKeyResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKeyResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKeyResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiKeyResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRole

`func (o *ApiKeyResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiKeyResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiKeyResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiKeyResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


