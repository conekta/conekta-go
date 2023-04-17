# PaymentMethodCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
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

### NewPaymentMethodCardResponse

`func NewPaymentMethodCardResponse(type_ string, id string, object string, createdAt int64, ) *PaymentMethodCardResponse`

NewPaymentMethodCardResponse instantiates a new PaymentMethodCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardResponseWithDefaults

`func NewPaymentMethodCardResponseWithDefaults() *PaymentMethodCardResponse`

NewPaymentMethodCardResponseWithDefaults instantiates a new PaymentMethodCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCardResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCardResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCardResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodCardResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodCardResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodCardResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *PaymentMethodCardResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodCardResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodCardResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *PaymentMethodCardResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentMethodCardResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentMethodCardResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentMethodCardResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetLast4

`func (o *PaymentMethodCardResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodCardResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodCardResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *PaymentMethodCardResponse) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetBin

`func (o *PaymentMethodCardResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PaymentMethodCardResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PaymentMethodCardResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *PaymentMethodCardResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethodCardResponse) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodCardResponse) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodCardResponse) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodCardResponse) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpMonth

`func (o *PaymentMethodCardResponse) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethodCardResponse) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethodCardResponse) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *PaymentMethodCardResponse) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *PaymentMethodCardResponse) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethodCardResponse) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethodCardResponse) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *PaymentMethodCardResponse) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetBrand

`func (o *PaymentMethodCardResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethodCardResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethodCardResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaymentMethodCardResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethodCardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodCardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodCardResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethodCardResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefault

`func (o *PaymentMethodCardResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PaymentMethodCardResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PaymentMethodCardResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PaymentMethodCardResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetVisibleOnCheckout

`func (o *PaymentMethodCardResponse) GetVisibleOnCheckout() bool`

GetVisibleOnCheckout returns the VisibleOnCheckout field if non-nil, zero value otherwise.

### GetVisibleOnCheckoutOk

`func (o *PaymentMethodCardResponse) GetVisibleOnCheckoutOk() (*bool, bool)`

GetVisibleOnCheckoutOk returns a tuple with the VisibleOnCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCheckout

`func (o *PaymentMethodCardResponse) SetVisibleOnCheckout(v bool)`

SetVisibleOnCheckout sets VisibleOnCheckout field to given value.

### HasVisibleOnCheckout

`func (o *PaymentMethodCardResponse) HasVisibleOnCheckout() bool`

HasVisibleOnCheckout returns a boolean if a field has been set.

### GetPaymentSourceStatus

`func (o *PaymentMethodCardResponse) GetPaymentSourceStatus() string`

GetPaymentSourceStatus returns the PaymentSourceStatus field if non-nil, zero value otherwise.

### GetPaymentSourceStatusOk

`func (o *PaymentMethodCardResponse) GetPaymentSourceStatusOk() (*string, bool)`

GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceStatus

`func (o *PaymentMethodCardResponse) SetPaymentSourceStatus(v string)`

SetPaymentSourceStatus sets PaymentSourceStatus field to given value.

### HasPaymentSourceStatus

`func (o *PaymentMethodCardResponse) HasPaymentSourceStatus() bool`

HasPaymentSourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


