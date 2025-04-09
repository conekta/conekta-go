# CheckoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** | Are the payment methods available for this link | [optional] 
**CanNotExpire** | Pointer to **bool** |  | [optional] 
**EmailsSent** | Pointer to **int32** |  | [optional] 
**ExcludeCardNetworks** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**FailureUrl** | Pointer to **string** |  | [optional] 
**Force3dsFlow** | Pointer to **bool** |  | [optional] 
**Id** | **string** |  | 
**Livemode** | **bool** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**MonthlyInstallmentsEnabled** | Pointer to **bool** |  | [optional] 
**MonthlyInstallmentsOptions** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** | Reason for charge | 
**NeedsShippingContact** | Pointer to **bool** |  | [optional] 
**Object** | **string** |  | 
**PaidPaymentsCount** | Pointer to **int32** |  | [optional] 
**PaymentsLimitCount** | Pointer to **NullableInt32** |  | [optional] 
**Recurrent** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**SmsSent** | Pointer to **int32** |  | [optional] 
**StartsAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SuccessUrl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckoutResponse

`func NewCheckoutResponse(id string, livemode bool, name string, object string, ) *CheckoutResponse`

NewCheckoutResponse instantiates a new CheckoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutResponseWithDefaults

`func NewCheckoutResponseWithDefaults() *CheckoutResponse`

NewCheckoutResponseWithDefaults instantiates a new CheckoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *CheckoutResponse) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *CheckoutResponse) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *CheckoutResponse) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *CheckoutResponse) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetCanNotExpire

`func (o *CheckoutResponse) GetCanNotExpire() bool`

GetCanNotExpire returns the CanNotExpire field if non-nil, zero value otherwise.

### GetCanNotExpireOk

`func (o *CheckoutResponse) GetCanNotExpireOk() (*bool, bool)`

GetCanNotExpireOk returns a tuple with the CanNotExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanNotExpire

`func (o *CheckoutResponse) SetCanNotExpire(v bool)`

SetCanNotExpire sets CanNotExpire field to given value.

### HasCanNotExpire

`func (o *CheckoutResponse) HasCanNotExpire() bool`

HasCanNotExpire returns a boolean if a field has been set.

### GetEmailsSent

`func (o *CheckoutResponse) GetEmailsSent() int32`

GetEmailsSent returns the EmailsSent field if non-nil, zero value otherwise.

### GetEmailsSentOk

`func (o *CheckoutResponse) GetEmailsSentOk() (*int32, bool)`

GetEmailsSentOk returns a tuple with the EmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsSent

`func (o *CheckoutResponse) SetEmailsSent(v int32)`

SetEmailsSent sets EmailsSent field to given value.

### HasEmailsSent

`func (o *CheckoutResponse) HasEmailsSent() bool`

HasEmailsSent returns a boolean if a field has been set.

### GetExcludeCardNetworks

`func (o *CheckoutResponse) GetExcludeCardNetworks() []map[string]interface{}`

GetExcludeCardNetworks returns the ExcludeCardNetworks field if non-nil, zero value otherwise.

### GetExcludeCardNetworksOk

`func (o *CheckoutResponse) GetExcludeCardNetworksOk() (*[]map[string]interface{}, bool)`

GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCardNetworks

`func (o *CheckoutResponse) SetExcludeCardNetworks(v []map[string]interface{})`

SetExcludeCardNetworks sets ExcludeCardNetworks field to given value.

### HasExcludeCardNetworks

`func (o *CheckoutResponse) HasExcludeCardNetworks() bool`

HasExcludeCardNetworks returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailureUrl

