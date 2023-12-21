# OrderNextActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectToUrl** | Pointer to [**OrderNextActionResponseRedirectToUrl**](OrderNextActionResponseRedirectToUrl.md) |  | [optional] 
**Type** | Pointer to **string** | Indicates the type of action to be taken | [optional] 

## Methods

### NewOrderNextActionResponse

`func NewOrderNextActionResponse() *OrderNextActionResponse`

NewOrderNextActionResponse instantiates a new OrderNextActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderNextActionResponseWithDefaults

`func NewOrderNextActionResponseWithDefaults() *OrderNextActionResponse`

NewOrderNextActionResponseWithDefaults instantiates a new OrderNextActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectToUrl

`func (o *OrderNextActionResponse) GetRedirectToUrl() OrderNextActionResponseRedirectToUrl`

GetRedirectToUrl returns the RedirectToUrl field if non-nil, zero value otherwise.

### GetRedirectToUrlOk

`func (o *OrderNextActionResponse) GetRedirectToUrlOk() (*OrderNextActionResponseRedirectToUrl, bool)`

GetRedirectToUrlOk returns a tuple with the RedirectToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToUrl

`func (o *OrderNextActionResponse) SetRedirectToUrl(v OrderNextActionResponseRedirectToUrl)`

SetRedirectToUrl sets RedirectToUrl field to given value.

### HasRedirectToUrl

`func (o *OrderNextActionResponse) HasRedirectToUrl() bool`

HasRedirectToUrl returns a boolean if a field has been set.

### GetType

`func (o *OrderNextActionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderNextActionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderNextActionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderNextActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


