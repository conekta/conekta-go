# ChargebackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**FollowupStatus** | Pointer to **string** |  | [optional] 
**ResponseFromClient** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]ChargebackFileResponse**](ChargebackFileResponse.md) |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ChargeId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**EvidenceDueBy** | Pointer to **int64** |  | [optional] 

## Methods

### NewChargebackResponse

`func NewChargebackResponse() *ChargebackResponse`

NewChargebackResponse instantiates a new ChargebackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargebackResponseWithDefaults

`func NewChargebackResponseWithDefaults() *ChargebackResponse`

NewChargebackResponseWithDefaults instantiates a new ChargebackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargebackResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargebackResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargebackResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargebackResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ChargebackResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargebackResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargebackResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChargebackResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ChargebackResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChargebackResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChargebackResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ChargebackResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetNote

`func (o *ChargebackResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ChargebackResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ChargebackResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ChargebackResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetFollowupStatus

`func (o *ChargebackResponse) GetFollowupStatus() string`

GetFollowupStatus returns the FollowupStatus field if non-nil, zero value otherwise.

### GetFollowupStatusOk

`func (o *ChargebackResponse) GetFollowupStatusOk() (*string, bool)`

GetFollowupStatusOk returns a tuple with the FollowupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupStatus

`func (o *ChargebackResponse) SetFollowupStatus(v string)`

SetFollowupStatus sets FollowupStatus field to given value.

### HasFollowupStatus

`func (o *ChargebackResponse) HasFollowupStatus() bool`

HasFollowupStatus returns a boolean if a field has been set.

### GetResponseFromClient

`func (o *ChargebackResponse) GetResponseFromClient() string`

GetResponseFromClient returns the ResponseFromClient field if non-nil, zero value otherwise.

### GetResponseFromClientOk

`func (o *ChargebackResponse) GetResponseFromClientOk() (*string, bool)`

GetResponseFromClientOk returns a tuple with the ResponseFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFromClient

`func (o *ChargebackResponse) SetResponseFromClient(v string)`

SetResponseFromClient sets ResponseFromClient field to given value.

### HasResponseFromClient

`func (o *ChargebackResponse) HasResponseFromClient() bool`

HasResponseFromClient returns a boolean if a field has been set.

### GetFiles

`func (o *ChargebackResponse) GetFiles() []ChargebackFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ChargebackResponse) GetFilesOk() (*[]ChargebackFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ChargebackResponse) SetFiles(v []ChargebackFileResponse)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ChargebackResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetObject

`func (o *ChargebackResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargebackResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargebackResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChargebackResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetChargeId

`func (o *ChargebackResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *ChargebackResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *ChargebackResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *ChargebackResponse) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChargebackResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargebackResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargebackResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChargebackResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvidenceDueBy

`func (o *ChargebackResponse) GetEvidenceDueBy() int64`

GetEvidenceDueBy returns the EvidenceDueBy field if non-nil, zero value otherwise.

### GetEvidenceDueByOk

`func (o *ChargebackResponse) GetEvidenceDueByOk() (*int64, bool)`

GetEvidenceDueByOk returns a tuple with the EvidenceDueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDueBy

`func (o *ChargebackResponse) SetEvidenceDueBy(v int64)`

SetEvidenceDueBy sets EvidenceDueBy field to given value.

### HasEvidenceDueBy

`func (o *ChargebackResponse) HasEvidenceDueBy() bool`

HasEvidenceDueBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


