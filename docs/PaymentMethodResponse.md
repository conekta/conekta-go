# PaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodResponse

`func NewPaymentMethodResponse(type_ string, id string, object string, createdAt int64, ) *PaymentMethodResponse`

NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponseWithDefaults

`func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse`

NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaymentMethodResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethodResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *PaymentMethodResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *PaymentMethodResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentMethodResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentMethodResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentMethodResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


