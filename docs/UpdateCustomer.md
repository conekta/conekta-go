# UpdateCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntifraudInfo** | Pointer to [**NullableUpdateCustomerAntifraudInfo**](UpdateCustomerAntifraudInfo.md) |  | [optional] 
**DefaultPaymentSourceId** | Pointer to **string** | It is a parameter that allows to identify in the response, the Conekta ID of a payment method (payment_id) | [optional] 
**Email** | Pointer to **string** | An email address is a series of customizable characters followed by a universal Internet symbol, the at symbol (@), the name of a host server, and a web domain ending (.mx, .com, .org, . net, etc). | [optional] 
**Name** | Pointer to **string** | Client&#39;s name | [optional] 
**Phone** | Pointer to **string** | Is the customer&#39;s phone number | [optional] 
**PlanId** | Pointer to **string** | Contains the ID of a plan, which could together with name, email and phone create a client directly to a subscription | [optional] 
**DefaultShippingContactId** | Pointer to **string** | It is a parameter that allows to identify in the response, the Conekta ID of the shipping address (shipping_contact) | [optional] 
**Corporate** | Pointer to **bool** | It is a value that allows identifying if the email is corporate or not. | [optional] [default to false]
**CustomReference** | Pointer to **string** | It is an undefined value. | [optional] 
**FiscalEntities** | Pointer to [**[]CustomerFiscalEntitiesRequest**](CustomerFiscalEntitiesRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**PaymentSources** | Pointer to [**[]CustomerPaymentMethodsRequest**](CustomerPaymentMethodsRequest.md) | Contains details of the payment methods that the customer has active or has used in Conekta | [optional] 
**ShippingContacts** | Pointer to [**[]CustomerShippingContacts**](CustomerShippingContacts.md) | Contains the detail of the shipping addresses that the client has active or has used in Conekta | [optional] 
**Subscription** | Pointer to [**SubscriptionRequest**](SubscriptionRequest.md) |  | [optional] 

## Methods

### NewUpdateCustomer

`func NewUpdateCustomer() *UpdateCustomer`

NewUpdateCustomer instantiates a new UpdateCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerWithDefaults

`func NewUpdateCustomerWithDefaults() *UpdateCustomer`

NewUpdateCustomerWithDefaults instantiates a new UpdateCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *UpdateCustomer) GetAntifraudInfo() UpdateCustomerAntifraudInfo`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *UpdateCustomer) GetAntifraudInfoOk() (*UpdateCustomerAntifraudInfo, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *UpdateCustomer) SetAntifraudInfo(v UpdateCustomerAntifraudInfo)`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *UpdateCustomer) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### SetAntifraudInfoNil

`func (o *UpdateCustomer) SetAntifraudInfoNil(b bool)`

 SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil

### UnsetAntifraudInfo
`func (o *UpdateCustomer) UnsetAntifraudInfo()`

UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
### GetDefaultPaymentSourceId

`func (o *UpdateCustomer) GetDefaultPaymentSourceId() string`

GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field if non-nil, zero value otherwise.

### GetDefaultPaymentSourceIdOk

`func (o *UpdateCustomer) GetDefaultPaymentSourceIdOk() (*string, bool)`

GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentSourceId

`func (o *UpdateCustomer) SetDefaultPaymentSourceId(v string)`

SetDefaultPaymentSourceId sets DefaultPaymentSourceId field to given value.

### HasDefaultPaymentSourceId

`func (o *UpdateCustomer) HasDefaultPaymentSourceId() bool`

HasDefaultPaymentSourceId returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateCustomer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateCustomer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateCustomer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateCustomer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPlanId

`func (o *UpdateCustomer) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UpdateCustomer) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UpdateCustomer) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UpdateCustomer) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetDefaultShippingContactId

`func (o *UpdateCustomer) GetDefaultShippingContactId() string`

GetDefaultShippingContactId returns the DefaultShippingContactId field if non-nil, zero value otherwise.

### GetDefaultShippingContactIdOk

`func (o *UpdateCustomer) GetDefaultShippingContactIdOk() (*string, bool)`

GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingContactId

`func (o *UpdateCustomer) SetDefaultShippingContactId(v string)`

SetDefaultShippingContactId sets DefaultShippingContactId field to given value.

### HasDefaultShippingContactId

`func (o *UpdateCustomer) HasDefaultShippingContactId() bool`

HasDefaultShippingContactId returns a boolean if a field has been set.

### GetCorporate

`func (o *UpdateCustomer) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *UpdateCustomer) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *UpdateCustomer) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *UpdateCustomer) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetCustomReference

`func (o *UpdateCustomer) GetCustomReference() string`

GetCustomReference returns the CustomReference field if non-nil, zero value otherwise.

### GetCustomReferenceOk

`func (o *UpdateCustomer) GetCustomReferenceOk() (*string, bool)`

GetCustomReferenceOk returns a tuple with the CustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReference

`func (o *UpdateCustomer) SetCustomReference(v string)`

SetCustomReference sets CustomReference field to given value.

### HasCustomReference

`func (o *UpdateCustomer) HasCustomReference() bool`

HasCustomReference returns a boolean if a field has been set.

### GetFiscalEntities

`func (o *UpdateCustomer) GetFiscalEntities() []CustomerFiscalEntitiesRequest`

GetFiscalEntities returns the FiscalEntities field if non-nil, zero value otherwise.

### GetFiscalEntitiesOk

`func (o *UpdateCustomer) GetFiscalEntitiesOk() (*[]CustomerFiscalEntitiesRequest, bool)`

GetFiscalEntitiesOk returns a tuple with the FiscalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntities

`func (o *UpdateCustomer) SetFiscalEntities(v []CustomerFiscalEntitiesRequest)`

SetFiscalEntities sets FiscalEntities field to given value.

### HasFiscalEntities

`func (o *UpdateCustomer) HasFiscalEntities() bool`

HasFiscalEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateCustomer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateCustomer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateCustomer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentSources

`func (o *UpdateCustomer) GetPaymentSources() []CustomerPaymentMethodsRequest`

GetPaymentSources returns the PaymentSources field if non-nil, zero value otherwise.

### GetPaymentSourcesOk

`func (o *UpdateCustomer) GetPaymentSourcesOk() (*[]CustomerPaymentMethodsRequest, bool)`

GetPaymentSourcesOk returns a tuple with the PaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSources

`func (o *UpdateCustomer) SetPaymentSources(v []CustomerPaymentMethodsRequest)`

SetPaymentSources sets PaymentSources field to given value.

### HasPaymentSources

`func (o *UpdateCustomer) HasPaymentSources() bool`

HasPaymentSources returns a boolean if a field has been set.

### GetShippingContacts

`func (o *UpdateCustomer) GetShippingContacts() []CustomerShippingContacts`

GetShippingContacts returns the ShippingContacts field if non-nil, zero value otherwise.

### GetShippingContactsOk

`func (o *UpdateCustomer) GetShippingContactsOk() (*[]CustomerShippingContacts, bool)`

GetShippingContactsOk returns a tuple with the ShippingContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContacts

`func (o *UpdateCustomer) SetShippingContacts(v []CustomerShippingContacts)`

SetShippingContacts sets ShippingContacts field to given value.

### HasShippingContacts

`func (o *UpdateCustomer) HasShippingContacts() bool`

HasShippingContacts returns a boolean if a field has been set.

### GetSubscription

`func (o *UpdateCustomer) GetSubscription() SubscriptionRequest`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UpdateCustomer) GetSubscriptionOk() (*SubscriptionRequest, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UpdateCustomer) SetSubscription(v SubscriptionRequest)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UpdateCustomer) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


