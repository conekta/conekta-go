# PaymentMethodCashResponseAllOfAgreements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to **string** | Agreement number, you can use this number to pay in the store/bbva | [optional] 
**Provider** | Pointer to **string** | Provider name, you can use this to know where to pay | [optional] 

## Methods

### NewPaymentMethodCashResponseAllOfAgreements

`func NewPaymentMethodCashResponseAllOfAgreements() *PaymentMethodCashResponseAllOfAgreements`

NewPaymentMethodCashResponseAllOfAgreements instantiates a new PaymentMethodCashResponseAllOfAgreements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCashResponseAllOfAgreementsWithDefaults

`func NewPaymentMethodCashResponseAllOfAgreementsWithDefaults() *PaymentMethodCashResponseAllOfAgreements`

NewPaymentMethodCashResponseAllOfAgreementsWithDefaults instantiates a new PaymentMethodCashResponseAllOfAgreements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *PaymentMethodCashResponseAllOfAgreements) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PaymentMethodCashResponseAllOfAgreements) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PaymentMethodCashResponseAllOfAgreements) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PaymentMethodCashResponseAllOfAgreements) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentMethodCashResponseAllOfAgreements) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethodCashResponseAllOfAgreements) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethodCashResponseAllOfAgreements) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentMethodCashResponseAllOfAgreements) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


