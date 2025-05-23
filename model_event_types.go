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

// EventTypes It is a parameter that allows to identify in the response, the type of event that is being generated.
type EventTypes string

// List of event_types
const (
	WEBHOOK_PING EventTypes = "webhook_ping"
	ORDER_PAID EventTypes = "order.paid"
	ORDER_EXPIRED EventTypes = "order.expired"
	ORDER_CANCELED EventTypes = "order.canceled"
	ORDER_PENDING_PAYMENT EventTypes = "order.pending_payment"
)

// All allowed values of EventTypes enum
var AllowedEventTypesEnumValues = []EventTypes{
	"webhook_ping",
	"order.paid",
	"order.expired",
	"order.canceled",
	"order.pending_payment",
}

func (v *EventTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypes(value)
	for _, existing := range AllowedEventTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypes", value)
}

// NewEventTypesFromValue returns a pointer to a valid EventTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypesFromValue(v string) (*EventTypes, error) {
	ev := EventTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypes: valid values are %v", v, AllowedEventTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypes) IsValid() bool {
	for _, existing := range AllowedEventTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_types value
func (v EventTypes) Ptr() *EventTypes {
	return &v
}

type NullableEventTypes struct {
	value *EventTypes
	isSet bool
}

func (v NullableEventTypes) Get() *EventTypes {
	return v.value
}

func (v *NullableEventTypes) Set(val *EventTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypes(val *EventTypes) *NullableEventTypes {
	return &NullableEventTypes{value: val, isSet: true}
}

func (v NullableEventTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

