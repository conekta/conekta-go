# ProductOrderResponse

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
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewProductOrderResponse

`func NewProductOrderResponse(name string, quantity int32, unitPrice int32, ) *ProductOrderResponse`

NewProductOrderResponse instantiates a new ProductOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOrderResponseWithDefaults

`func NewProductOrderResponseWithDefaults() *ProductOrderResponse`

NewProductOrderResponseWithDefaults instantiates a new ProductOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *ProductOrderResponse) GetAntifraudInfo() map[string]interface{}`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *ProductOrderResponse) GetAntifraudInfoOk() (*map[string]interface{}, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *ProductOrderResponse) SetAntifraudInfo(v map[string]interface{})`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *ProductOrderResponse) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### GetBrand

`func (o *ProductOrderResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ProductOrderResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ProductOrderResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ProductOrderResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDescription

`func (o *ProductOrderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductOrderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductOrderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductOrderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *ProductOrderResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductOrderResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductOrderResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductOrderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ProductOrderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductOrderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductOrderResponse) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *ProductOrderResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductOrderResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductOrderResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSku

`func (o *ProductOrderResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductOrderResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductOrderResponse) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductOrderResponse) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTags

`func (o *ProductOrderResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductOrderResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductOrderResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductOrderResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnitPrice

`func (o *ProductOrderResponse) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ProductOrderResponse) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ProductOrderResponse) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetId

`func (o *ProductOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ProductOrderResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProductOrderResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProductOrderResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProductOrderResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParentId

`func (o *ProductOrderResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductOrderResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductOrderResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductOrderResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


