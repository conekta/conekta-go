# CustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntifraudInfo** | Pointer to [**NullableCustomerAntifraudInfoResponse**](CustomerAntifraudInfoResponse.md) |  | [optional] 
**Corporate** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **int64** |  | 
**CustomReference** | Pointer to **string** |  | [optional] 
**DefaultFiscalEntityId** | Pointer to **NullableString** |  | [optional] 
**DefaultShippingContactId** | Pointer to **string** |  | [optional] 
**DefaultPaymentSourceId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FiscalEntities** | Pointer to [**CustomerFiscalEntitiesResponse**](CustomerFiscalEntitiesResponse.md) |  | [optional] 
**Id** | **string** |  | 
**Livemode** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**PaymentSources** | Pointer to [**CustomerPaymentMethodsResponse**](CustomerPaymentMethodsResponse.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**ShippingContacts** | Pointer to [**CustomerResponseShippingContacts**](CustomerResponseShippingContacts.md) |  | [optional] 
**Subscription** | Pointer to [**SubscriptionResponse**](SubscriptionResponse.md) |  | [optional] 

## Methods

### NewCustomerResponse

`func NewCustomerResponse(createdAt int64, id string, livemode bool, object string, ) *CustomerResponse`

NewCustomerResponse instantiates a new CustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseWithDefaults

`func NewCustomerResponseWithDefaults() *CustomerResponse`

NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *CustomerResponse) GetAntifraudInfo() CustomerAntifraudInfoResponse`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *CustomerResponse) GetAntifraudInfoOk() (*CustomerAntifraudInfoResponse, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *CustomerResponse) SetAntifraudInfo(v CustomerAntifraudInfoResponse)`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *CustomerResponse) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### SetAntifraudInfoNil

`func (o *CustomerResponse) SetAntifraudInfoNil(b bool)`

 SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil

### UnsetAntifraudInfo
`func (o *CustomerResponse) UnsetAntifraudInfo()`

UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
### GetCorporate

