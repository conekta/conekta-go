# PaymentMethodCashRecurrentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Agreements** | Pointer to [**[]PaymentMethodCashResponseAllOfAgreements**](PaymentMethodCashResponseAllOfAgreements.md) |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** | URL to the barcode image, reference is the same as barcode | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodCashRecurrentResponse

`func NewPaymentMethodCashRecurrentResponse(type_ string, id string, object string, createdAt int64, ) *PaymentMethodCashRecurrentResponse`

NewPaymentMethodCashRecurrentResponse instantiates a new PaymentMethodCashRecurrentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCashRecurrentResponseWithDefaults

`func NewPaymentMethodCashRecurrentResponseWithDefaults() *PaymentMethodCashRecurrentResponse`

NewPaymentMethodCashRecurrentResponseWithDefaults instantiates a new PaymentMethodCashRecurrentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCashRecurrentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCashRecurrentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCashRecurrentResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodCashRecurrentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodCashRecurrentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodCashRecurrentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodCashRecurrentResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodCashRecurrentResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodCashRecurrentResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *PaymentMethodCashRecurrentResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodCashRecurrentResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodCashRecurrentResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *PaymentMethodCashRecurrentResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentMethodCashRecurrentResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentMethodCashRecurrentResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentMethodCashRecurrentResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAgreements

`func (o *PaymentMethodCashRecurrentResponse) GetAgreements() []PaymentMethodCashResponseAllOfAgreements`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *PaymentMethodCashRecurrentResponse) GetAgreementsOk() (*[]PaymentMethodCashResponseAllOfAgreements, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *PaymentMethodCashRecurrentResponse) SetAgreements(v []PaymentMethodCashResponseAllOfAgreements)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *PaymentMethodCashRecurrentResponse) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetReference

`func (o *PaymentMethodCashRecurrentResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentMethodCashRecurrentResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentMethodCashRecurrentResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PaymentMethodCashRecurrentResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcode

`func (o *PaymentMethodCashRecurrentResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PaymentMethodCashRecurrentResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PaymentMethodCashRecurrentResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PaymentMethodCashRecurrentResponse) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *PaymentMethodCashRecurrentResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *PaymentMethodCashRecurrentResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *PaymentMethodCashRecurrentResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *PaymentMethodCashRecurrentResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodCashRecurrentResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCashRecurrentResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCashRecurrentResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodCashRecurrentResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentMethodCashRecurrentResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethodCashRecurrentResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethodCashRecurrentResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentMethodCashRecurrentResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


