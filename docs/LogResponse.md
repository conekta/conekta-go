# LogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**Id** | **string** |  | 
**IpAddress** | Pointer to **string** |  | [optional] 
**Livemode** | **bool** |  | 
**LoggableId** | Pointer to **NullableString** |  | [optional] 
**LoggableType** | Pointer to **NullableString** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**OauthTokenId** | Pointer to **NullableString** |  | [optional] 
**QueryString** | Pointer to **map[string]interface{}** |  | [optional] 
**Related** | Pointer to **string** |  | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestHeaders** | Pointer to **map[string]string** |  | [optional] 
**ResponseBody** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseHeaders** | Pointer to **map[string]string** |  | [optional] 
**SearchableTags** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserAccountId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewLogResponse

`func NewLogResponse(createdAt int64, id string, livemode bool, ) *LogResponse`

NewLogResponse instantiates a new LogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResponseWithDefaults

`func NewLogResponseWithDefaults() *LogResponse`

NewLogResponseWithDefaults instantiates a new LogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LogResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *LogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *LogResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLivemode

`func (o *LogResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *LogResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *LogResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetLoggableId

`func (o *LogResponse) GetLoggableId() string`

GetLoggableId returns the LoggableId field if non-nil, zero value otherwise.

### GetLoggableIdOk

`func (o *LogResponse) GetLoggableIdOk() (*string, bool)`

GetLoggableIdOk returns a tuple with the LoggableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableId

`func (o *LogResponse) SetLoggableId(v string)`

SetLoggableId sets LoggableId field to given value.

### HasLoggableId

`func (o *LogResponse) HasLoggableId() bool`

HasLoggableId returns a boolean if a field has been set.

### SetLoggableIdNil

`func (o *LogResponse) SetLoggableIdNil(b bool)`

 SetLoggableIdNil sets the value for LoggableId to be an explicit nil

### UnsetLoggableId
`func (o *LogResponse) UnsetLoggableId()`

UnsetLoggableId ensures that no value is present for LoggableId, not even an explicit nil
### GetLoggableType

`func (o *LogResponse) GetLoggableType() string`

GetLoggableType returns the LoggableType field if non-nil, zero value otherwise.

### GetLoggableTypeOk

`func (o *LogResponse) GetLoggableTypeOk() (*string, bool)`

GetLoggableTypeOk returns a tuple with the LoggableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableType

`func (o *LogResponse) SetLoggableType(v string)`

SetLoggableType sets LoggableType field to given value.

### HasLoggableType

`func (o *LogResponse) HasLoggableType() bool`

HasLoggableType returns a boolean if a field has been set.

### SetLoggableTypeNil

`func (o *LogResponse) SetLoggableTypeNil(b bool)`

 SetLoggableTypeNil sets the value for LoggableType to be an explicit nil

### UnsetLoggableType
`func (o *LogResponse) UnsetLoggableType()`

UnsetLoggableType ensures that no value is present for LoggableType, not even an explicit nil
### GetMethod

`func (o *LogResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOauthTokenId

`func (o *LogResponse) GetOauthTokenId() string`

GetOauthTokenId returns the OauthTokenId field if non-nil, zero value otherwise.

### GetOauthTokenIdOk

`func (o *LogResponse) GetOauthTokenIdOk() (*string, bool)`

GetOauthTokenIdOk returns a tuple with the OauthTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenId

`func (o *LogResponse) SetOauthTokenId(v string)`

SetOauthTokenId sets OauthTokenId field to given value.

### HasOauthTokenId

`func (o *LogResponse) HasOauthTokenId() bool`

HasOauthTokenId returns a boolean if a field has been set.

### SetOauthTokenIdNil

`func (o *LogResponse) SetOauthTokenIdNil(b bool)`

 SetOauthTokenIdNil sets the value for OauthTokenId to be an explicit nil

### UnsetOauthTokenId
`func (o *LogResponse) UnsetOauthTokenId()`

UnsetOauthTokenId ensures that no value is present for OauthTokenId, not even an explicit nil
### GetQueryString

`func (o *LogResponse) GetQueryString() map[string]interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogResponse) GetQueryStringOk() (*map[string]interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogResponse) SetQueryString(v map[string]interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LogResponse) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetRelated

`func (o *LogResponse) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *LogResponse) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *LogResponse) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *LogResponse) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRequestBody

`func (o *LogResponse) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *LogResponse) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *LogResponse) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *LogResponse) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *LogResponse) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *LogResponse) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *LogResponse) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *LogResponse) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseBody

`func (o *LogResponse) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *LogResponse) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *LogResponse) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *LogResponse) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *LogResponse) GetResponseHeaders() map[string]string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *LogResponse) GetResponseHeadersOk() (*map[string]string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *LogResponse) SetResponseHeaders(v map[string]string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *LogResponse) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetSearchableTags

`func (o *LogResponse) GetSearchableTags() []string`

GetSearchableTags returns the SearchableTags field if non-nil, zero value otherwise.

### GetSearchableTagsOk

`func (o *LogResponse) GetSearchableTagsOk() (*[]string, bool)`

GetSearchableTagsOk returns a tuple with the SearchableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableTags

`func (o *LogResponse) SetSearchableTags(v []string)`

SetSearchableTags sets SearchableTags field to given value.

### HasSearchableTags

`func (o *LogResponse) HasSearchableTags() bool`

HasSearchableTags returns a boolean if a field has been set.

### GetStatus

`func (o *LogResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LogResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LogResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *LogResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LogResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserAccountId

`func (o *LogResponse) GetUserAccountId() string`

GetUserAccountId returns the UserAccountId field if non-nil, zero value otherwise.

### GetUserAccountIdOk

`func (o *LogResponse) GetUserAccountIdOk() (*string, bool)`

GetUserAccountIdOk returns a tuple with the UserAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountId

`func (o *LogResponse) SetUserAccountId(v string)`

SetUserAccountId sets UserAccountId field to given value.

### HasUserAccountId

`func (o *LogResponse) HasUserAccountId() bool`

HasUserAccountId returns a boolean if a field has been set.

### GetVersion

`func (o *LogResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


