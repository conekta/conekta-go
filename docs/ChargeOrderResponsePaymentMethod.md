# ChargeOrderResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**Agreement** | Pointer to **string** | Agreement ID | [optional] 
**AuthCode** | Pointer to **string** |  | [optional] 
**CashierId** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
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
**RedirectUrl** | Pointer to **string** | URL to redirect the customer to complete the payment | [optional] 
**SuccessUrl** | Pointer to **string** | URL to redirect the customer after a successful payment | [optional] 

## Methods

### NewChargeOrderResponsePaymentMethod

`func NewChargeOrderResponsePaymentMethod(object string, expiresAt int64, productType string, ) *ChargeOrderResponsePaymentMethod`

NewChargeOrderResponsePaymentMethod instantiates a new ChargeOrderResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeOrderResponsePaymentMethodWithDefaults

`func NewChargeOrderResponsePaymentMethodWithDefaults() *ChargeOrderResponsePaymentMethod`

NewChargeOrderResponsePaymentMethodWithDefaults instantiates a new ChargeOrderResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargeOrderResponsePaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeOrderResponsePaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeOrderResponsePaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargeOrderResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *ChargeOrderResponsePaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeOrderResponsePaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeOrderResponsePaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.


### GetAgreement

`func (o *ChargeOrderResponsePaymentMethod) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ChargeOrderResponsePaymentMethod) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ChargeOrderResponsePaymentMethod) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ChargeOrderResponsePaymentMethod) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAuthCode

`func (o *ChargeOrderResponsePaymentMethod) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeOrderResponsePaymentMethod) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeOrderResponsePaymentMethod) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeOrderResponsePaymentMethod) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCashierId

`func (o *ChargeOrderResponsePaymentMethod) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *ChargeOrderResponsePaymentMethod) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *ChargeOrderResponsePaymentMethod) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *ChargeOrderResponsePaymentMethod) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *ChargeOrderResponsePaymentMethod) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *ChargeOrderResponsePaymentMethod) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *ChargeOrderResponsePaymentMethod) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ChargeOrderResponsePaymentMethod) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ChargeOrderResponsePaymentMethod) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ChargeOrderResponsePaymentMethod) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ChargeOrderResponsePaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeOrderResponsePaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeOrderResponsePaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.


### GetProductType

`func (o *ChargeOrderResponsePaymentMethod) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ChargeOrderResponsePaymentMethod) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ChargeOrderResponsePaymentMethod) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetServiceName

`func (o *ChargeOrderResponsePaymentMethod) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ChargeOrderResponsePaymentMethod) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ChargeOrderResponsePaymentMethod) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *ChargeOrderResponsePaymentMethod) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChargeOrderResponsePaymentMethod) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChargeOrderResponsePaymentMethod) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChargeOrderResponsePaymentMethod) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ChargeOrderResponsePaymentMethod) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ChargeOrderResponsePaymentMethod) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *ChargeOrderResponsePaymentMethod) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ChargeOrderResponsePaymentMethod) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *ChargeOrderResponsePaymentMethod) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetCustomerIpAddress

`func (o *ChargeOrderResponsePaymentMethod) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *ChargeOrderResponsePaymentMethod) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *ChargeOrderResponsePaymentMethod) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *ChargeOrderResponsePaymentMethod) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.

### GetAccountType

`func (o *ChargeOrderResponsePaymentMethod) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ChargeOrderResponsePaymentMethod) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ChargeOrderResponsePaymentMethod) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ChargeOrderResponsePaymentMethod) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBrand

`func (o *ChargeOrderResponsePaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ChargeOrderResponsePaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ChargeOrderResponsePaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ChargeOrderResponsePaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetContractId

`func (o *ChargeOrderResponsePaymentMethod) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ChargeOrderResponsePaymentMethod) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ChargeOrderResponsePaymentMethod) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ChargeOrderResponsePaymentMethod) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCountry

`func (o *ChargeOrderResponsePaymentMethod) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ChargeOrderResponsePaymentMethod) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ChargeOrderResponsePaymentMethod) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ChargeOrderResponsePaymentMethod) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpMonth

