# CustomerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Corporate** | Pointer to **bool** |  | [optional] [default to false]
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerInfoResponse

`func NewCustomerInfoResponse() *CustomerInfoResponse`

NewCustomerInfoResponse instantiates a new CustomerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInfoResponseWithDefaults

`func NewCustomerInfoResponseWithDefaults() *CustomerInfoResponse`

NewCustomerInfoResponseWithDefaults instantiates a new CustomerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerInfoResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerInfoResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerInfoResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerInfoResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerInfoResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerInfoResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCorporate

`func (o *CustomerInfoResponse) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *CustomerInfoResponse) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *CustomerInfoResponse) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *CustomerInfoResponse) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *CustomerInfoResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerInfoResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerInfoResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerInfoResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


