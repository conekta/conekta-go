/*
Conekta API

Conekta sdk

API version: 2.1.0
Contact: engineering@conekta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conekta

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateRiskRulesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRiskRulesData{}

// CreateRiskRulesData struct for CreateRiskRulesData
type CreateRiskRulesData struct {
	// Description of the rule
	Description string `json:"description"`
	// Field to be used for the rule
	Field string `json:"field"`
	// Value to be used for the rule
	Value string `json:"value"`
}

type _CreateRiskRulesData CreateRiskRulesData

// NewCreateRiskRulesData instantiates a new CreateRiskRulesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRiskRulesData(description string, field string, value string) *CreateRiskRulesData {
	this := CreateRiskRulesData{}
	this.Description = description
	this.Field = field
	this.Value = value
	return &this
}

// NewCreateRiskRulesDataWithDefaults instantiates a new CreateRiskRulesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRiskRulesDataWithDefaults() *CreateRiskRulesData {
	this := CreateRiskRulesData{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateRiskRulesData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRulesData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateRiskRulesData) SetDescription(v string) {
	o.Description = v
}

// GetField returns the Field field value
func (o *CreateRiskRulesData) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRulesData) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *CreateRiskRulesData) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value
func (o *CreateRiskRulesData) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateRiskRulesData) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateRiskRulesData) SetValue(v string) {
	o.Value = v
}

func (o CreateRiskRulesData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRiskRulesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["field"] = o.Field
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CreateRiskRulesData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"field",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRiskRulesData := _CreateRiskRulesData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRiskRulesData)

	if err != nil {
		return err
	}

	*o = CreateRiskRulesData(varCreateRiskRulesData)

	return err
}

type NullableCreateRiskRulesData struct {
	value *CreateRiskRulesData
	isSet bool
}

func (v NullableCreateRiskRulesData) Get() *CreateRiskRulesData {
	return v.value
}

func (v *NullableCreateRiskRulesData) Set(val *CreateRiskRulesData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRiskRulesData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRiskRulesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRiskRulesData(val *CreateRiskRulesData) *NullableCreateRiskRulesData {
	return &NullableCreateRiskRulesData{value: val, isSet: true}
}

func (v NullableCreateRiskRulesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRiskRulesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


