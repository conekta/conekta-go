# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]DetailsError**](DetailsError.md) |  | [optional] 
**LogId** | Pointer to **NullableString** | log id | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewModelError

`func NewModelError() *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ModelError) GetDetails() []DetailsError`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ModelError) GetDetailsOk() (*[]DetailsError, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ModelError) SetDetails(v []DetailsError)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ModelError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLogId

`func (o *ModelError) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *ModelError) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *ModelError) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *ModelError) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### SetLogIdNil

`func (o *ModelError) SetLogIdNil(b bool)`

 SetLogIdNil sets the value for LogId to be an explicit nil

### UnsetLogId
`func (o *ModelError) UnsetLogId()`

UnsetLogId ensures that no value is present for LogId, not even an explicit nil
### GetType

`func (o *ModelError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *ModelError) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelError) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelError) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelError) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


