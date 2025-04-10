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

// checks if the ChargesDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargesDataResponse{}

// ChargesDataResponse struct for ChargesDataResponse
type ChargesDataResponse struct {
	Amount int32 `json:"amount"`
	Channel *ChargeResponseChannel `json:"channel,omitempty"`
	CreatedAt int64 `json:"created_at"`
	Currency string `json:"currency"`
	CustomerId *string `json:"customer_id,omitempty"`
	Description *string `json:"description,omitempty"`
	DeviceFingerprint *string `json:"device_fingerprint,omitempty"`
	FailureCode *string `json:"failure_code,omitempty"`
	FailureMessage *string `json:"failure_message,omitempty"`
	// Charge ID
	Id string `json:"id"`
	// Whether the charge was made in live mode or not
	Livemode bool `json:"livemode"`
	Object string `json:"object"`
	// Order ID
	OrderId string `json:"order_id"`
	// Payment date
	PaidAt NullableInt64 `json:"paid_at,omitempty"`
	PaymentMethod *ChargeResponsePaymentMethod `json:"payment_method,omitempty"`
	// Reference ID of the charge
	ReferenceId NullableString `json:"reference_id,omitempty"`
	Refunds NullableChargeResponseRefunds `json:"refunds,omitempty"`
	// Charge status
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _ChargesDataResponse ChargesDataResponse

// NewChargesDataResponse instantiates a new ChargesDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargesDataResponse(amount int32, createdAt int64, currency string, id string, livemode bool, object string, orderId string, status string) *ChargesDataResponse {
	this := ChargesDataResponse{}
	this.Amount = amount
	this.CreatedAt = createdAt
	this.Currency = currency
	this.Id = id
	this.Livemode = livemode
	this.Object = object
	this.OrderId = orderId
	this.Status = status
	return &this
}

// NewChargesDataResponseWithDefaults instantiates a new ChargesDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargesDataResponseWithDefaults() *ChargesDataResponse {
	this := ChargesDataResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ChargesDataResponse) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ChargesDataResponse) SetAmount(v int32) {
	o.Amount = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetChannel() ChargeResponseChannel {
	if o == nil || IsNil(o.Channel) {
		var ret ChargeResponseChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetChannelOk() (*ChargeResponseChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given ChargeResponseChannel and assigns it to the Channel field.
func (o *ChargesDataResponse) SetChannel(v ChargeResponseChannel) {
	o.Channel = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChargesDataResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChargesDataResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *ChargesDataResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ChargesDataResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ChargesDataResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChargesDataResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceFingerprint returns the DeviceFingerprint field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetDeviceFingerprint() string {
	if o == nil || IsNil(o.DeviceFingerprint) {
		var ret string
		return ret
	}
	return *o.DeviceFingerprint
}

// GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetDeviceFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceFingerprint) {
		return nil, false
	}
	return o.DeviceFingerprint, true
}

// HasDeviceFingerprint returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasDeviceFingerprint() bool {
	if o != nil && !IsNil(o.DeviceFingerprint) {
		return true
	}

	return false
}

// SetDeviceFingerprint gets a reference to the given string and assigns it to the DeviceFingerprint field.
func (o *ChargesDataResponse) SetDeviceFingerprint(v string) {
	o.DeviceFingerprint = &v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetFailureCode() string {
	if o == nil || IsNil(o.FailureCode) {
		var ret string
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetFailureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasFailureCode() bool {
	if o != nil && !IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given string and assigns it to the FailureCode field.
func (o *ChargesDataResponse) SetFailureCode(v string) {
	o.FailureCode = &v
}

// GetFailureMessage returns the FailureMessage field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetFailureMessage() string {
	if o == nil || IsNil(o.FailureMessage) {
		var ret string
		return ret
	}
	return *o.FailureMessage
}

// GetFailureMessageOk returns a tuple with the FailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FailureMessage) {
		return nil, false
	}
	return o.FailureMessage, true
}

// HasFailureMessage returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasFailureMessage() bool {
	if o != nil && !IsNil(o.FailureMessage) {
		return true
	}

	return false
}

// SetFailureMessage gets a reference to the given string and assigns it to the FailureMessage field.
func (o *ChargesDataResponse) SetFailureMessage(v string) {
	o.FailureMessage = &v
}

// GetId returns the Id field value
func (o *ChargesDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChargesDataResponse) SetId(v string) {
	o.Id = v
}

// GetLivemode returns the Livemode field value
func (o *ChargesDataResponse) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *ChargesDataResponse) SetLivemode(v bool) {
	o.Livemode = v
}

// GetObject returns the Object field value
func (o *ChargesDataResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ChargesDataResponse) SetObject(v string) {
	o.Object = v
}

// GetOrderId returns the OrderId field value
func (o *ChargesDataResponse) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *ChargesDataResponse) SetOrderId(v string) {
	o.OrderId = v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargesDataResponse) GetPaidAt() int64 {
	if o == nil || IsNil(o.PaidAt.Get()) {
		var ret int64
		return ret
	}
	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargesDataResponse) GetPaidAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// HasPaidAt returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasPaidAt() bool {
	if o != nil && o.PaidAt.IsSet() {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given NullableInt64 and assigns it to the PaidAt field.
func (o *ChargesDataResponse) SetPaidAt(v int64) {
	o.PaidAt.Set(&v)
}
// SetPaidAtNil sets the value for PaidAt to be an explicit nil
func (o *ChargesDataResponse) SetPaidAtNil() {
	o.PaidAt.Set(nil)
}

// UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
func (o *ChargesDataResponse) UnsetPaidAt() {
	o.PaidAt.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ChargesDataResponse) GetPaymentMethod() ChargeResponsePaymentMethod {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret ChargeResponsePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetPaymentMethodOk() (*ChargeResponsePaymentMethod, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given ChargeResponsePaymentMethod and assigns it to the PaymentMethod field.
func (o *ChargesDataResponse) SetPaymentMethod(v ChargeResponsePaymentMethod) {
	o.PaymentMethod = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargesDataResponse) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargesDataResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *ChargesDataResponse) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *ChargesDataResponse) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *ChargesDataResponse) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetRefunds returns the Refunds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargesDataResponse) GetRefunds() ChargeResponseRefunds {
	if o == nil || IsNil(o.Refunds.Get()) {
		var ret ChargeResponseRefunds
		return ret
	}
	return *o.Refunds.Get()
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargesDataResponse) GetRefundsOk() (*ChargeResponseRefunds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refunds.Get(), o.Refunds.IsSet()
}

