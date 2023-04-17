# ShippingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Shipping amount in cents | 
**Carrier** | Pointer to **string** | Carrier name for the shipment | [optional] 
**TrackingNumber** | Pointer to **string** | Tracking number can be used to track the shipment | [optional] 
**Method** | Pointer to **string** | Method of shipment | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Hash where the user can send additional information for each &#39;shipping&#39;. | [optional] 

## Methods

### NewShippingRequest

`func NewShippingRequest(amount int64, ) *ShippingRequest`

NewShippingRequest instantiates a new ShippingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingRequestWithDefaults

`func NewShippingRequestWithDefaults() *ShippingRequest`

NewShippingRequestWithDefaults instantiates a new ShippingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ShippingRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ShippingRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ShippingRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCarrier

`func (o *ShippingRequest) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ShippingRequest) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ShippingRequest) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ShippingRequest) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ShippingRequest) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ShippingRequest) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ShippingRequest) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ShippingRequest) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetMethod

`func (o *ShippingRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ShippingRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ShippingRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ShippingRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


