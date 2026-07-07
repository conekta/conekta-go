# PaymentMethodSpeiRecurrentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to **string** | Bank name for the SPEI payment method | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodSpeiRecurrentResponse

`func NewPaymentMethodSpeiRecurrentResponse(type_ string, id string, object string, createdAt int64, ) *PaymentMethodSpeiRecurrentResponse`

NewPaymentMethodSpeiRecurrentResponse instantiates a new PaymentMethodSpeiRecurrentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodSpeiRecurrentResponseWithDefaults

`func NewPaymentMethodSpeiRecurrentResponseWithDefaults() *PaymentMethodSpeiRecurrentResponse`

NewPaymentMethodSpeiRecurrentResponseWithDefaults instantiates a new PaymentMethodSpeiRecurrentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodSpeiRecurrentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodSpeiRecurrentResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodSpeiRecurrentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodSpeiRecurrentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodSpeiRecurrentResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodSpeiRecurrentResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *PaymentMethodSpeiRecurrentResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodSpeiRecurrentResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *PaymentMethodSpeiRecurrentResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentMethodSpeiRecurrentResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentMethodSpeiRecurrentResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetBank

`func (o *PaymentMethodSpeiRecurrentResponse) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *PaymentMethodSpeiRecurrentResponse) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *PaymentMethodSpeiRecurrentResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetReference

`func (o *PaymentMethodSpeiRecurrentResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodSpeiRecurrentResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentMethodSpeiRecurrentResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodSpeiRecurrentResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodSpeiRecurrentResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodSpeiRecurrentResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodSpeiRecurrentResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