`func (o *ChargeOrderResponsePaymentMethod) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *ChargeOrderResponsePaymentMethod) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *ChargeOrderResponsePaymentMethod) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *ChargeOrderResponsePaymentMethod) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *ChargeOrderResponsePaymentMethod) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *ChargeOrderResponsePaymentMethod) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *ChargeOrderResponsePaymentMethod) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *ChargeOrderResponsePaymentMethod) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFraudIndicators

`func (o *ChargeOrderResponsePaymentMethod) GetFraudIndicators() []map[string]interface{}`

GetFraudIndicators returns the FraudIndicators field if non-nil, zero value otherwise.

### GetFraudIndicatorsOk

`func (o *ChargeOrderResponsePaymentMethod) GetFraudIndicatorsOk() (*[]map[string]interface{}, bool)`

GetFraudIndicatorsOk returns a tuple with the FraudIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudIndicators

`func (o *ChargeOrderResponsePaymentMethod) SetFraudIndicators(v []map[string]interface{})`

SetFraudIndicators sets FraudIndicators field to given value.

### HasFraudIndicators

`func (o *ChargeOrderResponsePaymentMethod) HasFraudIndicators() bool`

HasFraudIndicators returns a boolean if a field has been set.

### GetIssuer

`func (o *ChargeOrderResponsePaymentMethod) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ChargeOrderResponsePaymentMethod) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ChargeOrderResponsePaymentMethod) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ChargeOrderResponsePaymentMethod) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLast4

`func (o *ChargeOrderResponsePaymentMethod) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *ChargeOrderResponsePaymentMethod) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *ChargeOrderResponsePaymentMethod) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *ChargeOrderResponsePaymentMethod) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *ChargeOrderResponsePaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeOrderResponsePaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChargeOrderResponsePaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBank

`func (o *ChargeOrderResponsePaymentMethod) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *ChargeOrderResponsePaymentMethod) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *ChargeOrderResponsePaymentMethod) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *ChargeOrderResponsePaymentMethod) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetClabe

`func (o *ChargeOrderResponsePaymentMethod) GetClabe() string`

GetClabe returns the Clabe field if non-nil, zero value otherwise.

### GetClabeOk

`func (o *ChargeOrderResponsePaymentMethod) GetClabeOk() (*string, bool)`

GetClabeOk returns a tuple with the Clabe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClabe

`func (o *ChargeOrderResponsePaymentMethod) SetClabe(v string)`

SetClabe sets Clabe field to given value.

### HasClabe

`func (o *ChargeOrderResponsePaymentMethod) HasClabe() bool`

HasClabe returns a boolean if a field has been set.

### GetDescription

`func (o *ChargeOrderResponsePaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeOrderResponsePaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeOrderResponsePaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChargeOrderResponsePaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ChargeOrderResponsePaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ChargeOrderResponsePaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExecutedAt

`func (o *ChargeOrderResponsePaymentMethod) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ChargeOrderResponsePaymentMethod) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ChargeOrderResponsePaymentMethod) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *ChargeOrderResponsePaymentMethod) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### SetExecutedAtNil

`func (o *ChargeOrderResponsePaymentMethod) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *ChargeOrderResponsePaymentMethod) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
### GetIssuingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountBank() string`

GetIssuingAccountBank returns the IssuingAccountBank field if non-nil, zero value otherwise.

### GetIssuingAccountBankOk

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountBankOk() (*string, bool)`

GetIssuingAccountBankOk returns a tuple with the IssuingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountBank(v string)`

SetIssuingAccountBank sets IssuingAccountBank field to given value.

### HasIssuingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) HasIssuingAccountBank() bool`

HasIssuingAccountBank returns a boolean if a field has been set.

### SetIssuingAccountBankNil

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountBankNil(b bool)`

 SetIssuingAccountBankNil sets the value for IssuingAccountBank to be an explicit nil

### UnsetIssuingAccountBank
`func (o *ChargeOrderResponsePaymentMethod) UnsetIssuingAccountBank()`

UnsetIssuingAccountBank ensures that no value is present for IssuingAccountBank, not even an explicit nil
### GetIssuingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountNumber() string`

GetIssuingAccountNumber returns the IssuingAccountNumber field if non-nil, zero value otherwise.

### GetIssuingAccountNumberOk

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountNumberOk() (*string, bool)`

GetIssuingAccountNumberOk returns a tuple with the IssuingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountNumber(v string)`

