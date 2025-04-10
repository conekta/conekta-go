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

// checks if the PaymentMethodBnplPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodBnplPayment{}

// PaymentMethodBnplPayment struct for PaymentMethodBnplPayment
type PaymentMethodBnplPayment struct {
	Type *string `json:"type,omitempty"`
	Object string `json:"object"`
	// URL to redirect the customer after a canceled payment
	CancelUrl *string `json:"cancel_url,omitempty"`
	// Expiration date of the charge
	ExpiresAt int64 `json:"expires_at"`
	// URL to redirect the customer after a failed payment
	FailureUrl *string `json:"failure_url,omitempty"`
	// Product type of the charge
	ProductType string `json:"product_type"`
	// URL to redirect the customer to complete the payment
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// URL to redirect the customer after a successful payment
	SuccessUrl *string `json:"success_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodBnplPayment PaymentMethodBnplPayment

// NewPaymentMethodBnplPayment instantiates a new PaymentMethodBnplPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodBnplPayment(object string, expiresAt int64, productType string) *PaymentMethodBnplPayment {
	this := PaymentMethodBnplPayment{}
	this.Object = object
	this.ExpiresAt = expiresAt
	this.ProductType = productType
	return &this
}

// NewPaymentMethodBnplPaymentWithDefaults instantiates a new PaymentMethodBnplPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodBnplPaymentWithDefaults() *PaymentMethodBnplPayment {
	this := PaymentMethodBnplPayment{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodBnplPayment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodBnplPayment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodBnplPayment) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value
func (o *PaymentMethodBnplPayment) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodBnplPayment) SetObject(v string) {
	o.Object = v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *PaymentMethodBnplPayment) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *PaymentMethodBnplPayment) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *PaymentMethodBnplPayment) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *PaymentMethodBnplPayment) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *PaymentMethodBnplPayment) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *PaymentMethodBnplPayment) GetFailureUrl() string {
	if o == nil || IsNil(o.FailureUrl) {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetFailureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *PaymentMethodBnplPayment) HasFailureUrl() bool {
	if o != nil && !IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *PaymentMethodBnplPayment) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetProductType returns the ProductType field value
func (o *PaymentMethodBnplPayment) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *PaymentMethodBnplPayment) SetProductType(v string) {
	o.ProductType = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *PaymentMethodBnplPayment) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PaymentMethodBnplPayment) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *PaymentMethodBnplPayment) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *PaymentMethodBnplPayment) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodBnplPayment) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *PaymentMethodBnplPayment) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *PaymentMethodBnplPayment) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

func (o PaymentMethodBnplPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodBnplPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.CancelUrl) {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	toSerialize["expires_at"] = o.ExpiresAt
	if !IsNil(o.FailureUrl) {
		toSerialize["failure_url"] = o.FailureUrl
	}
	toSerialize["product_type"] = o.ProductType
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["success_url"] = o.SuccessUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodBnplPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"expires_at",
		"product_type",
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

	varPaymentMethodBnplPayment := _PaymentMethodBnplPayment{}

	err = json.Unmarshal(data, &varPaymentMethodBnplPayment)

	if err != nil {
		return err
	}

	*o = PaymentMethodBnplPayment(varPaymentMethodBnplPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "object")
		delete(additionalProperties, "cancel_url")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "failure_url")
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "redirect_url")
		delete(additionalProperties, "success_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodBnplPayment struct {
	value *PaymentMethodBnplPayment
	isSet bool
}

func (v NullablePaymentMethodBnplPayment) Get() *PaymentMethodBnplPayment {
	return v.value
}

func (v *NullablePaymentMethodBnplPayment) Set(val *PaymentMethodBnplPayment) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodBnplPayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodBnplPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodBnplPayment(val *PaymentMethodBnplPayment) *NullablePaymentMethodBnplPayment {
	return &NullablePaymentMethodBnplPayment{value: val, isSet: true}
}

func (v NullablePaymentMethodBnplPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodBnplPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


