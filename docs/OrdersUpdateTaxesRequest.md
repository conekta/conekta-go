# OrdersUpdateTaxesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount to be collected for tax in cents | [optional] 
**Description** | Pointer to **string** | description or tax&#39;s name | [optional] 
**Metadata** | Pointer to  |  | [optional] 

## Methods

### NewOrdersUpdateTaxesRequest

`func NewOrdersUpdateTaxesRequest() *OrdersUpdateTaxesRequest`

NewOrdersUpdateTaxesRequest instantiates a new OrdersUpdateTaxesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersUpdateTaxesRequestWithDefaults

`func NewOrdersUpdateTaxesRequestWithDefaults() *OrdersUpdateTaxesRequest`

NewOrdersUpdateTaxesRequestWithDefaults instantiates a new OrdersUpdateTaxesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrdersUpdateTaxesRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrdersUpdateTaxesRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrdersUpdateTaxesRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrdersUpdateTaxesRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *OrdersUpdateTaxesRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrdersUpdateTaxesRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrdersUpdateTaxesRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrdersUpdateTaxesRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *OrdersUpdateTaxesRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrdersUpdateTaxesRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrdersUpdateTaxesRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrdersUpdateTaxesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *OrdersUpdateTaxesRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *OrdersUpdateTaxesRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


