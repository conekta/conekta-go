# CustomerPortalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** | Unique slug identifier for the portal | [optional] 
**SubscriptionId** | Pointer to **string** | Associated subscription ID | [optional] 
**CustomerId** | Pointer to **string** | Associated customer ID | [optional] 
**Livemode** | Pointer to **bool** | Whether this is a live or test mode portal | [optional] 
**Subscription** | Pointer to [**SubscriptionDetails**](SubscriptionDetails.md) |  | [optional] 
**Customer** | Pointer to [**CustomerDetails**](CustomerDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Customer portal ID | [optional] 
**CompanyId** | Pointer to **string** | Associated company ID | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp of creation | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix timestamp of last update | [optional] 
**PortalUrl** | Pointer to **string** | URL to access the customer portal | [optional] 

## Methods

### NewCustomerPortalResponse

`func NewCustomerPortalResponse() *CustomerPortalResponse`

NewCustomerPortalResponse instantiates a new CustomerPortalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPortalResponseWithDefaults

`func NewCustomerPortalResponseWithDefaults() *CustomerPortalResponse`

NewCustomerPortalResponseWithDefaults instantiates a new CustomerPortalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *CustomerPortalResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CustomerPortalResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CustomerPortalResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CustomerPortalResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CustomerPortalResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CustomerPortalResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CustomerPortalResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CustomerPortalResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomerPortalResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerPortalResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerPortalResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerPortalResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetLivemode

`func (o *CustomerPortalResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *CustomerPortalResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *CustomerPortalResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *CustomerPortalResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetSubscription

`func (o *CustomerPortalResponse) GetSubscription() SubscriptionDetails`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CustomerPortalResponse) GetSubscriptionOk() (*SubscriptionDetails, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CustomerPortalResponse) SetSubscription(v SubscriptionDetails)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CustomerPortalResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetCustomer

`func (o *CustomerPortalResponse) GetCustomer() CustomerDetails`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPortalResponse) GetCustomerOk() (*CustomerDetails, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPortalResponse) SetCustomer(v CustomerDetails)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerPortalResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetId

`func (o *CustomerPortalResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPortalResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPortalResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerPortalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompanyId

`func (o *CustomerPortalResponse) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CustomerPortalResponse) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CustomerPortalResponse) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CustomerPortalResponse) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetObject

`func (o *CustomerPortalResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerPortalResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerPortalResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CustomerPortalResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerPortalResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerPortalResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerPortalResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerPortalResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerPortalResponse) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerPortalResponse) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerPortalResponse) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerPortalResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPortalUrl

`func (o *CustomerPortalResponse) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *CustomerPortalResponse) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *CustomerPortalResponse) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *CustomerPortalResponse) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


