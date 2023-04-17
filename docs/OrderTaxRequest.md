# OrderTaxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount to be collected for tax in cents | 
**Description** | **string** | description or tax&#39;s name | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderTaxRequest

`func NewOrderTaxRequest(amount int64, description string, ) *OrderTaxRequest`

NewOrderTaxRequest instantiates a new OrderTaxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTaxRequestWithDefaults

`func NewOrderTaxRequestWithDefaults() *OrderTaxRequest`

NewOrderTaxRequestWithDefaults instantiates a new OrderTaxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrderTaxRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderTaxRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderTaxRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *OrderTaxRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderTaxRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderTaxRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *OrderTaxRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderTaxRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderTaxRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderTaxRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


