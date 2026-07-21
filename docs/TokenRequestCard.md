# TokenRequestCard

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

### NewTokenRequestCard

`func NewTokenRequestCard(cvc string, expMonth string, expYear string, name string, number string, ) *TokenRequestCard`

NewTokenRequestCard instantiates a new TokenRequestCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestCardWithDefaults

`func NewTokenRequestCardWithDefaults() *TokenRequestCard`

NewTokenRequestCardWithDefaults instantiates a new TokenRequestCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvc

`func (o *TokenRequestCard) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *TokenRequestCard) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *TokenRequestCard) SetCvc(v string)`

SetCvc sets Cvc field to given value.


### GetDeviceFingerprint

`func (o *TokenRequestCard) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *TokenRequestCard) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *TokenRequestCard) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *TokenRequestCard) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetExpMonth

`func (o *TokenRequestCard) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *TokenRequestCard) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *TokenRequestCard) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *TokenRequestCard) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *TokenRequestCard) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *TokenRequestCard) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.


### GetName

`func (o *TokenRequestCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenRequestCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenRequestCard) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *TokenRequestCard) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TokenRequestCard) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TokenRequestCard) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


