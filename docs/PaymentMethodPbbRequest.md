# PaymentMethodPbbRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the payment method | 
**ExpiresAt** | Pointer to **int64** | Expiration date of the payment method, in Unix timestamp format | [optional] 
**ProductType** | **string** | Product type of the payment method, use for the payment method to know the product type | 

## Methods

### NewPaymentMethodPbbRequest

`func NewPaymentMethodPbbRequest(type_ string, productType string, ) *PaymentMethodPbbRequest`

NewPaymentMethodPbbRequest instantiates a new PaymentMethodPbbRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodPbbRequestWithDefaults

`func NewPaymentMethodPbbRequestWithDefaults() *PaymentMethodPbbRequest`

NewPaymentMethodPbbRequestWithDefaults instantiates a new PaymentMethodPbbRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodPbbRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodPbbRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodPbbRequest) SetType(v string)`

SetType sets Type field to given value.


### GetExpiresAt

`func (o *PaymentMethodPbbRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodPbbRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodPbbRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodPbbRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProductType

`func (o *PaymentMethodPbbRequest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PaymentMethodPbbRequest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PaymentMethodPbbRequest) SetProductType(v string)`

SetProductType sets ProductType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


