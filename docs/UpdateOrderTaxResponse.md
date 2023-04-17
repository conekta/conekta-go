# UpdateOrderTaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount to be collected for tax in cents | 
**Description** | **string** | description or tax&#39;s name | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateOrderTaxResponse

`func NewUpdateOrderTaxResponse(amount int64, description string, id string, ) *UpdateOrderTaxResponse`

NewUpdateOrderTaxResponse instantiates a new UpdateOrderTaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderTaxResponseWithDefaults

`func NewUpdateOrderTaxResponseWithDefaults() *UpdateOrderTaxResponse`

NewUpdateOrderTaxResponseWithDefaults instantiates a new UpdateOrderTaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdateOrderTaxResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateOrderTaxResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateOrderTaxResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *UpdateOrderTaxResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrderTaxResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrderTaxResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *UpdateOrderTaxResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateOrderTaxResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateOrderTaxResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateOrderTaxResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *UpdateOrderTaxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrderTaxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrderTaxResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *UpdateOrderTaxResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UpdateOrderTaxResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UpdateOrderTaxResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *UpdateOrderTaxResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParentId

`func (o *UpdateOrderTaxResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateOrderTaxResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateOrderTaxResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateOrderTaxResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


