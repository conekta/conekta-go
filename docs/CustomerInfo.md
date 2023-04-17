# CustomerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Phone** | **string** |  | 
**Corporate** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerInfo

`func NewCustomerInfo(name string, email string, phone string, ) *CustomerInfo`

NewCustomerInfo instantiates a new CustomerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInfoWithDefaults

`func NewCustomerInfoWithDefaults() *CustomerInfo`

NewCustomerInfoWithDefaults instantiates a new CustomerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CustomerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *CustomerInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetCorporate

`func (o *CustomerInfo) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *CustomerInfo) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *CustomerInfo) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *CustomerInfo) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *CustomerInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


