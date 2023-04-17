# PaymentMethodCash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**AuthCode** | Pointer to **NullableInt32** |  | [optional] 
**CashierId** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodCash

`func NewPaymentMethodCash(object string, ) *PaymentMethodCash`

NewPaymentMethodCash instantiates a new PaymentMethodCash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCashWithDefaults

`func NewPaymentMethodCashWithDefaults() *PaymentMethodCash`

NewPaymentMethodCashWithDefaults instantiates a new PaymentMethodCash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCash) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCash) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCash) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodCash) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodCash) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodCash) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodCash) SetObject(v string)`

SetObject sets Object field to given value.


### GetAuthCode

`func (o *PaymentMethodCash) GetAuthCode() int32`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentMethodCash) GetAuthCodeOk() (*int32, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentMethodCash) SetAuthCode(v int32)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaymentMethodCash) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *PaymentMethodCash) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *PaymentMethodCash) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetCashierId

`func (o *PaymentMethodCash) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *PaymentMethodCash) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *PaymentMethodCash) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *PaymentMethodCash) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *PaymentMethodCash) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *PaymentMethodCash) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *PaymentMethodCash) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodCash) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodCash) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentMethodCash) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *PaymentMethodCash) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *PaymentMethodCash) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *PaymentMethodCash) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *PaymentMethodCash) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodCash) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCash) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCash) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodCash) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetServiceName

`func (o *PaymentMethodCash) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PaymentMethodCash) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PaymentMethodCash) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PaymentMethodCash) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *PaymentMethodCash) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentMethodCash) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentMethodCash) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentMethodCash) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *PaymentMethodCash) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *PaymentMethodCash) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *PaymentMethodCash) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *PaymentMethodCash) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *PaymentMethodCash) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *PaymentMethodCash) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


