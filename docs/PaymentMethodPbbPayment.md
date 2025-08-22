# PaymentMethodPbbPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**DeepLink** | **string** | Deep link for the payment, use for mobile apps/flows | 
**ExpiresAt** | **int64** | Expiration date of the charge | 
**ProductType** | **string** | Product type of the charge | 
**RedirectUrl** | **string** | URL to redirect the customer to complete the payment | 
**Reference** | **string** | Reference for the payment | 

## Methods

### NewPaymentMethodPbbPayment

`func NewPaymentMethodPbbPayment(object string, deepLink string, expiresAt int64, productType string, redirectUrl string, reference string, ) *PaymentMethodPbbPayment`

NewPaymentMethodPbbPayment instantiates a new PaymentMethodPbbPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodPbbPaymentWithDefaults

`func NewPaymentMethodPbbPaymentWithDefaults() *PaymentMethodPbbPayment`

NewPaymentMethodPbbPaymentWithDefaults instantiates a new PaymentMethodPbbPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodPbbPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodPbbPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodPbbPayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodPbbPayment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodPbbPayment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodPbbPayment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodPbbPayment) SetObject(v string)`

SetObject sets Object field to given value.


### GetDeepLink

`func (o *PaymentMethodPbbPayment) GetDeepLink() string`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *PaymentMethodPbbPayment) GetDeepLinkOk() (*string, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *PaymentMethodPbbPayment) SetDeepLink(v string)`

SetDeepLink sets DeepLink field to given value.


### GetExpiresAt

`func (o *PaymentMethodPbbPayment) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodPbbPayment) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodPbbPayment) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetProductType

`func (o *PaymentMethodPbbPayment) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PaymentMethodPbbPayment) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PaymentMethodPbbPayment) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRedirectUrl

`func (o *PaymentMethodPbbPayment) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentMethodPbbPayment) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentMethodPbbPayment) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetReference

`func (o *PaymentMethodPbbPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodPbbPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodPbbPayment) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


