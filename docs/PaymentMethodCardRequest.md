# PaymentMethodCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of payment method | 
**TokenId** | **string** | Token id that will be used to create a \&quot;card\&quot; type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards. | 

## Methods

### NewPaymentMethodCardRequest

`func NewPaymentMethodCardRequest(type_ string, tokenId string, ) *PaymentMethodCardRequest`

NewPaymentMethodCardRequest instantiates a new PaymentMethodCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardRequestWithDefaults

`func NewPaymentMethodCardRequestWithDefaults() *PaymentMethodCardRequest`

NewPaymentMethodCardRequestWithDefaults instantiates a new PaymentMethodCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCardRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCardRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCardRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTokenId

`func (o *PaymentMethodCardRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentMethodCardRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentMethodCardRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


