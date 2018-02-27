package order

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.OrderParams) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.New(ord, p, "/orders")
	return ord, err
}
