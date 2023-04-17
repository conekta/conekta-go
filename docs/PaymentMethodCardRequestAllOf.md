# PaymentMethodCardRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **string** | Token id that will be used to create a \&quot;card\&quot; type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards. | [optional] 

## Methods

### NewPaymentMethodCardRequestAllOf

`func NewPaymentMethodCardRequestAllOf() *PaymentMethodCardRequestAllOf`

NewPaymentMethodCardRequestAllOf instantiates a new PaymentMethodCardRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardRequestAllOfWithDefaults

`func NewPaymentMethodCardRequestAllOfWithDefaults() *PaymentMethodCardRequestAllOf`

NewPaymentMethodCardRequestAllOfWithDefaults instantiates a new PaymentMethodCardRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *PaymentMethodCardRequestAllOf) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentMethodCardRequestAllOf) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentMethodCardRequestAllOf) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PaymentMethodCardRequestAllOf) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


