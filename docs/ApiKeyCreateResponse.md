# ApiKeyCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationToken** | Pointer to **string** | It is occupied as a user when authenticated with basic authentication, with a blank password. This value will only appear once, in the request to create a new key | [optional] 
**Active** | Pointer to **bool** | Indicates if the api key is active | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in seconds with the creation date of the api key | [optional] 
**Description** | Pointer to **string** | Detail of the use that will be given to the api key | [optional] 
**Id** | Pointer to **string** | Unique identifier of the api key | [optional] 
**Livemode** | Pointer to **bool** | Indicates if the api key is in live mode | [optional] 
**Object** | Pointer to **string** | Object name, value is api_key | [optional] 
**Prefix** | Pointer to **string** | The first few characters of the authentication_token | [optional] 
**Role** | Pointer to **string** | Indicates the user account private&#x3D;owner or public&#x3D;public | [optional] 

## Methods

### NewApiKeyCreateResponse

`func NewApiKeyCreateResponse() *ApiKeyCreateResponse`

NewApiKeyCreateResponse instantiates a new ApiKeyCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreateResponseWithDefaults

`func NewApiKeyCreateResponseWithDefaults() *ApiKeyCreateResponse`

NewApiKeyCreateResponseWithDefaults instantiates a new ApiKeyCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationToken

`func (o *ApiKeyCreateResponse) GetAuthenticationToken() string`

GetAuthenticationToken returns the AuthenticationToken field if non-nil, zero value otherwise.

### GetAuthenticationTokenOk

`func (o *ApiKeyCreateResponse) GetAuthenticationTokenOk() (*string, bool)`

GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationToken

`func (o *ApiKeyCreateResponse) SetAuthenticationToken(v string)`

SetAuthenticationToken sets AuthenticationToken field to given value.

### HasAuthenticationToken

`func (o *ApiKeyCreateResponse) HasAuthenticationToken() bool`

HasAuthenticationToken returns a boolean if a field has been set.

### GetActive

`func (o *ApiKeyCreateResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyCreateResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyCreateResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyCreateResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKeyCreateResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyCreateResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyCreateResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKeyCreateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiKeyCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKeyCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *ApiKeyCreateResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *ApiKeyCreateResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *ApiKeyCreateResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *ApiKeyCreateResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *ApiKeyCreateResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ApiKeyCreateResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ApiKeyCreateResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ApiKeyCreateResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPrefix

`func (o *ApiKeyCreateResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKeyCreateResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKeyCreateResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiKeyCreateResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRole

`func (o *ApiKeyCreateResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiKeyCreateResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiKeyCreateResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiKeyCreateResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


