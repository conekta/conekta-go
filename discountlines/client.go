package discountlines

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.DiscountLinesParams) (*conekta.DiscountLines, *conekta.Error) {
	dl := &conekta.DiscountLines{}
	err := conekta.New(dl, p, "/orders/"+p.ID+"/discount_lines")
	return dl, err
}
