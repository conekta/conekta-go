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

// checks if the CustomerShippingContacts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContacts{}

// CustomerShippingContacts [Shipping](https://developers.conekta.com/v2.2.0/reference/createcustomershippingcontacts) details, required in case of sending a shipping. If we do not receive a shipping_contact on the order, the default shipping_contact of the customer will be used.
type CustomerShippingContacts struct {
	// Phone contact
	Phone *string `json:"phone,omitempty"`
	// Name of the person who will receive the order
	Receiver *string `json:"receiver,omitempty"`
	// The street names between which the order will be delivered.
	BetweenStreets *string `json:"between_streets,omitempty"`
	Address CustomerShippingContactsAddress `json:"address"`
	ParentId *string `json:"parent_id,omitempty"`
	Default NullableBool `json:"default,omitempty"`
	Deleted NullableBool `json:"deleted,omitempty"`
	// Metadata associated with the shipping contact
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerShippingContacts CustomerShippingContacts

// NewCustomerShippingContacts instantiates a new CustomerShippingContacts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContacts(address CustomerShippingContactsAddress) *CustomerShippingContacts {
	this := CustomerShippingContacts{}
	this.Address = address
	return &this
}

// NewCustomerShippingContactsWithDefaults instantiates a new CustomerShippingContacts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsWithDefaults() *CustomerShippingContacts {
	this := CustomerShippingContacts{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerShippingContacts) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerShippingContacts) SetPhone(v string) {
	o.Phone = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *CustomerShippingContacts) GetReceiver() string {
	if o == nil || IsNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetReceiverOk() (*string, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *CustomerShippingContacts) SetReceiver(v string) {
	o.Receiver = &v
}

// GetBetweenStreets returns the BetweenStreets field value if set, zero value otherwise.
func (o *CustomerShippingContacts) GetBetweenStreets() string {
	if o == nil || IsNil(o.BetweenStreets) {
		var ret string
		return ret
	}
	return *o.BetweenStreets
}

// GetBetweenStreetsOk returns a tuple with the BetweenStreets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetBetweenStreetsOk() (*string, bool) {
	if o == nil || IsNil(o.BetweenStreets) {
		return nil, false
	}
	return o.BetweenStreets, true
}

// HasBetweenStreets returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasBetweenStreets() bool {
	if o != nil && !IsNil(o.BetweenStreets) {
		return true
	}

	return false
}

// SetBetweenStreets gets a reference to the given string and assigns it to the BetweenStreets field.
func (o *CustomerShippingContacts) SetBetweenStreets(v string) {
	o.BetweenStreets = &v
}

// GetAddress returns the Address field value
func (o *CustomerShippingContacts) GetAddress() CustomerShippingContactsAddress {
	if o == nil {
		var ret CustomerShippingContactsAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetAddressOk() (*CustomerShippingContactsAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustomerShippingContacts) SetAddress(v CustomerShippingContactsAddress) {
	o.Address = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CustomerShippingContacts) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CustomerShippingContacts) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContacts) GetDefault() bool {
	if o == nil || IsNil(o.Default.Get()) {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContacts) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *CustomerShippingContacts) SetDefault(v bool) {
	o.Default.Set(&v)
}
// SetDefaultNil sets the value for Default to be an explicit nil
func (o *CustomerShippingContacts) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *CustomerShippingContacts) UnsetDefault() {
	o.Default.Unset()
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContacts) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted.Get()) {
		var ret bool
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContacts) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableBool and assigns it to the Deleted field.
func (o *CustomerShippingContacts) SetDeleted(v bool) {
	o.Deleted.Set(&v)
}
// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *CustomerShippingContacts) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *CustomerShippingContacts) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerShippingContacts) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContacts) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerShippingContacts) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerShippingContacts) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CustomerShippingContacts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContacts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	if !IsNil(o.BetweenStreets) {
		toSerialize["between_streets"] = o.BetweenStreets
	}
	toSerialize["address"] = o.Address
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerShippingContacts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
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

	varCustomerShippingContacts := _CustomerShippingContacts{}

	err = json.Unmarshal(data, &varCustomerShippingContacts)

	if err != nil {
		return err
	}

	*o = CustomerShippingContacts(varCustomerShippingContacts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "phone")
		delete(additionalProperties, "receiver")
		delete(additionalProperties, "between_streets")
		delete(additionalProperties, "address")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "default")
		delete(additionalProperties, "deleted")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerShippingContacts struct {
	value *CustomerShippingContacts
	isSet bool
}

func (v NullableCustomerShippingContacts) Get() *CustomerShippingContacts {
	return v.value
}

func (v *NullableCustomerShippingContacts) Set(val *CustomerShippingContacts) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContacts) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContacts(val *CustomerShippingContacts) *NullableCustomerShippingContacts {
	return &NullableCustomerShippingContacts{value: val, isSet: true}
}

func (v NullableCustomerShippingContacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


