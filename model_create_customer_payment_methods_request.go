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

// CreateCustomerPaymentMethodsRequest - Contains details of the payment methods that the customer has active or has used in Conekta
type CreateCustomerPaymentMethodsRequest struct {
	PaymentMethodCashRequest *PaymentMethodCashRequest
	PaymentMethodSpeiRequest *PaymentMethodSpeiRequest
	PaymentMethodTokenRequest *PaymentMethodTokenRequest
}

// PaymentMethodCashRequestAsCreateCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodCashRequest wrapped in CreateCustomerPaymentMethodsRequest
func PaymentMethodCashRequestAsCreateCustomerPaymentMethodsRequest(v *PaymentMethodCashRequest) CreateCustomerPaymentMethodsRequest {
	return CreateCustomerPaymentMethodsRequest{
		PaymentMethodCashRequest: v,
	}
}

// PaymentMethodSpeiRequestAsCreateCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodSpeiRequest wrapped in CreateCustomerPaymentMethodsRequest
func PaymentMethodSpeiRequestAsCreateCustomerPaymentMethodsRequest(v *PaymentMethodSpeiRequest) CreateCustomerPaymentMethodsRequest {
	return CreateCustomerPaymentMethodsRequest{
		PaymentMethodSpeiRequest: v,
	}
}

// PaymentMethodTokenRequestAsCreateCustomerPaymentMethodsRequest is a convenience function that returns PaymentMethodTokenRequest wrapped in CreateCustomerPaymentMethodsRequest
func PaymentMethodTokenRequestAsCreateCustomerPaymentMethodsRequest(v *PaymentMethodTokenRequest) CreateCustomerPaymentMethodsRequest {
	return CreateCustomerPaymentMethodsRequest{
		PaymentMethodTokenRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateCustomerPaymentMethodsRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaymentMethodCashRequest
	err = json.Unmarshal(data, &dst.PaymentMethodCashRequest)
	if err == nil {
		jsonPaymentMethodCashRequest, _ := json.Marshal(dst.PaymentMethodCashRequest)
		if string(jsonPaymentMethodCashRequest) == "{}" { // empty struct
			dst.PaymentMethodCashRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodCashRequest = nil
	}

	// try to unmarshal data into PaymentMethodSpeiRequest
	err = json.Unmarshal(data, &dst.PaymentMethodSpeiRequest)
	if err == nil {
		jsonPaymentMethodSpeiRequest, _ := json.Marshal(dst.PaymentMethodSpeiRequest)
		if string(jsonPaymentMethodSpeiRequest) == "{}" { // empty struct
			dst.PaymentMethodSpeiRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodSpeiRequest = nil
	}

	// try to unmarshal data into PaymentMethodTokenRequest
	err = json.Unmarshal(data, &dst.PaymentMethodTokenRequest)
	if err == nil {
		jsonPaymentMethodTokenRequest, _ := json.Marshal(dst.PaymentMethodTokenRequest)
		if string(jsonPaymentMethodTokenRequest) == "{}" { // empty struct
			dst.PaymentMethodTokenRequest = nil
		} else {
			match++
		}
	} else {
		dst.PaymentMethodTokenRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaymentMethodCashRequest = nil
		dst.PaymentMethodSpeiRequest = nil
		dst.PaymentMethodTokenRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateCustomerPaymentMethodsRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateCustomerPaymentMethodsRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateCustomerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodCashRequest != nil {
		return json.Marshal(&src.PaymentMethodCashRequest)
	}

	if src.PaymentMethodSpeiRequest != nil {
		return json.Marshal(&src.PaymentMethodSpeiRequest)
	}

	if src.PaymentMethodTokenRequest != nil {
		return json.Marshal(&src.PaymentMethodTokenRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateCustomerPaymentMethodsRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodCashRequest != nil {
		return obj.PaymentMethodCashRequest
	}

	if obj.PaymentMethodSpeiRequest != nil {
		return obj.PaymentMethodSpeiRequest
	}

	if obj.PaymentMethodTokenRequest != nil {
		return obj.PaymentMethodTokenRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateCustomerPaymentMethodsRequest struct {
	value *CreateCustomerPaymentMethodsRequest
	isSet bool
}

func (v NullableCreateCustomerPaymentMethodsRequest) Get() *CreateCustomerPaymentMethodsRequest {
	return v.value
}

func (v *NullableCreateCustomerPaymentMethodsRequest) Set(val *CreateCustomerPaymentMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerPaymentMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerPaymentMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerPaymentMethodsRequest(val *CreateCustomerPaymentMethodsRequest) *NullableCreateCustomerPaymentMethodsRequest {
	return &NullableCreateCustomerPaymentMethodsRequest{value: val, isSet: true}
}

func (v NullableCreateCustomerPaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerPaymentMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


