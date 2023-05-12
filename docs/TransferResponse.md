# TransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Amount in cents of the transfer. | [optional] 
**CreatedAt** | Pointer to **int64** | Date and time of creation of the transfer in Unix format. | [optional] 
**Currency** | Pointer to **string** | The currency of the transfer. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217) | [optional] 
**Id** | Pointer to **string** | Unique identifier of the transfer. | [optional] 
**Livemode** | Pointer to **bool** | Indicates whether the transfer was created in live mode or test mode. | [optional] 
**Destination** | Pointer to [**TransferDestinationResponse**](TransferDestinationResponse.md) |  | [optional] 
**Object** | Pointer to **string** | Object name, which is transfer. | [optional] 
**StatementDescription** | Pointer to **string** | Description of the transfer. | [optional] 
**StatementReference** | Pointer to **string** | Reference number of the transfer. | [optional] 
**Status** | Pointer to **string** | Code indicating transfer status. | [optional] 

## Methods

### NewTransferResponse

`func NewTransferResponse() *TransferResponse`

NewTransferResponse instantiates a new TransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResponseWithDefaults

`func NewTransferResponseWithDefaults() *TransferResponse`

NewTransferResponseWithDefaults instantiates a new TransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransferResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransferResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransferResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransferResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *TransferResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransferResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *TransferResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *TransferResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *TransferResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *TransferResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *TransferResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetDestination

`func (o *TransferResponse) GetDestination() TransferDestinationResponse`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TransferResponse) GetDestinationOk() (*TransferDestinationResponse, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TransferResponse) SetDestination(v TransferDestinationResponse)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *TransferResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetObject

`func (o *TransferResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TransferResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TransferResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TransferResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatementDescription

`func (o *TransferResponse) GetStatementDescription() string`

GetStatementDescription returns the StatementDescription field if non-nil, zero value otherwise.

### GetStatementDescriptionOk

`func (o *TransferResponse) GetStatementDescriptionOk() (*string, bool)`

GetStatementDescriptionOk returns a tuple with the StatementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescription

`func (o *TransferResponse) SetStatementDescription(v string)`

SetStatementDescription sets StatementDescription field to given value.

### HasStatementDescription

`func (o *TransferResponse) HasStatementDescription() bool`

HasStatementDescription returns a boolean if a field has been set.

### GetStatementReference

`func (o *TransferResponse) GetStatementReference() string`

GetStatementReference returns the StatementReference field if non-nil, zero value otherwise.

### GetStatementReferenceOk

`func (o *TransferResponse) GetStatementReferenceOk() (*string, bool)`

GetStatementReferenceOk returns a tuple with the StatementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementReference

`func (o *TransferResponse) SetStatementReference(v string)`

SetStatementReference sets StatementReference field to given value.

### HasStatementReference

`func (o *TransferResponse) HasStatementReference() bool`

HasStatementReference returns a boolean if a field has been set.

### GetStatus

`func (o *TransferResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransferResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


