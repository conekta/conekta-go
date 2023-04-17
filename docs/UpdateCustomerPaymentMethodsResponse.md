# UpdateCustomerPaymentMethodsResponse

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
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Bin** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**ExpMonth** | Pointer to **string** |  | [optional] 
**ExpYear** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**VisibleOnCheckout** | Pointer to **bool** |  | [optional] 
**PaymentSourceStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateCustomerPaymentMethodsResponse

`func NewUpdateCustomerPaymentMethodsResponse(type_ string, id string, object string, createdAt int64, ) *UpdateCustomerPaymentMethodsResponse`

NewUpdateCustomerPaymentMethodsResponse instantiates a new UpdateCustomerPaymentMethodsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerPaymentMethodsResponseWithDefaults

`func NewUpdateCustomerPaymentMethodsResponseWithDefaults() *UpdateCustomerPaymentMethodsResponse`

NewUpdateCustomerPaymentMethodsResponseWithDefaults instantiates a new UpdateCustomerPaymentMethodsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateCustomerPaymentMethodsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCustomerPaymentMethodsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UpdateCustomerPaymentMethodsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomerPaymentMethodsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *UpdateCustomerPaymentMethodsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UpdateCustomerPaymentMethodsResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *UpdateCustomerPaymentMethodsResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateCustomerPaymentMethodsResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *UpdateCustomerPaymentMethodsResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateCustomerPaymentMethodsResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateCustomerPaymentMethodsResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReference

`func (o *UpdateCustomerPaymentMethodsResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateCustomerPaymentMethodsResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateCustomerPaymentMethodsResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcode

`func (o *UpdateCustomerPaymentMethodsResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *UpdateCustomerPaymentMethodsResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *UpdateCustomerPaymentMethodsResponse) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *UpdateCustomerPaymentMethodsResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *UpdateCustomerPaymentMethodsResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *UpdateCustomerPaymentMethodsResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateCustomerPaymentMethodsResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateCustomerPaymentMethodsResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProvider

`func (o *UpdateCustomerPaymentMethodsResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UpdateCustomerPaymentMethodsResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UpdateCustomerPaymentMethodsResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLast4

`func (o *UpdateCustomerPaymentMethodsResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *UpdateCustomerPaymentMethodsResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *UpdateCustomerPaymentMethodsResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *UpdateCustomerPaymentMethodsResponse) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetBin

`func (o *UpdateCustomerPaymentMethodsResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *UpdateCustomerPaymentMethodsResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *UpdateCustomerPaymentMethodsResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardType

`func (o *UpdateCustomerPaymentMethodsResponse) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *UpdateCustomerPaymentMethodsResponse) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *UpdateCustomerPaymentMethodsResponse) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpMonth

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *UpdateCustomerPaymentMethodsResponse) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *UpdateCustomerPaymentMethodsResponse) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *UpdateCustomerPaymentMethodsResponse) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *UpdateCustomerPaymentMethodsResponse) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetBrand

`func (o *UpdateCustomerPaymentMethodsResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UpdateCustomerPaymentMethodsResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UpdateCustomerPaymentMethodsResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomerPaymentMethodsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomerPaymentMethodsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomerPaymentMethodsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefault

`func (o *UpdateCustomerPaymentMethodsResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UpdateCustomerPaymentMethodsResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UpdateCustomerPaymentMethodsResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetVisibleOnCheckout

`func (o *UpdateCustomerPaymentMethodsResponse) GetVisibleOnCheckout() bool`

GetVisibleOnCheckout returns the VisibleOnCheckout field if non-nil, zero value otherwise.

### GetVisibleOnCheckoutOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetVisibleOnCheckoutOk() (*bool, bool)`

GetVisibleOnCheckoutOk returns a tuple with the VisibleOnCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCheckout

`func (o *UpdateCustomerPaymentMethodsResponse) SetVisibleOnCheckout(v bool)`

SetVisibleOnCheckout sets VisibleOnCheckout field to given value.

### HasVisibleOnCheckout

`func (o *UpdateCustomerPaymentMethodsResponse) HasVisibleOnCheckout() bool`

HasVisibleOnCheckout returns a boolean if a field has been set.

### GetPaymentSourceStatus

`func (o *UpdateCustomerPaymentMethodsResponse) GetPaymentSourceStatus() string`

GetPaymentSourceStatus returns the PaymentSourceStatus field if non-nil, zero value otherwise.

### GetPaymentSourceStatusOk

`func (o *UpdateCustomerPaymentMethodsResponse) GetPaymentSourceStatusOk() (*string, bool)`

GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceStatus

`func (o *UpdateCustomerPaymentMethodsResponse) SetPaymentSourceStatus(v string)`

SetPaymentSourceStatus sets PaymentSourceStatus field to given value.

### HasPaymentSourceStatus

`func (o *UpdateCustomerPaymentMethodsResponse) HasPaymentSourceStatus() bool`

HasPaymentSourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


