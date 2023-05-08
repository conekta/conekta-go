# CompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The child company&#39;s unique identifier | [optional] 
**CreatedAt** | Pointer to **int64** | The resource&#39;s creation date (unix timestamp) | [optional] 
**Name** | Pointer to **string** | The child company&#39;s name | [optional] 
**Object** | Pointer to **string** | The resource&#39;s type | [optional] 
**ParentCompanyId** | Pointer to **string** | Id of the parent company | [optional] 
**UseParentFiscalData** | Pointer to **bool** | Whether the parent company&#39;s fiscal data is to be used for liquidation and tax purposes | [optional] 
**PayoutDestination** | Pointer to [**CompanyPayoutDestinationResponse**](CompanyPayoutDestinationResponse.md) |  | [optional] 
**FiscalInfo** | Pointer to [**CompanyFiscalInfoResponse**](CompanyFiscalInfoResponse.md) |  | [optional] 

## Methods

### NewCompanyResponse

`func NewCompanyResponse() *CompanyResponse`

NewCompanyResponse instantiates a new CompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseWithDefaults

`func NewCompanyResponseWithDefaults() *CompanyResponse`

NewCompanyResponseWithDefaults instantiates a new CompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CompanyResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CompanyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *CompanyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *CompanyResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CompanyResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CompanyResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CompanyResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetParentCompanyId

`func (o *CompanyResponse) GetParentCompanyId() string`

GetParentCompanyId returns the ParentCompanyId field if non-nil, zero value otherwise.

### GetParentCompanyIdOk

`func (o *CompanyResponse) GetParentCompanyIdOk() (*string, bool)`

GetParentCompanyIdOk returns a tuple with the ParentCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompanyId

`func (o *CompanyResponse) SetParentCompanyId(v string)`

SetParentCompanyId sets ParentCompanyId field to given value.

### HasParentCompanyId

`func (o *CompanyResponse) HasParentCompanyId() bool`

HasParentCompanyId returns a boolean if a field has been set.

### GetUseParentFiscalData

`func (o *CompanyResponse) GetUseParentFiscalData() bool`

GetUseParentFiscalData returns the UseParentFiscalData field if non-nil, zero value otherwise.

### GetUseParentFiscalDataOk

`func (o *CompanyResponse) GetUseParentFiscalDataOk() (*bool, bool)`

GetUseParentFiscalDataOk returns a tuple with the UseParentFiscalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParentFiscalData

`func (o *CompanyResponse) SetUseParentFiscalData(v bool)`

SetUseParentFiscalData sets UseParentFiscalData field to given value.

### HasUseParentFiscalData

`func (o *CompanyResponse) HasUseParentFiscalData() bool`

HasUseParentFiscalData returns a boolean if a field has been set.

### GetPayoutDestination

`func (o *CompanyResponse) GetPayoutDestination() CompanyPayoutDestinationResponse`

GetPayoutDestination returns the PayoutDestination field if non-nil, zero value otherwise.

### GetPayoutDestinationOk

`func (o *CompanyResponse) GetPayoutDestinationOk() (*CompanyPayoutDestinationResponse, bool)`

GetPayoutDestinationOk returns a tuple with the PayoutDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutDestination

`func (o *CompanyResponse) SetPayoutDestination(v CompanyPayoutDestinationResponse)`

SetPayoutDestination sets PayoutDestination field to given value.

### HasPayoutDestination

`func (o *CompanyResponse) HasPayoutDestination() bool`

HasPayoutDestination returns a boolean if a field has been set.

### GetFiscalInfo

`func (o *CompanyResponse) GetFiscalInfo() CompanyFiscalInfoResponse`

GetFiscalInfo returns the FiscalInfo field if non-nil, zero value otherwise.

### GetFiscalInfoOk

`func (o *CompanyResponse) GetFiscalInfoOk() (*CompanyFiscalInfoResponse, bool)`

GetFiscalInfoOk returns a tuple with the FiscalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalInfo

`func (o *CompanyResponse) SetFiscalInfo(v CompanyFiscalInfoResponse)`

SetFiscalInfo sets FiscalInfo field to given value.

### HasFiscalInfo

`func (o *CompanyResponse) HasFiscalInfo() bool`

HasFiscalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


