# DetailsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**DebugMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewDetailsError

`func NewDetailsError() *DetailsError`

NewDetailsError instantiates a new DetailsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsErrorWithDefaults

`func NewDetailsErrorWithDefaults() *DetailsError`

NewDetailsErrorWithDefaults instantiates a new DetailsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DetailsError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DetailsError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DetailsError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DetailsError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetParam

`func (o *DetailsError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *DetailsError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *DetailsError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *DetailsError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *DetailsError) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *DetailsError) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetMessage

`func (o *DetailsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DetailsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DetailsError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DetailsError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDebugMessage

`func (o *DetailsError) GetDebugMessage() string`

GetDebugMessage returns the DebugMessage field if non-nil, zero value otherwise.

### GetDebugMessageOk

`func (o *DetailsError) GetDebugMessageOk() (*string, bool)`

GetDebugMessageOk returns a tuple with the DebugMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMessage

`func (o *DetailsError) SetDebugMessage(v string)`

SetDebugMessage sets DebugMessage field to given value.

### HasDebugMessage

`func (o *DetailsError) HasDebugMessage() bool`

HasDebugMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


