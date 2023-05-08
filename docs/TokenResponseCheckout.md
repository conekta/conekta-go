# TokenResponseCheckout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** |  | [optional] 
**CanNotExpire** | Pointer to **bool** | Indicates if the checkout can not expire. | [optional] 
**EmailsSent** | Pointer to **int32** |  | [optional] 
**ExcludeCardNetworks** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** | Date and time when the checkout expires. | [optional] 
**FailureUrl** | Pointer to **string** | URL to redirect the customer to if the payment process fails. | [optional] 
**Force3dsFlow** | Pointer to **bool** | Indicates if the checkout forces the 3DS flow. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MonthlyInstallmentsEnabled** | Pointer to **bool** | Indicates if the checkout allows monthly installments. | [optional] 
**MonthlyInstallmentsOptions** | Pointer to **[]int32** | List of monthly installments options. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedsShippingContact** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** | Indicates the type of object, in this case checkout. | [optional] 
**OnDemandEnabled** | Pointer to **bool** | Indicates if the checkout allows on demand payments. | [optional] 
**PaidPaymentsCount** | Pointer to **int32** | Number of payments that have been paid. | [optional] 
**Recurrent** | Pointer to **bool** | Indicates if the checkout is recurrent. | [optional] 
**SmsSent** | Pointer to **int32** |  | [optional] 
**StartsAt** | Pointer to **int64** | Date and time when the checkout starts. | [optional] 
**Status** | Pointer to **string** | Status of the checkout. | [optional] 
**SuccessUrl** | Pointer to **string** | URL to redirect the customer to after the payment process is completed. | [optional] 
**Type** | Pointer to **string** | Type of checkout. | [optional] 

## Methods

### NewTokenResponseCheckout

`func NewTokenResponseCheckout() *TokenResponseCheckout`

NewTokenResponseCheckout instantiates a new TokenResponseCheckout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseCheckoutWithDefaults

`func NewTokenResponseCheckoutWithDefaults() *TokenResponseCheckout`

NewTokenResponseCheckoutWithDefaults instantiates a new TokenResponseCheckout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *TokenResponseCheckout) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *TokenResponseCheckout) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *TokenResponseCheckout) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *TokenResponseCheckout) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetCanNotExpire

`func (o *TokenResponseCheckout) GetCanNotExpire() bool`

GetCanNotExpire returns the CanNotExpire field if non-nil, zero value otherwise.

### GetCanNotExpireOk

`func (o *TokenResponseCheckout) GetCanNotExpireOk() (*bool, bool)`

GetCanNotExpireOk returns a tuple with the CanNotExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanNotExpire

`func (o *TokenResponseCheckout) SetCanNotExpire(v bool)`

SetCanNotExpire sets CanNotExpire field to given value.

### HasCanNotExpire

`func (o *TokenResponseCheckout) HasCanNotExpire() bool`

HasCanNotExpire returns a boolean if a field has been set.

### GetEmailsSent

`func (o *TokenResponseCheckout) GetEmailsSent() int32`

GetEmailsSent returns the EmailsSent field if non-nil, zero value otherwise.

### GetEmailsSentOk

`func (o *TokenResponseCheckout) GetEmailsSentOk() (*int32, bool)`

GetEmailsSentOk returns a tuple with the EmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsSent

`func (o *TokenResponseCheckout) SetEmailsSent(v int32)`

SetEmailsSent sets EmailsSent field to given value.

### HasEmailsSent

`func (o *TokenResponseCheckout) HasEmailsSent() bool`

HasEmailsSent returns a boolean if a field has been set.

### GetExcludeCardNetworks

`func (o *TokenResponseCheckout) GetExcludeCardNetworks() []string`

GetExcludeCardNetworks returns the ExcludeCardNetworks field if non-nil, zero value otherwise.

### GetExcludeCardNetworksOk

`func (o *TokenResponseCheckout) GetExcludeCardNetworksOk() (*[]string, bool)`

GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCardNetworks

`func (o *TokenResponseCheckout) SetExcludeCardNetworks(v []string)`

SetExcludeCardNetworks sets ExcludeCardNetworks field to given value.

### HasExcludeCardNetworks

`func (o *TokenResponseCheckout) HasExcludeCardNetworks() bool`

HasExcludeCardNetworks returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenResponseCheckout) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenResponseCheckout) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenResponseCheckout) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenResponseCheckout) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailureUrl

`func (o *TokenResponseCheckout) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *TokenResponseCheckout) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *TokenResponseCheckout) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *TokenResponseCheckout) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetForce3dsFlow

`func (o *TokenResponseCheckout) GetForce3dsFlow() bool`

GetForce3dsFlow returns the Force3dsFlow field if non-nil, zero value otherwise.

### GetForce3dsFlowOk

`func (o *TokenResponseCheckout) GetForce3dsFlowOk() (*bool, bool)`

GetForce3dsFlowOk returns a tuple with the Force3dsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce3dsFlow

`func (o *TokenResponseCheckout) SetForce3dsFlow(v bool)`

SetForce3dsFlow sets Force3dsFlow field to given value.

### HasForce3dsFlow

`func (o *TokenResponseCheckout) HasForce3dsFlow() bool`

HasForce3dsFlow returns a boolean if a field has been set.

### GetId

`func (o *TokenResponseCheckout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenResponseCheckout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenResponseCheckout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenResponseCheckout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *TokenResponseCheckout) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *TokenResponseCheckout) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *TokenResponseCheckout) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *TokenResponseCheckout) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMetadata

