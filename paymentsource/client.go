package paymentsource

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new payment source
// For details see https://developers.conekta.com/api#payment-source.
func Create(custID string, p *conekta.PaymentSourceCreateParams) (*conekta.PaymentSource, *conekta.Error) {
	ps := &conekta.PaymentSource{}
	err := conekta.MakeRequest("POST", "/customers/"+custID+"/payment_sources", p, ps)
	return ps, err
}

// Update updates a payment source
// For details see https://developers.conekta.com/api#update-shipping-contact
func Update(custID string, id string, p *conekta.PaymentSourceUpdateParams) (*conekta.PaymentSource, *conekta.Error) {
	ps := &conekta.PaymentSource{}
	err := conekta.MakeRequest("PUT", "/customers/"+custID+"/payment_sources/"+id, p, ps)
	return ps, err
}

// Delete deletes a payment source
// For details see https://developers.conekta.com/api#delete-payment-source.
func Delete(custID string, id string) (*conekta.PaymentSource, *conekta.Error) {
	ps := &conekta.PaymentSource{}
	err := conekta.MakeRequest("DELETE", "/customers/"+custID+"/payment_sources/"+id, &conekta.EmptyParams{}, ps)
	return ps, err
}

// Find gets a payment source
func Find(custID string, id string) (*conekta.PaymentSource, *conekta.Error) {
	ps := &conekta.PaymentSource{}
	err := conekta.MakeRequest("GET", "/customers/"+custID+"/payment_sources/"+id, &conekta.EmptyParams{}, ps)
	return ps, err
}

// All gets all payment sources from a customer
func All(custID string) (*conekta.PaymentSourceList, *conekta.Error) {
	ps := &conekta.PaymentSourceList{}
	err := conekta.MakeRequest("GET", "/customers/"+custID+"/payment_sources/", &conekta.EmptyParams{}, ps)
	return ps, err
}
