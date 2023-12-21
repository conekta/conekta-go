# PaymentMethodBankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**Bank** | Pointer to **string** |  | [optional] 
**Clabe** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExecutedAt** | Pointer to **NullableInt32** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
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

## Methods

### NewPaymentMethodBankTransfer

`func NewPaymentMethodBankTransfer(object string, ) *PaymentMethodBankTransfer`

NewPaymentMethodBankTransfer instantiates a new PaymentMethodBankTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodBankTransferWithDefaults

`func NewPaymentMethodBankTransferWithDefaults() *PaymentMethodBankTransfer`

NewPaymentMethodBankTransferWithDefaults instantiates a new PaymentMethodBankTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodBankTransfer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodBankTransfer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodBankTransfer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodBankTransfer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodBankTransfer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodBankTransfer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodBankTransfer) SetObject(v string)`

SetObject sets Object field to given value.


### GetBank

`func (o *PaymentMethodBankTransfer) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *PaymentMethodBankTransfer) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *PaymentMethodBankTransfer) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *PaymentMethodBankTransfer) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetClabe

`func (o *PaymentMethodBankTransfer) GetClabe() string`

GetClabe returns the Clabe field if non-nil, zero value otherwise.

### GetClabeOk

`func (o *PaymentMethodBankTransfer) GetClabeOk() (*string, bool)`

GetClabeOk returns a tuple with the Clabe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClabe

`func (o *PaymentMethodBankTransfer) SetClabe(v string)`

SetClabe sets Clabe field to given value.

### HasClabe

`func (o *PaymentMethodBankTransfer) HasClabe() bool`

HasClabe returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethodBankTransfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodBankTransfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodBankTransfer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodBankTransfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethodBankTransfer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethodBankTransfer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExecutedAt

`func (o *PaymentMethodBankTransfer) GetExecutedAt() int32`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *PaymentMethodBankTransfer) GetExecutedAtOk() (*int32, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *PaymentMethodBankTransfer) SetExecutedAt(v int32)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *PaymentMethodBankTransfer) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### SetExecutedAtNil

`func (o *PaymentMethodBankTransfer) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *PaymentMethodBankTransfer) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil
### GetExpiresAt

`func (o *PaymentMethodBankTransfer) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodBankTransfer) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodBankTransfer) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodBankTransfer) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIssuingAccountBank

`func (o *PaymentMethodBankTransfer) GetIssuingAccountBank() string`

GetIssuingAccountBank returns the IssuingAccountBank field if non-nil, zero value otherwise.

### GetIssuingAccountBankOk

`func (o *PaymentMethodBankTransfer) GetIssuingAccountBankOk() (*string, bool)`

GetIssuingAccountBankOk returns a tuple with the IssuingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountBank

`func (o *PaymentMethodBankTransfer) SetIssuingAccountBank(v string)`

SetIssuingAccountBank sets IssuingAccountBank field to given value.

### HasIssuingAccountBank

`func (o *PaymentMethodBankTransfer) HasIssuingAccountBank() bool`

HasIssuingAccountBank returns a boolean if a field has been set.

### SetIssuingAccountBankNil

`func (o *PaymentMethodBankTransfer) SetIssuingAccountBankNil(b bool)`

 SetIssuingAccountBankNil sets the value for IssuingAccountBank to be an explicit nil

### UnsetIssuingAccountBank
`func (o *PaymentMethodBankTransfer) UnsetIssuingAccountBank()`

UnsetIssuingAccountBank ensures that no value is present for IssuingAccountBank, not even an explicit nil
### GetIssuingAccountNumber

`func (o *PaymentMethodBankTransfer) GetIssuingAccountNumber() string`

GetIssuingAccountNumber returns the IssuingAccountNumber field if non-nil, zero value otherwise.

### GetIssuingAccountNumberOk

`func (o *PaymentMethodBankTransfer) GetIssuingAccountNumberOk() (*string, bool)`

GetIssuingAccountNumberOk returns a tuple with the IssuingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountNumber

`func (o *PaymentMethodBankTransfer) SetIssuingAccountNumber(v string)`

SetIssuingAccountNumber sets IssuingAccountNumber field to given value.

### HasIssuingAccountNumber

`func (o *PaymentMethodBankTransfer) HasIssuingAccountNumber() bool`

HasIssuingAccountNumber returns a boolean if a field has been set.

### SetIssuingAccountNumberNil

`func (o *PaymentMethodBankTransfer) SetIssuingAccountNumberNil(b bool)`

 SetIssuingAccountNumberNil sets the value for IssuingAccountNumber to be an explicit nil

### UnsetIssuingAccountNumber
`func (o *PaymentMethodBankTransfer) UnsetIssuingAccountNumber()`

