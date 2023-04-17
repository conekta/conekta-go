# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntifraudInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Brand** | Pointer to **string** | The brand of the item. | [optional] 
**Description** | Pointer to **string** | Short description of the item | [optional] 
**Metadata** | Pointer to **map[string]string** | It is a key/value hash that can hold custom fields. Maximum 100 elements and allows special characters. | [optional] 
**Name** | **string** | The name of the item. It will be displayed in the order. | 
**Quantity** | **int32** | The quantity of the item in the order. | 
**Sku** | Pointer to **string** | The stock keeping unit for the item. It is used to identify the item in the order. | [optional] 
**Tags** | Pointer to **[]string** | List of tags for the item. It is used to identify the item in the order. | [optional] 
**UnitPrice** | **int32** | The price of the item in cents. | 

## Methods

### NewProduct

`func NewProduct(name string, quantity int32, unitPrice int32, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *Product) GetAntifraudInfo() map[string]interface{}`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *Product) GetAntifraudInfoOk() (*map[string]interface{}, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *Product) SetAntifraudInfo(v map[string]interface{})`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *Product) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### GetBrand

`func (o *Product) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Product) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Product) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Product) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Product) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Product) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Product) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Product) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *Product) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Product) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Product) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSku

`func (o *Product) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Product) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Product) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Product) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTags

`func (o *Product) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Product) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Product) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Product) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnitPrice

`func (o *Product) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Product) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Product) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


