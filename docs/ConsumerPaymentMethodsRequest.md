# ConsumerPaymentMethodsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of payment method | 
**TokenId** | **string** | Token id that will be used to create a \&quot;card\&quot; type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards. | 
**ExpiresAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewConsumerPaymentMethodsRequest

`func NewConsumerPaymentMethodsRequest(type_ string, tokenId string, ) *ConsumerPaymentMethodsRequest`

NewConsumerPaymentMethodsRequest instantiates a new ConsumerPaymentMethodsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerPaymentMethodsRequestWithDefaults

`func NewConsumerPaymentMethodsRequestWithDefaults() *ConsumerPaymentMethodsRequest`

NewConsumerPaymentMethodsRequestWithDefaults instantiates a new ConsumerPaymentMethodsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConsumerPaymentMethodsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerPaymentMethodsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerPaymentMethodsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTokenId

`func (o *ConsumerPaymentMethodsRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ConsumerPaymentMethodsRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ConsumerPaymentMethodsRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetExpiresAt

`func (o *ConsumerPaymentMethodsRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConsumerPaymentMethodsRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConsumerPaymentMethodsRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ConsumerPaymentMethodsRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


