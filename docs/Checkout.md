# Checkout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | **[]string** | Those are the payment methods that will be available for the link | 
**ExpiresAt** | **int64** | It is the time when the link will expire. It is expressed in seconds since the Unix epoch. The valid range is from 2 to 365 days (the valid range will be taken from the next day of the creation date at 00:01 hrs)  | 
**MonthlyInstallmentsEnabled** | Pointer to **bool** | This flag allows you to specify if months without interest will be active. | [optional] 
**MonthlyInstallmentsOptions** | Pointer to **[]int32** | This field allows you to specify the number of months without interest. | [optional] 
**ThreeDsMode** | Pointer to **string** | Indicates the 3DS2 mode for the order, either smart or strict. | [optional] 
**Name** | **string** | Reason for charge | 
**NeedsShippingContact** | Pointer to **bool** | This flag allows you to fill in the shipping information at checkout. | [optional] 
**OnDemandEnabled** | Pointer to **NullableBool** | This flag allows you to specify if the link will be on demand. | [optional] 
**OrderTemplate** | [**CheckoutOrderTemplate**](CheckoutOrderTemplate.md) |  | 
**PaymentsLimitCount** | Pointer to **int32** | It is the number of payments that can be made through the link. | [optional] 
**Recurrent** | **bool** | false: single use. true: multiple payments | 
**Type** | **string** | It is the type of link that will be created. It must be a valid type. | 

## Methods

### NewCheckout

`func NewCheckout(allowedPaymentMethods []string, expiresAt int64, name string, orderTemplate CheckoutOrderTemplate, recurrent bool, type_ string, ) *Checkout`

NewCheckout instantiates a new Checkout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutWithDefaults

`func NewCheckoutWithDefaults() *Checkout`

NewCheckoutWithDefaults instantiates a new Checkout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *Checkout) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *Checkout) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *Checkout) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.


### GetExpiresAt

`func (o *Checkout) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Checkout) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Checkout) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetMonthlyInstallmentsEnabled

`func (o *Checkout) GetMonthlyInstallmentsEnabled() bool`

GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsEnabledOk

`func (o *Checkout) GetMonthlyInstallmentsEnabledOk() (*bool, bool)`

GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsEnabled

`func (o *Checkout) SetMonthlyInstallmentsEnabled(v bool)`

SetMonthlyInstallmentsEnabled sets MonthlyInstallmentsEnabled field to given value.

### HasMonthlyInstallmentsEnabled

`func (o *Checkout) HasMonthlyInstallmentsEnabled() bool`

HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.

### GetMonthlyInstallmentsOptions

`func (o *Checkout) GetMonthlyInstallmentsOptions() []int32`

GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOptionsOk

`func (o *Checkout) GetMonthlyInstallmentsOptionsOk() (*[]int32, bool)`

GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsOptions

`func (o *Checkout) SetMonthlyInstallmentsOptions(v []int32)`

SetMonthlyInstallmentsOptions sets MonthlyInstallmentsOptions field to given value.

### HasMonthlyInstallmentsOptions

`func (o *Checkout) HasMonthlyInstallmentsOptions() bool`

HasMonthlyInstallmentsOptions returns a boolean if a field has been set.

### GetThreeDsMode

`func (o *Checkout) GetThreeDsMode() string`

GetThreeDsMode returns the ThreeDsMode field if non-nil, zero value otherwise.

### GetThreeDsModeOk

`func (o *Checkout) GetThreeDsModeOk() (*string, bool)`

GetThreeDsModeOk returns a tuple with the ThreeDsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsMode

`func (o *Checkout) SetThreeDsMode(v string)`

SetThreeDsMode sets ThreeDsMode field to given value.

### HasThreeDsMode

`func (o *Checkout) HasThreeDsMode() bool`

HasThreeDsMode returns a boolean if a field has been set.

### GetName

`func (o *Checkout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Checkout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Checkout) SetName(v string)`

SetName sets Name field to given value.


### GetNeedsShippingContact

`func (o *Checkout) GetNeedsShippingContact() bool`

GetNeedsShippingContact returns the NeedsShippingContact field if non-nil, zero value otherwise.

### GetNeedsShippingContactOk

`func (o *Checkout) GetNeedsShippingContactOk() (*bool, bool)`

GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsShippingContact

`func (o *Checkout) SetNeedsShippingContact(v bool)`

SetNeedsShippingContact sets NeedsShippingContact field to given value.

### HasNeedsShippingContact

`func (o *Checkout) HasNeedsShippingContact() bool`

HasNeedsShippingContact returns a boolean if a field has been set.

### GetOnDemandEnabled

`func (o *Checkout) GetOnDemandEnabled() bool`

GetOnDemandEnabled returns the OnDemandEnabled field if non-nil, zero value otherwise.

### GetOnDemandEnabledOk

`func (o *Checkout) GetOnDemandEnabledOk() (*bool, bool)`

GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandEnabled

`func (o *Checkout) SetOnDemandEnabled(v bool)`

SetOnDemandEnabled sets OnDemandEnabled field to given value.

### HasOnDemandEnabled

`func (o *Checkout) HasOnDemandEnabled() bool`

HasOnDemandEnabled returns a boolean if a field has been set.

### SetOnDemandEnabledNil

`func (o *Checkout) SetOnDemandEnabledNil(b bool)`

 SetOnDemandEnabledNil sets the value for OnDemandEnabled to be an explicit nil

### UnsetOnDemandEnabled
`func (o *Checkout) UnsetOnDemandEnabled()`

UnsetOnDemandEnabled ensures that no value is present for OnDemandEnabled, not even an explicit nil
### GetOrderTemplate

`func (o *Checkout) GetOrderTemplate() CheckoutOrderTemplate`

GetOrderTemplate returns the OrderTemplate field if non-nil, zero value otherwise.

### GetOrderTemplateOk

`func (o *Checkout) GetOrderTemplateOk() (*CheckoutOrderTemplate, bool)`

GetOrderTemplateOk returns a tuple with the OrderTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTemplate

`func (o *Checkout) SetOrderTemplate(v CheckoutOrderTemplate)`

SetOrderTemplate sets OrderTemplate field to given value.


### GetPaymentsLimitCount

`func (o *Checkout) GetPaymentsLimitCount() int32`

GetPaymentsLimitCount returns the PaymentsLimitCount field if non-nil, zero value otherwise.

### GetPaymentsLimitCountOk

`func (o *Checkout) GetPaymentsLimitCountOk() (*int32, bool)`

GetPaymentsLimitCountOk returns a tuple with the PaymentsLimitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsLimitCount

`func (o *Checkout) SetPaymentsLimitCount(v int32)`

SetPaymentsLimitCount sets PaymentsLimitCount field to given value.

### HasPaymentsLimitCount

`func (o *Checkout) HasPaymentsLimitCount() bool`

HasPaymentsLimitCount returns a boolean if a field has been set.

### GetRecurrent

`func (o *Checkout) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *Checkout) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *Checkout) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.


### GetType

`func (o *Checkout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Checkout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Checkout) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


