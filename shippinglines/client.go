package shippinglines

import (
	conekta "github.com/conekta/conekta-go/v2"
)

// Create creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func Create(orderID string, p *conekta.ShippingLinesParams) (*conekta.ShippingLines, error) {
	sh := &conekta.ShippingLines{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/shipping_lines", p, sh)
	return sh, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(orderID string, id string, p *conekta.ShippingLinesParams) (*conekta.ShippingLines, error) {
	sh := &conekta.ShippingLines{}
	err := conekta.MakeRequest("PUT", "/orders/"+orderID+"/shipping_lines/"+id, p, sh)
	return sh, err
}

// Delete deletes a Order
func Delete(orderID string, id string) (*conekta.ShippingLines, error) {
	sh := &conekta.ShippingLines{}
	err := conekta.MakeRequest("DELETE", "/orders/"+orderID+"/shipping_lines/"+id, &conekta.OrderParams{}, sh)
	return sh, err
}
