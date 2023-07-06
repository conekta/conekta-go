# ApiKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates if the webhook key is active | [optional] 
**Description** | Pointer to **string** | A name or brief explanation of what this api key is used for | [optional] 

## Methods

### NewApiKeyUpdateRequest

`func NewApiKeyUpdateRequest() *ApiKeyUpdateRequest`

NewApiKeyUpdateRequest instantiates a new ApiKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyUpdateRequestWithDefaults

`func NewApiKeyUpdateRequestWithDefaults() *ApiKeyUpdateRequest`

NewApiKeyUpdateRequestWithDefaults instantiates a new ApiKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


