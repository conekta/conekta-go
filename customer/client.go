package customer

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.CustomerParams) (*conekta.Customer, *conekta.Error) {
	cust := &conekta.Customer{}
	err := conekta.New(cust, p, "/customers")
	return cust, err
}
