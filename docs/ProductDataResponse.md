# ProductDataResponse

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

### NewProductDataResponse

`func NewProductDataResponse(name string, quantity int32, unitPrice int32, ) *ProductDataResponse`

NewProductDataResponse instantiates a new ProductDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDataResponseWithDefaults

`func NewProductDataResponseWithDefaults() *ProductDataResponse`

NewProductDataResponseWithDefaults instantiates a new ProductDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntifraudInfo

`func (o *ProductDataResponse) GetAntifraudInfo() map[string]interface{}`

GetAntifraudInfo returns the AntifraudInfo field if non-nil, zero value otherwise.

### GetAntifraudInfoOk

`func (o *ProductDataResponse) GetAntifraudInfoOk() (*map[string]interface{}, bool)`

GetAntifraudInfoOk returns a tuple with the AntifraudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntifraudInfo

`func (o *ProductDataResponse) SetAntifraudInfo(v map[string]interface{})`

SetAntifraudInfo sets AntifraudInfo field to given value.

### HasAntifraudInfo

`func (o *ProductDataResponse) HasAntifraudInfo() bool`

HasAntifraudInfo returns a boolean if a field has been set.

### GetBrand

`func (o *ProductDataResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ProductDataResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ProductDataResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ProductDataResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDescription

`func (o *ProductDataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductDataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductDataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductDataResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *ProductDataResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductDataResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductDataResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductDataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ProductDataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductDataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductDataResponse) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *ProductDataResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductDataResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductDataResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSku

`func (o *ProductDataResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductDataResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductDataResponse) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductDataResponse) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTags

`func (o *ProductDataResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductDataResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductDataResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductDataResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnitPrice

`func (o *ProductDataResponse) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ProductDataResponse) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ProductDataResponse) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetId

`func (o *ProductDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductDataResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductDataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ProductDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ProductDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ProductDataResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ProductDataResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParentId

`func (o *ProductDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


