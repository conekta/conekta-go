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
)

// checks if the ChargeUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeUpdateRequest{}

// ChargeUpdateRequest requested field for update a charge
type ChargeUpdateRequest struct {
	// custom reference id
	ReferenceId *string `json:"reference_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChargeUpdateRequest ChargeUpdateRequest

// NewChargeUpdateRequest instantiates a new ChargeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeUpdateRequest() *ChargeUpdateRequest {
	this := ChargeUpdateRequest{}
	return &this
}

// NewChargeUpdateRequestWithDefaults instantiates a new ChargeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeUpdateRequestWithDefaults() *ChargeUpdateRequest {
	this := ChargeUpdateRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *ChargeUpdateRequest) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeUpdateRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeUpdateRequest) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *ChargeUpdateRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

func (o ChargeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChargeUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varChargeUpdateRequest := _ChargeUpdateRequest{}

	err = json.Unmarshal(data, &varChargeUpdateRequest)

	if err != nil {
		return err
	}

	*o = ChargeUpdateRequest(varChargeUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reference_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChargeUpdateRequest struct {
	value *ChargeUpdateRequest
	isSet bool
}

func (v NullableChargeUpdateRequest) Get() *ChargeUpdateRequest {
	return v.value
}

func (v *NullableChargeUpdateRequest) Set(val *ChargeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeUpdateRequest(val *ChargeUpdateRequest) *NullableChargeUpdateRequest {
	return &NullableChargeUpdateRequest{value: val, isSet: true}
}

func (v NullableChargeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


