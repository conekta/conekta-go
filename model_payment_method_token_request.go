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

// checks if the PaymentMethodTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodTokenRequest{}

// PaymentMethodTokenRequest struct for PaymentMethodTokenRequest
type PaymentMethodTokenRequest struct {
	// Type of payment method
	Type string `json:"type"`
	// Token id that will be used to create a \"card\" type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards.
	TokenId string `json:"token_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodTokenRequest PaymentMethodTokenRequest

// NewPaymentMethodTokenRequest instantiates a new PaymentMethodTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodTokenRequest(type_ string, tokenId string) *PaymentMethodTokenRequest {
	this := PaymentMethodTokenRequest{}
	this.Type = type_
	this.TokenId = tokenId
	return &this
}

// NewPaymentMethodTokenRequestWithDefaults instantiates a new PaymentMethodTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodTokenRequestWithDefaults() *PaymentMethodTokenRequest {
	this := PaymentMethodTokenRequest{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodTokenRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodTokenRequest) SetType(v string) {
	o.Type = v
}

// GetTokenId returns the TokenId field value
func (o *PaymentMethodTokenRequest) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenRequest) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *PaymentMethodTokenRequest) SetTokenId(v string) {
	o.TokenId = v
}

func (o PaymentMethodTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["token_id"] = o.TokenId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"token_id",
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

	varPaymentMethodTokenRequest := _PaymentMethodTokenRequest{}

	err = json.Unmarshal(data, &varPaymentMethodTokenRequest)

	if err != nil {
		return err
	}

	*o = PaymentMethodTokenRequest(varPaymentMethodTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodTokenRequest struct {
	value *PaymentMethodTokenRequest
	isSet bool
}

func (v NullablePaymentMethodTokenRequest) Get() *PaymentMethodTokenRequest {
	return v.value
}

func (v *NullablePaymentMethodTokenRequest) Set(val *PaymentMethodTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodTokenRequest(val *PaymentMethodTokenRequest) *NullablePaymentMethodTokenRequest {
	return &NullablePaymentMethodTokenRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

