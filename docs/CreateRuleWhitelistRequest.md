# CreateRuleWhitelistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the rule | 
**Field** | **string** | Field to be used for the rule | 
**Value** | **string** | Value to be used for the rule | 

## Methods

### NewCreateRuleWhitelistRequest

`func NewCreateRuleWhitelistRequest(description string, field string, value string, ) *CreateRuleWhitelistRequest`

NewCreateRuleWhitelistRequest instantiates a new CreateRuleWhitelistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleWhitelistRequestWithDefaults

`func NewCreateRuleWhitelistRequestWithDefaults() *CreateRuleWhitelistRequest`

NewCreateRuleWhitelistRequestWithDefaults instantiates a new CreateRuleWhitelistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateRuleWhitelistRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRuleWhitelistRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRuleWhitelistRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetField

`func (o *CreateRuleWhitelistRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CreateRuleWhitelistRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CreateRuleWhitelistRequest) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *CreateRuleWhitelistRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateRuleWhitelistRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateRuleWhitelistRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


