# Checkout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** | Those are the payment methods that will be available for the link. This field is mutually exclusive with excluded_payment_methods. | [optional] 
**ExcludedPaymentMethods** | Pointer to **[]string** | Payment methods to be excluded from the checkout. This field is mutually exclusive with allowed_payment_methods. | [optional] 
**ExcludeCardNetworks** | Pointer to **[]string** | List of card networks to exclude from the checkout. This field is only applicable for card payments. | [optional] 
**ExpiresAt** | **int64** | It is the time when the link will expire.  It is expressed in seconds since the Unix epoch. The valid range is from 5 minutes to 365 days from the creation date.  | 
**MonthlyInstallmentsEnabled** | Pointer to **bool** | This flag allows you to specify if months without interest will be active. | [optional] 
**MonthlyInstallmentsOptions** | Pointer to **[]int32** | This field allows you to specify the number of months without interest. | [optional] 
**ThreeDsMode** | Pointer to **string** | Indicates the 3DS2 mode for the order, either smart or strict. This property is only applicable when 3DS is enabled. When 3DS is disabled, this field should be null. | [optional] 
**Name** | **string** | Reason for charge | 
**NeedsShippingContact** | Pointer to **bool** | This flag allows you to fill in the shipping information at checkout. | [optional] 
**OnDemandEnabled** | Pointer to **bool** | This flag allows you to specify if the link will be on demand. | [optional] 
**PlanIds** | Pointer to **[]string** | It is a list of plan IDs that will be associated with the order. | [optional] 
**OrderTemplate** | [**CheckoutOrderTemplate**](CheckoutOrderTemplate.md) |  | 
**PaymentsLimitCount** | Pointer to **int32** | It is the number of payments that can be made through the link. | [optional] 
**SuccessUrl** | Pointer to **string** | The URL to redirect to after a successful payment. | [optional] 
**Recurrent** | **bool** | false: single use. true: multiple payments | 
**Type** | **string** | It is the type of link that will be created. It must be a valid type. | 

## Methods

### NewCheckout

`func NewCheckout(expiresAt int64, name string, orderTemplate CheckoutOrderTemplate, recurrent bool, type_ string, ) *Checkout`

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

### HasAllowedPaymentMethods

`func (o *Checkout) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetExcludedPaymentMethods

`func (o *Checkout) GetExcludedPaymentMethods() []string`

GetExcludedPaymentMethods returns the ExcludedPaymentMethods field if non-nil, zero value otherwise.

### GetExcludedPaymentMethodsOk

`func (o *Checkout) GetExcludedPaymentMethodsOk() (*[]string, bool)`

GetExcludedPaymentMethodsOk returns a tuple with the ExcludedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPaymentMethods

`func (o *Checkout) SetExcludedPaymentMethods(v []string)`

SetExcludedPaymentMethods sets ExcludedPaymentMethods field to given value.

### HasExcludedPaymentMethods

`func (o *Checkout) HasExcludedPaymentMethods() bool`

HasExcludedPaymentMethods returns a boolean if a field has been set.

### GetExcludeCardNetworks

`func (o *Checkout) GetExcludeCardNetworks() []string`

GetExcludeCardNetworks returns the ExcludeCardNetworks field if non-nil, zero value otherwise.

### GetExcludeCardNetworksOk

`func (o *Checkout) GetExcludeCardNetworksOk() (*[]string, bool)`

GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCardNetworks

`func (o *Checkout) SetExcludeCardNetworks(v []string)`

SetExcludeCardNetworks sets ExcludeCardNetworks field to given value.

### HasExcludeCardNetworks

`func (o *Checkout) HasExcludeCardNetworks() bool`

HasExcludeCardNetworks returns a boolean if a field has been set.

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

### GetPlanIds

`func (o *Checkout) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *Checkout) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *Checkout) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *Checkout) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

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

### GetSuccessUrl

`func (o *Checkout) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *Checkout) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *Checkout) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *Checkout) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

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


