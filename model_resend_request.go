/*
Conekta API

Conekta sdk

API version: 2.2.0
Contact: engineering@conekta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conekta

import (
	"encoding/json"
	"fmt"
)

// checks if the ResendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResendRequest{}

// ResendRequest struct for ResendRequest
type ResendRequest struct {
	// webhooks ids to resend event
	WebhooksIds []string `json:"webhooks_ids"`
	AdditionalProperties map[string]interface{}
}

type _ResendRequest ResendRequest

// NewResendRequest instantiates a new ResendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResendRequest(webhooksIds []string) *ResendRequest {
	this := ResendRequest{}
	this.WebhooksIds = webhooksIds
	return &this
}

// NewResendRequestWithDefaults instantiates a new ResendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResendRequestWithDefaults() *ResendRequest {
	this := ResendRequest{}
	return &this
}

// GetWebhooksIds returns the WebhooksIds field value
func (o *ResendRequest) GetWebhooksIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WebhooksIds
}

// GetWebhooksIdsOk returns a tuple with the WebhooksIds field value
// and a boolean to check if the value has been set.
func (o *ResendRequest) GetWebhooksIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhooksIds, true
}

// SetWebhooksIds sets field value
func (o *ResendRequest) SetWebhooksIds(v []string) {
	o.WebhooksIds = v
}

func (o ResendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhooks_ids"] = o.WebhooksIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResendRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhooks_ids",
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

	varResendRequest := _ResendRequest{}

	err = json.Unmarshal(data, &varResendRequest)

	if err != nil {
		return err
	}

	*o = ResendRequest(varResendRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "webhooks_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResendRequest struct {
	value *ResendRequest
	isSet bool
}

func (v NullableResendRequest) Get() *ResendRequest {
	return v.value
}

func (v *NullableResendRequest) Set(val *ResendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendRequest(val *ResendRequest) *NullableResendRequest {
	return &NullableResendRequest{value: val, isSet: true}
}

func (v NullableResendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