SetIssuingAccountNumber sets IssuingAccountNumber field to given value.

### HasIssuingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) HasIssuingAccountNumber() bool`

HasIssuingAccountNumber returns a boolean if a field has been set.

### SetIssuingAccountNumberNil

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountNumberNil(b bool)`

 SetIssuingAccountNumberNil sets the value for IssuingAccountNumber to be an explicit nil

### UnsetIssuingAccountNumber
`func (o *ChargeOrderResponsePaymentMethod) UnsetIssuingAccountNumber()`

UnsetIssuingAccountNumber ensures that no value is present for IssuingAccountNumber, not even an explicit nil
### GetIssuingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountHolderName() string`

GetIssuingAccountHolderName returns the IssuingAccountHolderName field if non-nil, zero value otherwise.

### GetIssuingAccountHolderNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountHolderNameOk() (*string, bool)`

GetIssuingAccountHolderNameOk returns a tuple with the IssuingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountHolderName(v string)`

SetIssuingAccountHolderName sets IssuingAccountHolderName field to given value.

### HasIssuingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) HasIssuingAccountHolderName() bool`

HasIssuingAccountHolderName returns a boolean if a field has been set.

### SetIssuingAccountHolderNameNil

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountHolderNameNil(b bool)`

 SetIssuingAccountHolderNameNil sets the value for IssuingAccountHolderName to be an explicit nil

### UnsetIssuingAccountHolderName
`func (o *ChargeOrderResponsePaymentMethod) UnsetIssuingAccountHolderName()`

UnsetIssuingAccountHolderName ensures that no value is present for IssuingAccountHolderName, not even an explicit nil
### GetIssuingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountTaxId() string`

GetIssuingAccountTaxId returns the IssuingAccountTaxId field if non-nil, zero value otherwise.

### GetIssuingAccountTaxIdOk

`func (o *ChargeOrderResponsePaymentMethod) GetIssuingAccountTaxIdOk() (*string, bool)`

GetIssuingAccountTaxIdOk returns a tuple with the IssuingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountTaxId(v string)`

SetIssuingAccountTaxId sets IssuingAccountTaxId field to given value.

### HasIssuingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) HasIssuingAccountTaxId() bool`

HasIssuingAccountTaxId returns a boolean if a field has been set.

### SetIssuingAccountTaxIdNil

`func (o *ChargeOrderResponsePaymentMethod) SetIssuingAccountTaxIdNil(b bool)`

 SetIssuingAccountTaxIdNil sets the value for IssuingAccountTaxId to be an explicit nil

### UnsetIssuingAccountTaxId
`func (o *ChargeOrderResponsePaymentMethod) UnsetIssuingAccountTaxId()`

UnsetIssuingAccountTaxId ensures that no value is present for IssuingAccountTaxId, not even an explicit nil
### GetPaymentAttempts

`func (o *ChargeOrderResponsePaymentMethod) GetPaymentAttempts() []map[string]interface{}`

GetPaymentAttempts returns the PaymentAttempts field if non-nil, zero value otherwise.

### GetPaymentAttemptsOk

`func (o *ChargeOrderResponsePaymentMethod) GetPaymentAttemptsOk() (*[]map[string]interface{}, bool)`

GetPaymentAttemptsOk returns a tuple with the PaymentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttempts

`func (o *ChargeOrderResponsePaymentMethod) SetPaymentAttempts(v []map[string]interface{})`

SetPaymentAttempts sets PaymentAttempts field to given value.

### HasPaymentAttempts

`func (o *ChargeOrderResponsePaymentMethod) HasPaymentAttempts() bool`

HasPaymentAttempts returns a boolean if a field has been set.

### GetReceivingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountHolderName() string`

GetReceivingAccountHolderName returns the ReceivingAccountHolderName field if non-nil, zero value otherwise.

### GetReceivingAccountHolderNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountHolderNameOk() (*string, bool)`

