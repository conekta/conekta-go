# LogResponseForRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** |  | 
**Id** | **string** |  | 
**IpAddress** | Pointer to **string** |  | [optional] 
**Livemode** | **bool** |  | 
**LoggableId** | Pointer to **string** |  | [optional] 
**LoggableType** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**OauthTokenId** | Pointer to **string** |  | [optional] 
**QueryString** | Pointer to **map[string]interface{}** |  | [optional] 
**Related** | Pointer to **string** |  | [optional] 
**RequestBody** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestHeaders** | Pointer to  |  | [optional] 
**ResponseBody** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseHeaders** | Pointer to  |  | [optional] 
**SearchableTags** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserAccountId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewLogResponseForRequest

`func NewLogResponseForRequest(createdAt int64, id string, livemode bool, ) *LogResponseForRequest`

NewLogResponseForRequest instantiates a new LogResponseForRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResponseForRequestWithDefaults

`func NewLogResponseForRequestWithDefaults() *LogResponseForRequest`

NewLogResponseForRequestWithDefaults instantiates a new LogResponseForRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LogResponseForRequest) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogResponseForRequest) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogResponseForRequest) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *LogResponseForRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogResponseForRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogResponseForRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *LogResponseForRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LogResponseForRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LogResponseForRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LogResponseForRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLivemode

`func (o *LogResponseForRequest) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *LogResponseForRequest) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *LogResponseForRequest) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetLoggableId

`func (o *LogResponseForRequest) GetLoggableId() string`

GetLoggableId returns the LoggableId field if non-nil, zero value otherwise.

### GetLoggableIdOk

`func (o *LogResponseForRequest) GetLoggableIdOk() (*string, bool)`

GetLoggableIdOk returns a tuple with the LoggableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableId

`func (o *LogResponseForRequest) SetLoggableId(v string)`

SetLoggableId sets LoggableId field to given value.

### HasLoggableId

`func (o *LogResponseForRequest) HasLoggableId() bool`

HasLoggableId returns a boolean if a field has been set.

### GetLoggableType

`func (o *LogResponseForRequest) GetLoggableType() string`

GetLoggableType returns the LoggableType field if non-nil, zero value otherwise.

### GetLoggableTypeOk

`func (o *LogResponseForRequest) GetLoggableTypeOk() (*string, bool)`

GetLoggableTypeOk returns a tuple with the LoggableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggableType

`func (o *LogResponseForRequest) SetLoggableType(v string)`

SetLoggableType sets LoggableType field to given value.

### HasLoggableType

`func (o *LogResponseForRequest) HasLoggableType() bool`

HasLoggableType returns a boolean if a field has been set.

### GetMethod

`func (o *LogResponseForRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogResponseForRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogResponseForRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogResponseForRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOauthTokenId

`func (o *LogResponseForRequest) GetOauthTokenId() string`

GetOauthTokenId returns the OauthTokenId field if non-nil, zero value otherwise.

### GetOauthTokenIdOk

`func (o *LogResponseForRequest) GetOauthTokenIdOk() (*string, bool)`

GetOauthTokenIdOk returns a tuple with the OauthTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthTokenId

`func (o *LogResponseForRequest) SetOauthTokenId(v string)`

SetOauthTokenId sets OauthTokenId field to given value.

### HasOauthTokenId

`func (o *LogResponseForRequest) HasOauthTokenId() bool`

HasOauthTokenId returns a boolean if a field has been set.

### GetQueryString

`func (o *LogResponseForRequest) GetQueryString() map[string]interface{}`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogResponseForRequest) GetQueryStringOk() (*map[string]interface{}, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogResponseForRequest) SetQueryString(v map[string]interface{})`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LogResponseForRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetRelated

`func (o *LogResponseForRequest) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *LogResponseForRequest) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *LogResponseForRequest) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *LogResponseForRequest) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRequestBody

`func (o *LogResponseForRequest) GetRequestBody() map[string]interface{}`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *LogResponseForRequest) GetRequestBodyOk() (*map[string]interface{}, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *LogResponseForRequest) SetRequestBody(v map[string]interface{})`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *LogResponseForRequest) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *LogResponseForRequest) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *LogResponseForRequest) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *LogResponseForRequest) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *LogResponseForRequest) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *LogResponseForRequest) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *LogResponseForRequest) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetResponseBody

`func (o *LogResponseForRequest) GetResponseBody() map[string]interface{}`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *LogResponseForRequest) GetResponseBodyOk() (*map[string]interface{}, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *LogResponseForRequest) SetResponseBody(v map[string]interface{})`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *LogResponseForRequest) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *LogResponseForRequest) GetResponseHeaders() map[string]string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *LogResponseForRequest) GetResponseHeadersOk() (*map[string]string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *LogResponseForRequest) SetResponseHeaders(v map[string]string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *LogResponseForRequest) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *LogResponseForRequest) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *LogResponseForRequest) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetSearchableTags

`func (o *LogResponseForRequest) GetSearchableTags() []string`

GetSearchableTags returns the SearchableTags field if non-nil, zero value otherwise.

### GetSearchableTagsOk

`func (o *LogResponseForRequest) GetSearchableTagsOk() (*[]string, bool)`

GetSearchableTagsOk returns a tuple with the SearchableTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableTags

`func (o *LogResponseForRequest) SetSearchableTags(v []string)`

SetSearchableTags sets SearchableTags field to given value.

### HasSearchableTags

`func (o *LogResponseForRequest) HasSearchableTags() bool`

HasSearchableTags returns a boolean if a field has been set.

### GetStatus

`func (o *LogResponseForRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogResponseForRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogResponseForRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogResponseForRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LogResponseForRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogResponseForRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogResponseForRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LogResponseForRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *LogResponseForRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogResponseForRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogResponseForRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LogResponseForRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserAccountId

`func (o *LogResponseForRequest) GetUserAccountId() string`

GetUserAccountId returns the UserAccountId field if non-nil, zero value otherwise.

### GetUserAccountIdOk

`func (o *LogResponseForRequest) GetUserAccountIdOk() (*string, bool)`

GetUserAccountIdOk returns a tuple with the UserAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountId

`func (o *LogResponseForRequest) SetUserAccountId(v string)`

SetUserAccountId sets UserAccountId field to given value.

### HasUserAccountId

`func (o *LogResponseForRequest) HasUserAccountId() bool`

HasUserAccountId returns a boolean if a field has been set.

### GetVersion

`func (o *LogResponseForRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogResponseForRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogResponseForRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogResponseForRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


