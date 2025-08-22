# ChargeResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**Agreement** | Pointer to **string** | Agreement ID | [optional] 
**AuthCode** | Pointer to **string** |  | [optional] 
**CashierId** | Pointer to **NullableString** |  | [optional] 
**Reference** | **string** | Reference for the payment | 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | **int64** | Expiration date of the charge | 
**ProductType** | **string** | Product type of the charge | 
**ServiceName** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **string** |  | [optional] 
**CustomerIpAddress** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** | Account type of the card | [optional] 
**Brand** | Pointer to **string** | Brand of the card | [optional] 
**ContractId** | Pointer to **string** | Id sent for recurrent charges. | [optional] 
**Country** | Pointer to **string** | Country of the card | [optional] 
**ExpMonth** | Pointer to **string** | Expiration month of the card | [optional] 
**ExpYear** | Pointer to **string** | Expiration year of the card | [optional] 
**FraudIndicators** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Issuer** | Pointer to **string** | Issuer of the card | [optional] 
**Last4** | Pointer to **string** | Last 4 digits of the card | [optional] 
**Name** | Pointer to **string** | Name of the cardholder | [optional] 
**Bank** | Pointer to **string** |  | [optional] 
**Clabe** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExecutedAt** | Pointer to **NullableString** |  | [optional] 
**IssuingAccountBank** | Pointer to **NullableString** |  | [optional] 
**IssuingAccountNumber** | Pointer to **NullableString** |  | [optional] 
**IssuingAccountHolderName** | Pointer to **NullableString** |  | [optional] 
**IssuingAccountTaxId** | Pointer to **NullableString** |  | [optional] 
**PaymentAttempts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ReceivingAccountHolderName** | Pointer to **NullableString** |  | [optional] 
**ReceivingAccountNumber** | Pointer to **string** |  | [optional] 
**ReceivingAccountBank** | Pointer to **string** |  | [optional] 
**ReceivingAccountTaxId** | Pointer to **NullableString** |  | [optional] 
**ReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**TrackingCode** | Pointer to **NullableString** |  | [optional] 
**CancelUrl** | Pointer to **string** | URL to redirect the customer after a canceled payment | [optional] 
**FailureUrl** | Pointer to **string** | URL to redirect the customer after a failed payment | [optional] 
**RedirectUrl** | **string** | URL to redirect the customer to complete the payment | 
**SuccessUrl** | Pointer to **string** | URL to redirect the customer after a successful payment | [optional] 
**DeepLink** | **string** | Deep link for the payment, use for mobile apps/flows | 

## Methods

### NewChargeResponsePaymentMethod

`func NewChargeResponsePaymentMethod(object string, reference string, expiresAt int64, productType string, redirectUrl string, deepLink string, ) *ChargeResponsePaymentMethod`

NewChargeResponsePaymentMethod instantiates a new ChargeResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeResponsePaymentMethodWithDefaults

`func NewChargeResponsePaymentMethodWithDefaults() *ChargeResponsePaymentMethod`

NewChargeResponsePaymentMethodWithDefaults instantiates a new ChargeResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargeResponsePaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeResponsePaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeResponsePaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargeResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *ChargeResponsePaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeResponsePaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeResponsePaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.


### GetAgreement

`func (o *ChargeResponsePaymentMethod) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ChargeResponsePaymentMethod) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ChargeResponsePaymentMethod) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ChargeResponsePaymentMethod) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAuthCode

`func (o *ChargeResponsePaymentMethod) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeResponsePaymentMethod) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeResponsePaymentMethod) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeResponsePaymentMethod) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCashierId

`func (o *ChargeResponsePaymentMethod) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *ChargeResponsePaymentMethod) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *ChargeResponsePaymentMethod) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *ChargeResponsePaymentMethod) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *ChargeResponsePaymentMethod) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *ChargeResponsePaymentMethod) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *ChargeResponsePaymentMethod) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ChargeResponsePaymentMethod) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ChargeResponsePaymentMethod) SetReference(v string)`

SetReference sets Reference field to given value.


### GetBarcodeUrl

`func (o *ChargeResponsePaymentMethod) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *ChargeResponsePaymentMethod) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *ChargeResponsePaymentMethod) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *ChargeResponsePaymentMethod) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ChargeResponsePaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeResponsePaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeResponsePaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetProductType

`func (o *ChargeResponsePaymentMethod) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ChargeResponsePaymentMethod) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ChargeResponsePaymentMethod) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetServiceName

