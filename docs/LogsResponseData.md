# LogsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
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

### NewLogsResponseData

`func NewLogsResponseData() *LogsResponseData`

NewLogsResponseData instantiates a new LogsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsResponseDataWithDefaults

`func NewLogsResponseDataWithDefaults() *LogsResponseData`

NewLogsResponseDataWithDefaults instantiates a new LogsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LogsResponseData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogsResponseData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogsResponseData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LogsResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LogsResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogsResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *LogsResponseData) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogsResponseData) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogsResponseData) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogsResponseData) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLivemode

`func (o *LogsResponseData) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *LogsResponseData) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *LogsResponseData) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *LogsResponseData) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetLoggableId

`func (o *LogsResponseData) GetLoggableId() string`

GetLoggableId returns the LoggableId field if non-nil, zero value otherwise.

### GetLoggableIdOk

`func (o *LogsResponseData) GetLoggableIdOk() (*string, bool)`

GetLoggableIdOk returns a tuple with the LoggableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableId

`func (o *LogsResponseData) SetLoggableId(v string)`

SetLoggableId sets LoggableId field to given value.

### HasLoggableId

`func (o *LogsResponseData) HasLoggableId() bool`

HasLoggableId returns a boolean if a field has been set.

### SetLoggableIdNil

`func (o *LogsResponseData) SetLoggableIdNil(b bool)`

 SetLoggableIdNil sets the value for LoggableId to be an explicit nil

### UnsetLoggableId
`func (o *LogsResponseData) UnsetLoggableId()`

UnsetLoggableId ensures that no value is present for LoggableId, not even an explicit nil
### GetLoggableType

`func (o *LogsResponseData) GetLoggableType() string`

GetLoggableType returns the LoggableType field if non-nil, zero value otherwise.

### GetLoggableTypeOk

`func (o *LogsResponseData) GetLoggableTypeOk() (*string, bool)`

GetLoggableTypeOk returns a tuple with the LoggableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableType

`func (o *LogsResponseData) SetLoggableType(v string)`

SetLoggableType sets LoggableType field to given value.

### HasLoggableType

`func (o *LogsResponseData) HasLoggableType() bool`

HasLoggableType returns a boolean if a field has been set.

### SetLoggableTypeNil

`func (o *LogsResponseData) SetLoggableTypeNil(b bool)`

 SetLoggableTypeNil sets the value for LoggableType to be an explicit nil

### UnsetLoggableType
`func (o *LogsResponseData) UnsetLoggableType()`

UnsetLoggableType ensures that no value is present for LoggableType, not even an explicit nil
### GetMethod

`func (o *LogsResponseData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogsResponseData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogsResponseData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogsResponseData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOauthTokenId

`func (o *LogsResponseData) GetOauthTokenId() string`

GetOauthTokenId returns the OauthTokenId field if non-nil, zero value otherwise.

### GetOauthTokenIdOk

`func (o *LogsResponseData) GetOauthTokenIdOk() (*string, bool)`

GetOauthTokenIdOk returns a tuple with the OauthTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenId

`func (o *LogsResponseData) SetOauthTokenId(v string)`

SetOauthTokenId sets OauthTokenId field to given value.

### HasOauthTokenId

`func (o *LogsResponseData) HasOauthTokenId() bool`

HasOauthTokenId returns a boolean if a field has been set.

### SetOauthTokenIdNil

`func (o *LogsResponseData) SetOauthTokenIdNil(b bool)`

 SetOauthTokenIdNil sets the value for OauthTokenId to be an explicit nil

### UnsetOauthTokenId
`func (o *LogsResponseData) UnsetOauthTokenId()`

UnsetOauthTokenId ensures that no value is present for OauthTokenId, not even an explicit nil
### GetQueryString

`func (o *LogsResponseData) GetQueryString() map[string]interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogsResponseData) GetQueryStringOk() (*map[string]interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogsResponseData) SetQueryString(v map[string]interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LogsResponseData) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetRelated

`func (o *LogsResponseData) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *LogsResponseData) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *LogsResponseData) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *LogsResponseData) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRequestBody

`func (o *LogsResponseData) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *LogsResponseData) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *LogsResponseData) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *LogsResponseData) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *LogsResponseData) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *LogsResponseData) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *LogsResponseData) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *LogsResponseData) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseBody

`func (o *LogsResponseData) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *LogsResponseData) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *LogsResponseData) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *LogsResponseData) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *LogsResponseData) GetResponseHeaders() map[string]string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *LogsResponseData) GetResponseHeadersOk() (*map[string]string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *LogsResponseData) SetResponseHeaders(v map[string]string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *LogsResponseData) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetSearchableTags

`func (o *LogsResponseData) GetSearchableTags() []string`

GetSearchableTags returns the SearchableTags field if non-nil, zero value otherwise.

### GetSearchableTagsOk

`func (o *LogsResponseData) GetSearchableTagsOk() (*[]string, bool)`

GetSearchableTagsOk returns a tuple with the SearchableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableTags

`func (o *LogsResponseData) SetSearchableTags(v []string)`

SetSearchableTags sets SearchableTags field to given value.

### HasSearchableTags

`func (o *LogsResponseData) HasSearchableTags() bool`

HasSearchableTags returns a boolean if a field has been set.

### GetStatus

`func (o *LogsResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogsResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogsResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogsResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LogsResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogsResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogsResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LogsResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *LogsResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogsResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogsResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LogsResponseData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserAccountId

`func (o *LogsResponseData) GetUserAccountId() string`

GetUserAccountId returns the UserAccountId field if non-nil, zero value otherwise.

### GetUserAccountIdOk

`func (o *LogsResponseData) GetUserAccountIdOk() (*string, bool)`

GetUserAccountIdOk returns a tuple with the UserAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountId

`func (o *LogsResponseData) SetUserAccountId(v string)`

SetUserAccountId sets UserAccountId field to given value.

### HasUserAccountId

`func (o *LogsResponseData) HasUserAccountId() bool`

HasUserAccountId returns a boolean if a field has been set.

### GetVersion

`func (o *LogsResponseData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogsResponseData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogsResponseData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogsResponseData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


