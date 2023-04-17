# UpdateProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntifraudInfo** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UnitPrice** | Pointer to **int64** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateProduct

`func NewUpdateProduct() *UpdateProduct`

NewUpdateProduct instantiates a new UpdateProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductWithDefaults

`func NewUpdateProductWithDefaults() *UpdateProduct`

NewUpdateProductWithDefaults instantiates a new UpdateProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *UpdateProduct) GetAntifraudInfo() map[string]map[string]interface{}`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *UpdateProduct) GetAntifraudInfoOk() (*map[string]map[string]interface{}, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *UpdateProduct) SetAntifraudInfo(v map[string]map[string]interface{})`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *UpdateProduct) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSku

`func (o *UpdateProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *UpdateProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *UpdateProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *UpdateProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *UpdateProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnitPrice

`func (o *UpdateProduct) GetUnitPrice() int64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *UpdateProduct) GetUnitPriceOk() (*int64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *UpdateProduct) SetUnitPrice(v int64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *UpdateProduct) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *UpdateProduct) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateProduct) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateProduct) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpdateProduct) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTags

`func (o *UpdateProduct) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateProduct) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateProduct) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateProduct) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBrand

`func (o *UpdateProduct) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UpdateProduct) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UpdateProduct) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UpdateProduct) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateProduct) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateProduct) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateProduct) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateProduct) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


