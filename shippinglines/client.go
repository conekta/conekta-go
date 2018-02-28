package shippinglines

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.ShippingLinesParams) (*conekta.ShippingLines, *conekta.Error) {
	sh := &conekta.ShippingLines{}
	err := conekta.New(sh, p, "/orders/"+p.ID+"/shipping_lines")
	return sh, err
}
