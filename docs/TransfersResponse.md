# TransfersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Amount in cents of the transfer. | [optional] 
**CreatedAt** | Pointer to **int64** | Date and time of creation of the transfer. | [optional] 
**Currency** | Pointer to **string** | The currency of the transfer. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217) | [optional] 
**Id** | Pointer to **string** | Unique identifier of the transfer. | [optional] 
**Livemode** | Pointer to **bool** | Indicates whether the transfer was created in live mode or test mode. | [optional] 
**Method** | Pointer to [**TransferMethodResponse**](TransferMethodResponse.md) |  | [optional] 
**Object** | Pointer to **string** | Object name, which is transfer. | [optional] 
**StatementDescription** | Pointer to **string** | Description of the transfer. | [optional] 
**StatementReference** | Pointer to **string** | Reference number of the transfer. | [optional] 
**Status** | Pointer to **string** | Code indicating transfer status. | [optional] 

## Methods

### NewTransfersResponse

`func NewTransfersResponse() *TransfersResponse`

NewTransfersResponse instantiates a new TransfersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransfersResponseWithDefaults

`func NewTransfersResponseWithDefaults() *TransfersResponse`

NewTransfersResponseWithDefaults instantiates a new TransfersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransfersResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransfersResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransfersResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransfersResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransfersResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransfersResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransfersResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransfersResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *TransfersResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransfersResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransfersResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransfersResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *TransfersResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransfersResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransfersResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransfersResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *TransfersResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *TransfersResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *TransfersResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *TransfersResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMethod

`func (o *TransfersResponse) GetMethod() TransferMethodResponse`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransfersResponse) GetMethodOk() (*TransferMethodResponse, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransfersResponse) SetMethod(v TransferMethodResponse)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransfersResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetObject

`func (o *TransfersResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TransfersResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TransfersResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TransfersResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatementDescription

`func (o *TransfersResponse) GetStatementDescription() string`

GetStatementDescription returns the StatementDescription field if non-nil, zero value otherwise.

### GetStatementDescriptionOk

`func (o *TransfersResponse) GetStatementDescriptionOk() (*string, bool)`

GetStatementDescriptionOk returns a tuple with the StatementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescription

`func (o *TransfersResponse) SetStatementDescription(v string)`

SetStatementDescription sets StatementDescription field to given value.

### HasStatementDescription

`func (o *TransfersResponse) HasStatementDescription() bool`

HasStatementDescription returns a boolean if a field has been set.

### GetStatementReference

`func (o *TransfersResponse) GetStatementReference() string`

GetStatementReference returns the StatementReference field if non-nil, zero value otherwise.

### GetStatementReferenceOk

`func (o *TransfersResponse) GetStatementReferenceOk() (*string, bool)`

GetStatementReferenceOk returns a tuple with the StatementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementReference

`func (o *TransfersResponse) SetStatementReference(v string)`

SetStatementReference sets StatementReference field to given value.

### HasStatementReference

`func (o *TransfersResponse) HasStatementReference() bool`

HasStatementReference returns a boolean if a field has been set.

### GetStatus

`func (o *TransfersResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransfersResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransfersResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransfersResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


