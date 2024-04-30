# PayoutOrderResponseCustomerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerCustomReference** | Pointer to **NullableString** | Custom reference | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Corporate** | Pointer to **bool** |  | [optional] [default to false]
**Object** | Pointer to **string** |  | [optional] 
**Id** | **string** | The id of the customer. | 

## Methods

### NewPayoutOrderResponseCustomerInfo

`func NewPayoutOrderResponseCustomerInfo(id string, ) *PayoutOrderResponseCustomerInfo`

NewPayoutOrderResponseCustomerInfo instantiates a new PayoutOrderResponseCustomerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOrderResponseCustomerInfoWithDefaults

`func NewPayoutOrderResponseCustomerInfoWithDefaults() *PayoutOrderResponseCustomerInfo`

NewPayoutOrderResponseCustomerInfoWithDefaults instantiates a new PayoutOrderResponseCustomerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerCustomReference

`func (o *PayoutOrderResponseCustomerInfo) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *PayoutOrderResponseCustomerInfo) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *PayoutOrderResponseCustomerInfo) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *PayoutOrderResponseCustomerInfo) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.

### SetCustomerCustomReferenceNil

`func (o *PayoutOrderResponseCustomerInfo) SetCustomerCustomReferenceNil(b bool)`

 SetCustomerCustomReferenceNil sets the value for CustomerCustomReference to be an explicit nil

### UnsetCustomerCustomReference
`func (o *PayoutOrderResponseCustomerInfo) UnsetCustomerCustomReference()`

UnsetCustomerCustomReference ensures that no value is present for CustomerCustomReference, not even an explicit nil
### GetName

`func (o *PayoutOrderResponseCustomerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayoutOrderResponseCustomerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayoutOrderResponseCustomerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PayoutOrderResponseCustomerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *PayoutOrderResponseCustomerInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayoutOrderResponseCustomerInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayoutOrderResponseCustomerInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayoutOrderResponseCustomerInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *PayoutOrderResponseCustomerInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PayoutOrderResponseCustomerInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PayoutOrderResponseCustomerInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PayoutOrderResponseCustomerInfo) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCorporate

`func (o *PayoutOrderResponseCustomerInfo) GetCorporate() bool`

GetCorporate returns the Corporate field if non-nil, zero value otherwise.

### GetCorporateOk

`func (o *PayoutOrderResponseCustomerInfo) GetCorporateOk() (*bool, bool)`

GetCorporateOk returns a tuple with the Corporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporate

`func (o *PayoutOrderResponseCustomerInfo) SetCorporate(v bool)`

SetCorporate sets Corporate field to given value.

### HasCorporate

`func (o *PayoutOrderResponseCustomerInfo) HasCorporate() bool`

HasCorporate returns a boolean if a field has been set.

### GetObject

`func (o *PayoutOrderResponseCustomerInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PayoutOrderResponseCustomerInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PayoutOrderResponseCustomerInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PayoutOrderResponseCustomerInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetId

`func (o *PayoutOrderResponseCustomerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutOrderResponseCustomerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutOrderResponseCustomerInfo) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