UnsetIssuingAccountNumber ensures that no value is present for IssuingAccountNumber, not even an explicit nil
### GetIssuingAccountHolderName

`func (o *PaymentMethodBankTransfer) GetIssuingAccountHolderName() string`

GetIssuingAccountHolderName returns the IssuingAccountHolderName field if non-nil, zero value otherwise.

### GetIssuingAccountHolderNameOk

`func (o *PaymentMethodBankTransfer) GetIssuingAccountHolderNameOk() (*string, bool)`

GetIssuingAccountHolderNameOk returns a tuple with the IssuingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountHolderName

`func (o *PaymentMethodBankTransfer) SetIssuingAccountHolderName(v string)`

SetIssuingAccountHolderName sets IssuingAccountHolderName field to given value.

### HasIssuingAccountHolderName

`func (o *PaymentMethodBankTransfer) HasIssuingAccountHolderName() bool`

HasIssuingAccountHolderName returns a boolean if a field has been set.

### SetIssuingAccountHolderNameNil

`func (o *PaymentMethodBankTransfer) SetIssuingAccountHolderNameNil(b bool)`

 SetIssuingAccountHolderNameNil sets the value for IssuingAccountHolderName to be an explicit nil

### UnsetIssuingAccountHolderName
`func (o *PaymentMethodBankTransfer) UnsetIssuingAccountHolderName()`

UnsetIssuingAccountHolderName ensures that no value is present for IssuingAccountHolderName, not even an explicit nil
### GetIssuingAccountTaxId

`func (o *PaymentMethodBankTransfer) GetIssuingAccountTaxId() string`

GetIssuingAccountTaxId returns the IssuingAccountTaxId field if non-nil, zero value otherwise.

### GetIssuingAccountTaxIdOk

`func (o *PaymentMethodBankTransfer) GetIssuingAccountTaxIdOk() (*string, bool)`

GetIssuingAccountTaxIdOk returns a tuple with the IssuingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAccountTaxId

`func (o *PaymentMethodBankTransfer) SetIssuingAccountTaxId(v string)`

SetIssuingAccountTaxId sets IssuingAccountTaxId field to given value.

### HasIssuingAccountTaxId

`func (o *PaymentMethodBankTransfer) HasIssuingAccountTaxId() bool`

HasIssuingAccountTaxId returns a boolean if a field has been set.

### SetIssuingAccountTaxIdNil

`func (o *PaymentMethodBankTransfer) SetIssuingAccountTaxIdNil(b bool)`

 SetIssuingAccountTaxIdNil sets the value for IssuingAccountTaxId to be an explicit nil

### UnsetIssuingAccountTaxId
`func (o *PaymentMethodBankTransfer) UnsetIssuingAccountTaxId()`

UnsetIssuingAccountTaxId ensures that no value is present for IssuingAccountTaxId, not even an explicit nil
### GetPaymentAttempts

`func (o *PaymentMethodBankTransfer) GetPaymentAttempts() []map[string]interface{}`

GetPaymentAttempts returns the PaymentAttempts field if non-nil, zero value otherwise.

### GetPaymentAttemptsOk

`func (o *PaymentMethodBankTransfer) GetPaymentAttemptsOk() (*[]map[string]interface{}, bool)`

GetPaymentAttemptsOk returns a tuple with the PaymentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttempts

`func (o *PaymentMethodBankTransfer) SetPaymentAttempts(v []map[string]interface{})`

SetPaymentAttempts sets PaymentAttempts field to given value.

### HasPaymentAttempts

`func (o *PaymentMethodBankTransfer) HasPaymentAttempts() bool`

HasPaymentAttempts returns a boolean if a field has been set.

### GetReceivingAccountHolderName

`func (o *PaymentMethodBankTransfer) GetReceivingAccountHolderName() string`

GetReceivingAccountHolderName returns the ReceivingAccountHolderName field if non-nil, zero value otherwise.

### GetReceivingAccountHolderNameOk

`func (o *PaymentMethodBankTransfer) GetReceivingAccountHolderNameOk() (*string, bool)`

GetReceivingAccountHolderNameOk returns a tuple with the ReceivingAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountHolderName

`func (o *PaymentMethodBankTransfer) SetReceivingAccountHolderName(v string)`

SetReceivingAccountHolderName sets ReceivingAccountHolderName field to given value.

### HasReceivingAccountHolderName

`func (o *PaymentMethodBankTransfer) HasReceivingAccountHolderName() bool`

HasReceivingAccountHolderName returns a boolean if a field has been set.

### SetReceivingAccountHolderNameNil

