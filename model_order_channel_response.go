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
)

// checks if the OrderChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderChannelResponse{}

// OrderChannelResponse struct for OrderChannelResponse
type OrderChannelResponse struct {
	Segment *string `json:"segment,omitempty"`
	CheckoutRequestId *string `json:"checkout_request_id,omitempty"`
	CheckoutRequestType *string `json:"checkout_request_type,omitempty"`
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderChannelResponse OrderChannelResponse

// NewOrderChannelResponse instantiates a new OrderChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderChannelResponse() *OrderChannelResponse {
	this := OrderChannelResponse{}
	return &this
}

// NewOrderChannelResponseWithDefaults instantiates a new OrderChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderChannelResponseWithDefaults() *OrderChannelResponse {
	this := OrderChannelResponse{}
	return &this
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *OrderChannelResponse) GetSegment() string {
	if o == nil || IsNil(o.Segment) {
		var ret string
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChannelResponse) GetSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *OrderChannelResponse) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given string and assigns it to the Segment field.
func (o *OrderChannelResponse) SetSegment(v string) {
	o.Segment = &v
}

// GetCheckoutRequestId returns the CheckoutRequestId field value if set, zero value otherwise.
func (o *OrderChannelResponse) GetCheckoutRequestId() string {
	if o == nil || IsNil(o.CheckoutRequestId) {
		var ret string
		return ret
	}
	return *o.CheckoutRequestId
}

// GetCheckoutRequestIdOk returns a tuple with the CheckoutRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChannelResponse) GetCheckoutRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutRequestId) {
		return nil, false
	}
	return o.CheckoutRequestId, true
}

// HasCheckoutRequestId returns a boolean if a field has been set.
func (o *OrderChannelResponse) HasCheckoutRequestId() bool {
	if o != nil && !IsNil(o.CheckoutRequestId) {
		return true
	}

	return false
}

// SetCheckoutRequestId gets a reference to the given string and assigns it to the CheckoutRequestId field.
func (o *OrderChannelResponse) SetCheckoutRequestId(v string) {
	o.CheckoutRequestId = &v
}

// GetCheckoutRequestType returns the CheckoutRequestType field value if set, zero value otherwise.
func (o *OrderChannelResponse) GetCheckoutRequestType() string {
	if o == nil || IsNil(o.CheckoutRequestType) {
		var ret string
		return ret
	}
	return *o.CheckoutRequestType
}

// GetCheckoutRequestTypeOk returns a tuple with the CheckoutRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChannelResponse) GetCheckoutRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutRequestType) {
		return nil, false
	}
	return o.CheckoutRequestType, true
}

// HasCheckoutRequestType returns a boolean if a field has been set.
func (o *OrderChannelResponse) HasCheckoutRequestType() bool {
	if o != nil && !IsNil(o.CheckoutRequestType) {
		return true
	}

	return false
}

// SetCheckoutRequestType gets a reference to the given string and assigns it to the CheckoutRequestType field.
func (o *OrderChannelResponse) SetCheckoutRequestType(v string) {
	o.CheckoutRequestType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderChannelResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderChannelResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderChannelResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderChannelResponse) SetId(v string) {
	o.Id = &v
}

func (o OrderChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.CheckoutRequestId) {
		toSerialize["checkout_request_id"] = o.CheckoutRequestId
	}
	if !IsNil(o.CheckoutRequestType) {
		toSerialize["checkout_request_type"] = o.CheckoutRequestType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderChannelResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderChannelResponse := _OrderChannelResponse{}

	err = json.Unmarshal(data, &varOrderChannelResponse)

	if err != nil {
		return err
	}

	*o = OrderChannelResponse(varOrderChannelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "segment")
		delete(additionalProperties, "checkout_request_id")
		delete(additionalProperties, "checkout_request_type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderChannelResponse struct {
	value *OrderChannelResponse
	isSet bool
}

func (v NullableOrderChannelResponse) Get() *OrderChannelResponse {
	return v.value
}

func (v *NullableOrderChannelResponse) Set(val *OrderChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderChannelResponse(val *OrderChannelResponse) *NullableOrderChannelResponse {
	return &NullableOrderChannelResponse{value: val, isSet: true}
}

func (v NullableOrderChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


