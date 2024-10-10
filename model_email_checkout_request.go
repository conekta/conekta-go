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
	"fmt"
)

// checks if the EmailCheckoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailCheckoutRequest{}

// EmailCheckoutRequest struct for EmailCheckoutRequest
type EmailCheckoutRequest struct {
	Email string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _EmailCheckoutRequest EmailCheckoutRequest

// NewEmailCheckoutRequest instantiates a new EmailCheckoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCheckoutRequest(email string) *EmailCheckoutRequest {
	this := EmailCheckoutRequest{}
	this.Email = email
	return &this
}

// NewEmailCheckoutRequestWithDefaults instantiates a new EmailCheckoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCheckoutRequestWithDefaults() *EmailCheckoutRequest {
	this := EmailCheckoutRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailCheckoutRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailCheckoutRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailCheckoutRequest) SetEmail(v string) {
	o.Email = v
}

func (o EmailCheckoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailCheckoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailCheckoutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varEmailCheckoutRequest := _EmailCheckoutRequest{}

	err = json.Unmarshal(data, &varEmailCheckoutRequest)

	if err != nil {
		return err
	}

	*o = EmailCheckoutRequest(varEmailCheckoutRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailCheckoutRequest struct {
	value *EmailCheckoutRequest
	isSet bool
}

func (v NullableEmailCheckoutRequest) Get() *EmailCheckoutRequest {
	return v.value
}

func (v *NullableEmailCheckoutRequest) Set(val *EmailCheckoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCheckoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCheckoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCheckoutRequest(val *EmailCheckoutRequest) *NullableEmailCheckoutRequest {
	return &NullableEmailCheckoutRequest{value: val, isSet: true}
}

func (v NullableEmailCheckoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCheckoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


