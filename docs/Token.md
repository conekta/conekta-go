# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**NullableTokenCard**](TokenCard.md) |  | [optional] 
**Checkout** | Pointer to [**NullableTokenCheckout**](TokenCheckout.md) |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *Token) GetCard() TokenCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *Token) GetCardOk() (*TokenCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *Token) SetCard(v TokenCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *Token) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *Token) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *Token) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetCheckout

`func (o *Token) GetCheckout() TokenCheckout`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *Token) GetCheckoutOk() (*TokenCheckout, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *Token) SetCheckout(v TokenCheckout)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *Token) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### SetCheckoutNil

`func (o *Token) SetCheckoutNil(b bool)`

 SetCheckoutNil sets the value for Checkout to be an explicit nil

### UnsetCheckout
`func (o *Token) UnsetCheckout()`

UnsetCheckout ensures that no value is present for Checkout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


