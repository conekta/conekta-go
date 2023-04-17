# UpdateOrderTaxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount to be collected for tax in cents | [optional] 
**Description** | Pointer to **string** | description or tax&#39;s name | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateOrderTaxRequest

`func NewUpdateOrderTaxRequest() *UpdateOrderTaxRequest`

NewUpdateOrderTaxRequest instantiates a new UpdateOrderTaxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderTaxRequestWithDefaults

`func NewUpdateOrderTaxRequestWithDefaults() *UpdateOrderTaxRequest`

NewUpdateOrderTaxRequestWithDefaults instantiates a new UpdateOrderTaxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdateOrderTaxRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateOrderTaxRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateOrderTaxRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateOrderTaxRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOrderTaxRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOrderTaxRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOrderTaxRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOrderTaxRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateOrderTaxRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateOrderTaxRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateOrderTaxRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateOrderTaxRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