`func (o *PaymentMethodBankTransfer) SetReceivingAccountHolderNameNil(b bool)`

 SetReceivingAccountHolderNameNil sets the value for ReceivingAccountHolderName to be an explicit nil

### UnsetReceivingAccountHolderName
`func (o *PaymentMethodBankTransfer) UnsetReceivingAccountHolderName()`

UnsetReceivingAccountHolderName ensures that no value is present for ReceivingAccountHolderName, not even an explicit nil
### GetReceivingAccountNumber

`func (o *PaymentMethodBankTransfer) GetReceivingAccountNumber() string`

GetReceivingAccountNumber returns the ReceivingAccountNumber field if non-nil, zero value otherwise.

### GetReceivingAccountNumberOk

`func (o *PaymentMethodBankTransfer) GetReceivingAccountNumberOk() (*string, bool)`

GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountNumber

`func (o *PaymentMethodBankTransfer) SetReceivingAccountNumber(v string)`

SetReceivingAccountNumber sets ReceivingAccountNumber field to given value.

### HasReceivingAccountNumber

`func (o *PaymentMethodBankTransfer) HasReceivingAccountNumber() bool`

HasReceivingAccountNumber returns a boolean if a field has been set.

### GetReceivingAccountBank

`func (o *PaymentMethodBankTransfer) GetReceivingAccountBank() string`

GetReceivingAccountBank returns the ReceivingAccountBank field if non-nil, zero value otherwise.

### GetReceivingAccountBankOk

`func (o *PaymentMethodBankTransfer) GetReceivingAccountBankOk() (*string, bool)`

GetReceivingAccountBankOk returns a tuple with the ReceivingAccountBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountBank

`func (o *PaymentMethodBankTransfer) SetReceivingAccountBank(v string)`

SetReceivingAccountBank sets ReceivingAccountBank field to given value.

### HasReceivingAccountBank

`func (o *PaymentMethodBankTransfer) HasReceivingAccountBank() bool`

HasReceivingAccountBank returns a boolean if a field has been set.

### GetReceivingAccountTaxId

`func (o *PaymentMethodBankTransfer) GetReceivingAccountTaxId() string`

GetReceivingAccountTaxId returns the ReceivingAccountTaxId field if non-nil, zero value otherwise.

### GetReceivingAccountTaxIdOk

`func (o *PaymentMethodBankTransfer) GetReceivingAccountTaxIdOk() (*string, bool)`

GetReceivingAccountTaxIdOk returns a tuple with the ReceivingAccountTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountTaxId

`func (o *PaymentMethodBankTransfer) SetReceivingAccountTaxId(v string)`

SetReceivingAccountTaxId sets ReceivingAccountTaxId field to given value.

### HasReceivingAccountTaxId

`func (o *PaymentMethodBankTransfer) HasReceivingAccountTaxId() bool`

HasReceivingAccountTaxId returns a boolean if a field has been set.

### SetReceivingAccountTaxIdNil

`func (o *PaymentMethodBankTransfer) SetReceivingAccountTaxIdNil(b bool)`

 SetReceivingAccountTaxIdNil sets the value for ReceivingAccountTaxId to be an explicit nil

### UnsetReceivingAccountTaxId
`func (o *PaymentMethodBankTransfer) UnsetReceivingAccountTaxId()`

UnsetReceivingAccountTaxId ensures that no value is present for ReceivingAccountTaxId, not even an explicit nil
### GetReferenceNumber

`func (o *PaymentMethodBankTransfer) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *PaymentMethodBankTransfer) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *PaymentMethodBankTransfer) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *PaymentMethodBankTransfer) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *PaymentMethodBankTransfer) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *PaymentMethodBankTransfer) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetTrackingCode

`func (o *PaymentMethodBankTransfer) GetTrackingCode() string`

GetTrackingCode returns the TrackingCode field if non-nil, zero value otherwise.

### GetTrackingCodeOk

`func (o *PaymentMethodBankTransfer) GetTrackingCodeOk() (*string, bool)`

GetTrackingCodeOk returns a tuple with the TrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCode

`func (o *PaymentMethodBankTransfer) SetTrackingCode(v string)`

SetTrackingCode sets TrackingCode field to given value.

### HasTrackingCode

`func (o *PaymentMethodBankTransfer) HasTrackingCode() bool`

HasTrackingCode returns a boolean if a field has been set.

### SetTrackingCodeNil

`func (o *PaymentMethodBankTransfer) SetTrackingCodeNil(b bool)`

 SetTrackingCodeNil sets the value for TrackingCode to be an explicit nil

### UnsetTrackingCode
`func (o *PaymentMethodBankTransfer) UnsetTrackingCode()`

UnsetTrackingCode ensures that no value is present for TrackingCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


