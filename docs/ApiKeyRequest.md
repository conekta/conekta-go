# ApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Indicates if the api key is active | 
**Description** | **string** | Detail of the use that will be given to the api key | 
**Role** | **string** |  | 

## Methods

### NewApiKeyRequest

`func NewApiKeyRequest(active bool, description string, role string, ) *ApiKeyRequest`

NewApiKeyRequest instantiates a new ApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyRequestWithDefaults

`func NewApiKeyRequestWithDefaults() *ApiKeyRequest`

NewApiKeyRequestWithDefaults instantiates a new ApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDescription

`func (o *ApiKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRole

`func (o *ApiKeyRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiKeyRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiKeyRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


