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

// checks if the OrderFiscalEntityAddressResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderFiscalEntityAddressResponseAllOf{}

// OrderFiscalEntityAddressResponseAllOf struct for OrderFiscalEntityAddressResponseAllOf
type OrderFiscalEntityAddressResponseAllOf struct {
	Object *string `json:"object,omitempty"`
}

// NewOrderFiscalEntityAddressResponseAllOf instantiates a new OrderFiscalEntityAddressResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFiscalEntityAddressResponseAllOf() *OrderFiscalEntityAddressResponseAllOf {
	this := OrderFiscalEntityAddressResponseAllOf{}
	return &this
}

// NewOrderFiscalEntityAddressResponseAllOfWithDefaults instantiates a new OrderFiscalEntityAddressResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFiscalEntityAddressResponseAllOfWithDefaults() *OrderFiscalEntityAddressResponseAllOf {
	this := OrderFiscalEntityAddressResponseAllOf{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderFiscalEntityAddressResponseAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFiscalEntityAddressResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderFiscalEntityAddressResponseAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderFiscalEntityAddressResponseAllOf) SetObject(v string) {
	o.Object = &v
}

func (o OrderFiscalEntityAddressResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderFiscalEntityAddressResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderFiscalEntityAddressResponseAllOf struct {
	value *OrderFiscalEntityAddressResponseAllOf
	isSet bool
}

func (v NullableOrderFiscalEntityAddressResponseAllOf) Get() *OrderFiscalEntityAddressResponseAllOf {
	return v.value
}

func (v *NullableOrderFiscalEntityAddressResponseAllOf) Set(val *OrderFiscalEntityAddressResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFiscalEntityAddressResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFiscalEntityAddressResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFiscalEntityAddressResponseAllOf(val *OrderFiscalEntityAddressResponseAllOf) *NullableOrderFiscalEntityAddressResponseAllOf {
	return &NullableOrderFiscalEntityAddressResponseAllOf{value: val, isSet: true}
}

func (v NullableOrderFiscalEntityAddressResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFiscalEntityAddressResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


