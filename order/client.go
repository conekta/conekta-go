package order

import (
	"encoding/json"

	conekta "github.com/conekta/conekta-go"
	"github.com/google/go-querystring/query"
)

// Create creates a new order
func Create(p *conekta.OrderParams, customHeaders ...interface{}) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders", p, ord, customHeaders...)

	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, ord)
		return ord, nil
	}
	return ord, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(id string, p *conekta.OrderParams) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("PUT", "/orders/"+id, p, ord)
	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, ord)
		return ord, nil
	}
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

// Where gets a order by querying parameter
func Where(params interface{}) (*conekta.OrderList, error) {
	ord := &conekta.OrderList{}
	v, _ := query.Values(params)
	err := conekta.MakeRequest("GET", "/orders?"+v.Encode(), &conekta.OrderParams{}, ord)
	return ord, err
}

// Refund refund an order
func Refund(orderID string, params *conekta.OrderRefundParams) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/refund", params, ord)
	return ord, err
}

// Void voids an order
func Void(orderID string) (*conekta.Order, error) {
	ord := &conekta.Order{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/void", &conekta.EmptyParams{}, ord)
	return ord, err
}
