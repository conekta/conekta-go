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

// checks if the PayoutOrderPayoutsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayoutOrderPayoutsItem{}

// PayoutOrderPayoutsItem struct for PayoutOrderPayoutsItem
type PayoutOrderPayoutsItem struct {
	// The amount of the payout.
	Amount int32 `json:"amount"`
	// The currency in which the payout is made.
	Currency string `json:"currency"`
	// The expiration date of the payout.
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// The id of the payout.
	Id string `json:"id"`
	// The live mode of the payout.
	Livemode bool `json:"livemode"`
	// The object of the payout.
	Object string `json:"object"`
	// The id of the payout order.
	PayoutOrderId *string `json:"payout_order_id,omitempty"`
	// The status of the payout.
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PayoutOrderPayoutsItem PayoutOrderPayoutsItem

// NewPayoutOrderPayoutsItem instantiates a new PayoutOrderPayoutsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutOrderPayoutsItem(amount int32, currency string, id string, livemode bool, object string) *PayoutOrderPayoutsItem {
	this := PayoutOrderPayoutsItem{}
	this.Amount = amount
	this.Currency = currency
	this.Id = id
	this.Livemode = livemode
	this.Object = object
	return &this
}

// NewPayoutOrderPayoutsItemWithDefaults instantiates a new PayoutOrderPayoutsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutOrderPayoutsItemWithDefaults() *PayoutOrderPayoutsItem {
	this := PayoutOrderPayoutsItem{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PayoutOrderPayoutsItem) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PayoutOrderPayoutsItem) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *PayoutOrderPayoutsItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PayoutOrderPayoutsItem) SetCurrency(v string) {
	o.Currency = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PayoutOrderPayoutsItem) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PayoutOrderPayoutsItem) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PayoutOrderPayoutsItem) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *PayoutOrderPayoutsItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PayoutOrderPayoutsItem) SetId(v string) {
	o.Id = v
}

// GetLivemode returns the Livemode field value
func (o *PayoutOrderPayoutsItem) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *PayoutOrderPayoutsItem) SetLivemode(v bool) {
	o.Livemode = v
}

// GetObject returns the Object field value
func (o *PayoutOrderPayoutsItem) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PayoutOrderPayoutsItem) SetObject(v string) {
	o.Object = v
}

// GetPayoutOrderId returns the PayoutOrderId field value if set, zero value otherwise.
func (o *PayoutOrderPayoutsItem) GetPayoutOrderId() string {
	if o == nil || IsNil(o.PayoutOrderId) {
		var ret string
		return ret
	}
	return *o.PayoutOrderId
}

// GetPayoutOrderIdOk returns a tuple with the PayoutOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetPayoutOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayoutOrderId) {
		return nil, false
	}
	return o.PayoutOrderId, true
}

// HasPayoutOrderId returns a boolean if a field has been set.
func (o *PayoutOrderPayoutsItem) HasPayoutOrderId() bool {
	if o != nil && !IsNil(o.PayoutOrderId) {
		return true
	}

	return false
}

// SetPayoutOrderId gets a reference to the given string and assigns it to the PayoutOrderId field.
func (o *PayoutOrderPayoutsItem) SetPayoutOrderId(v string) {
	o.PayoutOrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PayoutOrderPayoutsItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutOrderPayoutsItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PayoutOrderPayoutsItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PayoutOrderPayoutsItem) SetStatus(v string) {
	o.Status = &v
}

func (o PayoutOrderPayoutsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutOrderPayoutsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	toSerialize["livemode"] = o.Livemode
	toSerialize["object"] = o.Object
	if !IsNil(o.PayoutOrderId) {
		toSerialize["payout_order_id"] = o.PayoutOrderId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PayoutOrderPayoutsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"id",
		"livemode",
		"object",
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

	varPayoutOrderPayoutsItem := _PayoutOrderPayoutsItem{}

	err = json.Unmarshal(data, &varPayoutOrderPayoutsItem)

	if err != nil {
		return err
	}

	*o = PayoutOrderPayoutsItem(varPayoutOrderPayoutsItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "livemode")
		delete(additionalProperties, "object")
		delete(additionalProperties, "payout_order_id")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayoutOrderPayoutsItem struct {
	value *PayoutOrderPayoutsItem
	isSet bool
}

func (v NullablePayoutOrderPayoutsItem) Get() *PayoutOrderPayoutsItem {
	return v.value
}

func (v *NullablePayoutOrderPayoutsItem) Set(val *PayoutOrderPayoutsItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutOrderPayoutsItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutOrderPayoutsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutOrderPayoutsItem(val *PayoutOrderPayoutsItem) *NullablePayoutOrderPayoutsItem {
	return &NullablePayoutOrderPayoutsItem{value: val, isSet: true}
}

func (v NullablePayoutOrderPayoutsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutOrderPayoutsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


