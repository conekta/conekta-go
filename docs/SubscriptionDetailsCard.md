# SubscriptionDetailsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ExpMonth** | Pointer to **string** |  | [optional] 
**ExpYear** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PaymentSourceStatus** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**CustomerCustomReference** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionDetailsCard

`func NewSubscriptionDetailsCard() *SubscriptionDetailsCard`

NewSubscriptionDetailsCard instantiates a new SubscriptionDetailsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDetailsCardWithDefaults

`func NewSubscriptionDetailsCardWithDefaults() *SubscriptionDetailsCard`

NewSubscriptionDetailsCardWithDefaults instantiates a new SubscriptionDetailsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionDetailsCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDetailsCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDetailsCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionDetailsCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionDetailsCard) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionDetailsCard) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionDetailsCard) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionDetailsCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetActive

`func (o *SubscriptionDetailsCard) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubscriptionDetailsCard) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubscriptionDetailsCard) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubscriptionDetailsCard) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionDetailsCard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionDetailsCard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionDetailsCard) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionDetailsCard) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetExpMonth

`func (o *SubscriptionDetailsCard) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *SubscriptionDetailsCard) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *SubscriptionDetailsCard) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *SubscriptionDetailsCard) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *SubscriptionDetailsCard) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *SubscriptionDetailsCard) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *SubscriptionDetailsCard) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *SubscriptionDetailsCard) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetBrand

`func (o *SubscriptionDetailsCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *SubscriptionDetailsCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *SubscriptionDetailsCard) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *SubscriptionDetailsCard) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetLast4

`func (o *SubscriptionDetailsCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *SubscriptionDetailsCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *SubscriptionDetailsCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *SubscriptionDetailsCard) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionDetailsCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionDetailsCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionDetailsCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionDetailsCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentSourceStatus

`func (o *SubscriptionDetailsCard) GetPaymentSourceStatus() string`

GetPaymentSourceStatus returns the PaymentSourceStatus field if non-nil, zero value otherwise.

### GetPaymentSourceStatusOk

`func (o *SubscriptionDetailsCard) GetPaymentSourceStatusOk() (*string, bool)`

GetPaymentSourceStatusOk returns a tuple with the PaymentSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceStatus

`func (o *SubscriptionDetailsCard) SetPaymentSourceStatus(v string)`

SetPaymentSourceStatus sets PaymentSourceStatus field to given value.

### HasPaymentSourceStatus

`func (o *SubscriptionDetailsCard) HasPaymentSourceStatus() bool`

HasPaymentSourceStatus returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionDetailsCard) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionDetailsCard) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionDetailsCard) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionDetailsCard) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerCustomReference

`func (o *SubscriptionDetailsCard) GetCustomerCustomReference() string`

GetCustomerCustomReference returns the CustomerCustomReference field if non-nil, zero value otherwise.

### GetCustomerCustomReferenceOk

`func (o *SubscriptionDetailsCard) GetCustomerCustomReferenceOk() (*string, bool)`

GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomReference

`func (o *SubscriptionDetailsCard) SetCustomerCustomReference(v string)`

SetCustomerCustomReference sets CustomerCustomReference field to given value.

### HasCustomerCustomReference

`func (o *SubscriptionDetailsCard) HasCustomerCustomReference() bool`

HasCustomerCustomReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


