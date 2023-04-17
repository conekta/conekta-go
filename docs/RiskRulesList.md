# RiskRulesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]RiskRulesListDataInner**](RiskRulesListDataInner.md) |  | [optional] 

## Methods

### NewRiskRulesList

`func NewRiskRulesList() *RiskRulesList`

NewRiskRulesList instantiates a new RiskRulesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRulesListWithDefaults

`func NewRiskRulesListWithDefaults() *RiskRulesList`

NewRiskRulesListWithDefaults instantiates a new RiskRulesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *RiskRulesList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *RiskRulesList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *RiskRulesList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *RiskRulesList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetObject

`func (o *RiskRulesList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RiskRulesList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RiskRulesList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RiskRulesList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *RiskRulesList) GetData() []RiskRulesListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RiskRulesList) GetDataOk() (*[]RiskRulesListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RiskRulesList) SetData(v []RiskRulesListDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *RiskRulesList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


