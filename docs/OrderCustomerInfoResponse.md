# OrderCustomerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerCustomReference** | Pointer to **NullableString** | Custom reference | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Corporate** | Pointer to **bool** |  | [optional] [default to false]
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderCustomerInfoResponse

`func NewOrderCustomerInfoResponse() *OrderCustomerInfoResponse`

NewOrderCustomerInfoResponse instantiates a new OrderCustomerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCustomerInfoResponseWithDefaults

`func NewOrderCustomerInfoResponseWithDefaults() *OrderCustomerInfoResponse`

NewOrderCustomerInfoResponseWithDefaults instantiates a new OrderCustomerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerCustomReference

`func (o *OrderCustomerInfoResponse) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *OrderCustomerInfoResponse) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *OrderCustomerInfoResponse) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *OrderCustomerInfoResponse) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.

### SetCustomerCustomReferenceNil

`func (o *OrderCustomerInfoResponse) SetCustomerCustomReferenceNil(b bool)`

 SetCustomerCustomReferenceNil sets the value for CustomerCustomReference to be an explicit nil

### UnsetCustomerCustomReference
`func (o *OrderCustomerInfoResponse) UnsetCustomerCustomReference()`

UnsetCustomerCustomReference ensures that no value is present for CustomerCustomReference, not even an explicit nil
### GetName

`func (o *OrderCustomerInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderCustomerInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderCustomerInfoResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderCustomerInfoResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *OrderCustomerInfoResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrderCustomerInfoResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrderCustomerInfoResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrderCustomerInfoResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *OrderCustomerInfoResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrderCustomerInfoResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrderCustomerInfoResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrderCustomerInfoResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCorporate

`func (o *OrderCustomerInfoResponse) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *OrderCustomerInfoResponse) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *OrderCustomerInfoResponse) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *OrderCustomerInfoResponse) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *OrderCustomerInfoResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderCustomerInfoResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderCustomerInfoResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderCustomerInfoResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