// HasRefunds returns a boolean if a field has been set.
func (o *ChargesDataResponse) HasRefunds() bool {
	if o != nil && o.Refunds.IsSet() {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given NullableChargeResponseRefunds and assigns it to the Refunds field.
func (o *ChargesDataResponse) SetRefunds(v ChargeResponseRefunds) {
	o.Refunds.Set(&v)
}
// SetRefundsNil sets the value for Refunds to be an explicit nil
func (o *ChargesDataResponse) SetRefundsNil() {
	o.Refunds.Set(nil)
}

// UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
func (o *ChargesDataResponse) UnsetRefunds() {
	o.Refunds.Unset()
}

// GetStatus returns the Status field value
func (o *ChargesDataResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChargesDataResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChargesDataResponse) SetStatus(v string) {
	o.Status = v
}

func (o ChargesDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargesDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["currency"] = o.Currency
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeviceFingerprint) {
		toSerialize["device_fingerprint"] = o.DeviceFingerprint
	}
	if !IsNil(o.FailureCode) {
		toSerialize["failure_code"] = o.FailureCode
	}
	if !IsNil(o.FailureMessage) {
		toSerialize["failure_message"] = o.FailureMessage
	}
	toSerialize["id"] = o.Id
	toSerialize["livemode"] = o.Livemode
	toSerialize["object"] = o.Object
	toSerialize["order_id"] = o.OrderId
	if o.PaidAt.IsSet() {
		toSerialize["paid_at"] = o.PaidAt.Get()
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if o.Refunds.IsSet() {
		toSerialize["refunds"] = o.Refunds.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChargesDataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"created_at",
		"currency",
		"id",
		"livemode",
		"object",
		"order_id",
		"status",
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

	varChargesDataResponse := _ChargesDataResponse{}

	err = json.Unmarshal(data, &varChargesDataResponse)

	if err != nil {
		return err
	}

	*o = ChargesDataResponse(varChargesDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device_fingerprint")
		delete(additionalProperties, "failure_code")
		delete(additionalProperties, "failure_message")
		delete(additionalProperties, "id")
		delete(additionalProperties, "livemode")
		delete(additionalProperties, "object")
		delete(additionalProperties, "order_id")
		delete(additionalProperties, "paid_at")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "reference_id")
		delete(additionalProperties, "refunds")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChargesDataResponse struct {
	value *ChargesDataResponse
	isSet bool
}

func (v NullableChargesDataResponse) Get() *ChargesDataResponse {
	return v.value
}

func (v *NullableChargesDataResponse) Set(val *ChargesDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargesDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargesDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargesDataResponse(val *ChargesDataResponse) *NullableChargesDataResponse {
	return &NullableChargesDataResponse{value: val, isSet: true}
}

func (v NullableChargesDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargesDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


