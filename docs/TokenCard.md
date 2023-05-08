# TokenCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cvc** | **string** | It is a value that allows identifying the security code of the card. | 
**DeviceFingerprint** | Pointer to **string** | It is a value that allows identifying the device fingerprint. | [optional] 
**ExpMonth** | **string** | It is a value that allows identifying the expiration month of the card. | 
**ExpYear** | **string** | It is a value that allows identifying the expiration year of the card. | 
**Name** | **string** | It is a value that allows identifying the name of the cardholder. | 
**Number** | **string** | It is a value that allows identifying the number of the card. | 

## Methods

### NewTokenCard

`func NewTokenCard(cvc string, expMonth string, expYear string, name string, number string, ) *TokenCard`

NewTokenCard instantiates a new TokenCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCardWithDefaults

`func NewTokenCardWithDefaults() *TokenCard`

NewTokenCardWithDefaults instantiates a new TokenCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvc

`func (o *TokenCard) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *TokenCard) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *TokenCard) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetDeviceFingerprint

`func (o *TokenCard) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *TokenCard) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *TokenCard) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *TokenCard) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetExpMonth

`func (o *TokenCard) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *TokenCard) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *TokenCard) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *TokenCard) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *TokenCard) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *TokenCard) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.


### GetName

`func (o *TokenCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenCard) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *TokenCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TokenCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TokenCard) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


