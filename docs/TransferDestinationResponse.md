# TransferDestinationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to **string** | Name of the account holder. | [optional] 
**AccountNumber** | Pointer to **string** | Account number of the bank account. | [optional] 
**Bank** | Pointer to **string** | Name of the bank. | [optional] 
**CreatedAt** | Pointer to **int64** | Date and time of creation of the transfer. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the transfer. | [optional] 
**Object** | Pointer to **string** | Object name, which is bank_transfer_payout_method. | [optional] 
**PayeeId** | Pointer to **string** | Unique identifier of the payee. | [optional] 
**Type** | Pointer to **string** | Type of the payee. | [optional] 

## Methods

### NewTransferDestinationResponse

`func NewTransferDestinationResponse() *TransferDestinationResponse`

NewTransferDestinationResponse instantiates a new TransferDestinationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDestinationResponseWithDefaults

`func NewTransferDestinationResponseWithDefaults() *TransferDestinationResponse`

NewTransferDestinationResponseWithDefaults instantiates a new TransferDestinationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *TransferDestinationResponse) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *TransferDestinationResponse) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *TransferDestinationResponse) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *TransferDestinationResponse) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetAccountNumber

`func (o *TransferDestinationResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *TransferDestinationResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *TransferDestinationResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *TransferDestinationResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBank

`func (o *TransferDestinationResponse) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *TransferDestinationResponse) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *TransferDestinationResponse) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *TransferDestinationResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransferDestinationResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransferDestinationResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransferDestinationResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransferDestinationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *TransferDestinationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferDestinationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferDestinationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferDestinationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *TransferDestinationResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TransferDestinationResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TransferDestinationResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TransferDestinationResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPayeeId

`func (o *TransferDestinationResponse) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *TransferDestinationResponse) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *TransferDestinationResponse) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *TransferDestinationResponse) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetType

`func (o *TransferDestinationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferDestinationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferDestinationResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransferDestinationResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


