# PaymentMethodCashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodCashResponse

`func NewPaymentMethodCashResponse(type_ string, id string, object string, createdAt int64, ) *PaymentMethodCashResponse`

NewPaymentMethodCashResponse instantiates a new PaymentMethodCashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCashResponseWithDefaults

`func NewPaymentMethodCashResponseWithDefaults() *PaymentMethodCashResponse`

NewPaymentMethodCashResponseWithDefaults instantiates a new PaymentMethodCashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCashResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCashResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCashResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodCashResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodCashResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodCashResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodCashResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodCashResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodCashResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *PaymentMethodCashResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodCashResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodCashResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *PaymentMethodCashResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentMethodCashResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentMethodCashResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentMethodCashResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReference

`func (o *PaymentMethodCashResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodCashResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodCashResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentMethodCashResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcode

`func (o *PaymentMethodCashResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PaymentMethodCashResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PaymentMethodCashResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PaymentMethodCashResponse) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *PaymentMethodCashResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *PaymentMethodCashResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *PaymentMethodCashResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *PaymentMethodCashResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodCashResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCashResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCashResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodCashResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentMethodCashResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethodCashResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethodCashResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentMethodCashResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


