# ShippingLinesDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Shipping amount in cents | 
**Carrier** | Pointer to **string** | Carrier name for the shipment | [optional] 
**TrackingNumber** | Pointer to **string** | Tracking number can be used to track the shipment | [optional] 
**Method** | Pointer to **string** | Method of shipment | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Hash where the user can send additional information for each &#39;shipping&#39;. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewShippingLinesDataResponse

`func NewShippingLinesDataResponse(amount int64, ) *ShippingLinesDataResponse`

NewShippingLinesDataResponse instantiates a new ShippingLinesDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingLinesDataResponseWithDefaults

`func NewShippingLinesDataResponseWithDefaults() *ShippingLinesDataResponse`

NewShippingLinesDataResponseWithDefaults instantiates a new ShippingLinesDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ShippingLinesDataResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ShippingLinesDataResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ShippingLinesDataResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCarrier

`func (o *ShippingLinesDataResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ShippingLinesDataResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ShippingLinesDataResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ShippingLinesDataResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ShippingLinesDataResponse) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ShippingLinesDataResponse) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ShippingLinesDataResponse) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ShippingLinesDataResponse) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetMethod

`func (o *ShippingLinesDataResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ShippingLinesDataResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ShippingLinesDataResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ShippingLinesDataResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingLinesDataResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingLinesDataResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingLinesDataResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingLinesDataResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *ShippingLinesDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingLinesDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingLinesDataResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingLinesDataResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ShippingLinesDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ShippingLinesDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ShippingLinesDataResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ShippingLinesDataResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParentId

`func (o *ShippingLinesDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ShippingLinesDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ShippingLinesDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ShippingLinesDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


