# UpdatePaymentMethodsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the payment method holder | [optional] 
**ExpiresAt** | Pointer to **int64** | The expiration date of the payment method in Unix timestamp format | [optional] 

## Methods

### NewUpdatePaymentMethodsCard

`func NewUpdatePaymentMethodsCard() *UpdatePaymentMethodsCard`

NewUpdatePaymentMethodsCard instantiates a new UpdatePaymentMethodsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodsCardWithDefaults

`func NewUpdatePaymentMethodsCardWithDefaults() *UpdatePaymentMethodsCard`

NewUpdatePaymentMethodsCardWithDefaults instantiates a new UpdatePaymentMethodsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePaymentMethodsCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePaymentMethodsCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePaymentMethodsCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePaymentMethodsCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UpdatePaymentMethodsCard) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdatePaymentMethodsCard) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdatePaymentMethodsCard) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdatePaymentMethodsCard) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


