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

// checks if the CustomerAntifraudInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAntifraudInfoResponse{}

// CustomerAntifraudInfoResponse struct for CustomerAntifraudInfoResponse
type CustomerAntifraudInfoResponse struct {
	FirstPaidAt *int32 `json:"first_paid_at,omitempty"`
	AccountCreatedAt *int64 `json:"account_created_at,omitempty"`
}

// NewCustomerAntifraudInfoResponse instantiates a new CustomerAntifraudInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAntifraudInfoResponse() *CustomerAntifraudInfoResponse {
	this := CustomerAntifraudInfoResponse{}
	return &this
}

// NewCustomerAntifraudInfoResponseWithDefaults instantiates a new CustomerAntifraudInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAntifraudInfoResponseWithDefaults() *CustomerAntifraudInfoResponse {
	this := CustomerAntifraudInfoResponse{}
	return &this
}

// GetFirstPaidAt returns the FirstPaidAt field value if set, zero value otherwise.
func (o *CustomerAntifraudInfoResponse) GetFirstPaidAt() int32 {
	if o == nil || IsNil(o.FirstPaidAt) {
		var ret int32
		return ret
	}
	return *o.FirstPaidAt
}

// GetFirstPaidAtOk returns a tuple with the FirstPaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAntifraudInfoResponse) GetFirstPaidAtOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstPaidAt) {
		return nil, false
	}
	return o.FirstPaidAt, true
}

// HasFirstPaidAt returns a boolean if a field has been set.
func (o *CustomerAntifraudInfoResponse) HasFirstPaidAt() bool {
	if o != nil && !IsNil(o.FirstPaidAt) {
		return true
	}

	return false
}

// SetFirstPaidAt gets a reference to the given int32 and assigns it to the FirstPaidAt field.
func (o *CustomerAntifraudInfoResponse) SetFirstPaidAt(v int32) {
	o.FirstPaidAt = &v
}

// GetAccountCreatedAt returns the AccountCreatedAt field value if set, zero value otherwise.
func (o *CustomerAntifraudInfoResponse) GetAccountCreatedAt() int64 {
	if o == nil || IsNil(o.AccountCreatedAt) {
		var ret int64
		return ret
	}
	return *o.AccountCreatedAt
}

// GetAccountCreatedAtOk returns a tuple with the AccountCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAntifraudInfoResponse) GetAccountCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountCreatedAt) {
		return nil, false
	}
	return o.AccountCreatedAt, true
}

// HasAccountCreatedAt returns a boolean if a field has been set.
func (o *CustomerAntifraudInfoResponse) HasAccountCreatedAt() bool {
	if o != nil && !IsNil(o.AccountCreatedAt) {
		return true
	}

	return false
}

// SetAccountCreatedAt gets a reference to the given int64 and assigns it to the AccountCreatedAt field.
func (o *CustomerAntifraudInfoResponse) SetAccountCreatedAt(v int64) {
	o.AccountCreatedAt = &v
}

func (o CustomerAntifraudInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAntifraudInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstPaidAt) {
		toSerialize["first_paid_at"] = o.FirstPaidAt
	}
	if !IsNil(o.AccountCreatedAt) {
		toSerialize["account_created_at"] = o.AccountCreatedAt
	}
	return toSerialize, nil
}

type NullableCustomerAntifraudInfoResponse struct {
	value *CustomerAntifraudInfoResponse
	isSet bool
}

func (v NullableCustomerAntifraudInfoResponse) Get() *CustomerAntifraudInfoResponse {
	return v.value
}

func (v *NullableCustomerAntifraudInfoResponse) Set(val *CustomerAntifraudInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAntifraudInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAntifraudInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAntifraudInfoResponse(val *CustomerAntifraudInfoResponse) *NullableCustomerAntifraudInfoResponse {
	return &NullableCustomerAntifraudInfoResponse{value: val, isSet: true}
}

func (v NullableCustomerAntifraudInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAntifraudInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