`func (o *ChargeResponsePaymentMethod) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ChargeResponsePaymentMethod) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ChargeResponsePaymentMethod) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ChargeResponsePaymentMethod) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *ChargeResponsePaymentMethod) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChargeResponsePaymentMethod) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChargeResponsePaymentMethod) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChargeResponsePaymentMethod) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ChargeResponsePaymentMethod) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ChargeResponsePaymentMethod) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *ChargeResponsePaymentMethod) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ChargeResponsePaymentMethod) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ChargeResponsePaymentMethod) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *ChargeResponsePaymentMethod) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetCustomerIpAddress

`func (o *ChargeResponsePaymentMethod) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *ChargeResponsePaymentMethod) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *ChargeResponsePaymentMethod) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *ChargeResponsePaymentMethod) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.

### GetAccountType

`func (o *ChargeResponsePaymentMethod) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ChargeResponsePaymentMethod) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ChargeResponsePaymentMethod) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ChargeResponsePaymentMethod) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBrand

`func (o *ChargeResponsePaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ChargeResponsePaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ChargeResponsePaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ChargeResponsePaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetContractId

`func (o *ChargeResponsePaymentMethod) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ChargeResponsePaymentMethod) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ChargeResponsePaymentMethod) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ChargeResponsePaymentMethod) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCountry

`func (o *ChargeResponsePaymentMethod) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ChargeResponsePaymentMethod) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ChargeResponsePaymentMethod) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ChargeResponsePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpMonth

`func (o *ChargeResponsePaymentMethod) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *ChargeResponsePaymentMethod) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *ChargeResponsePaymentMethod) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *ChargeResponsePaymentMethod) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *ChargeResponsePaymentMethod) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *ChargeResponsePaymentMethod) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *ChargeResponsePaymentMethod) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *ChargeResponsePaymentMethod) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFraudIndicators

`func (o *ChargeResponsePaymentMethod) GetFraudIndicators() []map[string]interface{}`

GetFraudIndicators returns the FraudIndicators field if non-nil, zero value otherwise.

### GetFraudIndicatorsOk

`func (o *ChargeResponsePaymentMethod) GetFraudIndicatorsOk() (*[]map[string]interface{}, bool)`

GetFraudIndicatorsOk returns a tuple with the FraudIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudIndicators

`func (o *ChargeResponsePaymentMethod) SetFraudIndicators(v []map[string]interface{})`

SetFraudIndicators sets FraudIndicators field to given value.

### HasFraudIndicators

`func (o *ChargeResponsePaymentMethod) HasFraudIndicators() bool`

HasFraudIndicators returns a boolean if a field has been set.

### GetIssuer

`func (o *ChargeResponsePaymentMethod) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ChargeResponsePaymentMethod) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ChargeResponsePaymentMethod) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ChargeResponsePaymentMethod) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLast4

`func (o *ChargeResponsePaymentMethod) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *ChargeResponsePaymentMethod) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *ChargeResponsePaymentMethod) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *ChargeResponsePaymentMethod) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *ChargeResponsePaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeResponsePaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeResponsePaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChargeResponsePaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBank

`func (o *ChargeResponsePaymentMethod) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *ChargeResponsePaymentMethod) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *ChargeResponsePaymentMethod) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *ChargeResponsePaymentMethod) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetClabe

`func (o *ChargeResponsePaymentMethod) GetClabe() string`

GetClabe returns the Clabe field if non-nil, zero value otherwise.

### GetClabeOk

`func (o *ChargeResponsePaymentMethod) GetClabeOk() (*string, bool)`

GetClabeOk returns a tuple with the Clabe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClabe

`func (o *ChargeResponsePaymentMethod) SetClabe(v string)`

SetClabe sets Clabe field to given value.

### HasClabe

`func (o *ChargeResponsePaymentMethod) HasClabe() bool`

HasClabe returns a boolean if a field has been set.

### GetDescription

`func (o *ChargeResponsePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeResponsePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeResponsePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargeResponsePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ChargeResponsePaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ChargeResponsePaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExecutedAt

`func (o *ChargeResponsePaymentMethod) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ChargeResponsePaymentMethod) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ChargeResponsePaymentMethod) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *ChargeResponsePaymentMethod) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### SetExecutedAtNil

`func (o *ChargeResponsePaymentMethod) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *ChargeResponsePaymentMethod) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
### GetIssuingAccountBank

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountBank() string`

GetIssuingAccountBank returns the IssuingAccountBank field if non-nil, zero value otherwise.

### GetIssuingAccountBankOk

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountBankOk() (*string, bool)`

GetIssuingAccountBankOk returns a tuple with the IssuingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountBank

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountBank(v string)`

SetIssuingAccountBank sets IssuingAccountBank field to given value.

### HasIssuingAccountBank

`func (o *ChargeResponsePaymentMethod) HasIssuingAccountBank() bool`

HasIssuingAccountBank returns a boolean if a field has been set.

