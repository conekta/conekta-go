# ChargeDataPaymentMethodCashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | Pointer to **NullableInt32** |  | [optional] 
**CashierId** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeDataPaymentMethodCashResponse

`func NewChargeDataPaymentMethodCashResponse() *ChargeDataPaymentMethodCashResponse`

NewChargeDataPaymentMethodCashResponse instantiates a new ChargeDataPaymentMethodCashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeDataPaymentMethodCashResponseWithDefaults

`func NewChargeDataPaymentMethodCashResponseWithDefaults() *ChargeDataPaymentMethodCashResponse`

NewChargeDataPaymentMethodCashResponseWithDefaults instantiates a new ChargeDataPaymentMethodCashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *ChargeDataPaymentMethodCashResponse) GetAuthCode() int32`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeDataPaymentMethodCashResponse) GetAuthCodeOk() (*int32, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeDataPaymentMethodCashResponse) SetAuthCode(v int32)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeDataPaymentMethodCashResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *ChargeDataPaymentMethodCashResponse) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *ChargeDataPaymentMethodCashResponse) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetCashierId

`func (o *ChargeDataPaymentMethodCashResponse) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *ChargeDataPaymentMethodCashResponse) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *ChargeDataPaymentMethodCashResponse) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *ChargeDataPaymentMethodCashResponse) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *ChargeDataPaymentMethodCashResponse) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *ChargeDataPaymentMethodCashResponse) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *ChargeDataPaymentMethodCashResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ChargeDataPaymentMethodCashResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ChargeDataPaymentMethodCashResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ChargeDataPaymentMethodCashResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *ChargeDataPaymentMethodCashResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *ChargeDataPaymentMethodCashResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *ChargeDataPaymentMethodCashResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *ChargeDataPaymentMethodCashResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ChargeDataPaymentMethodCashResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeDataPaymentMethodCashResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeDataPaymentMethodCashResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeDataPaymentMethodCashResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetServiceName

`func (o *ChargeDataPaymentMethodCashResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ChargeDataPaymentMethodCashResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ChargeDataPaymentMethodCashResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ChargeDataPaymentMethodCashResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *ChargeDataPaymentMethodCashResponse) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChargeDataPaymentMethodCashResponse) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChargeDataPaymentMethodCashResponse) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChargeDataPaymentMethodCashResponse) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ChargeDataPaymentMethodCashResponse) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ChargeDataPaymentMethodCashResponse) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *ChargeDataPaymentMethodCashResponse) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ChargeDataPaymentMethodCashResponse) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ChargeDataPaymentMethodCashResponse) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *ChargeDataPaymentMethodCashResponse) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


