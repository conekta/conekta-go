package shippinglines

import (
	conekta "github.com/conekta/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(orderID string, p *conekta.ShippingLinesParams) (*conekta.ShippingLines, *conekta.Error) {
	sh := &conekta.ShippingLines{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/shipping_lines", p, sh)
	return sh, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(orderID string, id string, p *conekta.OrderParams) (*conekta.Order, *conekta.Error) {
	sh := &conekta.Order{}
	err := conekta.MakeRequest("PUT", "/orders/"+orderID+"/shipping_lines/"+id, p, sh)
	return sh, err
}

// Delete deletes a Order
func Delete(orderID string, id string) (*conekta.Order, *conekta.Error) {
	sh := &conekta.Order{}
	err := conekta.MakeRequest("DELETE", "/orders/"+orderID+"/shipping_lines/"+id, &conekta.OrderParams{}, sh)
	return sh, err
}