### SetIssuingAccountBankNil

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountBankNil(b bool)`

 SetIssuingAccountBankNil sets the value for IssuingAccountBank to be an explicit nil

### UnsetIssuingAccountBank
`func (o *ChargeResponsePaymentMethod) UnsetIssuingAccountBank()`

UnsetIssuingAccountBank ensures that no value is present for IssuingAccountBank, not even an explicit nil
### GetIssuingAccountNumber

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountNumber() string`

GetIssuingAccountNumber returns the IssuingAccountNumber field if non-nil, zero value otherwise.

### GetIssuingAccountNumberOk

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountNumberOk() (*string, bool)`

GetIssuingAccountNumberOk returns a tuple with the IssuingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountNumber

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountNumber(v string)`

SetIssuingAccountNumber sets IssuingAccountNumber field to given value.

### HasIssuingAccountNumber

`func (o *ChargeResponsePaymentMethod) HasIssuingAccountNumber() bool`

HasIssuingAccountNumber returns a boolean if a field has been set.

### SetIssuingAccountNumberNil

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountNumberNil(b bool)`

 SetIssuingAccountNumberNil sets the value for IssuingAccountNumber to be an explicit nil

### UnsetIssuingAccountNumber
`func (o *ChargeResponsePaymentMethod) UnsetIssuingAccountNumber()`

UnsetIssuingAccountNumber ensures that no value is present for IssuingAccountNumber, not even an explicit nil
### GetIssuingAccountHolderName

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountHolderName() string`

GetIssuingAccountHolderName returns the IssuingAccountHolderName field if non-nil, zero value otherwise.

### GetIssuingAccountHolderNameOk

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountHolderNameOk() (*string, bool)`

GetIssuingAccountHolderNameOk returns a tuple with the IssuingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountHolderName

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountHolderName(v string)`

SetIssuingAccountHolderName sets IssuingAccountHolderName field to given value.

### HasIssuingAccountHolderName

`func (o *ChargeResponsePaymentMethod) HasIssuingAccountHolderName() bool`

HasIssuingAccountHolderName returns a boolean if a field has been set.

### SetIssuingAccountHolderNameNil

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountHolderNameNil(b bool)`

 SetIssuingAccountHolderNameNil sets the value for IssuingAccountHolderName to be an explicit nil

### UnsetIssuingAccountHolderName
`func (o *ChargeResponsePaymentMethod) UnsetIssuingAccountHolderName()`

UnsetIssuingAccountHolderName ensures that no value is present for IssuingAccountHolderName, not even an explicit nil
### GetIssuingAccountTaxId

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountTaxId() string`

GetIssuingAccountTaxId returns the IssuingAccountTaxId field if non-nil, zero value otherwise.

### GetIssuingAccountTaxIdOk

`func (o *ChargeResponsePaymentMethod) GetIssuingAccountTaxIdOk() (*string, bool)`

GetIssuingAccountTaxIdOk returns a tuple with the IssuingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountTaxId

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountTaxId(v string)`

SetIssuingAccountTaxId sets IssuingAccountTaxId field to given value.

### HasIssuingAccountTaxId

`func (o *ChargeResponsePaymentMethod) HasIssuingAccountTaxId() bool`

HasIssuingAccountTaxId returns a boolean if a field has been set.

### SetIssuingAccountTaxIdNil

`func (o *ChargeResponsePaymentMethod) SetIssuingAccountTaxIdNil(b bool)`

 SetIssuingAccountTaxIdNil sets the value for IssuingAccountTaxId to be an explicit nil

### UnsetIssuingAccountTaxId
`func (o *ChargeResponsePaymentMethod) UnsetIssuingAccountTaxId()`

UnsetIssuingAccountTaxId ensures that no value is present for IssuingAccountTaxId, not even an explicit nil
### GetPaymentAttempts

`func (o *ChargeResponsePaymentMethod) GetPaymentAttempts() []map[string]interface{}`

GetPaymentAttempts returns the PaymentAttempts field if non-nil, zero value otherwise.

### GetPaymentAttemptsOk

`func (o *ChargeResponsePaymentMethod) GetPaymentAttemptsOk() (*[]map[string]interface{}, bool)`

GetPaymentAttemptsOk returns a tuple with the PaymentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttempts

`func (o *ChargeResponsePaymentMethod) SetPaymentAttempts(v []map[string]interface{})`

SetPaymentAttempts sets PaymentAttempts field to given value.

### HasPaymentAttempts

`func (o *ChargeResponsePaymentMethod) HasPaymentAttempts() bool`

HasPaymentAttempts returns a boolean if a field has been set.

### GetReceivingAccountHolderName

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountHolderName() string`

GetReceivingAccountHolderName returns the ReceivingAccountHolderName field if non-nil, zero value otherwise.

### GetReceivingAccountHolderNameOk

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountHolderNameOk() (*string, bool)`

GetReceivingAccountHolderNameOk returns a tuple with the ReceivingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountHolderName

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountHolderName(v string)`

