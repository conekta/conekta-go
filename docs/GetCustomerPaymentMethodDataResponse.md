# GetCustomerPaymentMethodDataResponse

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
**Bank** | Pointer to **string** | Bank name for the SPEI payment method | [optional] 

## Methods

### NewGetCustomerPaymentMethodDataResponse

`func NewGetCustomerPaymentMethodDataResponse(type_ string, id string, object string, createdAt int64, ) *GetCustomerPaymentMethodDataResponse`

NewGetCustomerPaymentMethodDataResponse instantiates a new GetCustomerPaymentMethodDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerPaymentMethodDataResponseWithDefaults

`func NewGetCustomerPaymentMethodDataResponseWithDefaults() *GetCustomerPaymentMethodDataResponse`

NewGetCustomerPaymentMethodDataResponseWithDefaults instantiates a new GetCustomerPaymentMethodDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCustomerPaymentMethodDataResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCustomerPaymentMethodDataResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCustomerPaymentMethodDataResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GetCustomerPaymentMethodDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCustomerPaymentMethodDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCustomerPaymentMethodDataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *GetCustomerPaymentMethodDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetCustomerPaymentMethodDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetCustomerPaymentMethodDataResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *GetCustomerPaymentMethodDataResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCustomerPaymentMethodDataResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCustomerPaymentMethodDataResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *GetCustomerPaymentMethodDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetCustomerPaymentMethodDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetCustomerPaymentMethodDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetCustomerPaymentMethodDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetAgreements

`func (o *GetCustomerPaymentMethodDataResponse) GetAgreements() []PaymentMethodCashResponseAllOfAgreements`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *GetCustomerPaymentMethodDataResponse) GetAgreementsOk() (*[]PaymentMethodCashResponseAllOfAgreements, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *GetCustomerPaymentMethodDataResponse) SetAgreements(v []PaymentMethodCashResponseAllOfAgreements)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *GetCustomerPaymentMethodDataResponse) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetReference

`func (o *GetCustomerPaymentMethodDataResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetCustomerPaymentMethodDataResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetCustomerPaymentMethodDataResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetCustomerPaymentMethodDataResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcode

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetCustomerPaymentMethodDataResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *GetCustomerPaymentMethodDataResponse) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetCustomerPaymentMethodDataResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProvider

`func (o *GetCustomerPaymentMethodDataResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetCustomerPaymentMethodDataResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetCustomerPaymentMethodDataResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetCustomerPaymentMethodDataResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetLast4

`func (o *GetCustomerPaymentMethodDataResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *GetCustomerPaymentMethodDataResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *GetCustomerPaymentMethodDataResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *GetCustomerPaymentMethodDataResponse) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetBin

`func (o *GetCustomerPaymentMethodDataResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GetCustomerPaymentMethodDataResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *GetCustomerPaymentMethodDataResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardType

`func (o *GetCustomerPaymentMethodDataResponse) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *GetCustomerPaymentMethodDataResponse) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *GetCustomerPaymentMethodDataResponse) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *GetCustomerPaymentMethodDataResponse) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpMonth

`func (o *GetCustomerPaymentMethodDataResponse) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *GetCustomerPaymentMethodDataResponse) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *GetCustomerPaymentMethodDataResponse) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *GetCustomerPaymentMethodDataResponse) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *GetCustomerPaymentMethodDataResponse) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *GetCustomerPaymentMethodDataResponse) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *GetCustomerPaymentMethodDataResponse) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *GetCustomerPaymentMethodDataResponse) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetBrand

`func (o *GetCustomerPaymentMethodDataResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetCustomerPaymentMethodDataResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *GetCustomerPaymentMethodDataResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetName

`func (o *GetCustomerPaymentMethodDataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCustomerPaymentMethodDataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCustomerPaymentMethodDataResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCustomerPaymentMethodDataResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefault

`func (o *GetCustomerPaymentMethodDataResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GetCustomerPaymentMethodDataResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GetCustomerPaymentMethodDataResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *GetCustomerPaymentMethodDataResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetVisibleOnCheckout

`func (o *GetCustomerPaymentMethodDataResponse) GetVisibleOnCheckout() bool`

GetVisibleOnCheckout returns the VisibleOnCheckout field if non-nil, zero value otherwise.

### GetVisibleOnCheckoutOk

`func (o *GetCustomerPaymentMethodDataResponse) GetVisibleOnCheckoutOk() (*bool, bool)`

GetVisibleOnCheckoutOk returns a tuple with the VisibleOnCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCheckout

`func (o *GetCustomerPaymentMethodDataResponse) SetVisibleOnCheckout(v bool)`

SetVisibleOnCheckout sets VisibleOnCheckout field to given value.

### HasVisibleOnCheckout

`func (o *GetCustomerPaymentMethodDataResponse) HasVisibleOnCheckout() bool`

HasVisibleOnCheckout returns a boolean if a field has been set.

### GetPaymentSourceStatus

`func (o *GetCustomerPaymentMethodDataResponse) GetPaymentSourceStatus() string`

GetPaymentSourceStatus returns the PaymentSourceStatus field if non-nil, zero value otherwise.

### GetPaymentSourceStatusOk

`func (o *GetCustomerPaymentMethodDataResponse) GetPaymentSourceStatusOk() (*string, bool)`

GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceStatus

`func (o *GetCustomerPaymentMethodDataResponse) SetPaymentSourceStatus(v string)`

SetPaymentSourceStatus sets PaymentSourceStatus field to given value.

### HasPaymentSourceStatus

`func (o *GetCustomerPaymentMethodDataResponse) HasPaymentSourceStatus() bool`

HasPaymentSourceStatus returns a boolean if a field has been set.

### GetBank

`func (o *GetCustomerPaymentMethodDataResponse) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *GetCustomerPaymentMethodDataResponse) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *GetCustomerPaymentMethodDataResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


