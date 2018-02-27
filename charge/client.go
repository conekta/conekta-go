package charge

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api?#create-charge
func New(c *conekta.ChargeParams) (*conekta.Charge, *conekta.Error) {
	ch := &conekta.Charge{}
	err := conekta.New(ch, c, "/orders/"+c.ID+"/charges")
	return ch, err
}
