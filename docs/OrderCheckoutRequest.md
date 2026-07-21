# OrderCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** | Are the payment methods available for this link. For subscriptions, only &#39;card&#39; is allowed due to the recurring nature of the payments. This field is mutually exclusive with excluded_payment_methods. | [optional] 
**ExcludedPaymentMethods** | Pointer to **[]string** | Payment methods to be excluded from the checkout. This field is mutually exclusive with allowed_payment_methods. | [optional] 
**ExcludeCardNetworks** | Pointer to **[]string** | List of card networks to exclude from the checkout. This field is only applicable for card payments. | [optional] 
**PlanIds** | Pointer to **[]string** | List of plan IDs that will be available for subscription. This field is required for subscription payments. | [optional] 
**ExpiresAt** | Pointer to **int64** | It is the time when the link will expire.  It is expressed in seconds since the Unix epoch. The valid range is from 5 minutes to 365 days from the creation date.  | [optional] 
**FailureUrl** | Pointer to **string** | Redirection url back to the site in case of failed payment, applies only to HostedPayment. | [optional] 
**ForceSaveCard** | Pointer to **bool** | Indicates whether the card used for the payment should be saved for future purchases. This field is only applicable for card payments. | [optional] 
**MonthlyInstallmentsEnabled** | Pointer to **bool** |  | [optional] 
**MonthlyInstallmentsOptions** | Pointer to **[]int32** |  | [optional] 
**MaxFailedRetries** | Pointer to **int32** | Number of retries allowed before the checkout is marked as failed | [optional] 
**Name** | Pointer to **string** | Reason for payment | [optional] 
**OnDemandEnabled** | Pointer to **bool** |  | [optional] 
**RedirectionTime** | Pointer to **int32** | number of seconds to wait before redirecting to the success_url | [optional] 
**SuccessUrl** | Pointer to **string** | Redirection url back to the site in case of successful payment, applies only to HostedPayment | [optional] 
**Type** | Pointer to **string** | Required. This field represents the type of checkout, which determines the user experience during the payment process. &#39;HostedPayment&#39; will redirect the customer to a Conekta-hosted page to complete the payment, while &#39;Integration&#39; allows the payment process to be handled entirely on your site using Conekta&#39;s APIs and SDKs. | [optional] 

## Methods

### NewOrderCheckoutRequest

`func NewOrderCheckoutRequest() *OrderCheckoutRequest`

NewOrderCheckoutRequest instantiates a new OrderCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCheckoutRequestWithDefaults

`func NewOrderCheckoutRequestWithDefaults() *OrderCheckoutRequest`

NewOrderCheckoutRequestWithDefaults instantiates a new OrderCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *OrderCheckoutRequest) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *OrderCheckoutRequest) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *OrderCheckoutRequest) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *OrderCheckoutRequest) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetExcludedPaymentMethods

`func (o *OrderCheckoutRequest) GetExcludedPaymentMethods() []string`

GetExcludedPaymentMethods returns the ExcludedPaymentMethods field if non-nil, zero value otherwise.

### GetExcludedPaymentMethodsOk

`func (o *OrderCheckoutRequest) GetExcludedPaymentMethodsOk() (*[]string, bool)`

GetExcludedPaymentMethodsOk returns a tuple with the ExcludedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPaymentMethods

`func (o *OrderCheckoutRequest) SetExcludedPaymentMethods(v []string)`

SetExcludedPaymentMethods sets ExcludedPaymentMethods field to given value.

### HasExcludedPaymentMethods

`func (o *OrderCheckoutRequest) HasExcludedPaymentMethods() bool`

HasExcludedPaymentMethods returns a boolean if a field has been set.

### GetExcludeCardNetworks

`func (o *OrderCheckoutRequest) GetExcludeCardNetworks() []string`

GetExcludeCardNetworks returns the ExcludeCardNetworks field if non-nil, zero value otherwise.

### GetExcludeCardNetworksOk

`func (o *OrderCheckoutRequest) GetExcludeCardNetworksOk() (*[]string, bool)`

GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCardNetworks

`func (o *OrderCheckoutRequest) SetExcludeCardNetworks(v []string)`

SetExcludeCardNetworks sets ExcludeCardNetworks field to given value.

### HasExcludeCardNetworks

`func (o *OrderCheckoutRequest) HasExcludeCardNetworks() bool`

HasExcludeCardNetworks returns a boolean if a field has been set.

### GetPlanIds

`func (o *OrderCheckoutRequest) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *OrderCheckoutRequest) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *OrderCheckoutRequest) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *OrderCheckoutRequest) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrderCheckoutRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrderCheckoutRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrderCheckoutRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrderCheckoutRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailureUrl

`func (o *OrderCheckoutRequest) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *OrderCheckoutRequest) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *OrderCheckoutRequest) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *OrderCheckoutRequest) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetForceSaveCard

`func (o *OrderCheckoutRequest) GetForceSaveCard() bool`

GetForceSaveCard returns the ForceSaveCard field if non-nil, zero value otherwise.

### GetForceSaveCardOk

`func (o *OrderCheckoutRequest) GetForceSaveCardOk() (*bool, bool)`

GetForceSaveCardOk returns a tuple with the ForceSaveCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSaveCard

`func (o *OrderCheckoutRequest) SetForceSaveCard(v bool)`

SetForceSaveCard sets ForceSaveCard field to given value.

### HasForceSaveCard

`func (o *OrderCheckoutRequest) HasForceSaveCard() bool`

HasForceSaveCard returns a boolean if a field has been set.

### GetMonthlyInstallmentsEnabled

`func (o *OrderCheckoutRequest) GetMonthlyInstallmentsEnabled() bool`

GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsEnabledOk

`func (o *OrderCheckoutRequest) GetMonthlyInstallmentsEnabledOk() (*bool, bool)`

GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsEnabled

`func (o *OrderCheckoutRequest) SetMonthlyInstallmentsEnabled(v bool)`

SetMonthlyInstallmentsEnabled sets MonthlyInstallmentsEnabled field to given value.

### HasMonthlyInstallmentsEnabled

`func (o *OrderCheckoutRequest) HasMonthlyInstallmentsEnabled() bool`

HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.

### GetMonthlyInstallmentsOptions

`func (o *OrderCheckoutRequest) GetMonthlyInstallmentsOptions() []int32`

GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOptionsOk

`func (o *OrderCheckoutRequest) GetMonthlyInstallmentsOptionsOk() (*[]int32, bool)`

GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsOptions

`func (o *OrderCheckoutRequest) SetMonthlyInstallmentsOptions(v []int32)`

SetMonthlyInstallmentsOptions sets MonthlyInstallmentsOptions field to given value.

### HasMonthlyInstallmentsOptions

`func (o *OrderCheckoutRequest) HasMonthlyInstallmentsOptions() bool`

HasMonthlyInstallmentsOptions returns a boolean if a field has been set.

### GetMaxFailedRetries

`func (o *OrderCheckoutRequest) GetMaxFailedRetries() int32`

GetMaxFailedRetries returns the MaxFailedRetries field if non-nil, zero value otherwise.

### GetMaxFailedRetriesOk

`func (o *OrderCheckoutRequest) GetMaxFailedRetriesOk() (*int32, bool)`

GetMaxFailedRetriesOk returns a tuple with the MaxFailedRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedRetries

`func (o *OrderCheckoutRequest) SetMaxFailedRetries(v int32)`

SetMaxFailedRetries sets MaxFailedRetries field to given value.

### HasMaxFailedRetries

`func (o *OrderCheckoutRequest) HasMaxFailedRetries() bool`

HasMaxFailedRetries returns a boolean if a field has been set.

### GetName

`func (o *OrderCheckoutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderCheckoutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderCheckoutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderCheckoutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnDemandEnabled

`func (o *OrderCheckoutRequest) GetOnDemandEnabled() bool`

GetOnDemandEnabled returns the OnDemandEnabled field if non-nil, zero value otherwise.

### GetOnDemandEnabledOk

`func (o *OrderCheckoutRequest) GetOnDemandEnabledOk() (*bool, bool)`

GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandEnabled

`func (o *OrderCheckoutRequest) SetOnDemandEnabled(v bool)`

SetOnDemandEnabled sets OnDemandEnabled field to given value.

### HasOnDemandEnabled

`func (o *OrderCheckoutRequest) HasOnDemandEnabled() bool`

HasOnDemandEnabled returns a boolean if a field has been set.

### GetRedirectionTime

`func (o *OrderCheckoutRequest) GetRedirectionTime() int32`

GetRedirectionTime returns the RedirectionTime field if non-nil, zero value otherwise.

### GetRedirectionTimeOk

`func (o *OrderCheckoutRequest) GetRedirectionTimeOk() (*int32, bool)`

GetRedirectionTimeOk returns a tuple with the RedirectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionTime

`func (o *OrderCheckoutRequest) SetRedirectionTime(v int32)`

SetRedirectionTime sets RedirectionTime field to given value.

### HasRedirectionTime

`func (o *OrderCheckoutRequest) HasRedirectionTime() bool`

HasRedirectionTime returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *OrderCheckoutRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *OrderCheckoutRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *OrderCheckoutRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *OrderCheckoutRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetType

`func (o *OrderCheckoutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCheckoutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCheckoutRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderCheckoutRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


