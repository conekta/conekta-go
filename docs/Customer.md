# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntifraudInfo** | Pointer to [**NullableCustomerAntifraudInfo**](CustomerAntifraudInfo.md) |  | [optional] 
**Corporate** | Pointer to **bool** | It is a value that allows identifying if the email is corporate or not. | [optional] [default to false]
**CustomReference** | Pointer to **string** | It is an undefined value. | [optional] 
**Email** | **string** | An email address is a series of customizable characters followed by a universal Internet symbol, the at symbol (@), the name of a host server, and a web domain ending (.mx, .com, .org, . net, etc). | 
**DefaultPaymentSourceId** | Pointer to **string** | It is a parameter that allows to identify in the response, the Conekta ID of a payment method (payment_id) | [optional] 
**DefaultShippingContactId** | Pointer to **string** | It is a parameter that allows to identify in the response, the Conekta ID of the shipping address (shipping_contact) | [optional] 
**FiscalEntities** | Pointer to [**[]CustomerFiscalEntitiesRequest**](CustomerFiscalEntitiesRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | Client&#39;s name | 
**PaymentSources** | Pointer to [**[]ConsumerPaymentMethodsRequest**](ConsumerPaymentMethodsRequest.md) | Contains details of the payment methods that the customer has active or has used in Conekta | [optional] 
**Phone** | **string** | Is the customer&#39;s phone number | 
**PlanId** | Pointer to **string** | Contains the ID of a plan, which could together with name, email and phone create a client directly to a subscription | [optional] 
**ShippingContacts** | Pointer to [**[]CustomerShippingContacts**](CustomerShippingContacts.md) | Contains the detail of the shipping addresses that the client has active or has used in Conekta | [optional] 
**Subscription** | Pointer to [**SubscriptionRequest**](SubscriptionRequest.md) |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer(email string, name string, phone string, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *Customer) GetAntifraudInfo() CustomerAntifraudInfo`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *Customer) GetAntifraudInfoOk() (*CustomerAntifraudInfo, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *Customer) SetAntifraudInfo(v CustomerAntifraudInfo)`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *Customer) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### SetAntifraudInfoNil

`func (o *Customer) SetAntifraudInfoNil(b bool)`

 SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil

### UnsetAntifraudInfo
`func (o *Customer) UnsetAntifraudInfo()`

UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
### GetCorporate

`func (o *Customer) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *Customer) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *Customer) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *Customer) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetCustomReference

`func (o *Customer) GetCustomReference() string`

GetCustomReference returns the CustomReference field if non-nil, zero value otherwise.

### GetCustomReferenceOk

`func (o *Customer) GetCustomReferenceOk() (*string, bool)`

GetCustomReferenceOk returns a tuple with the CustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReference

`func (o *Customer) SetCustomReference(v string)`

SetCustomReference sets CustomReference field to given value.

### HasCustomReference

`func (o *Customer) HasCustomReference() bool`

HasCustomReference returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDefaultPaymentSourceId

`func (o *Customer) GetDefaultPaymentSourceId() string`

GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field if non-nil, zero value otherwise.

### GetDefaultPaymentSourceIdOk

`func (o *Customer) GetDefaultPaymentSourceIdOk() (*string, bool)`

GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentSourceId

`func (o *Customer) SetDefaultPaymentSourceId(v string)`

SetDefaultPaymentSourceId sets DefaultPaymentSourceId field to given value.

### HasDefaultPaymentSourceId

`func (o *Customer) HasDefaultPaymentSourceId() bool`

HasDefaultPaymentSourceId returns a boolean if a field has been set.

### GetDefaultShippingContactId

`func (o *Customer) GetDefaultShippingContactId() string`

GetDefaultShippingContactId returns the DefaultShippingContactId field if non-nil, zero value otherwise.

### GetDefaultShippingContactIdOk

`func (o *Customer) GetDefaultShippingContactIdOk() (*string, bool)`

GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingContactId

`func (o *Customer) SetDefaultShippingContactId(v string)`

SetDefaultShippingContactId sets DefaultShippingContactId field to given value.

### HasDefaultShippingContactId

`func (o *Customer) HasDefaultShippingContactId() bool`

HasDefaultShippingContactId returns a boolean if a field has been set.

### GetFiscalEntities

`func (o *Customer) GetFiscalEntities() []CustomerFiscalEntitiesRequest`

GetFiscalEntities returns the FiscalEntities field if non-nil, zero value otherwise.

### GetFiscalEntitiesOk

`func (o *Customer) GetFiscalEntitiesOk() (*[]CustomerFiscalEntitiesRequest, bool)`

GetFiscalEntitiesOk returns a tuple with the FiscalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntities

`func (o *Customer) SetFiscalEntities(v []CustomerFiscalEntitiesRequest)`

SetFiscalEntities sets FiscalEntities field to given value.

### HasFiscalEntities

`func (o *Customer) HasFiscalEntities() bool`

HasFiscalEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *Customer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Customer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Customer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Customer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentSources

`func (o *Customer) GetPaymentSources() []ConsumerPaymentMethodsRequest`

GetPaymentSources returns the PaymentSources field if non-nil, zero value otherwise.

### GetPaymentSourcesOk

`func (o *Customer) GetPaymentSourcesOk() (*[]ConsumerPaymentMethodsRequest, bool)`

GetPaymentSourcesOk returns a tuple with the PaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSources

`func (o *Customer) SetPaymentSources(v []ConsumerPaymentMethodsRequest)`

SetPaymentSources sets PaymentSources field to given value.

### HasPaymentSources

`func (o *Customer) HasPaymentSources() bool`

HasPaymentSources returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPlanId

`func (o *Customer) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Customer) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Customer) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *Customer) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetShippingContacts

`func (o *Customer) GetShippingContacts() []CustomerShippingContacts`

GetShippingContacts returns the ShippingContacts field if non-nil, zero value otherwise.

### GetShippingContactsOk

`func (o *Customer) GetShippingContactsOk() (*[]CustomerShippingContacts, bool)`

GetShippingContactsOk returns a tuple with the ShippingContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContacts

`func (o *Customer) SetShippingContacts(v []CustomerShippingContacts)`

SetShippingContacts sets ShippingContacts field to given value.

### HasShippingContacts

`func (o *Customer) HasShippingContacts() bool`

HasShippingContacts returns a boolean if a field has been set.

### GetSubscription

`func (o *Customer) GetSubscription() SubscriptionRequest`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Customer) GetSubscriptionOk() (*SubscriptionRequest, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Customer) SetSubscription(v SubscriptionRequest)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *Customer) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


