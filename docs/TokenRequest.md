# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**TokenRequestCard**](TokenRequestCard.md) |  | [optional] 
**Checkout** | Pointer to [**TokenRequestCheckout**](TokenRequestCheckout.md) |  | [optional] 

## Methods

### NewTokenRequest

`func NewTokenRequest() *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *TokenRequest) GetCard() TokenRequestCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *TokenRequest) GetCardOk() (*TokenRequestCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *TokenRequest) SetCard(v TokenRequestCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *TokenRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCheckout

`func (o *TokenRequest) GetCheckout() TokenRequestCheckout`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *TokenRequest) GetCheckoutOk() (*TokenRequestCheckout, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *TokenRequest) SetCheckout(v TokenRequestCheckout)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *TokenRequest) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