`func (o *CustomerResponse) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *CustomerResponse) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *CustomerResponse) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *CustomerResponse) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomReference

`func (o *CustomerResponse) GetCustomReference() string`

GetCustomReference returns the CustomReference field if non-nil, zero value otherwise.

### GetCustomReferenceOk

`func (o *CustomerResponse) GetCustomReferenceOk() (*string, bool)`

GetCustomReferenceOk returns a tuple with the CustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReference

`func (o *CustomerResponse) SetCustomReference(v string)`

SetCustomReference sets CustomReference field to given value.

### HasCustomReference

`func (o *CustomerResponse) HasCustomReference() bool`

HasCustomReference returns a boolean if a field has been set.

### GetDefaultFiscalEntityId

`func (o *CustomerResponse) GetDefaultFiscalEntityId() string`

GetDefaultFiscalEntityId returns the DefaultFiscalEntityId field if non-nil, zero value otherwise.

### GetDefaultFiscalEntityIdOk

`func (o *CustomerResponse) GetDefaultFiscalEntityIdOk() (*string, bool)`

GetDefaultFiscalEntityIdOk returns a tuple with the DefaultFiscalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFiscalEntityId

`func (o *CustomerResponse) SetDefaultFiscalEntityId(v string)`

SetDefaultFiscalEntityId sets DefaultFiscalEntityId field to given value.

### HasDefaultFiscalEntityId

`func (o *CustomerResponse) HasDefaultFiscalEntityId() bool`

HasDefaultFiscalEntityId returns a boolean if a field has been set.

### SetDefaultFiscalEntityIdNil

`func (o *CustomerResponse) SetDefaultFiscalEntityIdNil(b bool)`

 SetDefaultFiscalEntityIdNil sets the value for DefaultFiscalEntityId to be an explicit nil

### UnsetDefaultFiscalEntityId
`func (o *CustomerResponse) UnsetDefaultFiscalEntityId()`

UnsetDefaultFiscalEntityId ensures that no value is present for DefaultFiscalEntityId, not even an explicit nil
### GetDefaultShippingContactId

`func (o *CustomerResponse) GetDefaultShippingContactId() string`

GetDefaultShippingContactId returns the DefaultShippingContactId field if non-nil, zero value otherwise.

### GetDefaultShippingContactIdOk

`func (o *CustomerResponse) GetDefaultShippingContactIdOk() (*string, bool)`

GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingContactId

`func (o *CustomerResponse) SetDefaultShippingContactId(v string)`

SetDefaultShippingContactId sets DefaultShippingContactId field to given value.

### HasDefaultShippingContactId

`func (o *CustomerResponse) HasDefaultShippingContactId() bool`

HasDefaultShippingContactId returns a boolean if a field has been set.

### GetDefaultPaymentSourceId

`func (o *CustomerResponse) GetDefaultPaymentSourceId() string`

GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field if non-nil, zero value otherwise.

### GetDefaultPaymentSourceIdOk

`func (o *CustomerResponse) GetDefaultPaymentSourceIdOk() (*string, bool)`

GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentSourceId

`func (o *CustomerResponse) SetDefaultPaymentSourceId(v string)`

SetDefaultPaymentSourceId sets DefaultPaymentSourceId field to given value.

### HasDefaultPaymentSourceId

`func (o *CustomerResponse) HasDefaultPaymentSourceId() bool`

HasDefaultPaymentSourceId returns a boolean if a field has been set.

### SetDefaultPaymentSourceIdNil

`func (o *CustomerResponse) SetDefaultPaymentSourceIdNil(b bool)`

 SetDefaultPaymentSourceIdNil sets the value for DefaultPaymentSourceId to be an explicit nil

### UnsetDefaultPaymentSourceId
`func (o *CustomerResponse) UnsetDefaultPaymentSourceId()`

UnsetDefaultPaymentSourceId ensures that no value is present for DefaultPaymentSourceId, not even an explicit nil
### GetEmail

`func (o *CustomerResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFiscalEntities

`func (o *CustomerResponse) GetFiscalEntities() CustomerFiscalEntitiesResponse`

GetFiscalEntities returns the FiscalEntities field if non-nil, zero value otherwise.

### GetFiscalEntitiesOk

`func (o *CustomerResponse) GetFiscalEntitiesOk() (*CustomerFiscalEntitiesResponse, bool)`

GetFiscalEntitiesOk returns a tuple with the FiscalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntities

`func (o *CustomerResponse) SetFiscalEntities(v CustomerFiscalEntitiesResponse)`

SetFiscalEntities sets FiscalEntities field to given value.

### HasFiscalEntities

`func (o *CustomerResponse) HasFiscalEntities() bool`

HasFiscalEntities returns a boolean if a field has been set.

### GetId

`func (o *CustomerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLivemode

`func (o *CustomerResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *CustomerResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *CustomerResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetName

`func (o *CustomerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *CustomerResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetPaymentSources

`func (o *CustomerResponse) GetPaymentSources() CustomerPaymentMethodsResponse`

GetPaymentSources returns the PaymentSources field if non-nil, zero value otherwise.

### GetPaymentSourcesOk

`func (o *CustomerResponse) GetPaymentSourcesOk() (*CustomerPaymentMethodsResponse, bool)`

GetPaymentSourcesOk returns a tuple with the PaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSources

`func (o *CustomerResponse) SetPaymentSources(v CustomerPaymentMethodsResponse)`

SetPaymentSources sets PaymentSources field to given value.

### HasPaymentSources

`func (o *CustomerResponse) HasPaymentSources() bool`

HasPaymentSources returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetShippingContacts

`func (o *CustomerResponse) GetShippingContacts() CustomerResponseShippingContacts`

GetShippingContacts returns the ShippingContacts field if non-nil, zero value otherwise.

### GetShippingContactsOk

`func (o *CustomerResponse) GetShippingContactsOk() (*CustomerResponseShippingContacts, bool)`

GetShippingContactsOk returns a tuple with the ShippingContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContacts

`func (o *CustomerResponse) SetShippingContacts(v CustomerResponseShippingContacts)`

SetShippingContacts sets ShippingContacts field to given value.

### HasShippingContacts

`func (o *CustomerResponse) HasShippingContacts() bool`

HasShippingContacts returns a boolean if a field has been set.

### GetSubscription

`func (o *CustomerResponse) GetSubscription() SubscriptionResponse`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CustomerResponse) GetSubscriptionOk() (*SubscriptionResponse, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CustomerResponse) SetSubscription(v SubscriptionResponse)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CustomerResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


