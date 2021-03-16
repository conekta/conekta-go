package checkout

import (
	"encoding/json"
	"fmt"

	conekta "github.com/conekta/conekta-go"
	"github.com/google/go-querystring/query"
)

type sendEmailParams struct {
	Email string `json:"email,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *sendEmailParams) Bytes() []byte {
	r, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return r
}

type sendSmsParams struct {
	Phone string `json:"sms,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *sendSmsParams) Bytes() []byte {
	r, err := json.Marshal(c)
	if err != nil {
		return []byte{}
	}
	return r
}

// Create creates a new checkout
func Create(p *conekta.CheckoutParams, customHeaders ...interface{}) (*conekta.Checkout, error) {
	checkout := &conekta.Checkout{}
	err := conekta.MakeRequest("POST", "/checkouts", p, checkout, customHeaders...)
	return checkout, err
}

// Cancel cancels only a oxxo Checkout
func Cancel(checkoutID string) (*conekta.Checkout, error) {
	checkout := &conekta.Checkout{}
	err := conekta.MakeRequest("PUT", fmt.Sprintf("/checkouts/%v/cancel", checkoutID), &conekta.EmptyParams{}, checkout)
	return checkout, err
}

// Find gets a checkout by id
func Find(id string) (*conekta.Checkout, error) {
	checkout := &conekta.Checkout{}
	err := conekta.MakeRequest("GET", fmt.Sprintf("/checkouts/%v", id), &conekta.EmptyParams{}, checkout)
	return checkout, err
}

// Where gets a checkout by querying parameter
func Where(params interface{}) (*conekta.CheckoutList, error) {
	checkout := &conekta.CheckoutList{}
	v, _ := query.Values(params)
	err := conekta.MakeRequest("GET", fmt.Sprintf("/checkouts?%v", v.Encode()), &conekta.EmptyParams{}, checkout)
	return checkout, err
}

// Sends an email
func SendEmail(checkoutID string, email string) (*conekta.Checkout, error) {
	checkout := &conekta.Checkout{}

	err := conekta.MakeRequest("POST", fmt.Sprintf("/checkouts/%v/email", checkoutID), &sendEmailParams{Email: email}, checkout)
	return checkout, err
}

// Sends an SMS
func SendSms(checkoutID string, phone string) (*conekta.Checkout, error) {
	checkout := &conekta.Checkout{}
	err := conekta.MakeRequest("POST", fmt.Sprintf("/checkouts/%v/sms", checkoutID), &sendSmsParams{Phone: phone}, checkout)
	return checkout, err
}
