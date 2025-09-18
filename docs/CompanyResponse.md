# CompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the company. | 
**Name** | **string** | The name of the company. | 
**Active** | **bool** | Indicates if the company is active. | 
**AccountStatus** | **string** | The current status of the company&#39;s account. | 
**ParentCompanyId** | Pointer to **NullableString** | The identifier of the parent company, if any. | [optional] 
**OnboardingStatus** | **string** | The current status of the company&#39;s onboarding process. | 
**Documents** | [**[]CompanyResponseDocumentsInner**](CompanyResponseDocumentsInner.md) | A list of documents related to the company. | 
**CreatedAt** | **int64** | Timestamp of when the company was created. | 
**Object** | **string** | The type of object, typically \&quot;company\&quot;. | 
**ThreeDsEnabled** | Pointer to **bool** | Indicates if 3DS authentication is enabled for the company. | [optional] 
**ThreeDsMode** | Pointer to **NullableString** | The 3DS mode for the company, either &#39;smart&#39; or &#39;strict&#39;. This property is only applicable when three_ds_enabled is true. When three_ds_enabled is false, this field will be null. | [optional] 

## Methods

### NewCompanyResponse

`func NewCompanyResponse(id string, name string, active bool, accountStatus string, onboardingStatus string, documents []CompanyResponseDocumentsInner, createdAt int64, object string, ) *CompanyResponse`

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


### GetActive

`func (o *CompanyResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CompanyResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CompanyResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAccountStatus

`func (o *CompanyResponse) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *CompanyResponse) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *CompanyResponse) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.


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

### SetParentCompanyIdNil

`func (o *CompanyResponse) SetParentCompanyIdNil(b bool)`

 SetParentCompanyIdNil sets the value for ParentCompanyId to be an explicit nil

### UnsetParentCompanyId
`func (o *CompanyResponse) UnsetParentCompanyId()`

UnsetParentCompanyId ensures that no value is present for ParentCompanyId, not even an explicit nil
### GetOnboardingStatus

`func (o *CompanyResponse) GetOnboardingStatus() string`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *CompanyResponse) GetOnboardingStatusOk() (*string, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingStatus

`func (o *CompanyResponse) SetOnboardingStatus(v string)`

SetOnboardingStatus sets OnboardingStatus field to given value.


### GetDocuments

`func (o *CompanyResponse) GetDocuments() []CompanyResponseDocumentsInner`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CompanyResponse) GetDocumentsOk() (*[]CompanyResponseDocumentsInner, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CompanyResponse) SetDocuments(v []CompanyResponseDocumentsInner)`

SetDocuments sets Documents field to given value.


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


### GetThreeDsEnabled

`func (o *CompanyResponse) GetThreeDsEnabled() bool`

GetThreeDsEnabled returns the ThreeDsEnabled field if non-nil, zero value otherwise.

### GetThreeDsEnabledOk

`func (o *CompanyResponse) GetThreeDsEnabledOk() (*bool, bool)`

GetThreeDsEnabledOk returns a tuple with the ThreeDsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsEnabled

`func (o *CompanyResponse) SetThreeDsEnabled(v bool)`

SetThreeDsEnabled sets ThreeDsEnabled field to given value.

### HasThreeDsEnabled

`func (o *CompanyResponse) HasThreeDsEnabled() bool`

HasThreeDsEnabled returns a boolean if a field has been set.

### GetThreeDsMode

`func (o *CompanyResponse) GetThreeDsMode() string`

GetThreeDsMode returns the ThreeDsMode field if non-nil, zero value otherwise.

### GetThreeDsModeOk

`func (o *CompanyResponse) GetThreeDsModeOk() (*string, bool)`

GetThreeDsModeOk returns a tuple with the ThreeDsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsMode

`func (o *CompanyResponse) SetThreeDsMode(v string)`

SetThreeDsMode sets ThreeDsMode field to given value.

### HasThreeDsMode

`func (o *CompanyResponse) HasThreeDsMode() bool`

HasThreeDsMode returns a boolean if a field has been set.

### SetThreeDsModeNil

`func (o *CompanyResponse) SetThreeDsModeNil(b bool)`

 SetThreeDsModeNil sets the value for ThreeDsMode to be an explicit nil

### UnsetThreeDsMode
`func (o *CompanyResponse) UnsetThreeDsMode()`

UnsetThreeDsMode ensures that no value is present for ThreeDsMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


