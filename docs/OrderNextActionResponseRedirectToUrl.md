# OrderNextActionResponseRedirectToUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | pay.conekta.com/{id} Indicates the url of the Conekta component to authenticate the flow through 3DS2. | [optional] 
**ReturnUrl** | Pointer to **string** | Indicates the url to which the 3DS2 flow returns at the end, when the integration is redirected. | [optional] 

## Methods

### NewOrderNextActionResponseRedirectToUrl

`func NewOrderNextActionResponseRedirectToUrl() *OrderNextActionResponseRedirectToUrl`

NewOrderNextActionResponseRedirectToUrl instantiates a new OrderNextActionResponseRedirectToUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderNextActionResponseRedirectToUrlWithDefaults

`func NewOrderNextActionResponseRedirectToUrlWithDefaults() *OrderNextActionResponseRedirectToUrl`

NewOrderNextActionResponseRedirectToUrlWithDefaults instantiates a new OrderNextActionResponseRedirectToUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrderNextActionResponseRedirectToUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrderNextActionResponseRedirectToUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrderNextActionResponseRedirectToUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrderNextActionResponseRedirectToUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetReturnUrl

`func (o *OrderNextActionResponseRedirectToUrl) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OrderNextActionResponseRedirectToUrl) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OrderNextActionResponseRedirectToUrl) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OrderNextActionResponseRedirectToUrl) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


