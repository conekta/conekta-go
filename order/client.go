package order

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new order
func Create(p *conekta.OrderParams, customHeaders ...interface{}) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders", p, ord, customHeaders...)
	return ord, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(id string, p *conekta.OrderParams) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("PUT", "/orders/"+id, p, ord)
	return ord, err
}

// Capture deletes a Order
func Capture(id string) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders/"+id+"/capture", &conekta.OrderParams{}, ord)
	return ord, err
}

// Cancel cancels only a oxxo Order
func Cancel(orderID string) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/cancel", &conekta.OrderParams{}, ord)
	return ord, err
}

// Find gets a order by id
func Find(id string) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("GET", "/orders/"+id, &conekta.OrderParams{}, ord)
	return ord, err
}
