# ChargeResponseRefundsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**AuthCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**ExpiresAt** | Pointer to **int64** | refund expiration date | [optional] 
**Id** | **string** |  | 
**Object** | **string** |  | 
**Status** | Pointer to **string** | refund status | [optional] 

## Methods

### NewChargeResponseRefundsData

`func NewChargeResponseRefundsData(amount int64, createdAt int64, id string, object string, ) *ChargeResponseRefundsData`

NewChargeResponseRefundsData instantiates a new ChargeResponseRefundsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeResponseRefundsDataWithDefaults

`func NewChargeResponseRefundsDataWithDefaults() *ChargeResponseRefundsData`

NewChargeResponseRefundsDataWithDefaults instantiates a new ChargeResponseRefundsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargeResponseRefundsData) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargeResponseRefundsData) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargeResponseRefundsData) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAuthCode

`func (o *ChargeResponseRefundsData) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeResponseRefundsData) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeResponseRefundsData) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeResponseRefundsData) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargeResponseRefundsData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeResponseRefundsData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeResponseRefundsData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *ChargeResponseRefundsData) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeResponseRefundsData) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeResponseRefundsData) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeResponseRefundsData) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ChargeResponseRefundsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeResponseRefundsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeResponseRefundsData) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ChargeResponseRefundsData) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeResponseRefundsData) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeResponseRefundsData) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *ChargeResponseRefundsData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeResponseRefundsData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeResponseRefundsData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargeResponseRefundsData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