GetReceivingAccountHolderNameOk returns a tuple with the ReceivingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountHolderName(v string)`

SetReceivingAccountHolderName sets ReceivingAccountHolderName field to given value.

### HasReceivingAccountHolderName

`func (o *ChargeOrderResponsePaymentMethod) HasReceivingAccountHolderName() bool`

HasReceivingAccountHolderName returns a boolean if a field has been set.

### SetReceivingAccountHolderNameNil

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountHolderNameNil(b bool)`

 SetReceivingAccountHolderNameNil sets the value for ReceivingAccountHolderName to be an explicit nil

### UnsetReceivingAccountHolderName
`func (o *ChargeOrderResponsePaymentMethod) UnsetReceivingAccountHolderName()`

UnsetReceivingAccountHolderName ensures that no value is present for ReceivingAccountHolderName, not even an explicit nil
### GetReceivingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountNumber() string`

GetReceivingAccountNumber returns the ReceivingAccountNumber field if non-nil, zero value otherwise.

### GetReceivingAccountNumberOk

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountNumberOk() (*string, bool)`

GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountNumber(v string)`

SetReceivingAccountNumber sets ReceivingAccountNumber field to given value.

### HasReceivingAccountNumber

`func (o *ChargeOrderResponsePaymentMethod) HasReceivingAccountNumber() bool`

HasReceivingAccountNumber returns a boolean if a field has been set.

### GetReceivingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountBank() string`

GetReceivingAccountBank returns the ReceivingAccountBank field if non-nil, zero value otherwise.

### GetReceivingAccountBankOk

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountBankOk() (*string, bool)`

GetReceivingAccountBankOk returns a tuple with the ReceivingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountBank(v string)`

SetReceivingAccountBank sets ReceivingAccountBank field to given value.

### HasReceivingAccountBank

`func (o *ChargeOrderResponsePaymentMethod) HasReceivingAccountBank() bool`

HasReceivingAccountBank returns a boolean if a field has been set.

### GetReceivingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountTaxId() string`

GetReceivingAccountTaxId returns the ReceivingAccountTaxId field if non-nil, zero value otherwise.

### GetReceivingAccountTaxIdOk

`func (o *ChargeOrderResponsePaymentMethod) GetReceivingAccountTaxIdOk() (*string, bool)`

GetReceivingAccountTaxIdOk returns a tuple with the ReceivingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountTaxId(v string)`

SetReceivingAccountTaxId sets ReceivingAccountTaxId field to given value.

### HasReceivingAccountTaxId

`func (o *ChargeOrderResponsePaymentMethod) HasReceivingAccountTaxId() bool`

HasReceivingAccountTaxId returns a boolean if a field has been set.

### SetReceivingAccountTaxIdNil

`func (o *ChargeOrderResponsePaymentMethod) SetReceivingAccountTaxIdNil(b bool)`

 SetReceivingAccountTaxIdNil sets the value for ReceivingAccountTaxId to be an explicit nil

### UnsetReceivingAccountTaxId
`func (o *ChargeOrderResponsePaymentMethod) UnsetReceivingAccountTaxId()`

UnsetReceivingAccountTaxId ensures that no value is present for ReceivingAccountTaxId, not even an explicit nil
### GetReferenceNumber

`func (o *ChargeOrderResponsePaymentMethod) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ChargeOrderResponsePaymentMethod) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ChargeOrderResponsePaymentMethod) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ChargeOrderResponsePaymentMethod) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *ChargeOrderResponsePaymentMethod) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *ChargeOrderResponsePaymentMethod) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetTrackingCode

`func (o *ChargeOrderResponsePaymentMethod) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *ChargeOrderResponsePaymentMethod) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *ChargeOrderResponsePaymentMethod) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *ChargeOrderResponsePaymentMethod) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *ChargeOrderResponsePaymentMethod) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *ChargeOrderResponsePaymentMethod) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil
### GetCancelUrl

`func (o *ChargeOrderResponsePaymentMethod) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *ChargeOrderResponsePaymentMethod) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *ChargeOrderResponsePaymentMethod) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *ChargeOrderResponsePaymentMethod) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *ChargeOrderResponsePaymentMethod) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *ChargeOrderResponsePaymentMethod) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ChargeOrderResponsePaymentMethod) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ChargeOrderResponsePaymentMethod) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ChargeOrderResponsePaymentMethod) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *ChargeOrderResponsePaymentMethod) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *ChargeOrderResponsePaymentMethod) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *ChargeOrderResponsePaymentMethod) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