`func (o *TokenResponseCheckout) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TokenResponseCheckout) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TokenResponseCheckout) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TokenResponseCheckout) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonthlyInstallmentsEnabled

`func (o *TokenResponseCheckout) GetMonthlyInstallmentsEnabled() bool`

GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsEnabledOk

`func (o *TokenResponseCheckout) GetMonthlyInstallmentsEnabledOk() (*bool, bool)`

GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsEnabled

`func (o *TokenResponseCheckout) SetMonthlyInstallmentsEnabled(v bool)`

SetMonthlyInstallmentsEnabled sets MonthlyInstallmentsEnabled field to given value.

### HasMonthlyInstallmentsEnabled

`func (o *TokenResponseCheckout) HasMonthlyInstallmentsEnabled() bool`

HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.

### GetMonthlyInstallmentsOptions

`func (o *TokenResponseCheckout) GetMonthlyInstallmentsOptions() []int32`

GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOptionsOk

`func (o *TokenResponseCheckout) GetMonthlyInstallmentsOptionsOk() (*[]int32, bool)`

GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsOptions

`func (o *TokenResponseCheckout) SetMonthlyInstallmentsOptions(v []int32)`

SetMonthlyInstallmentsOptions sets MonthlyInstallmentsOptions field to given value.

### HasMonthlyInstallmentsOptions

`func (o *TokenResponseCheckout) HasMonthlyInstallmentsOptions() bool`

HasMonthlyInstallmentsOptions returns a boolean if a field has been set.

### GetName

`func (o *TokenResponseCheckout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenResponseCheckout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenResponseCheckout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenResponseCheckout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedsShippingContact

`func (o *TokenResponseCheckout) GetNeedsShippingContact() bool`

GetNeedsShippingContact returns the NeedsShippingContact field if non-nil, zero value otherwise.

### GetNeedsShippingContactOk

`func (o *TokenResponseCheckout) GetNeedsShippingContactOk() (*bool, bool)`

GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsShippingContact

`func (o *TokenResponseCheckout) SetNeedsShippingContact(v bool)`

SetNeedsShippingContact sets NeedsShippingContact field to given value.

### HasNeedsShippingContact

`func (o *TokenResponseCheckout) HasNeedsShippingContact() bool`

HasNeedsShippingContact returns a boolean if a field has been set.

### GetObject

`func (o *TokenResponseCheckout) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TokenResponseCheckout) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TokenResponseCheckout) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TokenResponseCheckout) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOnDemandEnabled

`func (o *TokenResponseCheckout) GetOnDemandEnabled() bool`

GetOnDemandEnabled returns the OnDemandEnabled field if non-nil, zero value otherwise.

### GetOnDemandEnabledOk

`func (o *TokenResponseCheckout) GetOnDemandEnabledOk() (*bool, bool)`

GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandEnabled

`func (o *TokenResponseCheckout) SetOnDemandEnabled(v bool)`

SetOnDemandEnabled sets OnDemandEnabled field to given value.

### HasOnDemandEnabled

`func (o *TokenResponseCheckout) HasOnDemandEnabled() bool`

HasOnDemandEnabled returns a boolean if a field has been set.

### GetPaidPaymentsCount

`func (o *TokenResponseCheckout) GetPaidPaymentsCount() int32`

GetPaidPaymentsCount returns the PaidPaymentsCount field if non-nil, zero value otherwise.

### GetPaidPaymentsCountOk

`func (o *TokenResponseCheckout) GetPaidPaymentsCountOk() (*int32, bool)`

GetPaidPaymentsCountOk returns a tuple with the PaidPaymentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidPaymentsCount

`func (o *TokenResponseCheckout) SetPaidPaymentsCount(v int32)`

SetPaidPaymentsCount sets PaidPaymentsCount field to given value.

### HasPaidPaymentsCount

`func (o *TokenResponseCheckout) HasPaidPaymentsCount() bool`

HasPaidPaymentsCount returns a boolean if a field has been set.

### GetRecurrent

`func (o *TokenResponseCheckout) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *TokenResponseCheckout) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *TokenResponseCheckout) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *TokenResponseCheckout) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetSmsSent

`func (o *TokenResponseCheckout) GetSmsSent() int32`

GetSmsSent returns the SmsSent field if non-nil, zero value otherwise.

### GetSmsSentOk

`func (o *TokenResponseCheckout) GetSmsSentOk() (*int32, bool)`

GetSmsSentOk returns a tuple with the SmsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSent

`func (o *TokenResponseCheckout) SetSmsSent(v int32)`

SetSmsSent sets SmsSent field to given value.

### HasSmsSent

`func (o *TokenResponseCheckout) HasSmsSent() bool`

HasSmsSent returns a boolean if a field has been set.

### GetStartsAt

`func (o *TokenResponseCheckout) GetStartsAt() int64`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *TokenResponseCheckout) GetStartsAtOk() (*int64, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *TokenResponseCheckout) SetStartsAt(v int64)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *TokenResponseCheckout) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *TokenResponseCheckout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenResponseCheckout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenResponseCheckout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenResponseCheckout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *TokenResponseCheckout) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *TokenResponseCheckout) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *TokenResponseCheckout) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *TokenResponseCheckout) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetType

`func (o *TokenResponseCheckout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenResponseCheckout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenResponseCheckout) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenResponseCheckout) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