`func (o *CheckoutResponse) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *CheckoutResponse) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *CheckoutResponse) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *CheckoutResponse) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetForce3dsFlow

`func (o *CheckoutResponse) GetForce3dsFlow() bool`

GetForce3dsFlow returns the Force3dsFlow field if non-nil, zero value otherwise.

### GetForce3dsFlowOk

`func (o *CheckoutResponse) GetForce3dsFlowOk() (*bool, bool)`

GetForce3dsFlowOk returns a tuple with the Force3dsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce3dsFlow

`func (o *CheckoutResponse) SetForce3dsFlow(v bool)`

SetForce3dsFlow sets Force3dsFlow field to given value.

### HasForce3dsFlow

`func (o *CheckoutResponse) HasForce3dsFlow() bool`

HasForce3dsFlow returns a boolean if a field has been set.

### GetId

`func (o *CheckoutResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLivemode

`func (o *CheckoutResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *CheckoutResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *CheckoutResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetMetadata

`func (o *CheckoutResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonthlyInstallmentsEnabled

`func (o *CheckoutResponse) GetMonthlyInstallmentsEnabled() bool`

GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsEnabledOk

`func (o *CheckoutResponse) GetMonthlyInstallmentsEnabledOk() (*bool, bool)`

GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsEnabled

`func (o *CheckoutResponse) SetMonthlyInstallmentsEnabled(v bool)`

SetMonthlyInstallmentsEnabled sets MonthlyInstallmentsEnabled field to given value.

### HasMonthlyInstallmentsEnabled

`func (o *CheckoutResponse) HasMonthlyInstallmentsEnabled() bool`

HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.

### GetMonthlyInstallmentsOptions

`func (o *CheckoutResponse) GetMonthlyInstallmentsOptions() []int32`

GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field if non-nil, zero value otherwise.

### GetMonthlyInstallmentsOptionsOk

`func (o *CheckoutResponse) GetMonthlyInstallmentsOptionsOk() (*[]int32, bool)`

GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyInstallmentsOptions

`func (o *CheckoutResponse) SetMonthlyInstallmentsOptions(v []int32)`

SetMonthlyInstallmentsOptions sets MonthlyInstallmentsOptions field to given value.

### HasMonthlyInstallmentsOptions

`func (o *CheckoutResponse) HasMonthlyInstallmentsOptions() bool`

HasMonthlyInstallmentsOptions returns a boolean if a field has been set.

### GetName

`func (o *CheckoutResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNeedsShippingContact

`func (o *CheckoutResponse) GetNeedsShippingContact() bool`

GetNeedsShippingContact returns the NeedsShippingContact field if non-nil, zero value otherwise.

### GetNeedsShippingContactOk

`func (o *CheckoutResponse) GetNeedsShippingContactOk() (*bool, bool)`

GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsShippingContact

`func (o *CheckoutResponse) SetNeedsShippingContact(v bool)`

SetNeedsShippingContact sets NeedsShippingContact field to given value.

### HasNeedsShippingContact

`func (o *CheckoutResponse) HasNeedsShippingContact() bool`

HasNeedsShippingContact returns a boolean if a field has been set.

### GetObject

`func (o *CheckoutResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CheckoutResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CheckoutResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetPaidPaymentsCount

`func (o *CheckoutResponse) GetPaidPaymentsCount() int32`

GetPaidPaymentsCount returns the PaidPaymentsCount field if non-nil, zero value otherwise.

### GetPaidPaymentsCountOk

`func (o *CheckoutResponse) GetPaidPaymentsCountOk() (*int32, bool)`

GetPaidPaymentsCountOk returns a tuple with the PaidPaymentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidPaymentsCount

`func (o *CheckoutResponse) SetPaidPaymentsCount(v int32)`

SetPaidPaymentsCount sets PaidPaymentsCount field to given value.

### HasPaidPaymentsCount

`func (o *CheckoutResponse) HasPaidPaymentsCount() bool`

HasPaidPaymentsCount returns a boolean if a field has been set.

### GetPaymentsLimitCount

`func (o *CheckoutResponse) GetPaymentsLimitCount() int32`

GetPaymentsLimitCount returns the PaymentsLimitCount field if non-nil, zero value otherwise.

### GetPaymentsLimitCountOk

`func (o *CheckoutResponse) GetPaymentsLimitCountOk() (*int32, bool)`

GetPaymentsLimitCountOk returns a tuple with the PaymentsLimitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsLimitCount

`func (o *CheckoutResponse) SetPaymentsLimitCount(v int32)`

SetPaymentsLimitCount sets PaymentsLimitCount field to given value.

### HasPaymentsLimitCount

`func (o *CheckoutResponse) HasPaymentsLimitCount() bool`

HasPaymentsLimitCount returns a boolean if a field has been set.

### SetPaymentsLimitCountNil

`func (o *CheckoutResponse) SetPaymentsLimitCountNil(b bool)`

 SetPaymentsLimitCountNil sets the value for PaymentsLimitCount to be an explicit nil

### UnsetPaymentsLimitCount
`func (o *CheckoutResponse) UnsetPaymentsLimitCount()`

UnsetPaymentsLimitCount ensures that no value is present for PaymentsLimitCount, not even an explicit nil
### GetRecurrent

`func (o *CheckoutResponse) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *CheckoutResponse) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *CheckoutResponse) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *CheckoutResponse) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetSlug

`func (o *CheckoutResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CheckoutResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CheckoutResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CheckoutResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSmsSent

`func (o *CheckoutResponse) GetSmsSent() int32`

GetSmsSent returns the SmsSent field if non-nil, zero value otherwise.

### GetSmsSentOk

`func (o *CheckoutResponse) GetSmsSentOk() (*int32, bool)`

GetSmsSentOk returns a tuple with the SmsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSent

`func (o *CheckoutResponse) SetSmsSent(v int32)`

SetSmsSent sets SmsSent field to given value.

### HasSmsSent

`func (o *CheckoutResponse) HasSmsSent() bool`

HasSmsSent returns a boolean if a field has been set.

### GetStartsAt

`func (o *CheckoutResponse) GetStartsAt() int32`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CheckoutResponse) GetStartsAtOk() (*int32, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CheckoutResponse) SetStartsAt(v int32)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *CheckoutResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *CheckoutResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckoutResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CheckoutResponse) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutResponse) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutResponse) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CheckoutResponse) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetType

`func (o *CheckoutResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *CheckoutResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CheckoutResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


