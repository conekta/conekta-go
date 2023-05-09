# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkout** | Pointer to [**NullableTokenResponseCheckout**](TokenResponseCheckout.md) |  | [optional] 
**Id** | **string** | Unique identifier for the token generated by Conekta. | 
**Livemode** | **bool** | Indicates whether the token is in live mode or test mode. | 
**Object** | **string** | Indicates the type of object, in this case token | 
**Used** | **bool** | Indicates if the token has been used | 

## Methods

### NewTokenResponse

`func NewTokenResponse(id string, livemode bool, object string, used bool, ) *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckout

`func (o *TokenResponse) GetCheckout() TokenResponseCheckout`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *TokenResponse) GetCheckoutOk() (*TokenResponseCheckout, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *TokenResponse) SetCheckout(v TokenResponseCheckout)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *TokenResponse) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### SetCheckoutNil

`func (o *TokenResponse) SetCheckoutNil(b bool)`

 SetCheckoutNil sets the value for Checkout to be an explicit nil

### UnsetCheckout
`func (o *TokenResponse) UnsetCheckout()`

UnsetCheckout ensures that no value is present for Checkout, not even an explicit nil
### GetId

`func (o *TokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLivemode

`func (o *TokenResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *TokenResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *TokenResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetObject

`func (o *TokenResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TokenResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TokenResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetUsed

`func (o *TokenResponse) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *TokenResponse) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *TokenResponse) SetUsed(v bool)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