SetReceivingAccountHolderName sets ReceivingAccountHolderName field to given value.

### HasReceivingAccountHolderName

`func (o *ChargeResponsePaymentMethod) HasReceivingAccountHolderName() bool`

HasReceivingAccountHolderName returns a boolean if a field has been set.

### SetReceivingAccountHolderNameNil

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountHolderNameNil(b bool)`

 SetReceivingAccountHolderNameNil sets the value for ReceivingAccountHolderName to be an explicit nil

### UnsetReceivingAccountHolderName
`func (o *ChargeResponsePaymentMethod) UnsetReceivingAccountHolderName()`

UnsetReceivingAccountHolderName ensures that no value is present for ReceivingAccountHolderName, not even an explicit nil
### GetReceivingAccountNumber

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountNumber() string`

GetReceivingAccountNumber returns the ReceivingAccountNumber field if non-nil, zero value otherwise.

### GetReceivingAccountNumberOk

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountNumberOk() (*string, bool)`

GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountNumber

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountNumber(v string)`

SetReceivingAccountNumber sets ReceivingAccountNumber field to given value.

### HasReceivingAccountNumber

`func (o *ChargeResponsePaymentMethod) HasReceivingAccountNumber() bool`

HasReceivingAccountNumber returns a boolean if a field has been set.

### GetReceivingAccountBank

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountBank() string`

GetReceivingAccountBank returns the ReceivingAccountBank field if non-nil, zero value otherwise.

### GetReceivingAccountBankOk

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountBankOk() (*string, bool)`

GetReceivingAccountBankOk returns a tuple with the ReceivingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountBank

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountBank(v string)`

SetReceivingAccountBank sets ReceivingAccountBank field to given value.

### HasReceivingAccountBank

`func (o *ChargeResponsePaymentMethod) HasReceivingAccountBank() bool`

HasReceivingAccountBank returns a boolean if a field has been set.

### GetReceivingAccountTaxId

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountTaxId() string`

GetReceivingAccountTaxId returns the ReceivingAccountTaxId field if non-nil, zero value otherwise.

### GetReceivingAccountTaxIdOk

`func (o *ChargeResponsePaymentMethod) GetReceivingAccountTaxIdOk() (*string, bool)`

GetReceivingAccountTaxIdOk returns a tuple with the ReceivingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountTaxId

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountTaxId(v string)`

SetReceivingAccountTaxId sets ReceivingAccountTaxId field to given value.

### HasReceivingAccountTaxId

`func (o *ChargeResponsePaymentMethod) HasReceivingAccountTaxId() bool`

HasReceivingAccountTaxId returns a boolean if a field has been set.

### SetReceivingAccountTaxIdNil

`func (o *ChargeResponsePaymentMethod) SetReceivingAccountTaxIdNil(b bool)`

 SetReceivingAccountTaxIdNil sets the value for ReceivingAccountTaxId to be an explicit nil

### UnsetReceivingAccountTaxId
`func (o *ChargeResponsePaymentMethod) UnsetReceivingAccountTaxId()`

UnsetReceivingAccountTaxId ensures that no value is present for ReceivingAccountTaxId, not even an explicit nil
### GetReferenceNumber

`func (o *ChargeResponsePaymentMethod) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ChargeResponsePaymentMethod) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ChargeResponsePaymentMethod) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ChargeResponsePaymentMethod) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *ChargeResponsePaymentMethod) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *ChargeResponsePaymentMethod) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetTrackingCode

`func (o *ChargeResponsePaymentMethod) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ChargeResponsePaymentMethod) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ChargeResponsePaymentMethod) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ChargeResponsePaymentMethod) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ChargeResponsePaymentMethod) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ChargeResponsePaymentMethod) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetCancelUrl

`func (o *ChargeResponsePaymentMethod) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *ChargeResponsePaymentMethod) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *ChargeResponsePaymentMethod) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *ChargeResponsePaymentMethod) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *ChargeResponsePaymentMethod) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *ChargeResponsePaymentMethod) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *ChargeResponsePaymentMethod) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *ChargeResponsePaymentMethod) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ChargeResponsePaymentMethod) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ChargeResponsePaymentMethod) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ChargeResponsePaymentMethod) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetSuccessUrl

`func (o *ChargeResponsePaymentMethod) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *ChargeResponsePaymentMethod) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *ChargeResponsePaymentMethod) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *ChargeResponsePaymentMethod) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetDeepLink

`func (o *ChargeResponsePaymentMethod) GetDeepLink() string`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *ChargeResponsePaymentMethod) GetDeepLinkOk() (*string, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *ChargeResponsePaymentMethod) SetDeepLink(v string)`

SetDeepLink sets DeepLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


