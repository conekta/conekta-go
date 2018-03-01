package order

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new order
func Create(p *conekta.OrderParams) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders", p, ord)
	return ord, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(id string, p *conekta.OrderParams) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("PUT", "/orders/"+id, p, ord)
	return ord, err
}

// Delete deletes a Order
func Delete(id string) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("DELETE", "/orders/"+id, &conekta.OrderParams{}, ord)
	return ord, err
}

// Find gets a ordomer by id
func Find(id string) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("GET", "/orders/"+id, &conekta.OrderParams{}, ord)
	return ord, err
}

//CreateDiscountLine create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func CreateDiscountLine(p *conekta.DiscountLines) (*conekta.Order, *conekta.Error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders/"+p.ID+"/discount_lines", p, ord)
	return ord, err
}

// UpdateDiscountLine updates a DiscountLine inside order
// For details see https://developers.conekta.com/api#update-discount-line
func UpdateDiscountLine(id string, p *conekta.DiscountLines) (*conekta.DiscountLines, *conekta.Error) {
	ord := &conekta.DiscountLines{}
	err := conekta.MakeRequest("PUT", "/orders/"+id+"/discount_lines", p, ord)
	return ord, err
}

// DeleteDiscountLine deletes a Order Discount Line
func DeleteDiscountLine(id string) (*conekta.DiscountLines, *conekta.Error) {
	ord := &conekta.DiscountLines{}
	err := conekta.MakeRequest("DELETE", "/orders/"+id, &conekta.OrderParams{}, ord)
	return ord, err
}
